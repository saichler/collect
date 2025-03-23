package config

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
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

func (this *ConfigServicePoint) Post(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	device := pb.(*types.Device)
	this.configCenter.Add(device)
	if this.controller != nil {
		resourcs.Logger().Info("Start Polling Device ", device.Id, " ", device.ServiceName)
		this.controller.StartPolling(device.Id, device.ServiceName)
	}
	return nil, nil
}
func (this *ConfigServicePoint) Put(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Patch(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Delete(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Get(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) GetCopy(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Failed(pb proto.Message, resourcs common.IResources, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *ConfigServicePoint) ServiceName() string {
	return ServiceName
}
func (this *ConfigServicePoint) Transactional() bool { return false }
func (this *ConfigServicePoint) ServiceModel() proto.Message {
	return &types.Device{}
}
func (this *ConfigServicePoint) ReplicationCount() int {
	return 0
}
func (this *ConfigServicePoint) ReplicationScore() int {
	return 0
}
