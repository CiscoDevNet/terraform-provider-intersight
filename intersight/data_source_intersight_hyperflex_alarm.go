package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHyperflexAlarm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHyperflexAlarmRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"acknowledged_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acknowledged_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"acknowledged_time_as_utc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_uu_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
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
			"triggered_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"triggered_time_as_utc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceHyperflexAlarmRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
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
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.HyperflexApi.GetHyperflexAlarmList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.HyperflexAlarmList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to HyperflexAlarm: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexAlarm{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("acknowledged", (s.GetAcknowledged())); err != nil {
				return fmt.Errorf("error occurred while setting property Acknowledged: %+v", err)
			}
			if err := d.Set("acknowledged_by", (s.GetAcknowledgedBy())); err != nil {
				return fmt.Errorf("error occurred while setting property AcknowledgedBy: %+v", err)
			}
			if err := d.Set("acknowledged_time", (s.GetAcknowledgedTime())); err != nil {
				return fmt.Errorf("error occurred while setting property AcknowledgedTime: %+v", err)
			}
			if err := d.Set("acknowledged_time_as_utc", (s.GetAcknowledgedTimeAsUtc())); err != nil {
				return fmt.Errorf("error occurred while setting property AcknowledgedTimeAsUtc: %+v", err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("cluster", flattenMapHyperflexClusterRelationship(s.GetCluster(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Cluster: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}
			if err := d.Set("entity_data", (s.GetEntityData())); err != nil {
				return fmt.Errorf("error occurred while setting property EntityData: %+v", err)
			}
			if err := d.Set("entity_name", (s.GetEntityName())); err != nil {
				return fmt.Errorf("error occurred while setting property EntityName: %+v", err)
			}
			if err := d.Set("entity_type", (s.GetEntityType())); err != nil {
				return fmt.Errorf("error occurred while setting property EntityType: %+v", err)
			}
			if err := d.Set("entity_uu_id", (s.GetEntityUuId())); err != nil {
				return fmt.Errorf("error occurred while setting property EntityUuId: %+v", err)
			}
			if err := d.Set("message", (s.GetMessage())); err != nil {
				return fmt.Errorf("error occurred while setting property Message: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("status", (s.GetStatus())); err != nil {
				return fmt.Errorf("error occurred while setting property Status: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("triggered_time", (s.GetTriggeredTime())); err != nil {
				return fmt.Errorf("error occurred while setting property TriggeredTime: %+v", err)
			}
			if err := d.Set("triggered_time_as_utc", (s.GetTriggeredTimeAsUtc())); err != nil {
				return fmt.Errorf("error occurred while setting property TriggeredTimeAsUtc: %+v", err)
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return fmt.Errorf("error occurred while setting property Uuid: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
