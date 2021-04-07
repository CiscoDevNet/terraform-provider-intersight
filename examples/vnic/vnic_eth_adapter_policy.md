### Resource Creation

```hcl
resource "intersight_vnic_eth_adapter_policy" "vnic_eth_adapter_policy" {
    name        = "vnic_eth_adapter_policy1"
    description = "vnic ethernet adapter policy"
    interrupt_scaling   = true
    uplink_failback_timeout = 5
    interrupt_settings  = {
        object_type = "vnic.EthInterruptSettings"
        coalescing_time     = 5
        coalescing_type     = "MIN"
        count       = 10
        mode        = "MSI"
    }
    nvgre_settings  = {
        object_type = "vnic.NvgreSettings"
        enabled     = false
    }
    vxlan_settings  = {
        object_type = "vnic.VxlanSettings"
        enabled     = false
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```