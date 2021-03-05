package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirmwareUnsupportedVersionUpgrade() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareUnsupportedVersionUpgradeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_error": {
				Description: "Any error encountered. Set to empty when download is in progress or completed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_progress": {
				Description: "The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"download_retries": {
				Description: "The number of retries the plugin attempted before succeeding or failing the download.",
				Type:        schema.TypeInt,
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
			"upgrade_status": {
				Description: "Workflow status of firmware upgrade.\n* `None` - Upgrade status is none when upgrade is in progress.\n* `Completed` - Upgrade completed successfully.\n* `Failed` - Upgrade status is failed when upgrade has failed.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFirmwareUnsupportedVersionUpgrade().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFirmwareUnsupportedVersionUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareUnsupportedVersionUpgrade{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("download_error"); ok {
		x := (v.(string))
		o.SetDownloadError(x)
	}
	if v, ok := d.GetOk("download_progress"); ok {
		x := int64(v.(int))
		o.SetDownloadProgress(x)
	}
	if v, ok := d.GetOk("download_retries"); ok {
		x := int64(v.(int))
		o.SetDownloadRetries(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("upgrade_status"); ok {
		x := (v.(string))
		o.SetUpgradeStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareUnsupportedVersionUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareUnsupportedVersionUpgradeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of FirmwareUnsupportedVersionUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.FirmwareUnsupportedVersionUpgradeList.GetCount()
	var i int32
	var firmwareUnsupportedVersionUpgradeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareUnsupportedVersionUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching FirmwareUnsupportedVersionUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.FirmwareUnsupportedVersionUpgradeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for FirmwareUnsupportedVersionUpgrade data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["checksum"] = flattenMapConnectorFileChecksum(s.GetChecksum(), d)
				temp["class_id"] = (s.GetClassId())

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)

				temp["distributable"] = flattenMapFirmwareDistributableRelationship(s.GetDistributable(), d)
				temp["download_error"] = (s.GetDownloadError())
				temp["download_progress"] = (s.GetDownloadProgress())
				temp["download_retries"] = (s.GetDownloadRetries())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["physical_identity"] = flattenMapEquipmentPhysicalIdentityRelationship(s.GetPhysicalIdentity(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["upgrade_status"] = (s.GetUpgradeStatus())
				firmwareUnsupportedVersionUpgradeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(firmwareUnsupportedVersionUpgradeResults))
	if err := d.Set("results", firmwareUnsupportedVersionUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(firmwareUnsupportedVersionUpgradeResults[0]["moid"].(string))
	return de
}
