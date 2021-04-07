### Resource Creation

```hcl
resource "intersight_hyperflex_app_catalog" "hyperflex_app_catalog1" {
    feature_limit_internal {
      moid = var.hyperflex_feature_limit_internal
      object_type = "hyperflex.FeatureLimitInternal"
    }
    hxdp_versions [
      {
        moid = var.hyperflex_hxdp_version
        object_type = "hyperflex.HxdpVersion"
      }
    ]
    hyperflex_capability_infos [
      {
        moid = var.hyperflex_capability_info
        object_type = "hyperflex.CapabilityInfo"
      }
    ]
    hyperflex_software_compatibility_infos [
      {
        moid = var.hcl_hyperflex_software_compatibility_info
        object_type = "hcl.HyperflexSoftwareCompatibilityInfo"
      }
    ]
    server_firmware_version {
      moid = var.hyperflex_server_firmware_version
      object_type = "hyperflex.ServerFirmwareVersion"
    }
    server_model {
      moid = var.hyperflex_server_model
      object_type = "hyperflex.ServerModel"
    }
    software_distributions [
      {
        moid = var.hyperflex_software_distribution_entry
        object_type = "hyperflex.SoftwareDistributionEntry"
      }
    ]
    version = "1.0.248"
}
```