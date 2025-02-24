package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
	"testing"
	"time"
)

func TestOneCollectorTwoParsers(t *testing.T) {

	sw := createVNet()
	sleep()
	col := createCollectionService(boot.CreateSNMPBootPolls())
	sleep()
	par1 := createParsingService(0, &types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	par2 := createParsingService(1, &types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	sleep()
	cli := createClient()
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

	cli.Multicast(types2.Action_POST, 0, config.TOPIC, device1)
	cli.Multicast(types2.Action_POST, 0, config.TOPIC, device2)

	time.Sleep(2 * time.Second)

	if !checkInventory(ip1, par1.Resources(), t) {
		return
	}
	if !checkInventory(ip2, par2.Resources(), t) {
		return
	}
}

func checkInventory(ip string, resours interfaces.IResources, t *testing.T) bool {
	ic := inventory.Inventory(resours)
	box := ic.ElementByKey(ip).(*types.NetworkBox)
	if box == nil {
		log.Fail(t, ip, " Expected box to be non-nil")
		return false
	}

	if box.Info == nil {
		log.Fail(t, ip, " Expected box info to be non-nil")
		return false
	}

	if box.Info.SysName == "" {
		log.Fail(t, ip, " Expected box info sysname to not be blank")
		return false
	}

	if box.Info.Vendor == "" {
		log.Fail(t, ip, " Expected box vendor to not be blank")
		return false
	}
	return true
}
