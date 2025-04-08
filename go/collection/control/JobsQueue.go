package control

import (
	"errors"
	"github.com/saichler/collect/go/collection/base"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
	"sync"
	"time"
)

type JobsQueue struct {
	deviceId     string
	hostId       string
	jobs         []*types.Job
	jobsMap      map[string]*types.Job
	mtx          *sync.Mutex
	resources    common.IResources
	cServiceArea uint16
	dServiceArea uint16
	serviceName  string
}

func NewJobsQueue(deviceId, hostId string, resources common.IResources, serviceName string, cServiceArea, dServiceArea uint16) *JobsQueue {
	jq := &JobsQueue{}
	jq.resources = resources
	jq.mtx = &sync.Mutex{}
	jq.jobs = make([]*types.Job, 0)
	jq.jobsMap = make(map[string]*types.Job)
	jq.deviceId = deviceId
	jq.hostId = hostId
	jq.cServiceArea = cServiceArea
	jq.dServiceArea = dServiceArea
	jq.serviceName = serviceName + base.Parsing_Suffix
	return jq
}

func (this *JobsQueue) newJob(name, vendor, series, family, software, hardware, version string, cadence, timeout int64) *types.Job {
	pc := polling.Polling(this.resources, this.cServiceArea)
	poll := pc.PollByKey(name, vendor, series, family, software, hardware, version)
	if poll == nil {
		return nil
	}
	job := &types.Job{}
	job.PollName = poll.Name
	job.Cadence = cadence
	job.Timeout = timeout
	job.DeviceId = this.deviceId
	job.HostId = this.hostId
	job.CServiceArea = int32(this.cServiceArea)
	job.DServiceArea = int32(this.dServiceArea)
	job.ServiceName = this.serviceName

	if job.Cadence == 0 {
		job.Cadence = poll.DefaultCadence
	}
	if job.Timeout == 0 {
		job.Timeout = poll.DefaultTimeout
	}
	return job
}

func (this *JobsQueue) newJobs(groupName, vendor, series, family, software, hardware, version string) []*types.Job {
	pc := polling.Polling(this.resources, this.cServiceArea)
	polls := pc.PollsByGroup(groupName, vendor, series, family, software, hardware, version)
	jobs := make([]*types.Job, 0)
	for _, poll := range polls {
		job := &types.Job{}
		job.DeviceId = this.deviceId
		job.HostId = this.hostId
		job.PollName = poll.Name
		job.Cadence = poll.DefaultCadence
		job.Timeout = poll.DefaultTimeout
		job.DServiceArea = int32(this.dServiceArea)
		job.CServiceArea = int32(this.cServiceArea)
		job.ServiceName = this.serviceName
		jobs = append(jobs, job)
	}
	return jobs
}

func (this *JobsQueue) InsertJob(name, vendor, series, family, software, hardware, version string, cadence, timeout int64) error {
	job := this.newJob(name, vendor, series, family, software, hardware, version, cadence, timeout)
	if job == nil {
		return errors.New("cannot find poll to create job")
	}
	this.mtx.Lock()
	defer this.mtx.Unlock()

	old, ok := this.jobsMap[job.PollName]
	if !ok {
		this.jobsMap[job.PollName] = job
		this.jobs = append(this.jobs, job)
	} else {
		old.Started = 0
		old.Ended = 0
	}
	return nil
}

func (this *JobsQueue) Pop() (*types.Job, int64) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	var job *types.Job
	index := -1
	now := time.Now().Unix()
	waitTimeTillNext := int64(5)
	for i, j := range this.jobs {
		timeSinceExecuted := now - j.Ended
		if timeSinceExecuted >= j.Cadence {
			job = j
			index = i
			break
		} else {
			timeTillNextExecution := j.Cadence - timeSinceExecuted
			if timeTillNextExecution < waitTimeTillNext {
				waitTimeTillNext = timeTillNextExecution
			}
		}
	}
	this.moveToLast(index)
	return job, waitTimeTillNext
}

func (this *JobsQueue) moveToLast(index int) {
	if index != -1 {
		swap := make([]*types.Job, 0)
		job := this.jobs[index]
		swap = append(swap, this.jobs[0:index]...)
		swap = append(swap, this.jobs[index+1:]...)
		swap = append(swap, job)
		for i, j := range swap {
			this.jobs[i] = j
		}
	}
}

func MarkStart(job *types.Job) {
	job.Started = time.Now().Unix()
	job.Ended = 0
	job.Result = nil
	job.Error = ""
}

func MarkEnded(job *types.Job) {
	job.Ended = time.Now().Unix()
}
