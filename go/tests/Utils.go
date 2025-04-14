package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/types"
	types3 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"os"
	"sync"
	"time"
)

var home, _ = os.LookupEnv("HOME")
var admin1 = home + "/admin.conf"
var context1 = "kubernetes-admin@kubernetes"
var admin2 = home + "/lab.conf"
var context2 = "lab"

const (
	InvServiceName = "NetworkBox"
	K8sServiceName = "Cluster"
)

func sleep() {
	time.Sleep(time.Millisecond * 100)
}

type CollectorListener struct {
	resources common.IResources
	expected  int
	received  int
	cond      *sync.Cond
	ph        *control.DirectParsingHandler
	area      int32
}

func activateDeviceAndPollConfigServices(vnic common.IVirtualNetworkInterface,
	controller *control.Controller, polls []*types.PollConfig) {
	vnic.Resources().ServicePoints().AddServicePointType(&device_config.DeviceConfigServicePoint{})
	vnic.Resources().ServicePoints().AddServicePointType(&poll_config.PollConfigServicePoint{})
	vnic.Resources().ServicePoints().Activate(device_config.ServicePointType, device_config.ServiceName,
		device_config.ServiceArea, vnic.Resources(), vnic, controller)
	vnic.Resources().ServicePoints().Activate(poll_config.ServicePointType, poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
	pc := poll_config.PollConfig(vnic.Resources())
	pc.AddAll(polls)
}

func deActivateDeviceAndPollConfigServices(vnic common.IVirtualNetworkInterface) {
	vnic.Resources().ServicePoints().DeActivate(device_config.ServiceName, device_config.ServiceArea, vnic.Resources(), vnic)
	vnic.Resources().ServicePoints().DeActivate(poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
}

func activateParsingAndPollConfigServices(vnic common.IVirtualNetworkInterface,
	pService *types.DeviceServiceInfo, elem interface{}, primaryKey string, polls []*types.PollConfig) {
	vnic.Resources().ServicePoints().AddServicePointType(&parsing.ParsingServicePoint{})
	vnic.Resources().ServicePoints().AddServicePointType(&poll_config.PollConfigServicePoint{})
	vnic.Resources().ServicePoints().Activate(parsing.ServicePointType, pService.ServiceName,
		uint16(pService.ServiceArea), vnic.Resources(), vnic, elem, primaryKey)
	vnic.Resources().ServicePoints().Activate(poll_config.ServicePointType, poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)

	vnic.Resources().Registry().RegisterEnums(types3.NodeStatus_value)
	vnic.Resources().Registry().RegisterEnums(types3.PodStatus_value)

	pc := poll_config.PollConfig(vnic.Resources())
	pc.AddAll(polls)
}

func deActivateParsingAndPollConfigServices(vnic common.IVirtualNetworkInterface, pService *types.DeviceServiceInfo) {
	vnic.Resources().ServicePoints().DeActivate(pService.ServiceName, uint16(pService.ServiceArea), vnic.Resources(), vnic)
	vnic.Resources().ServicePoints().DeActivate(poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
}

func activateInventoryService(vnic common.IVirtualNetworkInterface, iService *types.DeviceServiceInfo,
	elem interface{}, primaryKey string) {
	vnic.Resources().ServicePoints().AddServicePointType(&inventory.InventoryServicePoint{})
	vnic.Resources().ServicePoints().Activate(inventory.ServicePointType, iService.ServiceName,
		uint16(iService.ServiceArea), vnic.Resources(), vnic, primaryKey, elem)
}

func deActivateInventoryService(vnic common.IVirtualNetworkInterface, iService *types.DeviceServiceInfo) {
	vnic.Resources().ServicePoints().DeActivate(iService.ServiceName, uint16(iService.ServiceArea), vnic.Resources(), vnic)
}

/*
func createResources(alias string) common.IResources {
	reg := registry.NewRegistry()
	secure, err := common.LoadSecurityProvider("security.so", "../../../")
	if err != nil {
		panic("Failed to load security provider")
	}
	cfg := &types2.SysConfig{MaxDataSize: resources.DEFAULT_MAX_DATA_SIZE,
		RxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		TxQueueSize: resources.DEFAULT_QUEUE_SIZE,
		LocalAlias:  alias}
	ins := introspecting.NewIntrospect(reg)
	sps := service_points.NewServicePoints(ins, cfg)

	ress := resources.NewResources(reg, secure, sps, Log, nil, nil, cfg, ins)
	ress.ServicePoints().AddServicePointType(&device_config.DeviceConfigServicePoint{})
	ress.ServicePoints().AddServicePointType(&poll_config.PollConfigServicePoint{})
	return ress
}*/

func (l *CollectorListener) JobCompleted(job *types.Job) {
	if l.ph != nil {
		l.ph.HandleCollectNotification(job)
	}
	pc := poll_config.PollConfig(l.resources)
	poll := pc.PollByName(job.PollName)
	if poll == nil {
		l.resources.Logger().Error("cannot find poll for uuid ", job.PollName)
		return
	}
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.received++
	var result interface{}
	/*
		if poll.Name == "version" ||
			poll.Name == "clock" ||
			poll.Name == "timezone" ||
			poll.Name == "te-tunnel-id" {
			result = string(job.Result)
		} else {
			enc := object.NewDecode(job.Result, 0, "", l.resources.Registry())
			val, _ := enc.Get()
			result = val
			m := val.(*types.Map)
			for key, value := range m.Data {
				enc = object.NewDecode(value, 0, "", l.resources.Registry())
				v, _ := enc.Get()
				str, ok := v.(string)
				if !ok {
					byts, ok := v.([]byte)
					if ok {
						str = string(byts)
					}
				}
				if str != "" {
					fmt.Println("key:", key, " value:", str)
				} else {
					fmt.Println("key:", key, " value:", v)
				}
			}
		}*/
	fmt.Println(poll.Name, ":", result)
	if l.received >= l.expected {
		l.cond.Broadcast()
	}
}

/*
func CreateCommands() ([]*model.CollectCommand, map[string]string) {
	cVersion := &model.CollectCommand{}
	cVersion.Id = "version"
	cVersion.Enabled = true
	cVersion.Cadence = 300
	cVersion.Operation = model.CollectOperation_Get
	cVersion.Protocol = model.CollectProtocol_SSH
	cVersion.What = "show version"

	cSystem := &model.CollectCommand{}
	cSystem.Id = "system"
	cSystem.Enabled = true
	cSystem.Cadence = 300
	cSystem.Operation = model.CollectOperation_Map
	cSystem.Protocol = model.CollectProtocol_SNMPV2
	cSystem.What = ".1.3.6.1.2.1.1"

	cClock := &model.CollectCommand{}
	cClock.Id = "clock"
	cClock.Enabled = true
	cClock.Cadence = 900
	cClock.Operation = model.CollectOperation_Get
	cClock.Protocol = model.CollectProtocol_SSH
	cClock.What = "show clock"

	cTimezone := &model.CollectCommand{}
	cTimezone.Id = "timezone"
	cTimezone.Enabled = true
	cTimezone.Cadence = 900
	cTimezone.Operation = model.CollectOperation_Get
	cTimezone.Protocol = model.CollectProtocol_SSH
	cTimezone.What = "show running-config | inc clock"

	cTeTunnelId := &model.CollectCommand{}
	cTeTunnelId.Id = "te-tunnel-id"
	cTeTunnelId.Enabled = true
	cTeTunnelId.Cadence = 900
	cTeTunnelId.Operation = model.CollectOperation_Get
	cTeTunnelId.Protocol = model.CollectProtocol_SSH
	cTeTunnelId.What = "show mpls traffic-eng igp-areas"
	m := make(map[string]string)
	m[cSystem.Id] = cSystem.Id
	m[cVersion.Id] = cVersion.Id
	m[cClock.Id] = cClock.Id
	m[cTimezone.Id] = cTimezone.Id
	m[cTeTunnelId.Id] = cTeTunnelId.Id

	return []*model.CollectCommand{cVersion, cSystem, cClock, cTimezone, cTeTunnelId}, m
}*/

func CreateDevice(ip string, serviceArea uint16) *types.DeviceConfig {
	device := &types.DeviceConfig{}
	device.DeviceId = ip
	device.InventoryService = &types.DeviceServiceInfo{ServiceName: InvServiceName, ServiceArea: int32(serviceArea)}
	device.ParsingService = &types.DeviceServiceInfo{ServiceName: InvServiceName + "P", ServiceArea: int32(serviceArea)}
	device.Hosts = make(map[string]*types.HostConfig)
	host := &types.HostConfig{}
	host.DeviceId = device.DeviceId

	host.Configs = make(map[int32]*types.ConnectionConfig)
	device.Hosts[device.DeviceId] = host

	sshConfig := &types.ConnectionConfig{}
	sshConfig.Protocol = types.Protocol_SSH
	sshConfig.Port = 22
	sshConfig.Addr = ip
	sshConfig.Username = "admin"
	sshConfig.Password = "admin"
	sshConfig.Terminal = "vt100"
	sshConfig.Timeout = 15

	host.Configs[int32(sshConfig.Protocol)] = sshConfig

	snmpConfig := &types.ConnectionConfig{}
	snmpConfig.Protocol = types.Protocol_SNMPV2
	snmpConfig.Addr = ip
	snmpConfig.Port = 161
	snmpConfig.Timeout = 15
	snmpConfig.ReadCommunity = "public"

	host.Configs[int32(snmpConfig.Protocol)] = snmpConfig

	return device
}

func CreateCluster(kubeconfig, context string, serviceArea int32) *types.DeviceConfig {
	device := &types.DeviceConfig{}
	device.DeviceId = context
	device.InventoryService = &types.DeviceServiceInfo{ServiceName: K8sServiceName, ServiceArea: int32(serviceArea)}
	device.ParsingService = &types.DeviceServiceInfo{ServiceName: K8sServiceName + "P", ServiceArea: int32(serviceArea)}
	device.Hosts = make(map[string]*types.HostConfig)
	host := &types.HostConfig{}
	host.DeviceId = device.DeviceId

	host.Configs = make(map[int32]*types.ConnectionConfig)
	device.Hosts[device.DeviceId] = host

	k8sConfig := &types.ConnectionConfig{}
	k8sConfig.KubeConfig = kubeconfig
	k8sConfig.KukeContext = context
	k8sConfig.Protocol = types.Protocol_K8s

	host.Configs[int32(k8sConfig.Protocol)] = k8sConfig

	return device
}
