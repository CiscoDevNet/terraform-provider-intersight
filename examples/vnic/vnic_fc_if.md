### Resource Creation

```hcl
resource "intersight_vnic_fc_if" "fc1" {
  name = "fc0"
  order = 1
  placement {
    id = "1"
    pci_link = 0
    uplink = 0
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid = intersight_vnic_san_connectivity_policy.vnic_san1.id
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid = intersight_vnic_fc_network_policy.v_fc_network1.id
  }
  fc_adapter_policy {
    moid = intersight_vnic_fc_adapter_policy.v_fc_adapter1.id
  }
  fc_qos_policy {
    moid = intersight_vnic_fc_qos_policy.v_fc_qos1.id
  }
}
```