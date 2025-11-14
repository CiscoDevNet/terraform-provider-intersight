---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_if"
description: |-
        Virtual Ethernet Interface.

---

# Resource: intersight_vnic_eth_if
Virtual Ethernet Interface.
## Usage Example
### Resource Creation

```hcl
resource "intersight_vnic_eth_if" "eth1" {
  name  = "eth0"
  order = 0
  placement {
    id          = "1"
    pci_link    = 0
    uplink      = 0
    object_type = "vnic.PlacementSettings"
  }
  cdn {
    value       = "VIC-1-eth00"
    nr_source   = "user"
    object_type = "vnic.Cdn"
  }
  usnic_settings {
    cos         = 5
    nr_count    = 0
    object_type = "vnic.UsnicSettings"
  }
  vmq_settings {
    enabled             = true
    multi_queue_support = false
    num_interrupts      = 1
    num_vmqs            = 1
    object_type         = "vnic.VmqSettings"
  }
  lan_connectivity_policy {
    moid        = var.vnic_lan1
    object_type = "vnic.LanConnectivityPolicy"
  }
  eth_network_policy {
    moid        = var.v_eth_network1
    object_type = "vnic.EthNetworkPolicy"
  }
  eth_adapter_policy {
    moid        = var.v_eth_adapter1
    object_type = "vnic.EthAdapterPolicy"
  }
  eth_qos_policy {
    moid        = var.v_eth_qos1
    object_type = "vnic.EthQosPolicy"
  }
}

variable "vnic_lan1" {
  type        = string
  description = "Moid of vnic.LanConnectivityPolicy"
}

variable "v_eth_network1" {
  type        = string
  description = "Moid of vnic.EthNetworkPolicy"
}

variable "v_eth_adapter1" {
  type        = string
  description = "Moid of vnic.EthAdapterPolicy"
}

variable "v_eth_qos1" {
  type        = string
  description = "Moid of vnic.EthQosPolicy"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cdn`:(HashMap) - Consistent Device Naming configuration for the virtual NIC. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `nr_source`:(string) Source of the CDN. It can either be user specified or be the same as the vNIC name.* `vnic` - Source of the CDN is the same as the vNIC name.* `user` - Source of the CDN is specified by the user. 
  + `value`:(string) The CDN value entered in case of user defined mode. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `eth_adapter_policy`:(HashMap) - A reference to a vnicEthAdapterPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `eth_network_policy`:(HashMap) - A reference to a vnicEthNetworkPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `eth_qos_policy`:(HashMap) - A reference to a vnicEthQosPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `fabric_eth_network_control_policy`:(HashMap) - A reference to a fabricEthNetworkControlPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `fabric_eth_network_group_policy`:(Array) An array of relationships to fabricEthNetworkGroupPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `failover_enabled`:(bool) Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster. 
* `ip_lease`:(HashMap) -(ReadOnly) A reference to a ippoolIpLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `iscsi_boot_policy`:(HashMap) - A reference to a vnicIscsiBootPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `iscsi_ip_v4_address_allocation_type`:(string)(ReadOnly) Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.* `None` - Type indicates that there is no IP associated to an vnic.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `iscsi_ip_v4_config`:(HashMap) -(ReadOnly) IPV4 configurations such as Netmask, Gateway and DNS for iSCSI vNIC. 
This complex property has following sub-properties:
  + `gateway`:(string) IP address of the default IPv4 gateway. 
  + `netmask`:(string) A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `primary_dns`:(string) IP Address of the primary Domain Name System (DNS) server. 
  + `secondary_dns`:(string) IP Address of the secondary Domain Name System (DNS) server. 
* `iscsi_ip_v6_address_allocation_type`:(string)(ReadOnly) Static/Pool/DHCP Type of IPv6 address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.* `None` - Type indicates that there is no IP associated to an vnic.* `DHCP` - The IP address is assigned using DHCP, if available.* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool. 
* `iscsi_ip_v6_config`:(HashMap) -(ReadOnly) IPv6 configurations such as Prefix, Gateway and DNS for iSCSI vNIC. 
This complex property has following sub-properties:
  + `gateway`:(string) IP address of the default IPv6 gateway. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `prefix`:(int) A prefix length which masks the  IP address and divides the IP address into network address and host address. 
  + `primary_dns`:(string) IP Address of the primary Domain Name System (DNS) server. 
  + `secondary_dns`:(string) IP Address of the secondary Domain Name System (DNS) server. 
* `iscsi_ipv4_address`:(string)(ReadOnly) IP address associated to the vNIC. 
* `iscsi_ipv6_address`:(string)(ReadOnly) IPv6 address associated to the iSCSI vNIC. 
* `lan_connectivity_policy`:(HashMap) - A reference to a vnicLanConnectivityPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `lcp_vnic`:(HashMap) -(ReadOnly) A reference to a vnicEthIf resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mac_address`:(string)(ReadOnly) The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy. 
* `mac_address_type`:(string) Type of allocation selected to assign a MAC address for the vnic.* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface. 
* `mac_lease`:(HashMap) - A reference to a macpoolLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mac_pool`:(HashMap) - A reference to a macpoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual ethernet interface. 
* `old_info`:(HashMap) -(ReadOnly) Deployed information related to lease and vif id. 
This complex property has following sub-properties:
  + `mac_address`:(string)(ReadOnly) Old MAC Address associated with the interface. 
  + `mac_address_type`:(string)(ReadOnly) Type of allocation selected to assign a MAC address for the vnic.* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface. 
  + `mac_lease`:(HashMap) -(ReadOnly) Mac Address Lease taken for the vnic configured in the Lan Connectivity Policy. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `mac_pool`:(HashMap) -(ReadOnly) Mac Address Pool configured in the Lan Connectivity Policy. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `standby_vif_id`:(int)(ReadOnly) Old Standby Vif id that was associated with the interface. 
  + `static_mac_address`:(string)(ReadOnly) The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use thefollowing MAC prefix 00:25:B5:xx:xx:xx. 
  + `vif_id`:(int)(ReadOnly) Old Vif id that was associated with the interface. 
* `order`:(int) The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two. 
* `overridden_list`:
                (Array of schema.TypeString) -
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
* `pin_group_name`:(string) Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic. 
* `placement`:(HashMap) - Placement Settings for the virtual interface. 
This complex property has following sub-properties:
  + `auto_pci_link`:(bool) Enable or disable automatic assignment of the PCI Link in a dual-link adapter. This option applies only to 13xx series VICs that support dual-link. If enabled, the system determines the placement of the vNIC/vHBA on either of the PCI Links. 
  + `auto_slot_id`:(bool) Enable or disable automatic assignment of the VIC slot ID. If enabled and the server has only one VIC, the same VIC is chosen for all the vNICs. If enabled and the server has multiple VICs, the vNIC/vHBA are deployed on the first VIC. The Slot ID determines the first VIC. MLOM is the first Slot ID and the ID increments to 2, 3, and so on. 
  + `id`:(string) PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `pci_link`:(int) The PCI Link used as transport for the virtual interface. PCI Link is only applicable for select Cisco UCS VIC 1300 models (UCSC-PCIE-C40Q-03, UCSB-MLOM-40G-03, UCSB-VIC-M83-8P) that support two PCI links. The value, if specified, for any other VIC model will be ignored. 
  + `pci_link_assignment_mode`:(string) If the autoPciLink is disabled, the user can either choose to place the vNICs manually or based on a policy.If the autoPciLink is enabled, it will be set to None.* `Custom` - The user needs to specify the PCI Link manually.* `Load-Balanced` - The system will uniformly distribute the interfaces across the PCI Links.* `None` - Assignment is not applicable and will be set when the AutoPciLink is set to true. 
  + `switch_id`:(string) The fabric port to which the vNICs will be associated.* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.* `A` - Fabric A of the FI cluster.* `B` - Fabric B of the FI cluster. 
  + `uplink`:(int) Adapter port on which the virtual interface will be created. 
* `profile`:(HashMap) - A reference to a policyAbstractConfigProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sp_vnics`:(Array)(ReadOnly) An array of relationships to vnicEthIf resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `src_template`:(HashMap) - A reference to a vnicVnicTemplate resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `sriov_settings`:(HashMap) - Single Root Input Output Virtualization (SR-IOV) Settings that enable one physical ethernet socket to appear as multiple NICs to the hypervisor. 
This complex property has following sub-properties:
  + `comp_count_per_vf`:(int) Completion Queue resources per Virtual Function (VF). 
  + `enabled`:(bool) If enabled, sets Single Root Input Output Virtualization (SR-IOV) on this vNIC. 
  + `int_count_per_vf`:(int) Interrupt Count resources per Virtual Function (VF). 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `rx_count_per_vf`:(int) Receive Queue resources per Virtual Function (VF). 
  + `tx_count_per_vf`:(int) Transmit Queue resources per Virtual Function (VF). 
  + `vf_count`:(int) Number of Virtual Functions (VF) to be created for this vNIC. Valid values are 1 to 64 when SR-IOV is enabled. 
* `stale`:(bool)(ReadOnly) An EthIf is marked stale if it was deployed to the endpoint and the LAN Connectivity Policy associated with the server profile does not have this EthIf anymore. This maybe due to the LAN Connectivity Policy being removed from the server profile or a different LAN Connectivity Policy is attached which does not include any EthIf with the same name. 
* `standby_vif_id`:(int)(ReadOnly) The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path. 
* `static_mac_address`:(string) The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.To ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use thefollowing MAC prefix 00:25:B5:xx:xx:xx. 
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
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `template_actions`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `params`:(Array)
This complex property has following sub-properties:
    + `name`:(string) The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action.* `None` - The default parameter that implies that no action parameter is required for the template action.* `SyncType` - The parameter that describes the type of sync action such as SyncOne or SyncFailed supported on any template or derived object.* `SyncTimer` - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds.* `OverriddenList` - The parameter applicable in attach operation indicating the configurations that must override the template configurations. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) The action parameter value is based on the action parameter type. Supported action parameters and their values are-a) Name - SyncType, Supported Values - SyncFailed, SyncOne.b) Name - SyncTimer, Supported Values - 0 to 60 seconds.c) Name - OverriddenList, Supported Values - Comma Separated list of overridable configurations. 
  + `type`:(string) The action type to be executed.* `Sync` - The action to merge values from the template to its derived objects.* `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types.* `Detach` - The action to detach the current derived object from its attached template.* `Attach` - The action to attach the current object to the specified template. 
* `template_sync_errors`:(Array)
This complex property has following sub-properties:
  + `message`:(string)(ReadOnly) The localized message based on the locale setting of the user's context providing the error description. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `type`:(string)(ReadOnly) The error type that indicates the point of failure.* `Transient` - Any error which is a runtime error due to some other action in progress on the derived object that is blocking the sync action. This error type is retriable.For example, when vNIC Template is updated, but the derived vNIC or vNICs are part of a LAN Connectivity policy associated with a profile being deployed to endpoint. In this scenario, the derived vNIC update displays this error.* `Validation` - When the template sync on the derived object fails due to an incorrect configuration, it displays a validation error. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails due to the policy attach errors.* `User` - Any configuration error due to incorrect or invalid input and that requires user intervention for correction, is displayed under this category. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails. This happens when the policyis not applicable to the server assigned to the server profile, such as the Power policy that is not applicable for UCS Rack servers.* `Internal` - Any application internal errors are displayed under this category. This error type is considered fatal and not retried. 
* `template_sync_status`:(string)(ReadOnly) The sync status of the current MO wrt the attached Template MO.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `usnic_settings`:(HashMap) - User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets. 
This complex property has following sub-properties:
  + `cos`:(int) Class of Service to be used for traffic on the usNIC. 
  + `nr_count`:(int) Number of usNIC interfaces to be created. When usNIC is enabled, the valid values are from 1 to 225. When usNIC is disabled, the default value is 0. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `usnic_adapter_policy`:(string) Ethernet Adapter policy to be associated with the usNICs. 
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
* `vif_id`:(int)(ReadOnly) The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC. 
* `vmq_settings`:(HashMap) - Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS. 
This complex property has following sub-properties:
  + `enabled`:(bool) Enables VMQ feature on the virtual interface. 
  + `multi_queue_support`:(bool) Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM. 
  + `num_interrupts`:(int) The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server. 
  + `num_sub_vnics`:(int) The number of sub vNICs to be created. 
  + `num_vmqs`:(int) The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vmmq_adapter_policy`:(string) Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs. 


## Import
`intersight_vnic_eth_if` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_eth_if.example 1234567890987654321abcde
``` 
