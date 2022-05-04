### Resource Creation

```hcl
resource "intersight_os_install" "os1" {
  name = "InstallTemplatee165"
  server {
    object_type = "compute.RackUnit"
    moid        = var.server_moid
  }
  image {
    object_type = "softwarerepository.OperatingSystemFile"
    moid        = var.osf1
  }
  osdu_image {
    moid        = var.scu1
    object_type = "firmware.ServerConfigurationUtilityDistributable"
  }
  answers {
    answer_file = var.answer_file
    nr_source   = "File"
    object_type = "os.Answers"
  }
  description    = "Install Template 5"
  install_method = "vMedia"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
  type        = string
  description = "value for organization"
}

variable "answer_file" {
  type        = string
  description = "value for answer file"
}

variable "server_moid" {
  type        = string
  description = "value for moid"
}

variable "osf1" {
  type        = string
  description = "value for osf1"
}

variable "scu1" {
  type        = string
  description = "value for scu1"
}
```