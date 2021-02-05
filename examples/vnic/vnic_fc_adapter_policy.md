### Resource Creation

```hcl
resource "intersight_vnic_fc_adapter_policy" "v_fc_adapter1" {
  name = "v_fc_adapter1"
  error_detection_timeout = 100000
  organization {
    object_type = "organization.Organization"
    moid = var.organization
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
```