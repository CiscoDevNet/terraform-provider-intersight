### Resource Creation

```hcl
resource "intersight_resource_group" "resource_group1" {
    name = "resource_group1"
    per_type_combined_selector = [{
        combined_selector = "( Tags/any(tt/Key eq 'Intersight.LicenseTier' and t/Value eq Essential) )"
        empty_filter = false
        object_type = "resource.PerTypeCombinedSelector"
        selector_object_type = "compute.Blade"
    }]
    qualifier = "Allow-Selectors"
    selectors = [{
        object_type = "resource.Selector"
        selector = "/api/v1/asset/DeviceRegistrations?$filter=Moid in("
        intersight_asset_device_registrations_registeration1.id ")"
    }]
    account {
        object_type = "iam.Account"
        moid = intersight_iam_account.account1.id
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```