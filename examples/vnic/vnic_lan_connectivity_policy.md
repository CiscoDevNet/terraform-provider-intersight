### Resource Creation

```hcl
resource "intersight_vnic_lan_connectivity_policy" "vnic_lan_connectivity_policy" {
    name        = "vnic_lan_connectivity_policy1"
    description = "vnic lan connectivity policy"
    iqn_allocation_type = "Static"
    placement_mode      = "custom"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```