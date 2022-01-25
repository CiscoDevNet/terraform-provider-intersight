### Resource Creation

```hcl
resource "intersight_workflow_workflow_info" "workflow_workflow_info1" {
  name         = "workflow_workflow_info1"
  pause_reason = null
  action       = "Create"
  properties {
    object_type     = "workflow.WorkflowInfoProperties"
    retryable       = false
    rollback_action = "Disabled"
  }
  success_workflow_cleanup_duration = 2160
  wait_reason                       = null
  workflow_meta_type                = "SystemDefined"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  workflow_definition {
    object_type = "workflow.WorkflowDefinition"
    moid        = var.workflow_workflow_definition
  }
}

 variable "workflow_workflow_definition" {
   type = string
   description = "<moid workflow workflow definition>"
 }
```