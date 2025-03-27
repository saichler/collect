package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/collect/go/types"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/types/go/common"
	types2 "github.com/saichler/types/go/types"
	"testing"
	"time"
)

func TestOneCollectorTwoParsers(t *testing.T) {
	sw := createVNet(vNetPort1)
	sleep()
	col := createCollectionService(0, vNetPort1, boot.CreateSNMPBootPolls())
	sleep()
	par1 := createParsingService(0, vNetPort1, &types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	par2 := createParsingService(1, vNetPort1, &types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	sleep()
	cli := createClient(vNetPort1)
	sleep()

	defer func() {
		cli.Shutdown()
		par1.Shutdown()
		par2.Shutdown()
		col.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	ip1 := "192.168.86.198"
	ip2 := "192.168.86.179"

	//assign device 1 to parser in area 0
	device1 := CreateDevice(ip1, 0)
	//assign device 2 to parser in area 1
	device2 := CreateDevice(ip2, 1)
	cli.Multicast(config.ServiceName, 0, types2.Action_POST, device1)
	cli.Multicast(config.ServiceName, 0, types2.Action_POST, device2)

	time.Sleep(2 * time.Second)

	if !checkInventory(ip1, par1.Resources(), t, 0) {
		return
	}
	if !checkInventory(ip2, par2.Resources(), t, 1) {
		return
	}
}

func checkInventory(ip string, resours common.IResources, t *testing.T, serviceArea int32) bool {
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
