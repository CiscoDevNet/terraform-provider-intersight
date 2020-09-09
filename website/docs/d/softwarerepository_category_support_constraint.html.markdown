
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `constraint_id`:(string) Identifier for this managed object. 
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `min_version`:(string) Minimum image version from where the models can be supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `parse_from_image_name`:(bool) Fields which tells if the constraint is based on image name parsing. 
