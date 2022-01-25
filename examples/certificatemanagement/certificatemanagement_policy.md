### Resource Creation

```hcl
resource "intersight_certificatemanagement_policy" "certificate1" {
  description = "sample certificate"
  name        = "certificate1"
  organization {
    moid        = var.organization
    object_type = "organization.Organization"
  }
  certificates {
    certificate {
      pem_certificate = var.pem_certificate1
    }
    enabled    = true
    privatekey = var.privatekey
  }
}

variable "privatekey" {
  type        = string
  sensitive   = true
  description = "Private key base64 encoded value"
}
variable "pem_certificate1" {
  type        = string
  sensitive   = true
  description = "Pem certificate base64 encoded value"
}
```
