
---
layout: "intersight"
page_title: "Intersight: intersight_vnic_lcp_status"
sidebar_current: "docs-intersight-data-source-vnic-lcp-status"
description: |-
An internal MO to check if a LCP can be deployed or not on a specific Server Profile.
---

# Data Source: intersight_vnic._lcp_status
An internal MO to check if a LCP can be deployed or not on a specific Server Profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `reason`:(string) The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. 
* `status`:(string) Indicates if the LCP is ready for Deploy or not.* `ok` - No issues with the LCP/SCP/VIF.* `error` - The LCP/SCP/VIF cannot be deployed due to error.* `validating` - Validation in progress for the LCP. 
