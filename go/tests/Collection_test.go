package tests

import (
	"fmt"
	"github.com/saichler/collect/go/collection/control"
	"github.com/saichler/collect/go/collection/device_config"
	"github.com/saichler/collect/go/collection/poll_config"
	"github.com/saichler/collect/go/collection/poll_config/boot"
	"sync"
	"testing"
)

func TestCollectionController(t *testing.T) {
	serviceArea := uint16(0)
	resourcs := createResources("collect")

	device_config.RegisterConfigCenter(serviceArea, resourcs, nil, nil)
	poll_config.RegisterPollCenter(serviceArea, resourcs, nil)
	pc := poll_config.Polling(resourcs, serviceArea)
	pc.AddAll(boot.CreateSNMPBootPolls())

	l := &CollectorListener{}
	l.cond = sync.NewCond(&sync.Mutex{})
	l.resources = resourcs
	cont := control.NewController(l, resourcs, 0)
	/*
		cmds, commands := CreateCommands()
		for _, cmd := range cmds {
			cont.AddUpdateCommand(cmd)
		}*/
	device := CreateDevice("192.168.86.179", serviceArea)
	l.expected = 1
	cc := device_config.Configs(resourcs, serviceArea)

	cc.Add(device)
	cont.StartPolling(device.Id, "")
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.cond.Wait()
	fmt.Println("Test Ended")
}
