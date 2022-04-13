### Resource Creation

```hcl
resource "intersight_server_config_import" "server_profile_import1" {
  profile_name = "server_profile_import1"
  description  = "server_profile_import"

  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  server {
    object_type = "compute.RackUnit"
    moid        = var.compute_rack_unit
  }
  server_profile {
    object_type = "server.Profile"
    moid        = var.profile
  }
}

variable "organization" {
  type        = string
  description = "value for organization"
}

variable "profile" {
  type        = string
  description = "value for profile"
}

variable "compute_rack_unit" {
  type        = string
  description = "value for compute rack unit"
}
```