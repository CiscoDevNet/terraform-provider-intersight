### Resource Creation

```hcl
resource "intersight_fabric_appliance_pc_role" "fabric_appliance_pc_role1" {
  pc_id       = 100
  admin_speed = "Auto"
  mode        = "trunk"
  priority    = "Best Effort"
}
```