### Resource Creation

```hcl
resource "intersight_workflow_workflow_definition" "workflow_workflow_definition1" {
    name = "workflow_workflow_definition1"
    output_definition = [{
        object_type = "workflow.PrimitiveDataType"
        name = "status"
        label = "status"
        required = false
        default {
            object_type = "workflow.DefaultValue"
            override = false
            value = null
        }
        display_meta {
            object_type = "workflow.DisplayMeta"
            inventory_selector = true
            widget_type = None
        }
    }]
    tasks = [{
        description = "Create a vm"
        label = "StartTask"
        name = "StartTask"
        next_task = "NewCloudVirtualMachineAndMonitor1"
        object_type = "workflow.StartTask"
    }]
    ui_rendering_data {
        "Positions" = [{
            name = "StartTask"
            x = 209
            y = 79
        }]
        version = 1
        catalog {
            object_type = "workflow.Catalog"
            moid = var.workflow_catalog
        }
        workflow_metadata {
            object_type = "workflow.WorkflowMetadata"
            moid = var.workflow_workflow_metadata
        }
    }
}
```