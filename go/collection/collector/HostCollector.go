package collector

import (
	"errors"
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"github.com/saichler/collect/go/collection/protocols/k8s"
	"github.com/saichler/collect/go/collection/protocols/snmp"
	"github.com/saichler/collect/go/collection/protocols/ssh"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8utils/go/utils/maps"
	"github.com/saichler/l8types/go/ifs"
	"time"
)

type HostCollector struct {
	resources    ifs.IResources
	deviceConfig *types.DeviceConfig
	hostId       string
	handler      base.IJobCompleteHandler
	collectors   *maps.SyncMap
	jobsQueue    *JobsQueue
	running      bool
}

func newHostCollector(deviceConfig *types.DeviceConfig, hostId string, resources ifs.IResources, handler base.IJobCompleteHandler) *HostCollector {
	hc := &HostCollector{}
	hc.deviceConfig = deviceConfig
	hc.hostId = hostId
	hc.collectors = maps.NewSyncMap()
	hc.resources = resources
	hc.handler = handler
	hc.jobsQueue = NewJobsQueue(deviceConfig.DeviceId, hostId, hc.resources, deviceConfig.InventoryService, deviceConfig.ParsingService)
	hc.running = true
	return hc
}

func (this *HostCollector) update() error {
	host := this.deviceConfig.Hosts[this.hostId]
	for _, config := range host.Configs {
		exist := this.collectors.Contains(config.Protocol)
		if !exist {
			col, err := newProtocolCollector(config, this.resources)
			if err != nil {
				return this.resources.Logger().Error(err)
			}
			if col != nil {
				this.collectors.Put(config.Protocol, col)
			}
		}
	}

	pollCenter := poll_config.PollConfig(this.resources)
	bootPollList := pollCenter.PollsByGroup(boot.BOOT_GROUP, "", "", "", "", "", "")
	for _, pollName := range bootPollList {
		err := this.jobsQueue.InsertJob(pollName.Name, "", "", "", "", "", "", 0, 0)
		if err != nil {
			this.resources.Logger().Error(err)
		}
	}

	return nil
}

func (this *HostCollector) stop() {
	this.running = false
	this.collectors.Iterate(func(k, v interface{}) {
		c := v.(base.ProtocolCollector)
		c.Disconnect()
	})
	this.collectors = nil
	this.jobsQueue.Shutdown()
}

func (this *HostCollector) start() error {
	host := this.deviceConfig.Hosts[this.hostId]
	for _, config := range host.Configs {
		col, err := newProtocolCollector(config, this.resources)
		if err != nil {
			this.resources.Logger().Error(err)
		}
		if col != nil {
			this.collectors.Put(config.Protocol, col)
		}
	}

	pollCenter := poll_config.PollConfig(this.resources)
	bootPollList := pollCenter.PollsByGroup(boot.BOOT_GROUP, "", "", "", "", "", "")
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
			this.resources.Logger().Info("Poped job ", job.PollName)
		} else {
			this.resources.Logger().Info("No Job, waitTime ", waitTime)
		}
		if job != nil {
			poll := pc.PollByName(job.PollName)
			if poll == nil {
				this.resources.Logger().Error("cannot find poll for uuid ", job.PollName)
				continue
			}
			MarkStart(job)
			c, ok := this.collectors.Get(poll.Protocol)
			if !ok {
				MarkEnded(job)
				continue
			}
			c.(base.ProtocolCollector).Exec(job)
			MarkEnded(job)
			if this.running {
				this.handler.JobCompleted(job)
			}
		} else {
			this.resources.Logger().Info("No more jobs, next job in ", waitTime, " seconds.")
			time.Sleep(time.Second * time.Duration(waitTime))
		}
	}
	this.resources.Logger().Info("Host collection for device ", this.deviceConfig.DeviceId, " host ", this.hostId, " has ended.")
	this.resources = nil
}

func newProtocolCollector(config *types.ConnectionConfig, resource ifs.IResources) (base.ProtocolCollector, error) {
	var protocolCollector base.ProtocolCollector
	if config.Protocol == types.Protocol_SSH {
		protocolCollector = &ssh.SshCollector{}
	} else if config.Protocol == types.Protocol_SNMPV2 {
		protocolCollector = &snmp.SNMPCollector{}
	} else if config.Protocol == types.Protocol_K8s {
		protocolCollector = &k8s.Kubernetes{}
	} else {
		return nil, errors.New("Unknown Protocol " + config.Protocol.String())
	}
	err := protocolCollector.Init(config, resource)
	return protocolCollector, err
}
