### Resource Creation

```hcl
resource "intersight_iam_resource_roles" "iam_resource_roles1" {
  parent {
    moid        = var.iam_permission
    object_type = "iam.Permission"
  }
  permission {
    moid        = var.iam_permission
    object_type = "iam.Permission"
  }
  resource {
    moid        = var.organization_organization
    object_type = "organization.Organization"
  }
  roles = [
    {
      moid        = var.iam_role
      object_type = "iam.Role"
    }
  ]
}
```