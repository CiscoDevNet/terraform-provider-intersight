
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_latest_maintained_release"
sidebar_current: "docs-intersight-data-source-niaapi-dcnm-latest-maintained-release"
description: |-
This contains the latest maintained release information for each release on DCNM.
---

# Data Source: intersight_niaapi._dcnm_latest_maintained_release
This contains the latest maintained release information for each release on DCNM.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `maintained_release`:(string) Lastest maintained release. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `software_release`:(string) Software release version string. 
* `version_tag`:(string) Long lived version or short lived version. 
