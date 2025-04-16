package tests

import (
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	. "github.com/saichler/l8test/go/infra/t_resources"
	types3 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestK8s1Collector2Parsers(t *testing.T) {
	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	polls := boot.CreateK8sBootPolls()

	cfg := topo.VnicByVnetNum(2, 4)
	par1 := topo.VnicByVnetNum(3, 1)
	par2 := topo.VnicByVnetNum(1, 1)
	inv1 := topo.VnicByVnetNum(1, 3)
	inv2 := topo.VnicByVnetNum(2, 3)

	activateDeviceAndPollConfigServices(cfg, 0, polls)

	activateParsingAndPollConfigServices(par1, cluster1.ParsingService,
		&types3.Cluster{}, "Name", polls)
	activateParsingAndPollConfigServices(par2, cluster2.ParsingService,
		&types3.Cluster{}, "Name", polls)

	activateInventoryService(inv1, cluster1.InventoryService, &types3.Cluster{}, "Name")
	activateInventoryService(inv2, cluster2.InventoryService, &types3.Cluster{}, "Name")

	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
		deActivateParsingAndPollConfigServices(par1, cluster1.ParsingService)
		deActivateParsingAndPollConfigServices(par2, cluster2.ParsingService)
		deActivateInventoryService(inv1, cluster1.InventoryService)
		deActivateInventoryService(inv2, cluster2.InventoryService)
	}()
	sleep()

	cli := topo.VnicByVnetNum(1, 2)
	cli.Multicast(device_config.ServiceName, 0, common.POST, cluster1)
	cli.Multicast(device_config.ServiceName, 0, common.POST, cluster2)

	time.Sleep(2 * time.Second)

	if !checkCluster(inv1.Resources(), context1, t, 0) {
		return
	}

	if !checkCluster(inv2.Resources(), context2, t, 1) {
		return
	}
}

func checkCluster(resourcs common.IResources, context string, t *testing.T, serviceArea uint16) bool {
	ic := inventory.Inventory(resourcs, "Cluster", serviceArea)
	k8sCluster := ic.ElementByKey(context).(*types3.Cluster)
	if k8sCluster == nil {
		Log.Fail(t, context, " Expected K8s Cluster to be non-nil")
		return false
	}

	if k8sCluster.Nodes == nil {
		Log.Fail(t, context, " Expected K8s Cluster nodes to be non-nil")
		return false
	}
	if len(k8sCluster.Nodes) != 6 {
		Log.Fail(t, context, " Expected K8s Cluster nodes to be 6")
		return false
	}

	if k8sCluster.Pods == nil {
		Log.Fail(t, context, " Expected K8s Cluster pods to be non-nil")
		return false
	}

	if len(k8sCluster.Pods) < 17 {
		Log.Fail(t, context, " Expected K8s Cluster pods to be at least 17")
		return false
	}
	for _, pod := range k8sCluster.Pods {
		if pod.Status != types3.PodStatus_Running {
			Log.Fail(t, context, " Expected K8s Pod to be Running ", pod.Status.String())
			return false
		}
		if pod.Ready == nil || pod.Ready.Count == 0 {
			Log.Fail(t, context, " Expected K8s Pod state to be Ready ", pod.Ready.Count, "/", pod.Ready.Outof)
			return false
		}
	}
	return true
}
