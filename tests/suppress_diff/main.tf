terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.22"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}

data "intersight_organization_organization" "default" {
  name = "default"
}

resource "intersight_hyperflex_cluster_profile" "hyperflex_cluster_profile1" {
  storage_data_vlan {
    name    = "hx-storage-data"
    vlan_id = 27
  }
  mgmt_ip_address    = "10.225.68.237"
  mac_address_prefix = "00:25:B5:D5"
  mgmt_platform      = "EDGE"
  description        = "This is hyperflex cluster profile"
  tags {
    key   = "test"
    value = "ucsback-10G-3nodehx-cluster-"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  name = "hyperflex_cluster_profile1"
}

resource "intersight_networkconfig_policy" "network_config1" {
  name                     = "network_config1"
  description              = "demo network configuration policy"
  enable_dynamic_dns       = false
  preferred_ipv6dns_server = null
  enable_ipv6              = false
  enable_ipv6dns_from_dhcp = false
  preferred_ipv4dns_server = "10.10.10.1"
  alternate_ipv4dns_server = "10.10.10.1"
  alternate_ipv6dns_server = null
  dynamic_dns_domain       = ""
  enable_ipv4dns_from_dhcp = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_vnic_eth_adapter_policy" "vnic_eth_adapter_policy_1" {
  name                    = "vnic_eth_adapter_policy_11"
  rss_settings            = true
  uplink_failback_timeout = 5
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }

  nvgre_settings {
    enabled = true
  }

  arfs_settings {
    enabled = true
  }

  roce_settings {
    enabled = false
  }

  interrupt_settings {
    coalescing_time = 125
    coalescing_type = "MIN"
    nr_count        = 4
    mode            = "MSI"
  }
  completion_queue_settings {
    nr_count  = 4
    ring_size = 1
  }
  rx_queue_settings {
    nr_count  = 4
    ring_size = 512
  }
  tx_queue_settings {
    nr_count  = 4
    ring_size = 512
  }
  tcp_offload_settings {
    large_receive = true
    large_send    = true
    rx_checksum   = true
    tx_checksum   = true
  }
}

resource "intersight_kubernetes_virtual_machine_infra_config_policy" "vm_infra_config_policy" {
  name        = "TFC-Test-1"
  description = "VM Infrastructure Config Policy used for Kubernetes deployment to vSphere"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results[0].moid
  }
  vm_config {
    object_type = "kubernetes.EsxiVirtualMachineInfraConfig"
	interfaces = ["8", "4"]
    additional_properties = jsonencode({
      Cluster    =  "abc"
      Datastore  = "datastore"
      Passphrase =  "password123"
    })
  }
}
