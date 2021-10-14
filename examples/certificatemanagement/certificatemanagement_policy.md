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
            pem_certificate = var.pem_certificate
        }
        enabled = true
        privatekey = var.privatekey
    }
}
```