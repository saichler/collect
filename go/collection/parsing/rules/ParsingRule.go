package rules

import (
	"errors"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/shared/go/share/interfaces"
	"strconv"
)

type ParsingRule interface {
	Name() string
	ParamNames() []string
	Parse(interfaces.IResources, map[string]interface{}, map[string]*types.Parameter, interface{}) error
}

func getStringInput(resources interfaces.IResources, input interface{}, params map[string]*types.Parameter) (string, error) {
	m, ok := input.(*types.Map)
	if ok {
		from := params[From]
		if from == nil {
			return "", resources.Logger().Error("missing 'from' key in map input")
		}
		strData := m.Data[from.Value]
		enc := object.NewDecode(strData, 0, "Map", resources.Registry())
		strInt, _ := enc.Get()
		str, ok := strInt.(string)
		if ok {
			return str, nil
		}
		byts, ok := strInt.([]byte)
		if ok {
			return string(byts), nil
		}
		return "", resources.Logger().Error("'from' key not a string")
	}
	byts, ok := input.([]byte)
	if ok {
		return string(byts), nil
	}
	return "", resources.Logger().Error("'from' key not a []byte")
}

func getIntInput(workSpace map[string]interface{}, paramName string) (int, error) {
	v, ok := workSpace[paramName].(string)
	if !ok {
		return -1, errors.New("'" + paramName + "' does not exist")
	}
	i, e := strconv.Atoi(v)
	if e != nil {
		return -1, e
	}
	return i, nil
}
