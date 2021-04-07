### Resource Creation

```hcl
resource "intersight_capability_sioc_module_descriptor" "capability_sioc_module_descriptor1" {
    capabilities [
      {
        
        object_type = "capability.SiocModuleManufacturingDef"
        moid = var.capability_sioc_module_manufacturing_def
      }
      {
        object_type = "capability.SiocModuleCapabilityDef"
        moid = var.capability_sioc_module_capability_def
      }
    ]
    description = "capability of sioc module descriptor"
    model = "UCS-S3260-PCISIOC"
    revision = "0"
    vendor = "Cisco Systems Inc"
    version = "ge 3.1(1.1)"
}
```