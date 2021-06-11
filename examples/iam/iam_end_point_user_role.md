### Resource Creation

```hcl
resource "intersight_iam_end_point_user_role" "iam_end_point_user_role1" {
  enabled  = true
  password = "ChangeMe"
  end_point_role {
    moid        = data.intersight_iam_end_point_role.admin_role.results[0].moid
    object_type = data.intersight_iam_end_point_role.admin_role.results[0].object_type
  }
  end_point_user {
    moid        = intersight_iam_end_point_user.iam_end_point_user1.moid
    object_type = intersight_iam_end_point_user.iam_end_point_user1.object_type
  }
  end_point_user_policy {
    moid        = intersight_iam_end_point_user_policy.user_policy1.moid
    object_type = intersight_iam_end_point_user_policy.user_policy1.object_type
  }
}
```