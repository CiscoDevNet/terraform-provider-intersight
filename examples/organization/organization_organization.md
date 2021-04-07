### Resource Creation

```hcl
resource "intersight_organization_organization" "organization1" {
  name = "organization1"
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  resource_groups = [{
    moid        = intersight_resource_group.resource1.id
    object_type = "resource.Group"
  }]

}
```