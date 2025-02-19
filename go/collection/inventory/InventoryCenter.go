package inventory

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
)

type InventoryCenter struct {
	boxes *cache.Cache
}

func newInventoryCenter(resources interfaces.IResources, listener cache.ICacheListener) *InventoryCenter {
	this := &InventoryCenter{}
	this.boxes = cache.NewModelCache(resources.Config().LocalUuid, listener, resources.Introspector())
	node, _ := resources.Introspector().Inspect(&types.NetworkBox{})
	resources.Introspector().AddDecorator(types2.DecoratorType_Primary, []string{"Id"}, node)
	return this
}

func (this *InventoryCenter) Add(box *types.NetworkBox) {
	this.boxes.Put(box.Id, box)
}

func (this *InventoryCenter) Update(box *types.NetworkBox) {
	this.boxes.Update(box.Id, box)
}

func (this *InventoryCenter) BoxById(id string) *types.NetworkBox {
	box, _ := this.boxes.Get(id).(*types.NetworkBox)
	return box
}

func Inventory(resource interfaces.IResources) *InventoryCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(TOPIC)
	if !ok {
		return nil
	}
	return (sp.(*InventoryServicePoint)).inventoryCenter
}
