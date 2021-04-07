### Resource Creation

```hcl
resource "intersight_config_importer" "config_importer1" {
    name = "config_exporter1"
    items [{
        object_type = "config.MoRef"
        moid = var.config_mo_ref
    }]
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
}
```