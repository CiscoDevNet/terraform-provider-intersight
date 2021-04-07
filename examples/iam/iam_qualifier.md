### Resource Creation

```hcl
resource "intersight_iam_qualifier" "iam_qualifier1" {
    user_group {
        name = "iam_user_group1"
        idp {
            domain_name = "cisco.com"
            enable_single_logout = true
            name = "Cisco"
            account {
                object_type = "iam.Account"
                moid = intersight_iam_account.account1.id
            }
            type = "saml"
        }
        permissions [
            {
                moid = var.iam_permission
                object_type = "iam.Permission"
            }
        ]
    }
}
```