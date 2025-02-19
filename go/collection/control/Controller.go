package control

import (
	"errors"
	"github.com/saichler/collect/go/collection/common"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/protocols/snmp"
	"github.com/saichler/collect/go/collection/protocols/ssh"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/share/strings"
	"sync"
)

type Controller struct {
	hcollectors         map[string]*HostCollector
	mtx                 *sync.Mutex
	notificationHandler common.CollectNotificationHandler
	resources           interfaces.IResources
}

func NewController(handler common.CollectNotificationHandler, resources interfaces.IResources) *Controller {
	controller := &Controller{}
	controller.resources = resources
	controller.hcollectors = make(map[string]*HostCollector)
	controller.mtx = &sync.Mutex{}
	controller.notificationHandler = handler
	resources.Registry().Register(&types.Map{})
	resources.Registry().Register(&types.Table{})
	return controller
}

func newProtocolCollector(config *types.Config, resource interfaces.IResources) (common.ProtocolCollector, error) {
	var protocolCollector common.ProtocolCollector
	if config.Protocol == types.Protocol_SSH {
		protocolCollector = &ssh.SshCollector{}
	} else if config.Protocol == types.Protocol_SNMPV2 {
		protocolCollector = &snmp.SNMPCollector{}
	} else {
		return nil, errors.New("Unknown Protocol " + config.Protocol.String())
	}
	err := protocolCollector.Init(config, resource)
	return protocolCollector, err
}

func (this *Controller) StartPolling(deviceId string) error {
	cc := config.Configs(this.resources)
	device := cc.DeviceById(deviceId)
	if device == nil {
		return errors.New("device with id " + deviceId + " does not exist")
	}
	for _, host := range device.Hosts {
		hostCol, _ := this.hostCollector(deviceId, host.Id)
		hostCol.start()
	}
	return nil
}

func hcKey(deviceId, hostId string) string {
	return strings.New(deviceId, hostId).String()
}

func (this *Controller) hostCollector(deviceId, hostId string) (*HostCollector, bool) {
	key := hcKey(deviceId, hostId)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	hc, ok := this.hcollectors[key]
	if ok {
		return hc, ok
	}
	hc = newHostCollector(deviceId, hostId, this)
	this.hcollectors[key] = hc
	return hc, ok
}

func (this *Controller) jobComplete(job *types.Job) {
	/*
		pc := polling.Polling(this.resources)
		poll := pc.PollByUuid(job.PollUuid)
		if poll == nil {
			this.resources.Logger().Error("cannot find poll for uuid ", job.PollUuid)
			return
		}
		if job.Error == "" && poll.Attributes != nil {
			inv := inventory.Inventory(this.resources)
			box := inv.BoxById(job.DeviceId)
			if box == nil {
				box = &types.NetworkBox{}
				box.Id = job.DeviceId
				inv.Add(box)
			}
			parsing.Parser.Parse(job, box, this.resources)
			inv.Update(box)
		}*/
	if this.notificationHandler != nil {
		this.notificationHandler.HandleCollectNotification(job)
	}
}
