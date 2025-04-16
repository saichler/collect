package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/collector"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
	tear()
}

func TestCollectionController(t *testing.T) {

	cfg := topo.VnicByVnetNum(2, 4)

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = cfg.Resources()
	cont := collector.NewDeviceCollector(l, cfg.Resources())
	activateDeviceAndPollConfigServices(cfg, 0, boot.CreateSNMPBootPolls(), cont)
	defer func() {
		deActivateDeviceAndPollConfigServices(cfg, 0)
	}()

	serviceArea := uint16(0)
	device := CreateDevice("192.168.86.179", serviceArea)
	l.expected = 1
	cc := device_config.Configs(cfg.Resources(), 0)

	cc.Add(device)
	cont.StartPolling(device)
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}
