### Resource Creation

```hcl
resource "intersight_fabric_port_mode" "fabric_port_mode1" {
  custom_mode   = "FibreChannel"
  port_id_end   = 5000
  port_id_start = 1000
  slot_id       = 100
}
```