package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/layer8/go/overlay/vnet"
	vnic2 "github.com/saichler/layer8/go/overlay/vnic"
	"github.com/saichler/reflect/go/reflect/inspect"
	"github.com/saichler/servicepoints/go/points/service_points"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/share/registry"
	"github.com/saichler/shared/go/share/resources"
	"github.com/saichler/shared/go/share/shallow_security"
	"github.com/saichler/shared/go/types"
)

const (
	PORT = 30000
)

func createSwitch() *vnet.VNet {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	config := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "vnet",
		Topics:      map[string]bool{}}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, config)

	res := resources.NewResources(reg, security, sps, log, nil, nil, config, ins)
	res.Config().SwitchPort = PORT
	sw := vnet.NewVNet(res)
	sw.Start()
	return sw
}

func createCollectionService() interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "collector",
		Topics:      map[string]bool{}}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().SwitchPort = PORT

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)

	l := control.NewParsingCenterNotifier(vnic)
	controller := control.NewController(l, resourcs)

	config.RegisterConfigCenter(resourcs, nil, controller)
	polling.RegisterPollCenter(resourcs, nil)

	vnic.Start()

	return vnic
}

func createParsingService() interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "parsing",
		Topics:      map[string]bool{}}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().SwitchPort = PORT

	polling.RegisterPollCenter(resourcs, nil)
	parsing.RegisterParsingServicePoint(resourcs)

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)
	vnic.Start()

	return vnic
}

func createClient() interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "parsing",
		Topics:      map[string]bool{}}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().SwitchPort = PORT

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)
	vnic.Start()

	return vnic
}
