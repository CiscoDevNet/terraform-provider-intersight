
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_catalog"
sidebar_current: "docs-intersight-data-source-softwarerepository-catalog"
description: |-
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
---

# Data Source: intersight_softwarerepository._catalog
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the catalog. The names are populated and predefined during MO creation. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
