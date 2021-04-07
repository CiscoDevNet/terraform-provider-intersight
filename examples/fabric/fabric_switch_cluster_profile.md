### Resource Creation

```hcl
resource "intersight_fabric_switch_cluster_profile" "fabric_switch_cluster_profile1" {
  name        = "fabric_switch_cluster_profile"
  description = "fabric switch cluster profile"
  type        = "instance"
  config_context {
    object_type    = "policy.ConfigContext"
    control_action = "deploy"
    error_state    = "Config-error"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```