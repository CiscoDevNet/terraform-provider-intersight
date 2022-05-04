### Resource Creation

```hcl
resource "intersight_resource_group" "resource_group1" {
  name = "resource_group1"
  per_type_combined_selector {
    class_id    = "resource.PerTypeCombinedSelector"
    object_type = "resource.PerTypeCombinedSelector"
  }
  qualifier = "Allow-Selectors"
  selectors {
    object_type = "resource.Selector"
    class_id    = "resource.Selector"
    selector    = "/api/v1/asset/DeviceRegistrations?$filter=Moid in(\"intersight_asset_device_registrations_registeration1.id\")"
  }
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  organizations {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "account" {
  type        = string
  description = "Moid of iam.Account"
}

variable "organization" {
  type        = string
  description = "value for organization"
}
```