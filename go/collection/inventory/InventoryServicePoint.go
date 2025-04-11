package inventory

import (
	"github.com/saichler/types/go/common"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func RegisterInventoryCenter(serviceName string, serviceArea uint16, elem common.IElements, primaryKey string,
	resources common.IResources, vnic common.IVirtualNetworkInterface) {
	this := &InventoryServicePoint{}
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, elem, resources, vnic)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea, vnic)
	if err != nil {
		panic(err)
	}
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
func (this *InventoryServicePoint) EndPoint() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) ServiceName() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) Transactional() bool { return false }
func (this *InventoryServicePoint) ServiceModel() common.IElements {
	return this.inventoryCenter.element.(common.IElements)
}
func (this *InventoryServicePoint) ReplicationCount() int {
	return 0
}
func (this *InventoryServicePoint) ReplicationScore() int {
	return 0
}
