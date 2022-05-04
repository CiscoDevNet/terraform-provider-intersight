### Resource Creation

```hcl
resource "intersight_iam_private_key_spec" "iam_private_key_spec1" {
  certificate_request {
    object_type = "iam.CertificateRequest"
    moid        = var.certificate_request
  }
}

variable "certificate_request" {
  type        = string
  description = "value for certificate request"
}
```