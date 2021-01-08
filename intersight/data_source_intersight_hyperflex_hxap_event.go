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

func dataSourceHyperflexHxapEvent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHxapEventRead,
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
			"cluster": {
				Description: "A reference to a hyperflexHxapCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"first_time": {
				Description: "First timestamp of the event occurrence.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description: "Internally generated identity (UUID) of this event.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_time": {
				Description: "Last timestamp of the event occurrence.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"message": {
				Description: "Full description of the event.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
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
			"owner_name": {
				Description: "Name of the owner with which event is associated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"owner_type": {
				Description: "Type of the object with which event is associated (Host, Cluster, VM).\n* `Unknown` - Value is Unknown if the target type is unidentified.\n* `Cluster` - Cluster refers to HyperFlex AP Cluster.\n* `Host` - Host refers to server node which is part of HyperFlex AP Cluster.\n* `VM` - VM refers to Virtual machine available on a HyperFlex AP Node.\n* `Disk` - Disk refers to Virtual Disk available on a HyperFlex AP Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"owner_uuid": {
				Description: "UUID of the owner with which event is associated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"severity": {
				Description: "Severity level of the event (Info/Warning/Critical).\n* `None` - The Enum value None represents that there is no severity.\n* `Info` - The Enum value Info represents the Informational level of severity.\n* `Critical` - The Enum value Critical represents the Critical level of severity.\n* `Warning` - The Enum value Warning represents the Warning level of severity.\n* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.",
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
		},
	}
}

func dataSourceHyperflexHxapEventRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHxapEvent{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("first_time"); ok {
		x := (v.(string))
		o.SetFirstTime(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("last_time"); ok {
		x := (v.(string))
		o.SetLastTime(x)
	}
	if v, ok := d.GetOk("message"); ok {
		x := (v.(string))
		o.SetMessage(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("owner_name"); ok {
		x := (v.(string))
		o.SetOwnerName(x)
	}
	if v, ok := d.GetOk("owner_type"); ok {
		x := (v.(string))
		o.SetOwnerType(x)
	}
	if v, ok := d.GetOk("owner_uuid"); ok {
		x := (v.(string))
		o.SetOwnerUuid(x)
	}
	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.SetSeverity(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHxapEvent object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHxapEventList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexHxapEvent: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexHxapEvent list: %s", err.Error())
	}
	var s = &models.HyperflexHxapEventList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexHxapEvent list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for HyperflexHxapEvent did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexHxapEvent{}
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

			if err := d.Set("cluster", flattenMapHyperflexHxapClusterRelationship(s.GetCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Cluster: %s", err.Error())
			}
			if err := d.Set("first_time", (s.GetFirstTime())); err != nil {
				return diag.Errorf("error occurred while setting property FirstTime: %s", err.Error())
			}
			if err := d.Set("identity", (s.GetIdentity())); err != nil {
				return diag.Errorf("error occurred while setting property Identity: %s", err.Error())
			}
			if err := d.Set("last_time", (s.GetLastTime())); err != nil {
				return diag.Errorf("error occurred while setting property LastTime: %s", err.Error())
			}
			if err := d.Set("message", (s.GetMessage())); err != nil {
				return diag.Errorf("error occurred while setting property Message: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("owner_name", (s.GetOwnerName())); err != nil {
				return diag.Errorf("error occurred while setting property OwnerName: %s", err.Error())
			}
			if err := d.Set("owner_type", (s.GetOwnerType())); err != nil {
				return diag.Errorf("error occurred while setting property OwnerType: %s", err.Error())
			}
			if err := d.Set("owner_uuid", (s.GetOwnerUuid())); err != nil {
				return diag.Errorf("error occurred while setting property OwnerUuid: %s", err.Error())
			}
			if err := d.Set("severity", (s.GetSeverity())); err != nil {
				return diag.Errorf("error occurred while setting property Severity: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
