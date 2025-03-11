package polling

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC    = "Poll"
	ENDPOINT = "poll"
)

type PollServicePoint struct {
	pollCenter *PollCenter
}

func RegisterPollCenter(area int32, resources common.IResources, listener cache.ICacheListener) {
	psp := &PollServicePoint{}
	psp.pollCenter = newPollCenter(resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(area, &types.Poll{}, psp)
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
func (this *PollServicePoint) Topic() string {
	return TOPIC
}
func (this *PollServicePoint) Transactional() bool { return false }
