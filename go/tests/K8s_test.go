package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/probler/go/serializers"
	types3 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"sync"
	"testing"
	"time"
)

func TestK8sCollector(t *testing.T) {

	cfg := topo.VnicByVnetNum(2, 4)

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = cfg.Resources()
	cont := control.NewController(l, cfg.Resources())
	activateDeviceAndPollConfigServices(cfg, 0, cont, boot.CreateK8sBootPolls())
	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
	}()

	cluster := CreateCluster(admin1, context1, 0)

	cc := device_config.Configs(cfg.Resources(), 0)

	cc.Add(cluster)
	cont.StartPolling(cluster)

	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}

func TestParsingForK8s(t *testing.T) {

	cluster := CreateCluster(admin1, context1, 0)

	cfg := topo.VnicByVnetNum(2, 4)
	par := topo.VnicByVnetNum(3, 1)
	inv := topo.VnicByVnetNum(1, 3)

	cont := control.NewController(control.NewParsingCenterNotifier(cfg), cfg.Resources())
	activateDeviceAndPollConfigServices(cfg, 0, cont, boot.CreateK8sBootPolls())
	activateParsingAndPollConfigServices(par, cluster.ParsingService,
		&types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	activateInventoryService(inv, cluster.InventoryService, &types3.Cluster{}, "Name")

	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
		deActivateParsingAndPollConfigServices(par, cluster.ParsingService)
		deActivateInventoryService(inv, cluster.InventoryService)
	}()
	sleep()

	info, err := par.Resources().Registry().Info("ReadyState")
	if err != nil {
		Log.Fail(t, "Error getting registry info")
		return
	}
	info.AddSerializer(&serializers.Ready{})

	sleep()

	cli := topo.VnicByVnetNum(1, 2)
	cli.Multicast(device_config.ServiceName, 0, common.POST, cluster)

	time.Sleep(2 * time.Second)

	ic := inventory.Inventory(inv.Resources(), cluster.InventoryService.ServiceName, uint16(cluster.InventoryService.ServiceArea))
	var k8sCluster *types3.Cluster
	var ok bool
	for i := 0; i < 10; i++ {
		k8sCluster, ok = ic.ElementByKey(context1).(*types3.Cluster)
		if ok {
			break
		}
		time.Sleep(time.Second)
	}

	if k8sCluster == nil {
		Log.Fail(t, "Expected K8s Cluster to be non-nil")
		return
	}

	if k8sCluster.Nodes == nil {
		Log.Fail(t, "Expected K8s Cluster nodes to be non-nil")
		return
	}
	if len(k8sCluster.Nodes) != 6 {
		Log.Fail(t, "Expected K8s Cluster nodes to be 6")
		return
	}

	if k8sCluster.Pods == nil {
		Log.Fail(t, "Expected K8s Cluster pods to be non-nil")
		return
	}

	if len(k8sCluster.Pods) < 17 {
		Log.Fail(t, "Expected K8s Cluster pods to be 17")
		return
	}
	for _, pod := range k8sCluster.Pods {
		if pod.Status != types3.PodStatus_Running {
			Log.Fail(t, "Expected K8s Pod to be Running ", pod.Status.String())
			return
		}
		if pod.Ready == nil || pod.Ready.Count == 0 {
			Log.Fail(t, "Expected K8s Pod state to be Ready ", pod.Ready.Count, "/", pod.Ready.Outof)
			return
		}
	}
}
