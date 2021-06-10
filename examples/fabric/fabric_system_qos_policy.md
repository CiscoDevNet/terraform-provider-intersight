### Resource Creation

```hcl
resource "intersight_fabric_system_qos_policy" "fabric_system_qos_policy1" {
  name        = "fabric_system_qos_policy1"
  description = "demo fabric system qos policy"
  classes = [
    {
      admin_state        = "Enabled"
      bandwidth_percent  = 80
      cos                = 1
      mtu                = 2240
      multicast_optimize = true
      name               = "Best Effort"
      packet_drop        = true
    }
  ]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```