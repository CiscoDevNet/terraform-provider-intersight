### Resource Creation

```hcl
resource "intersight_iam_certificate_request" "iam_certificate_request1" {
  email_address = "email@example.com"
  name          = "iam_certificate_request1"
  self_signed   = false
  account {
    object_type = "iam.Account"
    moid        = var.intersight_iam_account
  }
  certificate {
    moid        = var.iam_certificate1
    object_type = "x509.Certificate"
  }
}

variable "intersight_iam_account" {
   type = string
   description = "MOID value for intersight_iam_account"
 }

 variable "iam_certificate1" {
   type = string
   description = "value for iam_certificate1"
}
```