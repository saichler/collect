package tests

import (
	"encoding/base64"
	"fmt"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/parsing"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/probler/go/serializers"
	types3 "github.com/saichler/probler/go/types"
	"github.com/saichler/l8types/go/ifs"
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
	resources ifs.IResources
	expected  int
	received  int
	cond      *sync.Cond
	ph        *DirectParsingHandler
	area      int32
}

func activateDeviceAndPollConfigServices(vnic ifs.IVNic, serviceArea uint16, polls []*types.PollConfig, args ...interface{}) {
	vnic.Resources().Services().RegisterServiceHandlerType(&device_config.DeviceConfigService{})
	vnic.Resources().Services().RegisterServiceHandlerType(&poll_config.PollConfigService{})
	vnic.Resources().Services().Activate(device_config.ServiceType, device_config.ServiceName,
		serviceArea, vnic.Resources(), vnic, args...)
	vnic.Resources().Services().Activate(poll_config.ServiceType, poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
	pc := poll_config.PollConfig(vnic.Resources())
	pc.AddAll(polls)
}

func deActivateDeviceAndPollConfigServices(vnic ifs.IVNic, serviceArea uint16) {
	vnic.Resources().Services().DeActivate(device_config.ServiceName, serviceArea, vnic.Resources(), vnic)
	vnic.Resources().Services().DeActivate(poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
}

func activateParsingAndPollConfigServices(vnic ifs.IVNic,
	pService *types.DeviceServiceInfo, elem interface{}, primaryKey string, polls []*types.PollConfig) {
	vnic.Resources().Services().RegisterServiceHandlerType(&parsing.ParsingService{})
	vnic.Resources().Services().RegisterServiceHandlerType(&poll_config.PollConfigService{})
	_, err := vnic.Resources().Services().Activate(parsing.ServiceType, pService.ServiceName,
		uint16(pService.ServiceArea), vnic.Resources(), vnic, elem, primaryKey)
	if err != nil {
		panic(err)
	}
	_, err = vnic.Resources().Services().Activate(poll_config.ServiceType, poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
	if err != nil {
		panic(err)
	}

	vnic.Resources().Registry().RegisterEnums(types3.NodeStatus_value)
	vnic.Resources().Registry().RegisterEnums(types3.PodStatus_value)
	info, _ := vnic.Resources().Registry().Info("ReadyState")
	if info != nil {
		info.AddSerializer(&serializers.Ready{})
	}
	pc := poll_config.PollConfig(vnic.Resources())
	pc.AddAll(polls)
}

func deActivateParsingAndPollConfigServices(vnic ifs.IVNic, pService *types.DeviceServiceInfo) {
	vnic.Resources().Services().DeActivate(pService.ServiceName, uint16(pService.ServiceArea), vnic.Resources(), vnic)
	vnic.Resources().Services().DeActivate(poll_config.ServiceName, poll_config.ServiceArea, vnic.Resources(), vnic)
}

func activateInventoryService(vnic ifs.IVNic, iService *types.DeviceServiceInfo,
	elem interface{}, primaryKey string) {
	vnic.Resources().Services().RegisterServiceHandlerType(&inventory.InventoryService{})
	vnic.Resources().Services().Activate(inventory.ServiceType, iService.ServiceName,
		uint16(iService.ServiceArea), vnic.Resources(), vnic, primaryKey, elem)
}

func deActivateInventoryService(vnic ifs.IVNic, iService *types.DeviceServiceInfo) {
	vnic.Resources().Services().DeActivate(iService.ServiceName, uint16(iService.ServiceArea), vnic.Resources(), vnic)
}

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
	device.InventoryService = &types.DeviceServiceInfo{ServiceName: K8sServiceName, ServiceArea: serviceArea}
	device.ParsingService = &types.DeviceServiceInfo{ServiceName: K8sServiceName + "P", ServiceArea: serviceArea}
	device.Hosts = make(map[string]*types.HostConfig)
	host := &types.HostConfig{}
	host.DeviceId = device.DeviceId

	host.Configs = make(map[int32]*types.ConnectionConfig)
	device.Hosts[device.DeviceId] = host

	k8sConfig := &types.ConnectionConfig{}
	data, err := os.ReadFile(kubeconfig)
	if err != nil {
		panic(err)
	}
	k8sConfig.KubeConfig = base64.StdEncoding.EncodeToString(data)
	k8sConfig.KukeContext = context
	k8sConfig.Protocol = types.Protocol_K8s

	host.Configs[int32(k8sConfig.Protocol)] = k8sConfig

	return device
}
