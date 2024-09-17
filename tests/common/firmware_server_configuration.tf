resource "intersight_firmware_server_configuration_utility_distributable" "scu1" {
  name = "SCU-6.0.4c nfs"
  nr_source {
    object_type = "softwarerepository.CifsServer"
    additional_properties = jsonencode({
      FileLocation = "10.10.10.20/iso.iso"
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
    moid        = data.intersight_softwarerepository_catalog.catalog.results[0].moid
    object_type = "softwarerepository.Catalog"
  }
}

data "intersight_softwarerepository_catalog" "catalog"{
  organization {
    class_id = "organization.Organization"
    moid = data.intersight_organization_organization.default.results.0.moid
  }
} 