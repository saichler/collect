package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/collection/polling/boot"
	"github.com/saichler/k8s_observer/go/serializers"
	types3 "github.com/saichler/k8s_observer/go/types"
	types2 "github.com/saichler/shared/go/types"
	"sync"
	"testing"
	"time"
)

func TestK8sCollector(t *testing.T) {
	resourcs := createResources("k8s")

	cluster := CreateCluster("/Users/orlyaicler/admin.conf", "kubernetes-admin@kubernetes", 0)

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = resourcs
	//l.ph = control.NewDirectParsingHandler(nil, resourcs)
	cont := control.NewController(l, resourcs)

	config.RegisterConfigCenter(0, resourcs, nil, cont)
	polling.RegisterPollCenter(0, resourcs, nil)

	l.expected = 1
	cc := config.Configs(resourcs)
	pp := polling.Polling(resourcs)

	pp.AddAll(boot.CreateK8sBootPolls())

	cc.Add(cluster)
	cont.StartPolling(cluster.Id, 0)

	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}

func TestParsingForK8s(t *testing.T) {

	sw := createVNet(vNetPort1)
	sleep()
	col := createCollectionService(0, vNetPort1, boot.CreateK8sBootPolls())
	sleep()
	par := createParsingService(0, vNetPort1, &types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
	sleep()
	cli := createClient(vNetPort1)
	sleep()

	par.Resources().Registry().RegisterEnums(types3.NodeStatus_value)
	par.Resources().Registry().RegisterEnums(types3.PodStatus_value)
	info, err := par.Resources().Registry().Info("ReadyState")
	if err != nil {
		log.Fail(t, "Error getting registry info")
		return
	}
	info.AddSerializer(&serializers.Ready{})
	defer func() {
		cli.Shutdown()
		par.Shutdown()
		col.Shutdown()
		sw.Shutdown()
	}()

	sleep()

	admin := "/Users/orlyaicler/admin.conf"
	context := "kubernetes-admin@kubernetes"

	cluster := CreateCluster(admin, context, 0)

	cli.Multicast(types2.Action_POST, 0, config.TOPIC, cluster)

	time.Sleep(2 * time.Second)

	ic := inventory.Inventory(par.Resources())
	k8sCluster := ic.ElementByKey(context).(*types3.Cluster)
	if k8sCluster == nil {
		log.Fail(t, "Expected K8s Cluster to be non-nil")
		return
	}

	if k8sCluster.Nodes == nil {
		log.Fail(t, "Expected K8s Cluster nodes to be non-nil")
		return
	}
	if len(k8sCluster.Nodes) != 6 {
		log.Fail(t, "Expected K8s Cluster nodes to be 6")
		return
	}

	if k8sCluster.Pods == nil {
		log.Fail(t, "Expected K8s Cluster pods to be non-nil")
		return
	}

	if len(k8sCluster.Pods) != 17 {
		log.Fail(t, "Expected K8s Cluster pods to be 17")
		return
	}
	for _, pod := range k8sCluster.Pods {
		if pod.Status != types3.PodStatus_Running {
			log.Fail(t, "Expected K8s Pod to be Running ", pod.Status.String())
			return
		}
		if pod.Ready == nil || pod.Ready.Count == 0 {
			log.Fail(t, "Expected K8s Pod state to be Ready ", pod.Ready.Count, "/", pod.Ready.Outof)
			return
		}
	}
}
