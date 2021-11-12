### Resource Creation

```hcl
resource "intersight_vnic_iscsi_boot_policy" "vnic_iscsi_boot_policy" {
  name                           = "vnic_iscsi_boot_policy1"
  description                    = "vnic iscsi boot policy"
  initiator_ip_source            = "DHCP"
  target_source_type             = "Auto"
  initiator_static_ip_v4_address = "10.1.1.1"
  chap {
    password        = "ChangeMeChange"
    user_id         = "user_1"
    object_type     = "vnic.IscsiAuthProfile"
    is_password_set = false

  }
  mutual_chap {
    password        = "ChangeMeChange1"
    user_id         = "user_1"
    is_password_set = false
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```