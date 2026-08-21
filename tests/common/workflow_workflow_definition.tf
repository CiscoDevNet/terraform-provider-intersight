resource "intersight_workflow_workflow_definition" "ProvisionNetwork" {
  name = "ProvisionNetwork"
  label = "TF Testing"
  default_version = true
  output_definition {
    object_type = "workflow.PrimitiveDataType"
    name        = "status"
    label       = "status"
    required    = false
    default {
      object_type = "workflow.DefaultValue"
      override    = false
      value       = null
    }
    display_meta {
      object_type        = "workflow.DisplayMeta"
      inventory_selector = true
    }
  }
  tasks {
    description = "Create a vm"
    label       = "StartTask"
    name        = "StartTask"
    object_type = "workflow.StartTask"
  }
  ui_rendering_data = jsonencode({
    Positions = {
      Name = "StartTask"
      X    = 209
      Y    = 79
    }
    Version = 1
  })
}
