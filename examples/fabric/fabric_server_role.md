### Resource Creation

```hcl
resource "intersight_fabric_server_role" "fabric_server_role1" {
  aggregate_port_id = 0
  port_id           = 16
  slot_id           = 1
  port_policy {
    moid = intersight_fabric_port_policy.fabric_port_policy1.moid
  }
}
```