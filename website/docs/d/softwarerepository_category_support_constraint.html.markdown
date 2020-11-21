
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_support_constraint"
sidebar_current: "docs-intersight-data-source-softwarerepository-category-support-constraint"
description: |-
Defines constraints for models which are supported from certain minimum image version.
---

# Data Source: intersight_softwarerepository._category_support_constraint
Defines constraints for models which are supported from certain minimum image version.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `constraint_id`:(string) Identifier for this managed object. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `min_version`:(string) Minimum image version from where the models can be supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `parse_from_image_name`:(bool) Fields which tells if the constraint is based on image name parsing. 
