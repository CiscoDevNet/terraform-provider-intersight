package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCapabilityAdapterUnitDescriptor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCapabilityAdapterUnitDescriptorCreate,
		ReadContext:   resourceCapabilityAdapterUnitDescriptorRead,
		UpdateContext: resourceCapabilityAdapterUnitDescriptorUpdate,
		DeleteContext: resourceCapabilityAdapterUnitDescriptorDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CombinedCustomizeDiff,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"adapter_generation": {
				Description:  "Generation of the adapter.\n* `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04.\n* `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02.\n* `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03.\n* `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{4, 2, 3, 5}),
				Optional:     true,
				Default:      4,
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
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"capabilities": {
				Description: "An array of relationships to capabilityCapability resources.",
				Type:        schema.TypeList,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Computed:    true,
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
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
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
				Default:     "capability.AdapterUnitDescriptor",
			},
			"connectivity_order": {
				Description: "Order in which the ports are connected; sequential or cyclic.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"description": {
				Description: "Detailed information about the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"ethernet_port_speed": {
				Description: "The port speed for ethernet ports in Mbps.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"features": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
							Default:     "capability.FeatureConfig",
						},
						"feature_name": {
							Description:  "Name of the feature that identifies the specific adapter configuration.\n* `RoCEv2` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2.\n* `RoCEv1` - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1.\n* `VMQ` - Capability indicator of the Virtual Machine Queue (VMQ) feature.\n* `VMMQ` - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature.\n* `VMQInterrupts` - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature.\n* `NVGRE` - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature.\n* `ARFS` - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature.\n* `VXLAN` - Capability indicator of the Virtual Extensible LAN (VXLAN) feature.\n* `usNIC` - Capability indicator of the User Space NIC (usNIC) feature.\n* `Advanced Filter` - Capability indicator of the Advanced Filter feature.\n* `Azure Stack Host QOS` - Capability indicator of the Azure Stack Host QOS feature.\n* `GENEVE Offload` - Capability indicator of the Generic Network Virtualization Encapsulation (Geneve) Offload feature.\n* `QinQ` - Capability indicator of the QinQ feature.\n* `SRIOV` - Capability indicator of the Single Root Input Output Virtualization (SR-IOV).\n* `Ether Channel Pinning` - Capability indicator of the Ether Channel Pinning feature.\n* `IPv6 Iscsi Boot` - Capability indicator of the Iscsi Boot via IPV6 protocol.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"RoCEv2", "RoCEv1", "VMQ", "VMMQ", "VMQInterrupts", "NVGRE", "ARFS", "VXLAN", "usNIC", "Advanced Filter", "Azure Stack Host QOS", "GENEVE Offload", "QinQ", "SRIOV", "Ether Channel Pinning", "IPv6 Iscsi Boot"}, false),
							Optional:     true,
							Default:      "RoCEv2",
						},
						"min_adapter_fw_version": {
							Description: "Firmware version of Adapter from which support for this feature is available.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"min_fw_version": {
							Description: "Firmware version of BMC from which support for this feature is available.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "capability.FeatureConfig",
						},
						"supported_fw_versions": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}},
						"supported_in_adapters": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}},
						"supported_in_generations": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type:         schema.TypeInt,
								ValidateFunc: validation.IntInSlice([]int{4, 2, 3, 5}),
							}},
						"unsupported_feature_matrix": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "capability.UnsupportedFeatureConfig",
									},
									"generation": {
										Description:  "The adapter generations that support this feature.\n* `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04.\n* `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02.\n* `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03.\n* `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.",
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntInSlice([]int{4, 2, 3, 5}),
										Optional:     true,
										Default:      4,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "capability.UnsupportedFeatureConfig",
									},
									"unsupportd_features": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringInSlice([]string{"RoCEv2", "RoCEv1", "VMQ", "VMMQ", "VMQInterrupts", "NVGRE", "ARFS", "VXLAN", "usNIC", "Advanced Filter", "Azure Stack Host QOS", "GENEVE Offload", "QinQ", "SRIOV", "Ether Channel Pinning", "IPv6 Iscsi Boot"}, false),
										}},
								},
							},
						},
						"validation_action": {
							Description:  "Action to be taken when validation does not succeed.\n* `Error` - Stop workflow execution by throwing error.\n* `Skip` - Remove the feature from configuration and continue workflow execution.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"Error", "Skip"}, false),
							Optional:     true,
							Default:      "Error",
						},
					},
				},
			},
			"fibre_channel_port_speed": {
				Description: "The port speed for fibre channel ports in Mbps.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fibre_channel_scsi_ioq_limit": {
				Description: "The number of SCSI I/O Queue resources to allocate.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"is_azure_qos_supported": {
				Description: "Indicates that the Azure Stack Host QoS feature is supported by this adapter.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"is_geneve_supported": {
				Description: "Indicates that the GENEVE offload feature is supported by this adapter.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"is_placement_applicable": {
				Description: "This field determines whether vNICs can be placed to the adapters. It is mandatory for all adapters. For third-party adapters, this field is set to 'false', meaning they will only be inventoried, and no LCP configuration will be applied.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"is_secure_boot_supported": {
				Description: "Indicates support for secure boot.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"max_eth_rx_ring_size": {
				Description: "Maximum Ring Size value for vNIC Receive Queue.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     4096,
			},
			"max_eth_tx_ring_size": {
				Description: "Maximum Ring Size value for vNIC Transmit Queue.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     4096,
			},
			"max_rocev2_interfaces": {
				Description: "Maximum number of vNIC interfaces that can be RoCEv2 enabled.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"model": {
				Description: "The model of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"num_dce_ports": {
				Description: "Number of Dce Ports for the adapter.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_pci_links": {
				Description: "Indicates number of PCI Links of the adapter.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "capability.AdapterUnitDescriptor",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"pci_link": {
				Description: "Indicates PCI Link status of adapter.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"prom_card_type": {
				Description: "Prom card type for the adapter.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description:  "The string representation of a tag key.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(1, 128),
							Optional:     true,
						},
						"value": {
							Description:  "The string representation of a tag value.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),
							Optional:     true,
						},
					},
				},
			},
			"vendor": {
				Description: "The vendor of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "The firmware or software version of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
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
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
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
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
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
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
			"vic_id": {
				Description: "Vic Id assigned for the adapter.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceCapabilityAdapterUnitDescriptorCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewCapabilityAdapterUnitDescriptorWithDefaults()

	if v, ok := d.GetOkExists("adapter_generation"); ok {
		x := int32(v.(int))
		o.SetAdapterGeneration(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("capabilities"); ok {
		x := make([]models.CapabilityCapabilityRelationship, 0)
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
			x = append(x, models.MoMoRefAsCapabilityCapabilityRelationship(o))
		}
		if len(x) > 0 {
			o.SetCapabilities(x)
		}
	}

	o.SetClassId("capability.AdapterUnitDescriptor")

	if v, ok := d.GetOk("connectivity_order"); ok {
		x := (v.(string))
		o.SetConnectivityOrder(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOkExists("ethernet_port_speed"); ok {
		x := int64(v.(int))
		o.SetEthernetPortSpeed(x)
	}

	if v, ok := d.GetOk("features"); ok {
		x := make([]models.CapabilityFeatureConfig, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewCapabilityFeatureConfigWithDefaults()
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
			o.SetClassId("capability.FeatureConfig")
			if v, ok := l["feature_name"]; ok {
				{
					x := (v.(string))
					o.SetFeatureName(x)
				}
			}
			if v, ok := l["min_adapter_fw_version"]; ok {
				{
					x := (v.(string))
					o.SetMinAdapterFwVersion(x)
				}
			}
			if v, ok := l["min_fw_version"]; ok {
				{
					x := (v.(string))
					o.SetMinFwVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["supported_fw_versions"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSupportedFwVersions(x)
					}
				}
			}
			if v, ok := l["supported_in_adapters"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSupportedInAdapters(x)
					}
				}
			}
			if v, ok := l["supported_in_generations"]; ok {
				{
					x := make([]int32, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(int32))
						}
					}
					if len(x) > 0 {
						o.SetSupportedInGenerations(x)
					}
				}
			}
			if v, ok := l["unsupported_feature_matrix"]; ok {
				{
					x := make([]models.CapabilityUnsupportedFeatureConfig, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewCapabilityUnsupportedFeatureConfigWithDefaults()
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
						o.SetClassId("capability.UnsupportedFeatureConfig")
						if v, ok := l["generation"]; ok {
							{
								x := int32(v.(int))
								o.SetGeneration(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["unsupportd_features"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									if y.Index(i).Interface() != nil {
										x = append(x, y.Index(i).Interface().(string))
									}
								}
								if len(x) > 0 {
									o.SetUnsupportdFeatures(x)
								}
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetUnsupportedFeatureMatrix(x)
					}
				}
			}
			if v, ok := l["validation_action"]; ok {
				{
					x := (v.(string))
					o.SetValidationAction(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetFeatures(x)
		}
	}

	if v, ok := d.GetOkExists("fibre_channel_port_speed"); ok {
		x := int64(v.(int))
		o.SetFibreChannelPortSpeed(x)
	}

	if v, ok := d.GetOkExists("fibre_channel_scsi_ioq_limit"); ok {
		x := int64(v.(int))
		o.SetFibreChannelScsiIoqLimit(x)
	}

	if v, ok := d.GetOkExists("is_azure_qos_supported"); ok {
		x := (v.(bool))
		o.SetIsAzureQosSupported(x)
	}

	if v, ok := d.GetOkExists("is_geneve_supported"); ok {
		x := (v.(bool))
		o.SetIsGeneveSupported(x)
	}

	if v, ok := d.GetOkExists("is_placement_applicable"); ok {
		x := (v.(bool))
		o.SetIsPlacementApplicable(x)
	}

	if v, ok := d.GetOkExists("is_secure_boot_supported"); ok {
		x := (v.(bool))
		o.SetIsSecureBootSupported(x)
	}

	if v, ok := d.GetOkExists("max_eth_rx_ring_size"); ok {
		x := int64(v.(int))
		o.SetMaxEthRxRingSize(x)
	}

	if v, ok := d.GetOkExists("max_eth_tx_ring_size"); ok {
		x := int64(v.(int))
		o.SetMaxEthTxRingSize(x)
	}

	if v, ok := d.GetOkExists("max_rocev2_interfaces"); ok {
		x := int64(v.(int))
		o.SetMaxRocev2Interfaces(x)
	}

	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOkExists("num_dce_ports"); ok {
		x := int64(v.(int))
		o.SetNumDcePorts(x)
	}

	if v, ok := d.GetOkExists("number_of_pci_links"); ok {
		x := int64(v.(int))
		o.SetNumberOfPciLinks(x)
	}

	o.SetObjectType("capability.AdapterUnitDescriptor")

	if v, ok := d.GetOkExists("pci_link"); ok {
		x := int64(v.(int))
		o.SetPciLink(x)
	}

	if v, ok := d.GetOk("prom_card_type"); ok {
		x := (v.(string))
		o.SetPromCardType(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
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
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	if v, ok := d.GetOk("vic_id"); ok {
		x := (v.(string))
		o.SetVicId(x)
	}

	r := conn.ApiClient.CapabilityApi.CreateCapabilityAdapterUnitDescriptor(conn.ctx).CapabilityAdapterUnitDescriptor(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating CapabilityAdapterUnitDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating CapabilityAdapterUnitDescriptor: %s", responseErr.Error())
	}
	if len(resultMo.GetMoid()) != 0 {
		log.Printf("Moid: %s", resultMo.GetMoid())
		d.SetId(resultMo.GetMoid())
	} else {
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		log.Printf("Mo: %v", resultMo)
	}
	if len(resultMo.GetMoid()) == 0 {
		return de
	}
	return append(de, resourceCapabilityAdapterUnitDescriptorRead(c, d, meta)...)
}

func resourceCapabilityAdapterUnitDescriptorRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	if len(d.Id()) == 0 {
		return de
	}
	conn := meta.(*Config)
	r := conn.ApiClient.CapabilityApi.GetCapabilityAdapterUnitDescriptorByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilityAdapterUnitDescriptor object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilityAdapterUnitDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching CapabilityAdapterUnitDescriptor: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("adapter_generation", (s.GetAdapterGeneration())); err != nil {
		return diag.Errorf("error occurred while setting property AdapterGeneration in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("capabilities", flattenListCapabilityCapabilityRelationship(s.GetCapabilities(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Capabilities in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("connectivity_order", (s.GetConnectivityOrder())); err != nil {
		return diag.Errorf("error occurred while setting property ConnectivityOrder in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("ethernet_port_speed", (s.GetEthernetPortSpeed())); err != nil {
		return diag.Errorf("error occurred while setting property EthernetPortSpeed in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("features", flattenListCapabilityFeatureConfig(s.GetFeatures(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Features in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("fibre_channel_port_speed", (s.GetFibreChannelPortSpeed())); err != nil {
		return diag.Errorf("error occurred while setting property FibreChannelPortSpeed in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("fibre_channel_scsi_ioq_limit", (s.GetFibreChannelScsiIoqLimit())); err != nil {
		return diag.Errorf("error occurred while setting property FibreChannelScsiIoqLimit in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("is_azure_qos_supported", (s.GetIsAzureQosSupported())); err != nil {
		return diag.Errorf("error occurred while setting property IsAzureQosSupported in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("is_geneve_supported", (s.GetIsGeneveSupported())); err != nil {
		return diag.Errorf("error occurred while setting property IsGeneveSupported in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("is_placement_applicable", (s.GetIsPlacementApplicable())); err != nil {
		return diag.Errorf("error occurred while setting property IsPlacementApplicable in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("is_secure_boot_supported", (s.GetIsSecureBootSupported())); err != nil {
		return diag.Errorf("error occurred while setting property IsSecureBootSupported in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("max_eth_rx_ring_size", (s.GetMaxEthRxRingSize())); err != nil {
		return diag.Errorf("error occurred while setting property MaxEthRxRingSize in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("max_eth_tx_ring_size", (s.GetMaxEthTxRingSize())); err != nil {
		return diag.Errorf("error occurred while setting property MaxEthTxRingSize in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("max_rocev2_interfaces", (s.GetMaxRocev2Interfaces())); err != nil {
		return diag.Errorf("error occurred while setting property MaxRocev2Interfaces in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("model", (s.GetModel())); err != nil {
		return diag.Errorf("error occurred while setting property Model in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("num_dce_ports", (s.GetNumDcePorts())); err != nil {
		return diag.Errorf("error occurred while setting property NumDcePorts in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("number_of_pci_links", (s.GetNumberOfPciLinks())); err != nil {
		return diag.Errorf("error occurred while setting property NumberOfPciLinks in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("pci_link", (s.GetPciLink())); err != nil {
		return diag.Errorf("error occurred while setting property PciLink in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("prom_card_type", (s.GetPromCardType())); err != nil {
		return diag.Errorf("error occurred while setting property PromCardType in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("vendor", (s.GetVendor())); err != nil {
		return diag.Errorf("error occurred while setting property Vendor in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return diag.Errorf("error occurred while setting property Version in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	if err := d.Set("vic_id", (s.GetVicId())); err != nil {
		return diag.Errorf("error occurred while setting property VicId in CapabilityAdapterUnitDescriptor object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceCapabilityAdapterUnitDescriptorUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilityAdapterUnitDescriptor{}

	if d.HasChange("adapter_generation") {
		v := d.Get("adapter_generation")
		x := int32(v.(int))
		o.SetAdapterGeneration(x)
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("capabilities") {
		v := d.Get("capabilities")
		x := make([]models.CapabilityCapabilityRelationship, 0)
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
			x = append(x, models.MoMoRefAsCapabilityCapabilityRelationship(o))
		}
		o.SetCapabilities(x)
	}

	o.SetClassId("capability.AdapterUnitDescriptor")

	if d.HasChange("connectivity_order") {
		v := d.Get("connectivity_order")
		x := (v.(string))
		o.SetConnectivityOrder(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("ethernet_port_speed") {
		v := d.Get("ethernet_port_speed")
		x := int64(v.(int))
		o.SetEthernetPortSpeed(x)
	}

	if d.HasChange("features") {
		v := d.Get("features")
		x := make([]models.CapabilityFeatureConfig, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityFeatureConfig{}
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
			o.SetClassId("capability.FeatureConfig")
			if v, ok := l["feature_name"]; ok {
				{
					x := (v.(string))
					o.SetFeatureName(x)
				}
			}
			if v, ok := l["min_adapter_fw_version"]; ok {
				{
					x := (v.(string))
					o.SetMinAdapterFwVersion(x)
				}
			}
			if v, ok := l["min_fw_version"]; ok {
				{
					x := (v.(string))
					o.SetMinFwVersion(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["supported_fw_versions"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSupportedFwVersions(x)
					}
				}
			}
			if v, ok := l["supported_in_adapters"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetSupportedInAdapters(x)
					}
				}
			}
			if v, ok := l["supported_in_generations"]; ok {
				{
					x := make([]int32, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(int32))
						}
					}
					if len(x) > 0 {
						o.SetSupportedInGenerations(x)
					}
				}
			}
			if v, ok := l["unsupported_feature_matrix"]; ok {
				{
					x := make([]models.CapabilityUnsupportedFeatureConfig, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewCapabilityUnsupportedFeatureConfigWithDefaults()
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
						o.SetClassId("capability.UnsupportedFeatureConfig")
						if v, ok := l["generation"]; ok {
							{
								x := int32(v.(int))
								o.SetGeneration(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["unsupportd_features"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									if y.Index(i).Interface() != nil {
										x = append(x, y.Index(i).Interface().(string))
									}
								}
								if len(x) > 0 {
									o.SetUnsupportdFeatures(x)
								}
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetUnsupportedFeatureMatrix(x)
					}
				}
			}
			if v, ok := l["validation_action"]; ok {
				{
					x := (v.(string))
					o.SetValidationAction(x)
				}
			}
			x = append(x, *o)
		}
		o.SetFeatures(x)
	}

	if d.HasChange("fibre_channel_port_speed") {
		v := d.Get("fibre_channel_port_speed")
		x := int64(v.(int))
		o.SetFibreChannelPortSpeed(x)
	}

	if d.HasChange("fibre_channel_scsi_ioq_limit") {
		v := d.Get("fibre_channel_scsi_ioq_limit")
		x := int64(v.(int))
		o.SetFibreChannelScsiIoqLimit(x)
	}

	if d.HasChange("is_azure_qos_supported") {
		v := d.Get("is_azure_qos_supported")
		x := (v.(bool))
		o.SetIsAzureQosSupported(x)
	}

	if d.HasChange("is_geneve_supported") {
		v := d.Get("is_geneve_supported")
		x := (v.(bool))
		o.SetIsGeneveSupported(x)
	}

	if d.HasChange("is_placement_applicable") {
		v := d.Get("is_placement_applicable")
		x := (v.(bool))
		o.SetIsPlacementApplicable(x)
	}

	if d.HasChange("is_secure_boot_supported") {
		v := d.Get("is_secure_boot_supported")
		x := (v.(bool))
		o.SetIsSecureBootSupported(x)
	}

	if d.HasChange("max_eth_rx_ring_size") {
		v := d.Get("max_eth_rx_ring_size")
		x := int64(v.(int))
		o.SetMaxEthRxRingSize(x)
	}

	if d.HasChange("max_eth_tx_ring_size") {
		v := d.Get("max_eth_tx_ring_size")
		x := int64(v.(int))
		o.SetMaxEthTxRingSize(x)
	}

	if d.HasChange("max_rocev2_interfaces") {
		v := d.Get("max_rocev2_interfaces")
		x := int64(v.(int))
		o.SetMaxRocev2Interfaces(x)
	}

	if d.HasChange("model") {
		v := d.Get("model")
		x := (v.(string))
		o.SetModel(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("num_dce_ports") {
		v := d.Get("num_dce_ports")
		x := int64(v.(int))
		o.SetNumDcePorts(x)
	}

	if d.HasChange("number_of_pci_links") {
		v := d.Get("number_of_pci_links")
		x := int64(v.(int))
		o.SetNumberOfPciLinks(x)
	}

	o.SetObjectType("capability.AdapterUnitDescriptor")

	if d.HasChange("pci_link") {
		v := d.Get("pci_link")
		x := int64(v.(int))
		o.SetPciLink(x)
	}

	if d.HasChange("prom_card_type") {
		v := d.Get("prom_card_type")
		x := (v.(string))
		o.SetPromCardType(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
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

	if d.HasChange("vendor") {
		v := d.Get("vendor")
		x := (v.(string))
		o.SetVendor(x)
	}

	if d.HasChange("nr_version") {
		v := d.Get("nr_version")
		x := (v.(string))
		o.SetVersion(x)
	}

	if d.HasChange("vic_id") {
		v := d.Get("vic_id")
		x := (v.(string))
		o.SetVicId(x)
	}

	r := conn.ApiClient.CapabilityApi.UpdateCapabilityAdapterUnitDescriptor(conn.ctx, d.Id()).CapabilityAdapterUnitDescriptor(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating CapabilityAdapterUnitDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating CapabilityAdapterUnitDescriptor: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceCapabilityAdapterUnitDescriptorRead(c, d, meta)...)
}

func resourceCapabilityAdapterUnitDescriptorDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.CapabilityApi.DeleteCapabilityAdapterUnitDescriptor(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "CapabilityAdapterUnitDescriptorDelete: CapabilityAdapterUnitDescriptor object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting CapabilityAdapterUnitDescriptor object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting CapabilityAdapterUnitDescriptor object: %s", deleteErr.Error())
	}
	return de
}
