
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `item`:(string) Message target type of the alerts. 
* `last_checked`:(string) Last checked time of the alerts. 
* `message`:(string) Detailed info about the alert. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `recommendation`:(string) Recommended fix for the alert. 
* `status`:(string) Status of the given alert. 
* `status_id`:(string) Status Id of the given alert. 
