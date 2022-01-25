### Resource Creation

```hcl
resource "intersight_iam_trust_point" "iam_trust_point1" {
   account {
     object_type = "iam.Account"
     moid        = var.iam_account
   }
   certificates = [{
     moid        = var.certificate_iam
     object_type = "x509.Certificate"
     enabled     = true
   }]
}

 variable "iam_account" {
   type = string
   description = "value for iam_account"
 }

  variable "certificate_iam" {
   type = string
   description = "value for certificate_iam"
 }
```