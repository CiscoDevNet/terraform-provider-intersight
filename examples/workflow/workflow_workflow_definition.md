### Resource Creation

```hcl
resource "intersight_workflow_workflow_definition" "workflow_workflow_definition1" {
  name  = "workflow_workflow_definition1"
  label = "Value for label"
  output_definition {
    class_id    = "workflow.StartTask"
    object_type = "workflow.PrimitiveDataType"
    name        = "status"
    label       = "Value for label"
    required    = false
    description = "Value for description"
    default {
      object_type = "workflow.DefaultValue"
      override    = false
      class_id    = "workflow.DefaultValue"
    }
    display_meta {
      class_id           = "workflow.DisplayMeta"
      object_type        = "workflow.DisplayMeta"
      inventory_selector = true
    }
  }
  tasks {
    description = "Create a vm"
    label       = "value for lable"
    name        = "StartTask"
    object_type = "workflow.StartTask"
    class_id    = "workflow.StartTask"
  }
  ui_rendering_data = <<EOT
  "{
    positions {
      name = "StartTask"
      x    = 209
      y    = 79
    }
    version = 1
    catalog = {
      object_type = "workflow.Catalog"
      moid        = var.workflow_catalog
    }
    workflow_metadata = {
      object_type = "workflow.WorkflowMetadata"
      moid        = var.workflow_workflow_metadata
    }
  }"
EOT
}

variable "workflow_catalog" {
  type        = string
  description = "<moid workflow workflow catalog>"
}

variable "workflow_workflow_metadata" {
  type        = string
  description = "<moid workflow workflow metadata>"
}
```