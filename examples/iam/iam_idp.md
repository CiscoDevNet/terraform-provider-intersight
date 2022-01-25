### Resource Creation

```hcl
resource "intersight_iam_idp" "idp1" {
  domain_name          = "cisco.com"
  enable_single_logout = true
  name                 = "Cisco"
   account {
     object_type = "iam.Account"
     moid        = var.iam_account
   }
   parent {
     moid        = var.iam_system
     object_type = "iam.System"
   }
   system {
     moid        = var.iam_system
     object_type = "iam.System"
   }
  type = "saml"
}


 variable "iam_system" {
   type = string
   description = "value for iam system"
 }

 variable "iam_account" {
   type = string
   description = "value for iam account"
 }
```