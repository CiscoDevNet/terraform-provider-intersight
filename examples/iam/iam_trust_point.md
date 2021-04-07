### Resource Creation

```hcl
resource "intersight_iam_trust_point" "iam_trust_point1" {
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  certificates = [{
    moid        = intersight_certificate_iam.iam1.id
    object_type = "x509.Certificate"
    enabled     = true
  }]
}
```