package inventory

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"reflect"
	"strings"
)

type InventoryCenter struct {
	elements            *cache.Cache
	elementType         reflect.Type
	primaryKeyAttribute string
	resources           common.IResources
	serviceName         string
	serviceArea         int32
	element             interface{}
}

func newInventoryCenter(serviceName string, serviceArea int32, primaryKeyAttribute string,
	element interface{}, resources common.IResources, listener cache.ICacheListener) *InventoryCenter {
	this := &InventoryCenter{}
	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.element = element
	this.elementType = reflect.ValueOf(element).Elem().Type()
	this.resources = resources
	this.primaryKeyAttribute = primaryKeyAttribute
	this.elements = cache.NewModelCache(this.serviceName, this.serviceArea, this.elementType.Name(),
		resources.Config().LocalUuid, listener, resources.Introspector())
	node, _ := resources.Introspector().Inspect(element)
	resources.Introspector().AddDecorator(types2.DecoratorType_Primary, []string{primaryKeyAttribute}, node)
	return this
}

func (this *InventoryCenter) Add(elem interface{}) {
	key := primaryKeyValue(this.primaryKeyAttribute, elem, this.resources)
	if key != "" {
		this.elements.Put(key, elem)
	}
}

func (this *InventoryCenter) Update(elem interface{}) {
	key := primaryKeyValue(this.primaryKeyAttribute, elem, this.resources)
	if key != "" {
		this.elements.Update(key, elem)
	}
}

func (this *InventoryCenter) ElementByKey(key string) interface{} {
	return this.elements.Get(key)
}

func removeParsingSuffix(serviceName string) string {
	index := strings.Index(serviceName, base.Parsing_Suffix)
	if index == -1 {
		return serviceName
	}
	return serviceName[:index]
}

func Inventory(resource common.IResources, serviceName string, serviceArea int32) *InventoryCenter {
	serviceName = removeParsingSuffix(serviceName)
	sp, ok := resource.ServicePoints().ServicePointHandler(serviceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*InventoryServicePoint)).inventoryCenter
}
