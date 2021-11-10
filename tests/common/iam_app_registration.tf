
resource "intersight_iam_app_registration" "iam_app_registration1" {
  client_name         = "name1"
  client_type         = "confidential"
  revoke              = true
  renew_client_secret = true
  description         = "value for description"
}


