package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricPortOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricPortOperationRead,
		Schema: map[string]*schema.Schema{
			"admin_state": {
				Description: "Admin configured state to disable the port.\n* `Enabled` - Admin configured Enabled State.\n* `Disabled` - Admin configured Disabled State.",
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
			"config_state": {
				Description: "The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain.\n* `None` - Nil value when no action has been triggered by the user.\n* `Applied` - User configured settings are in applied state.\n* `Applying` - User settings are being applied on the target server.\n* `Failed` - User configured settings could not be applied.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFabricPortOperation().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricPortOperationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricPortOperation{}
	if v, ok := d.GetOk("admin_state"); ok {
		x := (v.(string))
		o.SetAdminState(x)
	}
	if v, ok := d.GetOk("aggregate_port_id"); ok {
		x := int64(v.(int))
		o.SetAggregatePortId(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("config_state"); ok {
		x := (v.(string))
		o.SetConfigState(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricPortOperation object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricPortOperationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricPortOperation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricPortOperationList.GetCount()
	var i int32
	var fabricPortOperationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricPortOperationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricPortOperation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricPortOperationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricPortOperation data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_state"] = (s.GetAdminState())
				temp["aggregate_port_id"] = (s.GetAggregatePortId())
				temp["class_id"] = (s.GetClassId())
				temp["config_state"] = (s.GetConfigState())
				temp["moid"] = (s.GetMoid())

				temp["network_element"] = flattenMapNetworkElementRelationship(s.GetNetworkElement(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["port_id"] = (s.GetPortId())
				temp["slot_id"] = (s.GetSlotId())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				fabricPortOperationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricPortOperationResults))
	if err := d.Set("results", fabricPortOperationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricPortOperationResults[0]["moid"].(string))
	return de
}
