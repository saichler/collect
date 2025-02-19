package config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
)

type ConfigCenter struct {
	devices *cache.Cache
}

func newConfigCenter(resources interfaces.IResources, listener cache.ICacheListener) *ConfigCenter {
	this := &ConfigCenter{}
	this.devices = cache.NewModelCache(resources.Config().LocalUuid, listener, resources.Introspector())
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
	device, _ := this.devices.Get(deviceId).(*types.Device)
	if device == nil {
		return nil
	}
	return device.Hosts[hostId].Configs
}

func Configs(resource interfaces.IResources) *ConfigCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(TOPIC)
	if !ok {
		return nil
	}
	return (sp.(*ConfigServicePoint)).configCenter
}
