### Resource Creation

```hcl
resource "intersight_iam_ldap_provider" "iam_ldap_provider1" {
    port = 636
    server = "10.1.1.1"
    iam_ldap_policy {
        moid = var.iam_ldap_policy
        object_type = "iam.LdapPolicy"
        enabled                = true
        enable_dns             = true
        user_search_precedence = "LocalUserDb"
        organization {
            object_type = "organization.Organization"
            moid = var.organization
        }
    }
}
```