resource "intersight_storage_drive_security_policy" "storage_drive_security_policy1" {
  key_setting {
    key_type    = "Kmip"
    object_type = "storage.KeySetting"
    remote_key {
      auth_credentials {
        password           = "password"
        use_authentication = true
        username           = "kmip_user"
      }
      existing_key = "existing key"
      object_type = "storage.RemoteKeySetting"
      primary_server {
        enable_drive_security = true
        ip_address            = "kmip-primary.example.com"
        object_type           = "storage.KmipServer"
        port                  = 5696
        timeout               = 60
      }
      server_certificate = var.pem_certificate
    }
  }
  name        = "test_remote_key"
  object_type = "storage.DriveSecurityPolicy"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}