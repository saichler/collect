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

var vNetPort1 uint32 = 30000
var vNetPort2 uint32 = 40000

func createVNet(port uint32) *vnet.VNet {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	config := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "vnet"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, config)

	res := resources.NewResources(reg, security, sps, log, nil, nil, config, ins)
	res.Config().VnetPort = port
	sw := vnet.NewVNet(res)
	sw.Start()
	return sw
}

func createCollectionService(vlanId int32, port uint32, polls []*types2.Poll) interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "collector"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = port

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)

	l := control.NewParsingCenterNotifier(vnic)
	controller := control.NewController(l, resourcs)

	config.RegisterConfigCenter(vlanId, resourcs, nil, controller)
	polling.RegisterPollCenter(vlanId, resourcs, nil)
	pc := polling.Polling(resourcs)
	pc.AddAll(polls)

	vnic.Start()

	return vnic
}

func createParsingService(vlanId int32, port uint32, pb proto.Message, key string, polls []*types2.Poll) interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "parsing"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = port

	polling.RegisterPollCenter(vlanId, resourcs, nil)
	pc := polling.Polling(resourcs)
	pc.AddAll(polls)
	parsing.RegisterParsingServicePoint(vlanId, pb, key, resourcs)

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)
	vnic.Start()

	return vnic
}

func createClient(port uint32) interfaces.IVirtualNetworkInterface {
	reg := registry.NewRegistry()
	security := shallow_security.CreateShallowSecurityProvider()
	cfg := &types.VNicConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  "parsing"}
	ins := inspect.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	resourcs := resources.NewResources(reg, security, sps, log, nil, nil, cfg, ins)
	resourcs.Config().VnetPort = port

	vnic := vnic2.NewVirtualNetworkInterface(resourcs, nil)
	vnic.Start()

	return vnic
}
