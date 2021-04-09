### Resource Creation

```hcl
resource "intersight_iam_certificate" "iam_certificate1" {
  certificate {
    moid        = intersight_certificate_iam.iam1.id
    object_type = "x509.Certificate"
    enabled     = true
  }
}
```