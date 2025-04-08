package config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
)

type ConfigCenter struct {
	devices *cache.Cache
}

func newConfigCenter(serviceArea uint16, resources common.IResources, listener cache.ICacheListener) *ConfigCenter {
	this := &ConfigCenter{}
	this.devices = cache.NewModelCache(ServiceName, serviceArea, "Device", resources.SysConfig().LocalUuid, listener, resources.Introspector())
	return this
}

func (this *ConfigCenter) Add(device *types.Device) {
	this.devices.Put(device.Id, device)
}

func (this *ConfigCenter) DeviceById(id string) *types.Device {
	device, _ := this.devices.Get(id).(*types.Device)
	return device
}

func (this *ConfigCenter) HostConfigs(deviceId, hostId string) map[int32]*types.Config {
	if this == nil {
		panic("nil")
	}
	device, _ := this.devices.Get(deviceId).(*types.Device)
	if device == nil {
		return nil
	}
	return device.Hosts[hostId].Configs
}

func Configs(resource common.IResources, serviceArea uint16) *ConfigCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(ServiceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*ConfigServicePoint)).configCenter
}
