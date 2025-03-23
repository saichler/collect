package parsing

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
	"strings"
)

type ParsingServicePoint struct {
	resources   common.IResources
	serviceName string
	serviceArea int32
}

func RegisterParsingServicePoint(serviceName string, serviceArea int32, elem proto.Message,
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

func (this *ParsingServicePoint) Post(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	job := pb.(*types.Job)
	resourcs.Logger().Debug("Job ", job.PollName, " completed!")
	JobComplete(job, this.resources)
	return nil, nil
}
func (this *ParsingServicePoint) Put(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Patch(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Delete(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Get(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) GetCopy(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) Failed(pb proto.Message, resourcs common.IResources, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *ParsingServicePoint) EndPoint() string {
	return strings.ToLower(this.ServiceName())
}
func (this *ParsingServicePoint) ServiceName() string {
	return this.serviceName + base.Parsing_Suffix
}
func (this *ParsingServicePoint) Transactional() bool { return false }
func (this *ParsingServicePoint) ServiceModel() proto.Message {
	return &types.Job{}
}
func (this *ParsingServicePoint) ReplicationCount() int {
	return 0
}
func (this *ParsingServicePoint) ReplicationScore() int {
	return 0
}
