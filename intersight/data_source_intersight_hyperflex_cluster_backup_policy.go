package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexClusterBackupPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterBackupPolicyRead,
		Schema: map[string]*schema.Schema{
			"backup_data_store_name": {
				Description: "Backup datastore name prefix used during the auto creation of the datastore. All VMs created in this datastore will be automatically backed up.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"backup_data_store_size": {
				Description: "Replication data store size in backupDataStoreSizeUnit.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"backup_data_store_size_unit": {
				Description: "Replication data store size.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"replication_pair_name_prefix": {
				Description: "Replication cluster pairing name prefix.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snapshot_retention_count": {
				Description: "Number of snapshots that will be retained as part of the Multi Point in Time support.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexClusterBackupPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterBackupPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterBackupPolicy{}
	if v, ok := d.GetOk("backup_data_store_name"); ok {
		x := (v.(string))
		o.SetBackupDataStoreName(x)
	}
	if v, ok := d.GetOk("backup_data_store_size"); ok {
		x := int64(v.(int))
		o.SetBackupDataStoreSize(x)
	}
	if v, ok := d.GetOk("backup_data_store_size_unit"); ok {
		x := (v.(string))
		o.SetBackupDataStoreSizeUnit(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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
	if v, ok := d.GetOk("replication_pair_name_prefix"); ok {
		x := (v.(string))
		o.SetReplicationPairNamePrefix(x)
	}
	if v, ok := d.GetOk("snapshot_retention_count"); ok {
		x := int64(v.(int))
		o.SetSnapshotRetentionCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexClusterBackupPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexClusterBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexClusterBackupPolicyList.GetCount()
	var i int32
	var hyperflexClusterBackupPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexClusterBackupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexClusterBackupPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexClusterBackupPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["backup_data_store_name"] = (s.GetBackupDataStoreName())
				temp["backup_data_store_size"] = (s.GetBackupDataStoreSize())
				temp["backup_data_store_size_unit"] = (s.GetBackupDataStoreSizeUnit())

				temp["backup_target"] = flattenMapHyperflexClusterRelationship(s.GetBackupTarget(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profiles"] = flattenListHyperflexClusterProfileRelationship(s.GetClusterProfiles(), d)
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["replication_pair_name_prefix"] = (s.GetReplicationPairNamePrefix())

				temp["replication_schedule"] = flattenMapHyperflexReplicationSchedule(s.GetReplicationSchedule(), d)
				temp["snapshot_retention_count"] = (s.GetSnapshotRetentionCount())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				hyperflexClusterBackupPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterBackupPolicyResults))
	if err := d.Set("results", hyperflexClusterBackupPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterBackupPolicyResults[0]["moid"].(string))
	return de
}
