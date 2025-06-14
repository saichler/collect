package device_config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8services/go/services/dcache"
	"github.com/saichler/l8types/go/ifs"
)

type DeviceConfigCenter struct {
	devices ifs.IDistributedCache
}

func newConfigCenter(serviceName string, serviceArea uint16, resources ifs.IResources, listener ifs.IServiceCacheListener) *DeviceConfigCenter {
	this := &DeviceConfigCenter{}
	this.devices = dcache.NewDistributedCache(serviceName, serviceArea, "Device", resources.SysConfig().LocalUuid, nil, resources)
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

func Configs(resource ifs.IResources, serviceArea uint16) *DeviceConfigCenter {
	sp, ok := resource.Services().ServiceHandler(ServiceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*DeviceConfigService)).configCenter
}
