package inventory

import (
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"reflect"
)

type InventoryCenter struct {
	elements            *cache.Cache
	elemType            reflect.Type
	primaryKeyAttribute string
	resources           common.IResources
}

func newInventoryCenter(primaryKeyAttribute string, element interface{}, resources common.IResources, listener cache.ICacheListener) *InventoryCenter {
	this := &InventoryCenter{}
	this.resources = resources
	this.primaryKeyAttribute = primaryKeyAttribute
	this.elements = cache.NewModelCache(resources.Config().LocalUuid, listener, resources.Introspector())
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

func Inventory(resource common.IResources) *InventoryCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(TOPIC)
	if !ok {
		return nil
	}
	return (sp.(*InventoryServicePoint)).inventoryCenter
}
