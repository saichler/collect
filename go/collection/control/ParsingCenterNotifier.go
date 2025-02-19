package control

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
)

type ParsingCenterNotifier struct {
	nic interfaces.IVirtualNetworkInterface
}

func NewParsingCenterNotifier(nic interfaces.IVirtualNetworkInterface) *ParsingCenterNotifier {
	jn := &ParsingCenterNotifier{}
	jn.nic = nic
	return jn
}

func (this *ParsingCenterNotifier) HandleCollectNotification(job *types.Job) {
	this.nic.Do(types2.Action_POST, "Job", job)
}
