---
subcategory: "power"
layout: "intersight"
page_title: "Intersight: intersight_power_policy"
description: |-
        Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.

---

# Resource: intersight_power_policy
Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `allocated_budget`:(int) Sets the limit for the maximum input power consumption by the chassis (in Watts). Set to 0 for no limit. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `dynamic_rebalancing`:(string) Sets the dynamic power rebalancing mode of the chassis. If enabled, this mode allows the chassis to dynamically reallocate the power between servers depending on their power usage.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `extended_power_capacity`:(string) Sets the Extended Power Capacity of the Chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies.  This option is only supported for Cisco UCS X series Chassis.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `power_priority`:(string) Sets the Power Priority of the Server. This priority is used to determine the initial power allocation for servers. This field is only supported for Cisco UCS B series and X series servers.* `Low` - Set the Power Priority to Low.* `Medium` - Set the Power Priority to Medium.* `High` - Set the Power Priority to High. 
* `power_profiling`:(string) Sets the Power Profiling of the Server. If Enabled, this field allows the power manager to run power profiling  utility to determine the power needs of the server.  This field is only supported for Cisco UCS X series servers.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `power_restore_state`:(string) Sets the Power Restore State of the Server. In the absence of Intersight connectivity, the chassis/server will use this policy  to recover the host power after a power loss event.* `AlwaysOff` - Set the Power Restore Mode to Off.* `AlwaysOn` - Set the Power Restore Mode to On.* `LastState` - Set the Power Restore Mode to LastState. 
* `power_save_mode`:(string) Sets the power save mode of the chassis. If the requested power budget is less than available power capacity,  the additional PSUs not required to comply with redundancy policy are placed in power save mode.* `Enabled` - Set the value to Enabled.* `Disabled` - Set the value to Disabled. 
* `processor_package_power_limit`:(string) Sets the Processor Package Power Limit (PPL) of a server. PPL refers to the amount of power that a CPU can draw from the power supply. The Processor Package Power Limit (PPL) feature is currently available exclusively on Cisco UCS C225/C245 M8 servers.* `Default` - Set the Package Power Limit to the platform defined default value.* `Maximum` - Set the Package Power Limit to the platform defined maximum value.* `Minimum` - Set the Package Power Limit to the platform defined minimum value. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `redundancy_mode`:(string) Sets the Power Redundancy Mode of the Chassis.  Redundancy Mode determines the number of PSUs the chassis keeps as redundant.  N+2 mode is only supported for Cisco UCS X series Chassis.* `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis.* `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained.* `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy.* `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_power_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_power_policy.example 1234567890987654321abcde
``` 
