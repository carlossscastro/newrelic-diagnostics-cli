package minion

import (
	log "github.com/newrelic/newrelic-diagnostics-cli/logger"
	"github.com/newrelic/newrelic-diagnostics-cli/tasks"
)

// RegisterWith - will register any plugins in this package
func RegisterWith(registrationFunc func(tasks.Task, bool)) {
	log.Debug("Registering Synthetics/Minion/*")

	registrationFunc(SyntheticsMinionDetect{}, false)
	registrationFunc(SyntheticsMinionConfigValidate{}, false)
	registrationFunc(SyntheticsMinionHordeConnect{}, false)
	registrationFunc(SyntheticsMinionDetectCPM{executeCommand: tasks.CmdExecutor}, true)
	registrationFunc(SyntheticsMinionCollectLogs{executeCommand: tasks.BufferedCommandExec}, true)
	registrationFunc(SyntheticsMinionCollectK8sInfo{}, false)
}
