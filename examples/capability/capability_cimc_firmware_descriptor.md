### Resource Creation

```hcl
resource "intersight_capability_cimc_firmware_descriptor" "capability_cimc_firmware_descriptor1" {
  capabilities {
    moid        = var.capability_cimc_firmware_descriptor
    object_type = "capability.CimcFirmwareDescriptor"
    class_id    = "capability.CimcFirmwareDescriptor"
  }
  description = "capability of cimc firmware descriptor"
  model       = "UCSC-C240-M5L"
  revision    = "0"
  vendor      = "Cisco Systems Inc"

}

variable "capability_cimc_firmware_descriptor" {
  type        = string
  description = "Moid of capability.CimcFirmwareDescriptor Mo"
}
```