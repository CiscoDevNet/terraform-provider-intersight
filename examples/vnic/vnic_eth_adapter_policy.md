### Resource Creation

```hcl
resource "intersight_vnic_eth_adapter_policy" "v_eth_adapter1" {
  name                    = "v_eth_adapter1"
  rss_settings            = true
  uplink_failback_timeout = 5
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
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

variable "organization" {
  type        = string
  description = "<value for organization>"
}
```