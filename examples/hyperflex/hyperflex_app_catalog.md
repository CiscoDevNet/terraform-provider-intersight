### Resource Creation

```hcl
resource "intersight_hyperflex_app_catalog" "hyperflex_app_catalog1" {
  feature_limit_internal {
    moid        = var.hyperflex_feature_limit_internal
    object_type = "hyperflex.FeatureLimitInternal"
  }
  hxdp_versions = [
    {
	  additional_properties = ""
	  selector = ""
	  class_id = "hyperflex.HxdpVersion"
      object_type = "hyperflex.HxdpVersion"
	  moid = var.hyperflex_hxdp_version
    }
  ]
  hyperflex_capability_infos = [
    {
	  additional_properties = ""
	  selector = ""
	  class_id = "hyperflex.CapabilityInfo"
      object_type = "hyperflex.CapabilityInfo"
	  moid = var.hyperflex_capability_info
    }
  ]
  hyperflex_software_compatibility_infos = [
    {
	  additional_properties = ""
	  selector = ""
      object_type = "hcl.HyperflexSoftwareCompatibilityInfo"
	  class_id = "hcl.HyperflexSoftwareCompatibilityInfo"
	  moid = var.hcl_hyperflex_software_compatibility_info
    }
  ]
  server_firmware_version {
    moid        = var.hyperflex_server_firmware_version
    object_type = "hyperflex.ServerFirmwareVersion"
  }
  server_model {
    moid        = var.hyperflex_server_model
    object_type = "hyperflex.ServerModel"
  }
  software_distributions = [
    {
      object_type = "hyperflex.SoftwareDistributionEntry"
	  class_id = "hyperflex.SoftwareDistributionEntry"
	  additional_properties = ""
	  selector = ""
	  moid = var.hyperflex_software_distribution_entry
    }
  ]
}

variable "hyperflex_feature_limit_internal" {
  type = string
  description = "MOID of hyperflex.FeatureLimitInternal"
}

variable "hyperflex_hxdp_version" {
  type = string
  description = "MOID of hyperflex.HxdpVersion"
}

variable "hyperflex_capability_info" {
  type = string
  description = "MOID of hyperflex.CapabilityInfo"
}

variable "hcl_hyperflex_software_compatibility_info" {
  type = string
  description = "MOID of hcl.HyperflexSoftwareCompatibilityInfo"
}

variable "hyperflex_server_firmware_version" {
  type = string
  description = "MOID of hyperflex.ServerFirmwareVersion"
}

variable "hyperflex_server_model" {
  type = string
  description = "MOID of hyperflex.ServerModel"
}

variable "hyperflex_software_distribution_entry" {
  type = string
  description = "MOID of hyperflex.SoftwareDistributionEntry"
}
```
