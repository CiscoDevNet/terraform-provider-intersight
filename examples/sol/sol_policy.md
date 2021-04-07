### Resource Creation

```hcl
resource "intersight_sol_policy" "sol1" {
  name      = "sol2"
  enabled   = false
  baud_rate = 9600
  com_port  = "com1"
  ssh_port  = 1096
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