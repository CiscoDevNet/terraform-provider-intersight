package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicEthIfInventory() *schema.Resource {
	var subSchema = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"cdn": {
			Description: "Consistent Device Naming configuration for the virtual NIC.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_source": {
						Description: "Source of the CDN. It can either be user specified or be the same as the vNIC name.\n* `vnic` - Source of the CDN is the same as the vNIC name.\n* `user` - Source of the CDN is specified by the user.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The CDN value entered in case of user defined mode.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_mo_id": {
			Description: "Device ID of the entity from where inventory is reported.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"eth_adapter_policy": {
			Description: "A reference to a vnicEthAdapterPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"eth_network_policy": {
			Description: "A reference to a vnicEthNetworkPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"eth_qos_policy": {
			Description: "A reference to a vnicEthQosPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fabric_eth_network_control_policy": {
			Description: "A reference to a fabricEthNetworkControlPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fabric_eth_network_group_policy": {
			Description: "An array of relationships to fabricEthNetworkGroupPolicyInventory resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"failover_enabled": {
			Description: "Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"ip_lease": {
			Description: "A reference to a ippoolIpLease resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_boot_policy": {
			Description: "A reference to a vnicIscsiBootPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_ip_v4_address_allocation_type": {
			Description: "Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.\n* `None` - Type indicates that there is no IP associated to an vnic.\n* `DHCP` - The IP address is assigned using DHCP, if available.\n* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.\n* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"iscsi_ip_v4_config": {
			Description: "IPV4 configurations such as Netmask, Gateway and DNS for iSCSI vNIC.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"gateway": {
						Description: "IP address of the default IPv4 gateway.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"netmask": {
						Description: "A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"primary_dns": {
						Description: "IP Address of the primary Domain Name System (DNS) server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"secondary_dns": {
						Description: "IP Address of the secondary Domain Name System (DNS) server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_ipv4_address": {
			Description: "IP address associated to the vNIC.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"lan_connectivity_policy": {
			Description: "A reference to a vnicLanConnectivityPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"lcp_vnic": {
			Description: "A reference to a vnicEthIfInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mac_address": {
			Description: "The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_address_type": {
			Description: "Type of allocation selected to assign a MAC address for the vnic.\n* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.\n* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_lease": {
			Description: "A reference to a macpoolLease resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mac_pool": {
			Description: "A reference to a macpoolPool resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "Name of the virtual ethernet interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"order": {
			Description: "The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"pin_group_name": {
			Description: "Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"placement": {
			Description: "Placement Settings for the virtual interface.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"id": {
						Description: "PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"pci_link": {
						Description: "The PCI Link used as transport for the virtual interface. This field is applicable only for VIC 1385 model (UCSC-PCIE-C40Q-03) which support two PCI links. The value, if specified, for any other VIC model will be ignored.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"switch_id": {
						Description: "The fabric port to which the vNICs will be associated.\n* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.\n* `A` - Fabric A of the FI cluster.\n* `B` - Fabric B of the FI cluster.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"uplink": {
						Description: "Adapter port on which the virtual interface will be created.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sp_vnics": {
			Description: "An array of relationships to vnicEthIfInventory resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"standby_vif_id": {
			Description: "The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"static_mac_address": {
			Description: "The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.\nTo ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the\nfollowing MAC prefix 00:25:B5:xx:xx:xx.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"usnic_settings": {
			Description: "User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cos": {
						Description: "Class of Service to be used for traffic on the usNIC.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"nr_count": {
						Description: "Number of usNIC interfaces to be created. Valid values are 1 to 255 when the usNIC is enabled. When the usNIC is disabled, the default value is 0.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"usnic_adapter_policy": {
						Description: "Ethernet Adapter policy to be associated with the usNICs.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vif_id": {
			Description: "The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmq_settings": {
			Description: "Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"enabled": {
						Description: "Enables VMQ feature on the virtual interface.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"multi_queue_support": {
						Description: "Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"num_interrupts": {
						Description: "The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_sub_vnics": {
						Description: "The number of sub vNICs to be created.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_vmqs": {
						Description: "The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vmmq_adapter_policy": {
						Description: "Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	var model = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"cdn": {
			Description: "Consistent Device Naming configuration for the virtual NIC.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_source": {
						Description: "Source of the CDN. It can either be user specified or be the same as the vNIC name.\n* `vnic` - Source of the CDN is the same as the vNIC name.\n* `user` - Source of the CDN is specified by the user.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The CDN value entered in case of user defined mode.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"device_mo_id": {
			Description: "Device ID of the entity from where inventory is reported.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"eth_adapter_policy": {
			Description: "A reference to a vnicEthAdapterPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"eth_network_policy": {
			Description: "A reference to a vnicEthNetworkPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"eth_qos_policy": {
			Description: "A reference to a vnicEthQosPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fabric_eth_network_control_policy": {
			Description: "A reference to a fabricEthNetworkControlPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fabric_eth_network_group_policy": {
			Description: "An array of relationships to fabricEthNetworkGroupPolicyInventory resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"failover_enabled": {
			Description: "Enabling failover ensures that traffic from the vNIC automatically fails over to the secondary Fabric Interconnect, in case the specified Fabric Interconnect path goes down. Failover applies only to Cisco VICs that are connected to a Fabric Interconnect cluster.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"ip_lease": {
			Description: "A reference to a ippoolIpLease resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_boot_policy": {
			Description: "A reference to a vnicIscsiBootPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_ip_v4_address_allocation_type": {
			Description: "Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.\n* `None` - Type indicates that there is no IP associated to an vnic.\n* `DHCP` - The IP address is assigned using DHCP, if available.\n* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.\n* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"iscsi_ip_v4_config": {
			Description: "IPV4 configurations such as Netmask, Gateway and DNS for iSCSI vNIC.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"gateway": {
						Description: "IP address of the default IPv4 gateway.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"netmask": {
						Description: "A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"primary_dns": {
						Description: "IP Address of the primary Domain Name System (DNS) server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"secondary_dns": {
						Description: "IP Address of the secondary Domain Name System (DNS) server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"iscsi_ipv4_address": {
			Description: "IP address associated to the vNIC.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"lan_connectivity_policy": {
			Description: "A reference to a vnicLanConnectivityPolicyInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"lcp_vnic": {
			Description: "A reference to a vnicEthIfInventory resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mac_address": {
			Description: "The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_address_type": {
			Description: "Type of allocation selected to assign a MAC address for the vnic.\n* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.\n* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mac_lease": {
			Description: "A reference to a macpoolLease resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mac_pool": {
			Description: "A reference to a macpoolPool resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "Name of the virtual ethernet interface.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"order": {
			Description: "The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"pin_group_name": {
			Description: "Pingroup name associated to vNIC for static pinning. LCP deploy will resolve pingroup name and fetches the correspoding uplink port/port channel to pin the vNIC traffic.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"placement": {
			Description: "Placement Settings for the virtual interface.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"id": {
						Description: "PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"pci_link": {
						Description: "The PCI Link used as transport for the virtual interface. This field is applicable only for VIC 1385 model (UCSC-PCIE-C40Q-03) which support two PCI links. The value, if specified, for any other VIC model will be ignored.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"switch_id": {
						Description: "The fabric port to which the vNICs will be associated.\n* `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.\n* `A` - Fabric A of the FI cluster.\n* `B` - Fabric B of the FI cluster.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"uplink": {
						Description: "Adapter port on which the virtual interface will be created.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sp_vnics": {
			Description: "An array of relationships to vnicEthIfInventory resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"standby_vif_id": {
			Description: "The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"static_mac_address": {
			Description: "The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.\nTo ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the\nfollowing MAC prefix 00:25:B5:xx:xx:xx.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"usnic_settings": {
			Description: "User Space NIC Settings that enable low-latency and higher throughput by bypassing the kernel layer when sending/receiving packets.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cos": {
						Description: "Class of Service to be used for traffic on the usNIC.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"nr_count": {
						Description: "Number of usNIC interfaces to be created. Valid values are 1 to 255 when the usNIC is enabled. When the usNIC is disabled, the default value is 0.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"usnic_adapter_policy": {
						Description: "Ethernet Adapter policy to be associated with the usNICs.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vif_id": {
			Description: "The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vmq_settings": {
			Description: "Virtual Machine Queue Settings for the virtual interface that allow efficient transfer of network traffic to the guest OS.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"enabled": {
						Description: "Enables VMQ feature on the virtual interface.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"multi_queue_support": {
						Description: "Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"num_interrupts": {
						Description: "The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_sub_vnics": {
						Description: "The number of sub vNICs to be created.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"num_vmqs": {
						Description: "The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vmmq_adapter_policy": {
						Description: "Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceVnicEthIfInventoryRead,
		Schema:      model}
}

func dataSourceVnicEthIfInventoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicEthIfInventory{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("cdn"); ok {
		p := make([]models.VnicCdn, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VnicCdn{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("vnic.Cdn")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["nr_source"]; ok {
				{
					x := (v.(string))
					o.SetSource(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCdn(x)
		}
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("eth_adapter_policy"); ok {
		p := make([]models.VnicEthAdapterPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicEthAdapterPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetEthAdapterPolicy(x)
		}
	}

	if v, ok := d.GetOk("eth_network_policy"); ok {
		p := make([]models.VnicEthNetworkPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicEthNetworkPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetEthNetworkPolicy(x)
		}
	}

	if v, ok := d.GetOk("eth_qos_policy"); ok {
		p := make([]models.VnicEthQosPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicEthQosPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetEthQosPolicy(x)
		}
	}

	if v, ok := d.GetOk("fabric_eth_network_control_policy"); ok {
		p := make([]models.FabricEthNetworkControlPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsFabricEthNetworkControlPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetFabricEthNetworkControlPolicy(x)
		}
	}

	if v, ok := d.GetOk("fabric_eth_network_group_policy"); ok {
		x := make([]models.FabricEthNetworkGroupPolicyInventoryRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsFabricEthNetworkGroupPolicyInventoryRelationship(o))
		}
		o.SetFabricEthNetworkGroupPolicy(x)
	}

	if v, ok := d.GetOkExists("failover_enabled"); ok {
		x := (v.(bool))
		o.SetFailoverEnabled(x)
	}

	if v, ok := d.GetOk("ip_lease"); ok {
		p := make([]models.IppoolIpLeaseRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIppoolIpLeaseRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetIpLease(x)
		}
	}

	if v, ok := d.GetOk("iscsi_boot_policy"); ok {
		p := make([]models.VnicIscsiBootPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicIscsiBootPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetIscsiBootPolicy(x)
		}
	}

	if v, ok := d.GetOk("iscsi_ip_v4_address_allocation_type"); ok {
		x := (v.(string))
		o.SetIscsiIpV4AddressAllocationType(x)
	}

	if v, ok := d.GetOk("iscsi_ip_v4_config"); ok {
		p := make([]models.IppoolIpV4Config, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.IppoolIpV4Config{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("ippool.IpV4Config")
			if v, ok := l["gateway"]; ok {
				{
					x := (v.(string))
					o.SetGateway(x)
				}
			}
			if v, ok := l["netmask"]; ok {
				{
					x := (v.(string))
					o.SetNetmask(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["primary_dns"]; ok {
				{
					x := (v.(string))
					o.SetPrimaryDns(x)
				}
			}
			if v, ok := l["secondary_dns"]; ok {
				{
					x := (v.(string))
					o.SetSecondaryDns(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetIscsiIpV4Config(x)
		}
	}

	if v, ok := d.GetOk("iscsi_ipv4_address"); ok {
		x := (v.(string))
		o.SetIscsiIpv4Address(x)
	}

	if v, ok := d.GetOk("lan_connectivity_policy"); ok {
		p := make([]models.VnicLanConnectivityPolicyInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicLanConnectivityPolicyInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetLanConnectivityPolicy(x)
		}
	}

	if v, ok := d.GetOk("lcp_vnic"); ok {
		p := make([]models.VnicEthIfInventoryRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsVnicEthIfInventoryRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetLcpVnic(x)
		}
	}

	if v, ok := d.GetOk("mac_address"); ok {
		x := (v.(string))
		o.SetMacAddress(x)
	}

	if v, ok := d.GetOk("mac_address_type"); ok {
		x := (v.(string))
		o.SetMacAddressType(x)
	}

	if v, ok := d.GetOk("mac_lease"); ok {
		p := make([]models.MacpoolLeaseRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMacpoolLeaseRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetMacLease(x)
		}
	}

	if v, ok := d.GetOk("mac_pool"); ok {
		p := make([]models.MacpoolPoolRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMacpoolPoolRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetMacPool(x)
		}
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOkExists("order"); ok {
		x := int64(v.(int))
		o.SetOrder(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("pin_group_name"); ok {
		x := (v.(string))
		o.SetPinGroupName(x)
	}

	if v, ok := d.GetOk("placement"); ok {
		p := make([]models.VnicPlacementSettings, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VnicPlacementSettings{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("vnic.PlacementSettings")
			if v, ok := l["id"]; ok {
				{
					x := (v.(string))
					o.SetId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["pci_link"]; ok {
				{
					x := int64(v.(int))
					o.SetPciLink(x)
				}
			}
			if v, ok := l["switch_id"]; ok {
				{
					x := (v.(string))
					o.SetSwitchId(x)
				}
			}
			if v, ok := l["uplink"]; ok {
				{
					x := int64(v.(int))
					o.SetUplink(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPlacement(x)
		}
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("sp_vnics"); ok {
		x := make([]models.VnicEthIfInventoryRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsVnicEthIfInventoryRelationship(o))
		}
		o.SetSpVnics(x)
	}

	if v, ok := d.GetOkExists("standby_vif_id"); ok {
		x := int64(v.(int))
		o.SetStandbyVifId(x)
	}

	if v, ok := d.GetOk("static_mac_address"); ok {
		x := (v.(string))
		o.SetStaticMacAddress(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOk("usnic_settings"); ok {
		p := make([]models.VnicUsnicSettings, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VnicUsnicSettings{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("vnic.UsnicSettings")
			if v, ok := l["cos"]; ok {
				{
					x := int64(v.(int))
					o.SetCos(x)
				}
			}
			if v, ok := l["nr_count"]; ok {
				{
					x := int64(v.(int))
					o.SetCount(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["usnic_adapter_policy"]; ok {
				{
					x := (v.(string))
					o.SetUsnicAdapterPolicy(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetUsnicSettings(x)
		}
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if v, ok := d.GetOkExists("vif_id"); ok {
		x := int64(v.(int))
		o.SetVifId(x)
	}

	if v, ok := d.GetOk("vmq_settings"); ok {
		p := make([]models.VnicVmqSettings, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.VnicVmqSettings{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("vnic.VmqSettings")
			if v, ok := l["enabled"]; ok {
				{
					x := (v.(bool))
					o.SetEnabled(x)
				}
			}
			if v, ok := l["multi_queue_support"]; ok {
				{
					x := (v.(bool))
					o.SetMultiQueueSupport(x)
				}
			}
			if v, ok := l["num_interrupts"]; ok {
				{
					x := int64(v.(int))
					o.SetNumInterrupts(x)
				}
			}
			if v, ok := l["num_sub_vnics"]; ok {
				{
					x := int64(v.(int))
					o.SetNumSubVnics(x)
				}
			}
			if v, ok := l["num_vmqs"]; ok {
				{
					x := int64(v.(int))
					o.SetNumVmqs(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["vmmq_adapter_policy"]; ok {
				{
					x := (v.(string))
					o.SetVmmqAdapterPolicy(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmqSettings(x)
		}
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicEthIfInventory object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthIfInventoryList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of VnicEthIfInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of VnicEthIfInventory: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for VnicEthIfInventory data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var vnicEthIfInventoryResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthIfInventoryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching VnicEthIfInventory: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching VnicEthIfInventory: %s", responseErr.Error())
		}
		results := resMo.VnicEthIfInventoryList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["cdn"] = flattenMapVnicCdn(s.GetCdn(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["eth_adapter_policy"] = flattenMapVnicEthAdapterPolicyInventoryRelationship(s.GetEthAdapterPolicy(), d)

				temp["eth_network_policy"] = flattenMapVnicEthNetworkPolicyInventoryRelationship(s.GetEthNetworkPolicy(), d)

				temp["eth_qos_policy"] = flattenMapVnicEthQosPolicyInventoryRelationship(s.GetEthQosPolicy(), d)

				temp["fabric_eth_network_control_policy"] = flattenMapFabricEthNetworkControlPolicyInventoryRelationship(s.GetFabricEthNetworkControlPolicy(), d)

				temp["fabric_eth_network_group_policy"] = flattenListFabricEthNetworkGroupPolicyInventoryRelationship(s.GetFabricEthNetworkGroupPolicy(), d)
				temp["failover_enabled"] = (s.GetFailoverEnabled())

				temp["ip_lease"] = flattenMapIppoolIpLeaseRelationship(s.GetIpLease(), d)

				temp["iscsi_boot_policy"] = flattenMapVnicIscsiBootPolicyInventoryRelationship(s.GetIscsiBootPolicy(), d)
				temp["iscsi_ip_v4_address_allocation_type"] = (s.GetIscsiIpV4AddressAllocationType())

				temp["iscsi_ip_v4_config"] = flattenMapIppoolIpV4Config(s.GetIscsiIpV4Config(), d)
				temp["iscsi_ipv4_address"] = (s.GetIscsiIpv4Address())

				temp["lan_connectivity_policy"] = flattenMapVnicLanConnectivityPolicyInventoryRelationship(s.GetLanConnectivityPolicy(), d)

				temp["lcp_vnic"] = flattenMapVnicEthIfInventoryRelationship(s.GetLcpVnic(), d)
				temp["mac_address"] = (s.GetMacAddress())
				temp["mac_address_type"] = (s.GetMacAddressType())

				temp["mac_lease"] = flattenMapMacpoolLeaseRelationship(s.GetMacLease(), d)

				temp["mac_pool"] = flattenMapMacpoolPoolRelationship(s.GetMacPool(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["order"] = (s.GetOrder())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["pin_group_name"] = (s.GetPinGroupName())

				temp["placement"] = flattenMapVnicPlacementSettings(s.GetPlacement(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["sp_vnics"] = flattenListVnicEthIfInventoryRelationship(s.GetSpVnics(), d)
				temp["standby_vif_id"] = (s.GetStandbyVifId())
				temp["static_mac_address"] = (s.GetStaticMacAddress())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["usnic_settings"] = flattenMapVnicUsnicSettings(s.GetUsnicSettings(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vif_id"] = (s.GetVifId())

				temp["vmq_settings"] = flattenMapVnicVmqSettings(s.GetVmqSettings(), d)
				vnicEthIfInventoryResults = append(vnicEthIfInventoryResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(vnicEthIfInventoryResults))
	if err := d.Set("results", vnicEthIfInventoryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicEthIfInventoryResults[0]["moid"].(string))
	return de
}
