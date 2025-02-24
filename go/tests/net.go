package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/collection/polling"
	types2 "github.com/saichler/collect/go/types"
	"github.com/saichler/layer8/go/overlay/vnet"
	vnic2 "github.com/saichler/layer8/go/overlay/vnic"
	"github.com/saichler/reflect/go/reflect/inspect"
	"github.com/saichler/servicepoints/go/points/service_points"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/share/registry"
	"github.com/saichler/shared/go/share/resources"
	"github.com/saichler/shared/go/share/shallow_security"
	"github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	PORT = 30000
)

func createVNet() *vnet.VNet {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	config := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "vnet"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, config)

	res := resources.NewResources(reg, security, sps, log, nil, nil, config, ins)
	res.Config().VnetPort = PORT
	sw := vnet.NewVNet(res)
	sw.Start()
	return sw
}

func createCollectionService(polls []*types2.Poll) interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "collector"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = PORT

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)

	l := control.NewParsingCenterNotifier(vnic)
	controller := control.NewController(l, resourcs)

	config.RegisterConfigCenter(cfg.Area, resourcs, nil, controller)
	polling.RegisterPollCenter(cfg.Area, resourcs, nil)
	pc := polling.Polling(resourcs)
	pc.AddAll(polls)

	vnic.Start()

	return vnic
}

func createParsingService(area int32, pb proto.Message, key string, polls []*types2.Poll) interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "parsing",
		Area:        area}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = PORT

	polling.RegisterPollCenter(cfg.Area, resourcs, nil)
	pc := polling.Polling(resourcs)
	pc.AddAll(polls)
	parsing.RegisterParsingServicePoint(cfg.Area, pb, key, resourcs)

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
		LocalAlias:  "parsing"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = PORT

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)
	vnic.Start()

	return vnic
}
