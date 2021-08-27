variable "api_key" {
  type = string
  description = "API Key Id from Intersight"
}

variable "secretkey" {
  type = string
  description = "Secret Key File Path or String"
}

variable "endpoint" {
  type = string
  description = "Endpoint URL"
}

variable "encryption_key" {
  type = string
  description = "Encryption key for IPMIOverLan Policy"
}

variable "auth_password" {
  type = string
  description = "Auth Password for SNMP Policy"
}

variable "privacy_password" {
  type = string
  description = "Privacy Password for SNMP Policy"
}

variable "base_properties_password_1" {
  type = string
  description = "Password field in Base properties in LDAP policy"
}

variable "organization" {
  type = string
  description = "Organization moid"
}
