### Resource Creation

```hcl
resource "intersight_workflow_batch_api_executor" "workflow_batch_api_executor1" {
  name                  = "workflow_batch_api_executor1"
  retry_from_failed_api = false
  task_definition {
    object_type = "workflow.TaskDefinition"
    moid        = var.workflow_task_definition
  }
}
```