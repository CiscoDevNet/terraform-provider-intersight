data "intersight_hyperflex_hxdp_version" "hxdp_version"{	
}


resource "intersight_hyperflex_software_version_policy" "hyperflex_software_version_policy1" {
  hxdp_version = data.intersight_hyperflex_hxdp_version.hxdp_version.results[length(data.intersight_hyperflex_hxdp_version.hxdp_version.results)-1].nr_version
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  name = "hyperflex_software_version_policy1"
}
