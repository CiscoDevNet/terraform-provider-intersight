package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCapabilitySiocModuleDescriptor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilitySiocModuleDescriptorRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Detailed information about the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "The model of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
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
			"revision": {
				Description: "Revision for the SIOC module.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vendor": {
				Description: "The vendor of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "The firmware or software version of the endpoint, for which this capability information is applicable.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCapabilitySiocModuleDescriptor().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCapabilitySiocModuleDescriptorRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilitySiocModuleDescriptor{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("revision"); ok {
		x := (v.(string))
		o.SetRevision(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilitySiocModuleDescriptor object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySiocModuleDescriptorList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CapabilitySiocModuleDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CapabilitySiocModuleDescriptorList.GetCount()
	var i int32
	var capabilitySiocModuleDescriptorResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySiocModuleDescriptorList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilitySiocModuleDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CapabilitySiocModuleDescriptorList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CapabilitySiocModuleDescriptor data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["capabilities"] = flattenListCapabilityCapabilityRelationship(s.GetCapabilities(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["revision"] = (s.GetRevision())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())
				capabilitySiocModuleDescriptorResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilitySiocModuleDescriptorResults))
	if err := d.Set("results", capabilitySiocModuleDescriptorResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilitySiocModuleDescriptorResults[0]["moid"].(string))
	return de
}
