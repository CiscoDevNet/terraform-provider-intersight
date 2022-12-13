package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWorkflowServiceItemActionDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowServiceItemActionDefinitionCreate,
		ReadContext:   resourceWorkflowServiceItemActionDefinitionRead,
		UpdateContext: resourceWorkflowServiceItemActionDefinitionUpdate,
		DeleteContext: resourceWorkflowServiceItemActionDefinitionDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CustomizeTagDiff,
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
			"action_type": {
				Description:  "Type of actionDefinition which decides on how to trigger the action.\n* `External` - External actions definition can be triggered by enduser to perform actions on the service item. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again.\n* `Internal` - Internal action definition is used to trigger periodic actions on the service item instance.\n* `Repetitive` - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created service item.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"External", "Internal", "Repetitive"}, false),
				Optional:     true,
				Default:      "External",
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"allowed_instance_states": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"NotCreated", "InProgress", "Failed", "Okay", "Decommissioned"}, false),
				}},
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "workflow.ServiceItemActionDefinition",
			},
			"core_workflows": {
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
							Description: "Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"description": {
							Description: "The description of this workflow instance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of ActionDefinition inputDefinition to workflow definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description:  "The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,64}$"), ""),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"nr_version": {
							Description: "The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"workflow_definition_name": {
							Description: "The qualified name of workflow that should be executed.",
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
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"description": {
				Description: "The description for this action which provides information on what are the pre-requisites to use this action on the service item and what features are supported by this action.",
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
										Description:  "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"None", "Radio", "Dropdown", "GridSelector", "DrawerSelector"}, false),
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
			"label": {
				Description:  "A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]+[\\sa-zA-Z0-9_.:-]{1,92}$"), ""),
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
				Description:  "The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a service item definition.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.:-]{1,64}$"), ""),
				Optional:     true,
				ForceNew:     true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "workflow.ServiceItemActionDefinition",
			},
			"output_parameters": {
				Description: "The output mappings from workflows in the action definition to the service item output definition. Any output from core or post-core workflow can be mapped to service item output definition. The output can be referred using the name of the workflow definition and the output name in the following format '${<ServiceItemActionWorkflowDefinition.Name>.output.<outputName>'.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"periodicity": {
				Description:  "Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity.",
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 604800),
				Optional:     true,
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
			"post_core_workflows": {
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
							Description: "Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"description": {
							Description: "The description of this workflow instance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of ActionDefinition inputDefinition to workflow definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description:  "The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,64}$"), ""),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"nr_version": {
							Description: "The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"workflow_definition_name": {
							Description: "The qualified name of workflow that should be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"pre_core_workflows": {
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
							Description: "Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"description": {
							Description: "The description of this workflow instance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of ActionDefinition inputDefinition to workflow definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description:  "The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,64}$"), ""),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"nr_version": {
							Description: "The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"workflow_definition_name": {
							Description: "The qualified name of workflow that should be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"restrict_on_private_appliance": {
				Description: "The flag to indicate that action is restricted on a Private Virtual Appliance.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"service_item_definition": {
				Description: "A reference to a workflowServiceItemDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"stop_workflows": {
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
							Description: "Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"description": {
							Description: "The description of this workflow instance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of ActionDefinition inputDefinition to workflow definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description:  "The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,64}$"), ""),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"nr_version": {
							Description: "The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"workflow_definition_name": {
							Description: "The qualified name of workflow that should be executed.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
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
			"validation_information": {
				Description: "The current validation state and associated validation errors when state is invalid.",
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
			"validation_workflows": {
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
							Description: "Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"description": {
							Description: "The description of this workflow instance.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"input_parameters": {
							Description: "Capture the mapping of ActionDefinition inputDefinition to workflow definition.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"label": {
							Description: "A user defined label identifier of the workflow used for UI display.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description:  "The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,64}$"), ""),
							Optional:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "workflow.ServiceItemActionWorkflowDefinition",
						},
						"nr_version": {
							Description: "The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"workflow_definition_name": {
							Description: "The qualified name of workflow that should be executed.",
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
			"workflow_definition": {
				Description: "A reference to a workflowWorkflowDefinition resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		},
	}
}

func resourceWorkflowServiceItemActionDefinitionCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewWorkflowServiceItemActionDefinitionWithDefaults()

	if v, ok := d.GetOk("action_type"); ok {
		x := (v.(string))
		o.SetActionType(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("allowed_instance_states"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		if len(x) > 0 {
			o.SetAllowedInstanceStates(x)
		}
	}

	o.SetClassId("workflow.ServiceItemActionDefinition")

	if v, ok := d.GetOk("core_workflows"); ok {
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowServiceItemActionWorkflowDefinitionWithDefaults()
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetCoreWorkflows(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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

	o.SetObjectType("workflow.ServiceItemActionDefinition")

	if v, ok := d.GetOk("output_parameters"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			x2 := x1.(map[string]interface{})
			o.SetOutputParameters(x2)
		}
	}

	if v, ok := d.GetOkExists("periodicity"); ok {
		x := int64(v.(int))
		o.SetPeriodicity(x)
	}

	if v, ok := d.GetOk("post_core_workflows"); ok {
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowServiceItemActionWorkflowDefinitionWithDefaults()
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPostCoreWorkflows(x)
		}
	}

	if v, ok := d.GetOk("pre_core_workflows"); ok {
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowServiceItemActionWorkflowDefinitionWithDefaults()
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPreCoreWorkflows(x)
		}
	}

	if v, ok := d.GetOkExists("restrict_on_private_appliance"); ok {
		x := (v.(bool))
		o.SetRestrictOnPrivateAppliance(x)
	}

	if v, ok := d.GetOk("service_item_definition"); ok {
		p := make([]models.WorkflowServiceItemDefinitionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowServiceItemDefinitionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetServiceItemDefinition(x)
		}
	}

	if v, ok := d.GetOk("stop_workflows"); ok {
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowServiceItemActionWorkflowDefinitionWithDefaults()
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetStopWorkflows(x)
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

	if v, ok := d.GetOk("validation_workflows"); ok {
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewWorkflowServiceItemActionWorkflowDefinitionWithDefaults()
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetValidationWorkflows(x)
		}
	}

	r := conn.ApiClient.WorkflowApi.CreateWorkflowServiceItemActionDefinition(conn.ctx).WorkflowServiceItemActionDefinition(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating WorkflowServiceItemActionDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating WorkflowServiceItemActionDefinition: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceWorkflowServiceItemActionDefinitionRead(c, d, meta)...)
}

func resourceWorkflowServiceItemActionDefinitionRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	r := conn.ApiClient.WorkflowApi.GetWorkflowServiceItemActionDefinitionByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "WorkflowServiceItemActionDefinition object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching WorkflowServiceItemActionDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching WorkflowServiceItemActionDefinition: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("action_type", (s.GetActionType())); err != nil {
		return diag.Errorf("error occurred while setting property ActionType in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("allowed_instance_states", (s.GetAllowedInstanceStates())); err != nil {
		return diag.Errorf("error occurred while setting property AllowedInstanceStates in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("core_workflows", flattenListWorkflowServiceItemActionWorkflowDefinition(s.GetCoreWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property CoreWorkflows in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("input_definition", flattenListWorkflowBaseDataType(s.GetInputDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property InputDefinition in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("label", (s.GetLabel())); err != nil {
		return diag.Errorf("error occurred while setting property Label in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("output_parameters", flattenAdditionalProperties(s.GetOutputParameters())); err != nil {
		return diag.Errorf("error occurred while setting property OutputParameters in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("periodicity", (s.GetPeriodicity())); err != nil {
		return diag.Errorf("error occurred while setting property Periodicity in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("post_core_workflows", flattenListWorkflowServiceItemActionWorkflowDefinition(s.GetPostCoreWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PostCoreWorkflows in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("pre_core_workflows", flattenListWorkflowServiceItemActionWorkflowDefinition(s.GetPreCoreWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PreCoreWorkflows in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("restrict_on_private_appliance", (s.GetRestrictOnPrivateAppliance())); err != nil {
		return diag.Errorf("error occurred while setting property RestrictOnPrivateAppliance in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("service_item_definition", flattenMapWorkflowServiceItemDefinitionRelationship(s.GetServiceItemDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ServiceItemDefinition in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("stop_workflows", flattenListWorkflowServiceItemActionWorkflowDefinition(s.GetStopWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property StopWorkflows in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("validation_information", flattenMapWorkflowValidationInformation(s.GetValidationInformation(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ValidationInformation in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("validation_workflows", flattenListWorkflowServiceItemActionWorkflowDefinition(s.GetValidationWorkflows(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ValidationWorkflows in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	if err := d.Set("workflow_definition", flattenMapWorkflowWorkflowDefinitionRelationship(s.GetWorkflowDefinition(), d)); err != nil {
		return diag.Errorf("error occurred while setting property WorkflowDefinition in WorkflowServiceItemActionDefinition object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceWorkflowServiceItemActionDefinitionUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.WorkflowServiceItemActionDefinition{}

	if d.HasChange("action_type") {
		v := d.Get("action_type")
		x := (v.(string))
		o.SetActionType(x)
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

	if d.HasChange("allowed_instance_states") {
		v := d.Get("allowed_instance_states")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetAllowedInstanceStates(x)
	}

	o.SetClassId("workflow.ServiceItemActionDefinition")

	if d.HasChange("core_workflows") {
		v := d.Get("core_workflows")
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowServiceItemActionWorkflowDefinition{}
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetCoreWorkflows(x)
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
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

	o.SetObjectType("workflow.ServiceItemActionDefinition")

	if d.HasChange("output_parameters") {
		v := d.Get("output_parameters")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			x2 := x1.(map[string]interface{})
			o.SetOutputParameters(x2)
		}
	}

	if d.HasChange("periodicity") {
		v := d.Get("periodicity")
		x := int64(v.(int))
		o.SetPeriodicity(x)
	}

	if d.HasChange("post_core_workflows") {
		v := d.Get("post_core_workflows")
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowServiceItemActionWorkflowDefinition{}
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPostCoreWorkflows(x)
	}

	if d.HasChange("pre_core_workflows") {
		v := d.Get("pre_core_workflows")
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowServiceItemActionWorkflowDefinition{}
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPreCoreWorkflows(x)
	}

	if d.HasChange("restrict_on_private_appliance") {
		v := d.Get("restrict_on_private_appliance")
		x := (v.(bool))
		o.SetRestrictOnPrivateAppliance(x)
	}

	if d.HasChange("service_item_definition") {
		v := d.Get("service_item_definition")
		p := make([]models.WorkflowServiceItemDefinitionRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowServiceItemDefinitionRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetServiceItemDefinition(x)
		}
	}

	if d.HasChange("stop_workflows") {
		v := d.Get("stop_workflows")
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowServiceItemActionWorkflowDefinition{}
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetStopWorkflows(x)
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

	if d.HasChange("validation_workflows") {
		v := d.Get("validation_workflows")
		x := make([]models.WorkflowServiceItemActionWorkflowDefinition, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.WorkflowServiceItemActionWorkflowDefinition{}
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
			o.SetClassId("workflow.ServiceItemActionWorkflowDefinition")
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
			if v, ok := l["nr_version"]; ok {
				{
					x := int64(v.(int))
					o.SetVersion(x)
				}
			}
			if v, ok := l["workflow_definition_name"]; ok {
				{
					x := (v.(string))
					o.SetWorkflowDefinitionName(x)
				}
			}
			x = append(x, *o)
		}
		o.SetValidationWorkflows(x)
	}

	r := conn.ApiClient.WorkflowApi.UpdateWorkflowServiceItemActionDefinition(conn.ctx, d.Id()).WorkflowServiceItemActionDefinition(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating WorkflowServiceItemActionDefinition: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating WorkflowServiceItemActionDefinition: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceWorkflowServiceItemActionDefinitionRead(c, d, meta)...)
}

func resourceWorkflowServiceItemActionDefinitionDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.WorkflowApi.DeleteWorkflowServiceItemActionDefinition(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "WorkflowServiceItemActionDefinitionDelete: WorkflowServiceItemActionDefinition object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting WorkflowServiceItemActionDefinition object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting WorkflowServiceItemActionDefinition object: %s", deleteErr.Error())
	}
	return de
}
