package control

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/collection/polling/boot"
	"sync"
	"time"
)

type HostCollector struct {
	controller   *Controller
	deviceId     string
	hostId       string
	serviceName  string
	cServiceArea int32
	dServiceArea int32
	collectors   map[int32]base.ProtocolCollector
	jobsQueue    *JobsQueue
	mtx          *sync.Mutex
	running      bool
}

func newHostCollector(deviceId, hoistId, serviceName string, cServiceArea, dServiceArea int32, controller *Controller) *HostCollector {
	hc := &HostCollector{}
	hc.deviceId = deviceId
	hc.hostId = hoistId
	hc.collectors = make(map[int32]base.ProtocolCollector)
	hc.controller = controller
	hc.jobsQueue = NewJobsQueue(deviceId, hoistId, controller.resources, serviceName, cServiceArea, dServiceArea)
	hc.mtx = &sync.Mutex{}
	hc.running = true
	hc.cServiceArea = cServiceArea
	hc.dServiceArea = dServiceArea
	hc.serviceName = serviceName
	return hc
}

func (this *HostCollector) update() error {
	cc := config.Configs(this.controller.resources, this.cServiceArea)
	configs := cc.HostConfigs(this.deviceId, this.hostId)
	for _, config := range configs {
		this.mtx.Lock()
		_, exist := this.collectors[int32(config.Protocol)]
		this.mtx.Unlock()

		if !exist {
			col, err := newProtocolCollector(config, this.controller.resources)
			if err != nil {
				this.controller.resources.Logger().Error(err)
			}
			if col != nil {
				this.mtx.Lock()
				this.collectors[int32(config.Protocol)] = col
				this.mtx.Unlock()
			}
		}
	}

	pc := polling.Polling(this.controller.resources, this.cServiceArea)
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
}

func (this *HostCollector) start() error {
	cc := config.Configs(this.controller.resources, this.cServiceArea)
	configs := cc.HostConfigs(this.deviceId, this.hostId)
	for _, config := range configs {
		col, err := newProtocolCollector(config, this.controller.resources)
		if err != nil {
			this.controller.resources.Logger().Error(err)
		}
		if col != nil {
			this.mtx.Lock()
			this.collectors[int32(config.Protocol)] = col
			this.mtx.Unlock()
		}
	}

	pc := polling.Polling(this.controller.resources, this.cServiceArea)
	bootPollList := pc.PollsByGroup(boot.BOOT_GROUP, "", "", "", "", "", "")
	for _, pollName := range bootPollList {
		this.jobsQueue.InsertJob(pollName.Name, "", "", "", "", "", "", 0, 0)
	}

	go this.collect()

	return nil
}

func (this *HostCollector) collect() {
	this.controller.resources.Logger().Info("** Starting Collection on host ", this.hostId)
	pc := polling.Polling(this.controller.resources, this.cServiceArea)
	for this.running {
		job, waitTime := this.jobsQueue.Pop()
		if job != nil {
			poll := pc.PollByName(job.PollName)
			if poll == nil {
				this.controller.resources.Logger().Error("cannot find poll for uuid ", job.PollName)
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
			this.controller.jobComplete(job)
		} else {
			time.Sleep(time.Second * time.Duration(waitTime))
		}
	}
	this.controller.resources.Logger().Info("Host collection for device ", this.deviceId, " host ", this.hostId, " has ended.")
}
