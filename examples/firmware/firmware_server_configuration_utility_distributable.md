### Resource Creation

```hcl
resource "intersight_firmware_server_configuration_utility_distributable" "scu1" {
  name = "SCU-6.0.4c nfs"
  nr_source {
    object_type = "softwarerepository.CifsServer"
    additional_properties = jsonencode({
      FileLocation = "/path/to/filelocation"
      RemoteIp     = "10.10.10.1"
      RemoteShare  = "/path/to/remote/share"
      RemoteFile   = "/path/to/remote/file"
      Username     = "user"
      Password     = "ChangeMe"
      MountOption  = "sec=ntlm"
    })
  }
  vendor     = "Cisco"
  nr_version = "6.0.(4c)"
  supported_models = [
    "C-series"
  ]
  description = "Cisco SCU-6.0(4c)"
  catalog {
    moid        = var.catalog
    object_type = "softwarerepository.Catalog"
  }
}

variable "catalog" {
  type        = string
  description = "Moid of catalog for firmware_server_configuration_utility_distributable"
}
```