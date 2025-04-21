package inventory

import (
	"github.com/saichler/reflect/go/reflect/introspecting"
	"github.com/saichler/servicepoints/go/points/dcache"
	"github.com/saichler/types/go/common"
	"reflect"
)

type InventoryCenter struct {
	elements            common.IDistributedCache
	elementType         reflect.Type
	primaryKeyAttribute string
	resources           common.IResources
	serviceName         string
	serviceArea         uint16
	element             interface{}
}

func newInventoryCenter(serviceName string, serviceArea uint16, primaryKeyAttribute string,
	element interface{}, resources common.IResources, listener common.IServicePointCacheListener) *InventoryCenter {
	this := &InventoryCenter{}
	this.serviceName = serviceName
	this.serviceArea = serviceArea
	this.element = element
	this.elementType = reflect.ValueOf(element).Elem().Type()
	this.resources = resources
	this.primaryKeyAttribute = primaryKeyAttribute
	this.elements = dcache.NewDistributedCache(this.serviceName, this.serviceArea, this.elementType.Name(),
		resources.SysConfig().LocalUuid, listener, resources.Introspector())
	node, _ := resources.Introspector().Inspect(element)
	introspecting.AddPrimaryKeyDecorator(node, primaryKeyAttribute)
	return this
}

func (this *InventoryCenter) Add(elem interface{}, isNotification bool) {
	_, ok := elem.(common.IElements)
	if ok {
		panic("")
	}
	key := primaryKeyValue(this.primaryKeyAttribute, elem, this.resources)
	if key != "" {
		this.elements.Put(key, elem, isNotification)
	}
}

func (this *InventoryCenter) Update(elem interface{}, isNotification bool) {
	key := primaryKeyValue(this.primaryKeyAttribute, elem, this.resources)
	if key != "" {
		this.elements.Update(key, elem, isNotification)
	}
}

func (this *InventoryCenter) ElementByKey(key string) interface{} {
	return this.elements.Get(key)
}

func (this *InventoryCenter) ElementByElement(elem interface{}) interface{} {
	key := primaryKeyValue(this.primaryKeyAttribute, elem, this.resources)
	return this.elements.Get(key)
}

func Inventory(resource common.IResources, serviceName string, serviceArea uint16) *InventoryCenter {
	//serviceName = serviceName
	sp, ok := resource.ServicePoints().ServicePointHandler(serviceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*InventoryServicePoint)).inventoryCenter
}
