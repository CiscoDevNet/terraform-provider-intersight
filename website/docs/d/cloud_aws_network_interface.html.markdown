---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_aws_network_interface"
description: |-
        AWS Network Interface object is represented here.It is Virtual interface that can be attached to
        an instance in a Virtual Private Cloud (VPC).

---

# Data Source: intersight_cloud_aws_network_interface
AWS Network Interface object is represented here.It is Virtual interface that can be attached to
an instance in a Virtual Private Cloud (VPC).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_aws_network_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_type`:(string) Type or model of the virtual network interface card.* `Unknown` - The type of the network adaptor type is unknown.* `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC.* `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support.* `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.* `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features.* `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later.* `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets.* `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length.* `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance.* `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \ virtual\  devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \ emulated\  devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware.* `` - Default network adaptor type supported by the hypervisor. 
* `cidr`:(string) CIDR scheme for defining an IP block. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User friendly description of network interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) Internally generated identity of this network interface. 
* `mac_address`:(string) MAC address assigned to the virtual network interface card. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual network interface card. 
* `nic_create_time`:(string) Time when this network interface is created. 
* `private_dns_name`:(string) The private DNS hostname name assigned to the network interface. 
* `public_dns_name`:(string) The public DNS hostname name assigned to the network interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the network interface. If the network interface is not attached to an instance, the statusis available; if a network interface is attached to an instance the status is in-use. 
 
