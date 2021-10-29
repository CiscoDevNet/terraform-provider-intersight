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

variable "privatekey" {
  type = string
  description = "Private key base64 encoded value"
}

variable "pem_certificate" {
  type = string
  description = "Pem certificate base64 encoded value"
}
