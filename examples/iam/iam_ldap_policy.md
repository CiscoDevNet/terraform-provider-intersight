### Resource Creation

```hcl
resource "intersight_iam_ldap_policy" "ldap1" {
  name                   = "ldap1"
  description            = "test policy"
  enabled                = true
  enable_dns             = true
  user_search_precedence = "LocalUserDb"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  base_properties {
    attribute                  = "CiscoAvPair"
    base_dn                    = "DC=QATCSLABTPI02DC=ciscoDC=com"
    bind_dn                    = "CN=administratorCN=UsersDC=QATCSLABTPI02DC=ciscoDC=com"
    bind_method                = "Anonymous"
    domain                     = "QATCSLABTPI02.cisco.com"
    enable_encryption          = true
    enable_group_authorization = true
    filter                     = "sAMAccountName"
    group_attribute            = "memberOf"
    nested_group_search_depth  = 128
    timeout                    = 180
    object_type                = "iam.LdapBaseProperties"
  }
  dns_parameters {
    nr_source     = "Extracted"
    search_forest = "xyz"
    search_domain = "abc"
    object_type   = "iam.LdapDnsParameters"
  }
}

variable "organization" {
  type        = string
  description = "value for organization organization"
}
```