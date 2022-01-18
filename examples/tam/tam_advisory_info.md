### Resource Creation

```hcl
resource "intersight_tam_advisory_info" "tam_advisory_info" {
  state = "active"
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
}

variable "account"{
  type = string
  description = "Moid of iam.Account Mo"
}
```