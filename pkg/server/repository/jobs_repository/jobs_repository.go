package jobs_repository

import (
	"github.com/ovvesley/scientific-workflow-k8s/pkg/server/connector"
	"github.com/ovvesley/scientific-workflow-k8s/pkg/server/repository"
	"github.com/ovvesley/scientific-workflow-k8s/pkg/server/workflow"
)

type ActivityRepository struct {
	tableName string
}

var TableName = "activities"
var Columns = "(ID INTEGER PRIMARY KEY AUTOINCREMENT, workflow_id INTEGER, namespace TEXT, name TEXT, raw_activity TEXT, status INTEGER, depend_on_activity INTEGER)"

func New() *ActivityRepository {

	database := connector.Database{}
	c := database.Connect()
	err := repository.CreateOrVerifyTable(c, TableName, Columns)
	if err != nil {
		return nil
	}

	err = c.Close()
	if err != nil {
		return nil
	}

	return &ActivityRepository{
		tableName: TableName,
	}
}

func (w *ActivityRepository) Create(namespace string, workflowId int, activities []workflow.WorkflowActivities) error {

	database := connector.Database{}
	c := database.Connect()

	var lastActivity int64 = 0
	for _, activity := range activities {
		rawActivity := activity.GetBase64Activities()
		result, err := c.Exec("INSERT INTO "+w.tableName+" (workflow_id, namespace, name, raw_activity, status, depend_on_activity) VALUES (?, ?, ?, ?, ?, ?)", workflowId, namespace, activity.Name, rawActivity, 0, lastActivity)
		if err != nil {
			return err
		}
		lastActivity, _ = result.LastInsertId()
	}

	err := c.Close()

	if err != nil {
		return err
	}
	return nil
}
