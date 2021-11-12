### Resource Creation

```hcl
resource "intersight_capability_sioc_module_descriptor" "capability_sioc_module_descriptor1" {
  capabilities = [
    {

      object_type           = "capability.SiocModuleManufacturingDef"
      moid                  = var.capability_sioc_module_manufacturing_def
      class_id              = "capability.SiocModuleManufacturingDef"
      additional_properties = ""
      selector              = ""
    },
    {
      object_type           = "capability.SiocModuleCapabilityDef"
      moid                  = var.capability_sioc_module_capability_def
      class_id              = "capability.SiocModuleCapabilityDef"
      additional_properties = ""
      selector              = ""
    }
  ]
  description = "capability of sioc module descriptor"
  model       = "UCS-S3260-PCISIOC"
  revision    = "0"
  vendor      = "Cisco Systems Inc"

}

variable "capability_sioc_module_capability_def" {
  type        = string
  description = "moid of capability.SiocModuleCapabilityDef Mo"
}

variable "capability_sioc_module_manufacturing_def" {
  type        = string
  description = "moid of capability.SiocModuleManufacturingDef Mo"
}
```