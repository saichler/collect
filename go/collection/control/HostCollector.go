package control

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
	"sync"
	"time"
)

type HostCollector struct {
	resources          common.IResources
	jobCompleteHandler base.IJobCompleteHandler
	deviceId           string
	hostId             string
	iService           *types.DeviceServiceInfo
	pService           *types.DeviceServiceInfo
	collectors         map[int32]base.ProtocolCollector
	jobsQueue          *JobsQueue
	mtx                *sync.Mutex
	running            bool
}

func newHostCollector(hoistId string, device *types.DeviceConfig, resources common.IResources, handler base.IJobCompleteHandler) *HostCollector {
	hc := &HostCollector{}
	hc.deviceId = device.DeviceId
	hc.hostId = hoistId
	hc.collectors = make(map[int32]base.ProtocolCollector)
	hc.resources = resources
	hc.jobCompleteHandler = handler

	hc.iService = device.InventoryService
	hc.pService = device.ParsingService

	hc.jobsQueue = NewJobsQueue(device.DeviceId, hoistId, hc.resources, device.InventoryService, device.ParsingService)
	hc.mtx = &sync.Mutex{}
	hc.running = true

	return hc
}

func (this *HostCollector) update() error {
	cc := device_config.Configs(this.resources)
	configs := cc.HostConnectionConfigs(this.deviceId, this.hostId)
	for _, config := range configs {
		this.mtx.Lock()
		_, exist := this.collectors[int32(config.Protocol)]
		this.mtx.Unlock()

		if !exist {
			col, err := newProtocolCollector(config, this.resources)
			if err != nil {
				this.resources.Logger().Error(err)
			}
			if col != nil {
				this.mtx.Lock()
				this.collectors[int32(config.Protocol)] = col
				this.mtx.Unlock()
			}
		}
	}

	pc := poll_config.PollConfig(this.resources)
	bootPollList := pc.PollsByGroup(boot.BOOT_GROUP, "", "", "", "", "", "")
	for _, pollName := range bootPollList {
		this.jobsQueue.InsertJob(pollName.Name, "", "", "", "", "", "", 0, 0)
	}

	return nil
}

func (this *HostCollector) stop() {
	this.running = false
	this.mtx.Lock()
	defer this.mtx.Unlock()
	for _, collector := range this.collectors {
		collector.Disconnect()
	}
	this.collectors = nil
	this.jobsQueue.Shutdown()
}

func (this *HostCollector) start() error {
	cc := device_config.Configs(this.resources)
	configs := cc.HostConnectionConfigs(this.deviceId, this.hostId)
	for _, config := range configs {
		col, err := newProtocolCollector(config, this.resources)
		if err != nil {
			this.resources.Logger().Error(err)
		}
		if col != nil {
			this.mtx.Lock()
			this.collectors[int32(config.Protocol)] = col
			this.mtx.Unlock()
		}
	}

	pc := poll_config.PollConfig(this.resources)
	bootPollList := pc.PollsByGroup(boot.BOOT_GROUP, "", "", "", "", "", "")
	for _, pollName := range bootPollList {
		this.jobsQueue.InsertJob(pollName.Name, "", "", "", "", "", "", 0, 0)
	}

	go this.collect()

	return nil
}

func (this *HostCollector) collect() {
	this.resources.Logger().Info("** Starting Collection on host ", this.hostId)
	pc := poll_config.PollConfig(this.resources)
	for this.running {
		job, waitTime := this.jobsQueue.Pop()
		if job != nil {
			poll := pc.PollByName(job.PollName)
			if poll == nil {
				this.resources.Logger().Error("cannot find poll for uuid ", job.PollName)
				continue
			}
			MarkStart(job)
			this.mtx.Lock()
			col, ok := this.collectors[int32(poll.Protocol)]
			this.mtx.Unlock()
			if !ok {
				MarkEnded(job)
				continue
			}
			col.Exec(job)
			MarkEnded(job)
			if this.running {
				this.jobCompleteHandler.JobCompleted(job)
			}
		} else {
			time.Sleep(time.Second * time.Duration(waitTime))
		}
	}
	this.resources.Logger().Info("Host collection for device ", this.deviceId, " host ", this.hostId, " has ended.")
	this.resources = nil
}
