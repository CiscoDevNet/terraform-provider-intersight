### Resource Creation

```hcl
resource "intersight_iam_trust_point" "iam_trust_point1" {
  account {
    object_type = "iam.Account"
    moid        = var.iam_account
  }
  certificates {
    object_type = "x509.Certificate"
	pem_certificate = ""
  }
}

variable "iam_account" {
  type        = string
  description = "value for iam_account"
}
```