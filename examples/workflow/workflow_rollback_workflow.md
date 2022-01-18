### Resource Creation

```hcl
resource "intersight_workflow_rollback_workflow" "workflow_rollback_workflow" {
  action                   = "start"
  continue_on_task_failure = false
  primary_workflow {
    object_type = "workflow/WorkflowInfos"
    moid        = intersight_workflow_workflow_info.workflow_workflow_info1.moid
  }
  selected_tasks = [{
    additional_properties = ""
    class_id              = "workflow.RollbackWorkflowTask"
    description           = ""
    rollback_completed    = false
    rollback_task_name    = ""
    status                = ""
    task_path             = ""
    object_type           = "workflow.RollbackWorkflowTask"
    name                  = "taskInfo_name"
    ref_name              = "reference_name_task_info_rollback"
    task_info_moid        = intersight_workflow_workflow_info.workflow_workflow_info1.moid
  }]
}
```