### Resource Creation

```hcl
resource "intersight_fabric_uplink_pc_role" "fabric_uplink_pc_role1" {
    pcid = 100
    ports [
        {
            port_id = 1
            aggregate_port_id = 0
            slot_id = 100
        }
    ]
    admin_speed = "Auto"
    ud_ld_admin_state = "Enabled"
}
```