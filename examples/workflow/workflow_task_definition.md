### Resource Creation

```hcl
resource "intersight_workflow_task_definition" "workflow_task_definition1" {
    name = "workflow_task_definition1"
    properties {
        object_type = "workflow.Properties"
        external_meta = true
        retry_count = 6
        retry_delay = 60
        retry_policy = "Fixed"
        support_status = "Supported"
        time_out = 60
        time_retry = "retry"
    }
    label = "inventory.ScopedInventoryTask"
    version = 1
    catalog {
        object_type = "workflow.Catalog"
        moid = var.workflow_catalog
    }
    interface_task {
        object_type = "workflow.TaskDefinition"
        moid = var.workflow_task_definition
    }
    task_metadata {
        object_type = "workflow.TaskMetadata"
        moid = var.workflow_task_metadata
    }

}
```