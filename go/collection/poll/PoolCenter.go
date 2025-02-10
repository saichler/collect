package poll

import (
	"errors"
	"github.com/google/uuid"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/share/strings"
	"sync"
)

type PollCenter struct {
	uuid2Poll *cache.Cache
	key2uuid  map[string]string
	groups    map[string]map[string]string
	log       interfaces.ILogger
	mtx       *sync.RWMutex
}

func newPollCenter(resources interfaces.IResources, listener cache.ICacheListener) *PollCenter {
	pc := &PollCenter{}
	pc.uuid2Poll = cache.NewModelCache(resources.Config().LocalUuid, listener, resources.Introspector())
	pc.key2uuid = make(map[string]string)
	pc.groups = make(map[string]map[string]string)
	pc.log = resources.Logger()
	pc.mtx = &sync.RWMutex{}
	return pc
}

func (this *PollCenter) getPollUuid(key string) (string, bool) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	pollUuid, ok := this.key2uuid[key]
	return pollUuid, ok
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

func (this *PollCenter) deleteFromKey2Uuid(key string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(this.key2uuid, key)
}

func (this *PollCenter) deleteExisting(poll *types.Poll, pollUuid, key string) {
	if poll.Uuid != "" && poll.Uuid != pollUuid {
		this.log.Error("provided poll uuid is different than the uuid by key for existing poll, ignoring it")
		poll.Uuid = ""
	}
	existPoll := this.uuid2Poll.Get(pollUuid).(*types.Poll)
	if existPoll.Groups != nil {
		for _, gName := range existPoll.Groups {
			gEntry := this.getGroup(gName)
			if gEntry != nil {
				this.deleteFromGroup(gEntry, key)
			}
		}
	}
	this.deleteFromKey2Uuid(key)
	this.uuid2Poll.Delete(pollUuid)
}

func (this *PollCenter) Add(poll *types.Poll) error {
	if poll.What == "" {
		return errors.New("poll does not contain a What value")
	}
	if poll.Name == "" {
		return errors.New("poll does not contain a Name")
	}

	key := this.PollKey(poll)
	pollUuid, ok := this.getPollUuid(key)

	if ok {
		this.deleteExisting(poll, pollUuid, key)
	} else {
		pollUuid = poll.Uuid
	}

	if pollUuid == "" {
		pollUuid = uuid.New().String()
	}
	poll.Uuid = pollUuid
	this.uuid2Poll.Put(pollUuid, poll)

	this.mtx.Lock()
	defer this.mtx.Unlock()

	this.key2uuid[key] = pollUuid
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

func (this *PollCenter) PollByUuid(uuid string) *types.Poll {
	poll, _ := this.uuid2Poll.Get(uuid).(*types.Poll)
	return poll
}

func (this *PollCenter) PollByKey(args ...string) *types.Poll {
	if args == nil || len(args) == 0 {
		return nil
	}
	if len(args) == 1 {
		pollUuid := this.key2uuid[args[0]]
		poll, _ := this.uuid2Poll.Get(pollUuid).(*types.Poll)
		return poll
	}
	buff := strings.New()
	buff.Add(args[0])
	for i := 1; i < len(args); i++ {
		addToKey(args[i], buff)
	}
	p, ok := this.getPollUuid(buff.String())
	if ok {
		poll, _ := this.uuid2Poll.Get(p).(*types.Poll)
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

func Poll(resource interfaces.IResources) *PollCenter {
	sp, ok := resource.ServicePoints().ServicePointHandler(TOPIC)
	if !ok {
		return nil
	}
	return (sp.(*PollServicePoint)).pollCenter
}
