
---
layout: "intersight"
page_title: "Intersight: intersight_os_configuration_file"
sidebar_current: "docs-intersight-data-source-os-configuration-file"
description: |-
A ConfigurationFile is an OS specific answer file that helps with the unattended
installation.
The file can also be a template file with placeholders instead of actual answers.
Intersight supports the golang template syntax specified in https://golang.org/pkg/text/template/.
The template supports placeholders for all the properties of os.Answers MO type
as well as any additional user-defined properties. The values for these placeholders
shall be given during OS installation in the form of os.Answers type and 'additionalProperties' in
os.OsInstall object.
---

# Data Source: intersight_os._configuration_file
A ConfigurationFile is an OS specific answer file that helps with the unattended
installation.
The file can also be a template file with placeholders instead of actual answers.
Intersight supports the golang template syntax specified in https://golang.org/pkg/text/template/.
The template supports placeholders for all the properties of os.Answers MO type
as well as any additional user-defined properties. The values for these placeholders
shall be given during OS installation in the form of os.Answers type and 'additionalProperties' in
os.OsInstall object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `file_content`:(string) The content of the entire configuration file is stored as value. The contentcan either be a static file content or a template content.The template is expected to conform to the golang template syntax. The valuesfrom os.Answers properties will be used to populate this template. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS ConfigurationFile that uniquely identifies the configuration file. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `supported`:(bool) An internal property that is used to distinguish between the pre-canned OSconfiguration file entries and user provided entries. 
