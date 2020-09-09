
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_diagnostic_messages"
sidebar_current: "docs-intersight-data-source-iaas-diagnostic-messages"
description: |-
Gets diagnostics messages from UCSD.
---

# Data Source: intersight_iaas._diagnostic_messages
Gets diagnostics messages from UCSD.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) Message category of the alerts. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `item`:(string) Message target type of the alerts. 
* `last_checked`:(string) Last checked time of the alerts. 
* `message`:(string) Detailed info about the alert. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `recommendation`:(string) Recommended fix for the alert. 
* `status`:(string) Status of the given alert. 
* `status_id`:(string) Status Id of the given alert. 
