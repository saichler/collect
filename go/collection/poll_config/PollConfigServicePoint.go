package poll_config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

const (
	ServiceName      = "PollConfig"
	ServicePointType = "PollConfigServicePoint"
)

type PollConfigServicePoint struct {
	pollCenter *PollConfigCenter
}

func (this PollConfigServicePoint) Activate(serviceName string, serviceArea uint16,
	r common.IResources, l common.IServicePointCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.PollConfig{})
	this.pollCenter = newPollConfigCenter(serviceArea, r, l)
	return nil
}

func (this *PollConfigServicePoint) Post(pb common.IElements, resourcs common.IResources) common.IElements {
	hp := pb.Element().(*types.PollConfig)
	this.pollCenter.Add(hp)
	return nil
}
func (this *PollConfigServicePoint) Put(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) Patch(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) Delete(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) Get(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) GetCopy(pb common.IElements, resourcs common.IResources) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) Failed(pb common.IElements, resourcs common.IResources, msg common.IMessage) common.IElements {
	return nil
}
func (this *PollConfigServicePoint) Transactional() bool { return false }

func (this *PollConfigServicePoint) ReplicationCount() int {
	return 0
}
func (this *PollConfigServicePoint) ReplicationScore() int {
	return 0
}
