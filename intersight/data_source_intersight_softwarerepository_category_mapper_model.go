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

func dataSourceSoftwarerepositoryCategoryMapperModel() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSoftwarerepositoryCategoryMapperModelRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Description: "The category of the model series.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"dist_tag": {
				Description: "The distributable tag value of the model series.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"image_type": {
				Description: "The type of image based on the endpoint it can upgrade. For example, ucs-bundle-6400-infra.4.1.2a.bin can upgrade ucs managed fabric interconnects, so the image type is UCS Managed Fabric Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Description: "An unique identifer for a capability descriptor.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"regex_pattern": {
				Description: "The regex that all images of this model follow.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"series_id": {
				Description: "Cisco hardware model series.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceSoftwarerepositoryCategoryMapperModel().Schema},
				Computed: true,
			}},
	}
}

func dataSourceSoftwarerepositoryCategoryMapperModelRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.SoftwarerepositoryCategoryMapperModel{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("category"); ok {
		x := (v.(string))
		o.SetCategory(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("dist_tag"); ok {
		x := (v.(string))
		o.SetDistTag(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("image_type"); ok {
		x := (v.(string))
		o.SetImageType(x)
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
	if v, ok := d.GetOk("regex_pattern"); ok {
		x := (v.(string))
		o.SetRegexPattern(x)
	}
	if v, ok := d.GetOk("series_id"); ok {
		x := (v.(string))
		o.SetSeriesId(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of SoftwarerepositoryCategoryMapperModel object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategoryMapperModelList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of SoftwarerepositoryCategoryMapperModel: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of SoftwarerepositoryCategoryMapperModel: %s", responseErr.Error())
	}
	count := countResponse.SoftwarerepositoryCategoryMapperModelList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for SoftwarerepositoryCategoryMapperModel data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var softwarerepositoryCategoryMapperModelResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.SoftwarerepositoryApi.GetSoftwarerepositoryCategoryMapperModelList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching SoftwarerepositoryCategoryMapperModel: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching SoftwarerepositoryCategoryMapperModel: %s", responseErr.Error())
		}
		results := resMo.SoftwarerepositoryCategoryMapperModelList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["category"] = (s.GetCategory())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["dist_tag"] = (s.GetDistTag())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["image_type"] = (s.GetImageType())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["regex_pattern"] = (s.GetRegexPattern())
				temp["series_id"] = (s.GetSeriesId())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["supported_models"] = (s.GetSupportedModels())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				softwarerepositoryCategoryMapperModelResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(softwarerepositoryCategoryMapperModelResults))
	if err := d.Set("results", softwarerepositoryCategoryMapperModelResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(softwarerepositoryCategoryMapperModelResults[0]["moid"].(string))
	return de
}
