
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `date`:(string) The date when this package is generated. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the Metadata file downloaded. 
* `metadata_filename`:(string) The Filename of this Metadata package. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `version`:(int) The version number of the Metadata package. 
