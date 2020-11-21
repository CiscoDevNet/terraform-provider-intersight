
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_lc"
sidebar_current: "docs-intersight-data-source-niatelemetry-lc"
description: |-
Object is available at Line Card scope.
---

# Data Source: intersight_niatelemetry._lc
Object is available at Line Card scope.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the line cards present. 
* `dn`:(string) Dn value for the line cards present. 
* `hardware_version`:(string) Hardware version of the line cards present. 
* `model`:(string) Model of the line cards present. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `operational_state`:(string) Opretaional state of the line cards present. 
* `power_state`:(string) Power state of the line cards present. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `redundancy_state`:(string) Redundancy state of the line cards present. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
