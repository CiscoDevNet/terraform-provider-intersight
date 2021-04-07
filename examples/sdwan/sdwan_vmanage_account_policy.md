### Resource Creation

```hcl
resource "intersight_sdwan_vmanage_account_policy" "sdwan_vmanage_account_policy1" {
    name = "sdwan_vmanage_account_policy1"
    description = "sdwan_vmanage_account_policy"
    endpoint_address = "<server fqdn>"
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
    profiles {
        object_type = "sdwan.Profile"
        moid = intersight_sdwan_profile.sdwan_profile1.id
    }
    user_name = "user_name to authenticate the vmanage_server"

}
```