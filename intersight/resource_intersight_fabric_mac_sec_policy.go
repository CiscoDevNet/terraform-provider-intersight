package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFabricMacSecPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFabricMacSecPolicyCreate,
		ReadContext:   resourceFabricMacSecPolicyRead,
		UpdateContext: resourceFabricMacSecPolicyUpdate,
		DeleteContext: resourceFabricMacSecPolicyDelete,
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
			"cipher_suite": {
				Description:  "Cipher suite to be used for MACsec encryption.\n* `GCM-AES-XPN-256` - An extended Cipher Suite of GCM-AES-256 used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) for enhanced security and scalability.\n* `GCM-AES-128` - This Cipher Suite employs the Advanced Encryption Standard (AES) with a 128-bit key in Galois/Counter Mode, offering both encryption and authentication.\n* `GCM-AES-256` - This Cipher Suite utilizes Advanced Encryption Standard (AES) with a 256-bit key in Galois/Counter Mode, offering a higher level of security compared to GCM-AES-128 due to the larger key size.\n* `GCM-AES-XPN-128` - An extended Cipher Suite of GCM-AES-128  used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) to enhance security and scalability.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"GCM-AES-XPN-256", "GCM-AES-128", "GCM-AES-256", "GCM-AES-XPN-128"}, false),
				Optional:     true,
				Default:      "GCM-AES-XPN-256",
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "fabric.MacSecPolicy",
			},
			"confidentiality_offset": {
				Description:  "The MACsec confidentiality offset specifies the number of bytes starting from the frame header. MACsec encrypts only the bytes after the offset in a frame.\n* `CONF-OFFSET-0` - A value of 0 means the entire ethernet frame is encrypted.\n* `CONF-OFFSET-30` - The first 30 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted.\n* `CONF-OFFSET-50` - The first 50 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"CONF-OFFSET-0", "CONF-OFFSET-30", "CONF-OFFSET-50"}, false),
				Optional:     true,
				Default:      "CONF-OFFSET-0",
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
				Description:  "Description of the policy.",
				Type:         schema.TypeString,
				ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^$|^[a-zA-Z0-9]+[\\x00-\\xFF]*$"), ""), StringLenMaximum(1024)),
				Optional:     true,
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
			"fallback_key_chain": {
				Description: "Fallback keychain for managing an alternative set of security keys to be used when a secure session cannot be established using the primary keychain.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Default:     "fabric.SecKeyChain",
						},
						"name": {
							Description:  "Must be a maximum of 63 characters, without spacing.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "fabric.SecKeyChain",
						},
						"sec_keys": {
							Type:       schema.TypeList,
							MaxItems:   64,
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
										Default:     "fabric.SecKey",
									},
									"cryptographic_algorithm": {
										Description:  "The cryptographic algorithm that employs the cipher-based message authentication code (CMAC) mode of operation with advanced encryption standard (AES).\n* `AES_256_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 256-bit key to generate a CMAC.\n* `AES_128_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 128-bit key to generate a CMAC.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"AES_256_CMAC", "AES_128_CMAC"}, false),
										Optional:     true,
										Default:      "AES_256_CMAC",
									},
									"id": {
										Description:  "Must have an even number of hexadecimal characters (including 0-9 and A-F, only) with a length between 2 and 64 characters. For example, \"10\", \"2000\", \"ABCD1234\".",
										Type:         schema.TypeString,
										ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[0-9A-Fa-f]{2}([0-9A-Fa-f]{2}){0,31}$"), ""), validation.StringLenBetween(2, 64)),
										Optional:     true,
									},
									"is_octet_string_set": {
										Description: "Indicates whether the value of the 'octetString' property has been set.",
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
											if val != nil {
												warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
											}
											return
										}},
									"key_type": {
										Description:  "The type of encryption used for the specified key.\n* `Type-0` - No encryption for the specified octetString.\n* `Type-6` - Proprietary advanced encryption standard for the specified octetString.\n* `Type-7` - Proprietary insecure encryption for the specified octetString.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"Type-0", "Type-6", "Type-7"}, false),
										Optional:     true,
										Default:      "Type-0",
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "fabric.SecKey",
									},
									"octet_string": {
										Description: "The key octet string is a shared secret used in cryptographic operations. The valid size and format of the octet string depend on the selected KeyCryptographicAlgorithm and KeyEncryptionType. It should start with the character 'J'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_duration": {
										Description:  "The key lifetime duration in seconds after the start time. If a non-zero value is configured for the duration, the end time configuration for the key lifetime is ignored.",
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483646),
										Optional:     true,
									},
									"send_lifetime_end_time": {
										Description: "The time of day and date when the key becomes inactive.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_infinite": {
										Description: "Indicates that the key remains active indefinitely after the specified start time. When this parameter is set, the end time and duration configurations for the key lifetime are ignored.",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
									},
									"send_lifetime_start_time": {
										Description: "The time of day and date when the key becomes active.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_time_zone": {
										Description:  "The time zone used for key lifetime configurations.\n* `UTC` - The Universal Time (UTC) for key lifetime configurations.\n* `Local` - The local time zone of the device for key lifetime configurations.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"UTC", "Local"}, false),
										Optional:     true,
										Default:      "UTC",
									},
									"send_lifetime_unlimited": {
										Description: "Indicates that the key is always active. When this parameter is set, all other key lifetime configurations are ignored.",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
									},
								},
							},
						},
					},
				},
			},
			"include_icv_indicator": {
				Description: "Configures inclusion of the optional integrity check value (ICV) indicator as part of the transmitted MACsec key agreement protocol data unit (PDU).",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"key_server_priority": {
				Description:  "The key server is selected by comparing key-server priority values during MACsec key agreement (MKA) message exchange between peer devices. Valid values range from 0 to 255. The lower the value, the higher the chance it will be selected as the key server.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Default:      16,
			},
			"mac_sec_ea_pol": {
				Description: "Extensible authentication protocol over LAN (EAPoL). MACsec transmits MACsec key agreement (MKA) protocol data units (PDUs) using EAPoL packets to establish a secure session.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Default:     "fabric.MacSecEaPol",
						},
						"ea_pol_ethertype": {
							Description:  "Ethertype to use in extensible authentication protocol over LAN (EAPoL) frames for MACsec key agreement (MKA) protocol data units (PDUs). The range is between 0x600 - 0xffff.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^0x([6-9A-Fa-f][0-9A-Fa-f]{2}|[A-Fa-f0-9][0-9A-Fa-f]{3})$"), ""),
							Optional:     true,
							Default:      "34958",
						},
						"ea_pol_mac_address": {
							Description:  "MAC address to use in extensible authentication protocol over LAN (EAPoL) for MACsec key agreement (MKA) protocol data units (PDUs). EAPol mac address should not be equal to all-zero (0000.0000.0000).",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[0-9A-Fa-f]{4}\\.[0-9A-Fa-f]{4}\\.[0-9A-Fa-f]{4}$"), ""),
							Optional:     true,
							Default:      "0180.C200.0003",
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "fabric.MacSecEaPol",
						},
					},
				},
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
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description:  "Name of the concrete policy.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.:-]{1,64}$"), ""),
				Optional:     true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "fabric.MacSecPolicy",
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
				ForceNew: true,
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
			"primary_key_chain": {
				Description: "Primary keychain for managing the default set of security keys for encryption and decryption.",
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Default:     "fabric.SecKeyChain",
						},
						"name": {
							Description:  "Must be a maximum of 63 characters, without spacing.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "fabric.SecKeyChain",
						},
						"sec_keys": {
							Type:       schema.TypeList,
							MaxItems:   64,
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
										Default:     "fabric.SecKey",
									},
									"cryptographic_algorithm": {
										Description:  "The cryptographic algorithm that employs the cipher-based message authentication code (CMAC) mode of operation with advanced encryption standard (AES).\n* `AES_256_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 256-bit key to generate a CMAC.\n* `AES_128_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 128-bit key to generate a CMAC.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"AES_256_CMAC", "AES_128_CMAC"}, false),
										Optional:     true,
										Default:      "AES_256_CMAC",
									},
									"id": {
										Description:  "Must have an even number of hexadecimal characters (including 0-9 and A-F, only) with a length between 2 and 64 characters. For example, \"10\", \"2000\", \"ABCD1234\".",
										Type:         schema.TypeString,
										ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[0-9A-Fa-f]{2}([0-9A-Fa-f]{2}){0,31}$"), ""), validation.StringLenBetween(2, 64)),
										Optional:     true,
									},
									"is_octet_string_set": {
										Description: "Indicates whether the value of the 'octetString' property has been set.",
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
											if val != nil {
												warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
											}
											return
										}},
									"key_type": {
										Description:  "The type of encryption used for the specified key.\n* `Type-0` - No encryption for the specified octetString.\n* `Type-6` - Proprietary advanced encryption standard for the specified octetString.\n* `Type-7` - Proprietary insecure encryption for the specified octetString.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"Type-0", "Type-6", "Type-7"}, false),
										Optional:     true,
										Default:      "Type-0",
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "fabric.SecKey",
									},
									"octet_string": {
										Description: "The key octet string is a shared secret used in cryptographic operations. The valid size and format of the octet string depend on the selected KeyCryptographicAlgorithm and KeyEncryptionType. It should start with the character 'J'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_duration": {
										Description:  "The key lifetime duration in seconds after the start time. If a non-zero value is configured for the duration, the end time configuration for the key lifetime is ignored.",
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 2147483646),
										Optional:     true,
									},
									"send_lifetime_end_time": {
										Description: "The time of day and date when the key becomes inactive.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_infinite": {
										Description: "Indicates that the key remains active indefinitely after the specified start time. When this parameter is set, the end time and duration configurations for the key lifetime are ignored.",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
									},
									"send_lifetime_start_time": {
										Description: "The time of day and date when the key becomes active.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"send_lifetime_time_zone": {
										Description:  "The time zone used for key lifetime configurations.\n* `UTC` - The Universal Time (UTC) for key lifetime configurations.\n* `Local` - The local time zone of the device for key lifetime configurations.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"UTC", "Local"}, false),
										Optional:     true,
										Default:      "UTC",
									},
									"send_lifetime_unlimited": {
										Description: "Indicates that the key is always active. When this parameter is set, all other key lifetime configurations are ignored.",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
									},
								},
							},
						},
					},
				},
			},
			"replay_window_size": {
				Description:  "Defines the size of the replay protection window. It determines the number of packets that can be received out of order without being considered replay attacks.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 596000000),
				Optional:     true,
				Default:      148809600,
			},
			"sak_expiry_time": {
				Description:  "Time in seconds to force secure association key (SAK) rekey. Valid range is from 60 to 2592000 seconds when configured. When not configured, the SAK rekey interval is determined based on the interface speed.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2592000),
				Optional:     true,
			},
			"security_policy": {
				Description:  "The security policy specifies the level of MACsec enforcement on network traffic passing through a given interface.\nShould secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow. Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured.\n* `Should-secure` - Should secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow.\n* `Must-secure` - Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Should-secure", "Must-secure"}, false),
				Optional:     true,
				Default:      "Should-secure",
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
						"ancestor_definitions": {
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
						"definition": {
							Description: "The definition is a reference to the tag definition object.\nThe tag definition object contains the properties of the tag such as name, type, and description.",
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
						"key": {
							Description:  "The string representation of a tag key.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(1, 256),
							Optional:     true,
						},
						"propagated": {
							Description: "Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"sys_tag": {
							Description: "Specifies whether the tag is user-defined or owned by the system.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"type": {
							Description: "An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.\n* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.\n* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \"A/B/C\", then \"A\" is the parent tag, \"B\" is the child tag of \"A\" and \"C\" is the child tag of \"B\".",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"value": {
							Description:  "The string representation of a tag value.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),
							Optional:     true,
						},
					},
				},
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
		},
	}
}

func resourceFabricMacSecPolicyCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewFabricMacSecPolicyWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("cipher_suite"); ok {
		x := (v.(string))
		o.SetCipherSuite(x)
	}

	o.SetClassId("fabric.MacSecPolicy")

	if v, ok := d.GetOk("confidentiality_offset"); ok {
		x := (v.(string))
		o.SetConfidentialityOffset(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("fallback_key_chain"); ok {
		p := make([]models.FabricSecKeyChain, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewFabricSecKeyChainWithDefaults()
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
			o.SetClassId("fabric.SecKeyChain")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["sec_keys"]; ok {
				{
					x := make([]models.FabricSecKey, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewFabricSecKeyWithDefaults()
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
						o.SetClassId("fabric.SecKey")
						if v, ok := l["cryptographic_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetCryptographicAlgorithm(x)
							}
						}
						if v, ok := l["id"]; ok {
							{
								x := (v.(string))
								o.SetId(x)
							}
						}
						if v, ok := l["key_type"]; ok {
							{
								x := (v.(string))
								o.SetKeyType(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["octet_string"]; ok {
							{
								x := (v.(string))
								o.SetOctetString(x)
							}
						}
						if v, ok := l["send_lifetime_duration"]; ok {
							{
								x := int64(v.(int))
								o.SetSendLifetimeDuration(x)
							}
						}
						if v, ok := l["send_lifetime_end_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeEndTime(x)
							}
						}
						if v, ok := l["send_lifetime_infinite"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeInfinite(x)
							}
						}
						if v, ok := l["send_lifetime_start_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeStartTime(x)
							}
						}
						if v, ok := l["send_lifetime_time_zone"]; ok {
							{
								x := (v.(string))
								o.SetSendLifetimeTimeZone(x)
							}
						}
						if v, ok := l["send_lifetime_unlimited"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeUnlimited(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetSecKeys(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetFallbackKeyChain(x)
		}
	}

	if v, ok := d.GetOkExists("include_icv_indicator"); ok {
		x := (v.(bool))
		o.SetIncludeIcvIndicator(x)
	}

	if v, ok := d.GetOkExists("key_server_priority"); ok {
		x := int64(v.(int))
		o.SetKeyServerPriority(x)
	}

	if v, ok := d.GetOk("mac_sec_ea_pol"); ok {
		p := make([]models.FabricMacSecEaPol, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewFabricMacSecEaPolWithDefaults()
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
			o.SetClassId("fabric.MacSecEaPol")
			if v, ok := l["ea_pol_ethertype"]; ok {
				{
					x := (v.(string))
					o.SetEaPolEthertype(x)
				}
			}
			if v, ok := l["ea_pol_mac_address"]; ok {
				{
					x := (v.(string))
					o.SetEaPolMacAddress(x)
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
			o.SetMacSecEaPol(x)
		}
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("fabric.MacSecPolicy")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("primary_key_chain"); ok {
		p := make([]models.FabricSecKeyChain, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewFabricSecKeyChainWithDefaults()
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
			o.SetClassId("fabric.SecKeyChain")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["sec_keys"]; ok {
				{
					x := make([]models.FabricSecKey, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewFabricSecKeyWithDefaults()
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
						o.SetClassId("fabric.SecKey")
						if v, ok := l["cryptographic_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetCryptographicAlgorithm(x)
							}
						}
						if v, ok := l["id"]; ok {
							{
								x := (v.(string))
								o.SetId(x)
							}
						}
						if v, ok := l["key_type"]; ok {
							{
								x := (v.(string))
								o.SetKeyType(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["octet_string"]; ok {
							{
								x := (v.(string))
								o.SetOctetString(x)
							}
						}
						if v, ok := l["send_lifetime_duration"]; ok {
							{
								x := int64(v.(int))
								o.SetSendLifetimeDuration(x)
							}
						}
						if v, ok := l["send_lifetime_end_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeEndTime(x)
							}
						}
						if v, ok := l["send_lifetime_infinite"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeInfinite(x)
							}
						}
						if v, ok := l["send_lifetime_start_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeStartTime(x)
							}
						}
						if v, ok := l["send_lifetime_time_zone"]; ok {
							{
								x := (v.(string))
								o.SetSendLifetimeTimeZone(x)
							}
						}
						if v, ok := l["send_lifetime_unlimited"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeUnlimited(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetSecKeys(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrimaryKeyChain(x)
		}
	}

	if v, ok := d.GetOkExists("replay_window_size"); ok {
		x := int64(v.(int))
		o.SetReplayWindowSize(x)
	}

	if v, ok := d.GetOkExists("sak_expiry_time"); ok {
		x := int64(v.(int))
		o.SetSakExpiryTime(x)
	}

	if v, ok := d.GetOk("security_policy"); ok {
		x := (v.(string))
		o.SetSecurityPolicy(x)
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
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	r := conn.ApiClient.FabricApi.CreateFabricMacSecPolicy(conn.ctx).FabricMacSecPolicy(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating FabricMacSecPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating FabricMacSecPolicy: %s", responseErr.Error())
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
	return append(de, resourceFabricMacSecPolicyRead(c, d, meta)...)
}

func resourceFabricMacSecPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	if len(d.Id()) == 0 {
		return de
	}
	conn := meta.(*Config)
	r := conn.ApiClient.FabricApi.GetFabricMacSecPolicyByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "FabricMacSecPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricMacSecPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching FabricMacSecPolicy: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("cipher_suite", (s.GetCipherSuite())); err != nil {
		return diag.Errorf("error occurred while setting property CipherSuite in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("confidentiality_offset", (s.GetConfidentialityOffset())); err != nil {
		return diag.Errorf("error occurred while setting property ConfidentialityOffset in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("fallback_key_chain", flattenMapFabricSecKeyChain(s.GetFallbackKeyChain(), d)); err != nil {
		return diag.Errorf("error occurred while setting property FallbackKeyChain in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("include_icv_indicator", (s.GetIncludeIcvIndicator())); err != nil {
		return diag.Errorf("error occurred while setting property IncludeIcvIndicator in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("key_server_priority", (s.GetKeyServerPriority())); err != nil {
		return diag.Errorf("error occurred while setting property KeyServerPriority in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("mac_sec_ea_pol", flattenMapFabricMacSecEaPol(s.GetMacSecEaPol(), d)); err != nil {
		return diag.Errorf("error occurred while setting property MacSecEaPol in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("primary_key_chain", flattenMapFabricSecKeyChain(s.GetPrimaryKeyChain(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PrimaryKeyChain in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("replay_window_size", (s.GetReplayWindowSize())); err != nil {
		return diag.Errorf("error occurred while setting property ReplayWindowSize in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("sak_expiry_time", (s.GetSakExpiryTime())); err != nil {
		return diag.Errorf("error occurred while setting property SakExpiryTime in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("security_policy", (s.GetSecurityPolicy())); err != nil {
		return diag.Errorf("error occurred while setting property SecurityPolicy in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in FabricMacSecPolicy object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in FabricMacSecPolicy object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceFabricMacSecPolicyUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricMacSecPolicy{}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("cipher_suite") {
		v := d.Get("cipher_suite")
		x := (v.(string))
		o.SetCipherSuite(x)
	}

	o.SetClassId("fabric.MacSecPolicy")

	if d.HasChange("confidentiality_offset") {
		v := d.Get("confidentiality_offset")
		x := (v.(string))
		o.SetConfidentialityOffset(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("fallback_key_chain") {
		v := d.Get("fallback_key_chain")
		p := make([]models.FabricSecKeyChain, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.FabricSecKeyChain{}
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
			o.SetClassId("fabric.SecKeyChain")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["sec_keys"]; ok {
				{
					x := make([]models.FabricSecKey, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewFabricSecKeyWithDefaults()
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
						o.SetClassId("fabric.SecKey")
						if v, ok := l["cryptographic_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetCryptographicAlgorithm(x)
							}
						}
						if v, ok := l["id"]; ok {
							{
								x := (v.(string))
								o.SetId(x)
							}
						}
						if v, ok := l["key_type"]; ok {
							{
								x := (v.(string))
								o.SetKeyType(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["octet_string"]; ok {
							{
								x := (v.(string))
								o.SetOctetString(x)
							}
						}
						if v, ok := l["send_lifetime_duration"]; ok {
							{
								x := int64(v.(int))
								o.SetSendLifetimeDuration(x)
							}
						}
						if v, ok := l["send_lifetime_end_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeEndTime(x)
							}
						}
						if v, ok := l["send_lifetime_infinite"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeInfinite(x)
							}
						}
						if v, ok := l["send_lifetime_start_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeStartTime(x)
							}
						}
						if v, ok := l["send_lifetime_time_zone"]; ok {
							{
								x := (v.(string))
								o.SetSendLifetimeTimeZone(x)
							}
						}
						if v, ok := l["send_lifetime_unlimited"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeUnlimited(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetSecKeys(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetFallbackKeyChain(x)
		}
	}

	if d.HasChange("include_icv_indicator") {
		v := d.Get("include_icv_indicator")
		x := (v.(bool))
		o.SetIncludeIcvIndicator(x)
	}

	if d.HasChange("key_server_priority") {
		v := d.Get("key_server_priority")
		x := int64(v.(int))
		o.SetKeyServerPriority(x)
	}

	if d.HasChange("mac_sec_ea_pol") {
		v := d.Get("mac_sec_ea_pol")
		p := make([]models.FabricMacSecEaPol, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.FabricMacSecEaPol{}
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
			o.SetClassId("fabric.MacSecEaPol")
			if v, ok := l["ea_pol_ethertype"]; ok {
				{
					x := (v.(string))
					o.SetEaPolEthertype(x)
				}
			}
			if v, ok := l["ea_pol_mac_address"]; ok {
				{
					x := (v.(string))
					o.SetEaPolMacAddress(x)
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
			o.SetMacSecEaPol(x)
		}
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("fabric.MacSecPolicy")

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if d.HasChange("primary_key_chain") {
		v := d.Get("primary_key_chain")
		p := make([]models.FabricSecKeyChain, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.FabricSecKeyChain{}
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
			o.SetClassId("fabric.SecKeyChain")
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["sec_keys"]; ok {
				{
					x := make([]models.FabricSecKey, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewFabricSecKeyWithDefaults()
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
						o.SetClassId("fabric.SecKey")
						if v, ok := l["cryptographic_algorithm"]; ok {
							{
								x := (v.(string))
								o.SetCryptographicAlgorithm(x)
							}
						}
						if v, ok := l["id"]; ok {
							{
								x := (v.(string))
								o.SetId(x)
							}
						}
						if v, ok := l["key_type"]; ok {
							{
								x := (v.(string))
								o.SetKeyType(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["octet_string"]; ok {
							{
								x := (v.(string))
								o.SetOctetString(x)
							}
						}
						if v, ok := l["send_lifetime_duration"]; ok {
							{
								x := int64(v.(int))
								o.SetSendLifetimeDuration(x)
							}
						}
						if v, ok := l["send_lifetime_end_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeEndTime(x)
							}
						}
						if v, ok := l["send_lifetime_infinite"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeInfinite(x)
							}
						}
						if v, ok := l["send_lifetime_start_time"]; ok {
							{
								x, _ := time.Parse(time.RFC1123, v.(string))
								o.SetSendLifetimeStartTime(x)
							}
						}
						if v, ok := l["send_lifetime_time_zone"]; ok {
							{
								x := (v.(string))
								o.SetSendLifetimeTimeZone(x)
							}
						}
						if v, ok := l["send_lifetime_unlimited"]; ok {
							{
								x := (v.(bool))
								o.SetSendLifetimeUnlimited(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetSecKeys(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetPrimaryKeyChain(x)
		}
	}

	if d.HasChange("replay_window_size") {
		v := d.Get("replay_window_size")
		x := int64(v.(int))
		o.SetReplayWindowSize(x)
	}

	if d.HasChange("sak_expiry_time") {
		v := d.Get("sak_expiry_time")
		x := int64(v.(int))
		o.SetSakExpiryTime(x)
	}

	if d.HasChange("security_policy") {
		v := d.Get("security_policy")
		x := (v.(string))
		o.SetSecurityPolicy(x)
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

	r := conn.ApiClient.FabricApi.UpdateFabricMacSecPolicy(conn.ctx, d.Id()).FabricMacSecPolicy(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating FabricMacSecPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating FabricMacSecPolicy: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceFabricMacSecPolicyRead(c, d, meta)...)
}

func resourceFabricMacSecPolicyDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.FabricApi.DeleteFabricMacSecPolicy(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "FabricMacSecPolicyDelete: FabricMacSecPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting FabricMacSecPolicy object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting FabricMacSecPolicy object: %s", deleteErr.Error())
	}
	return de
}
