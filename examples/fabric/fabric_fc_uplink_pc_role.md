### Resource Creation

```hcl
resource "intersight_fabric_fc_uplink_pc_role" "fabric_fc_uplink_pc_role1" {
  admin_speed  = "Auto"
  fill_pattern = "Idle"
  vsan_id      = 10
}
```