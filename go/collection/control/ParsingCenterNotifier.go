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

func (this *ParsingCenterNotifier) HandleCollectNotification(job *types.Job, area int32) {
	this.nic.Multicast(types2.CastMode_All, types2.Action_POST, area, "Job", job)
}
