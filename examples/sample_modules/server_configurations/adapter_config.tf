resource "intersight_adapter_config_policy" "tf_adapter_config" {
  name        = "tf_adapter_config"
  description = "test policy"
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  settings {
    object_type = "adapter.AdapterConfig"
    slot_id     = "1"
    eth_settings {
      lldp_enabled = true
      object_type  = "adapter.EthSettings"
    }
    fc_settings {
      object_type = "adapter.FcSettings"
      fip_enabled = true
    }
  }
  settings {
    object_type = "adapter.AdapterConfig"
    slot_id     = "MLOM"
    eth_settings {
      object_type  = "adapter.EthSettings"
      lldp_enabled = true
    }
    fc_settings {
      object_type = "adapter.FcSettings"
      fip_enabled = true
    }
  }
  tags {
    key   = "source"
    value = "terraform"
  }
}

/*
SAMPLE PAYLOAD
-----------------
AdapterConfigPolicyApi: {
  "Name": "AUTO_TEST_ACP1",
  "Settings": [
    {
      "SlotId": "1",
      "EthSettings": {"LldpEnabled": True},
      "FcSettings": {"FipEnabled": True}
    }
  ]
}
*/