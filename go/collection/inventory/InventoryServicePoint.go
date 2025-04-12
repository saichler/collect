package inventory

import (
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/types/go/common"
)

const (
	ServicePointType = "InventoryServicePoint"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func (this InventoryServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {
	primaryKey := args[0].(string)
	elem := object.New(nil, args[1])
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, elem, r, l)
	return nil
}

func (this *InventoryServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	this.inventoryCenter.Add(pb)
	return nil
}
func (this *InventoryServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Transactional() bool { return false }

func (this *InventoryServicePoint) ReplicationCount() int {
	return 0
}
func (this *InventoryServicePoint) ReplicationScore() int {
	return 0
}
