package inventory

import (
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/reflect/go/reflect/introspecting"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	"reflect"
	"strings"
)

type InventoryCenter struct {
	elements            *cache.Cache
	elementType         reflect.Type
	primaryKeyAttribute string
	resources           common.IResources
	serviceName         string
	serviceArea         uint16
	element             interface{}
}

func newInventoryCenter(serviceName string, serviceArea uint16, primaryKeyAttribute string,
	element common.IElements, resources common.IResources, listener common.IServicePointCacheListener) *InventoryCenter {
	this := &InventoryCenter{}
	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.element = element
	this.elementType = reflect.ValueOf(element.Element()).Elem().Type()
	this.resources = resources
	this.primaryKeyAttribute = primaryKeyAttribute
	this.elements = cache.NewModelCache(this.serviceName, this.serviceArea, this.elementType.Name(),
		resources.SysConfig().LocalUuid, listener, resources.Introspector())
	node, _ := resources.Introspector().Inspect(element.Element())
	introspecting.AddPrimaryKeyDecorator(node, primaryKeyAttribute)
	return this
}

func (this *InventoryCenter) Add(elem interface{}) {
	_, ok := elem.(common.IElements)
	if ok {
		panic("")
	}
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

func Inventory(resource common.IResources, serviceName string, serviceArea uint16) *InventoryCenter {
	serviceName = removeParsingSuffix(serviceName)
	sp, ok := resource.ServicePoints().ServicePointHandler(serviceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*InventoryServicePoint)).inventoryCenter
}
