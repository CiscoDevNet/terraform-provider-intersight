---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_unit_descriptor"
description: |-
  Descriptor that uniquely identifies an adapter.
---

# Data Source: intersight_capability_adapter_unit_descriptor
Descriptor that uniquely identifies an adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_adapter_unit_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_generation`:(int) Generation of the adapter.* `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04.* `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02.* `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03.* `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string. 
* `connectivity_order`:(string) Order in which the ports are connected; sequential or cyclic. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ethernet_port_speed`:(int) The port speed for ethernet ports in Mbps. 
* `fibre_channel_port_speed`:(int) The port speed for fibre channel ports in Mbps. 
* `fibre_channel_scsi_ioq_limit`:(int) The number of SCSI I/O Queue resources to allocate. 
* `is_azure_qos_supported`:(bool) Indicates that the Azure Stack Host QoS feature is supported by this adapter. 
* `is_geneve_supported`:(bool) Indicates that the GENEVE offload feature is supported by this adapter. 
* `max_rocev2_interfaces`:(int) Maximum number of vNIC interfaces that can be RoCEv2 enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_dce_ports`:(int) Number of Dce Ports for the adapter. 
* `pci_link`:(int) Indicates PCI Link status of adapter. 
* `prom_card_type`:(string) Prom card type for the adapter. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 