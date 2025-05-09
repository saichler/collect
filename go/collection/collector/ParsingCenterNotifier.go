package collector

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/layer8/go/overlay/health"
	"github.com/saichler/l8types/go/ifs"
)

type ParsingCenterNotifier struct {
	nic ifs.IVNic
}

func NewParsingCenterNotifier(nic ifs.IVNic) *ParsingCenterNotifier {
	jn := &ParsingCenterNotifier{}
	jn.nic = nic
	return jn
}

func (this *ParsingCenterNotifier) JobCompleted(job *types.Job) {
	this.nic.Resources().Logger().Info("Job ", job.PollName, " took:", (job.Ended - job.Started))
	dest, err := this.nic.Single(job.PService.ServiceName, uint16(job.PService.ServiceArea), ifs.POST, job)
	if err != nil {
		this.nic.Resources().Logger().Error("Failed to notify on job: ", err)
	}
	hc := health.Health(this.nic.Resources())
	hp := hc.HealthPoint(dest)
	this.nic.Resources().Logger().Info("Sent job complete to : ", hp.Alias)
}
