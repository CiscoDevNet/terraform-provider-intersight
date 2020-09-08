resource "intersight_vnic_fc_adapter_policy" "v_fc_adapter1" {
  name                    = "v_fc_adapter1"
  error_detection_timeout = 100000
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  error_recovery_settings {
    enabled           = false
    io_retry_count    = 255
    io_retry_timeout  = 50
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
  lun_count         = 1024
  lun_queue_depth   = 254

  plogi_settings {
    retries = 255
    timeout = 255000
  }
  resource_allocation_timeout = 100000

  rx_queue_settings {
    nr_count     = 1
    ring_size = 128
  }
  tx_queue_settings {
    nr_count     = 1
    ring_size = 128
  }


  scsi_queue_settings {
    nr_count     = 8
    ring_size = 152
  }

}

resource "intersight_vnic_fc_network_policy" "v_fc_network1" {
  name = "v_fc_network1"
  vsan_settings {
    id = 100
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_vnic_fc_qos_policy" "v_fc_qos1" {
  name                = "v_fc_qos1"
  rate_limit          = 10000
  cos                 = 6
  max_data_field_size = 2112
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_vnic_san_connectivity_policy" "vnic_san1" {
  name = "vnic_san1"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }

}

resource "intersight_vnic_fc_if" "fc1" {
  name  = "fc0"
  order = 1
  placement {
    id = "1"
    pci_link = 0
    uplink = 0
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.vnic_san1.id
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

resource "intersight_vnic_fc_if" "fc2" {
  name  = "fc0"
  order = 2
  placement {
    id = "MLOM"
    pci_link = 0
    uplink = 0
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.vnic_san1.id
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

resource "intersight_vnic_fc_if" "fc3" {
  name  = "fc1"
  order = 3
  placement {
    id = "MLOM"
    pci_link = 0
    uplink = 1
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.vnic_san1.id
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

resource "intersight_vnic_fc_if" "fc4" {
  name  = "fc1"
  order = 1
  placement {
    id = "1"
    pci_link = 1
    uplink = 1
  }
  persistent_bindings = true
  san_connectivity_policy {
    moid        = intersight_vnic_san_connectivity_policy.vnic_san1.id
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

/*
SAMPLE PAYLOAD
-----------------
VnicFcIfApi: [{
    "Name": "fc0",
    "Order": 2,
    "Placement": {
        "Type": "virtual",
        "Id": "1"
    },
    "PersistentBindings": True,
    "SanConnectivityPolicy": "",
    "FcNetworkPolicy": "",
                       "FcQosPolicy": "",
                       "FcAdapterPolicy": ""
},
    {
    "Name": "fc1",
    "Order": 3,
    "Placement": {
        "Type": "virtual",
        "Id": "1"
    },
    "PersistentBindings": True,
    "SanConnectivityPolicy":
    "",
    "FcNetworkPolicy": "",
    "FcQosPolicy": "",
    "FcAdapterPolicy": ""
}]
*/