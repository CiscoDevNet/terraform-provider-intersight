package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexClusterBackupPolicyDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterBackupPolicyDeploymentRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_data_store_name": {
				Description: "Backup data store name used during the auto creation of the datastore. All VMs created in this data store will be automatically backed up.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"backup_data_store_size": {
				Description: "Replication data store size in backupDataStoreSizeUnit.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"backup_data_store_size_unit": {
				Description: "Replication data store size.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description from corresponding ClusterBackupPolicy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"discovered": {
				Description: "True if record created by discovery on HyperFlex cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
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
				Description: "Name from corresponding ClusterBackupPolicy.",
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
			"policy_moid": {
				Description: "Deployed cluster policy moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"profile_moid": {
				Description: "Deployed cluster profile moid.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"replication_pair_name_prefix": {
				Description: "Replication cluster pairing name prefix.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"snapshot_retention_count": {
				Description: "Number of snapshots that will be retained as part of the Multi Point in Time support.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"source_detached": {
				Description: "True if policy was detached from source Hyperflex Cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"source_request_id": {
				Description: "Unique source cluster request ID allowing retry of the same logical request following a transient communication failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"source_uuid": {
				Description: "Uuid of the source Hyperflex Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"target_detached": {
				Description: "True if policy was detached from target Hyperflex Cluster.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"target_request_id": {
				Description: "Unique target cluster request ID allowing retry of the same logical request following a transient communication failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"target_uuid": {
				Description: "Uuid of the target Hyperflex Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexClusterBackupPolicyDeployment().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterBackupPolicyDeploymentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterBackupPolicyDeployment{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
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
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("discovered"); ok {
		x := (v.(bool))
		o.SetDiscovered(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
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
	if v, ok := d.GetOk("policy_moid"); ok {
		x := (v.(string))
		o.SetPolicyMoid(x)
	}
	if v, ok := d.GetOk("profile_moid"); ok {
		x := (v.(string))
		o.SetProfileMoid(x)
	}
	if v, ok := d.GetOk("replication_pair_name_prefix"); ok {
		x := (v.(string))
		o.SetReplicationPairNamePrefix(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("snapshot_retention_count"); ok {
		x := int64(v.(int))
		o.SetSnapshotRetentionCount(x)
	}
	if v, ok := d.GetOk("source_detached"); ok {
		x := (v.(bool))
		o.SetSourceDetached(x)
	}
	if v, ok := d.GetOk("source_request_id"); ok {
		x := (v.(string))
		o.SetSourceRequestId(x)
	}
	if v, ok := d.GetOk("source_uuid"); ok {
		x := (v.(string))
		o.SetSourceUuid(x)
	}
	if v, ok := d.GetOk("target_detached"); ok {
		x := (v.(bool))
		o.SetTargetDetached(x)
	}
	if v, ok := d.GetOk("target_request_id"); ok {
		x := (v.(string))
		o.SetTargetRequestId(x)
	}
	if v, ok := d.GetOk("target_uuid"); ok {
		x := (v.(string))
		o.SetTargetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexClusterBackupPolicyDeployment object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyDeploymentList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexClusterBackupPolicyDeployment: %s", responseErr.Error())
	}
	count := countResponse.HyperflexClusterBackupPolicyDeploymentList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexClusterBackupPolicyDeployment data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexClusterBackupPolicyDeploymentResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyDeploymentList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexClusterBackupPolicyDeployment: %s", responseErr.Error())
		}
		results := resMo.HyperflexClusterBackupPolicyDeploymentList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["backup_data_store_name"] = (s.GetBackupDataStoreName())
				temp["backup_data_store_size"] = (s.GetBackupDataStoreSize())
				temp["backup_data_store_size_unit"] = (s.GetBackupDataStoreSizeUnit())

				temp["backup_target"] = flattenMapHyperflexClusterRelationship(s.GetBackupTarget(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["description"] = (s.GetDescription())
				temp["discovered"] = (s.GetDiscovered())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["policy_moid"] = (s.GetPolicyMoid())
				temp["profile_moid"] = (s.GetProfileMoid())
				temp["replication_pair_name_prefix"] = (s.GetReplicationPairNamePrefix())

				temp["replication_schedule"] = flattenMapHyperflexReplicationSchedule(s.GetReplicationSchedule(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["snapshot_retention_count"] = (s.GetSnapshotRetentionCount())

				temp["source_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSourceCluster(), d)
				temp["source_detached"] = (s.GetSourceDetached())
				temp["source_request_id"] = (s.GetSourceRequestId())
				temp["source_uuid"] = (s.GetSourceUuid())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_detached"] = (s.GetTargetDetached())
				temp["target_request_id"] = (s.GetTargetRequestId())
				temp["target_uuid"] = (s.GetTargetUuid())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				hyperflexClusterBackupPolicyDeploymentResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterBackupPolicyDeploymentResults))
	if err := d.Set("results", hyperflexClusterBackupPolicyDeploymentResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterBackupPolicyDeploymentResults[0]["moid"].(string))
	return de
}
