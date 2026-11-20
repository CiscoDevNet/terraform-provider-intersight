resource "intersight_access_policy" "access1" {
  name        = "access1"
  description = "demo imc access policy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  configuration_type{
	configure_inband = false
	configure_out_of_band = true
  }
  out_of_band_ip_pool {
	object_type = "ippool.Pool"
	moid = intersight_ippool_pool.ippool_pool_tf.moid
  }
}
