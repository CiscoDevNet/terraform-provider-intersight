package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexSoftwareVersionPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexSoftwareVersionPolicyRead,
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
			"hxdp_version": {
				Description: "Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_version": {
				Description: "Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster.",
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
			"server_firmware_version": {
				Description: "Desired server firmware version to apply on the HyperFlex Cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexSoftwareVersionPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexSoftwareVersionPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexSoftwareVersionPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("hxdp_version"); ok {
		x := (v.(string))
		o.SetHxdpVersion(x)
	}
	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
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
	if v, ok := d.GetOk("server_firmware_version"); ok {
		x := (v.(string))
		o.SetServerFirmwareVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexSoftwareVersionPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareVersionPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexSoftwareVersionPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexSoftwareVersionPolicyList.GetCount()
	var i int32
	var hyperflexSoftwareVersionPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexSoftwareVersionPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexSoftwareVersionPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexSoftwareVersionPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexSoftwareVersionPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_profiles"] = flattenListHyperflexClusterProfileRelationship(s.GetClusterProfiles(), d)
				temp["description"] = (s.GetDescription())
				temp["hxdp_version"] = (s.GetHxdpVersion())

				temp["hxdp_version_info"] = flattenMapSoftwareHyperflexDistributableRelationship(s.GetHxdpVersionInfo(), d)
				temp["hypervisor_version"] = (s.GetHypervisorVersion())

				temp["hypervisor_version_info"] = flattenMapSoftwareHyperflexDistributableRelationship(s.GetHypervisorVersionInfo(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["server_firmware_version"] = (s.GetServerFirmwareVersion())

				temp["server_firmware_version_info"] = flattenMapFirmwareDistributableRelationship(s.GetServerFirmwareVersionInfo(), d)

				temp["server_firmware_versions"] = flattenListHyperflexServerFirmwareVersionInfo(s.GetServerFirmwareVersions(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["upgrade_types"] = (s.GetUpgradeTypes())
				hyperflexSoftwareVersionPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexSoftwareVersionPolicyResults))
	if err := d.Set("results", hyperflexSoftwareVersionPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexSoftwareVersionPolicyResults[0]["moid"].(string))
	return de
}
