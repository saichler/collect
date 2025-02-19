package polling

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/servicepoints/go/points/cache"
	"github.com/saichler/shared/go/share/interfaces"
	types2 "github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC    = "Poll"
	ENDPOINT = "poll"
)

type PollServicePoint struct {
	pollCenter *PollCenter
}

func RegisterPollCenter(resources interfaces.IResources, listener cache.ICacheListener) {
	psp := &PollServicePoint{}
	psp.pollCenter = newPollCenter(resources, listener)
	err := resources.ServicePoints().RegisterServicePoint(&types.Poll{}, psp)
	if err != nil {
		panic(err)
	}
}

func (this *PollServicePoint) Post(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	hp := pb.(*types.Poll)
	this.pollCenter.Add(hp)
	return nil, nil
}
func (this *PollServicePoint) Put(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Patch(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Delete(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Get(pb proto.Message, vnic interfaces.IVirtualNetworkInterface) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) Failed(pb proto.Message, vnic interfaces.IVirtualNetworkInterface, msg *types2.Message) (proto.Message, error) {
	return nil, nil
}
func (this *PollServicePoint) EndPoint() string {
	return ENDPOINT
}
func (this *PollServicePoint) Topic() string {
	return TOPIC
}
