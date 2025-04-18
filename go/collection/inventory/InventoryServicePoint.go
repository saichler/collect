package inventory

import (
	types2 "github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

const (
	ServicePointType = "InventoryServicePoint"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
	forwardService  *types2.DeviceServiceInfo
	nic             common.IVirtualNetworkInterface
}

func (this *InventoryServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {
	primaryKey := args[0].(string)
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, args[1], r, l)
	if len(args) == 3 {
		this.forwardService = args[2].(*types2.DeviceServiceInfo)
		this.nic = l.(common.IVirtualNetworkInterface)
	}
	return nil
}

func (this *InventoryServicePoint) DeActivate() error {
	this.inventoryCenter = nil
	return nil
}

func (this *InventoryServicePoint) Post(elements common.IElements, resourcs common.IResources) common.IElements {
	this.inventoryCenter.Add(elements.Element())
	if this.forwardService != nil {
		elem := this.inventoryCenter.ElementByElement(elements.Element())
		this.nic.Single(this.forwardService.ServiceName, uint16(this.forwardService.ServiceArea),
			common.POST, elem)
	}
	return nil
}
func (this *InventoryServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *InventoryServicePoint) Patch(elements common.IElements, resourcs common.IResources) common.IElements {
	this.inventoryCenter.Update(elements.Element())
	if this.forwardService != nil {
		elem := this.inventoryCenter.ElementByElement(elements.Element())
		this.nic.Single(this.forwardService.ServiceName, uint16(this.forwardService.ServiceArea),
			common.POST, elem)
	}
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
