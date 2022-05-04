resource "intersight_vnic_eth_adapter_policy" "tf_v_eth_adapter" {
  name                    = "tf_v_eth_adapter"
  rss_settings            = true
  uplink_failback_timeout = 5
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  vxlan_settings {
    enabled     = false
    object_type = "vnic.VxlanSettings"
  }

  nvgre_settings {
    enabled     = true
    object_type = "vnic.NvgreSettings"
  }

  arfs_settings {
    enabled     = true
    object_type = "vnic.ArfsSettings"
  }

  interrupt_settings {
    coalescing_time = 125
    coalescing_type = "MIN"
    nr_count        = 4
    mode            = "MSI"
    object_type     = "vnic.EthInterruptSettings"
  }
  completion_queue_settings {
    nr_count    = 4
    object_type = "vnic.CompletionQueueSettings"
  }
  rx_queue_settings {
    nr_count    = 4
    ring_size   = 512
    object_type = "vnic.EthRxQueueSettings"
  }
  tx_queue_settings {
    nr_count    = 4
    ring_size   = 512
    object_type = "vnic.EthTxQueueSettings"
  }
  tcp_offload_settings {
    large_receive = true
    large_send    = true
    rx_checksum   = true
    tx_checksum   = true
    object_type   = "vnic.TcpOffloadSettings"
  }
}

resource "intersight_vnic_eth_network_policy" "tf_v_eth_network" {
  name = "tf_v_eth_network"
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  vlan_settings {
    object_type  = "vnic.VlanSettings"
    default_vlan = 1
    mode         = "ACCESS"
  }
}

resource "intersight_vnic_eth_qos_policy" "tf_v_eth_qos" {
  name           = "tf_v_eth_qos"
  mtu            = 1500
  rate_limit     = 0
  cos            = 0
  trust_host_cos = false
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_vnic_lan_connectivity_policy" "tf_vnic_lan" {
  name = "tf_vnic_lan"
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_vnic_eth_if" "tf_eth1" {
  name  = "eth0"
  order = 0
  placement {
    id          = "1"
    pci_link    = 0
    uplink      = 0
    object_type = "vnic.PlacementSettings"
  }
  cdn {
    value       = "VIC-1-eth00"
    nr_source   = "user"
    object_type = "vnic.Cdn"
  }
  usnic_settings {
    cos         = 5
    nr_count    = 0
    object_type = "vnic.UsnicSettings"
  }
  vmq_settings {
    enabled             = true
    multi_queue_support = false
    num_interrupts      = 1
    num_vmqs            = 1
    object_type         = "vnic.VmqSettings"
  }
  lan_connectivity_policy {
    moid        = intersight_vnic_lan_connectivity_policy.tf_vnic_lan.id
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid        = intersight_vnic_eth_network_policy.tf_v_eth_network.id
    object_type = "vnic.EthNetworkPolicy"
  }
  eth_adapter_policy {
    moid        = intersight_vnic_eth_adapter_policy.tf_v_eth_adapter.id
    object_type = "vnic.EthAdapterPolicy"
  }
  eth_qos_policy {
    moid        = intersight_vnic_eth_qos_policy.tf_v_eth_qos.id
    object_type = "vnic.EthQosPolicy"
  }
}

resource "intersight_vnic_eth_if" "tf_eth2" {
  name  = "eth1"
  order = 0
  placement {
    id          = "1"
    pci_link    = 1
    uplink      = 0
    object_type = "vnic.PlacementSettings"
  }
  cdn {
    value       = "VIC-1-eth1"
    nr_source   = "vnic"
    object_type = "vnic.Cdn"
  }
  usnic_settings {
    cos         = 5
    nr_count    = 0
    object_type = "vnic.UsnicSettings"
  }
  vmq_settings {
    enabled             = true
    multi_queue_support = false
    num_interrupts      = 1
    num_vmqs            = 1
    object_type         = "vnic.VmqSettings"
  }
  lan_connectivity_policy {
    moid        = intersight_vnic_lan_connectivity_policy.tf_vnic_lan.id
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid        = intersight_vnic_eth_network_policy.tf_v_eth_network.id
    object_type = "vnic.EthNetworkPolicy"
  }
  eth_adapter_policy {
    moid        = intersight_vnic_eth_adapter_policy.tf_v_eth_adapter.id
    object_type = "vnic.EthAdapterPolicy"
  }
  eth_qos_policy {
    moid        = intersight_vnic_eth_qos_policy.tf_v_eth_qos.id
    object_type = "vnic.EthQosPolicy"
  }
}

resource "intersight_vnic_eth_if" "tf_eth3" {
  name  = "eth0"
  order = 0
  placement {
    id          = "MLOM"
    pci_link    = 0
    uplink      = 1
    object_type = "vnic.PlacementSettings"
  }
  cdn {
    value       = "VIC-MLOM-eth0"
    nr_source   = "vnic"
    object_type = "vnic.Cdn"
  }
  usnic_settings {
    cos         = 5
    nr_count    = 0
    object_type = "vnic.UsnicSettings"
  }
  vmq_settings {
    enabled        = false
    num_interrupts = 16
    num_sub_vnics  = 64
    num_vmqs       = 4
    object_type    = "vnic.VmqSettings"
  }
  lan_connectivity_policy {
    moid        = intersight_vnic_lan_connectivity_policy.tf_vnic_lan.id
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid        = intersight_vnic_eth_network_policy.tf_v_eth_network.id
    object_type = "vnic.EthNetworkPolicy"
  }
  eth_adapter_policy {
    moid        = intersight_vnic_eth_adapter_policy.tf_v_eth_adapter.id
    object_type = "vnic.EthAdapterPolicy"
  }
  eth_qos_policy {
    moid        = intersight_vnic_eth_qos_policy.tf_v_eth_qos.id
    object_type = "vnic.EthQosPolicy"
  }
}

resource "intersight_vnic_eth_if" "tf_eth4" {
  name  = "eth1"
  order = 1
  placement {
    id          = "MLOM"
    pci_link    = 0
    uplink      = 1
    object_type = "vnic.PlacementSettings"
  }
  cdn {
    value       = "VIC-MLOM-eth1"
    nr_source   = "vnic"
    object_type = "vnic.Cdn"
  }
  usnic_settings {
    cos         = 5
    nr_count    = 0
    object_type = "vnic.UsnicSettings"
  }
  vmq_settings {
    enabled        = true
    num_interrupts = 1
    num_vmqs       = 1
    object_type    = "vnic.VmqSettings"
  }
  lan_connectivity_policy {
    moid        = intersight_vnic_lan_connectivity_policy.tf_vnic_lan.id
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid        = intersight_vnic_eth_network_policy.tf_v_eth_network.id
    object_type = "vnic.EthNetworkPolicy"
  }
  eth_adapter_policy {
    moid        = intersight_vnic_eth_adapter_policy.tf_v_eth_adapter.id
    object_type = "vnic.EthAdapterPolicy"
  }
  eth_qos_policy {
    moid        = intersight_vnic_eth_qos_policy.tf_v_eth_qos.id
    object_type = "vnic.EthQosPolicy"
  }
}