package control

import (
	"errors"
	common2 "github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/device_config"
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
	serviceArea         uint16
}

func NewController(handler common2.CollectNotificationHandler, resources common.IResources, serviceArea uint16) *Controller {
	resources.Logger().Debug("*** Creating new controller for vnet ", resources.SysConfig().VnetPort)
	controller := &Controller{}
	controller.resources = resources
	controller.hcollectors = make(map[string]*HostCollector)
	controller.mtx = &sync.Mutex{}
	controller.notificationHandler = handler
	resources.Registry().Register(&types.CMap{})
	resources.Registry().Register(&types.CTable{})
	controller.serviceArea = serviceArea
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

func (this *Controller) StartPolling(deviceId, serviceName string) error {
	cc := deviceconfig.Configs(this.resources, this.serviceArea)
	device := cc.DeviceById(deviceId)
	if device == nil {
		return errors.New("device with id " + deviceId + " does not exist")
	}
	for _, host := range device.Hosts {
		hostCol, _ := this.hostCollector(deviceId, host.Id, serviceName,
			this.serviceArea, uint16(device.ServiceArea))
		hostCol.start()
	}
	return nil
}

func hcKey(deviceId, hostId string) string {
	return strings.New(deviceId, hostId).String()
}

func (this *Controller) hostCollector(deviceId, hostId, serviceName string, cServiceArea, dServiceArea uint16) (*HostCollector, bool) {
	key := hcKey(deviceId, hostId)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	hc, ok := this.hcollectors[key]
	if ok {
		return hc, ok
	}
	hc = newHostCollector(deviceId, hostId, serviceName, cServiceArea, dServiceArea, this)
	this.hcollectors[key] = hc
	return hc, ok
}

func (this *Controller) jobComplete(job *types.Job) {
	this.resources.Logger().Debug("Job Complete For ", job.DeviceId, " ", job.PollName)
	if this.notificationHandler != nil {
		this.notificationHandler.HandleCollectNotification(job)
	}
}
