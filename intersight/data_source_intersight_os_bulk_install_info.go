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

func dataSourceOsBulkInstallInfo() *schema.Resource {
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
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"file_content": {
			Description: "The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections.\nThe first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holds\nparameters which are specific to each server level data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"global_config": {
			Description: "The global parameter of the CSV file.",
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
					"configuration_file_name": {
						Description: "Name of the Configuration file.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configuration_source": {
						Description: "Configuration source for the OS Installation.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"install_method": {
						Description: "The install method to be used for OS installation - vMedia, iPXE.\nOnly vMedia is supported as of now.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"install_target_type": {
						Description: "The Prefill install Target Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
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
								},
							},
						},
					},
					"os_image_name": {
						Description: "The Operating System Image name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"scu_image_name": {
						Description: "The name of the Server Configuration Utilities Image.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"windows_edition": {
						Description: "The Windows OS edition, this property required only for Windows server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"is_file_content_set": {
			Description: "Indicates whether the value of the 'fileContent' property has been set.",
			Type:        schema.TypeBool,
			Optional:    true,
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
			Description: "The name of the CSV file, which holds the OS install parameters.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_state": {
			Description: "Denotes if the operating is pending, in_progress, completed_ok, completed_error.\n* `Pending` - The initial value of the OperStatus.\n* `InProgress` - The OperStatus value will be InProgress during execution.\n* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.\n* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.\n* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"os_image": {
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
		"scu_image": {
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
		"server_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
														"is_value_set": {
															Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
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
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"widget_type": {
															Description: "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
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
																					Description: "Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_) and must be at least 2 characters.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"value": {
																					Description: "Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_).",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																			},
																		},
																	},
																	"max": {
																		Description: "Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.",
																		Type:        schema.TypeFloat,
																		Optional:    true,
																	},
																	"min": {
																		Description: "Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.",
																		Type:        schema.TypeFloat,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"regex": {
																		Description: "When the parameter is a string this regular expression is used to ensure the value is valid.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
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
																	},
																	"selector": {
																		Description: "Field to hold an Intersight API along with an optional filter to narrow down the search options.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"selector_property": {
																		Description: "Selector properties to define HTTP method and 'body' in case of upsert operation.",
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
																				"body": {
																					Description: "Content of the request body to send for POST request.",
																					Type:        schema.TypeMap,
																					Elem: &schema.Schema{
																						Type: schema.TypeString,
																					}, Optional: true,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"method": {
																					Description: "The HTTP method to be used.\n* `GET` - The HTTP GET method requests a representation of the specified resource.\n* `POST` - The HTTP POST method sends data to the server.",
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
																	"value_attribute": {
																		Description: "A property from the Intersight object, value of which can be used as value for referenced input definition.",
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
											},
											"required": {
												Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
												Type:        schema.TypeBool,
												Optional:    true,
											},
										},
									},
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
											},
										},
									},
								},
								"is_answer_file_set": {
									Description: "Indicates whether the value of the 'answerFile' property has been set.",
									Type:        schema.TypeBool,
									Optional:    true,
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
								},
								"nameserver": {
									Description: "IP address of the name server to be configured in the OS.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"network_device": {
									Description: "Network Device where the IP address must be configured. Network Interface names and MAC address are supported.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
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
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"error_msgs": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"install_target": {
						Description: "The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
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
								},
							},
						},
					},
					"processed_install_target": {
						Description: "The target in which OS installation triggered, this is populated after processing the given data.",
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
								},
							},
						},
					},
					"serial_number": {
						Description: "The Serial Number of the server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"servers": {
			Description: "An array of relationships to computePhysical resources.",
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
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
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
		"validation_infos": {
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
					"error_msg": {
						Description: "Validation error message.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "The status of the validation step.\n* `NotValidated` - The validation not started.\n* `Valid` - The step status marked as valid when respective step validation condition is successful.\n* `Invalid` - The step status marked as invalid when respective step validation condition is failed.\n* `InProgress` - The validation is in progress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"step_name": {
						Description: "The validation step name.\n* `OS Install Schema Validation` - The step to validate the CSV file schema.\n* `OS Image Validation` - The Operating System Image parameter validation step.\n* `SCU Image Validation` - The SCU Image parameter validation step.\n* `Configuration source and file validation` - The Configuration Source and Configuration file validation step.\n* `Server level data validation` - The server level parameters validation.",
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
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"file_content": {
			Description: "The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections.\nThe first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holds\nparameters which are specific to each server level data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"global_config": {
			Description: "The global parameter of the CSV file.",
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
					"configuration_file_name": {
						Description: "Name of the Configuration file.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"configuration_source": {
						Description: "Configuration source for the OS Installation.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"install_method": {
						Description: "The install method to be used for OS installation - vMedia, iPXE.\nOnly vMedia is supported as of now.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"install_target_type": {
						Description: "The Prefill install Target Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
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
								},
							},
						},
					},
					"os_image_name": {
						Description: "The Operating System Image name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"scu_image_name": {
						Description: "The name of the Server Configuration Utilities Image.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"windows_edition": {
						Description: "The Windows OS edition, this property required only for Windows server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"is_file_content_set": {
			Description: "Indicates whether the value of the 'fileContent' property has been set.",
			Type:        schema.TypeBool,
			Optional:    true,
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
			Description: "The name of the CSV file, which holds the OS install parameters.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"oper_state": {
			Description: "Denotes if the operating is pending, in_progress, completed_ok, completed_error.\n* `Pending` - The initial value of the OperStatus.\n* `InProgress` - The OperStatus value will be InProgress during execution.\n* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.\n* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.\n* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"os_image": {
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
		"scu_image": {
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
		"server_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
														"is_value_set": {
															Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
															Type:        schema.TypeBool,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
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
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"widget_type": {
															Description: "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
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
																					Description: "Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_) and must be at least 2 characters.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"value": {
																					Description: "Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_).",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																			},
																		},
																	},
																	"max": {
																		Description: "Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.",
																		Type:        schema.TypeFloat,
																		Optional:    true,
																	},
																	"min": {
																		Description: "Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.",
																		Type:        schema.TypeFloat,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"regex": {
																		Description: "When the parameter is a string this regular expression is used to ensure the value is valid.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
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
																	},
																	"selector": {
																		Description: "Field to hold an Intersight API along with an optional filter to narrow down the search options.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"selector_property": {
																		Description: "Selector properties to define HTTP method and 'body' in case of upsert operation.",
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
																				"body": {
																					Description: "Content of the request body to send for POST request.",
																					Type:        schema.TypeMap,
																					Elem: &schema.Schema{
																						Type: schema.TypeString,
																					}, Optional: true,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"method": {
																					Description: "The HTTP method to be used.\n* `GET` - The HTTP GET method requests a representation of the specified resource.\n* `POST` - The HTTP POST method sends data to the server.",
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
																	"value_attribute": {
																		Description: "A property from the Intersight object, value of which can be used as value for referenced input definition.",
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
											},
											"required": {
												Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
												Type:        schema.TypeBool,
												Optional:    true,
											},
										},
									},
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
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
											},
										},
									},
								},
								"is_answer_file_set": {
									Description: "Indicates whether the value of the 'answerFile' property has been set.",
									Type:        schema.TypeBool,
									Optional:    true,
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
								},
								"nameserver": {
									Description: "IP address of the name server to be configured in the OS.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"network_device": {
									Description: "Network Device where the IP address must be configured. Network Interface names and MAC address are supported.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
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
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"error_msgs": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"install_target": {
						Description: "The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
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
								},
							},
						},
					},
					"processed_install_target": {
						Description: "The target in which OS installation triggered, this is populated after processing the given data.",
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
								},
							},
						},
					},
					"serial_number": {
						Description: "The Serial Number of the server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"servers": {
			Description: "An array of relationships to computePhysical resources.",
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
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
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
		"validation_infos": {
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
					"error_msg": {
						Description: "Validation error message.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"status": {
						Description: "The status of the validation step.\n* `NotValidated` - The validation not started.\n* `Valid` - The step status marked as valid when respective step validation condition is successful.\n* `Invalid` - The step status marked as invalid when respective step validation condition is failed.\n* `InProgress` - The validation is in progress.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"step_name": {
						Description: "The validation step name.\n* `OS Install Schema Validation` - The step to validate the CSV file schema.\n* `OS Image Validation` - The Operating System Image parameter validation step.\n* `SCU Image Validation` - The SCU Image parameter validation step.\n* `Configuration source and file validation` - The Configuration Source and Configuration file validation step.\n* `Server level data validation` - The server level parameters validation.",
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
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceOsBulkInstallInfoRead,
		Schema:      model}
}

func dataSourceOsBulkInstallInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsBulkInstallInfo{}
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

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("configuration_file"); ok {
		p := make([]models.OsConfigurationFileRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsOsConfigurationFileRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetConfigurationFile(x)
		}
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("file_content"); ok {
		x := (v.(string))
		o.SetFileContent(x)
	}

	if v, ok := d.GetOk("global_config"); ok {
		p := make([]models.OsGlobalConfig, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.OsGlobalConfig{}
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
			o.SetClassId("os.GlobalConfig")
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
			o.SetGlobalConfig(x)
		}
	}

	if v, ok := d.GetOkExists("is_file_content_set"); ok {
		x := (v.(bool))
		o.SetIsFileContentSet(x)
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

	if v, ok := d.GetOk("oper_state"); ok {
		x := (v.(string))
		o.SetOperState(x)
	}

	if v, ok := d.GetOk("organization"); ok {
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

	if v, ok := d.GetOk("os_image"); ok {
		p := make([]models.SoftwarerepositoryOperatingSystemFileRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsSoftwarerepositoryOperatingSystemFileRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOsImage(x)
		}
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

	if v, ok := d.GetOk("scu_image"); ok {
		p := make([]models.FirmwareServerConfigurationUtilityDistributableRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsFirmwareServerConfigurationUtilityDistributableRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetScuImage(x)
		}
	}

	if v, ok := d.GetOk("server_configs"); ok {
		x := make([]models.OsServerConfig, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OsServerConfig{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_parameters"]; ok {
				{
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
														o.SetValue(v)
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
											o.SetInputParameters(v)
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
																		if y.Index(i).Interface() != nil {
																			x = append(x, y.Index(i).Interface().(string))
																		}
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
															if v, ok := l["selector_property"]; ok {
																{
																	p := make([]models.WorkflowSelectorProperty, 0, 1)
																	s := v.([]interface{})
																	for i := 0; i < len(s); i++ {
																		l := s[i].(map[string]interface{})
																		o := models.NewWorkflowSelectorPropertyWithDefaults()
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
																		if v, ok := l["body"]; ok {
																			{
																				o.SetBody(v)
																			}
																		}
																		o.SetClassId("workflow.SelectorProperty")
																		if v, ok := l["method"]; ok {
																			{
																				x := (v.(string))
																				o.SetMethod(x)
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
																		o.SetSelectorProperty(x)
																	}
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
								o.SetValue(v)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetAdditionalParameters(x)
					}
				}
			}
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
			o.SetClassId("os.ServerConfig")
			if v, ok := l["error_msgs"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetErrorMsgs(x)
					}
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
		o.SetServerConfigs(x)
	}

	if v, ok := d.GetOk("servers"); ok {
		x := make([]models.ComputePhysicalRelationship, 0)
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
			x = append(x, models.MoMoRefAsComputePhysicalRelationship(o))
		}
		o.SetServers(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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

	if v, ok := d.GetOk("validation_infos"); ok {
		x := make([]models.OsValidationInformation, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OsValidationInformation{}
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
			o.SetClassId("os.ValidationInformation")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetValidationInfos(x)
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
		return diag.Errorf("json marshal of OsBulkInstallInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.OsApi.GetOsBulkInstallInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of OsBulkInstallInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of OsBulkInstallInfo: %s", responseErr.Error())
	}
	count := countResponse.OsBulkInstallInfoList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for OsBulkInstallInfo data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var osBulkInstallInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.OsApi.GetOsBulkInstallInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching OsBulkInstallInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching OsBulkInstallInfo: %s", responseErr.Error())
		}
		results := resMo.OsBulkInstallInfoList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["configuration_file"] = flattenMapOsConfigurationFileRelationship(s.GetConfigurationFile(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["global_config"] = flattenMapOsGlobalConfig(s.GetGlobalConfig(), d)
				temp["is_file_content_set"] = (s.GetIsFileContentSet())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["oper_state"] = (s.GetOperState())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["os_image"] = flattenMapSoftwarerepositoryOperatingSystemFileRelationship(s.GetOsImage(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["scu_image"] = flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(s.GetScuImage(), d)

				temp["server_configs"] = flattenListOsServerConfig(s.GetServerConfigs(), d)

				temp["servers"] = flattenListComputePhysicalRelationship(s.GetServers(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["validation_infos"] = flattenListOsValidationInformation(s.GetValidationInfos(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				osBulkInstallInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(osBulkInstallInfoResults))
	if err := d.Set("results", osBulkInstallInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(osBulkInstallInfoResults[0]["moid"].(string))
	return de
}
