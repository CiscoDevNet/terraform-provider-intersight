package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWorkloadBlueprint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkloadBlueprintCreate,
		ReadContext:   resourceWorkloadBlueprintRead,
		UpdateContext: resourceWorkloadBlueprintUpdate,
		DeleteContext: resourceWorkloadBlueprintDelete,
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
			"blueprint_dependencies": {
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
						"blueprint_definition": {
							Description: "The system resolved reference to the blueprint.",
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
						"blueprint_name": {
							Description: "The qualified name of blueprint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"catalog_moid": {
							Description: "Specify the catalog moid that this blueprint belongs too. When catalog moid is not specified then the catalog of the parent blueprint is used first and if no blueprint can be found in that catalog, then the system catalog is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.BlueprintDefinitionDependency",
						},
						"description": {
							Description: "The description of that describes the purpose of including this dependent blueprint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "If any input of the nested blueprint needs to be modified in some form to include data from the parent blueprint, then the input parameters can be specified here.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier for this blueprint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.BlueprintDefinitionDependency",
						},
						"nr_version": {
							Description: "The blueprint version. If not specified, the default version is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"catalog": {
				Description: "A reference to a workflowCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "workload.Blueprint",
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
			"default_version": {
				Description: "The flag to indicate that this is the default version of the blueprint.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"description": {
				Description:  "The description for this blueprint which provides information on what are the pre-requisites to deploy the blueprint and what features are supported on the blueprint.",
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
			"external_meta": {
				Description: "When set to false the blueprint is created for use within internal services.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"generated_object_definition": {
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
							Default:     "blueprint.GeneratedObjectMetadata",
						},
						"conditions": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							}},
						"name": {
							Description:  "Name for the generated object. The name given for reference in later parts of the blueprint definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
							Type:         schema.TypeString,
							ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+([a-zA-Z0-9-_]*[a-zA-Z0-9])*$"), ""), validation.StringLenBetween(1, 92)),
							Optional:     true,
						},
						"object_iteration": {
							Description: "A template for controlling loop behavior when generating multiple instances of the generated object. The resolved value of the template is expected to be an integer which is used to determine the number of iterations in the loop. Each iteration creates a new instance of the generated object.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_source": {
							Description: "The source object from which the object is generated.",
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.GeneratedObjectMetadata",
						},
						"pre_generate_operations": {
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
										Default:     "blueprint.PreGenerateOperation",
									},
									"generated_objects": {
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
													Default:     "blueprint.GeneratedObjectOperationTarget",
												},
												"name": {
													Description: "The name of the generated object metadata. The objects produced from this metadata are the target of the operation.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "blueprint.GeneratedObjectOperationTarget",
												},
												"property_parameters": {
													Description: "A list of key value pairs where key is the property path and value is the template to derive the value. The target of the pre-generation operation is updated with the resolved values of these parameters. This operation is only applicable if the operation is of type \"Update\".",
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
										Default:     "blueprint.PreGenerateOperation",
									},
									"operation": {
										Description:  "The operation to be performed before generating the object.\n* `Update` - The operation is to update the object via an HTTP PATCH operation.\n* `Delete` - The operation is to immediately delete the object via an HTTP DELETE operation.\n* `MarkDeleted` - The object is marked as deleted. This adds a tag to the object indicating that it should be deleted as part of the object cleanup task during workload deployment.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"Update", "Delete", "MarkDeleted"}, false),
										Optional:     true,
										Default:      "Update",
									},
								},
							},
						},
						"property_iteration": {
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
										Default:     "blueprint.PropertyIteration",
									},
									"iteration": {
										Description: "The expression to evaluate for each iteration. This expression can refer to the array input that is being iterated over. The value of this expression will be used to populate the array property in the generated object.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "The name of the array property in the generated object to populate.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "blueprint.PropertyIteration",
									},
								},
							},
						},
						"property_parameters": {
							Description: "A list of key value pairs where key is the property path and value is the template to derive the value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"required": {
							Description: "The flag to indicate if this object must be generated for this blueprint.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
					},
				},
			},
			"input_definition": {
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default": {
							Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
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
										Default:     "workflow.DefaultValue",
									},
									"is_value_set": {
										Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
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
										Default:     "workflow.DefaultValue",
									},
									"override": {
										Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
									"value": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
										Type:        schema.TypeString,
										Optional:    true,
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
										Default:     "workflow.DisplayMeta",
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
										Default:     "workflow.DisplayMeta",
									},
									"widget_type": {
										Description:  "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.\n* `MultiSelect` - Display the widget as a multi-select.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"None", "Radio", "Dropdown", "GridSelector", "DrawerSelector", "MultiSelect"}, false),
										Optional:     true,
										Default:      "None",
									},
								},
							},
						},
						"input_parameters": {
							Description: "JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description:  "Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ), forward slash (/) or an underscore (_). The first and last character in label must be an alphanumeric character.",
							Type:         schema.TypeString,
							ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+[\\sa-zA-Z0-9_'.:/-]{1,92}$"), ""), validation.StringLenBetween(1, 92)),
							Optional:     true,
						},
						"name": {
							Description:  "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
							Type:         schema.TypeString,
							ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+([a-zA-Z0-9-_]*[a-zA-Z0-9])*$"), ""), validation.StringLenBetween(1, 92)),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"required": {
							Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
			"input_operation_metadata": {
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
							Default:     "blueprint.InputOperationType",
						},
						"impact_type": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"None", "Reboot"}, false),
							}},
						"name": {
							Description: "The name of the input as defined within the definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.InputOperationType",
						},
						"operation_modes": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"InitialDeployment", "PostDeployment"}, false),
							}},
					},
				},
			},
			"label": {
				Description:  "A user friendly short name to identify the blueprint. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+[\\sa-zA-Z0-9_.:-]{1,31}$"), ""),
				Optional:     true,
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
				Description:  "The name for this blueprint. You can have multiple versions of the blueprint with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), or an underscore (_).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9_]{0,31}$"), ""),
				Optional:     true,
				ForceNew:     true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "workload.Blueprint",
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
			"platform_type": {
				Description:  "The Intersight platforms supported by this blueprint.\n* `None` - Default value is none, platform is not mentioned.\n* `UnifiedEdge` - The blueprint is created for Unified Edge platform.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"None", "UnifiedEdge"}, false),
				Optional:     true,
				Default:      "None",
			},
			"resource_definition": {
				Description: "The resource constraint definition for the blueprint.",
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
							Default:     "blueprint.ResourceConstraints",
						},
						"input_definition": {
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
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
													Default:     "workflow.DefaultValue",
												},
												"is_value_set": {
													Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
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
													Default:     "workflow.DefaultValue",
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeString,
													Optional:    true,
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
													Default:     "workflow.DisplayMeta",
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
													Default:     "workflow.DisplayMeta",
												},
												"widget_type": {
													Description:  "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.\n* `MultiSelect` - Display the widget as a multi-select.",
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"None", "Radio", "Dropdown", "GridSelector", "DrawerSelector", "MultiSelect"}, false),
													Optional:     true,
													Default:      "None",
												},
											},
										},
									},
									"input_parameters": {
										Description: "JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"label": {
										Description:  "Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ), forward slash (/) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:         schema.TypeString,
										ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+[\\sa-zA-Z0-9_'.:/-]{1,92}$"), ""), validation.StringLenBetween(1, 92)),
										Optional:     true,
									},
									"name": {
										Description:  "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:         schema.TypeString,
										ValidateFunc: validation.All(validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+([a-zA-Z0-9-_]*[a-zA-Z0-9])*$"), ""), validation.StringLenBetween(1, 92)),
										Optional:     true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.ResourceConstraints",
						},
					},
				},
			},
			"service_items": {
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
						"catalog_moid": {
							Description: "Specify the catalog moid that this service item belongs. When catalog moid is not specified then the catalog of the service item is used first and if no service item can be found in that catalog, then the system catalog is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.ServiceItemDefinition",
						},
						"description": {
							Description: "The description of that describes the purpose of including this service item within the blueprint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of blueprint inputs to service item inputs. The mapping for inputs can refer to blueprint inputs or to the generated objects which are part of this blueprint and all nested blueprints. The mapping will be referred using the convention of ${<BlueprintName>.input.<InputName>} or ${<BlueprintName>.generatedObject.<GeneratedObjectName>}.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier for this Service item.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "blueprint.ServiceItemDefinition",
						},
						"service_item_definition_name": {
							Description: "The qualified name of service item that should be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"nr_version": {
							Description: "The service item definition version. If not specified, the default version of the service item is used.",
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
			"validation_information": {
				Description: "The current validation state and associated validation errors when state is invalid. When state is valid, it means all referred dependent blueprints and service items are valid. All referred objects are present in the system and available for cloning.",
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
							Default:     "workflow.ValidationInformation",
						},
						"engine_state": {
							Description: "The state of workflow definition metadata in the workflow engine. The workflow definition must be successfully updated in the engine before workflows can be executed.\n* `NotUpdated` - The workflow and task definition metadata is not yet updated in the workflow engine.\n* `Updating` - The workflow and task definition metadata is in the processing of being updated in the workflow engine.\n* `UpdateFailed` - The workflow and task definition metadata failed to be updated in the workflow engine.\n* `Updated` - The workflow and task definition metadata was updated successfully in the workflow engine.",
							Type:        schema.TypeString,
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
							Default:     "workflow.ValidationInformation",
						},
						"state": {
							Description: "The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default).\n* `NotValidated` - The state when workflow definition has not been validated.\n* `Valid` - The state when workflow definition is valid.\n* `Invalid` - The state when workflow definition is invalid.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"validation_error": {
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
										Default:     "workflow.ValidationError",
									},
									"error_log": {
										Description: "Description of the error.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
											if val != nil {
												warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
											}
											return
										}},
									"field": {
										Description: "When populated this refers to the input or output field within the workflow or task.",
										Type:        schema.TypeString,
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
										Default:     "workflow.ValidationError",
									},
									"task_name": {
										Description: "The task name on which the error is found, when empty the error applies to the top level workflow.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
											if val != nil {
												warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
											}
											return
										}},
									"transition_name": {
										Description: "When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue.",
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
				},
			},
			"nr_version": {
				Description:  "The version of the blueprint to support multiple versions.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(1),
				Optional:     true,
				Default:      1,
				ForceNew:     true,
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

func resourceWorkloadBlueprintCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewWorkloadBlueprintWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("blueprint_dependencies"); ok {
		x := make([]models.BlueprintBlueprintDefinitionDependency, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewBlueprintBlueprintDefinitionDependencyWithDefaults()
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
			if v, ok := l["blueprint_name"]; ok {
				{
					x := (v.(string))
					o.SetBlueprintName(x)
				}
			}
			if v, ok := l["catalog_moid"]; ok {
				{
					x := (v.(string))
					o.SetCatalogMoid(x)
				}
			}
			o.SetClassId("blueprint.BlueprintDefinitionDependency")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["input_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
				}
			}
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetBlueprintDependencies(x)
		}
	}

	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("workload.Blueprint")

	if v, ok := d.GetOkExists("default_version"); ok {
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOkExists("external_meta"); ok {
		x := (v.(bool))
		o.SetExternalMeta(x)
	}

	if v, ok := d.GetOk("generated_object_definition"); ok {
		x := make([]models.BlueprintGeneratedObjectMetadata, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewBlueprintGeneratedObjectMetadataWithDefaults()
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
			o.SetClassId("blueprint.GeneratedObjectMetadata")
			if v, ok := l["conditions"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetConditions(x)
					}
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_iteration"]; ok {
				{
					x := (v.(string))
					o.SetObjectIteration(x)
				}
			}
			if v, ok := l["object_source"]; ok {
				{
					p := make([]models.BlueprintBaseGeneratedObjectSourceMetadata, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewBlueprintBaseGeneratedObjectSourceMetadataWithDefaults()
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
						o.SetClassId("blueprint.BaseGeneratedObjectSourceMetadata")
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
						o.SetObjectSource(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["pre_generate_operations"]; ok {
				{
					x := make([]models.BlueprintPreGenerateOperation, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewBlueprintPreGenerateOperationWithDefaults()
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
						o.SetClassId("blueprint.PreGenerateOperation")
						if v, ok := l["generated_objects"]; ok {
							{
								x := make([]models.BlueprintGeneratedObjectOperationTarget, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewBlueprintGeneratedObjectOperationTargetWithDefaults()
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
									o.SetClassId("blueprint.GeneratedObjectOperationTarget")
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
									if v, ok := l["property_parameters"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetPropertyParameters(x2)
											}
										}
									}
									x = append(x, *o)
								}
								if len(x) > 0 {
									o.SetGeneratedObjects(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["operation"]; ok {
							{
								x := (v.(string))
								o.SetOperation(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetPreGenerateOperations(x)
					}
				}
			}
			if v, ok := l["property_iteration"]; ok {
				{
					x := make([]models.BlueprintPropertyIteration, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewBlueprintPropertyIterationWithDefaults()
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
						o.SetClassId("blueprint.PropertyIteration")
						if v, ok := l["iteration"]; ok {
							{
								x := (v.(string))
								o.SetIteration(x)
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
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetPropertyIteration(x)
					}
				}
			}
			if v, ok := l["property_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetPropertyParameters(x2)
					}
				}
			}
			if v, ok := l["required"]; ok {
				{
					x := (v.(bool))
					o.SetRequired(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetGeneratedObjectDefinition(x)
		}
	}

	if v, ok := d.GetOk("input_definition"); ok {
		x := make([]models.WorkflowBaseDataType, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowBaseDataTypeWithDefaults()
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
			o.SetClassId("workflow.BaseDataType")
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
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetValue(x2)
								}
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
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
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
			if v, ok := l["required"]; ok {
				{
					x := (v.(bool))
					o.SetRequired(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetInputDefinition(x)
		}
	}

	if v, ok := d.GetOk("input_operation_metadata"); ok {
		x := make([]models.BlueprintInputOperationType, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewBlueprintInputOperationTypeWithDefaults()
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
			o.SetClassId("blueprint.InputOperationType")
			if v, ok := l["impact_type"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetImpactType(x)
					}
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
			if v, ok := l["operation_modes"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetOperationModes(x)
					}
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetInputOperationMetadata(x)
		}
	}

	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("workload.Blueprint")

	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}

	if v, ok := d.GetOk("resource_definition"); ok {
		p := make([]models.BlueprintResourceConstraints, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewBlueprintResourceConstraintsWithDefaults()
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
			o.SetClassId("blueprint.ResourceConstraints")
			if v, ok := l["input_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
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
						o.SetClassId("workflow.BaseDataType")
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
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetValue(x2)
											}
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
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetInputParameters(x2)
								}
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
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInputDefinition(x)
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
			o.SetResourceDefinition(x)
		}
	}

	if v, ok := d.GetOk("service_items"); ok {
		x := make([]models.BlueprintServiceItemDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewBlueprintServiceItemDefinitionWithDefaults()
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
			if v, ok := l["catalog_moid"]; ok {
				{
					x := (v.(string))
					o.SetCatalogMoid(x)
				}
			}
			o.SetClassId("blueprint.ServiceItemDefinition")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["input_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
				}
			}
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
			if v, ok := l["service_item_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetServiceItemDefinitionName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetServiceItems(x)
		}
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

	if v, ok := d.GetOkExists("nr_version"); ok {
		x := int64(v.(int))
		o.SetVersion(x)
	}

	r := conn.ApiClient.WorkloadApi.CreateWorkloadBlueprint(conn.ctx).WorkloadBlueprint(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating WorkloadBlueprint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating WorkloadBlueprint: %s", responseErr.Error())
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
	return append(de, resourceWorkloadBlueprintRead(c, d, meta)...)
}

func resourceWorkloadBlueprintRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	if len(d.Id()) == 0 {
		return de
	}
	conn := meta.(*Config)
	r := conn.ApiClient.WorkloadApi.GetWorkloadBlueprintByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "WorkloadBlueprint object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkloadBlueprint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching WorkloadBlueprint: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("blueprint_dependencies", flattenListBlueprintBlueprintDefinitionDependency(s.GetBlueprintDependencies(), d)); err != nil {
		return diag.Errorf("error occurred while setting property BlueprintDependencies in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("catalog", flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Catalog in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("default_version", (s.GetDefaultVersion())); err != nil {
		return diag.Errorf("error occurred while setting property DefaultVersion in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("external_meta", (s.GetExternalMeta())); err != nil {
		return diag.Errorf("error occurred while setting property ExternalMeta in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("generated_object_definition", flattenListBlueprintGeneratedObjectMetadata(s.GetGeneratedObjectDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property GeneratedObjectDefinition in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("input_definition", flattenListWorkflowBaseDataType(s.GetInputDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property InputDefinition in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("input_operation_metadata", flattenListBlueprintInputOperationType(s.GetInputOperationMetadata(), d)); err != nil {
		return diag.Errorf("error occurred while setting property InputOperationMetadata in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("label", (s.GetLabel())); err != nil {
		return diag.Errorf("error occurred while setting property Label in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("platform_type", (s.GetPlatformType())); err != nil {
		return diag.Errorf("error occurred while setting property PlatformType in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("resource_definition", flattenMapBlueprintResourceConstraints(s.GetResourceDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ResourceDefinition in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("service_items", flattenListBlueprintServiceItemDefinition(s.GetServiceItems(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ServiceItems in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("validation_information", flattenMapWorkflowValidationInformation(s.GetValidationInformation(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ValidationInformation in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return diag.Errorf("error occurred while setting property Version in WorkloadBlueprint object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in WorkloadBlueprint object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceWorkloadBlueprintUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkloadBlueprint{}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("blueprint_dependencies") {
		v := d.Get("blueprint_dependencies")
		x := make([]models.BlueprintBlueprintDefinitionDependency, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.BlueprintBlueprintDefinitionDependency{}
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
			if v, ok := l["blueprint_name"]; ok {
				{
					x := (v.(string))
					o.SetBlueprintName(x)
				}
			}
			if v, ok := l["catalog_moid"]; ok {
				{
					x := (v.(string))
					o.SetCatalogMoid(x)
				}
			}
			o.SetClassId("blueprint.BlueprintDefinitionDependency")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["input_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
				}
			}
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		o.SetBlueprintDependencies(x)
	}

	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("workload.Blueprint")

	if d.HasChange("default_version") {
		v := d.Get("default_version")
		x := (v.(bool))
		o.SetDefaultVersion(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("external_meta") {
		v := d.Get("external_meta")
		x := (v.(bool))
		o.SetExternalMeta(x)
	}

	if d.HasChange("generated_object_definition") {
		v := d.Get("generated_object_definition")
		x := make([]models.BlueprintGeneratedObjectMetadata, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.BlueprintGeneratedObjectMetadata{}
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
			o.SetClassId("blueprint.GeneratedObjectMetadata")
			if v, ok := l["conditions"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetConditions(x)
					}
				}
			}
			if v, ok := l["name"]; ok {
				{
					x := (v.(string))
					o.SetName(x)
				}
			}
			if v, ok := l["object_iteration"]; ok {
				{
					x := (v.(string))
					o.SetObjectIteration(x)
				}
			}
			if v, ok := l["object_source"]; ok {
				{
					p := make([]models.BlueprintBaseGeneratedObjectSourceMetadata, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewBlueprintBaseGeneratedObjectSourceMetadataWithDefaults()
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
						o.SetClassId("blueprint.BaseGeneratedObjectSourceMetadata")
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
						o.SetObjectSource(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["pre_generate_operations"]; ok {
				{
					x := make([]models.BlueprintPreGenerateOperation, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewBlueprintPreGenerateOperationWithDefaults()
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
						o.SetClassId("blueprint.PreGenerateOperation")
						if v, ok := l["generated_objects"]; ok {
							{
								x := make([]models.BlueprintGeneratedObjectOperationTarget, 0)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									o := models.NewBlueprintGeneratedObjectOperationTargetWithDefaults()
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
									o.SetClassId("blueprint.GeneratedObjectOperationTarget")
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
									if v, ok := l["property_parameters"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetPropertyParameters(x2)
											}
										}
									}
									x = append(x, *o)
								}
								if len(x) > 0 {
									o.SetGeneratedObjects(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["operation"]; ok {
							{
								x := (v.(string))
								o.SetOperation(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetPreGenerateOperations(x)
					}
				}
			}
			if v, ok := l["property_iteration"]; ok {
				{
					x := make([]models.BlueprintPropertyIteration, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewBlueprintPropertyIterationWithDefaults()
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
						o.SetClassId("blueprint.PropertyIteration")
						if v, ok := l["iteration"]; ok {
							{
								x := (v.(string))
								o.SetIteration(x)
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
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetPropertyIteration(x)
					}
				}
			}
			if v, ok := l["property_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetPropertyParameters(x2)
					}
				}
			}
			if v, ok := l["required"]; ok {
				{
					x := (v.(bool))
					o.SetRequired(x)
				}
			}
			x = append(x, *o)
		}
		o.SetGeneratedObjectDefinition(x)
	}

	if d.HasChange("input_definition") {
		v := d.Get("input_definition")
		x := make([]models.WorkflowBaseDataType, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowBaseDataType{}
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
			o.SetClassId("workflow.BaseDataType")
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
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetValue(x2)
								}
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
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
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
			if v, ok := l["required"]; ok {
				{
					x := (v.(bool))
					o.SetRequired(x)
				}
			}
			x = append(x, *o)
		}
		o.SetInputDefinition(x)
	}

	if d.HasChange("input_operation_metadata") {
		v := d.Get("input_operation_metadata")
		x := make([]models.BlueprintInputOperationType, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.BlueprintInputOperationType{}
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
			o.SetClassId("blueprint.InputOperationType")
			if v, ok := l["impact_type"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetImpactType(x)
					}
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
			if v, ok := l["operation_modes"]; ok {
				{
					x := make([]string, 0)
					y := reflect.ValueOf(v)
					for i := 0; i < y.Len(); i++ {
						if y.Index(i).Interface() != nil {
							x = append(x, y.Index(i).Interface().(string))
						}
					}
					if len(x) > 0 {
						o.SetOperationModes(x)
					}
				}
			}
			x = append(x, *o)
		}
		o.SetInputOperationMetadata(x)
	}

	if d.HasChange("label") {
		v := d.Get("label")
		x := (v.(string))
		o.SetLabel(x)
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

	o.SetObjectType("workload.Blueprint")

	if d.HasChange("platform_type") {
		v := d.Get("platform_type")
		x := (v.(string))
		o.SetPlatformType(x)
	}

	if d.HasChange("resource_definition") {
		v := d.Get("resource_definition")
		p := make([]models.BlueprintResourceConstraints, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.BlueprintResourceConstraints{}
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
			o.SetClassId("blueprint.ResourceConstraints")
			if v, ok := l["input_definition"]; ok {
				{
					x := make([]models.WorkflowBaseDataType, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewWorkflowBaseDataTypeWithDefaults()
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
						o.SetClassId("workflow.BaseDataType")
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
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetValue(x2)
											}
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
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetInputParameters(x2)
								}
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
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInputDefinition(x)
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
			o.SetResourceDefinition(x)
		}
	}

	if d.HasChange("service_items") {
		v := d.Get("service_items")
		x := make([]models.BlueprintServiceItemDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.BlueprintServiceItemDefinition{}
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
			if v, ok := l["catalog_moid"]; ok {
				{
					x := (v.(string))
					o.SetCatalogMoid(x)
				}
			}
			o.SetClassId("blueprint.ServiceItemDefinition")
			if v, ok := l["description"]; ok {
				{
					x := (v.(string))
					o.SetDescription(x)
				}
			}
			if v, ok := l["input_parameters"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetInputParameters(x2)
					}
				}
			}
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
			if v, ok := l["service_item_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetServiceItemDefinitionName(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			x = append(x, *o)
		}
		o.SetServiceItems(x)
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

	if d.HasChange("nr_version") {
		v := d.Get("nr_version")
		x := int64(v.(int))
		o.SetVersion(x)
	}

	r := conn.ApiClient.WorkloadApi.UpdateWorkloadBlueprint(conn.ctx, d.Id()).WorkloadBlueprint(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating WorkloadBlueprint: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating WorkloadBlueprint: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceWorkloadBlueprintRead(c, d, meta)...)
}

func resourceWorkloadBlueprintDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.WorkloadApi.DeleteWorkloadBlueprint(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "WorkloadBlueprintDelete: WorkloadBlueprint object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting WorkloadBlueprint object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting WorkloadBlueprint object: %s", deleteErr.Error())
	}
	return de
}
