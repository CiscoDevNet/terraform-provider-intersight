### Resource Creation

```hcl
resource "intersight_iam_idp" "idp1" {
  domain_name          = "cisco.com"
  enable_single_logout = true
  name                 = "Cisco"
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
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
```