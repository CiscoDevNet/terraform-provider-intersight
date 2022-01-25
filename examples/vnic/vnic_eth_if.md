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
    moid        = var.vnic_lan1
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid = var.v_eth_network1
  }
  eth_adapter_policy {
    moid = var.v_eth_adapter1
  }
  eth_qos_policy {
    moid = var.v_eth_qos1
  }
}

variable "vnic_lan1"{
  type = string
  description = "Moid of vnic.LanConnectivityPolicy"
}

variable "v_eth_network1"{
  type = string
  description = "Moid of vnic.EthNetworkPolicy"
}

variable "v_eth_adapter1"{
  type = string
  description = "Moid of vnic.EthAdapterPolicy"
}

variable "v_eth_qos1"{
  type = string
  description = "Moid of vnic.EthQosPolicy"
}
```