### Resource Creation

```hcl
resource "intersight_vnic_eth_if" "eth1" {
  name  = "eth0"
  order = 0
  placement {
    id       = "1"
    pci_link = 0
    uplink   = 0
  }
  cdn {
    value     = "VIC-1-eth00"
    nr_source = "user"
  }
  usnic_settings {
    cos      = 5
    nr_count = 0
  }
  vmq_settings {
    enabled             = true
    multi_queue_support = false
    num_interrupts      = 1
    num_vmqs            = 1
  }
  lan_connectivity_policy {
    moid        = intersight_vnic_lan_connectivity_policy.vnic_lan1.id
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid = intersight_vnic_eth_network_policy.v_eth_network1.id
  }
  eth_adapter_policy {
    moid = intersight_vnic_eth_adapter_policy.v_eth_adapter1.id
  }
  eth_qos_policy {
    moid = intersight_vnic_eth_qos_policy.v_eth_qos1.id
  }
}
```