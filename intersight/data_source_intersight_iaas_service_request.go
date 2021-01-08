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

func dataSourceIaasServiceRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaasServiceRequestRead,
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
			"duration": {
				Description: "Service request duration.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"initiating_user": {
				Description: "Service Request initiating user.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
			},
			"request_end_time": {
				Description: "Service request end time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_id": {
				Description: "Service request id of an SR.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_start_time": {
				Description: "Service request start time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_type": {
				Description: "Service request type of an SR.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "UCSD service request status.",
				Type:        schema.TypeString,
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
			"workflow_name": {
				Description: "Executed workflow name for an SR.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"workflow_steps": {
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
						"completed_time": {
							Description: "Completed time of the workflow step.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the workflow step.",
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
						"status": {
							Description: "Status of the workflow step.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"status_message": {
							Description: "Status message of the workflow step.",
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

func dataSourceIaasServiceRequestRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IaasServiceRequest{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("duration"); ok {
		x := (v.(string))
		o.SetDuration(x)
	}
	if v, ok := d.GetOk("initiating_user"); ok {
		x := (v.(string))
		o.SetInitiatingUser(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("request_end_time"); ok {
		x := (v.(string))
		o.SetRequestEndTime(x)
	}
	if v, ok := d.GetOk("request_id"); ok {
		x := (v.(string))
		o.SetRequestId(x)
	}
	if v, ok := d.GetOk("request_start_time"); ok {
		x := (v.(string))
		o.SetRequestStartTime(x)
	}
	if v, ok := d.GetOk("request_type"); ok {
		x := (v.(string))
		o.SetRequestType(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("workflow_name"); ok {
		x := (v.(string))
		o.SetWorkflowName(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IaasServiceRequest object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.IaasApi.GetIaasServiceRequestList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching IaasServiceRequest: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for IaasServiceRequest list: %s", err.Error())
	}
	var s = &models.IaasServiceRequestList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to IaasServiceRequest list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for IaasServiceRequest did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.IaasServiceRequest{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("duration", (s.GetDuration())); err != nil {
				return diag.Errorf("error occurred while setting property Duration: %s", err.Error())
			}
			if err := d.Set("initiating_user", (s.GetInitiatingUser())); err != nil {
				return diag.Errorf("error occurred while setting property InitiatingUser: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("request_end_time", (s.GetRequestEndTime())); err != nil {
				return diag.Errorf("error occurred while setting property RequestEndTime: %s", err.Error())
			}
			if err := d.Set("request_id", (s.GetRequestId())); err != nil {
				return diag.Errorf("error occurred while setting property RequestId: %s", err.Error())
			}
			if err := d.Set("request_start_time", (s.GetRequestStartTime())); err != nil {
				return diag.Errorf("error occurred while setting property RequestStartTime: %s", err.Error())
			}
			if err := d.Set("request_type", (s.GetRequestType())); err != nil {
				return diag.Errorf("error occurred while setting property RequestType: %s", err.Error())
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("workflow_name", (s.GetWorkflowName())); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowName: %s", err.Error())
			}

			if err := d.Set("workflow_steps", flattenListIaasWorkflowSteps(s.GetWorkflowSteps(), d)); err != nil {
				return diag.Errorf("error occurred while setting property WorkflowSteps: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
