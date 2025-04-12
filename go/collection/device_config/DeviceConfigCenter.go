package device_config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
)

type DeviceConfigCenter struct {
	devices *cache.Cache
}

func newConfigCenter(serviceName string, serviceArea uint16, resources common.IResources, listener common.IServicePointCacheListener) *DeviceConfigCenter {
	this := &DeviceConfigCenter{}
	this.devices = cache.NewModelCache(serviceName, serviceArea, "Device", resources.SysConfig().LocalUuid, listener, resources.Introspector())
	return this
}

func (this *DeviceConfigCenter) Add(device *types.Device) {
	this.devices.Put(device.Id, device)
}

func (this *DeviceConfigCenter) DeviceById(id string) *types.Device {
	device, _ := this.devices.Get(id).(*types.Device)
	return device
}

func (this *DeviceConfigCenter) HostConfigs(deviceId, hostId string) map[int32]*types.HostConfig {
	if this == nil {
		panic("nil")
	}
	device, _ := this.devices.Get(deviceId).(*types.Device)
	if device == nil {
		return nil
	}
	return device.Hosts[hostId].Configs
}

func Configs(resource common.IResources, serviceArea uint16) *DeviceConfigCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(ServiceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*DeviceConfigServicePoint)).configCenter
}
