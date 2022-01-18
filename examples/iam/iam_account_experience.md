### Resource Creation

```hcl
resource "intersight_iam_account_experience" "iam_account_experience1" {
  account {
    object_type = "iam.Account"
    moid        = var.iam_account
  }
  features = [
    {
      additional_properties = ""
      class_id              = "iam.FeatureDefinition"
      feature               = "UnrecognizedValue"
      object_type           = "iam.FeatureDefinition"
    }
  ]
   parent = {
     moid        = var.iam_account
     object_type = "iam.Account"
   }
}

 variable "iam_account" {
   type = string
   description = "value for iam_account"
}
```