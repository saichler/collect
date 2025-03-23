package parsing

import (
	"fmt"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

func JobComplete(job *types.Job, resources common.IResources) {
	pc := polling.Polling(resources, job.DServiceArea)
	poll := pc.PollByName(job.PollName)
	if poll == nil {
		resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}
	if job.Error == "" && poll.Attributes != nil {
		inv := inventory.Inventory(resources, job.ServiceName, job.DServiceArea)
		elem := inv.ElementByKey(job.DeviceId)
		if elem == nil {
			inv.AddEmpty(job.DeviceId)
			elem = inv.ElementByKey(job.DeviceId)
		}
		Parser.Parse(job, elem, resources)
		fmt.Println(elem)
		inv.Update(elem)
	}
}
