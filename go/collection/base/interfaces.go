package base

import (
	"github.com/saichler/collect/go/types"
	"github.com/saichler/l8types/go/ifs"
)

const (
	Parser_Suffix = "-Parser"
)

type IJobCompleteHandler interface {
	JobCompleted(*types.Job)
}

type ProtocolCollector interface {
	Init(*types.ConnectionConfig, ifs.IResources) error
	Protocol() types.Protocol
	Exec(*types.Job)
	Connect() error
	Disconnect() error
}

type IController interface {
	StartPolling(config *types.DeviceConfig) error
	Shutdown()
}
