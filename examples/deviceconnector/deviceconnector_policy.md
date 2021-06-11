### Resource Creation

```hcl
resource "intersight_deviceconnector_policy" "dcp1" {
  name            = "device_con1"
  description     = "demo device connector policy"
  lockout_enabled = true
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```