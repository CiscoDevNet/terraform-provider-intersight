### Resource Creation

```hcl
resource "intersight_fabric_vlan" "fabric_vlan1" {
  auto_allow_on_uplinks = true
  is_native             = true
  name                  = "fabric_vlan1"
  vlan_id               = 10
  eth_network_policy {
    moid = var.vnic_eth_network
  }
  multicast_policy {
    moid = var.fabric_multicast_policy
  }
}

variable "vnic_eth_network" {
  type        = string
  description = "Moid of vnic_eth_network"
}

variable "fabric_multicast_policy" {
  type        = string
  description = "Moid of fabric_multicast_policy"
}
```