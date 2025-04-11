package config

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/types/go/common"
)

const (
	ServiceName = "Config"
	ENDPOINT    = "config"
)

type ConfigServicePoint struct {
	configCenter *ConfigCenter
	controller   base.IController
}

func RegisterConfigCenter(serviceArea uint16, resources common.IResources, vnic common.IVirtualNetworkInterface,
	controller base.IController) {
	this := &ConfigServicePoint{}
	this.controller = controller
	this.configCenter = newConfigCenter(serviceArea, resources, vnic)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea, vnic)
	if err != nil {
		panic(err)
	}
}

func (this *ConfigServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	device := pb.Element().(*types.Device)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.Id, " ", device.ServiceName)
		this.controller.StartPolling(device.Id, device.ServiceName)
	}
	return nil
}
func (this *ConfigServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ConfigServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ConfigServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ConfigServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ConfigServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ConfigServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *ConfigServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *ConfigServicePoint) ServiceName() string {
	return ServiceName
}
func (this *ConfigServicePoint) Transactional() bool { return false }
func (this *ConfigServicePoint) ServiceModel() common.IElements {
	return object.New(nil, &types.Device{})
}
func (this *ConfigServicePoint) ReplicationCount() int {
	return 0
}
func (this *ConfigServicePoint) ReplicationScore() int {
	return 0
}
