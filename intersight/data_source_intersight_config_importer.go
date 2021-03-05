package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigImporter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigImporterRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_path": {
				Description: "The path to the archive in Intersight storage that has all the MOs\nto be imported.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"import_source": {
				Description: "The source of the archive in Intersight storage that has all the MOs\nto be imported.\n* `ImageRepo` - The 'ImageRepo' source if the source of exporter archive is image repository.\n* `URL` - The 'URL' source if the source of exported archive is a URL.",
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
				Description: "An identifier for the importer instance.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"skip_integrity_checks": {
				Description: "Specifies whether integrity checks must be skipped.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"status": {
				Description: "Status of the import operation.\n* `` - The operation has not started.\n* `InProgress` - The operation is in progress.\n* `Success` - The operation has succeeded.\n* `Failed` - The operation has failed.\n* `RollBackInitiated` - The rollback has been inititiated for import failure.\n* `RollBackFailed` - The rollback has failed for import failure.\n* `RollbackCompleted` - The rollback has completed for import failure.\n* `RollbackAborted` - The rollback has been aborted for import failure.\n* `OperationTimedOut` - The operation has timed out.",
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
				Elem:     &schema.Resource{Schema: resourceConfigImporter().Schema},
				Computed: true,
			}},
	}
}

func dataSourceConfigImporterRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ConfigImporter{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("import_path"); ok {
		x := (v.(string))
		o.SetImportPath(x)
	}
	if v, ok := d.GetOk("import_source"); ok {
		x := (v.(string))
		o.SetImportSource(x)
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
	if v, ok := d.GetOk("skip_integrity_checks"); ok {
		x := (v.(bool))
		o.SetSkipIntegrityChecks(x)
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
		return diag.Errorf("json marshal of ConfigImporter object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ConfigApi.GetConfigImporterList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of ConfigImporter: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.ConfigImporterList.GetCount()
	var i int32
	var configImporterResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ConfigApi.GetConfigImporterList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching ConfigImporter: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.ConfigImporterList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for ConfigImporter data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["import_path"] = (s.GetImportPath())
				temp["import_source"] = (s.GetImportSource())

				temp["imported_items"] = flattenListConfigImportedItemRelationship(s.GetImportedItems(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["skip_integrity_checks"] = (s.GetSkipIntegrityChecks())
				temp["status"] = (s.GetStatus())
				temp["status_message"] = (s.GetStatusMessage())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				configImporterResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(configImporterResults))
	if err := d.Set("results", configImporterResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(configImporterResults[0]["moid"].(string))
	return de
}
