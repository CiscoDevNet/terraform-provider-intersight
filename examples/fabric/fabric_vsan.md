### Resource Creation

```hcl
resource "intersight_fabric_vsan" "fabric_vsan1" {
    default_zoning = "Enabled"
    fc_zone_sharing_mode = "Active"
    fcoe_vlan = 100
    name = "fabric_vsan1"
    vsan_id = 10
}
```