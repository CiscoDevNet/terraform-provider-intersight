---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_impact"
description: |-
        Before submitting firmware upgrade operation, the estimate impact helps to know the list of components be impacted and require host reboot. This cannot be used for network share based upgrade.

---

# Resource: intersight_firmware_upgrade_impact
Before submitting firmware upgrade operation, the estimate impact helps to know the list of components be impacted and require host reboot. This cannot be used for network share based upgrade.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `chassis`:(Array) An array of relationships to equipmentChassis resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `components`:
                (Array of schema.TypeString) -
* `computation_state`:(string) Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since image is not present in FI. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `device`:(Array)(ReadOnly) An array of relationships to assetDeviceRegistration resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `distributable`:(HashMap) - A reference to a firmwareDistributable resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `exclude_component_list`:
                (Array of schema.TypeString) -
* `exclude_components`:
                (Array of schema.TypeString) -
* `impacts`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [firmware.ChassisUpgradeImpact](#firmwareChassisUpgradeImpact)
[firmware.ComponentImpact](#firmwareComponentImpact)
[firmware.FabricUpgradeImpact](#firmwareFabricUpgradeImpact)
[firmware.ServerUpgradeImpact](#firmwareServerUpgradeImpact)
  + `computation_error`:(string) Details for the error that occurred during the reboot validation analysis. 
  + `computation_status_detail`:(string) The computation status of the estimate operation for a component.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect.* `Failed` - Upgrade impact is not available due to an unknown error. 
  + `domain_name`:(string) The endpoint type or name. 
  + `end_point`:(string) A reference to a REST resource, uniquely identified by object type and MOID. 
  + `firmware_version`:(string) The current firmware version of the component. 
  + `impact_type`:(string) The impact type of the endpoint, whether a reboot is required or not.* `NoReboot` - A reboot is not required for the endpoint after upgrade.* `Reboot` - A reboot is required to the endpoint after upgrade. 
  + `model`:(string) The model details of the component. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `target_firmware_version`:(string) The target firmware version of the component. 
  + `version_impact`:(string) The change of version impact for the endpoint.* `None` - No change in version for the component.* `Upgrade` - The component will be upgraded.* `Downgrade` - The component will be downgraded. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `network_elements`:(Array) An array of relationships to networkElement resources. 
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
* `pci_node`:(Array) An array of relationships to pciNode resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `release`:(HashMap) - A reference to a softwarerepositoryRelease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server`:(Array) An array of relationships to computePhysical resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `summary`:(string) The summary on the component or components getting impacted by the upgrade.* `Basic` - Summary of a single instance involved in the upgrade operation.* `Detail` - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
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
* `xfm_upgrade_option`:(string) XFM upgrade option Full or Partial Disruption.* `none` - If no option is selected for exclusion.* `full-shutdown` - PSX Switch in XFM will be upgraded in single action.* `partial-shutdown` - PSX Switch in XFM will be upgraded one after other. 


## Import
`intersight_firmware_upgrade_impact` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_firmware_upgrade_impact.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [firmware.ChassisUpgradeImpact](#argument-reference)
Impact of the chassis endpoint during the upgrade operation.
* `impact_detail`:(Array)
This complex property has following sub-properties:
  + `component`:(string) Impact on the component after the upgrade.* `ALL` - This represents all the components.* `ALL,HDD` - This represents all the components plus the HDDs.* `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).* `Storage` - This represents the storage controller components.* `None` - This represents none of the components.* `NXOS` - This represents NXOS components.* `ESU` - This represents ESU components.* `IOM` - This represents IOM components.* `PSU` - This represents PSU components.* `CIMC` - This represents CIMC components.* `BIOS` - This represents BIOS components.* `PCIE` - This represents PCIE components.* `Drive` - This represents Drive components.* `DIMM` - This represents DIMM components.* `BoardController` - This represents Board Controller components.* `StorageController` - This represents Storage Controller components.* `Storage-Sasexpander` - This represents Storage Sas-Expander components.* `Storage-U.2` - This represents U2 Storage Controller components.* `HBA` - This represents HBA components.* `GPU` - This represents GPU components.* `SasExpander` - This represents SasExpander components.* `MSwitch` - This represents mSwitch components.* `CMC` - This represents CMC components.* `PSX` - This represents PSX components. 
  + `computation_error`:(string) Details for the error that occurred during the reboot validation analysis. 
  + `computation_status_detail`:(string) The computation status of the estimate operation for a component.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect.* `Failed` - Upgrade impact is not available due to an unknown error. 
  + `domain_name`:(string) The endpoint type or name. 
  + `end_point`:(string) A reference to a REST resource, uniquely identified by object type and MOID. 
  + `firmware_version`:(string) The current firmware version of the component. 
  + `impact_type`:(string) The impact type of the endpoint, whether a reboot is required or not.* `NoReboot` - A reboot is not required for the endpoint after upgrade.* `Reboot` - A reboot is required to the endpoint after upgrade. 
  + `model`:(string) The model details of the component. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_firmware_version`:(string) The target firmware version of the component. 
  + `version_impact`:(string) The change of version impact for the endpoint.* `None` - No change in version for the component.* `Upgrade` - The component will be upgraded.* `Downgrade` - The component will be downgraded. 
* `impact_servers_detail`:(Array)
This complex property has following sub-properties:
  + `display_name`:(string) Display name of the server impacted by the upgrade. 
  + `name`:(string) Name of the server impacted by the upgrade. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `reboot_at`:(string) The reboot impact for the server endpoint during the upgrade operation.* `Phase1` - Server will be rebooted in the first phase of upgrade.* `Phase2` - Server will be rebooted in the second phase of upgrade. 
  + `server`:(HashMap) - The server on which this upgrade operation is performed. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `name`:(string) Name of the chassis that will be affected by the upgrade. 
* `user_label`:(string) Details for the chassis that will be impacted by the upgrade. 

### [firmware.ComponentImpact](#argument-reference)
Impact for the components such as CIMC, BIOS, and disk during the upgrade operation.
* `component`:(string) Impact on the component after the upgrade.* `ALL` - This represents all the components.* `ALL,HDD` - This represents all the components plus the HDDs.* `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).* `Storage` - This represents the storage controller components.* `None` - This represents none of the components.* `NXOS` - This represents NXOS components.* `ESU` - This represents ESU components.* `IOM` - This represents IOM components.* `PSU` - This represents PSU components.* `CIMC` - This represents CIMC components.* `BIOS` - This represents BIOS components.* `PCIE` - This represents PCIE components.* `Drive` - This represents Drive components.* `DIMM` - This represents DIMM components.* `BoardController` - This represents Board Controller components.* `StorageController` - This represents Storage Controller components.* `Storage-Sasexpander` - This represents Storage Sas-Expander components.* `Storage-U.2` - This represents U2 Storage Controller components.* `HBA` - This represents HBA components.* `GPU` - This represents GPU components.* `SasExpander` - This represents SasExpander components.* `MSwitch` - This represents mSwitch components.* `CMC` - This represents CMC components.* `PSX` - This represents PSX components. 

### [firmware.FabricUpgradeImpact](#argument-reference)
Impact for the Fabric Interconnect endpoint during the upgrade operation.
* `impact_detail`:(Array)
This complex property has following sub-properties:
  + `component`:(string) Impact on the component after the upgrade.* `ALL` - This represents all the components.* `ALL,HDD` - This represents all the components plus the HDDs.* `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).* `Storage` - This represents the storage controller components.* `None` - This represents none of the components.* `NXOS` - This represents NXOS components.* `ESU` - This represents ESU components.* `IOM` - This represents IOM components.* `PSU` - This represents PSU components.* `CIMC` - This represents CIMC components.* `BIOS` - This represents BIOS components.* `PCIE` - This represents PCIE components.* `Drive` - This represents Drive components.* `DIMM` - This represents DIMM components.* `BoardController` - This represents Board Controller components.* `StorageController` - This represents Storage Controller components.* `Storage-Sasexpander` - This represents Storage Sas-Expander components.* `Storage-U.2` - This represents U2 Storage Controller components.* `HBA` - This represents HBA components.* `GPU` - This represents GPU components.* `SasExpander` - This represents SasExpander components.* `MSwitch` - This represents mSwitch components.* `CMC` - This represents CMC components.* `PSX` - This represents PSX components. 
  + `computation_error`:(string) Details for the error that occurred during the reboot validation analysis. 
  + `computation_status_detail`:(string) The computation status of the estimate operation for a component.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect.* `Failed` - Upgrade impact is not available due to an unknown error. 
  + `domain_name`:(string) The endpoint type or name. 
  + `end_point`:(string) A reference to a REST resource, uniquely identified by object type and MOID. 
  + `firmware_version`:(string) The current firmware version of the component. 
  + `impact_type`:(string) The impact type of the endpoint, whether a reboot is required or not.* `NoReboot` - A reboot is not required for the endpoint after upgrade.* `Reboot` - A reboot is required to the endpoint after upgrade. 
  + `model`:(string) The model details of the component. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_firmware_version`:(string) The target firmware version of the component. 
  + `version_impact`:(string) The change of version impact for the endpoint.* `None` - No change in version for the component.* `Upgrade` - The component will be upgraded.* `Downgrade` - The component will be downgraded. 
* `serial`:(string) Details for the Fabric Interconnect that will be impacted by the upgrade. 

### [firmware.ServerUpgradeImpact](#argument-reference)
Impact of the server endpoint during the upgrade operation.
* `impact_detail`:(Array)
This complex property has following sub-properties:
  + `component`:(string) Impact on the component after the upgrade.* `ALL` - This represents all the components.* `ALL,HDD` - This represents all the components plus the HDDs.* `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).* `Storage` - This represents the storage controller components.* `None` - This represents none of the components.* `NXOS` - This represents NXOS components.* `ESU` - This represents ESU components.* `IOM` - This represents IOM components.* `PSU` - This represents PSU components.* `CIMC` - This represents CIMC components.* `BIOS` - This represents BIOS components.* `PCIE` - This represents PCIE components.* `Drive` - This represents Drive components.* `DIMM` - This represents DIMM components.* `BoardController` - This represents Board Controller components.* `StorageController` - This represents Storage Controller components.* `Storage-Sasexpander` - This represents Storage Sas-Expander components.* `Storage-U.2` - This represents U2 Storage Controller components.* `HBA` - This represents HBA components.* `GPU` - This represents GPU components.* `SasExpander` - This represents SasExpander components.* `MSwitch` - This represents mSwitch components.* `CMC` - This represents CMC components.* `PSX` - This represents PSX components. 
  + `computation_error`:(string) Details for the error that occurred during the reboot validation analysis. 
  + `computation_status_detail`:(string) The computation status of the estimate operation for a component.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since the image is not present in the Fabric Interconnect.* `Failed` - Upgrade impact is not available due to an unknown error. 
  + `domain_name`:(string) The endpoint type or name. 
  + `end_point`:(string) A reference to a REST resource, uniquely identified by object type and MOID. 
  + `firmware_version`:(string) The current firmware version of the component. 
  + `impact_type`:(string) The impact type of the endpoint, whether a reboot is required or not.* `NoReboot` - A reboot is not required for the endpoint after upgrade.* `Reboot` - A reboot is required to the endpoint after upgrade. 
  + `model`:(string) The model details of the component. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `target_firmware_version`:(string) The target firmware version of the component. 
  + `version_impact`:(string) The change of version impact for the endpoint.* `None` - No change in version for the component.* `Upgrade` - The component will be upgraded.* `Downgrade` - The component will be downgraded. 
* `name`:(string) Name of the server impacted by the upgrade. 
* `user_label`:(string) Details about the server which will be impacted by the upgrade. 
  
