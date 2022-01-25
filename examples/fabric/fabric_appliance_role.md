### Resource Creation

```hcl
resource "intersight_fabric_appliance_role" "fabric_appliance_role1" {
  aggregate_port_id = 0
  slot_id           = 4
  admin_speed       = "Auto"
  fec               = "Auto"
  mode              = "trunk"
  priority          = "Best Effort"
}
```