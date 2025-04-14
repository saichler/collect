package tests

import (
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"github.com/saichler/collect/go/types"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestParsingAndInventory(t *testing.T) {
	ip := "192.168.86.179"
	device := CreateDevice(ip, 0)
	cfg := topo.VnicByVnetNum(2, 4)
	par := topo.VnicByVnetNum(3, 1)
	inv := topo.VnicByVnetNum(1, 3)

	cont := control.NewController(control.NewParsingCenterNotifier(cfg), cfg.Resources())
	activateDeviceAndPollConfigServices(cfg, 0, cont, boot.CreateSNMPBootPolls())
	activateParsingAndPollConfigServices(par, device.ParsingService,
		&types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	activateInventoryService(inv, device.InventoryService, &types.NetworkBox{}, "Id")
	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
		deActivateParsingAndPollConfigServices(par, device.ParsingService)
		deActivateInventoryService(inv, device.InventoryService)
	}()
	sleep()

	Log.Info("Test Multicast")
	cli := topo.VnicByVnetNum(1, 2)
	cli.Multicast(device_config.ServiceName, 0, common.POST, device)

	ic := inventory.Inventory(inv.Resources(), device.InventoryService.ServiceName, uint16(device.InventoryService.ServiceArea))

	var box *types.NetworkBox
	var ok bool
	for i := 0; i < 10; i++ {
		box, ok = ic.ElementByKey(ip).(*types.NetworkBox)
		if ok {
			break
		}
		time.Sleep(time.Second)
	}

	if box == nil {
		Log.Fail(t, "Expected box to be non-nil")
		return
	}

	if box.Info == nil {
		Log.Fail(t, "Expected box info to be non-nil")
		return
	}

	if box.Info.SysName == "" {
		Log.Fail(t, "Expected box info sysname to not be blank")
		return
	}

	if box.Info.Vendor == "" {
		Log.Fail(t, "Expected box vendor to not be blank")
		return
	}
}
