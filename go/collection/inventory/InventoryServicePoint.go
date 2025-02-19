package inventory

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC    = "NetworkBox"
	ENDPOINT = "networkbox"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func RegisterInventoryCenter(resources interfaces.IResources, listener cache.ICacheListener) {
	this := &InventoryServicePoint{}
	this.inventoryCenter = newInventoryCenter(resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(&types.NetworkBox{}, this)
	if err != nil {
		panic(err)
	}
}

func (this *InventoryServicePoint) Post(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	box := pb.(*types.NetworkBox)
	this.inventoryCenter.Add(box)
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
func (this *InventoryServicePoint) Failed(pb proto.Message, vnic interfaces.IVirtualNetworkInterface, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *InventoryServicePoint) Topic() string {
	return TOPIC
}
