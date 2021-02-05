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

func dataSourceHyperflexClusterBackupPolicyDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterBackupPolicyDeploymentRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
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
			"backup_target": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
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
			"replication_schedule": {
				Description: "Backup policy replication schedule.",
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
						"backup_interval": {
							Description: "Time interval between two copies in minutes.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"snapshot_retention_count": {
				Description: "Number of snapshots that will be retained as part of the Multi Point in Time support.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"source_cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
				Computed: true,
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
		},
	}
}

func dataSourceHyperflexClusterBackupPolicyDeploymentRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexClusterBackupPolicyDeployment{}
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
	if v, ok := d.GetOk("discovered"); ok {
		x := (v.(bool))
		o.SetDiscovered(x)
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
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterBackupPolicyDeploymentList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching HyperflexClusterBackupPolicyDeployment: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexClusterBackupPolicyDeployment list: %s", err.Error())
	}
	var s = &models.HyperflexClusterBackupPolicyDeploymentList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexClusterBackupPolicyDeployment list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexClusterBackupPolicyDeployment data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexClusterBackupPolicyDeployment data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexClusterBackupPolicyDeployment{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("backup_data_store_name", (s.GetBackupDataStoreName())); err != nil {
				return diag.Errorf("error occurred while setting property BackupDataStoreName: %s", err.Error())
			}
			if err := d.Set("backup_data_store_size", (s.GetBackupDataStoreSize())); err != nil {
				return diag.Errorf("error occurred while setting property BackupDataStoreSize: %s", err.Error())
			}
			if err := d.Set("backup_data_store_size_unit", (s.GetBackupDataStoreSizeUnit())); err != nil {
				return diag.Errorf("error occurred while setting property BackupDataStoreSizeUnit: %s", err.Error())
			}

			if err := d.Set("backup_target", flattenMapHyperflexClusterRelationship(s.GetBackupTarget(), d)); err != nil {
				return diag.Errorf("error occurred while setting property BackupTarget: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("discovered", (s.GetDiscovered())); err != nil {
				return diag.Errorf("error occurred while setting property Discovered: %s", err.Error())
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

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Organization: %s", err.Error())
			}
			if err := d.Set("policy_moid", (s.GetPolicyMoid())); err != nil {
				return diag.Errorf("error occurred while setting property PolicyMoid: %s", err.Error())
			}
			if err := d.Set("profile_moid", (s.GetProfileMoid())); err != nil {
				return diag.Errorf("error occurred while setting property ProfileMoid: %s", err.Error())
			}
			if err := d.Set("replication_pair_name_prefix", (s.GetReplicationPairNamePrefix())); err != nil {
				return diag.Errorf("error occurred while setting property ReplicationPairNamePrefix: %s", err.Error())
			}

			if err := d.Set("replication_schedule", flattenMapHyperflexReplicationSchedule(s.GetReplicationSchedule(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ReplicationSchedule: %s", err.Error())
			}
			if err := d.Set("snapshot_retention_count", (s.GetSnapshotRetentionCount())); err != nil {
				return diag.Errorf("error occurred while setting property SnapshotRetentionCount: %s", err.Error())
			}

			if err := d.Set("source_cluster", flattenMapHyperflexClusterRelationship(s.GetSourceCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SourceCluster: %s", err.Error())
			}
			if err := d.Set("source_detached", (s.GetSourceDetached())); err != nil {
				return diag.Errorf("error occurred while setting property SourceDetached: %s", err.Error())
			}
			if err := d.Set("source_request_id", (s.GetSourceRequestId())); err != nil {
				return diag.Errorf("error occurred while setting property SourceRequestId: %s", err.Error())
			}
			if err := d.Set("source_uuid", (s.GetSourceUuid())); err != nil {
				return diag.Errorf("error occurred while setting property SourceUuid: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("target_detached", (s.GetTargetDetached())); err != nil {
				return diag.Errorf("error occurred while setting property TargetDetached: %s", err.Error())
			}
			if err := d.Set("target_request_id", (s.GetTargetRequestId())); err != nil {
				return diag.Errorf("error occurred while setting property TargetRequestId: %s", err.Error())
			}
			if err := d.Set("target_uuid", (s.GetTargetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property TargetUuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
