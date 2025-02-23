package common

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/shared/go/share/interfaces"
)

type CollectNotificationHandler interface {
	HandleCollectNotification(*types.Job, int32)
}

type ProtocolCollector interface {
	Init(*types.Config, interfaces.IResources) error
	Protocol() types.Protocol
	Exec(*types.Job)
	Connect() error
	Disconnect() error
}

type IController interface {
	StartPolling(string, int32) error
}
