
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_release_note"
sidebar_current: "docs-intersight-data-source-appliance-release-note"
description: |-
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
---

# Data Source: intersight_appliance._release_note
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `version`:(string) Version number of the pending upgrade. 
