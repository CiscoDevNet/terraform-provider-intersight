package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexVmRestoreOperation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexVmRestoreOperationRead,
		Schema: map[string]*schema.Schema{
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
			"new_name": {
				Description: "New name for the Virtual Machine after recovery.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"power_on": {
				Description: "Power on the Virtual Machine after recovery.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexVmRestoreOperation().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexVmRestoreOperationRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexVmRestoreOperation{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("new_name"); ok {
		x := (v.(string))
		o.SetNewName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("power_on"); ok {
		x := (v.(bool))
		o.SetPowerOn(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexVmRestoreOperation object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmRestoreOperationList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexVmRestoreOperation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexVmRestoreOperationList.GetCount()
	var i int32
	var hyperflexVmRestoreOperationResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmRestoreOperationList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexVmRestoreOperation: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexVmRestoreOperationList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexVmRestoreOperation data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["new_name"] = (s.GetNewName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["power_on"] = (s.GetPowerOn())

				temp["restore_edge_cluster_moid"] = flattenMapHyperflexClusterRelationship(s.GetRestoreEdgeClusterMoid(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["vm_backup_info_moid"] = flattenMapHyperflexVmBackupInfoRelationship(s.GetVmBackupInfoMoid(), d)

				temp["vm_snapshot_info_moid"] = flattenMapHyperflexVmSnapshotInfoRelationship(s.GetVmSnapshotInfoMoid(), d)
				hyperflexVmRestoreOperationResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexVmRestoreOperationResults))
	if err := d.Set("results", hyperflexVmRestoreOperationResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexVmRestoreOperationResults[0]["moid"].(string))
	return de
}
