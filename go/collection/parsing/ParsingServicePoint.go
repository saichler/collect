package parsing

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

const (
	ServicePointType = "ParsingServicePoint"
)

type ParsingServicePoint struct {
	resources  common.IResources
	elem       interface{}
	primaryKey string
	vnic       common.IVirtualNetworkInterface
}

func (this *ParsingServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {

	this.resources = r
	this.resources.Registry().Register(&types.CMap{})
	this.resources.Registry().Register(&types.CTable{})
	this.resources.Registry().Register(&types.Job{})
	this.elem = args[0]
	this.primaryKey = args[1].(string)
	vnic, ok := l.(common.IVirtualNetworkInterface)
	if ok {
		this.vnic = vnic
	}
	this.resources.Introspector().Inspect(this.elem)
	return nil
}

func (this *ParsingServicePoint) DeActivate() error {
	this.vnic = nil
	this.resources = nil
	this.elem = nil
	return nil
}

func (this *ParsingServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	job := pb.Element().(*types.Job)
	resourcs.Logger().Debug("Job ", job.PollName, " completed!")
	this.JobComplete(job, this.resources)
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
