package parsing

import (
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
)

func jobComplete(job *types.Job, resources interfaces.IResources) {
	pc := polling.Polling(resources)
	poll := pc.PollByName(job.PollName)
	if poll == nil {
		resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}
	if job.Error == "" && poll.Attributes != nil {
		inv := inventory.Inventory(resources)
		box := inv.BoxById(job.DeviceId)
		if box == nil {
			box = &types.NetworkBox{}
			box.Id = job.DeviceId
			inv.Add(box)
		}
		Parser.Parse(job, box, resources)
		inv.Update(box)
	}
}
