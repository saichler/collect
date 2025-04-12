package tests

import (
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	. "github.com/saichler/l8test/go/infra/t_resources"
	"github.com/saichler/probler/go/serializers"
	types3 "github.com/saichler/probler/go/types"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

func TestK8s1Collector2Parsers(t *testing.T) {

	sw := createVNet(vNetPort1)
	sleep()
	col := createCollectionService(0, vNetPort1, boot.CreateK8sBootPolls())
	sleep()
	par1 := createParsingService(0, vNetPort1, &types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	par2 := createParsingService(1, vNetPort1, &types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	sleep()
	cli := createClient(vNetPort1)
	sleep()

	par1.Resources().Registry().RegisterEnums(types3.NodeStatus_value)
	par1.Resources().Registry().RegisterEnums(types3.PodStatus_value)
	par2.Resources().Registry().RegisterEnums(types3.NodeStatus_value)
	par2.Resources().Registry().RegisterEnums(types3.PodStatus_value)

	info, err := par1.Resources().Registry().Info("ReadyState")
	if err != nil {
		Log.Fail(t, "Error getting registry info")
		return
	}
	info.AddSerializer(&serializers.Ready{})

	info, err = par2.Resources().Registry().Info("ReadyState")
	if err != nil {
		Log.Fail(t, "Error getting registry info")
		return
	}
	info.AddSerializer(&serializers.Ready{})

	defer func() {
		cli.Shutdown()
		par1.Shutdown()
		par2.Shutdown()
		col.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	cli.Multicast(deviceconfig.ServiceName, 0, common.POST, cluster1)
	cli.Multicast(deviceconfig.ServiceName, 0, common.POST, cluster2)

	time.Sleep(2 * time.Second)

	if !checkCluster(par1.Resources(), context1, t, 0) {
		return
	}

	if !checkCluster(par2.Resources(), context2, t, 1) {
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
