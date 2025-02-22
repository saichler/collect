package parsing

import (
	"fmt"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
)

func JobComplete(job *types.Job, resources interfaces.IResources) {
	fmt.Println(string(job.Result))
	pc := polling.Polling(resources)
	poll := pc.PollByName(job.PollName)
	if poll == nil {
		resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}
	if job.Error == "" && poll.Attributes != nil {
		inv := inventory.Inventory(resources)
		elem := inv.ElementByKey(job.DeviceId)
		if elem == nil {
			elem = inv.AddEmpty(job.DeviceId)
		}
		Parser.Parse(job, elem, resources)
		inv.Update(elem)
	}
}
