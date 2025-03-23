package control

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
)

type ParsingCenterNotifier struct {
	nic common.IVirtualNetworkInterface
}

func NewParsingCenterNotifier(nic common.IVirtualNetworkInterface) *ParsingCenterNotifier {
	jn := &ParsingCenterNotifier{}
	jn.nic = nic
	return jn
}

func (this *ParsingCenterNotifier) HandleCollectNotification(job *types.Job) {
	err := this.nic.Multicast(job.ServiceName, job.DServiceArea, types2.Action_POST, job)
	if err != nil {
		this.nic.Resources().Logger().Error("Failed to notify on job: ", err)
	}
}
