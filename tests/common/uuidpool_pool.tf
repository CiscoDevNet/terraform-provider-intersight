#resource "intersight_uuidpool_pool" "uuidpool_pool1" {
#  name             = "uuidpool_pool1"
#  description      = "uuidpool_pool"
#  assignment_order = "default"
#  size             = 774325
#  prefix           = "123e4567-e89b-42d3"
#  uuid_suffix_blocks = [{
#    additional_properties = ""
#    class_id              = "uuidpool_UuidBlock"
#    object_type           = "uuidpool.UuidBlock"
#    from                  = "BC2e-fBED62ea0FfD" #"123e4567-e89b-42d3"
#    to                    = "E6DB-CfB3782a41Ec" #"123e4567-e89b-84e6"
#    size                  = 4
#  }]
#  organization {
#    object_type = "organization.Organization"
#    moid        = data.intersight_organization_organization.default.results.0.moid
#  }
#}
