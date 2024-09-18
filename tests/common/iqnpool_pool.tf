resource "intersight_iqnpool_pool" "iqnpool_pool1" {
  name = "ippool_pool1"
  description = "demo_ippool pool"
  assignment_order = "sequential"
  prefix = "iqn.2023-06.abc.com"
  
  iqn_suffix_blocks {
    object_type = "iqnpool.IqnSuffixBlock"
    suffix = "iscsi01"
    from = 0
    to = 20
  }
  organization {
    object_type = "organization.Organization"
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}