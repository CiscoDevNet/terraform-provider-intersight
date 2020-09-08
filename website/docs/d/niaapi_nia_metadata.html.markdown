
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nia_metadata"
sidebar_current: "docs-intersight-data-source-niaapi-nia-metadata"
description: |-
Contains the latest Metadata available for download from server.
---

# Data Source: intersight_niaapi._nia_metadata
Contains the latest Metadata available for download from server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `date`:(string) The date when this package is generated. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the Metadata file downloaded. 
* `metadata_filename`:(string) The Filename of this Metadata package. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `version`:(int) The version number of the Metadata package. 
