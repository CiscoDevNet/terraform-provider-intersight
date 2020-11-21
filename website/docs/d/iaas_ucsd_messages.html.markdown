
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_messages"
sidebar_current: "docs-intersight-data-source-iaas-ucsd-messages"
description: |-
Gets ucsd messages from UCSD.
---

# Data Source: intersight_iaas._ucsd_messages
Gets ucsd messages from UCSD.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `status_id`:(string) Last checked time of the alerts. 
