### Resource Creation

```hcl
resource "intersight_fabric_switch_profile" "fabric_switch_profile1" {
  name        = "fabric_switch_profile1"
  description = "fabric switch profile"
  type        = "instance"
  action      = "Deploy"
  config_context {
    object_type    = "policy.ConfigContext"
    control_action = "deploy"
    error_state    = "Config-error"
  }
}
```