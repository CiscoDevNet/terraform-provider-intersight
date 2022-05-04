### Resource Creation

```hcl
resource "intersight_vnic_fc_network_policy" "v_fc_network1" {
  name = "v_fc_network1"
  vsan_settings {
    id          = 100
    object_type = "vnic.VsanSettings"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
  type        = string
  description = "<value for organization>"
}
```