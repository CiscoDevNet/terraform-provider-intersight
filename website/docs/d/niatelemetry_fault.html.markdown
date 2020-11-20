
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_fault"
sidebar_current: "docs-intersight-data-source-niatelemetry-fault"
description: |-
Object is available at Fault scope in a fabric and provides details about a fault occurred.
---

# Data Source: intersight_niatelemetry._fault
Object is available at Fault scope in a fabric and provides details about a fault occurred.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `code`:(string) Code of the fault present. 
* `description`:(string) Description of the fault present. 
* `dn`:(string) Dn value for the fault present. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `severity`:(string) Severity of the fault present. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters/sites. 
* `type`:(string) Type of the fault present. 
