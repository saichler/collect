package inventory

import (
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	"github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func RegisterInventoryCenter(serviceName string, serviceArea int32, elem proto.Message, primaryKey string,
	resources common.IResources, listener cache.ICacheListener) {
	this := &InventoryServicePoint{}
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, elem, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea)
	if err != nil {
		panic(err)
	}
}

func (this *InventoryServicePoint) Post(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	this.inventoryCenter.Add(pb)
	return nil, nil
}
func (this *InventoryServicePoint) Put(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Patch(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Delete(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Get(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) GetCopy(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) Failed(pb proto.Message, resourcs common.IResources, msg *types.Message) (proto.Message, error) {
	return nil, nil
}
func (this *InventoryServicePoint) EndPoint() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) ServiceName() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) Transactional() bool { return false }
func (this *InventoryServicePoint) ServiceModel() proto.Message {
	return this.inventoryCenter.element.(proto.Message)
}
func (this *InventoryServicePoint) ReplicationCount() int {
	return 0
}
func (this *InventoryServicePoint) ReplicationScore() int {
	return 0
}
