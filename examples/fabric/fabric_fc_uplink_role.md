### Resource Creation

```hcl
resource "intersight_fabric_fc_uplink_role" "fabric_fc_uplink_role1" {
  aggregate_port_id = 0
  slot_id           = 100
  admin_speed       = "Auto"
  fill_pattern      = "Idle"
  vsanid            = 10
}
```