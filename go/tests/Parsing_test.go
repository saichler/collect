package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/collect/go/types"
	. "github.com/saichler/shared/go/tests/infra"
	types2 "github.com/saichler/types/go/types"
	"testing"
	"time"
)

func TestParsingAndInventory(t *testing.T) {

	sw := createVNet(vNetPort1)
	sleep()
	col := createCollectionService(0, vNetPort1, boot.CreateSNMPBootPolls())
	sleep()
	par := createParsingService(0, vNetPort1, &types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	sleep()
	cli := createClient(vNetPort1)
	sleep()

	defer func() {
		cli.Shutdown()
		par.Shutdown()
		col.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	ip := "192.168.86.179"

	device := CreateDevice(ip, 0)

	cli.Multicast(types2.CastMode_All, types2.Action_POST, 0, config.TOPIC, device)

	time.Sleep(2 * time.Second)

	ic := inventory.Inventory(par.Resources())
	box := ic.ElementByKey(ip).(*types.NetworkBox)
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
