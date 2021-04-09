### Resource Creation

```hcl
resource "intersight_ipmioverlan_policy" "ipmi1" {
  name           = "ipmi1"
  description    = "demo ipmi policy"
  enabled        = true
  privilege      = "admin"
  encryption_key = var.encryption_key
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