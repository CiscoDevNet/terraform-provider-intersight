package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricFcoeUplinkRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricFcoeUplinkRoleRead,
		Schema: map[string]*schema.Schema{
			"admin_speed": {
				Description: "Admin configured speed for the port.\n* `Auto` - Admin configurable speed AUTO ( default ).\n* `1Gbps` - Admin configurable speed 1Gbps.\n* `10Gbps` - Admin configurable speed 10Gbps.\n* `25Gbps` - Admin configurable speed 25Gbps.\n* `40Gbps` - Admin configurable speed 40Gbps.\n* `100Gbps` - Admin configurable speed 100Gbps.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"aggregate_port_id": {
				Description: "Breakout port Identifier of the Switch Interface.\nWhen a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.\nWhen a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,\ne.g. the id of the port on the switch.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fec": {
				Description: "Forward error correction configuration for the port.\n* `Auto` - Forward error correction option 'Auto'.\n* `Cl91` - Forward error correction option 'cl91'.\n* `Cl74` - Forward error correction option 'cl74'.",
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
			"port_id": {
				Description: "Port Identifier of the Switch/FEX/Chassis Interface.\nWhen a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,\ne.g. the id of the port on the switch, FEX or chassis.\nWhen a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"slot_id": {
				Description: "Slot Identifier of the Switch/FEX/Chassis Interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"udld_admin_state": {
				Description: "Admin configured state for UDLD for this port.\n* `Disabled` - Admin configured Disabled State.\n* `Enabled` - Admin configured Enabled State.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricFcoeUplinkRole().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricFcoeUplinkRoleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricFcoeUplinkRole{}
	if v, ok := d.GetOk("admin_speed"); ok {
		x := (v.(string))
		o.SetAdminSpeed(x)
	}
	if v, ok := d.GetOk("aggregate_port_id"); ok {
		x := int64(v.(int))
		o.SetAggregatePortId(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("fec"); ok {
		x := (v.(string))
		o.SetFec(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("port_id"); ok {
		x := int64(v.(int))
		o.SetPortId(x)
	}
	if v, ok := d.GetOk("slot_id"); ok {
		x := int64(v.(int))
		o.SetSlotId(x)
	}
	if v, ok := d.GetOk("udld_admin_state"); ok {
		x := (v.(string))
		o.SetUdldAdminState(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricFcoeUplinkRole object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricFcoeUplinkRoleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricFcoeUplinkRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricFcoeUplinkRoleList.GetCount()
	var i int32
	var fabricFcoeUplinkRoleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricFcoeUplinkRoleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricFcoeUplinkRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricFcoeUplinkRoleList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricFcoeUplinkRole data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_speed"] = (s.GetAdminSpeed())
				temp["aggregate_port_id"] = (s.GetAggregatePortId())
				temp["class_id"] = (s.GetClassId())
				temp["fec"] = (s.GetFec())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["port_id"] = (s.GetPortId())

				temp["port_policy"] = flattenMapFabricPortPolicyRelationship(s.GetPortPolicy(), d)
				temp["slot_id"] = (s.GetSlotId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["udld_admin_state"] = (s.GetUdldAdminState())
				fabricFcoeUplinkRoleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricFcoeUplinkRoleResults))
	if err := d.Set("results", fabricFcoeUplinkRoleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricFcoeUplinkRoleResults[0]["moid"].(string))
	return de
}
