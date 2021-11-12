### Resource Creation

```hcl
resource "intersight_resource_group" "resource_group1" {
  name = "resource_group1"
  per_type_combined_selector = [{
    additional_properties = ""
    class_id              = "resource.PerTypeCombinedSelector"
    combined_selector     = "( Tags/any(tt/Key eq \"Intersight.LicenseTier\" and t/Value eq Essential) )"
    empty_filter          = false
    object_type           = "resource.PerTypeCombinedSelector"
    selector_object_type  = "compute.Blade"
  }]
  qualifier = "Allow-Selectors"
  selectors = [{
    object_type           = "resource.Selector"
    additional_properties = ""
    class_id              = "resource.Selector"
    selector              = "/api/v1/asset/DeviceRegistrations?$filter=Moid in(\"intersight_asset_device_registrations_registeration1.id\")"
  }]
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  organizations {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "account"{
  type = string
  description = "Moid of iam.Account"
}

 variable "organization" {
   type = string
   description = "value for organization"
 }
```