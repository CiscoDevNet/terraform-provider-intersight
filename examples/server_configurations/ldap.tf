resource "intersight_iam_ldap_policy" "ldap1" {
  name                   = "ldap1"
  description            = "test policy"
  enabled                = true
  enable_dns             = true
  user_search_precedence = "LocalUserDb"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  base_properties {
    attribute                  = "CiscoAvPair"
    base_dn                    = "DC=QATCSLABTPI02,DC=cisco,DC=com"
    bind_dn                    = "CN=administrator,CN=Users,DC=QATCSLABTPI02,DC=cisco,DC=com"
    bind_method                = "Anonymous"
    domain                     = "QATCSLABTPI02.cisco.com"
    enable_encryption          = true
    enable_group_authorization = true
    filter                     = "sAMAccountName"
    group_attribute            = "memberOf"
    nested_group_search_depth  = 128
    timeout                    = 180
  }
  dns_parameters {
    nr_source        = "Extracted"
    search_forest = "xyz"
    search_domain = "abc"
  }
}

resource "intersight_iam_ldap_policy" "ldap2" {
  name                   = "ldap2"
  description            = "test policy"
  enabled                = true
  enable_dns             = true
  user_search_precedence = "LocalUserDb"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  base_properties {
    attribute                  = "CiscoAvPair"
    base_dn                    = "DC=QATCSLABTPI02,DC=cisco,DC=com"
    bind_dn                    = "CN=administrator,CN=Users,DC=QATCSLABTPI02,DC=cisco,DC=com"
    bind_method                = "Anonymous"
    domain                     = "QATCSLABTPI02.cisco.com"
    enable_encryption          = true
    enable_group_authorization = true
    filter                     = "sAMAccountName"
    group_attribute            = "memberOf"
    password                   = var.base_properties_password_1
    nested_group_search_depth  = 128
    timeout                    = 180
  }
  dns_parameters {
    nr_source        = "Extracted"
    search_forest = "xyz"
    search_domain = "abc"
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

/*
SAMPLE PAYLOAD
-----------------
IamLdapPolicyApi: {
    "Name": "AUTO_LDAP_POLICY_CRR",
    "Description": "Test policy",
    "Tags": [{"Key": "ldap", "Value": "Ldap_policy"}],
    "Enabled": True,
    "EnableDns": False,
    "BaseProperties": {
        "BaseDn": "DC=new,DC=com",
        "Domain": "new.com",
        "EnableEncryption": True,
        "BindDn": "admin",
        "Timeout": 180,
        "BindMethod": "Anonymous",
        "Filter": "sAMAccountName",
        "Attribute": "CiscoAvPair",
        "GroupAttribute": "memberOf",
        "NestedGroupSearchDepth": 128,
        "EnableGroupAuthorization": True
    },
    "DnsParameters": {
        "Source": "Extracted",
        "SearchForest": "",
        "SearchDomain": ""
    },
    "UserSearchPrecedence": "LocalUserDb"
}
*/
