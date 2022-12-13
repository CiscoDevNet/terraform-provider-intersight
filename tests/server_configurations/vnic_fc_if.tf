resource "intersight_vnic_fc_adapter_policy" "tf_fc_adapter" {
  name = "tf_fc_adapter"
  error_detection_timeout = 100000
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
  error_recovery_settings {
    enabled = false
    io_retry_count = 255
    io_retry_timeout = 50
    link_down_timeout = 240000
    port_down_timeout = 240000
  }

  flogi_settings {
    retries = 0
    timeout = 255000
  }

  interrupt_settings {
    mode = "MSIx"
  }

  io_throttle_count = 1024
  lun_count = 1024
  lun_queue_depth = 254

  plogi_settings {
    retries = 255
    timeout = 255000
  }
  resource_allocation_timeout = 100000

  rx_queue_settings {
    nr_count = 1
    ring_size = 128
  }
  tx_queue_settings {
    nr_count = 1
    ring_size = 128
  }


  scsi_queue_settings {
    nr_count = 8
    ring_size = 152
  }

}

resource "intersight_vnic_fc_network_policy" "tf_fc_network" {
  name = "tf_fc_network"
  vsan_settings {
    id = 100
  }
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_vnic_fc_qos_policy" "tf_fc_qos" {
  name = "tf_fc_qos"
  rate_limit = 10000
  cos = 6
  max_data_field_size = 2112
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_vnic_san_connectivity_policy" "tf_vnic_san" {
  name = "tf_vnic_san"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

# resource "intersight_vnic_fc_if" "tf_fc1" {
#   name = "fc0"
#   order = 1
#   placement {
#     id = "1"
#     pci_link = 0
#     uplink = 0
#   }
#   persistent_bindings = true
#   san_connectivity_policy {
#     moid = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
#   }
#   fc_network_policy {
#     moid = intersight_vnic_fc_network_policy.tf_fc_network.id
#   }
#   fc_adapter_policy {
#     moid = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
#   }
#   fc_qos_policy {
#     moid = intersight_vnic_fc_qos_policy.tf_fc_qos.id
#   }
# }

# resource "intersight_vnic_fc_if" "tf_fc2" {
#   name = "fc0"
#   order = 2
#   placement {
#     id = "MLOM"
#     pci_link = 0
#     uplink = 0
#   }
#   persistent_bindings = true
#   san_connectivity_policy {
#     moid = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
#   }
#   fc_network_policy {
#     moid = intersight_vnic_fc_network_policy.tf_fc_network.id
#   }
#   fc_adapter_policy {
#     moid = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
#   }
#   fc_qos_policy {
#     moid = intersight_vnic_fc_qos_policy.tf_fc_qos.id
#   }
# }

# resource "intersight_vnic_fc_if" "tf_fc3" {
#   name = "fc1"
#   order = 3
#   placement {
#     id = "MLOM"
#     pci_link = 0
#     uplink = 1
#   }
#   persistent_bindings = true
#   san_connectivity_policy {
#     moid = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
#     object_type = "vnic.SanConnectivityPolicy"
#   }
#   fc_network_policy {
#     moid = intersight_vnic_fc_network_policy.tf_fc_network.id
#   }
#   fc_adapter_policy {
#     moid = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
#   }
#   fc_qos_policy {
#     moid = intersight_vnic_fc_qos_policy.tf_fc_qos.id
#   }
# }

# resource "intersight_vnic_fc_if" "tf_fc4" {
#   name = "fc1"
#   order = 1
#   placement {
#     id = "1"
#     pci_link = 1
#     uplink = 1
#   }
#   persistent_bindings = true
#   san_connectivity_policy {
#     moid = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
#   }
#   fc_network_policy {
#     moid = intersight_vnic_fc_network_policy.tf_fc_network.id
#   }
#   fc_adapter_policy {
#     moid = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
#   }
#   fc_qos_policy {
#     moid = intersight_vnic_fc_qos_policy.tf_fc_qos.id
#   }
# }