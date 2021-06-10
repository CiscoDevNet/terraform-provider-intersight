### Resource Creation

```hcl
resource "intersight_iam_ldap_group" "iam_ldap_group1" {
  domain = ".example.com"
  name   = "ldap_group1"
  end_point_role {
    moid = data.intersight_iam_end_point_role.imc_admin.results[0].moid
  }
  ldap_policy {
    moid = intersight_iam_ldap_policy.policy1.moid
  }
}
```