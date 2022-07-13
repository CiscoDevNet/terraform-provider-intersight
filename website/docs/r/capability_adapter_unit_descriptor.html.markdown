---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_unit_descriptor"
description: |-
        Descriptor that uniquely identifies an adapter.

---

# Resource: intersight_capability_adapter_unit_descriptor
Descriptor that uniquely identifies an adapter.
## Usage Example
### Resource Creation

```hcl
resource "intersight_capability_adapter_unit_descriptor" "capability_adapter_unit_descriptor1" {
  description = "capability adapter unit descriptor"
  model       = "Cisco Systems Inc"
  capabilities {
    moid        = var.capability_adapter_unit_descriptor
    object_type = "capability.AdapterUnitDescriptor"
    class_id    = "capability.AdapterUnitDescriptor"
  }
  connectivity_order  = "sequential"
  ethernet_port_speed = 40
}

variable "capability_adapter_unit_descriptor" {
  type        = string
  description = "Moid of capability.AdapterUnitDescriptor Mo"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `adapter_generation`:(int) Generation of the adapter.* `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04.* `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02.* `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03.* `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `capabilities`:(Array) An array of relationships to capabilityCapability resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `connectivity_order`:(string) Order in which the ports are connected; sequential or cyclic. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `ethernet_port_speed`:(int) The port speed for ethernet ports in Mbps. 
* `features`:(Array)
This complex property has following sub-properties:
  + `feature_name`:(string) Name of the feature that identifies the specific adapter configuration.* `RoCEv2` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2.* `RoCEv1` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1.* `VMQ` - Capability indicator of the Virtual Machine Queue (VMQ) feature.* `VMMQ` - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature.* `VMQInterrupts` - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature.* `NVGRE` - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature.* `ARFS` - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature.* `VXLAN` - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. 
  + `min_fw_version`:(string) Firmware version from which support for this feature is available. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `supported_fw_versions`:
                (Array of schema.TypeString) -
  + `supported_in_adapters`:
                (Array of schema.TypeString) -
  + `supported_in_generations`:
                (Array of schema.TypeInt) -
* `fibre_channel_port_speed`:(int) The port speed for fibre channel ports in Mbps. 
* `fibre_channel_scsi_ioq_limit`:(int) The number of SCSI I/O Queue resources to allocate. 
* `is_azure_qos_supported`:(bool) Indicates that the Azure Stack Host QoS feature is supported by this adapter. 
* `is_geneve_supported`:(bool) Indicates that the GENEVE offload feature is supported by this adapter. 
* `max_eth_rx_ring_size`:(int) Maximum Ring Size value for vNIC Receive Queue. 
* `max_eth_tx_ring_size`:(int) Maximum Ring Size value for vNIC Transmit Queue. 
* `max_rocev2_interfaces`:(int) Maximum number of vNIC interfaces that can be RoCEv2 enabled. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_dce_ports`:(int) Number of Dce Ports for the adapter. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pci_link`:(int) Indicates PCI Link status of adapter. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `prom_card_type`:(string) Prom card type for the adapter. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
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
`intersight_capability_adapter_unit_descriptor` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_adapter_unit_descriptor.example 1234567890987654321abcde
``` 
