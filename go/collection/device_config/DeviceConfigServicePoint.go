package device_config

import (
	"github.com/saichler/collect/go/collection/base"
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

func (this DeviceConfigServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.Device{})
	this.controller = args[0].(base.IController)
	this.configCenter = newConfigCenter(serviceName, serviceArea, r, l)
	return nil
}

func (this *DeviceConfigServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	device := pb.Element().(*types.Device)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.Id, " ", device.ServiceName)
		this.controller.StartPolling(device.Id, device.ServiceName)
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
