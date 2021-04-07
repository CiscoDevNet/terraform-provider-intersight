### Resource Creation

```hcl
resource "intersight_fabric_vlan" "fabric_vlan1" {
  auto_allow_on_uplinks = true
  is_native             = true
  name                  = "fabric_vlan1"
  vlan_id               = 10
  eth_network_policy {
    moid = intersight_vnic_eth_network_policy.v_eth_network1.id
  }
}
```