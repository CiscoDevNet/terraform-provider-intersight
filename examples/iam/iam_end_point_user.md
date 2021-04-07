### Resource Creation

```hcl
resource "intersight_iam_end_point_user" "iam_end_point_user1" {
    name = "iam_end_point_user1"
    end_point_user_role {
        moid = var.iam_end_point_user_role
        object_type = "iam.endPointUserRole"
        enabled = true
        password = "ChangeMe"
        
    }
    organization {
        moid = var.organization_organization
        object_type = "organization.Organization"
    }
}
```