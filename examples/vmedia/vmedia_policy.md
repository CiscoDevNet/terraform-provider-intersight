### Resource Creation
```hcl
resource "intersight_vmedia_policy" "vmedia1" {
  name          = "vmedia1"
  description   = "test policy"
  enabled       = true
  encryption    = true
  low_power_usb = true

  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```