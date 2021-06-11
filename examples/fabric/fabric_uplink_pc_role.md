### Resource Creation

```hcl
resource "intersight_fabric_uplink_pc_role" "fabric_uplink_pc_role1" {
  pcid = 100
  ports = [
    {
      port_id           = 1
      aggregate_port_id = 0
      slot_id           = 1
    },
    {
      port_id           = 2
      aggregate_port_id = 0
      slot_id           = 1
    }
  ]
  admin_speed       = "Auto"
  port_policy {
    moid = intersight_fabric_port_policy.fabric_port_policy1.moid
  }
}
```