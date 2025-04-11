package polling

import (
	"errors"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/strings"
	"github.com/saichler/types/go/common"
	"sync"
)

type PollCenter struct {
	name2Poll *cache.Cache
	key2Name  map[string]string
	groups    map[string]map[string]string
	log       common.ILogger
	mtx       *sync.RWMutex
}

func newPollCenter(serviceArea uint16, resources common.IResources, listener common.IServicePointCacheListener) *PollCenter {
	pc := &PollCenter{}
	pc.name2Poll = cache.NewModelCache(ServiceName, serviceArea, "Poll",
		resources.SysConfig().LocalUuid, listener, resources.Introspector())
	pc.key2Name = make(map[string]string)
	pc.groups = make(map[string]map[string]string)
	pc.log = resources.Logger()
	pc.mtx = &sync.RWMutex{}
	return pc
}

func (this *PollCenter) getPollName(key string) (string, bool) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	pollName, ok := this.key2Name[key]
	return pollName, ok
}

func (this *PollCenter) getGroup(name string) map[string]string {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	return this.groups[name]
}

func (this *PollCenter) deleteFromGroup(gEntry map[string]string, key string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(gEntry, key)
}

func (this *PollCenter) deleteFromKey2Name(key string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(this.key2Name, key)
}

func (this *PollCenter) deleteExisting(poll *types.Poll, key string) {
	existPoll := this.name2Poll.Get(poll.Name).(*types.Poll)
	if existPoll.Groups != nil {
		for _, gName := range existPoll.Groups {
			gEntry := this.getGroup(gName)
			if gEntry != nil {
				this.deleteFromGroup(gEntry, key)
			}
		}
	}
	this.deleteFromKey2Name(key)
	this.name2Poll.Delete(poll.Name)
}

func (this *PollCenter) AddAll(polls []*types.Poll) {
	for _, poll := range polls {
		this.Add(poll)
	}
}

func (this *PollCenter) Add(poll *types.Poll) error {
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

	this.name2Poll.Put(poll.Name, poll)

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

func (this *PollCenter) PollKey(poll *types.Poll) string {
	return pollKey(poll.Name, poll.Vendor, poll.Series, poll.Family, poll.Software, poll.Hardware, poll.Version)
}

func (this *PollCenter) PollByName(name string) *types.Poll {
	poll, _ := this.name2Poll.Get(name).(*types.Poll)
	return poll
}

func (this *PollCenter) PollByKey(args ...string) *types.Poll {
	if args == nil || len(args) == 0 {
		return nil
	}
	if len(args) == 1 {
		pollName := this.key2Name[args[0]]
		poll, _ := this.name2Poll.Get(pollName).(*types.Poll)
		return poll
	}
	buff := strings.New()
	buff.Add(args[0])
	for i := 1; i < len(args); i++ {
		addToKey(args[i], buff)
	}
	p, ok := this.getPollName(buff.String())
	if ok {
		poll, _ := this.name2Poll.Get(p).(*types.Poll)
		return poll
	}
	return this.PollByKey(args[0 : len(args)-1]...)
}

func (this *PollCenter) Names(groupName, vendor, series, family, software, hardware, version string) []string {
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

func (this *PollCenter) PollsByGroup(groupName, vendor, series, family, software, hardware, version string) []*types.Poll {
	names := this.Names(groupName, vendor, series, family, software, hardware, version)
	result := make([]*types.Poll, 0)
	for _, name := range names {
		poll := this.PollByKey(name, vendor, series, family, software, hardware, version)
		if poll != nil {
			result = append(result, poll)
		}
	}
	return result
}

func Polling(resource common.IResources, serviceArea uint16) *PollCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(ServiceName, serviceArea)
	if !ok {
		return nil
	}
	return (sp.(*PollServicePoint)).pollCenter
}
