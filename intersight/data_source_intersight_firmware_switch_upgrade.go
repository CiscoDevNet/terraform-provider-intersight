package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirmwareSwitchUpgrade() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareSwitchUpgradeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_fabric_evacuation": {
				Description: "The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.",
				Type:        schema.TypeBool,
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
			"skip_estimate_impact": {
				Description: "User has the option to skip the estimate impact calculation.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"status": {
				Description: "Status of the upgrade operation.\n* `NONE` - Upgrade status is not populated.\n* `IN_PROGRESS` - The upgrade is in progress.\n* `SUCCESSFUL` - The upgrade successfully completed.\n* `FAILED` - The upgrade shows failed status.\n* `TERMINATED` - The upgrade has been terminated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"upgrade_type": {
				Description: "Desired upgrade mode to choose either direct download based upgrade or network share upgrade.\n* `direct_upgrade` - Upgrade mode is direct download.\n* `network_upgrade` - Upgrade mode is network upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFirmwareSwitchUpgrade().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFirmwareSwitchUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareSwitchUpgrade{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("enable_fabric_evacuation"); ok {
		x := (v.(bool))
		o.SetEnableFabricEvacuation(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("skip_estimate_impact"); ok {
		x := (v.(bool))
		o.SetSkipEstimateImpact(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("upgrade_type"); ok {
		x := (v.(string))
		o.SetUpgradeType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareSwitchUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareSwitchUpgradeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FirmwareSwitchUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FirmwareSwitchUpgradeList.GetCount()
	var i int32
	var firmwareSwitchUpgradeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareSwitchUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FirmwareSwitchUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FirmwareSwitchUpgradeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FirmwareSwitchUpgrade data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)

				temp["direct_download"] = flattenMapFirmwareDirectDownload(s.GetDirectDownload(), d)

				temp["distributable"] = flattenMapFirmwareDistributableRelationship(s.GetDistributable(), d)
				temp["enable_fabric_evacuation"] = (s.GetEnableFabricEvacuation())

				temp["file_server"] = flattenMapSoftwarerepositoryFileServer(s.GetFileServer(), d)
				temp["moid"] = (s.GetMoid())

				temp["network_elements"] = flattenListNetworkElementRelationship(s.GetNetworkElements(), d)

				temp["network_share"] = flattenMapFirmwareNetworkShare(s.GetNetworkShare(), d)
				temp["object_type"] = (s.GetObjectType())

				temp["release"] = flattenMapSoftwarerepositoryReleaseRelationship(s.GetRelease(), d)
				temp["skip_estimate_impact"] = (s.GetSkipEstimateImpact())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["upgrade_impact"] = flattenMapFirmwareUpgradeImpactStatusRelationship(s.GetUpgradeImpact(), d)

				temp["upgrade_status"] = flattenMapFirmwareUpgradeStatusRelationship(s.GetUpgradeStatus(), d)
				temp["upgrade_type"] = (s.GetUpgradeType())
				firmwareSwitchUpgradeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(firmwareSwitchUpgradeResults))
	if err := d.Set("results", firmwareSwitchUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(firmwareSwitchUpgradeResults[0]["moid"].(string))
	return de
}
