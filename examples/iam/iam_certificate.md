### Resource Creation

```hcl
resource "intersight_iam_certificate" "iam_certificate1" {
  certificate {
    pem_certificate = var.pemcertificate1
    object_type     = "x509.Certificate"
  }
}


variable "pemcertificate1" {
  type        = string
  sensitive   = true
  description = "Pem certificate base64 encoded value"
}
```