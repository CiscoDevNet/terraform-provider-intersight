### Resource Creation

```hcl
resource "intersight_kubernetes_cluster_profile" "kubernetes_cluster_profile1" {
  description = "kubernetes cluster profile"
  name        = "kubernetes_cluster_profile1"
  type        = "instance"
  action      = "Deploy"
  config_context {
    control_action = "Deploy"
    error_state    = "Pre-config-error"
  }
  managed_mode = "Managed"
  status       = "Deploying"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  workflow_info {
    object_type = "workflow.WorkflowInfo"
    moid        = var.workflow_workflow_info
  }
}

variable "workflow_workflow_info" {
  type        = string
  description = "Moid of the workflow.WorkflowInfo "
}
```