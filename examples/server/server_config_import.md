### Resource Creation

```hcl
resource "intersight_server_config_import" "server_profile_import1" {
  profile_name = "server_profile_import1"
  description  = "server_profile_import"
  policy_prefix = [{
    object_type = "policies"
    moid        = var.policies
  }]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  server {
    object_type = "compute.RackUnit"
    moid        = var.compute.RackUnit
  }
  server_profile {
    object_type = "server.Profile"
    moid        = intersight_server_profile.server1.id
  }
}
```