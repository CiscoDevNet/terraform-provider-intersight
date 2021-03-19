package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOsConfigurationFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOsConfigurationFileCreate,
		ReadContext:   resourceOsConfigurationFileRead,
		UpdateContext: resourceOsConfigurationFileUpdate,
		DeleteContext: resourceOsConfigurationFileDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"catalog": {
				Description: "A reference to a osCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
							Computed:    true,
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
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the OS ConfigurationFile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"distributions": {
				Description: "An array of relationships to hclOperatingSystem resources.",
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
							Computed:    true,
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
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"file_content": {
				Description: "The content of the entire configuration file is stored as value. The content\ncan either be a static file content or a template content.\nThe template is expected to conform to the golang template syntax. The values\nfrom os.Answers properties will be used to populate this template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal": {
				Description: "The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the OS ConfigurationFile that uniquely identifies the configuration file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"placeholders": {
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
							Computed:    true,
						},
						"is_value_set": {
							Description: "Flag to indicate if value is set. Value will be used to check if any edit.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Description: "Definition of place holder.",
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
										Computed:    true,
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
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
													Computed:    true,
												},
												"is_value_set": {
													Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeMap,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													}, Optional: true,
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"description": {
										Description: "Provide a detailed description of the data type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"display_meta": {
										Description: "Captures the meta data needed for displaying workflow data types in Intersight User Interface.",
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
													Computed:    true,
												},
												"inventory_selector": {
													Description: "Inventory selector specified for primitive data property should be used in Intersight User Interface.",
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"widget_type": {
													Description: "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "None",
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"input_parameters": {
										Description: "JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.",
										Type:        schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										}, Optional: true,
									},
									"label": {
										Description: "Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"properties": {
										Description: "Primitive data type properties.",
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
													Computed:    true,
												},
												"constraints": {
													Description: "Constraints that must be applied to the parameter value supplied for this data type.",
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
																Computed:    true,
															},
															"enum_list": {
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
																			Computed:    true,
																		},
																		"label": {
																			Description: "Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_).",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"value": {
																			Description: "Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_).",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
																ConfigMode: schema.SchemaConfigModeAttr,
																Computed:   true,
															},
															"max": {
																Description: "Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.",
																Type:        schema.TypeFloat,
																Optional:    true,
															},
															"min": {
																Description: "Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.",
																Type:        schema.TypeFloat,
																Optional:    true,
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"regex": {
																Description: "When the parameter is a string this regular expression is used to ensure the value is valid.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
													ConfigMode: schema.SchemaConfigModeAttr,
													Computed:   true,
												},
												"inventory_selector": {
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
																Computed:    true,
															},
															"display_attributes": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString}},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"selector": {
																Description: "Field to hold an Intersight API along with an optional filter to narrow down the search options.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"value_attribute": {
																Description: "A property from the Intersight object, value of which can be used as value for referenced input definition.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
													ConfigMode: schema.SchemaConfigModeAttr,
													Computed:   true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"secure": {
													Description: "Intersight supports secure properties as task input/output. The values of\nthese properties are encrypted and stored in Intersight.\nThis flag marks the property to be secure when it is set to true.",
													Type:        schema.TypeBool,
													Optional:    true,
													ForceNew:    true,
												},
												"type": {
													Description: "Specify the enum type for primitive data type.\n* `string` - Enum to specify a string data type.\n* `integer` - Enum to specify an integer32 data type.\n* `float` - Enum to specify a float64 data type.\n* `boolean` - Enum to specify a boolean data type.\n* `json` - Enum to specify a json data type.\n* `enum` - Enum to specify a enum data type which is a list of pre-defined strings.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "string",
												},
											},
										},
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
						},
						"value": {
							Description: "Value for placeholder provided by user.",
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}, Optional: true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"supported": {
				Description: "An internal property that is used to distinguish between the pre-canned OS\nconfiguration file entries and user provided entries.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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
		},
	}
}

func resourceOsConfigurationFileCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewOsConfigurationFileWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.OsCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOsCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("os.ConfigurationFile")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("distributions"); ok {
		x := make([]models.HclOperatingSystemRelationship, 0)
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
			x = append(x, models.MoMoRefAsHclOperatingSystemRelationship(o))
		}
		if len(x) > 0 {
			o.SetDistributions(x)
		}
	}

	if v, ok := d.GetOk("file_content"); ok {
		x := (v.(string))
		o.SetFileContent(x)
	}

	if v, ok := d.GetOkExists("internal"); ok {
		x := v.(bool)
		o.SetInternal(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("os.ConfigurationFile")

	if v, ok := d.GetOk("placeholders"); ok {
		x := make([]models.OsPlaceHolder, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewOsPlaceHolderWithDefaults()
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
			o.SetClassId("os.PlaceHolder")
			if v, ok := l["is_value_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsValueSet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					p := make([]models.WorkflowPrimitiveDataType, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewWorkflowPrimitiveDataTypeWithDefaults()
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
						o.SetClassId("workflow.PrimitiveDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDefaultValueWithDefaults()
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
									o.SetClassId("workflow.DefaultValue")
									if v, ok := l["is_value_set"]; ok {
										{
											x := (v.(bool))
											o.SetIsValueSet(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDefault(x)
								}
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["display_meta"]; ok {
							{
								p := make([]models.WorkflowDisplayMeta, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDisplayMetaWithDefaults()
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
									o.SetClassId("workflow.DisplayMeta")
									if v, ok := l["inventory_selector"]; ok {
										{
											x := (v.(bool))
											o.SetInventorySelector(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["widget_type"]; ok {
										{
											x := (v.(string))
											o.SetWidgetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDisplayMeta(x)
								}
							}
						}
						if v, ok := l["input_parameters"]; ok {
							{
								x := v.(map[string]interface{})
								o.SetInputParameters(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
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
						if v, ok := l["properties"]; ok {
							{
								p := make([]models.WorkflowPrimitiveDataProperty, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowPrimitiveDataPropertyWithDefaults()
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
									o.SetClassId("workflow.PrimitiveDataProperty")
									if v, ok := l["constraints"]; ok {
										{
											p := make([]models.WorkflowConstraints, 0, 1)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												l := s[i].(map[string]interface{})
												o := models.NewWorkflowConstraintsWithDefaults()
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
												o.SetClassId("workflow.Constraints")
												if v, ok := l["enum_list"]; ok {
													{
														x := make([]models.WorkflowEnumEntry, 0)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															o := models.NewWorkflowEnumEntryWithDefaults()
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
															o.SetClassId("workflow.EnumEntry")
															if v, ok := l["label"]; ok {
																{
																	x := (v.(string))
																	o.SetLabel(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
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
															o.SetEnumList(x)
														}
													}
												}
												if v, ok := l["max"]; ok {
													{
														x := v.(float64)
														o.SetMax(x)
													}
												}
												if v, ok := l["min"]; ok {
													{
														x := v.(float64)
														o.SetMin(x)
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["regex"]; ok {
													{
														x := (v.(string))
														o.SetRegex(x)
													}
												}
												p = append(p, *o)
											}
											if len(p) > 0 {
												x := p[0]
												o.SetConstraints(x)
											}
										}
									}
									if v, ok := l["inventory_selector"]; ok {
										{
											x := make([]models.WorkflowMoReferenceProperty, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewWorkflowMoReferencePropertyWithDefaults()
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
												o.SetClassId("workflow.MoReferenceProperty")
												if v, ok := l["display_attributes"]; ok {
													{
														x := make([]string, 0)
														y := reflect.ValueOf(v)
														for i := 0; i < y.Len(); i++ {
															x = append(x, y.Index(i).Interface().(string))
														}
														if len(x) > 0 {
															o.SetDisplayAttributes(x)
														}
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
												if v, ok := l["value_attribute"]; ok {
													{
														x := (v.(string))
														o.SetValueAttribute(x)
													}
												}
												x = append(x, *o)
											}
											if len(x) > 0 {
												o.SetInventorySelector(x)
											}
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["secure"]; ok {
										{
											x := (v.(bool))
											o.SetSecure(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetProperties(x)
								}
							}
						}
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetType(x)
					}
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPlaceholders(x)
		}
	}

	if v, ok := d.GetOkExists("supported"); ok {
		x := v.(bool)
		o.SetSupported(x)
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

	r := conn.ApiClient.OsApi.CreateOsConfigurationFile(conn.ctx).OsConfigurationFile(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceOsConfigurationFileRead(c, d, meta)...)
}

func resourceOsConfigurationFileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.OsApi.GetOsConfigurationFileByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "OsConfigurationFile object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("catalog", flattenMapOsCatalogRelationship(s.GetCatalog(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Catalog in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("distributions", flattenListHclOperatingSystemRelationship(s.GetDistributions(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Distributions in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("file_content", (s.GetFileContent())); err != nil {
		return diag.Errorf("error occurred while setting property FileContent in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("internal", (s.GetInternal())); err != nil {
		return diag.Errorf("error occurred while setting property Internal in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("placeholders", flattenListOsPlaceHolder(s.GetPlaceholders(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Placeholders in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("supported", (s.GetSupported())); err != nil {
		return diag.Errorf("error occurred while setting property Supported in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in OsConfigurationFile object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceOsConfigurationFileUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsConfigurationFile{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.OsCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOsCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("os.ConfigurationFile")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("distributions") {
		v := d.Get("distributions")
		x := make([]models.HclOperatingSystemRelationship, 0)
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
			x = append(x, models.MoMoRefAsHclOperatingSystemRelationship(o))
		}
		if len(x) > 0 {
			o.SetDistributions(x)
		}
	}

	if d.HasChange("file_content") {
		v := d.Get("file_content")
		x := (v.(string))
		o.SetFileContent(x)
	}

	if d.HasChange("internal") {
		v := d.Get("internal")
		x := (v.(bool))
		o.SetInternal(x)
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

	o.SetObjectType("os.ConfigurationFile")

	if d.HasChange("placeholders") {
		v := d.Get("placeholders")
		x := make([]models.OsPlaceHolder, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OsPlaceHolder{}
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
			o.SetClassId("os.PlaceHolder")
			if v, ok := l["is_value_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsValueSet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					p := make([]models.WorkflowPrimitiveDataType, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewWorkflowPrimitiveDataTypeWithDefaults()
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
						o.SetClassId("workflow.PrimitiveDataType")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDefaultValueWithDefaults()
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
									o.SetClassId("workflow.DefaultValue")
									if v, ok := l["is_value_set"]; ok {
										{
											x := (v.(bool))
											o.SetIsValueSet(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := v.(map[string]interface{})
											o.SetValue(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDefault(x)
								}
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["display_meta"]; ok {
							{
								p := make([]models.WorkflowDisplayMeta, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDisplayMetaWithDefaults()
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
									o.SetClassId("workflow.DisplayMeta")
									if v, ok := l["inventory_selector"]; ok {
										{
											x := (v.(bool))
											o.SetInventorySelector(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["widget_type"]; ok {
										{
											x := (v.(string))
											o.SetWidgetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDisplayMeta(x)
								}
							}
						}
						if v, ok := l["input_parameters"]; ok {
							{
								x := v.(map[string]interface{})
								o.SetInputParameters(x)
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
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
						if v, ok := l["properties"]; ok {
							{
								p := make([]models.WorkflowPrimitiveDataProperty, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowPrimitiveDataPropertyWithDefaults()
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
									o.SetClassId("workflow.PrimitiveDataProperty")
									if v, ok := l["constraints"]; ok {
										{
											p := make([]models.WorkflowConstraints, 0, 1)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												l := s[i].(map[string]interface{})
												o := models.NewWorkflowConstraintsWithDefaults()
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
												o.SetClassId("workflow.Constraints")
												if v, ok := l["enum_list"]; ok {
													{
														x := make([]models.WorkflowEnumEntry, 0)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															o := models.NewWorkflowEnumEntryWithDefaults()
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
															o.SetClassId("workflow.EnumEntry")
															if v, ok := l["label"]; ok {
																{
																	x := (v.(string))
																	o.SetLabel(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
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
															o.SetEnumList(x)
														}
													}
												}
												if v, ok := l["max"]; ok {
													{
														x := v.(float64)
														o.SetMax(x)
													}
												}
												if v, ok := l["min"]; ok {
													{
														x := v.(float64)
														o.SetMin(x)
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["regex"]; ok {
													{
														x := (v.(string))
														o.SetRegex(x)
													}
												}
												p = append(p, *o)
											}
											if len(p) > 0 {
												x := p[0]
												o.SetConstraints(x)
											}
										}
									}
									if v, ok := l["inventory_selector"]; ok {
										{
											x := make([]models.WorkflowMoReferenceProperty, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewWorkflowMoReferencePropertyWithDefaults()
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
												o.SetClassId("workflow.MoReferenceProperty")
												if v, ok := l["display_attributes"]; ok {
													{
														x := make([]string, 0)
														y := reflect.ValueOf(v)
														for i := 0; i < y.Len(); i++ {
															x = append(x, y.Index(i).Interface().(string))
														}
														if len(x) > 0 {
															o.SetDisplayAttributes(x)
														}
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
												if v, ok := l["value_attribute"]; ok {
													{
														x := (v.(string))
														o.SetValueAttribute(x)
													}
												}
												x = append(x, *o)
											}
											if len(x) > 0 {
												o.SetInventorySelector(x)
											}
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["secure"]; ok {
										{
											x := (v.(bool))
											o.SetSecure(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetProperties(x)
								}
							}
						}
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetType(x)
					}
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := v.(map[string]interface{})
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPlaceholders(x)
		}
	}

	if d.HasChange("supported") {
		v := d.Get("supported")
		x := (v.(bool))
		o.SetSupported(x)
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
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	r := conn.ApiClient.OsApi.UpdateOsConfigurationFile(conn.ctx, d.Id()).OsConfigurationFile(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceOsConfigurationFileRead(c, d, meta)...)
}

func resourceOsConfigurationFileDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.OsApi.DeleteOsConfigurationFile(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting OsConfigurationFile object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
