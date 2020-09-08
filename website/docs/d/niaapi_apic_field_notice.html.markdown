
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `field_notice_desc`:(string) Field notice Description. 
* `field_notice_id`:(string) Fieldnotice Id of this notice. 
* `field_notice_url`:(string) Field notice URL link to the notice webpage. 
* `headline`:(string) The headline of this field notice. 
* `hwpid`:(string) Hardware PID for affected models. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `sw_release`:(string) Software Release number for affected versions. 
* `workaround_url`:(string) URL of workaround of this notice. 
