### Resource Creation

```hcl
resource "intersight_techsupportmanagement_collection_control_policy" "techsupportmanagement_collection_control_policy" {
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
  tech_support_collection = "Enable"

}

variable "account"{
  type = string
  description = "Moid of iam Account"
}
```