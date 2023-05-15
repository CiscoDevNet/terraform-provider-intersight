---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_supported_driver_name"
description: |-
        Supported driver names for a given product for the given operating system.

---

# Resource: intersight_hcl_supported_driver_name
Supported driver names for a given product for the given operating system.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `os_vendor`:(string) Vendor distributing the Operating System. 
* `os_version`:(string) Version of the Operating System. 
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
* `product_list`:(Array)
This complex property has following sub-properties:
  + `driver_names`:
                (Array of schema.TypeString) -
  + `error_code`:(string)(ReadOnly) Error code indicating the support status.* `Success` - The input combination is valid.* `Unknown` - Unknown API request to the service.* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.* `InvalidUcsVersion` - UCS Version is not in expected format.* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.* `OSUnknown` - OS vendor or version is not known as per the HCL database.* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.* `ProductUnknown` - Product is not known as per the HCL database.* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.* `InternalError` - API requests to the service have either failed or timed out.* `MarshallingError` - Error in JSON marshalling.* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. 
  + `firmwares`:(Array)
This complex property has following sub-properties:
    + `driver_name`:(string) Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3. 
    + `driver_version`:(string) Version of the Driver supported. 
    + `error_code`:(string)(ReadOnly) Error code for the support status.* `Success` - The input combination is valid.* `Unknown` - Unknown API request to the service.* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.* `InvalidUcsVersion` - UCS Version is not in expected format.* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.* `OSUnknown` - OS vendor or version is not known as per the HCL database.* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.* `ProductUnknown` - Product is not known as per the HCL database.* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.* `InternalError` - API requests to the service have either failed or timed out.* `MarshallingError` - Error in JSON marshalling.* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes. 
    + `firmware_version`:(string) Firmware version of the product/adapter supported. 
    + `id`:(string) Identifier of the firmware. 
    + `latest_driver`:(bool)(ReadOnly) True if the driver is latest recommended driver. 
    + `latest_firmware`:(bool)(ReadOnly) True if the firmware is latest recommended firmware. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `id`:(string) Identifier of the product. 
  + `model`:(string) Model/PID of the product/adapter. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `revision`:(string) Revision of the adapter model. 
  + `type`:(string) Type of the product/adapter say OCP, PT, GPU.* `` - Default type of the Product.* `Adapter` - Represents network adapter cards.* `StorageController` - Represents storage controllers.* `GPU` - Represents graphics cards. 
  + `vendor`:(string) Vendor of the product or adapter. 
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
`intersight_hcl_supported_driver_name` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hcl_supported_driver_name.example 1234567890987654321abcde
``` 
