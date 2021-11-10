
resource "intersight_fabric_port_policy" "fabric_port_policy1" {
  name         = "fabric_port_policy1"
  description  = "demo fabric port policy"
  device_model = "UCS-FI-6454"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
