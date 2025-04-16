package collector

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/maps"
	"github.com/saichler/shared/go/share/strings"
	"github.com/saichler/types/go/common"
)

type DeviceCollector struct {
	hostCollectors *maps.SyncMap
	handler        base.IJobCompleteHandler
	resources      common.IResources
}

func NewDeviceCollector(handler base.IJobCompleteHandler, resources common.IResources) *DeviceCollector {
	resources.Logger().Debug("*** Creating new collector for vnet ", resources.SysConfig().VnetPort)
	collector := &DeviceCollector{}
	collector.resources = resources
	collector.hostCollectors = maps.NewSyncMap()
	collector.handler = handler
	resources.Registry().Register(&types.CMap{})
	resources.Registry().Register(&types.CTable{})
	return collector
}

func (this *DeviceCollector) StartPolling(device *types.DeviceConfig) error {
	for _, host := range device.Hosts {
		hostCol, _ := this.hostCollector(host.DeviceId, device)
		hostCol.start()
	}
	return nil
}

func (this *DeviceCollector) Shutdown() {
	this.hostCollectors.Iterate(func(k, v interface{}) {
		h := v.(*HostCollector)
		h.stop()
	})
	this.hostCollectors = nil
}

func (this *DeviceCollector) hostCollector(hostId string, device *types.DeviceConfig) (*HostCollector, bool) {
	key := hostKey(device.DeviceId, hostId)
	h, ok := this.hostCollectors.Get(key)
	if ok {
		return h.(*HostCollector), ok
	}
	hc := newHostCollector(device, hostId, this.resources, this.handler)
	this.hostCollectors.Put(key, hc)
	return hc, ok
}

func hostKey(deviceId, hostId string) string {
	return strings.New(deviceId, hostId).String()
}
