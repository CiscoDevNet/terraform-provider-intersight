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

func getComputePhysicalSummarySchema() map[string]*schema.Schema {
	var schemaMap = make(map[string]*schema.Schema)
	schemaMap = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"admin_power_state": {
			Description: "The desired power state of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"alarm_summary": {
			Description: "The summary of alarm counts based on alarm serverity.",
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
					"critical": {
						Description: "The count of alarms that have severity type Critical.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"health": {
						Description: "Health of the managed end point. The highest severity computed from alarmSummary property is set as the health.\n* `Healthy` - The Enum value represents that the entity is healthy.\n* `Warning` - The Enum value Warning represents that the entity has one or more active warnings on it.\n* `Critical` - The Enum value Critical represents that the entity is in a critical state.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"info": {
						Description: "The count of alarms that have severity type Info.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"suppressed": {
						Description: "The flag that indicates whether suppression is enabled or not in the entity.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"suppressed_critical": {
						Description: "The count of active suppressed alarms that have severity type Critical.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"suppressed_info": {
						Description: "The count of active suppressed alarms that have severity type Info.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"suppressed_warning": {
						Description: "The count of active suppressed alarms that have severity type Warning.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"warning": {
						Description: "The count of alarms that have severity type Warning.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
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
		"asset_tag": {
			Description: "The user defined asset tag assigned to the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"available_memory": {
			Description: "Total memeory of the server in MB.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"bios_post_complete": {
			Description: "The BIOS POST completion status of the server.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"chassis_id": {
			Description: "The id of the chassis that the blade is discovered in.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"connection_status": {
			Description: "Connectivity Status of RackUnit to Switch - A or B or AB.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cooling_mode": {
			Description: "Cooling mode representation of the server, supported modes include Air and Immersion.\n* `Air` - Cooling mode of the device is set to Air.\n* `Immersion` - Cooling mode of the device is set to Immersion.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cpu_capacity": {
			Description: "Total processing capacity of the server.",
			Type:        schema.TypeFloat,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"custom_permission_resources": {
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
		"device_mo_id": {
			Description: "The MoId of the registered device that coresponds to the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dn": {
			Description: "The Distinguished Name unambiguously identifies an object in the system.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"equipment_chassis": {
			Description: "A reference to a equipmentChassis resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"fault_summary": {
			Description: "The fault summary for the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"firmware": {
			Description: "The firmware version of the Cisco Integrated Management Controller (CIMC) for this server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"front_panel_lock_state": {
			Description: "The actual front panel state of the server.\n* `None` - Front Panel of the server is set to None state. It is required so that the next frontPanelLockState operation can be triggered.\n* `Lock` - Front Panel of the server is set to Locked state.\n* `Unlock` - Front Panel of the server is set to Unlocked state.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"hardware_uuid": {
			Description: "The universally unique hardware identity of the server provided by the manufacturer.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"has_e3_s_support": {
			Description: "The flag to indicate server has the support for E3.S drives.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"inventory_device_info": {
			Description: "A reference to a inventoryDeviceInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"inventory_parent": {
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
		"ipv4_address": {
			Description: "The IPv4 address configured on the management interface of the Integrated Management Controller.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"is_upgraded": {
			Description: "This field indicates the compute status of the catalog values for the associated component or hardware.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"kvm_ip_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"address": {
						Description: "IP Address to be used for KVM.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"category": {
						Description: "Category of the Kvm IP Address.\n* `Equipment` - Ip Address assigned to an equipment.\n* `ServiceProfile` - Ip Address assigned to a Service Profile.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"default_gateway": {
						Description: "Default gateway property of KVM IP Address.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"dn": {
						Description: "The distinguished name for this managed object.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"http_port": {
						Description: "HTTP port of an IP Address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"https_port": {
						Description: "Secured HTTP port of an IP Address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"kvm_port": {
						Description: "Port number on which the KVM is running and used for connecting to KVM console.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"kvm_vlan": {
						Description: "VLAN Identifier of Inband IP Address.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"name": {
						Description: "Name to identify the KVM IP Address.\n* `Outband` - The user assigned Out of band IP Address.\n* `Inband` - The user assigned Inband IP Address.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"subnet": {
						Description: "Subnet detail of a KVM IP Address.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the KVM IP Address.\n* `MgmtInterface` - Ip Address of a Management Interface.\n* `VnicIpV4StaticAddr` - Static Ipv4 Address of a Virtual Network Interface.\n* `VnicIpV4PooledAddr` - Ipv4 Address of a Virtual Network Interface from an address pool.\n* `VnicIpV4MgmtPooledAddr` - Ipv4 Address of a Virtual Network Interface from a Managed address pool.\n* `VnicIpV6StaticAddr` - Static Ipv6 Address of a Virtual Network Interface.\n* `VnicIpV6MgmtPooledAddr` - Ipv6 Address of a Virtual Network Interface from a Managed address pool.\n* `VnicIpV4ProfDerivedAddr` - Server Profile derived Ipv4 Address of a Virtual Network Interface.\n* `MgmtIpV6ProfDerivedAddr` - Server Profile derived Ipv6 Address used for accessing server management services.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"kvm_server_state_enabled": {
			Description: "The KVM server state of the server.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"kvm_vendor": {
			Description: "The KVM Vendor for the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"last_power_state_changed_time": {
			Description: "The Last host power state changed time of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nr_lifecycle": {
			Description: "The lifecycle of the blade server.\n* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.\n* `Active` - Default Lifecycle State for a physical entity.\n* `Decommissioned` - Decommission Lifecycle state.\n* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.\n* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.\n* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.\n* `SecureEraseInProgress` - Secure Erase is in progress on given physical entity.\n* `ScrubInProgress` - Scrub is in progress on given physical entity.\n* `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity.\n* `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.\n* `Removed` - The blade server has been removed from its discovered slot, and not detected anywhere else. Blade inventory can be cleaned up by performing a software remove operation on the physically removed blade.\n* `Moved` - The blade server has been moved from its discovered location to a new location. Blade inventory can be updated by performing a rediscover operation on the moved blade.\n* `Replaced` - The blade server has been removed from its discovered location and another blade has been inserted in that location. Blade inventory can be cleaned up and updated by doing a software remove operation on the physically removed blade.\n* `MovedAndReplaced` - The blade server has been moved from its discovered location to a new location and another blade has been inserted into the old discovered location. Blade inventory can be updated by performing a rediscover operation on the moved blade.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"location_details": {
			Description: "Location details associated with the managed object. When a top level resource (i.e., Unified Edge Chassis) is assigned to a location, the details of the associated location are propagated to related inventory and configuration objects like servers, profiles, alarms and metrics to facilitate location based filtering.",
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
					"address": {
						Description: "The location's street address.",
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
								"address1": {
									Description: "The primary street address.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"address2": {
									Description: "Additional address information, such as suite number or floor.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"city": {
									Description: "The city where the address is located.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"country": {
									Description: "The country code in ISO 3166-1 alpha-2 format.\n* `Unknown` - The value Unknown is used when the country code is not known or applicable.\n* `AD` - The country code for Andorra.\n* `AE` - The country code for United Arab Emirates.\n* `AF` - The country code for Afghanistan.\n* `AG` - The country code for Antigua and Barbuda.\n* `AI` - The country code for Anguilla.\n* `AL` - The country code for Albania.\n* `AM` - The country code for Armenia.\n* `AO` - The country code for Angola.\n* `AQ` - The country code for Antarctica.\n* `AR` - The country code for Argentina.\n* `AS` - The country code for American Samoa.\n* `AT` - The country code for Austria.\n* `AU` - The country code for Australia.\n* `AW` - The country code for Aruba.\n* `AX` - The country code for Åland Islands.\n* `AZ` - The country code for Azerbaijan.\n* `BA` - The country code for Bosnia and Herzegovina.\n* `BB` - The country code for Barbados.\n* `BD` - The country code for Bangladesh.\n* `BE` - The country code for Belgium.\n* `BF` - The country code for Burkina Faso.\n* `BG` - The country code for Bulgaria.\n* `BH` - The country code for Bahrain.\n* `BI` - The country code for Burundi.\n* `BJ` - The country code for Benin.\n* `BL` - The country code for Saint Barthélemy.\n* `BM` - The country code for Bermuda.\n* `BN` - The country code for Brunei Darussalam.\n* `BO` - The country code for Plurinational State of Bolivia.\n* `BQ` - The country code for Sint Eustatius and Saba Bonaire.\n* `BR` - The country code for Brazil.\n* `BS` - The country code for Bahamas.\n* `BT` - The country code for Bhutan.\n* `BV` - The country code for Bouvet Island.\n* `BW` - The country code for Botswana.\n* `BZ` - The country code for Belize.\n* `CA` - The country code for Canada.\n* `CC` - The country code for Cocos (Keeling) Islands.\n* `CD` - The country code for Democratic Republic of the Congo.\n* `CF` - The country code for Central African Republic.\n* `CG` - The country code for Congo.\n* `CH` - The country code for Switzerland.\n* `CI` - The country code for Côte d'Ivoire.\n* `CK` - The country code for Cook Islands.\n* `CL` - The country code for Chile.\n* `CM` - The country code for Cameroon.\n* `CN` - The country code for China.\n* `CO` - The country code for Colombia.\n* `CR` - The country code for Costa Rica.\n* `CV` - The country code for Cabo Verde.\n* `CW` - The country code for Curaçao.\n* `CX` - The country code for Christmas Island.\n* `CY` - The country code for Cyprus.\n* `CZ` - The country code for Czechia.\n* `DE` - The country code for Germany.\n* `DJ` - The country code for Djibouti.\n* `DK` - The country code for Denmark.\n* `DM` - The country code for Dominica.\n* `DO` - The country code for Dominican Republic.\n* `DZ` - The country code for Algeria.\n* `EC` - The country code for Ecuador.\n* `EE` - The country code for Estonia.\n* `EG` - The country code for Egypt.\n* `EH` - The country code for Western Sahara.\n* `ER` - The country code for Eritrea.\n* `ES` - The country code for Spain.\n* `ET` - The country code for Ethiopia.\n* `FI` - The country code for Finland.\n* `FJ` - The country code for Fiji.\n* `FK` - The country code for Falkland Islands (Malvinas).\n* `FM` - The country code for Federated States of Micronesia.\n* `FO` - The country code for Faroe Islands.\n* `FR` - The country code for France.\n* `GA` - The country code for Gabon.\n* `GB` - The country code for United Kingdom of Great Britain and Northern Ireland.\n* `GD` - The country code for Grenada.\n* `GE` - The country code for Georgia.\n* `GF` - The country code for French Guiana.\n* `GG` - The country code for Guernsey.\n* `GH` - The country code for Ghana.\n* `GI` - The country code for Gibraltar.\n* `GL` - The country code for Greenland.\n* `GM` - The country code for Gambia.\n* `GN` - The country code for Guinea.\n* `GP` - The country code for Guadeloupe.\n* `GQ` - The country code for Equatorial Guinea.\n* `GR` - The country code for Greece.\n* `GS` - The country code for South Georgia and the South Sandwich Islands.\n* `GT` - The country code for Guatemala.\n* `GU` - The country code for Guam.\n* `GW` - The country code for Guinea-Bissau.\n* `GY` - The country code for Guyana.\n* `HK` - The country code for Hong Kong.\n* `HM` - The country code for Heard Island and McDonald Islands.\n* `HN` - The country code for Honduras.\n* `HR` - The country code for Croatia.\n* `HT` - The country code for Haiti.\n* `HU` - The country code for Hungary.\n* `ID` - The country code for Indonesia.\n* `IE` - The country code for Ireland.\n* `IL` - The country code for Israel.\n* `IM` - The country code for Isle of Man.\n* `IN` - The country code for India.\n* `IO` - The country code for British Indian Ocean Territory.\n* `IQ` - The country code for Iraq.\n* `IS` - The country code for Iceland.\n* `IT` - The country code for Italy.\n* `JE` - The country code for Jersey.\n* `JM` - The country code for Jamaica.\n* `JO` - The country code for Jordan.\n* `JP` - The country code for Japan.\n* `KE` - The country code for Kenya.\n* `KG` - The country code for Kyrgyzstan.\n* `KH` - The country code for Cambodia.\n* `KI` - The country code for Kiribati.\n* `KM` - The country code for Comoros.\n* `KN` - The country code for Saint Kitts and Nevis.\n* `KR` - The country code for Republic of Korea.\n* `KW` - The country code for Kuwait.\n* `KY` - The country code for Cayman Islands.\n* `KZ` - The country code for Kazakhstan.\n* `LA` - The country code for Lao People's Democratic Republic.\n* `LB` - The country code for Lebanon.\n* `LC` - The country code for Saint Lucia.\n* `LI` - The country code for Liechtenstein.\n* `LK` - The country code for Sri Lanka.\n* `LR` - The country code for Liberia.\n* `LS` - The country code for Lesotho.\n* `LT` - The country code for Lithuania.\n* `LU` - The country code for Luxembourg.\n* `LV` - The country code for Latvia.\n* `LY` - The country code for Libya.\n* `MA` - The country code for Morocco.\n* `MC` - The country code for Monaco.\n* `MD` - The country code for Republic of Moldova.\n* `ME` - The country code for Montenegro.\n* `MF` - The country code for Saint Martin (French part).\n* `MG` - The country code for Madagascar.\n* `MH` - The country code for Marshall Islands.\n* `MK` - The country code for North Macedonia.\n* `ML` - The country code for Mali.\n* `MM` - The country code for Myanmar.\n* `MN` - The country code for Mongolia.\n* `MO` - The country code for Macao.\n* `MP` - The country code for Northern Mariana Islands.\n* `MQ` - The country code for Martinique.\n* `MR` - The country code for Mauritania.\n* `MS` - The country code for Montserrat.\n* `MT` - The country code for Malta.\n* `MU` - The country code for Mauritius.\n* `MV` - The country code for Maldives.\n* `MW` - The country code for Malawi.\n* `MX` - The country code for Mexico.\n* `MY` - The country code for Malaysia.\n* `MZ` - The country code for Mozambique.\n* `NA` - The country code for Namibia.\n* `NC` - The country code for New Caledonia.\n* `NE` - The country code for Niger.\n* `NF` - The country code for Norfolk Island.\n* `NG` - The country code for Nigeria.\n* `NI` - The country code for Nicaragua.\n* `NL` - The country code for Kingdom of the Netherlands.\n* `NO` - The country code for Norway.\n* `NP` - The country code for Nepal.\n* `NR` - The country code for Nauru.\n* `NU` - The country code for Niue.\n* `NZ` - The country code for New Zealand.\n* `OM` - The country code for Oman.\n* `PA` - The country code for Panama.\n* `PE` - The country code for Peru.\n* `PF` - The country code for French Polynesia.\n* `PG` - The country code for Papua New Guinea.\n* `PH` - The country code for Philippines.\n* `PK` - The country code for Pakistan.\n* `PL` - The country code for Poland.\n* `PM` - The country code for Saint Pierre and Miquelon.\n* `PN` - The country code for Pitcairn.\n* `PR` - The country code for Puerto Rico.\n* `PS` - The country code for State of Palestine.\n* `PT` - The country code for Portugal.\n* `PW` - The country code for Palau.\n* `PY` - The country code for Paraguay.\n* `QA` - The country code for Qatar.\n* `RE` - The country code for Réunion.\n* `RO` - The country code for Romania.\n* `RS` - The country code for Serbia.\n* `RW` - The country code for Rwanda.\n* `SA` - The country code for Saudi Arabia.\n* `SB` - The country code for Solomon Islands.\n* `SC` - The country code for Seychelles.\n* `SD` - The country code for Sudan.\n* `SE` - The country code for Sweden.\n* `SG` - The country code for Singapore.\n* `SH` - The country code for Ascension and Tristan da Cunha Saint Helena.\n* `SI` - The country code for Slovenia.\n* `SJ` - The country code for Svalbard and Jan Mayen.\n* `SK` - The country code for Slovakia.\n* `SL` - The country code for Sierra Leone.\n* `SM` - The country code for San Marino.\n* `SN` - The country code for Senegal.\n* `SO` - The country code for Somalia.\n* `SR` - The country code for Suriname.\n* `SS` - The country code for South Sudan.\n* `ST` - The country code for Sao Tome and Principe.\n* `SV` - The country code for El Salvador.\n* `SX` - The country code for Sint Maarten (Dutch part).\n* `SZ` - The country code for Eswatini.\n* `TC` - The country code for Turks and Caicos Islands.\n* `TD` - The country code for Chad.\n* `TF` - The country code for French Southern Territories.\n* `TG` - The country code for Togo.\n* `TH` - The country code for Thailand.\n* `TJ` - The country code for Tajikistan.\n* `TK` - The country code for Tokelau.\n* `TL` - The country code for Timor-Leste.\n* `TM` - The country code for Turkmenistan.\n* `TN` - The country code for Tunisia.\n* `TO` - The country code for Tonga.\n* `TR` - The country code for Türkiye.\n* `TT` - The country code for Trinidad and Tobago.\n* `TV` - The country code for Tuvalu.\n* `TW` - The country code for Province of China Taiwan.\n* `TZ` - The country code for United Republic of Tanzania.\n* `UA` - The country code for Ukraine.\n* `UG` - The country code for Uganda.\n* `UM` - The country code for United States Minor Outlying Islands.\n* `US` - The country code for United States of America.\n* `UY` - The country code for Uruguay.\n* `UZ` - The country code for Uzbekistan.\n* `VA` - The country code for Holy See.\n* `VC` - The country code for Saint Vincent and the Grenadines.\n* `VE` - The country code for Bolivarian Republic of Venezuela.\n* `VG` - The country code for Virgin Islands (British).\n* `VI` - The country code for Virgin Islands (U.S.).\n* `VN` - The country code for Viet Nam.\n* `VU` - The country code for Vanuatu.\n* `WF` - The country code for Wallis and Futuna.\n* `WS` - The country code for Samoa.\n* `YE` - The country code for Yemen.\n* `YT` - The country code for Mayotte.\n* `ZA` - The country code for South Africa.\n* `ZM` - The country code for Zambia.\n* `ZW` - The country code for Zimbabwe.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"postal_code": {
									Description: "The postal or ZIP code for the address.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"state_province": {
									Description: "The state or province where the address is located.",
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
					"coordinates": {
						Description: "The location's longitude and latitude coordinates.",
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
								"latitude": {
									Description: "The latitude coordinate value.",
									Type:        schema.TypeFloat,
									Optional:    true,
								},
								"longitude": {
									Description: "The longitude coordinate value.",
									Type:        schema.TypeFloat,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"name": {
						Description: "A user provided name for the location.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"management_mode": {
			Description: "The management mode of the server.\n* `IntersightStandalone` - Intersight Standalone mode of operation.\n* `UCSM` - Unified Computing System Manager mode of operation.\n* `Intersight` - Intersight managed mode of operation.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"memory_speed": {
			Description: "The maximum memory speed in MHz available on the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mgmt_ip_address": {
			Description: "Management address of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"model": {
			Description: "This field identifies the model of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "The name of the UCS Fabric Interconnect cluster or Cisco Integrated Management Controller (CIMC). When this server is attached to a UCS Fabric Interconnect, the value of this property is the name of the UCS Fabric Interconnect along with chassis/server Id. When this server configured in standalone mode, the value of this property is the name of the Cisco Integrated Management Controller. when this server is configired in IMM mode, the value of this property contains model and chassis/server Id.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"num_adaptors": {
			Description: "The total number of network adapters present on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_cpu_cores": {
			Description: "The total number of CPU cores enabled on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_cpu_cores_enabled": {
			Description: "The total number of CPU cores enabled on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_cpus": {
			Description: "The total number of CPUs present on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_eth_host_interfaces": {
			Description: "The total number of vNICs which are visible to a host on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_fc_host_interfaces": {
			Description: "The total number of vHBAs which are visible to a host on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"num_threads": {
			Description: "The total number of threads the server is capable of handling.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_power_state": {
			Description: "The actual power state of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_reason": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"oper_state": {
			Description: "The operational state of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"operability": {
			Description: "The operability of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"package_version": {
			Description: "Bundle version which the firmware belongs to.",
			Type:        schema.TypeString,
			Optional:    true,
		},
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
		"personality": {
			Description: "Unique identity of added software personality.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"platform_type": {
			Description: "The platform type of the registered device - whether managed by UCSM or operating in standalone mode.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"presence": {
			Description: "This field identifies the presence (equipped) or absence of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"registered_device": {
			Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		"revision": {
			Description: "This field identifies the revision of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"rn": {
			Description: "The Relative Name uniquely identifies an object within a given context.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"scaled_mode": {
			Description: "The mode of the server that determines it is scaled.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"serial": {
			Description: "This field identifies the serial of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"server_id": {
			Description: "RackUnit ID that uniquely identifies the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"service_profile": {
			Description: "The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"slot_id": {
			Description: "The slot number in the chassis that the blade is discovered in.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"source_object_type": {
			Description: "Stores the source object type. This feild will either be RackUnit or Blade.",
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
					"ancestor_definitions": {
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
					"definition": {
						Description: "The definition is a reference to the tag definition object.\nThe tag definition object contains the properties of the tag such as name, type, and description.",
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
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"propagated": {
						Description: "Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"sys_tag": {
						Description: "Specifies whether the tag is user-defined or owned by the system.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"type": {
						Description: "An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.\n* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.\n* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".",
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
		"topology_scan_status": {
			Description: "To maintain the Topology workflow run status.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"total_memory": {
			Description: "The total memory available on the server.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tunneled_kvm": {
			Description: "The Tunneled vKVM status of the server.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"user_label": {
			Description: "The user defined label assigned to the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"uuid": {
			Description: "The universally unique identity of the server.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vendor": {
			Description: "This field identifies the vendor of the given component.",
			Type:        schema.TypeString,
			Optional:    true,
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
					"marked_for_deletion": {
						Description: "The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy.",
						Type:        schema.TypeBool,
						Optional:    true,
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
	}
	return schemaMap
}

func dataSourceComputePhysicalSummary() *schema.Resource {
	var subSchema = getComputePhysicalSummarySchema()
	var model = getComputePhysicalSummarySchema()
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceComputePhysicalSummaryRead,
		Schema:      model}
}

func dataSourceComputePhysicalSummaryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ComputePhysicalSummary{}
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

	if v, ok := d.GetOk("admin_power_state"); ok {
		x := (v.(string))
		o.SetAdminPowerState(x)
	}

	if v, ok := d.GetOk("alarm_summary"); ok {
		p := make([]models.ComputeAlarmSummary, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.ComputeAlarmSummary{}
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
			o.SetClassId("compute.AlarmSummary")
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
			o.SetAlarmSummary(x)
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

	if v, ok := d.GetOk("asset_tag"); ok {
		x := (v.(string))
		o.SetAssetTag(x)
	}

	if v, ok := d.GetOkExists("available_memory"); ok {
		x := int64(v.(int))
		o.SetAvailableMemory(x)
	}

	if v, ok := d.GetOkExists("bios_post_complete"); ok {
		x := (v.(bool))
		o.SetBiosPostComplete(x)
	}

	if v, ok := d.GetOk("chassis_id"); ok {
		x := (v.(string))
		o.SetChassisId(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}

	if v, ok := d.GetOk("cooling_mode"); ok {
		x := (v.(string))
		o.SetCoolingMode(x)
	}

	if v, ok := d.GetOk("cpu_capacity"); ok {
		x := float32(v.(float64))
		o.SetCpuCapacity(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("custom_permission_resources"); ok {
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
		o.SetCustomPermissionResources(x)
	}

	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.SetDeviceMoId(x)
	}

	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("equipment_chassis"); ok {
		p := make([]models.EquipmentChassisRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsEquipmentChassisRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetEquipmentChassis(x)
		}
	}

	if v, ok := d.GetOkExists("fault_summary"); ok {
		x := int64(v.(int))
		o.SetFaultSummary(x)
	}

	if v, ok := d.GetOk("firmware"); ok {
		x := (v.(string))
		o.SetFirmware(x)
	}

	if v, ok := d.GetOk("front_panel_lock_state"); ok {
		x := (v.(string))
		o.SetFrontPanelLockState(x)
	}

	if v, ok := d.GetOk("hardware_uuid"); ok {
		x := (v.(string))
		o.SetHardwareUuid(x)
	}

	if v, ok := d.GetOkExists("has_e3_s_support"); ok {
		x := (v.(bool))
		o.SetHasE3SSupport(x)
	}

	if v, ok := d.GetOk("inventory_device_info"); ok {
		p := make([]models.InventoryDeviceInfoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsInventoryDeviceInfoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetInventoryDeviceInfo(x)
		}
	}

	if v, ok := d.GetOk("inventory_parent"); ok {
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
			o.SetInventoryParent(x)
		}
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		x := (v.(string))
		o.SetIpv4Address(x)
	}

	if v, ok := d.GetOkExists("is_upgraded"); ok {
		x := (v.(bool))
		o.SetIsUpgraded(x)
	}

	if v, ok := d.GetOk("kvm_ip_addresses"); ok {
		x := make([]models.ComputeIpAddress, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.ComputeIpAddress{}
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
			o.SetClassId("compute.IpAddress")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetKvmIpAddresses(x)
	}

	if v, ok := d.GetOkExists("kvm_server_state_enabled"); ok {
		x := (v.(bool))
		o.SetKvmServerStateEnabled(x)
	}

	if v, ok := d.GetOk("kvm_vendor"); ok {
		x := (v.(string))
		o.SetKvmVendor(x)
	}

	if v, ok := d.GetOk("last_power_state_changed_time"); ok {
		x := (v.(string))
		o.SetLastPowerStateChangedTime(x)
	}

	if v, ok := d.GetOk("nr_lifecycle"); ok {
		x := (v.(string))
		o.SetLifecycle(x)
	}

	if v, ok := d.GetOk("location_details"); ok {
		p := make([]models.CommGeoLocationDetails, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CommGeoLocationDetails{}
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
			o.SetClassId("comm.GeoLocationDetails")
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
			o.SetLocationDetails(x)
		}
	}

	if v, ok := d.GetOk("management_mode"); ok {
		x := (v.(string))
		o.SetManagementMode(x)
	}

	if v, ok := d.GetOk("memory_speed"); ok {
		x := (v.(string))
		o.SetMemorySpeed(x)
	}

	if v, ok := d.GetOk("mgmt_ip_address"); ok {
		x := (v.(string))
		o.SetMgmtIpAddress(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	if v, ok := d.GetOkExists("num_adaptors"); ok {
		x := int64(v.(int))
		o.SetNumAdaptors(x)
	}

	if v, ok := d.GetOkExists("num_cpu_cores"); ok {
		x := int64(v.(int))
		o.SetNumCpuCores(x)
	}

	if v, ok := d.GetOkExists("num_cpu_cores_enabled"); ok {
		x := int64(v.(int))
		o.SetNumCpuCoresEnabled(x)
	}

	if v, ok := d.GetOkExists("num_cpus"); ok {
		x := int64(v.(int))
		o.SetNumCpus(x)
	}

	if v, ok := d.GetOkExists("num_eth_host_interfaces"); ok {
		x := int64(v.(int))
		o.SetNumEthHostInterfaces(x)
	}

	if v, ok := d.GetOkExists("num_fc_host_interfaces"); ok {
		x := int64(v.(int))
		o.SetNumFcHostInterfaces(x)
	}

	if v, ok := d.GetOkExists("num_threads"); ok {
		x := int64(v.(int))
		o.SetNumThreads(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("oper_power_state"); ok {
		x := (v.(string))
		o.SetOperPowerState(x)
	}

	if v, ok := d.GetOk("oper_reason"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOperReason(x)
	}

	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}

	if v, ok := d.GetOk("operability"); ok {
		x := (v.(string))
		o.SetOperability(x)
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

	if v, ok := d.GetOk("package_version"); ok {
		x := (v.(string))
		o.SetPackageVersion(x)
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

	if v, ok := d.GetOk("personality"); ok {
		x := (v.(string))
		o.SetPersonality(x)
	}

	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}

	if v, ok := d.GetOk("presence"); ok {
		x := (v.(string))
		o.SetPresence(x)
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}

	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.SetRn(x)
	}

	if v, ok := d.GetOk("scaled_mode"); ok {
		x := (v.(string))
		o.SetScaledMode(x)
	}

	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}

	if v, ok := d.GetOkExists("server_id"); ok {
		x := int64(v.(int))
		o.SetServerId(x)
	}

	if v, ok := d.GetOk("service_profile"); ok {
		x := (v.(string))
		o.SetServiceProfile(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOkExists("slot_id"); ok {
		x := int64(v.(int))
		o.SetSlotId(x)
	}

	if v, ok := d.GetOk("source_object_type"); ok {
		x := (v.(string))
		o.SetSourceObjectType(x)
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
			if v, ok := l["ancestor_definitions"]; ok {
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
						o.SetAncestorDefinitions(x)
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

	if v, ok := d.GetOk("topology_scan_status"); ok {
		x := (v.(string))
		o.SetTopologyScanStatus(x)
	}

	if v, ok := d.GetOkExists("total_memory"); ok {
		x := int64(v.(int))
		o.SetTotalMemory(x)
	}

	if v, ok := d.GetOkExists("tunneled_kvm"); ok {
		x := (v.(bool))
		o.SetTunneledKvm(x)
	}

	if v, ok := d.GetOk("user_label"); ok {
		x := (v.(string))
		o.SetUserLabel(x)
	}

	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ComputePhysicalSummary object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ComputeApi.GetComputePhysicalSummaryList(conn.ctx).Filter(getRequestParams(data)).Count(true).Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of ComputePhysicalSummary: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of ComputePhysicalSummary: %s", responseErr.Error())
	}
	count := countResponse.MoDocumentCount.GetCount()
	if count == 0 {
		return diag.Errorf("your query for ComputePhysicalSummary data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var computePhysicalSummaryResults = make([]map[string]interface{}, 0, 0)
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ComputeApi.GetComputePhysicalSummaryList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(*models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching ComputePhysicalSummary: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching ComputePhysicalSummary: %s", responseErr.Error())
		}
		results := resMo.ComputePhysicalSummaryList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for k := 0; k < len(results); k++ {
				var s = results[k]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_power_state"] = (s.GetAdminPowerState())

				temp["alarm_summary"] = flattenMapComputeAlarmSummary(s.GetAlarmSummary(), d)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["asset_tag"] = (s.GetAssetTag())
				temp["available_memory"] = (s.GetAvailableMemory())
				temp["bios_post_complete"] = (s.GetBiosPostComplete())
				temp["chassis_id"] = (s.GetChassisId())
				temp["class_id"] = (s.GetClassId())
				temp["connection_status"] = (s.GetConnectionStatus())
				temp["cooling_mode"] = (s.GetCoolingMode())
				temp["cpu_capacity"] = (s.GetCpuCapacity())

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["custom_permission_resources"] = flattenListMoBaseMoRelationship(s.GetCustomPermissionResources(), d)
				temp["device_mo_id"] = (s.GetDeviceMoId())
				temp["dn"] = (s.GetDn())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["equipment_chassis"] = flattenMapEquipmentChassisRelationship(s.GetEquipmentChassis(), d)
				temp["fault_summary"] = (s.GetFaultSummary())
				temp["firmware"] = (s.GetFirmware())
				temp["front_panel_lock_state"] = (s.GetFrontPanelLockState())
				temp["hardware_uuid"] = (s.GetHardwareUuid())
				temp["has_e3_s_support"] = (s.GetHasE3SSupport())

				temp["inventory_device_info"] = flattenMapInventoryDeviceInfoRelationship(s.GetInventoryDeviceInfo(), d)

				temp["inventory_parent"] = flattenMapMoBaseMoRelationship(s.GetInventoryParent(), d)
				temp["ipv4_address"] = (s.GetIpv4Address())
				temp["is_upgraded"] = (s.GetIsUpgraded())

				temp["kvm_ip_addresses"] = flattenListComputeIpAddress(s.GetKvmIpAddresses(), d)
				temp["kvm_server_state_enabled"] = (s.GetKvmServerStateEnabled())
				temp["kvm_vendor"] = (s.GetKvmVendor())
				temp["last_power_state_changed_time"] = (s.GetLastPowerStateChangedTime())
				temp["nr_lifecycle"] = (s.GetLifecycle())

				temp["location_details"] = flattenMapCommGeoLocationDetails(s.GetLocationDetails(), d)
				temp["management_mode"] = (s.GetManagementMode())
				temp["memory_speed"] = (s.GetMemorySpeed())
				temp["mgmt_ip_address"] = (s.GetMgmtIpAddress())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["num_adaptors"] = (s.GetNumAdaptors())
				temp["num_cpu_cores"] = (s.GetNumCpuCores())
				temp["num_cpu_cores_enabled"] = (s.GetNumCpuCoresEnabled())
				temp["num_cpus"] = (s.GetNumCpus())
				temp["num_eth_host_interfaces"] = (s.GetNumEthHostInterfaces())
				temp["num_fc_host_interfaces"] = (s.GetNumFcHostInterfaces())
				temp["num_threads"] = (s.GetNumThreads())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_power_state"] = (s.GetOperPowerState())
				temp["oper_reason"] = (s.GetOperReason())
				temp["oper_state"] = (s.GetOperState())
				temp["operability"] = (s.GetOperability())
				temp["owners"] = (s.GetOwners())
				temp["package_version"] = (s.GetPackageVersion())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["personality"] = (s.GetPersonality())
				temp["platform_type"] = (s.GetPlatformType())
				temp["presence"] = (s.GetPresence())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["revision"] = (s.GetRevision())
				temp["rn"] = (s.GetRn())
				temp["scaled_mode"] = (s.GetScaledMode())
				temp["serial"] = (s.GetSerial())
				temp["server_id"] = (s.GetServerId())
				temp["service_profile"] = (s.GetServiceProfile())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["slot_id"] = (s.GetSlotId())
				temp["source_object_type"] = (s.GetSourceObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["topology_scan_status"] = (s.GetTopologyScanStatus())
				temp["total_memory"] = (s.GetTotalMemory())
				temp["tunneled_kvm"] = (s.GetTunneledKvm())
				temp["user_label"] = (s.GetUserLabel())
				temp["uuid"] = (s.GetUuid())
				temp["vendor"] = (s.GetVendor())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				computePhysicalSummaryResults = append(computePhysicalSummaryResults, temp)
			}
		}
	}
	log.Println("length of results: ", len(computePhysicalSummaryResults))
	if err := d.Set("results", computePhysicalSummaryResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(computePhysicalSummaryResults[0]["moid"].(string))
	return de
}
