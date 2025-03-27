package polling

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
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

func (this *PollServicePoint) Post(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	hp := pb.Element().(*types.Poll)
	this.pollCenter.Add(hp)
	return nil
}
func (this *PollServicePoint) Put(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *PollServicePoint) Patch(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *PollServicePoint) Delete(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *PollServicePoint) Get(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *PollServicePoint) GetCopy(pb common.IMObjects, resourcs common.IResources) common.IMObjects {
	return nil
}
func (this *PollServicePoint) Failed(pb common.IMObjects, resourcs common.IResources, msg *types2.Message) common.IMObjects {
	return nil
}
func (this *PollServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *PollServicePoint) ServiceName() string {
	return ServiceName
}
func (this *PollServicePoint) Transactional() bool { return false }
func (this *PollServicePoint) ServiceModel() common.IMObjects {
	return object.New(nil, &types.Poll{})
}
func (this *PollServicePoint) ReplicationCount() int {
	return 0
}
func (this *PollServicePoint) ReplicationScore() int {
	return 0
}
