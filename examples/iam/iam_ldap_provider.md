### Resource Creation

```hcl
resource "intersight_iam_ldap_provider" "iam_ldap_provider1" {
  port   = 636
  server = "10.1.1.1"
  ldap_policy {
    moid = var.iam_ldap_policy
  }
}

variable "iam_ldap_policy" {
   type = string
   description = "value for iam_ldap_policy"
 }
```