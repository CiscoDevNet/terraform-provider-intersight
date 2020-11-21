
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_field_notice"
sidebar_current: "docs-intersight-data-source-niaapi-apic-field-notice"
description: |-
The field notice reporting bug and related software or hardware for APIC.
---

# Data Source: intersight_niaapi._apic_field_notice
The field notice reporting bug and related software or hardware for APIC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bugid`:(string) Bug Id associated with this notice. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `field_notice_desc`:(string) Field notice Description. 
* `field_notice_id`:(string) Fieldnotice Id of this notice. 
* `field_notice_url`:(string) Field notice URL link to the notice webpage. 
* `headline`:(string) The headline of this field notice. 
* `hwpid`:(string) Hardware PID for affected models. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `sw_release`:(string) Software Release number for affected versions. 
* `workaround_url`:(string) URL of workaround of this notice. 
