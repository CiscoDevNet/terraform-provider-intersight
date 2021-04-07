### Resource Creation

```hcl
resource "intersight_vnic_san_connectivity_policy" "vnic_san_connectivity_policy" {
    name        = "vnic_san_connectivity_policy1"
    description = "vnic san connectivity policy"
    placement_mode      = "custom"
    static_wwpn_address = "500a09829697c3ac"
    target_platform     = "FIAttached"
    wwpn_address_type   = "STATIC"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```