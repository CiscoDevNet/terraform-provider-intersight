package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsInstall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsInstallRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "User provided description about the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"install_method": {
				Description: "The install method to be used for OS installation - vMedia, iPXE. \nOnly vMedia is supported as of now.\n* `vMedia` - OS image is mounted as vMedia in target server for OS installation.",
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
				Description: "The name of the OS install configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceOsInstall().Schema},
				Computed: true,
			}},
	}
}

func dataSourceOsInstallRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsInstall{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("install_method"); ok {
		x := (v.(string))
		o.SetInstallMethod(x)
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of OsInstall object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.OsApi.GetOsInstallList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of OsInstall: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.OsInstallList.GetCount()
	var i int32
	var osInstallResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.OsApi.GetOsInstallList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching OsInstall: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.OsInstallList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for OsInstall data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["additional_parameters"] = flattenListOsPlaceHolder(s.GetAdditionalParameters(), d)
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["answers"] = flattenMapOsAnswers(s.GetAnswers(), d)
				temp["class_id"] = (s.GetClassId())

				temp["configuration_file"] = flattenMapOsConfigurationFileRelationship(s.GetConfigurationFile(), d)
				temp["description"] = (s.GetDescription())

				temp["image"] = flattenMapSoftwarerepositoryOperatingSystemFileRelationship(s.GetImage(), d)
				temp["install_method"] = (s.GetInstallMethod())

				temp["install_target"] = flattenMapOsInstallTarget(s.GetInstallTarget(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["operating_system_parameters"] = flattenMapOsOperatingSystemParameters(s.GetOperatingSystemParameters(), d)

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["osdu_image"] = flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(s.GetOsduImage(), d)

				temp["server"] = flattenMapComputePhysicalRelationship(s.GetServer(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["workflow_info"] = flattenMapWorkflowWorkflowInfoRelationship(s.GetWorkflowInfo(), d)
				osInstallResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(osInstallResults))
	if err := d.Set("results", osInstallResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(osInstallResults[0]["moid"].(string))
	return de
}
