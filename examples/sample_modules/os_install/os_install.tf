resource "intersight_os_install" "os1" {
  name = "InstallTemplatee165"
  server {
    object_type = "compute.RackUnit"
    moid = var.server_moid
  }
  image {
    object_type = "softwarerepository.OperatingSystemFile"
    moid = intersight_softwarerepository_operating_system_file.osf1.moid
  }
  osdu_image {
    moid = intersight_firmware_server_configuration_utility_distributable.scu1.moid
    object_type = "firmware.ServerConfigurationUtilityDistributable"
  }
  answers {
    answer_file = var.answer_file
    nr_source = "File"
  }
  description = "Install Template 5"
  install_method = "vMedia"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

/*
SAMPLE PAYLOAD
-----------------
{
  "Answers": {
    "AnswerFile": "",
    "Source": "File"
  },
  "Description": "Install TRemplate 5",
  "InstallMethod": "vMedia",
  "Name": "InstallTemplatee165",
  "Server": {
    "ObjectType": "compute.RackUnit",
    "Moid": ""
  },
  "Image": {
    "ObjectType": "softwarerepository.OperatingSystemFile",
    "Moid": ""
  },
  "OsduImage": {
    "Moid": "",
    "ObjectType": "firmware.ServerConfigurationUtilityDistributable"
  }
}
*/