### Resource Creation

```hcl
resource "intersight_hyperflex_software_version_policy" "hyperflex_software_version_policy1" {
    hxdp_version = "4.5(1a)"
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
    name = "hyperflex_software_version_policy1"
}
```