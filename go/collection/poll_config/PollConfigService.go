package poll_config

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8srlz/go/serialize/object"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "PollConfig"
	ServiceArea = 0
	ServiceType = "PollConfigService"
)

type PollConfigService struct {
	pollCenter *PollConfigCenter
}

func (this *PollConfigService) Activate(serviceName string, serviceArea uint16,
	r ifs.IResources, l ifs.IServiceCacheListener, args ...interface{}) error {
	r.Registry().Register(&types.PollConfig{})
	this.pollCenter = newPollConfigCenter(r, l)
	return nil
}

func (this *PollConfigService) DeActivate() error {
	this.pollCenter = nil
	return nil
}

func (this *PollConfigService) Post(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	pollConfig := pb.Element().(*types.PollConfig)
	vnic.Resources().Logger().Info("Added a poll config ", pollConfig.Name)
	return object.New(this.pollCenter.Add(pollConfig, pb.Notification()), &types.PollConfig{})
}
func (this *PollConfigService) Put(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	return nil
}
func (this *PollConfigService) Patch(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	return nil
}
func (this *PollConfigService) Delete(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	return nil
}
func (this *PollConfigService) Get(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	return nil
}
func (this *PollConfigService) GetCopy(pb ifs.IElements, vnic ifs.IVNic) ifs.IElements {
	return nil
}
func (this *PollConfigService) Failed(pb ifs.IElements, vnic ifs.IVNic, msg ifs.IMessage) ifs.IElements {
	return nil
}
func (this *PollConfigService) TransactionMethod() ifs.ITransactionMethod {
	return nil
}
func (this *PollConfigService) WebService() ifs.IWebService {
	return nil
}
