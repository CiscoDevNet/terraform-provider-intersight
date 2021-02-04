### Resource Creation
```hcl
resource "intersight_vnic_eth_network_policy" "v_eth_network1" {
  name = "v_eth_network1"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  vlan_settings {
    object_type = "vnic.VlanSettings"
    default_vlan = 1
    mode = "ACCESS"
  }
}
```