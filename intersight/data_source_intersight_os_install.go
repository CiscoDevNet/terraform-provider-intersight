package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsInstall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsInstallRead,
		Schema: map[string]*schema.Schema{
			"additional_parameters": {
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_value_set": {
							Description: "Flag to indicate if value is set. Value will be used to check if any edit.",
							Type:        schema.TypeBool,
							Optional:    true,
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
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
										Computed: true,
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
												},
												"inventory_selector": {
													Description: "Inventory selector specified for primitive data property should be used in Intersight User Interface.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"widget_type": {
													Description: "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
										Computed: true,
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
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
																		},
																		"label": {
																			Description: "Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_).",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
																Computed: true,
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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
													Computed: true,
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
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
													Computed: true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"secure": {
													Description: "Intersight supports secure properties as task input/output. The values of\nthese properties are encrypted and stored in Intersight.\nThis flag marks the property to be secure when it is set to true.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"type": {
													Description: "Specify the enum type for primitive data type.\n* `string` - Enum to specify a string data type.\n* `integer` - Enum to specify an integer32 data type.\n* `float` - Enum to specify a float64 data type.\n* `boolean` - Enum to specify a boolean data type.\n* `json` - Enum to specify a json data type.\n* `enum` - Enum to specify a enum data type which is a list of pre-defined strings.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
										Computed: true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
							Computed: true,
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
				Computed: true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"answers": {
				Description: "Answers provided by user for the unattended OS installation.",
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
						"answer_file": {
							Description: "If the source of the answers is a static file, the content of the file is stored as value\nin this property.\nThe value is mandatory only when the 'Source' property has been set to 'File'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"hostname": {
							Description: "Hostname to be configured for the server in the OS.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_config_type": {
							Description: "IP configuration type. Values are Static or Dynamic configuration of IP.\nIn case of static IP configuration, IP address, gateway and other details need\nto be populated. In case of dynamic the IP configuration is obtained dynamically\nfrom DHCP.\n* `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway.\n* `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_configuration": {
							Description: "In case of static IP configuration, IP address, netmask and gateway details are\nprovided.",
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
							Computed: true,
						},
						"is_answer_file_set": {
							Description: "Indicates whether the value of the 'answerFile' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"is_root_password_crypted": {
							Description: "Enable to indicate Root Password provided is encrypted.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"is_root_password_set": {
							Description: "Indicates whether the value of the 'rootPassword' property has been set.",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
						"nameserver": {
							Description: "IP address of the name server to be configured in the OS.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"product_key": {
							Description: "The product key to be used for a specific version of Windows installation.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"root_password": {
							Description: "Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.\nIntersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.\nFor more details on encrypting passwords, see Help Center.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_source": {
							Description: "Answer values can be provided from three sources - Embedded in OS image, static file,\nor as placeholder values for an answer file template.\nSource of the answers is given as value, Embedded/File/Template.\n'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'\noption indicates that the answers are provided as a file. 'Template' indicates that the\nplaceholders in the selected os.ConfigurationFile MO are replaced with values provided\nas os.Answers MO.\n* `None` - Indicates that answers is not sent and values must be populated from Install Template.  \n* `Embedded` - Indicates that the answer file is embedded within OS image.\n* `File` - Indicates that the answer file is a static content that has all thevalues populated.\n* `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"configuration_file": {
				Description: "A reference to a osConfigurationFile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"description": {
				Description: "User provided description about the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image": {
				Description: "A reference to a softwarerepositoryOperatingSystemFile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"install_method": {
				Description: "The install method to be used for OS installation - vMedia, iPXE. \nOnly vMedia is supported as of now.\n* `vMedia` - OS image is mounted as vMedia in target server for OS installation.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"install_target": {
				Description: "Install Target on which Operating system is installed.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The name of the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"operating_system_parameters": {
				Description: "Parameters specific to selected OS.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"osdu_image": {
				Description: "A reference to a firmwareServerConfigurationUtilityDistributable resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"server": {
				Description: "A reference to a computePhysical resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
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
			"workflow_info": {
				Description: "A reference to a workflowWorkflowInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
		},
	}
}

func dataSourceOsInstallRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsInstall{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("install_method"); ok {
		x := (v.(string))
		o.SetInstallMethod(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of OsInstall object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.OsApi.GetOsInstallList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching OsInstall: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for OsInstall list: %s", err.Error())
	}
	var s = &models.OsInstallList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to OsInstall list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for OsInstall did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.OsInstall{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}

			if err := d.Set("additional_parameters", flattenListOsPlaceHolder(s.GetAdditionalParameters(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalParameters: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("answers", flattenMapOsAnswers(s.GetAnswers(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Answers: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("configuration_file", flattenMapOsConfigurationFileRelationship(s.GetConfigurationFile(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ConfigurationFile: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}

			if err := d.Set("image", flattenMapSoftwarerepositoryOperatingSystemFileRelationship(s.GetImage(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Image: %s", err.Error())
			}
			if err := d.Set("install_method", (s.GetInstallMethod())); err != nil {
				return diag.Errorf("error occurred while setting property InstallMethod: %s", err.Error())
			}

			if err := d.Set("install_target", flattenMapOsInstallTarget(s.GetInstallTarget(), d)); err != nil {
				return diag.Errorf("error occurred while setting property InstallTarget: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("operating_system_parameters", flattenMapOsOperatingSystemParameters(s.GetOperatingSystemParameters(), d)); err != nil {
				return diag.Errorf("error occurred while setting property OperatingSystemParameters: %s", err.Error())
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}

			if err := d.Set("osdu_image", flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(s.GetOsduImage(), d)); err != nil {
				return diag.Errorf("error occurred while setting property OsduImage: %s", err.Error())
			}

			if err := d.Set("server", flattenMapComputePhysicalRelationship(s.GetServer(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Server: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("workflow_info", flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowInfo: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
