### Resource Creation

```hcl
resource "intersight_fabric_uplink_role" "fabric_uplink_role1" {
  aggregate_port_id = 0
  port_id           = 1
  slot_id           = 4
  admin_speed       = "Auto"
  fec               = "Auto"
}
```