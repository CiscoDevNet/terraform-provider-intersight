### Resource Creation

```hcl
resource "intersight_iam_end_point_user_role" "iam_end_point_user_role1" {
  enabled  = true
  password = "ChangeMe"
  end_point_role {
    
    moid        =  var.iam_end_point_role
    object_type = "iam.EndPointRole"
  }
  end_point_user {
    
    moid        = var.iam_end_point_user1
    object_type = "iam.EndPointUser"
  }
  end_point_user_policy {
    
    moid        = var.user_policy1
    object_type = "iam.EndPointUserPolicy"
  }
}

 variable "iam_end_point_user1" {
   type = string
   description = "value for iam_end_point_user1"
}

 variable "user_policy1" {
   type = string
   description = "value for user_policy1"
}

 variable "iam_end_point_role" {
   type = string
   description = "value for iam_end_point_role"
}
```