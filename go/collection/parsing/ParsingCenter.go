package parsing

import (
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
	"reflect"
)

func (this *ParsingServicePoint) JobComplete(job *types.Job, resources common.IResources) {
	pc := poll_config.PollConfig(resources)
	poll := pc.PollByName(job.PollName)

	if poll == nil {
		resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}

	if job.Error == "" && poll.Parsing != nil && poll.Parsing.Attributes != nil {
		newElem := reflect.New(reflect.ValueOf(this.elem).Elem().Type())
		field := newElem.Elem().FieldByName(this.primaryKey)
		field.Set(reflect.ValueOf(job.DeviceId))
		elem := newElem.Interface()
		err := Parser.Parse(job, elem, resources)
		if err != nil {
			panic(err)
		}
		if this.vnic == nil {
			resources.Logger().Error("No Vnic to notify inventory")
			return
		}
		this.vnic.Multicast(job.IService.ServiceName, uint16(job.IService.ServiceArea),
			common.PATCH, elem)
	}
}
