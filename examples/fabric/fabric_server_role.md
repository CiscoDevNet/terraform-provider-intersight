### Resource Creation

```hcl
resource "intersight_fabric_server_role" "fabric_server_role1" {
  aggregate_port_id = 0
  port_id           = 1
  slot_id           = 100
}
```