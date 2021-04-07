### Resource Creation

```hcl
resource "intersight_vnic_iscsi_adapter_policy" "vnic_iscsi_adapter_policy" {
    name        = "vnic_iscsi_adapter_policy1"
    description = "vnic iscsi adapter policy"
    rate_limit  = 5
    burst       = 4000
    max_data_field_size = 512
    cos     = 5
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```