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

func (this *DeviceConfigCenter) Shutdown() {
	this.devices = nil
}

func (this *DeviceConfigCenter) Add(device *types.DeviceConfig) {
	this.devices.Put(device.DeviceId, device)
}

func (this *DeviceConfigCenter) DeviceById(id string) *types.DeviceConfig {
	device, _ := this.devices.Get(id).(*types.DeviceConfig)
	return device
}

func (this *DeviceConfigCenter) HostConnectionConfigs(deviceId, hostId string) map[int32]*types.ConnectionConfig {
	if this == nil {
		panic("nil")
	}
	device, _ := this.devices.Get(deviceId).(*types.DeviceConfig)
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
