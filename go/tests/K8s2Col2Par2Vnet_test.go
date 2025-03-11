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

func TestK8s2Collector2Parsers2Vnet(t *testing.T) {
	sw1 := createVNet(vNetPort1)
	sw2 := createVNet(vNetPort2)
	sleep()
	col1 := createCollectionService(0, vNetPort1, boot.CreateK8sBootPolls())
	col2 := createCollectionService(1, vNetPort2, boot.CreateK8sBootPolls())
	sleep()
	par1 := createParsingService(0, vNetPort2, &types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	par2 := createParsingService(1, vNetPort2, &types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	sleep()
	cli := createClient(vNetPort2)
	sleep()
	sw1.ConnectNetworks("127.0.0.1", sw2.Resources().Config().VnetPort)
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
		sw1.Shutdown()
		sw2.Shutdown()
	}()

	sleep()

	admin1 := home + "/admin.conf"
	context1 := "kubernetes-admin@kubernetes"
	admin2 := home + "/lab.conf"
	context2 := "lab"

	cluster1 := CreateCluster(admin1, context1, 0)
	cluster2 := CreateCluster(admin2, context2, 1)

	cli.Multicast(types2.CastMode_All, types2.Action_POST, 1, config.TOPIC, cluster1)
	cli.Multicast(types2.CastMode_All, types2.Action_POST, 0, config.TOPIC, cluster2)

	time.Sleep(2 * time.Second)

	if !checkCluster(par1.Resources(), context1, t) {
		return
	}

	if !checkCluster(par2.Resources(), context2, t) {
		return
	}
}
