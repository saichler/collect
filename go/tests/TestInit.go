package tests

import (
	. "github.com/saichler/l8test/go/infra/t_resources"
	. "github.com/saichler/l8test/go/infra/t_topology"
	"github.com/saichler/layer8/go/overlay/protocol"
	"github.com/saichler/shared/go/share/logger"
	. "github.com/saichler/types/go/common"
	"time"
)

var topo *TestTopology
var slog = logger.NewLoggerDirectImpl(logger.NewFileLogMethod("/tmp/log"))

func init() {
	Log.SetLogLevel(Trace_Level)
}

func setup() {
	protocol.CountMessages = true
	setupTopology()
	printStats("Before Test")
}

func tear() {
	printStats("After Test")
	shutdownTopology()
	time.Sleep(time.Second)
	printStats("After Shutdown")
}

func reset(name string) {
	Log.Info("*** ", name, " end ***")
	topo.ResetHandlers()
}

func setupTopology() {
	topo = NewTestTopology(4, 20000, 30000, 40000)
}

func shutdownTopology() {
	topo.Shutdown()
}

func printStats(tag string) {
	slog.Info(tag, " total Messages in session=", protocol.MessagesCreated())
	slog.Info(tag, " total Property Changed Called=", protocol.PropertyChangedCalled())
}
