package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorageDiskGroupPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageDiskGroupPolicyRead,
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
			"raid_level": {
				Description: "The supported RAID level for the disk group.\n* `Raid0` - RAID 0 Stripe Raid Level.\n* `Raid1` - RAID 1 Mirror Raid Level.\n* `Raid5` - RAID 5 Mirror Raid Level.\n* `Raid6` - RAID 6 Mirror Raid Level.\n* `Raid10` - RAID 10 Mirror Raid Level.\n* `Raid50` - RAID 50 Mirror Raid Level.\n* `Raid60` - RAID 60 Mirror Raid Level.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"use_jbods": {
				Description: "Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceStorageDiskGroupPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceStorageDiskGroupPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageDiskGroupPolicy{}
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
	if v, ok := d.GetOk("raid_level"); ok {
		x := (v.(string))
		o.SetRaidLevel(x)
	}
	if v, ok := d.GetOk("use_jbods"); ok {
		x := (v.(bool))
		o.SetUseJbods(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageDiskGroupPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.StorageApi.GetStorageDiskGroupPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of StorageDiskGroupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.StorageDiskGroupPolicyList.GetCount()
	var i int32
	var storageDiskGroupPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageDiskGroupPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching StorageDiskGroupPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.StorageDiskGroupPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for StorageDiskGroupPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["dedicated_hot_spares"] = flattenListStorageLocalDisk(s.GetDedicatedHotSpares(), d)
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["raid_level"] = (s.GetRaidLevel())

				temp["span_groups"] = flattenListStorageSpanGroup(s.GetSpanGroups(), d)

				temp["storage_policies"] = flattenListStorageStoragePolicyRelationship(s.GetStoragePolicies(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["use_jbods"] = (s.GetUseJbods())
				storageDiskGroupPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(storageDiskGroupPolicyResults))
	if err := d.Set("results", storageDiskGroupPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(storageDiskGroupPolicyResults[0]["moid"].(string))
	return de
}
