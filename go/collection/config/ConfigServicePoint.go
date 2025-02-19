package config

import (
	"github.com/saichler/collect/go/collection/common"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC    = "Device"
	ENDPOINT = "device"
)

type ConfigServicePoint struct {
	configCenter *ConfigCenter
	controller   common.IController
}

func RegisterConfigCenter(resources interfaces.IResources, listener cache.ICacheListener,
	controller common.IController) {
	this := &ConfigServicePoint{}
	this.controller = controller
	this.configCenter = newConfigCenter(resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(&types.Device{}, this)
	if err != nil {
		panic(err)
	}
}

func (this *ConfigServicePoint) Post(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	device := pb.(*types.Device)
	this.configCenter.Add(device)
	if this.controller != nil {
		this.controller.StartPolling(device.Id)
	}
	return nil, nil
}
func (this *ConfigServicePoint) Put(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Patch(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Delete(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Get(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) Failed(pb proto.Message, vnic interfaces.IVirtualNetworkInterface, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *ConfigServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *ConfigServicePoint) Topic() string {
	return TOPIC
}
