### Resource Creation

```hcl
resource "intersight_iam_certificate_request" "iam_certificate_request1" {
  email_address = "email@example.com"
  name          = "iam_certificate_request1"
  self_signed   = false
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  certificate {
    moid        = intersight_certificate_iam.iam1.id
    object_type = "x509.Certificate"
    enabled     = true
  }
}
```