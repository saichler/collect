package inventory

import (
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

var TOPIC = ""
var ENDPOINT = ""

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func RegisterInventoryCenter(area int32, elem proto.Message, primaryKey string, resources interfaces.IResources, listener cache.ICacheListener) {
	this := &InventoryServicePoint{}
	this.inventoryCenter = newInventoryCenter(primaryKey, elem, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(area, elem, this)
	if err != nil {
		panic(err)
	}
	this.inventoryCenter.setTopic(elem, resources)
}

func (this *InventoryServicePoint) Post(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	this.inventoryCenter.Add(pb)
	return nil, nil
}
func (this *InventoryServicePoint) Put(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Patch(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Delete(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Get(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Failed(pb proto.Message, vnic interfaces.IVirtualNetworkInterface, msg *types.Message) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *InventoryServicePoint) Topic() string {
	return TOPIC
}
