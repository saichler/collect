package rules

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/reflect/go/reflect/property"
	"github.com/saichler/shared/go/share/interfaces"
	"strings"
)

type Contains struct{}

func (this *Contains) Name() string {
	return "Contains"
}

func (this *Contains) ParamNames() []string {
	return []string{"what"}
}

func (this *Contains) Parse(resources interfaces.IResources, workSpace map[string]interface{}, params map[string]*types.Parameter, any interface{}) error {
	input := workSpace[Input]
	what := params[What]
	output := params[Output]
	path := workSpace[PropertyId]

	if input == nil {
		return resources.Logger().Error("nil input for job")
	}
	if what == nil {
		return resources.Logger().Error("nil 'what' parameter")
	}
	if output == nil {
		return resources.Logger().Error("Nil 'output' parameter")
	}
	str, err := getStringInput(resources, input, params)
	if err != nil {
		return err
	}
	ok := strings.Contains(strings.ToLower(str), what.Value)
	if ok {
		if path != nil {
			instance, _ := property.PropertyOf(path.(string), resources.Introspector())
			if instance != nil {
				_, _, err := instance.Set(any, output.Value)
				if err != nil {
					return err
				}
			}
		}
		workSpace[Output] = output.Value
	}
	return nil
}
