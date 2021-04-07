### Resource Creation

```hcl
resource "intersight_vnic_fc_qos_policy" "vnic_fc_qos_policy" {
    name        = "vnic_fc_qos_policy1"
    description = "vnic fabric channel qos policy"
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