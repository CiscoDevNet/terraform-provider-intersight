### Resource Creation

```hcl
resource "intersight_certificatemanagement_policy" "certificatemanagement_policy1" {
  name        = "certificatemanagement_policy1"
  description = "certificate management policy"
  certificates = [{
    object_type = "certificatemanagement.Imc"
    certificate = {
      object_type = "x509.Certificate"
      enabled     = true
    }
  }]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```