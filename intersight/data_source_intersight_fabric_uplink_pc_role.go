package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricUplinkPcRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFabricUplinkPcRoleRead,
		Schema: map[string]*schema.Schema{
			"admin_speed": {
				Description: "Admin configured speed for the port.\n* `Auto` - Admin configurable speed AUTO ( default ).\n* `1Gbps` - Admin configurable speed 1Gbps.\n* `10Gbps` - Admin configurable speed 10Gbps.\n* `25Gbps` - Admin configurable speed 25Gbps.\n* `40Gbps` - Admin configurable speed 40Gbps.\n* `100Gbps` - Admin configurable speed 100Gbps.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"pc_id": {
				Description: "Unique Identifier of the port-channel, local to this switch.",
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
				Elem:     &schema.Resource{Schema: resourceFabricUplinkPcRole().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFabricUplinkPcRoleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FabricUplinkPcRole{}
	if v, ok := d.GetOk("admin_speed"); ok {
		x := (v.(string))
		o.SetAdminSpeed(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pc_id"); ok {
		x := int64(v.(int))
		o.SetPcId(x)
	}
	if v, ok := d.GetOk("udld_admin_state"); ok {
		x := (v.(string))
		o.SetUdldAdminState(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FabricUplinkPcRole object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FabricApi.GetFabricUplinkPcRoleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FabricUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FabricUplinkPcRoleList.GetCount()
	var i int32
	var fabricUplinkPcRoleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FabricApi.GetFabricUplinkPcRoleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FabricUplinkPcRole: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FabricUplinkPcRoleList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FabricUplinkPcRole data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["admin_speed"] = (s.GetAdminSpeed())
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["pc_id"] = (s.GetPcId())

				temp["port_policy"] = flattenMapFabricPortPolicyRelationship(s.GetPortPolicy(), d)

				temp["ports"] = flattenListFabricPortIdentifier(s.GetPorts(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["udld_admin_state"] = (s.GetUdldAdminState())
				fabricUplinkPcRoleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(fabricUplinkPcRoleResults))
	if err := d.Set("results", fabricUplinkPcRoleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(fabricUplinkPcRoleResults[0]["moid"].(string))
	return de
}
