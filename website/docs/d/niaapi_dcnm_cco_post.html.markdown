
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_cco_post"
sidebar_current: "docs-intersight-data-source-niaapi-dcnm-cco-post"
description: |-
The post reporting a new release is available for DCNM.
---

# Data Source: intersight_niaapi._dcnm_cco_post
The post reporting a new release is available for DCNM.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `post_date`:(string) The date when this new release notice is posted. 
* `post_type`:(string) The document type of this post. 
* `postid`:(string) Identificator of this inbox post. 
* `revision`:(string) Revision number of this notice. 
