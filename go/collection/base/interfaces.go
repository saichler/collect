package base

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

type CollectNotificationHandler interface {
	HandleCollectNotification(*types.Job, int32)
}

type ProtocolCollector interface {
	Init(*types.Config, common.IResources) error
	Protocol() types.Protocol
	Exec(*types.Job)
	Connect() error
	Disconnect() error
}

type IController interface {
	StartPolling(string, int32) error
}
