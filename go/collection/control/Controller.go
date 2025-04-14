package control

import (
	"errors"
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/protocols/k8s"
	"github.com/saichler/collect/go/collection/protocols/snmp"
	"github.com/saichler/collect/go/collection/protocols/ssh"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/strings"
	"github.com/saichler/types/go/common"
	"sync"
)

type Controller struct {
	hcollectors        map[string]*HostCollector
	mtx                *sync.Mutex
	jobCompleteHandler base.IJobCompleteHandler
	resources          common.IResources
}

func NewController(handler base.IJobCompleteHandler, resources common.IResources) *Controller {
	resources.Logger().Debug("*** Creating new controller for vnet ", resources.SysConfig().VnetPort)
	controller := &Controller{}
	controller.resources = resources
	controller.hcollectors = make(map[string]*HostCollector)
	controller.mtx = &sync.Mutex{}
	controller.jobCompleteHandler = handler
	resources.Registry().Register(&types.CMap{})
	resources.Registry().Register(&types.CTable{})
	return controller
}

func newProtocolCollector(config *types.ConnectionConfig, resource common.IResources) (base.ProtocolCollector, error) {
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

func (this *Controller) StartPolling(device *types.DeviceConfig) error {
	for _, host := range device.Hosts {
		hostCol, _ := this.hostCollector(host.DeviceId, device)
		hostCol.start()
	}
	return nil
}

func (this *Controller) Shutdown() {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	for _, h := range this.hcollectors {
		h.stop()
	}
	this.hcollectors = nil
}

func hcKey(deviceId, hostId string) string {
	return strings.New(deviceId, hostId).String()
}

func (this *Controller) hostCollector(hostId string, device *types.DeviceConfig) (*HostCollector, bool) {
	key := hcKey(device.DeviceId, hostId)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	hc, ok := this.hcollectors[key]
	if ok {
		return hc, ok
	}
	hc = newHostCollector(device, hostId, this.resources, this.jobCompleteHandler)
	this.hcollectors[key] = hc
	return hc, ok
}
