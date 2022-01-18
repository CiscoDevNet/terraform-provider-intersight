### Resource Creation

```hcl
resource "intersight_organization_organization" "organization1" {
  name = "organization1"
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  resource_groups = [{
    moid                  = var.resource_group
    object_type           = "resource.Group"
    additional_properties = ""
    class_id              = "resource.Group"
    selector              = null
  }]

}

variable "account"{
  type = string
  description = "Moid of iam.Account"
}

variable "resource_group"{
  type = string
  description = "Moid of resource.Group"
}
```