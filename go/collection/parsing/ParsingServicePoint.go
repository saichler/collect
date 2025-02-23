package parsing

import (
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC    = "Job"
	ENDPOINT = "job"
)

type ParsingServicePoint struct {
	resources interfaces.IResources
}

func RegisterParsingServicePoint(area int32, elem proto.Message, primaryKeyAttr string, resources interfaces.IResources) {
	this := &ParsingServicePoint{}
	this.resources = resources
	inventory.RegisterInventoryCenter(area, elem, primaryKeyAttr, resources, nil)
	err := resources.ServicePoints().RegisterServicePoint(area, &types.Job{}, this)
	resources.Registry().Register(&types.Map{})
	resources.Registry().Register(&types.Table{})
	if err != nil {
		panic(err)
	}
}

func (this *ParsingServicePoint) Post(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	job := pb.(*types.Job)
	vnic.Resources().Logger().Debug("Job ", job.PollName, " completed!")
	JobComplete(job, this.resources)
	return nil, nil
}
func (this *ParsingServicePoint) Put(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Patch(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Delete(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Get(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Failed(pb proto.Message, vnic interfaces.IVirtualNetworkInterface, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *ParsingServicePoint) Topic() string {
	return TOPIC
}
