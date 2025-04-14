package tests

import (
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	types2 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestK8s2Collector2Parsers2Vnet(t *testing.T) {

	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	polls := boot.CreateK8sBootPolls()

	cfg1 := topo.VnicByVnetNum(2, 4)
	cfg2 := topo.VnicByVnetNum(3, 4)
	par1 := topo.VnicByVnetNum(3, 1)
	par2 := topo.VnicByVnetNum(1, 1)
	inv1 := topo.VnicByVnetNum(1, 3)
	inv2 := topo.VnicByVnetNum(2, 3)

	cont1 := control.NewController(control.NewParsingCenterNotifier(cfg1), cfg1.Resources())
	activateDeviceAndPollConfigServices(cfg1, 0, cont1, polls)

	cont2 := control.NewController(control.NewParsingCenterNotifier(cfg2), cfg2.Resources())
	activateDeviceAndPollConfigServices(cfg2, 1, cont2, polls)

	activateParsingAndPollConfigServices(par1, cluster1.ParsingService,
		&types2.Cluster{}, "Name", polls)
	activateParsingAndPollConfigServices(par2, cluster2.ParsingService,
		&types2.Cluster{}, "Name", polls)

	activateInventoryService(inv1, cluster1.InventoryService, &types2.Cluster{}, "Name")
	activateInventoryService(inv2, cluster2.InventoryService, &types2.Cluster{}, "Name")

	defer func() {
		deActivateDeviceAndPollConfigServices(cfg1, 0)
		deActivateDeviceAndPollConfigServices(cfg2, 1)
		deActivateParsingAndPollConfigServices(par1, cluster1.ParsingService)
		deActivateParsingAndPollConfigServices(par2, cluster2.ParsingService)
		deActivateInventoryService(inv1, cluster1.InventoryService)
		deActivateInventoryService(inv2, cluster2.InventoryService)
	}()
	sleep()

	cli := topo.VnicByVnetNum(1, 2)
	cli.Multicast(device_config.ServiceName, 0, common.POST, cluster1)
	cli.Multicast(device_config.ServiceName, 1, common.POST, cluster2)

	time.Sleep(2 * time.Second)

	if !checkCluster(inv1.Resources(), context1, t, 0) {
		return
	}

	if !checkCluster(inv2.Resources(), context2, t, 1) {
		return
	}
}
