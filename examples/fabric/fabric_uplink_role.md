### Resource Creation

```hcl
resource "intersight_fabric_uplink_role" "fabric_uplink_role1" {
    aggregate_port_id = 0
    port_id = 1
    slot_id = 100
    admin_speed = "Auto"
    fec = "Auto"
    ud_ld_admin_state = "Enabled"
}
```