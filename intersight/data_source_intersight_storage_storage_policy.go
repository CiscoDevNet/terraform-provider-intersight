package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageStoragePolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageStoragePolicyRead,
		Schema: map[string]*schema.Schema{
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
			"retain_policy_virtual_drives": {
				Description: "Retains the virtual drives defined in policy if they exist already. If this flag is false, the existing virtual drives are removed and created again based on virtual drives in the policy.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"unused_disks_state": {
				Description: "Unused Disks State is used to specify the state, unconfigured good or jbod, in which the disks that are not used in this policy should be moved.\n* `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group.\n* `Jbod` - JBOD state where the disks start showing up to host os.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceStorageStoragePolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceStorageStoragePolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageStoragePolicy{}
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
	if v, ok := d.GetOk("retain_policy_virtual_drives"); ok {
		x := (v.(bool))
		o.SetRetainPolicyVirtualDrives(x)
	}
	if v, ok := d.GetOk("unused_disks_state"); ok {
		x := (v.(string))
		o.SetUnusedDisksState(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageStoragePolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageStoragePolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageStoragePolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageStoragePolicyList.GetCount()
	var i int32
	var storageStoragePolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageStoragePolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageStoragePolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageStoragePolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageStoragePolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["disk_group_policies"] = flattenListStorageDiskGroupPolicyRelationship(s.GetDiskGroupPolicies(), d)

				temp["global_hot_spares"] = flattenListStorageLocalDisk(s.GetGlobalHotSpares(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["profiles"] = flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)
				temp["retain_policy_virtual_drives"] = (s.GetRetainPolicyVirtualDrives())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["unused_disks_state"] = (s.GetUnusedDisksState())

				temp["virtual_drives"] = flattenListStorageVirtualDriveConfig(s.GetVirtualDrives(), d)
				storageStoragePolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageStoragePolicyResults))
	if err := d.Set("results", storageStoragePolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageStoragePolicyResults[0]["moid"].(string))
	return de
}
