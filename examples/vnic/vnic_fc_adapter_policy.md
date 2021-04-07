### Resource Creation

```hcl
resource "intersight_vnic_fc_adapter_policy" "vnic_fc_adapter_policy" {
    name        = "vnic_fc_adapter_policy1"
    description = "vnic fabric channel adapter policy"
    io_throttle_count   = 256
    lun_count   = 256
    lun_queue_depth = 256
    plogi_settings  = {
        object_type = "vnic.PlogiSettings"
        retries     = 8
        plogi_timeout   = 20000
        timeout     = 4000
    }
    interrupt_settings  = {
        object_type = "vnic.EthInterruptSettings"
        mode        = "MSI"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```