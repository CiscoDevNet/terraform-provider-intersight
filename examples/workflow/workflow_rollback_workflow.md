### Resource Creation

```hcl
resource "intersight_workflow_rollback_workflow" {
    action = "start"
    continue_on_Task_failure = false
    primary_workflow {
        object_type = "workflow/WorkflowInfos"
        moid = var.workflow_workflowInfos
    }
    selected_tasks = [{
        object_type = "workflow.RollbackWorkflowTask"
        name = "taskInfo_name"
        ref_name = "reference_name_task_info_rollback"
        task_info_moid = var.workflow.rollback_workflowTask
    }]
}
```