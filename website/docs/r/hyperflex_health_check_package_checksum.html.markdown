---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_package_checksum"
description: |-
  HyperFlex health check Debian Package SHA512 checksum.
---

# Resource: intersight_hyperflex_health_check_package_checksum
HyperFlex health check Debian Package SHA512 checksum.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `checksum`:(string) SHA512 checksum of the health check package. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check Debian package. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `package_name`:(string) HyperFlex health check Debian package file name. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `timestamp`:(string) Timestamp of last update of HyperFlex health check package checksum. 
* `nr_version`:(string) HyperFlex health check Debian Package Version. 


## Import
`intersight_hyperflex_health_check_package_checksum` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_health_check_package_checksum.example 1234567890987654321abcde
```