package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/inventory"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/collection/polling/boot"
	types3 "github.com/saichler/k8s_observer/go/types"
	types2 "github.com/saichler/shared/go/types"
	"sync"
	"testing"
	"time"
)

func TestK8sCollector(t *testing.T) {
	resourcs := createResources("k8s")

	cluster := CreateCluster("/Users/orlyaicler/admin.conf", "kubernetes-admin@kubernetes")

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = resourcs
	l.ph = control.NewDirectParsingHandler(nil, resourcs)
	cont := control.NewController(l, resourcs)

	config.RegisterConfigCenter(resourcs, nil, cont)
	polling.RegisterPollCenter(resourcs, nil)

	l.expected = 1
	cc := config.Configs(resourcs)
	pp := polling.Polling(resourcs)

	pp.AddAll(boot.CreateK8sBootPolls())

	cc.Add(cluster)
	cont.StartPolling(cluster.Id)

	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}

func TestParsingForK8s(t *testing.T) {

	sw := createSwitch()
	sleep()
	col := createCollectionService(boot.CreateK8sBootPolls())
	sleep()
	par := createParsingService(&types3.Cluster{}, "Name", boot.CreateK8sBootPolls())
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

	admin := "/Users/orlyaicler/admin.conf"
	context := "kubernetes-admin@kubernetes"

	cluster := CreateCluster(admin, context)

	cli.Do(types2.Action_POST, config.TOPIC, cluster)

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
}
