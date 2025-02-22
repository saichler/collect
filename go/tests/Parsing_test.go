package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/collect/go/types"
	types2 "github.com/saichler/shared/go/types"
	"testing"
	"time"
)

func TestParsingAndInventory(t *testing.T) {

	sw := createSwitch()
	sleep()
	col := createCollectionService(boot.CreateSNMPBootPolls())
	sleep()
	par := createParsingService(&types.NetworkBox{}, "Id", boot.CreateSNMPBootPolls())
	sleep()
	cli := createClient()
	sleep()

	defer func() {
		cli.Shutdown()
		par.Shutdown()
		col.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	ip := "192.168.86.179"

	device := CreateDevice(ip)

	cli.Do(types2.Action_POST, config.TOPIC, device)

	time.Sleep(2 * time.Second)

	ic := inventory.Inventory(par.Resources())
	box := ic.ElementByKey(ip).(*types.NetworkBox)
	if box == nil {
		log.Fail(t, "Expected box to be non-nil")
		return
	}

	if box.Info == nil {
		log.Fail(t, "Expected box info to be non-nil")
		return
	}

	if box.Info.SysName == "" {
		log.Fail(t, "Expected box info sysname to not be blank")
		return
	}

	if box.Info.Vendor == "" {
		log.Fail(t, "Expected box vendor to not be blank")
		return
	}
}
