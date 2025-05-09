package inventory

import (
	types2 "github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServicePointType = "InventoryServicePoint"
)

type InventoryServicePoint struct {
	inventoryCenter *InventoryCenter
	forwardService  *types2.DeviceServiceInfo
	nic             ifs.IVNic
}

func (this *InventoryServicePoint) Activate(serviceName string, serviceArea uint16,
	r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {
	r.Logger().Info("Activated Inventory on ", serviceName, " area ", serviceArea)
	primaryKey := args[0].(string)
	this.inventoryCenter = newInventoryCenter(serviceName, serviceArea, primaryKey, args[1], r, l)
	if len(args) == 3 {
		this.forwardService = args[2].(*types2.DeviceServiceInfo)
		this.nic = l.(ifs.IVNic)
		r.Logger().Info("Added forwarding to ", this.forwardService.ServiceName, " area ", this.forwardService.ServiceArea)
	}
	return nil
}

func (this *InventoryServicePoint) DeActivate() error {
	this.inventoryCenter = nil
	return nil
}

func (this *InventoryServicePoint) Post(elements ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	resourcs.Logger().Info("Post Received inventory item...")
	this.inventoryCenter.Add(elements.Element(), elements.Notification())
	if !elements.Notification() {
		go func() {
			if this.forwardService != nil {
				resourcs.Logger().Info("Forawrding Post to ", this.forwardService.ServiceName, " area ",
					this.forwardService.ServiceArea)
				elem := this.inventoryCenter.ElementByElement(elements.Element())
				resp := this.nic.SingleRequest(this.forwardService.ServiceName, uint16(this.forwardService.ServiceArea),
					ifs.POST, elem)
				if resp != nil && resp.Error() != nil {
					resourcs.Logger().Error(resp.Error().Error())
				} else {
					resourcs.Logger().Info("Post Finished to ", this.forwardService.ServiceName, " area ",
						this.forwardService.ServiceArea)
				}
			}
		}()
	}
	return nil
}

func (this *InventoryServicePoint) Put(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *InventoryServicePoint) Patch(elements ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	resourcs.Logger().Info("Patch Received inventory item...")
	this.inventoryCenter.Update(elements.Element(), elements.Notification())
	if !elements.Notification() {
		go func() {
			if this.forwardService != nil {
				resourcs.Logger().Info("Patch Forawrding to ", this.forwardService.ServiceName, " area ",
					this.forwardService.ServiceArea)
				elem := this.inventoryCenter.ElementByElement(elements.Element())
				resp := this.nic.SingleRequest(this.forwardService.ServiceName,
					uint16(this.forwardService.ServiceArea), ifs.POST, elem)
				if resp != nil && resp.Error() != nil {
					resourcs.Logger().Error(resp.Error().Error())
				} else {
					resourcs.Logger().Info("Patch Finished to ", this.forwardService.ServiceName, " area ",
						this.forwardService.ServiceArea)
				}
			}
		}()
	}
	return nil
}
func (this *InventoryServicePoint) Delete(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *InventoryServicePoint) Get(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *InventoryServicePoint) GetCopy(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *InventoryServicePoint) Failed(pb ifs.IElements, resourcs ifs.IResources, msg ifs.IMessage) ifs.IElements {
	return nil
}
func (this *InventoryServicePoint) TransactionMethod() ifs.ITransactionMethod {
	return nil
}

/*
func (this *InventoryServicePoint) Replication() bool {
	return false
}
func (this *InventoryServicePoint) ReplicationCount() int {
	return 0
}
func (this *InventoryServicePoint) KeyOf(elements ifs.IElements) string {
	return ""
}*/
