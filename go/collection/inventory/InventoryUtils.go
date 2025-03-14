package inventory

import (
	"github.com/saichler/types/go/common"
	"reflect"
	"strings"
)

func primaryKeyValue(attr string, any interface{}, resources common.IResources) string {
	if any == nil {
		resources.Logger().Error("element is nil")
		return ""
	}
	v := reflect.ValueOf(any)
	if v.Kind() != reflect.Ptr {
		resources.Logger().Error("element is not a pointer")
		return ""
	}
	field := v.Elem().FieldByName(attr)
	if !field.IsValid() {
		resources.Logger().Error("attribute " + attr + " not found")
		return ""
	}
	return field.String()
}

func (this *InventoryCenter) setTopic(any interface{}, resources common.IResources) {
	if any == nil {
		resources.Logger().Error("element is nil")
		return
	}
	v := reflect.ValueOf(any)
	if v.Kind() != reflect.Ptr {
		resources.Logger().Error("element is not a pointer")
		return
	}
	this.elemType = v.Elem().Type()
	TOPIC = v.Elem().Type().Name()
	ENDPOINT = strings.ToLower(TOPIC)
}

func (this *InventoryCenter) AddEmpty(key string) {
	elem := reflect.New(this.elemType)
	field := elem.Elem().FieldByName(this.primaryKeyAttribute)
	field.Set(reflect.ValueOf(key))
	this.Add(elem.Interface())
}
