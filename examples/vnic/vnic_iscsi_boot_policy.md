### Resource Creation

```hcl
resource "intersight_vnic_iscsi_boot_policy" "vnic_iscsi_boot_policy" {
    name        = "vnic_iscsi_boot_policy1"
    description = "vnic iscsi boot policy"
    dhcp_timeout    = 300
    connection_time_out = 300
    lun_busy_retry_count = 300
    initiator_ip_source = "DHCP"
    target_source_type  = "Auto"
    initiator_static_ip_v4_address = "10.1.1.1"
    chap    = {
        object_type = "vnic.IscsiAuthProfile"
        password    = "ChangeMe"
        user_id     = "user_1"
    }
    mutual_chap = {
        object_type = "vnic.IscsiAuthProfile"
        password    = "ChangeMe"
        user_id     = "user_1"
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```