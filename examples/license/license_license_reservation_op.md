### Resource Creation

```hcl
resource "intersight_license_license_reservation_op" "license_license_reservation_op1" {
  generate_request_code = true
  generate_return_code  = true
  account {
    object_type = "iam.Account"
    moid        = var.iam_account
  }
}

variable "iam_account" {
  type        = string
  description = "Moid of iam account"
}
```