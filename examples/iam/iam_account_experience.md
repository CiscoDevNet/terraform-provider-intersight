### Resource Creation

```hcl
resource "intersight_iam_account_experience" "iam_account_experience1" {
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  features = [
    {
      feature     = "UnrecognizedValue"
      object_type = "iam.FeatureDefinition"
    }
  ]
  parent = {
    moid        = var.iam_account
    object_type = "iam.Account"
  }
}
```