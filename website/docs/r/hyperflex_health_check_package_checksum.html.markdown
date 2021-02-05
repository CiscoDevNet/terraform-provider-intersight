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
* `checksum`:(string) SHA512 checksum of the health check package. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check Debian package. 
* `package_name`:(string) HyperFlex health check Debian package file name. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `timestamp`:(string) Timestamp of last update of HyperFlex health check package checksum. 
* `nr_version`:(string) HyperFlex health check Debian Package Version. 


## Import
`intersight_hyperflex_health_check_package_checksum` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_health_check_package_checksum.example 1234567890987654321abcde
``` 