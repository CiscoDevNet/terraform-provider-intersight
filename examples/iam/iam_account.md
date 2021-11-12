### Resource Creation

```hcl
resource "intersight_iam_account" "account1" {
  name = "iamaccount-1"
   resource_limits {
     moid        = var.iam_resource_limits
     object_type = "iam.ResourceLimits"
   }
   security_holder {
     moid        = var.iam_security_holder
     object_type = "iam.SecurityHolder"
   }
   session_limits {
     moid        = var.iam_session_limits
     object_type = "iam.SessionLimits"
   }
}

 variable "iam_resource_limits" {
   type = string
   description = "value for iam_resource_limits"
 }

 variable "iam_security_holder" {
   type = string
   description = "value for iam_security_holder"
 }

 variable "iam_session_limits" {
   type = string
   description = "value for iam_session_limits"
 }
```