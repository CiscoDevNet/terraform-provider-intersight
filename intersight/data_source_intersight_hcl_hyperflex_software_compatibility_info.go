package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHclHyperflexSoftwareCompatibilityInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHclHyperflexSoftwareCompatibilityInfoRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hxdp_version": {
				Description: "HXDP component software version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_version": {
				Description: "Hypervisor component software version.",
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
			"server_fw_version": {
				Description: "UCS Server Firmware component software version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHclHyperflexSoftwareCompatibilityInfo().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHclHyperflexSoftwareCompatibilityInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HclHyperflexSoftwareCompatibilityInfo{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("hxdp_version"); ok {
		x := (v.(string))
		o.SetHxdpVersion(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("server_fw_version"); ok {
		x := (v.(string))
		o.SetServerFwVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HclHyperflexSoftwareCompatibilityInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HclApi.GetHclHyperflexSoftwareCompatibilityInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HclHyperflexSoftwareCompatibilityInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HclHyperflexSoftwareCompatibilityInfoList.GetCount()
	var i int32
	var hclHyperflexSoftwareCompatibilityInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HclApi.GetHclHyperflexSoftwareCompatibilityInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HclHyperflexSoftwareCompatibilityInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HclHyperflexSoftwareCompatibilityInfoList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HclHyperflexSoftwareCompatibilityInfo data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["app_catalog"] = flattenMapHyperflexAppCatalogRelationship(s.GetAppCatalog(), d)
				temp["class_id"] = (s.GetClassId())

				temp["constraints"] = flattenListHclConstraint(s.GetConstraints(), d)
				temp["hxdp_version"] = (s.GetHxdpVersion())
				temp["hypervisor_type"] = (s.GetHypervisorType())
				temp["hypervisor_version"] = (s.GetHypervisorVersion())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["server_fw_version"] = (s.GetServerFwVersion())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				hclHyperflexSoftwareCompatibilityInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hclHyperflexSoftwareCompatibilityInfoResults))
	if err := d.Set("results", hclHyperflexSoftwareCompatibilityInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hclHyperflexSoftwareCompatibilityInfoResults[0]["moid"].(string))
	return de
}
