package config

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
)

const (
	ServiceName = "Config"
	ENDPOINT    = "config"
)

type ConfigServicePoint struct {
	configCenter *ConfigCenter
	controller   base.IController
}

func RegisterConfigCenter(serviceArea int32, resources common.IResources, listener cache.ICacheListener,
	controller base.IController) {
	this := &ConfigServicePoint{}
	this.controller = controller
	this.configCenter = newConfigCenter(serviceArea, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea)
	if err != nil {
		panic(err)
	}
}

var Count = 0

func (this *ConfigServicePoint) Post(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	device := pb.Element().(*types.Device)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.Id, " ", device.ServiceName)
		this.controller.StartPolling(device.Id, device.ServiceName)
	}
	return nil
}
func (this *ConfigServicePoint) Put(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) Patch(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) Delete(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) Get(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) GetCopy(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) Failed(pb common.IMObjects, resourcs common.IResources, msg *types2.Message) common.IMObjects {
	return nil
}
func (this *ConfigServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *ConfigServicePoint) ServiceName() string {
	return ServiceName
}
func (this *ConfigServicePoint) Transactional() bool { return false }
func (this *ConfigServicePoint) ServiceModel() common.IMObjects {
	return object.New(nil, &types.Device{})
}
func (this *ConfigServicePoint) ReplicationCount() int {
	return 0
}
func (this *ConfigServicePoint) ReplicationScore() int {
	return 0
}
