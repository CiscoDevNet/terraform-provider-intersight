resource "intersight_vnic_eth_adapter_policy" "v_eth_adapter1" {
  name = "v_eth_adapter1"
  rss_settings = true
  uplink_failback_timeout = 5
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  vxlan_settings {
    enabled = false
  }

  nvgre_settings {
    enabled = true
  }

  arfs_settings {
    enabled = true
  }

  interrupt_settings {
    coalescing_time = 125
    coalescing_type = "MIN"
    nr_count = 4
    mode = "MSI"
  }
  completion_queue_settings {
    nr_count = 4
    ring_size = 1
  }
  rx_queue_settings {
    nr_count = 4
    ring_size = 512
  }
  tx_queue_settings {
    nr_count = 4
    ring_size = 512
  }
  tcp_offload_settings {
    large_receive = true
    large_send = true
    rx_checksum = true
    tx_checksum = true
  }
}

resource "intersight_vnic_eth_network_policy" "v_eth_network1" {
  name = "v_eth_network1"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  vlan_settings {
    object_type = "vnic.VlanSettings"
    default_vlan = 1
    mode = "ACCESS"
  }
}

resource "intersight_vnic_eth_qos_policy" "v_eth_qos1" {
  name = "v_eth_qos1"
  mtu = 1500
  rate_limit = 0
  cos = 0
  trust_host_cos = false
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_vnic_lan_connectivity_policy" "vnic_lan1" {
  name = "vnic_lan1"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

resource "intersight_vnic_eth_if" "eth1" {
  name = "eth0"
  order = 0
  placement {
    id = "1"
    pci_link = 0
    uplink = 0
  }
  cdn {
    value = "VIC-1-eth00"
    nr_source = "user"
  }
  usnic_settings {
    cos = 5
    nr_count = 0
  }
  vmq_settings {
    enabled = true
    multi_queue_support = false
    num_interrupts = 1
    num_vmqs = 1
  }
  lan_connectivity_policy {
    moid = intersight_vnic_lan_connectivity_policy.vnic_lan1.id
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

resource "intersight_vnic_eth_if" "eth2" {
  name = "eth1"
  order = 0
  placement {
    id = "1"
    pci_link = 1
    uplink = 0
  }
  cdn {
    value = "VIC-1-eth1"
    nr_source = "vnic"
  }
  usnic_settings {
    cos = 5
    nr_count = 0
  }
  vmq_settings {
    enabled = true
    multi_queue_support = false
    num_interrupts = 1
    num_vmqs = 1
  }
  lan_connectivity_policy {
    moid = intersight_vnic_lan_connectivity_policy.vnic_lan1.id
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

resource "intersight_vnic_eth_if" "eth3" {
  name = "eth0"
  order = 0
  placement {
    id = "MLOM"
    pci_link = 0
    uplink = 1
  }
  cdn {
    value = "VIC-MLOM-eth0"
    nr_source = "vnic"
  }
  usnic_settings {
    cos = 5
    nr_count = 0
  }
  vmq_settings {
    enabled = false
    num_interrupts = 16
    num_sub_vnics = 64
    num_vmqs = 4
  }
  lan_connectivity_policy {
    moid = intersight_vnic_lan_connectivity_policy.vnic_lan1.id
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

resource "intersight_vnic_eth_if" "eth4" {
  name = "eth1"
  order = 1
  placement {
    id = "MLOM"
    pci_link = 0
    uplink = 1
  }
  cdn {
    value = "VIC-MLOM-eth1"
    nr_source = "vnic"
  }
  usnic_settings {
    cos = 5
    nr_count = 0
  }
  vmq_settings {
    enabled = true
    num_interrupts = 1
    num_vmqs = 1
  }
  lan_connectivity_policy {
    moid = intersight_vnic_lan_connectivity_policy.vnic_lan1.id
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

/*
SAMPLE PAYLOAD
-----------------
VnicEthIfApi: [{
  "Name": "eth0",
  "Cdn": {
    "Source": "user",
    "Value": "xyz0"
  },
  "Order": 0,
  "Placement": {
    "Type": "virtual",
    "Id": "MLOM",
    "Uplink": 1
  },
  "VmqSettings": {
    "Enabled": True
  },
  "LanConnectivityPolicy": {
    "ObjectType": "vnic.LanConnectivityPolicy",
    "Moid": ""
  },
  "EthNetworkPolicy": "",
  "EthQosPolicy": "",
  "EthAdapterPolicy": ""},
  {"Name": "eth1",
    "Cdn": {
      "Source": "user",
      "Value": "xyz1"
    },
    "Order": 1,
    "Placement": {
      "Type": "virtual",
      "Id": "MLOM",
      "Uplink": 1
    },
    "VmqSettings": {
      "Enabled": True
    },
    "LanConnectivityPolicy": {
      "ObjectType": "vnic.LanConnectivityPolicy",
      "Moid": ""
    },
    "EthNetworkPolicy": "",
    "EthQosPolicy": "",
    "EthAdapterPolicy": ""
  }]
*/