package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConnectorpackConnectorPackUpgrade() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectorpackConnectorPackUpgradeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector_pack_op_type": {
				Description: "The type of operation to be performed on UCS Director.\n* `Install` - Installs the requisite connector packs on UCS Director.\n* `Push` - Pushes the requisite connector packs to UCS Director.",
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
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceConnectorpackConnectorPackUpgrade().Schema},
				Computed: true,
			}},
	}
}

func dataSourceConnectorpackConnectorPackUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ConnectorpackConnectorPackUpgrade{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connector_pack_op_type"); ok {
		x := (v.(string))
		o.SetConnectorPackOpType(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ConnectorpackConnectorPackUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ConnectorpackApi.GetConnectorpackConnectorPackUpgradeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ConnectorpackConnectorPackUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ConnectorpackConnectorPackUpgradeList.GetCount()
	var i int32
	var connectorpackConnectorPackUpgradeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ConnectorpackApi.GetConnectorpackConnectorPackUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ConnectorpackConnectorPackUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ConnectorpackConnectorPackUpgradeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ConnectorpackConnectorPackUpgrade data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["connector_pack_op_type"] = (s.GetConnectorPackOpType())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["ucsd_info"] = flattenMapIaasUcsdInfoRelationship(s.GetUcsdInfo(), d)

				temp["workflow"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflow(), d)
				connectorpackConnectorPackUpgradeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(connectorpackConnectorPackUpgradeResults))
	if err := d.Set("results", connectorpackConnectorPackUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(connectorpackConnectorPackUpgradeResults[0]["moid"].(string))
	return de
}
