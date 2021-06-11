### Resource Creation

```hcl
resource "intersight_iam_end_point_user" "iam_end_point_user1" {
  name = "iam_end_point_user1"
  # mapping of user to role is performed by 
  # resource intersight_iam_end_point_user_role 

  organization {
    moid        = var.organization_organization
    object_type = "organization.Organization"
  }
}
```