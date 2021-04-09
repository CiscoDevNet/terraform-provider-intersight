### Resource Creation

```hcl
resource "intersight_ssh_policy" "ssh_policy1" {
  name        = "ssh_policy1"
  description = "ssh policy"
  enabled     = true
  port        = 22
  timeout     = 1800
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
```