package monitor

import (
	"github.com/ovvesley/scientific-workflow-k8s/pkg/server/services/monitor_change_workflow_service"
	"time"
)

const TimeToUpdateSeconds = 5

func StartMonitor() {
	for {
		handleMonitor()
		time.Sleep(TimeToUpdateSeconds * time.Second)

	}
}

func handleMonitor() {
	monitorChangeWorkflowService := monitor_change_workflow_service.New()
	monitorChangeWorkflowService.MonitorChangeWorkflow()
}
