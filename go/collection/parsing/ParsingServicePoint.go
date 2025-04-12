package parsing

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

type ParsingServicePoint struct {
	resources   common.IResources
	serviceName string
	serviceArea uint16
}

func (this ParsingServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {

	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.resources = r
	this.resources.Registry().Register(&types.CMap{})
	this.resources.Registry().Register(&types.CTable{})
	this.resources.Registry().Register(&types.Job{})

	this.resources.ServicePoints().AddServicePointType(&inventory.InventoryServicePoint{})
	this.resources.ServicePoints().Activate(inventory.ServicePointType, serviceName+base.Inventory_Suffix,
		serviceArea, r, l, args...)
	return nil
}

func (this *ParsingServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	job := pb.Element().(*types.Job)
	resourcs.Logger().Debug("Job ", job.PollName, " completed!")
	JobComplete(job, this.resources)
	return nil
}
func (this *ParsingServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ParsingServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ParsingServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ParsingServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ParsingServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *ParsingServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *ParsingServicePoint) Transactional() bool { return false }

func (this *ParsingServicePoint) ReplicationCount() int {
	return 0
}
func (this *ParsingServicePoint) ReplicationScore() int {
	return 0
}
