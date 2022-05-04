### Resource Creation

```hcl
resource "intersight_fabric_switch_control_policy" "fabric_switch_control_policy1" {
  name        = "fabric_switch_control_policy1"
  description = "fabric switch control policy"
  mac_aging_settings {
    mac_aging_option = "Custom"
    mac_aging_time   = 3000
    object_type      = "fabric.MacAgingSettings"
  }
  vlan_port_optimization_enabled = true
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```