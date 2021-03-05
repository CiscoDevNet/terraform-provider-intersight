package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsConfigurationFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsConfigurationFileRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the OS ConfigurationFile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"file_content": {
				Description: "The content of the entire configuration file is stored as value. The content\ncan either be a static file content or a template content.\nThe template is expected to conform to the golang template syntax. The values\nfrom os.Answers properties will be used to populate this template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal": {
				Description: "The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page.",
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
				Description: "The name of the OS ConfigurationFile that uniquely identifies the configuration file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"supported": {
				Description: "An internal property that is used to distinguish between the pre-canned OS\nconfiguration file entries and user provided entries.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceOsConfigurationFile().Schema},
				Computed: true,
			}},
	}
}

func dataSourceOsConfigurationFileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsConfigurationFile{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("file_content"); ok {
		x := (v.(string))
		o.SetFileContent(x)
	}
	if v, ok := d.GetOk("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
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
	if v, ok := d.GetOk("supported"); ok {
		x := (v.(bool))
		o.SetSupported(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of OsConfigurationFile object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.OsApi.GetOsConfigurationFileList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.OsConfigurationFileList.GetCount()
	var i int32
	var osConfigurationFileResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.OsApi.GetOsConfigurationFileList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.OsConfigurationFileList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for OsConfigurationFile data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["catalog"] = flattenMapOsCatalogRelationship(s.GetCatalog(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())

				temp["distributions"] = flattenListHclOperatingSystemRelationship(s.GetDistributions(), d)
				temp["file_content"] = (s.GetFileContent())
				temp["internal"] = (s.GetInternal())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["placeholders"] = flattenListOsPlaceHolder(s.GetPlaceholders(), d)
				temp["supported"] = (s.GetSupported())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				osConfigurationFileResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(osConfigurationFileResults))
	if err := d.Set("results", osConfigurationFileResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(osConfigurationFileResults[0]["moid"].(string))
	return de
}
