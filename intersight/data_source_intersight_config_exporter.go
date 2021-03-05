package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigExporter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigExporterRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"download_path": {
				Description: "Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired.",
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
				Description: "An identifier for the exporter instance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status of the export operation.\n* `` - The operation has not started.\n* `InProgress` - The operation is in progress.\n* `Success` - The operation has succeeded.\n* `Failed` - The operation has failed.\n* `RollBackInitiated` - The rollback has been inititiated for import failure.\n* `RollBackFailed` - The rollback has failed for import failure.\n* `RollbackCompleted` - The rollback has completed for import failure.\n* `RollbackAborted` - The rollback has been aborted for import failure.\n* `OperationTimedOut` - The operation has timed out.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status_message": {
				Description: "Status message associated with failures or progress indication.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceConfigExporter().Schema},
				Computed: true,
			}},
	}
}

func dataSourceConfigExporterRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ConfigExporter{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("download_path"); ok {
		x := (v.(string))
		o.SetDownloadPath(x)
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
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("status_message"); ok {
		x := (v.(string))
		o.SetStatusMessage(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ConfigExporter object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ConfigApi.GetConfigExporterList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ConfigExporter: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ConfigExporterList.GetCount()
	var i int32
	var configExporterResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ConfigApi.GetConfigExporterList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ConfigExporter: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ConfigExporterList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ConfigExporter data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["download_path"] = (s.GetDownloadPath())

				temp["exported_items"] = flattenListConfigExportedItemRelationship(s.GetExportedItems(), d)

				temp["items"] = flattenListConfigMoRef(s.GetItems(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["status"] = (s.GetStatus())
				temp["status_message"] = (s.GetStatusMessage())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				configExporterResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(configExporterResults))
	if err := d.Set("results", configExporterResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(configExporterResults[0]["moid"].(string))
	return de
}
