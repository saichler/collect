package polling

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
)

const (
	ServiceName = "Poll"
	ENDPOINT    = "poll"
)

type PollServicePoint struct {
	pollCenter *PollCenter
}

func RegisterPollCenter(serviceArea uint16, resources common.IResources, listener cache.ICacheListener) {
	psp := &PollServicePoint{}
	psp.pollCenter = newPollCenter(serviceArea, resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(psp, serviceArea)
	if err != nil {
		panic(err)
	}
}

func (this *PollServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	hp := pb.Element().(*types.Poll)
	this.pollCenter.Add(hp)
	return nil
}
func (this *PollServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *PollServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *PollServicePoint) ServiceName() string {
	return ServiceName
}
func (this *PollServicePoint) Transactional() bool { return false }
func (this *PollServicePoint) ServiceModel() common.IElements {
	return object.New(nil, &types.Poll{})
}
func (this *PollServicePoint) ReplicationCount() int {
	return 0
}
func (this *PollServicePoint) ReplicationScore() int {
	return 0
}
