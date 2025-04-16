package tests

import (
	"github.com/saichler/collect/go/collection/collector"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"github.com/saichler/collect/go/types"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestOneCollectorTwoParsers(t *testing.T) {
	ip2 := "192.168.86.179"
	ip1 := "192.168.86.198"

	device1 := CreateDevice(ip1, 0)
	device2 := CreateDevice(ip2, 1)

	cfg := topo.VnicByVnetNum(2, 4)
	par1 := topo.VnicByVnetNum(3, 1)
	par2 := topo.VnicByVnetNum(1, 1)
	inv1 := topo.VnicByVnetNum(1, 3)
	inv2 := topo.VnicByVnetNum(2, 3)

	cont := collector.NewDeviceCollector(collector.NewParsingCenterNotifier(cfg), cfg.Resources())
	activateDeviceAndPollConfigServices(cfg, 0, cont, boot.CreateSNMPBootPolls())

	activateParsingAndPollConfigServices(par1, device1.ParsingService,
		&types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	activateParsingAndPollConfigServices(par2, device2.ParsingService,
		&types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())

	activateInventoryService(inv1, device1.InventoryService, &types.NetworkBox{}, "Id")
	activateInventoryService(inv2, device2.InventoryService, &types.NetworkBox{}, "Id")

	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
		deActivateParsingAndPollConfigServices(par1, device1.ParsingService)
		deActivateParsingAndPollConfigServices(par2, device2.ParsingService)
		deActivateInventoryService(inv1, device1.InventoryService)
		deActivateInventoryService(inv2, device2.InventoryService)
	}()
	sleep()

	cli := topo.VnicByVnetNum(1, 2)
	cli.Multicast(device_config.ServiceName, 0, common.POST, device1)
	cli.Multicast(device_config.ServiceName, 0, common.POST, device2)

	ic := inventory.Inventory(inv2.Resources(), device2.InventoryService.ServiceName, uint16(device2.InventoryService.ServiceArea))

	for i := 0; i < 10; i++ {
		_, ok := ic.ElementByKey(ip2).(*types.NetworkBox)
		if ok {
			break
		}
		time.Sleep(time.Second)
	}

	if !checkInventory(ip1, inv1.Resources(), t, 0) {
		return
	}
	if !checkInventory(ip2, inv2.Resources(), t, 1) {
		return
	}
}

func checkInventory(ip string, resours common.IResources, t *testing.T, serviceArea uint16) bool {
	ic := inventory.Inventory(resours, "NetworkBox", serviceArea)
	box, _ := ic.ElementByKey(ip).(*types.NetworkBox)
	if box == nil {
		Log.Fail(t, ip, " Expected box to be non-nil")
		return false
	}

	if box.Info == nil {
		Log.Fail(t, ip, " Expected box info to be non-nil")
		return false
	}

	if box.Info.SysName == "" {
		Log.Fail(t, ip, " Expected box info sysname to not be blank")
		return false
	}

	if box.Info.Vendor == "" {
		Log.Fail(t, ip, " Expected box vendor to not be blank")
		return false
	}
	return true
}
