### Resource Creation
```hcl
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
```