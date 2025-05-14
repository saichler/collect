package device_config

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/collector"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8srlz/go/serialize/object"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "DeviceConfig"
	ServiceType = "DeviceConfigService"
)

type DeviceConfigService struct {
	configCenter *DeviceConfigCenter
	controller   base.IController
}

func (this *DeviceConfigService) Activate(serviceName string, serviceArea uint16,
	r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.DeviceConfig{})
	this.configCenter = newConfigCenter(ServiceName, serviceArea, r, l)
	if args == nil {
		vnic, ok := l.(ifs.IVNic)
		if ok {
			pt := collector.NewParsingCenterNotifier(vnic)
			this.controller = collector.NewDeviceCollector(pt, r)
		}
	} else {
		this.controller, _ = args[0].(base.IController)
	}
	return nil
}

func (this *DeviceConfigService) DeActivate() error {
	this.controller.Shutdown()
	this.configCenter.Shutdown()
	this.controller = nil
	this.configCenter = nil
	return nil
}

func (this *DeviceConfigService) Post(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	device := pb.Element().(*types.DeviceConfig)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.DeviceId)
		this.controller.StartPolling(device)
	}
	return object.New(nil, &types.DeviceConfig{})
}
func (this *DeviceConfigService) Put(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) Patch(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) Delete(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) Get(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) GetCopy(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) Failed(pb ifs.IElements, resourcs ifs.IResources, msg ifs.IMessage) ifs.IElements {
	return nil
}
func (this *DeviceConfigService) TransactionMethod() ifs.ITransactionMethod {
	return nil
}
func (this *DeviceConfigService) WebService() ifs.IWebService {
	return nil
}
