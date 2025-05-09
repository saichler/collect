package parsing

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServicePointType = "ParsingServicePoint"
)

type ParsingServicePoint struct {
	resources  ifs.IResources
	elem       interface{}
	primaryKey string
	vnic       ifs.IVNic
}

func (this *ParsingServicePoint) Activate(serviceName string, serviceArea uint16,
	r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {

	this.resources = r
	this.resources.Registry().Register(&types.CMap{})
	this.resources.Registry().Register(&types.CTable{})
	this.resources.Registry().Register(&types.Job{})
	this.elem = args[0]
	this.primaryKey = args[1].(string)
	vnic, ok := l.(ifs.IVNic)
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

func (this *ParsingServicePoint) Post(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	job := pb.Element().(*types.Job)
	resourcs.Logger().Info("Received Job ", job.PollName, " completed!")
	this.JobComplete(job, this.resources)
	return nil
}
func (this *ParsingServicePoint) Put(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) Patch(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) Delete(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) Get(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) GetCopy(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) Failed(pb ifs.IElements, resourcs ifs.IResources, msg ifs.IMessage) ifs.IElements {
	return nil
}
func (this *ParsingServicePoint) TransactionMethod() ifs.ITransactionMethod {
	return nil
}
