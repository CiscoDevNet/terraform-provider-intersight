### Resource Creation

```hcl
resource "intersight_fabric_system_qos_policy" "fabric_system_qos_policy1" {
  name        = "fabric_system_qos_policy1"
  description = "demo fabric system qos policy"
  classes {
    admin_state        = "Enabled"
    bandwidth_percent  = 80
    cos                = 255
    mtu                = 2240
    multicast_optimize = true
    name               = "Best Effort"
    packet_drop        = true
    class_id           = "fabric.QosClass"
    object_type        = "fabric.QosClass"
    weight             = 5
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```