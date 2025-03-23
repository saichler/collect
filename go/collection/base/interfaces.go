package base

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/types/go/common"
)

const (
	Parsing_Suffix = "_Parsing"
)

type CollectNotificationHandler interface {
	HandleCollectNotification(*types.Job)
}

type ProtocolCollector interface {
	Init(*types.Config, common.IResources) error
	Protocol() types.Protocol
	Exec(*types.Job)
	Connect() error
	Disconnect() error
}

type IController interface {
	StartPolling(string, string) error
}
