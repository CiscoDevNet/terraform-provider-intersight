---
subcategory: "softwarerepository"
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_category_support_constraint"
description: |-
  Defines constraints for models which are supported from certain minimum image version.
---

# Resource: intersight_softwarerepository_category_support_constraint
Defines constraints for models which are supported from certain minimum image version.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `constraint_id`:(string) Identifier for this managed object. 
* `filtered_models`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `min_version`:(string) Minimum version of the image. 
  + `name`:(string) Name of the contraint, used to identify constriant type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `platform_regex`:(string) Regular expression of the image name. 
  + `supported_models`:
                (Array of schema.TypeString) -
* `mdf_id`:(string) Cisco software repository image category identifier. 
* `min_version`:(string) Minimum image version from where the models can be supported. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `parse_from_image_name`:(bool) Fields which tells if the constraint is based on image name parsing. 
* `supported_models`:
                (Array of schema.TypeString) -
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_softwarerepository_category_support_constraint` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_softwarerepository_category_support_constraint.example 1234567890987654321abcde
```