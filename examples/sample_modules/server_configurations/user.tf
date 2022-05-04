resource "intersight_iam_end_point_user_policy" "tf_user_policy" {
  name        = "tf_user_policy"
  description = "test policy"

  password_properties {
    enforce_strong_password  = true
    enable_password_expiry   = true
    password_expiry_duration = 50
    password_history         = 5
    notification_period      = 1
    grace_period             = 2
    object_type              = "iam.EndPointPasswordProperties"
  }
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}