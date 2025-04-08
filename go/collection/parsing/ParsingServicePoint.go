package parsing

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	"strings"
)

type ParsingServicePoint struct {
	resources   common.IResources
	serviceName string
	serviceArea uint16
}

func RegisterParsingServicePoint(serviceName string, serviceArea uint16, elem common.IElements,
	primaryKeyAttr string, resources common.IResources, listener cache.ICacheListener) {
	this := &ParsingServicePoint{}
	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.resources = resources

	resources.Registry().Register(&types.Map{})
	resources.Registry().Register(&types.Table{})

	inventory.RegisterInventoryCenter(serviceName, serviceArea, elem, primaryKeyAttr, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(this, serviceArea)
	if err != nil {
		panic(err)
	}
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
func (this *ParsingServicePoint) EndPoint() string {
	return strings.ToLower(this.ServiceName())
}
func (this *ParsingServicePoint) ServiceName() string {
	return this.serviceName + base.Parsing_Suffix
}
func (this *ParsingServicePoint) Transactional() bool { return false }
func (this *ParsingServicePoint) ServiceModel() common.IElements {
	return object.New(nil, &types.Job{})
}
func (this *ParsingServicePoint) ReplicationCount() int {
	return 0
}
func (this *ParsingServicePoint) ReplicationScore() int {
	return 0
}
