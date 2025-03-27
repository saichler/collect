package inventory

import (
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	"github.com/saichler/types/go/types"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
}

func RegisterInventoryCenter(serviceName string, serviceArea int32, elem common.IMObjects, primaryKey string,
	resources common.IResources, listener cache.ICacheListener) {
	this := &InventoryServicePoint{}
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, elem, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea)
	if err != nil {
		panic(err)
	}
}

func (this *InventoryServicePoint) Post(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	this.inventoryCenter.Add(pb)
	return nil
}
func (this *InventoryServicePoint) Put(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) Patch(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) Delete(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) Get(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) GetCopy(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) Failed(pb common.IMObjects, resourcs common.IResources, msg *types.Message) common.IMObjects {
	return nil
}
func (this *InventoryServicePoint) EndPoint() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) ServiceName() string {
	return this.inventoryCenter.serviceName
}
func (this *InventoryServicePoint) Transactional() bool { return false }
func (this *InventoryServicePoint) ServiceModel() common.IMObjects {
	return this.inventoryCenter.element.(common.IMObjects)
}
func (this *InventoryServicePoint) ReplicationCount() int {
	return 0
}
func (this *InventoryServicePoint) ReplicationScore() int {
	return 0
}
