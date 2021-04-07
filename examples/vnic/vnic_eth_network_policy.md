### Resource Creation

```hcl
resource "intersight_vnic_eth_network_policy" "vnic_eth_network_policy" {
    name        = "vnic_ethernet_network_policy1"
    description = "vnic ethernet network policy"
    target_platform = "FIAttached"
    vlan_settings   = {
        object_type = "vnic.VlanSettings"
        default_vlan    = 10
        mode        = "TRUNK"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```