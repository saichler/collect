package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/config"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/collection/polling/boot"
	"sync"
	"testing"
)

func TestCollectionController(t *testing.T) {
	resourcs := createResources("collect")

	config.RegisterConfigCenter(0, resourcs, nil, nil)
	polling.RegisterPollCenter(0, resourcs, nil)
	pc := polling.Polling(resourcs)
	pc.AddAll(boot.CreateSNMPBootPolls())

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = resourcs
	cont := control.NewController(l, resourcs)
	/*
		cmds, commands := CreateCommands()
		for _, cmd := range cmds {
			cont.AddUpdateCommand(cmd)
		}*/
	device := CreateDevice("192.168.86.179")
	l.expected = 1
	cc := config.Configs(resourcs)

	cc.Add(device)
	cont.StartPolling(device.Id, 0)
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}
