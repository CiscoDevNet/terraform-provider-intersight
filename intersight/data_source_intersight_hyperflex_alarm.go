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

func dataSourceHyperflexAlarm() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexAlarmRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Description: "The acknowledgement state of the alarm. It is 'true' when the alarm is acknowledged and false otherwise.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_by": {
				Description: "The username of the user who acknowledged the alarm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_time": {
				Description: "The time when the alarm was acknowledged, represented as a Unix timestamp.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"acknowledged_time_as_utc": {
				Description: "The time when the alarm was acknowledged in ISO 6801 format.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
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
			"cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"description": {
				Description: "The description of the alarm which includes information about the fault that triggered the alarm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_data": {
				Description: "The data pertaining to this entity.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_name": {
				Description: "The name of entity which triggered the alarm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_type": {
				Description: "The type of entity which triggered the alarm. For example, this can be the cluster, a node, or VM running on the cluster.\n* `UNKNOWN` - The type of entity is not known.\n* `DISK` - The entity is a physical storage device.\n* `NODE` - The entity is a HyperFlex cluster node.\n* `CLUSTER` - The entity is the HyperFlex cluster itself.\n* `DATASTORE` - The entity is a logical datastore configured on the HyperFlex cluster.\n* `ZONE` - The entity is a logical or physical zone configured on the HyperFlex cluster.\n* `VIRTUALMACHINE` - The entity is a virtual machine running on the HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"entity_uu_id": {
				Description: "The unique identifier of the entity which triggered the alarm.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"message": {
				Description: "The localized message displayed to the user which describes the alarm.",
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
			"name": {
				Description: "The name of the alarm. This name identifies the type of alarm that was triggered.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "The severity of the alarm.\n* `UNKNOWN` - The alarm status is not known.\n* `CLEARED` - The event that triggered the alarm has been remedied and no longer requires the user's attention.\n* `INFO` - The alarm represents a message that does not require the user's immediate attention.\n* `WARNING` - The alarm represents a moderate fault.\n* `CRITICAL` - The alarm represents a critical fault.",
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
			"triggered_time": {
				Description: "The time when alarm was triggered as a Unix timestamp.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"triggered_time_as_utc": {
				Description: "The time when alarm was triggered in ISO 6801 UTC format.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "The unique identifier for this alarm instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceHyperflexAlarmRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexAlarm{}
	if v, ok := d.GetOk("acknowledged"); ok {
		x := (v.(bool))
		o.SetAcknowledged(x)
	}
	if v, ok := d.GetOk("acknowledged_by"); ok {
		x := (v.(string))
		o.SetAcknowledgedBy(x)
	}
	if v, ok := d.GetOk("acknowledged_time"); ok {
		x := int64(v.(int))
		o.SetAcknowledgedTime(x)
	}
	if v, ok := d.GetOk("acknowledged_time_as_utc"); ok {
		x := (v.(string))
		o.SetAcknowledgedTimeAsUtc(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("entity_data"); ok {
		x := (v.(string))
		o.SetEntityData(x)
	}
	if v, ok := d.GetOk("entity_name"); ok {
		x := (v.(string))
		o.SetEntityName(x)
	}
	if v, ok := d.GetOk("entity_type"); ok {
		x := (v.(string))
		o.SetEntityType(x)
	}
	if v, ok := d.GetOk("entity_uu_id"); ok {
		x := (v.(string))
		o.SetEntityUuId(x)
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.SetMessage(x)
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
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("triggered_time"); ok {
		x := int64(v.(int))
		o.SetTriggeredTime(x)
	}
	if v, ok := d.GetOk("triggered_time_as_utc"); ok {
		x := (v.(string))
		o.SetTriggeredTimeAsUtc(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexAlarm object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexAlarmList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexAlarm: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexAlarm list: %s", err.Error())
	}
	var s = &models.HyperflexAlarmList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexAlarm list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexAlarm data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexAlarm data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexAlarm{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("acknowledged", (s.GetAcknowledged())); err != nil {
				return diag.Errorf("error occurred while setting property Acknowledged: %s", err.Error())
			}
			if err := d.Set("acknowledged_by", (s.GetAcknowledgedBy())); err != nil {
				return diag.Errorf("error occurred while setting property AcknowledgedBy: %s", err.Error())
			}
			if err := d.Set("acknowledged_time", (s.GetAcknowledgedTime())); err != nil {
				return diag.Errorf("error occurred while setting property AcknowledgedTime: %s", err.Error())
			}
			if err := d.Set("acknowledged_time_as_utc", (s.GetAcknowledgedTimeAsUtc())); err != nil {
				return diag.Errorf("error occurred while setting property AcknowledgedTimeAsUtc: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster", flattenMapHyperflexClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("entity_data", (s.GetEntityData())); err != nil {
				return diag.Errorf("error occurred while setting property EntityData: %s", err.Error())
			}
			if err := d.Set("entity_name", (s.GetEntityName())); err != nil {
				return diag.Errorf("error occurred while setting property EntityName: %s", err.Error())
			}
			if err := d.Set("entity_type", (s.GetEntityType())); err != nil {
				return diag.Errorf("error occurred while setting property EntityType: %s", err.Error())
			}
			if err := d.Set("entity_uu_id", (s.GetEntityUuId())); err != nil {
				return diag.Errorf("error occurred while setting property EntityUuId: %s", err.Error())
			}
			if err := d.Set("message", (s.GetMessage())); err != nil {
				return diag.Errorf("error occurred while setting property Message: %s", err.Error())
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
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return diag.Errorf("error occurred while setting property Status: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("triggered_time", (s.GetTriggeredTime())); err != nil {
				return diag.Errorf("error occurred while setting property TriggeredTime: %s", err.Error())
			}
			if err := d.Set("triggered_time_as_utc", (s.GetTriggeredTimeAsUtc())); err != nil {
				return diag.Errorf("error occurred while setting property TriggeredTimeAsUtc: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
