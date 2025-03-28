package k8s

import (
	"github.com/google/uuid"
	"github.com/saichler/collect/go/collection/polling"
	"github.com/saichler/collect/go/types"
	"github.com/saichler/serializer/go/serialize/object"
	"github.com/saichler/shared/go/share/strings"
	"github.com/saichler/types/go/common"
	"os"
	"os/exec"
)

type Kubernetes struct {
	resources common.IResources
	config    *types.Config
}

func (this *Kubernetes) Init(config *types.Config, resources common.IResources) error {
	this.resources = resources
	this.config = config
	return nil
}

func (this *Kubernetes) Protocol() types.Protocol {
	return types.Protocol_K8s
}

func (this *Kubernetes) Exec(job *types.Job) {
	pollCenter := polling.Polling(this.resources, job.CServiceArea)
	pll := pollCenter.PollByName(job.PollName)
	if pll == nil {
		this.resources.Logger().Error("cannot find poll for name ", job.PollName)
		return
	}

	script := strings.New("kubectl --kubeconfig=")
	script.Add(this.config.KubeConfig)
	script.Add(" --context=")
	script.Add(this.config.KukeContext)
	script.Add(" ")
	script.Add(pll.What)
	script.Add("\n")

	id := uuid.New().String()
	in := "./" + id + ".sh"
	defer os.Remove(in)
	os.WriteFile(in, script.Bytes(), 0777)
	c := exec.Command("bash", "-c", in, "2>&1")
	o, e := c.Output()
	if e != nil {
		job.Error = e.Error()
	}
	obj := object.NewEncode([]byte{}, 0)
	obj.Add(string(o))
	job.Result = obj.Data()
}

func (this *Kubernetes) Connect() error {
	return nil
}

func (this *Kubernetes) Disconnect() error {
	return nil
}
