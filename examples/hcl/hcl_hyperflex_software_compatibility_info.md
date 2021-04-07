### Resource Creation

```hcl
resource "intersight_hcl_hyperflex_software_compatibility_info" "hcl_hyperflex_software_compatibility_info1" {
    app_catalog {
      moid = var.hyperflex_catalog
      ObjectType = "hyperflex.AppCatalog"
    }
    constraints [
      {
        constraint_name = "supportedOperations"
        constraint_value = "UPGRADE"
        object_type = "hcl.Constraint"
      }
    ]
    hxdp_version = "4.5(1a)-39020"
    hypervisor_type = "ESXi"
    hypervisor_version = "HX-ESXi-7.0U1-17325551-Cisco-Custom-7.1.0.4"
    parent {
        moid = var.hyperflex_catalog
        object_type = "hyperflex.AppCatalog"
    }
    server_fw_version = "4.0(1a)"
}
```