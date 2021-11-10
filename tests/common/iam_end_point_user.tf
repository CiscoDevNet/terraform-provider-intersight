
resource "intersight_iam_end_point_user" "iam_end_point_user1" {
  name = "iam_ept_user1"

  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}

