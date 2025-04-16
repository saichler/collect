package device_config

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/collector"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

const (
	ServiceName      = "DeviceConfig"
	ServicePointType = "DeviceConfigServicePoint"
)

type DeviceConfigServicePoint struct {
	configCenter *DeviceConfigCenter
	controller   base.IController
}

func (this *DeviceConfigServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.DeviceConfig{})
	this.configCenter = newConfigCenter(ServiceName, serviceArea, r, l)
	if args == nil {
		vnic, ok := l.(common.IVirtualNetworkInterface)
		if ok {
			pt := collector.NewParsingCenterNotifier(vnic)
			this.controller = collector.NewDeviceCollector(pt, r)
		}
	} else {
		this.controller, _ = args[0].(base.IController)
	}
	return nil
}

func (this *DeviceConfigServicePoint) DeActivate() error {
	this.controller.Shutdown()
	this.configCenter.Shutdown()
	this.controller = nil
	this.configCenter = nil
	return nil
}

func (this *DeviceConfigServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	device := pb.Element().(*types.DeviceConfig)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.DeviceId)
		this.controller.StartPolling(device)
	}
	return nil
}
func (this *DeviceConfigServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *DeviceConfigServicePoint) Transactional() bool { return false }
func (this *DeviceConfigServicePoint) ReplicationCount() int {
	return 0
}
func (this *DeviceConfigServicePoint) ReplicationScore() int {
	return 0
}
