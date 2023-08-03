resource "intersight_certificatemanagement_policy" "certificate1" {
  description = "sample certificate"
  name        = "certificate1"
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  certificates {
      certificate {
            pem_certificate = var.pem_certificate
            object_type = "certificatemanagement.Imc"
        }
        enabled = true
    }
}

