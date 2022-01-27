#
#resource "intersight_sdwan_router_node" "sdwan_router_node1" {
#  name = "sdwan_router_node1"
#  network_configuration = [{
#    object_type           = "sdwan.NetworkConfigurationType"
#    network_type          = "LAN"
#    port_group            = 19
#    vlan                  = 1000
#    additional_properties = ""
#    class_id              = "sdwan.NetworkConfigurationType"
#  }]
#  uuid = "123e4567-e89b-42d3-a456-556642440000"
#  organization {
#    object_type = "organization.Organization"
#    moid        = data.intersight_organization_organization.default.results.0.moid
#  }
#  profile {
#    object_type = "sdwan.Profile"
#    moid        = intersight_sdwan_profile.sdwan_profile1.id
#  }
#}
