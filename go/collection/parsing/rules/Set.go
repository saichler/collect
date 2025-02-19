package rules

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/reflect/go/reflect/property"
	"github.com/saichler/shared/go/share/interfaces"
)

type Set struct{}

func (this *Set) Name() string {
	return "Set"
}

func (this *Set) ParamNames() []string {
	return []string{}
}

func (this *Set) Parse(resources interfaces.IResources, workSpace map[string]interface{}, params map[string]*types.Parameter, any interface{}) error {
	input := workSpace["input"]
	path := workSpace["InstancePath"]

	if input == nil {
		return resources.Logger().Error("nil input for job")
	}

	str, err := getStringInput(resources, input, params)
	if err != nil {
		return err
	}

	if path != nil {
		instance, err := property.PropertyOf(path.(string), resources.Introspector())
		if err != nil {
			return resources.Logger().Error("error parsing instance path", err.Error())
		}
		if instance != nil {
			instance.Set(any, str)
		}
	}
	workSpace["output"] = str
	return nil
}
