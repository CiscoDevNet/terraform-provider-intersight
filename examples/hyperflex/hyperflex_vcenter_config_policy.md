### Resource Creation

```hcl
resource "intersight_hyperflex_vcenter_config_policy" "hyperflex_vcenter_config_policy1" {
  hostname    = "10.10.10.1"
  username    = "admin@vsphere.local"
  password    = "Changeme"
  data_center = "ucsback"
  sso_url     = ""
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  name = "hyperflex_vcenter_config_policy1"
}
```