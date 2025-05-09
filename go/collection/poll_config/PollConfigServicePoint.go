package poll_config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8srlz/go/serialize/object"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName      = "PollConfig"
	ServiceArea      = 0
	ServicePointType = "PollConfigServicePoint"
)

type PollConfigServicePoint struct {
	pollCenter *PollConfigCenter
}

func (this *PollConfigServicePoint) Activate(serviceName string, serviceArea uint16,
	r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.PollConfig{})
	this.pollCenter = newPollConfigCenter(r, l)
	return nil
}

func (this *PollConfigServicePoint) DeActivate() error {
	this.pollCenter = nil
	return nil
}

func (this *PollConfigServicePoint) Post(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	pollConfig := pb.Element().(*types.PollConfig)
	resourcs.Logger().Info("Added a poll config ", pollConfig.Name)
	return object.New(this.pollCenter.Add(pollConfig, pb.Notification()), &types.PollConfig{})
}
func (this *PollConfigServicePoint) Put(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) Patch(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) Delete(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) Get(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) GetCopy(pb ifs.IElements, resourcs ifs.IResources) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) Failed(pb ifs.IElements, resourcs ifs.IResources, msg ifs.IMessage) ifs.IElements {
	return nil
}
func (this *PollConfigServicePoint) TransactionMethod() ifs.ITransactionMethod {
	return nil
}
