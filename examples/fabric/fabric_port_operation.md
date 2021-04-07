### Resource Creation

```hcl
resource "intersight_fabric_port_operation" "fabric_port_operation1" {
    aggregate_port_id = 0
    port_id = 1
    slot_id = 100
    config_state = "Applied"
}
```