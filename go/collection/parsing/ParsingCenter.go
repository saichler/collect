package parsing

import (
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
	"reflect"
)

func (this *ParsingService) JobComplete(job *types.Job, resources ifs.IResources) {
	pc := poll_config.PollConfig(resources)
	poll := pc.PollByName(job.PollName)

	if poll == nil {
		resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}

	if job.Error != "" {
		resources.Logger().Error("job error: ", job.Error)
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
		_, err = this.vnic.Single(job.IService.ServiceName, uint16(job.IService.ServiceArea),
			ifs.PATCH, elem)
		if err != nil {
			this.vnic.Resources().Logger().Error(err.Error())
		}
		this.vnic.Resources().Logger().Info("Sent model to ", job.IService.ServiceName,
			" area ", job.IService.ServiceArea)
	}
}
