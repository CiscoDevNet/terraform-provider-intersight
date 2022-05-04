### Resource Creation

```hcl
resource "intersight_vnic_fc_if" "fc1" {
  name  = "fc0"
  order = 1
  placement {
    id          = "1"
    pci_link    = 0
    uplink      = 0
    object_type = "vnic.PlacementSettings"
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = var.vnic_san1
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid        = var.v_fc_network1
    object_type = "vnic.FcNetworkPolicy"
  }
  fc_adapter_policy {
    moid        = var.v_fc_adapter1
    object_type = "vnic.FcAdapterPolicy"
  }
  fc_qos_policy {
    moid        = var.v_fc_qos1
    object_type = "vnic.FcQosPolicy"
  }
}

variable "vnic_san1" {
  type        = string
  description = "Moid of vnic.SanConnectivityPolicy"
}

variable "v_fc_network1" {
  type        = string
  description = "Moid of vnic.FcNetworkPolicy"
}

variable "v_fc_adapter1" {
  type        = string
  description = "Moid of vnic.FcAdapterPolicy"
}

variable "v_fc_qos1" {
  type        = string
  description = "Moid of vnic.FcQosPolicy"
}
```