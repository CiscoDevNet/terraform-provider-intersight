resource "intersight_vnic_fc_adapter_policy" "tf_fc_adapter" {
  name                    = "tf_fc_adapter"
  error_detection_timeout = 100000
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  error_recovery_settings {
    enabled           = false
    io_retry_count    = 255
    io_retry_timeout  = 50
    link_down_timeout = 240000
    port_down_timeout = 240000
    object_type       = "vnic.FcErrorRecoverySettings"
  }

  flogi_settings {
    retries     = 0
    timeout     = 255000
    object_type = "vnic.FlogiSettings"
  }

  interrupt_settings {
    mode        = "MSIx"
    object_type = "vnic.FcInterruptSettings"
  }

  io_throttle_count = 1024
  lun_count         = 1024
  lun_queue_depth   = 254

  plogi_settings {
    retries     = 255
    timeout     = 255000
    object_type = "vnic.PlogiSettings"
  }
  resource_allocation_timeout = 100000

  rx_queue_settings {
    ring_size   = 128
    object_type = "vnic.FcQueueSettings"
  }
  tx_queue_settings {
    ring_size   = 128
    object_type = "vnic.FcQueueSettings"
  }

  scsi_queue_settings {
    nr_count    = 8
    ring_size   = 152
    object_type = "vnic.ScsiQueueSettings"
  }

}

resource "intersight_vnic_fc_network_policy" "tf_fc_network" {
  name = "tf_fc_network"
  vsan_settings {
    id          = 100
    object_type = "vnic.VsanSettings"
  }
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_vnic_fc_qos_policy" "tf_fc_qos" {
  name                = "tf_fc_qos"
  rate_limit          = 10000
  cos                 = 6
  max_data_field_size = 2112
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_vnic_san_connectivity_policy" "tf_vnic_san" {
  name = "tf_vnic_san"
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_vnic_fc_if" "tf_fc1" {
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
    moid        = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid        = intersight_vnic_fc_network_policy.tf_fc_network.id
    object_type = "vnic.FcNetworkPolicy"
  }
  fc_adapter_policy {
    moid        = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
    object_type = "vnic.FcAdapterPolicy"
  }
  fc_qos_policy {
    moid        = intersight_vnic_fc_qos_policy.tf_fc_qos.id
    object_type = "vnic.FcQosPolicy"
  }
}

resource "intersight_vnic_fc_if" "tf_fc2" {
  name  = "fc0"
  order = 2
  placement {
    id          = "MLOM"
    pci_link    = 0
    uplink      = 0
    object_type = "vnic.PlacementSettings"
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid        = intersight_vnic_fc_network_policy.tf_fc_network.id
    object_type = "vnic.FcNetworkPolicy"
  }
  fc_adapter_policy {
    moid        = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
    object_type = "vnic.FcAdapterPolicy"
  }
  fc_qos_policy {
    moid        = intersight_vnic_fc_qos_policy.tf_fc_qos.id
    object_type = "vnic.FcQosPolicy"
  }
}

resource "intersight_vnic_fc_if" "tf_fc3" {
  name  = "fc1"
  order = 3
  placement {
    id          = "MLOM"
    pci_link    = 0
    uplink      = 1
    object_type = "vnic.PlacementSettings"
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid        = intersight_vnic_fc_network_policy.tf_fc_network.id
    object_type = "vnic.FcNetworkPolicy"
  }
  fc_adapter_policy {
    moid        = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
    object_type = "vnic.FcAdapterPolicy"
  }
  fc_qos_policy {
    moid        = intersight_vnic_fc_qos_policy.tf_fc_qos.id
    object_type = "vnic.FcQosPolicy"
  }
}

resource "intersight_vnic_fc_if" "tf_fc4" {
  name  = "fc1"
  order = 1
  placement {
    id          = "1"
    pci_link    = 1
    uplink      = 1
    object_type = "vnic.PlacementSettings"
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.tf_vnic_san.id
    object_type = "vnic.SanConnectivityPolicy"
  }
  fc_network_policy {
    moid        = intersight_vnic_fc_network_policy.tf_fc_network.id
    object_type = "vnic.FcNetworkPolicy"
  }
  fc_adapter_policy {
    moid        = intersight_vnic_fc_adapter_policy.tf_fc_adapter.id
    object_type = "vnic.FcAdapterPolicy"
  }
  fc_qos_policy {
    moid        = intersight_vnic_fc_qos_policy.tf_fc_qos.id
    object_type = "vnic.FcQosPolicy"
  }
}