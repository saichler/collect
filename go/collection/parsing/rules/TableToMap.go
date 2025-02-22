package rules

import (
	"errors"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/reflect/go/reflect/property"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/shared/go/share/interfaces"
	strings2 "github.com/saichler/shared/go/share/strings"
	"reflect"
	"strings"
)

type TableToMap struct{}

func (this *TableToMap) Name() string {
	return "TableToMap"
}

func (this *TableToMap) ParamNames() []string {
	return []string{""}
}

func (this *TableToMap) Parse(resources interfaces.IResources, workSpace map[string]interface{}, params map[string]*types.Parameter, any interface{}) error {
	table, ok := workSpace[Output].(*types.Table)
	if !ok {
		return errors.New("Workspace had an invalid output object")
	}

	propertyId := workSpace[PropertyId].(string)
	toString := strings2.New()
	toString.TypesPrefix = true

	for _, row := range table.Rows {
		pid := strings2.New(propertyId)
		for i := 0; i < len(table.Columns); i++ {
			if i == 0 {
				if len(row.Data[0]) == 0 {
					break
				}
				val := getValue(row.Data[0], resources)
				pid.Add("<")
				pid.Add(toString.ToString(reflect.ValueOf(val)))
				pid.Add(">.")
			}

			key := strings2.New(pid.String())
			attrName := getAttributeNameFromColumn(table.Columns[int32(i)])
			key.Add(attrName)
			
			prop, err := property.PropertyOf(key.String(), resources.Introspector())
			if err != nil {
				resources.Logger().Error(err.Error())
				continue
			}

			data := row.Data[int32(i)]
			obj := object.NewDecode(data, 0, "", resources.Registry())
			val, err := obj.Get()

			if err != nil {
				resources.Logger().Error(err.Error())
				continue
			}
			_, _, err = prop.Set(any, val)
			if err != nil {
				resources.Logger().Error(err.Error())
				continue
			}
		}
	}
	return nil
}

func getAttributeNameFromColumn(value interface{}) string {
	colName := strings.TrimSpace(value.(string))
	colName = strings.ToLower(colName)
	index := strings.LastIndex(colName, "-")
	if index == -1 {
		return colName
	}
	return strings2.New(colName[0:index], colName[index+1:]).String()
}

func getValue(data []byte, resources interfaces.IResources) interface{} {
	obj := object.NewDecode(data, 0, "", resources.Registry())
	val, _ := obj.Get()
	return val
}
