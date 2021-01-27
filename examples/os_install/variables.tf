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

variable "scu_password_1" {
  type = string
  description = "Password for source in SCU"
}

variable "scu_password_2" {
  type = string
  description = "Password for source in SCU"
}

variable "os_file_password_1" {
  type = string
  description = "Password for esx6u3.iso remote file in software repo OS"
}

variable "organization" {
  type = string
  description = "Organization moid"
}

variable "server_moid" {
  type = string
  description = "Rack unit's moid"
}

variable "catalog" {
  type = string
  description = "Catalog moid"
}

variable "answer_file" {
  type = string
  default = ""
  description = "Answer file content"
}