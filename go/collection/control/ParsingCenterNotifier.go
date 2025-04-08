package control

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
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
	err := this.nic.Multicast(job.ServiceName, uint16(job.DServiceArea), common.POST, job)
	if err != nil {
		this.nic.Resources().Logger().Error("Failed to notify on job: ", err)
	}
}
