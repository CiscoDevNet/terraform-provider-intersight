### Resource Creation

```hcl
resource "intersight_tam_advisory_count" "tam_advisory_count" {
  advisory_count = 2
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```
