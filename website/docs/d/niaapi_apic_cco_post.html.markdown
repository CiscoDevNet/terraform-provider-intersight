
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_cco_post"
sidebar_current: "docs-intersight-data-source-niaapi-apic-cco-post"
description: |-
The post reporting a new release is available for APIC.
---

# Data Source: intersight_niaapi._apic_cco_post
The post reporting a new release is available for APIC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `post_date`:(string) The date when this new release notice is posted. 
* `post_type`:(string) The document type of this post. 
* `postid`:(string) Identificator of this inbox post. 
* `revision`:(string) Revision number of this notice. 
