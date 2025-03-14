package control

import (
	"errors"
	common2 "github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/protocols/k8s"
	"github.com/saichler/collect/go/collection/protocols/snmp"
	"github.com/saichler/collect/go/collection/protocols/ssh"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/strings"
	"github.com/saichler/types/go/common"
	"sync"
)

type Controller struct {
	hcollectors         map[string]*HostCollector
	mtx                 *sync.Mutex
	notificationHandler common2.CollectNotificationHandler
	resources           common.IResources
}

func NewController(handler common2.CollectNotificationHandler, resources common.IResources) *Controller {
	resources.Logger().Debug("*** Creating new controller for vnet ", resources.Config().VnetPort)
	controller := &Controller{}
	controller.resources = resources
	controller.hcollectors = make(map[string]*HostCollector)
	controller.mtx = &sync.Mutex{}
	controller.notificationHandler = handler
	resources.Registry().Register(&types.Map{})
	resources.Registry().Register(&types.Table{})
	return controller
}

func newProtocolCollector(config *types.Config, resource common.IResources) (common2.ProtocolCollector, error) {
	var protocolCollector common2.ProtocolCollector
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

func (this *Controller) StartPolling(deviceId string, area int32) error {
	cc := config.Configs(this.resources)
	device := cc.DeviceById(deviceId)
	if device == nil {
		return errors.New("device with id " + deviceId + " does not exist")
	}
	for _, host := range device.Hosts {
		hostCol, _ := this.hostCollector(deviceId, host.Id, area)
		hostCol.start()
	}
	return nil
}

func hcKey(deviceId, hostId string) string {
	return strings.New(deviceId, hostId).String()
}

func (this *Controller) hostCollector(deviceId, hostId string, area int32) (*HostCollector, bool) {
	key := hcKey(deviceId, hostId)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	hc, ok := this.hcollectors[key]
	if ok {
		return hc, ok
	}
	hc = newHostCollector(deviceId, hostId, area, this)
	this.hcollectors[key] = hc
	return hc, ok
}

func (this *Controller) jobComplete(job *types.Job, area int32) {
	if this.notificationHandler != nil {
		this.notificationHandler.HandleCollectNotification(job, area)
	}
}
