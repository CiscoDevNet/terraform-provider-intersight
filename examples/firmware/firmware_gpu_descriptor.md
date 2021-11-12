### Resource Creation

```hcl
resource "intersight_firmware_gpu_descriptor" "firmware_gpu_descriptor1" {
  description = "firmware gpu descriptor"
  model       = "UCSC-GPU-V340"
  vendor      = "Cisco"
  revision    = "0"
}
```