package polling

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	ServiceName = "Poll"
	ENDPOINT    = "poll"
)

type PollServicePoint struct {
	pollCenter *PollCenter
}

func RegisterPollCenter(serviceArea int32, resources common.IResources, listener cache.ICacheListener) {
	psp := &PollServicePoint{}
	psp.pollCenter = newPollCenter(serviceArea, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(psp, serviceArea)
	if err != nil {
		panic(err)
	}
}

func (this *PollServicePoint) Post(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	hp := pb.(*types.Poll)
	this.pollCenter.Add(hp)
	return nil, nil
}
func (this *PollServicePoint) Put(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Patch(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Delete(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Get(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) GetCopy(pb proto.Message, resourcs common.IResources) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Failed(pb proto.Message, resourcs common.IResources, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *PollServicePoint) ServiceName() string {
	return ServiceName
}
func (this *PollServicePoint) Transactional() bool { return false }
func (this *PollServicePoint) ServiceModel() proto.Message {
	return &types.Poll{}
}
func (this *PollServicePoint) ReplicationCount() int {
	return 0
}
func (this *PollServicePoint) ReplicationScore() int {
	return 0
}
