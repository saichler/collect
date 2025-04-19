package tests

import (
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	types2 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestK8s2Collector2Parsers(t *testing.T) {
	time.Sleep(2 * time.Second)
	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	polls := boot.CreateK8sBootPolls()

	cfg1 := topo.VnicByVnetNum(1, 1)
	cfg2 := topo.VnicByVnetNum(1, 2)
	par1 := topo.VnicByVnetNum(1, 3)
	par2 := topo.VnicByVnetNum(1, 4)
	inv1 := topo.VnicByVnetNum(1, 1)
	inv2 := topo.VnicByVnetNum(1, 2)

	activateDeviceAndPollConfigServices(cfg1, 0, polls)

	activateDeviceAndPollConfigServices(cfg2, 1, polls)

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

	time.Sleep(2 * time.Second)

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
