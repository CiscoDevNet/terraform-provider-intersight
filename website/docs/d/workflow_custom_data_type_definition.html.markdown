
---
layout: "intersight"
page_title: "Intersight: intersight_workflow_custom_data_type_definition"
sidebar_current: "docs-intersight-data-source-workflow-custom-data-type-definition"
description: |-
Captures a customized data type definition that can be used for task or workflow input/output.  This can be reused across multiple tasks and workflow definitions.
---

# Data Source: intersight_workflow._custom_data_type_definition
Captures a customized data type definition that can be used for task or workflow input/output.  This can be reused across multiple tasks and workflow definitions.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `composite_type`:(bool) When true this data type definition is a collection of type definitions to represent composite data like JSON. 
* `description`:(string) A human-friendly description of this custom data type indicating it's domain and usage. 
* `label`:(string) A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
