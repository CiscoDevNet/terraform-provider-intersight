
---
layout: "intersight"
page_title: "Intersight: intersight_os_supported_version"
sidebar_current: "docs-intersight-data-source-os-supported-version"
description: |-
The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
---

# Data Source: intersight_os._supported_version
The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `version_name`:(string) The OsInstall Supported Operating System Version Name. 
