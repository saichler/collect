package poll_config

import (
	"errors"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8services/go/services/dcache"
	"github.com/saichler/l8utils/go/utils/strings"
	"github.com/saichler/l8types/go/ifs"
	"sync"
)

type PollConfigCenter struct {
	name2Poll ifs.IDistributedCache
	key2Name  map[string]string
	groups    map[string]map[string]string
	log       ifs.ILogger
	mtx       *sync.RWMutex
}

func newPollConfigCenter(resources ifs.IResources, listener ifs.IServiceCacheListener) *PollConfigCenter {
	pc := &PollConfigCenter{}
	pc.name2Poll = dcache.NewDistributedCache(ServiceName, ServiceArea, "PollConfig",
		resources.SysConfig().LocalUuid, listener, resources)
	pc.key2Name = make(map[string]string)
	pc.groups = make(map[string]map[string]string)
	pc.log = resources.Logger()
	pc.mtx = &sync.RWMutex{}
	return pc
}

func (this *PollConfigCenter) getPollName(key string) (string, bool) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	pollName, ok := this.key2Name[key]
	return pollName, ok
}

func (this *PollConfigCenter) getGroup(name string) map[string]string {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	return this.groups[name]
}

func (this *PollConfigCenter) deleteFromGroup(gEntry map[string]string, key string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(gEntry, key)
}

func (this *PollConfigCenter) deleteFromKey2Name(key string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(this.key2Name, key)
}

func (this *PollConfigCenter) deleteExisting(poll *types.PollConfig, key string) {
	existPoll := this.name2Poll.Get(poll.Name).(*types.PollConfig)
	if existPoll.Groups != nil {
		for _, gName := range existPoll.Groups {
			gEntry := this.getGroup(gName)
			if gEntry != nil {
				this.deleteFromGroup(gEntry, key)
			}
		}
	}
	this.deleteFromKey2Name(key)
}

func (this *PollConfigCenter) AddAll(polls []*types.PollConfig) {
	for _, poll := range polls {
		this.Add(poll, false)
	}
}

func (this *PollConfigCenter) Add(poll *types.PollConfig, isNotification bool) error {
	if poll.What == "" {
		return errors.New("poll does not contain a What value")
	}
	if poll.Name == "" {
		return errors.New("poll does not contain a Name")
	}

	key := this.PollKey(poll)
	_, ok := this.getPollName(key)

	if ok {
		this.deleteExisting(poll, key)
	}

	this.name2Poll.Put(poll.Name, poll, isNotification)

	this.mtx.Lock()
	defer this.mtx.Unlock()

	this.key2Name[key] = poll.Name
	if poll.Groups != nil {
		for _, gName := range poll.Groups {
			gEntry, ok := this.groups[gName]
			if !ok {
				this.groups[gName] = make(map[string]string)
				gEntry = this.groups[gName]
			}
			gEntry[key] = poll.Name
		}
	}
	return nil
}

func (this *PollConfigCenter) Update(poll *types.PollConfig, isNotification bool) error {
	if poll.What == "" {
		return errors.New("poll does not contain a What value")
	}
	if poll.Name == "" {
		return errors.New("poll does not contain a Name")
	}

	key := this.PollKey(poll)
	_, ok := this.getPollName(key)

	if ok {
		this.deleteExisting(poll, key)
	}

	this.name2Poll.Put(poll.Name, poll, isNotification)

	this.mtx.Lock()
	defer this.mtx.Unlock()

	this.key2Name[key] = poll.Name
	if poll.Groups != nil {
		for _, gName := range poll.Groups {
			gEntry, ok := this.groups[gName]
			if !ok {
				this.groups[gName] = make(map[string]string)
				gEntry = this.groups[gName]
			}
			gEntry[key] = poll.Name
		}
	}
	return nil
}

func (this *PollConfigCenter) PollKey(poll *types.PollConfig) string {
	return pollKey(poll.Name, poll.Vendor, poll.Series, poll.Family, poll.Software, poll.Hardware, poll.Version)
}

func (this *PollConfigCenter) PollByName(name string) *types.PollConfig {
	poll, _ := this.name2Poll.Get(name).(*types.PollConfig)
	return poll
}

func (this *PollConfigCenter) PollByKey(args ...string) *types.PollConfig {
	if args == nil || len(args) == 0 {
		return nil
	}
	if len(args) == 1 {
		pollName := this.key2Name[args[0]]
		poll, _ := this.name2Poll.Get(pollName).(*types.PollConfig)
		return poll
	}
	buff := strings.New()
	buff.Add(args[0])
	for i := 1; i < len(args); i++ {
		addToKey(args[i], buff)
	}
	p, ok := this.getPollName(buff.String())
	if ok {
		poll, _ := this.name2Poll.Get(p).(*types.PollConfig)
		return poll
	}
	return this.PollByKey(args[0 : len(args)-1]...)
}

func (this *PollConfigCenter) Names(groupName, vendor, series, family, software, hardware, version string) []string {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	result := make([]string, 0)
	group, ok := this.groups[groupName]
	if !ok {
		return result
	}
	for _, name := range group {
		result = append(result, name)
	}
	return result
}

func (this *PollConfigCenter) PollsByGroup(groupName, vendor, series, family, software, hardware, version string) []*types.PollConfig {
	names := this.Names(groupName, vendor, series, family, software, hardware, version)
	result := make([]*types.PollConfig, 0)
	for _, name := range names {
		poll := this.PollByKey(name, vendor, series, family, software, hardware, version)
		if poll != nil {
			result = append(result, poll)
		}
	}
	return result
}

func PollConfig(resource ifs.IResources) *PollConfigCenter {
	sp, ok := resource.Services().ServicePointHandler(ServiceName, ServiceArea)
	if !ok {
		return nil
	}
	return (sp.(*PollConfigServicePoint)).pollCenter
}
