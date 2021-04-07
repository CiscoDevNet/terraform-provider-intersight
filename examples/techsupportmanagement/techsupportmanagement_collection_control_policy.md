### Resource Creation

```hcl
resource "intersight_techsupportmanagement_collection_control_policy" {
    account {
        object_type = "iam.Account"
        moid = intersight_iam_account.account1.id
    }
    tech_support_collection = "Enable"

}
```