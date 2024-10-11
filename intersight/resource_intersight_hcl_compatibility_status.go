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

func resourceHclCompatibilityStatus() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHclCompatibilityStatusCreate,
		ReadContext:   resourceHclCompatibilityStatusRead,
		DeleteContext: resourceHclCompatibilityStatusDelete,
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
				}, ForceNew: true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
				ForceNew:         true,
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
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ForceNew: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "hcl.CompatibilityStatus",
				ForceNew:    true,
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
				}, ForceNew: true,
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
				}, ForceNew: true,
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
				}, ForceNew: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "hcl.CompatibilityStatus",
				ForceNew:    true,
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				}, ForceNew: true,
			},
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
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ForceNew: true,
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
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
							ForceNew:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
				ForceNew: true,
			},
			"profile_list": {
				Type:       schema.TypeList,
				MinItems:   1,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "hcl.HardwareCompatibilityProfile",
							ForceNew:    true,
						},
						"driver_iso_url": {
							Description: "Url for the ISO with the drivers supported for the server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"error_code": {
							Description: "Error code indicating the compatibility status.\n* `Success` - The input combination is valid.\n* `Unknown` - Unknown API request to the service.\n* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.\n* `InvalidUcsVersion` - UCS Version is not in expected format.\n* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.\n* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.\n* `OSUnknown` - OS vendor or version is not known as per the HCL database.\n* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.\n* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.\n* `ProductUnknown` - Product is not known as per the HCL database.\n* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.\n* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.\n* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.\n* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.\n* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.\n* `InternalError` - API requests to the service have either failed or timed out.\n* `MarshallingError` - Error in JSON marshalling.\n* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}, ForceNew: true,
						},
						"id": {
							Description: "Identifier of the hardware compatibility profile.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "hcl.HardwareCompatibilityProfile",
							ForceNew:    true,
						},
						"os_vendor": {
							Description: "Vendor of the Operating System running on the server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"os_version": {
							Description: "Version of the Operating System running on the server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"personality": {
							Description: "Personality indicating the personality with the which the sever is used in a Hyperflex environment.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"processor_model": {
							Description: "Model of the processor present in the server.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"products": {
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
										ForceNew:         true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "hcl.Product",
										ForceNew:    true,
									},
									"driver_names": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, ForceNew: true,
									},
									"error_code": {
										Description: "Error code indicating the support status.\n* `Success` - The input combination is valid.\n* `Unknown` - Unknown API request to the service.\n* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.\n* `InvalidUcsVersion` - UCS Version is not in expected format.\n* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.\n* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.\n* `OSUnknown` - OS vendor or version is not known as per the HCL database.\n* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.\n* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.\n* `ProductUnknown` - Product is not known as per the HCL database.\n* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.\n* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.\n* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.\n* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.\n* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.\n* `InternalError` - API requests to the service have either failed or timed out.\n* `MarshallingError` - Error in JSON marshalling.\n* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
											if val != nil {
												warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
											}
											return
										}, ForceNew: true,
									},
									"firmwares": {
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
													ForceNew:         true,
												},
												"class_id": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "hcl.Firmware",
													ForceNew:    true,
												},
												"driver_name": {
													Description: "Protocol for which the driver is provided. E.g.  enic, fnic, lsi_mr3.",
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
												},
												"driver_version": {
													Description: "Version of the Driver supported.",
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
												},
												"error_code": {
													Description: "Error code for the support status.\n* `Success` - The input combination is valid.\n* `Unknown` - Unknown API request to the service.\n* `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database.\n* `InvalidUcsVersion` - UCS Version is not in expected format.\n* `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database.\n* `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database.\n* `OSUnknown` - OS vendor or version is not known as per the HCL database.\n* `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database.\n* `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported.\n* `ProductUnknown` - Product is not known as per the HCL database.\n* `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version.\n* `DriverNameNotSupported` - Driver protocol or name is not supported for the given product.\n* `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination.\n* `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination.\n* `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination.\n* `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination.\n* `InternalError` - API requests to the service have either failed or timed out.\n* `MarshallingError` - Error in JSON marshalling.\n* `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
														if val != nil {
															warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
														}
														return
													}, ForceNew: true,
												},
												"firmware_version": {
													Description: "Firmware version of the product/adapter supported.",
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
												},
												"id": {
													Description: "Identifier of the firmware.",
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
												},
												"latest_driver": {
													Description: "True if the driver is latest recommended driver.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
														if val != nil {
															warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
														}
														return
													}, ForceNew: true,
												},
												"latest_firmware": {
													Description: "True if the firmware is latest recommended firmware.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
														if val != nil {
															warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
														}
														return
													}, ForceNew: true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "hcl.Firmware",
													ForceNew:    true,
												},
											},
										},
										ForceNew: true,
									},
									"id": {
										Description: "Identifier of the product.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"model": {
										Description: "Model/PID of the product/adapter.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "hcl.Product",
										ForceNew:    true,
									},
									"revision": {
										Description: "Revision of the adapter model.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"type": {
										Description:  "Type of the product/adapter say OCP, PT, GPU.\n* `` - Default type of the Product.\n* `Adapter` - Represents network adapter cards.\n* `StorageController` - Represents storage controllers.\n* `GPU` - Represents graphics cards.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"", "Adapter", "StorageController", "GPU"}, false),
										Optional:     true,
										Default:      "",
										ForceNew:     true,
									},
									"vendor": {
										Description: "Vendor of the product or adapter.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ForceNew: true,
						},
						"server_model": {
							Description: "Model of the server as returned by UCSM/CIMC XML API.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"server_revision": {
							Description: "Revision of the server model.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"ucs_version": {
							Description: "Version of the UCS software.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"version_type": {
							Description:  "Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.\n* `UCSM` - The server is managed by UCS Manager.\n* `IMC` - The server is standalone managed by CIMC.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"UCSM", "IMC"}, false),
							Optional:     true,
							Default:      "UCSM",
							ForceNew:     true,
						},
					},
				},
				ForceNew: true,
			},
			"request_type": {
				Description:  "Type of the request to be served.\n* `FillSupportedVersions` - Responds with the supported firmware and driver versions. The API doesn't expect firmware and driver versions to be passed in the request and ignores if passed.\n* `CheckCompatibility` - Checks the compatibility for the given firmware and driver versions. This request type expects the firmware and driver versions to filled and the service validates the values and responds back with the error codes.\n* `GetRecommendedDrivers` - Responds with the supported drivers. The API expects firmware version to be filled. The API populates driver ISO url for the given server model. Today the link is same for all servers managed by UCSM whereas it depends on the model for Standalone servers.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"FillSupportedVersions", "CheckCompatibility", "GetRecommendedDrivers"}, false),
				Optional:     true,
				Default:      "FillSupportedVersions",
				ForceNew:     true,
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
				}, ForceNew: true,
			},
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
							ForceNew:         true,
						},
						"key": {
							Description:  "The string representation of a tag key.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(1, 128),
							Optional:     true,
							ForceNew:     true,
						},
						"value": {
							Description:  "The string representation of a tag value.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),
							Optional:     true,
							ForceNew:     true,
						},
					},
				},
				ForceNew: true,
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
							ForceNew:         true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
							ForceNew:    true,
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
										ForceNew:         true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
										ForceNew:    true,
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ForceNew: true,
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
							}, ForceNew: true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
							ForceNew:    true,
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
										ForceNew:         true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
										ForceNew:    true,
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ForceNew:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
							ForceNew: true,
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
							}, ForceNew: true,
						},
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
							}, ForceNew: true,
						},
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
							}, ForceNew: true,
						},
					},
				},
				ForceNew: true,
			},
		},
	}
}

func resourceHclCompatibilityStatusCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewHclCompatibilityStatusWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("hcl.CompatibilityStatus")

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("hcl.CompatibilityStatus")

	if v, ok := d.GetOk("profile_list"); ok {
		x := make([]models.HclHardwareCompatibilityProfile, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewHclHardwareCompatibilityProfileWithDefaults()
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
			o.SetClassId("hcl.HardwareCompatibilityProfile")
			if v, ok := l["driver_iso_url"]; ok {
				{
					x := (v.(string))
					o.SetDriverIsoUrl(x)
				}
			}
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
			if v, ok := l["os_vendor"]; ok {
				{
					x := (v.(string))
					o.SetOsVendor(x)
				}
			}
			if v, ok := l["os_version"]; ok {
				{
					x := (v.(string))
					o.SetOsVersion(x)
				}
			}
			if v, ok := l["personality"]; ok {
				{
					x := (v.(string))
					o.SetPersonality(x)
				}
			}
			if v, ok := l["processor_model"]; ok {
				{
					x := (v.(string))
					o.SetProcessorModel(x)
				}
			}
			if v, ok := l["products"]; ok {
				{
					x := make([]models.HclProduct, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewHclProductWithDefaults()
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
						o.SetClassId("hcl.Product")
						if v, ok := l["driver_names"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									if y.Index(i).Interface() != nil {
										x = append(x, y.Index(i).Interface().(string))
									}
								}
								if len(x) > 0 {
									o.SetDriverNames(x)
								}
							}
						}
						if v, ok := l["firmwares"]; ok {
							{
								x := make([]models.HclFirmware, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewHclFirmwareWithDefaults()
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
									o.SetClassId("hcl.Firmware")
									if v, ok := l["driver_name"]; ok {
										{
											x := (v.(string))
											o.SetDriverName(x)
										}
									}
									if v, ok := l["driver_version"]; ok {
										{
											x := (v.(string))
											o.SetDriverVersion(x)
										}
									}
									if v, ok := l["firmware_version"]; ok {
										{
											x := (v.(string))
											o.SetFirmwareVersion(x)
										}
									}
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
									x = append(x, *o)
								}
								if len(x) > 0 {
									o.SetFirmwares(x)
								}
							}
						}
						if v, ok := l["id"]; ok {
							{
								x := (v.(string))
								o.SetId(x)
							}
						}
						if v, ok := l["model"]; ok {
							{
								x := (v.(string))
								o.SetModel(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["revision"]; ok {
							{
								x := (v.(string))
								o.SetRevision(x)
							}
						}
						if v, ok := l["type"]; ok {
							{
								x := (v.(string))
								o.SetType(x)
							}
						}
						if v, ok := l["vendor"]; ok {
							{
								x := (v.(string))
								o.SetVendor(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetProducts(x)
					}
				}
			}
			if v, ok := l["server_model"]; ok {
				{
					x := (v.(string))
					o.SetServerModel(x)
				}
			}
			if v, ok := l["server_revision"]; ok {
				{
					x := (v.(string))
					o.SetServerRevision(x)
				}
			}
			if v, ok := l["ucs_version"]; ok {
				{
					x := (v.(string))
					o.SetUcsVersion(x)
				}
			}
			if v, ok := l["version_type"]; ok {
				{
					x := (v.(string))
					o.SetVersionType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetProfileList(x)
		}
	}

	if v, ok := d.GetOk("request_type"); ok {
		x := (v.(string))
		o.SetRequestType(x)
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

	r := conn.ApiClient.HclApi.CreateHclCompatibilityStatus(conn.ctx).HclCompatibilityStatus(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating HclCompatibilityStatus: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating HclCompatibilityStatus: %s", responseErr.Error())
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	log.Printf("Mo: %v", resultMo)
	return de
}

func resourceHclCompatibilityStatusRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", d)
	var de diag.Diagnostics
	return de
}

func resourceHclCompatibilityStatusDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	var warning = diag.Diagnostic{Severity: diag.Warning, Summary: "HclCompatibilityStatus does not allow delete functionality"}
	de = append(de, warning)
	return de
}
