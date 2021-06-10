### Resource Creation

```hcl
resource "intersight_iam_ldap_provider" "iam_ldap_provider1" {
  port   = 636
  server = "10.1.1.1"
  ldap_policy {
    moid = intersight_iam_ldap_policy.policy1.moid
  }
}
```