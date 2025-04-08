package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/collect/go/types"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/types/go/common"
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

	cli.Multicast(config.ServiceName, 0, common.POST, device)

	time.Sleep(1 * time.Second)

	ic := inventory.Inventory(par.Resources(), InvServiceName, 0)
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
