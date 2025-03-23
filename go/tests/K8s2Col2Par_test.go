package tests

import (
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/k8s_observer/go/serializers"
	types3 "github.com/saichler/k8s_observer/go/types"
	. "github.com/saichler/shared/go/tests/infra"
	types2 "github.com/saichler/types/go/types"
	"testing"
	"time"
)

func TestK8s2Collector2Parsers(t *testing.T) {
	sw := createVNet(vNetPort1)
	sleep()
	col1 := createCollectionService(0, vNetPort1, boot.CreateK8sBootPolls())
	col2 := createCollectionService(1, vNetPort1, boot.CreateK8sBootPolls())
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
		col1.Shutdown()
		col2.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	cli.Multicast(config.ServiceName, 1, types2.Action_POST, cluster1)
	cli.Multicast(config.ServiceName, 0, types2.Action_POST, cluster2)

	time.Sleep(2 * time.Second)

	if !checkCluster(par1.Resources(), context1, t, 0) {
		return
	}

	if !checkCluster(par2.Resources(), context2, t, 1) {
		return
	}
}
