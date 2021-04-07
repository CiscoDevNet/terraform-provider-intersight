### Resource Creation

```hcl
resource "intersight_vnic_fc_network_policy" "vnic_vnic_fc_network_policy" {
    name        = "vnic_fc_network_policy1"
    description = "vnic fabric channel network policy"
    vsan_settings {
        object_type = "vnic.VsanSettings"
        DefaultVlanId   = 10
        id          = 20
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```