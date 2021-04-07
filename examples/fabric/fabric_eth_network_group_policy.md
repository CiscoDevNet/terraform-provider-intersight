### Resource Creation

```hcl
resource "intersight_fabric_eth_network_group_policy" "fabric_eth_network_group_policy1" {
    name    = "AUTO_FabricEthNetworkGroupPolicy"
    target_platform     = "FIAttached"
    vlan_settings    {
        default_vlan    = 1
        allowed_vlans   = "313"
        mode            = "ACCESS"
    }
}
```