package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexHealthCheckPackageChecksum() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexHealthCheckPackageChecksumRead,
		Schema: map[string]*schema.Schema{
			"checksum": {
				Description: "SHA512 checksum of the health check package.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
				Description: "Name of the health check Debian package.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"package_name": {
				Description: "HyperFlex health check Debian package file name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"timestamp": {
				Description: "Timestamp of last update of HyperFlex health check package checksum.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "HyperFlex health check Debian Package Version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceHyperflexHealthCheckPackageChecksum().Schema},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexHealthCheckPackageChecksumRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckPackageChecksum{}
	if v, ok := d.GetOk("checksum"); ok {
		x := (v.(string))
		o.SetChecksum(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
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
	if v, ok := d.GetOk("package_name"); ok {
		x := (v.(string))
		o.SetPackageName(x)
	}
	if v, ok := d.GetOk("timestamp"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetTimestamp(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexHealthCheckPackageChecksum object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckPackageChecksumList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexHealthCheckPackageChecksum: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexHealthCheckPackageChecksumList.GetCount()
	var i int32
	var hyperflexHealthCheckPackageChecksumResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckPackageChecksumList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexHealthCheckPackageChecksum: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexHealthCheckPackageChecksumList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexHealthCheckPackageChecksum data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["checksum"] = (s.GetChecksum())
				temp["class_id"] = (s.GetClassId())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["package_name"] = (s.GetPackageName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["timestamp"] = (s.GetTimestamp()).String()
				temp["nr_version"] = (s.GetVersion())
				hyperflexHealthCheckPackageChecksumResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexHealthCheckPackageChecksumResults))
	if err := d.Set("results", hyperflexHealthCheckPackageChecksumResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexHealthCheckPackageChecksumResults[0]["moid"].(string))
	return de
}
