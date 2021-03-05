package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHclExemptedCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHclExemptedCatalogRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"comments": {
				Description: "Reason for the exemption.",
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
				Description: "A unique descriptive name of the exemption.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"os_vendor": {
				Description: "Vendor of the Operating System.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_version": {
				Description: "Version of the Operating system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_name": {
				Description: "Name of the processor supported for the server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"product_type": {
				Description: "Type of the product/adapter say GPU for graphic cards.\n* `` - Default type of the Product.\n* `Adapter` - Represents network adapter cards.\n* `StorageController` - Represents storage controllers.\n* `GPU` - Represents graphics cards.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"server_pid": {
				Description: "Three part ID representing the server model as returned by UCSM/CIMC XML APIs.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ucs_version": {
				Description: "Version of the UCS software.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_type": {
				Description: "Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"comments": {
						Description: "Reason for the exemption.",
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
						Description: "A unique descriptive name of the exemption.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"os_vendor": {
						Description: "Vendor of the Operating System.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"os_version": {
						Description: "Version of the Operating system.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"processor_name": {
						Description: "Name of the processor supported for the server.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"product_models": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"product_type": {
						Description: "Type of the product/adapter say GPU for graphic cards.\n* `` - Default type of the Product.\n* `Adapter` - Represents network adapter cards.\n* `StorageController` - Represents storage controllers.\n* `GPU` - Represents graphics cards.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"server_pid": {
						Description: "Three part ID representing the server model as returned by UCSM/CIMC XML APIs.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"ucs_version": {
						Description: "Version of the UCS software.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHclExemptedCatalogRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HclExemptedCatalog{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("comments"); ok {
		x := (v.(string))
		o.SetComments(x)
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
	if v, ok := d.GetOk("os_vendor"); ok {
		x := (v.(string))
		o.SetOsVendor(x)
	}
	if v, ok := d.GetOk("os_version"); ok {
		x := (v.(string))
		o.SetOsVersion(x)
	}
	if v, ok := d.GetOk("processor_name"); ok {
		x := (v.(string))
		o.SetProcessorName(x)
	}
	if v, ok := d.GetOk("product_type"); ok {
		x := (v.(string))
		o.SetProductType(x)
	}
	if v, ok := d.GetOk("server_pid"); ok {
		x := (v.(string))
		o.SetServerPid(x)
	}
	if v, ok := d.GetOk("ucs_version"); ok {
		x := (v.(string))
		o.SetUcsVersion(x)
	}
	if v, ok := d.GetOk("version_type"); ok {
		x := (v.(string))
		o.SetVersionType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HclExemptedCatalog object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HclApi.GetHclExemptedCatalogList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HclExemptedCatalog: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HclExemptedCatalogList.GetCount()
	var i int32
	var hclExemptedCatalogResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HclApi.GetHclExemptedCatalogList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HclExemptedCatalog: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HclExemptedCatalogList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HclExemptedCatalog data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["comments"] = (s.GetComments())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["os_vendor"] = (s.GetOsVendor())
				temp["os_version"] = (s.GetOsVersion())
				temp["processor_name"] = (s.GetProcessorName())
				temp["product_models"] = (s.GetProductModels())
				temp["product_type"] = (s.GetProductType())
				temp["server_pid"] = (s.GetServerPid())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["ucs_version"] = (s.GetUcsVersion())
				temp["version_type"] = (s.GetVersionType())
				hclExemptedCatalogResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hclExemptedCatalogResults))
	if err := d.Set("results", hclExemptedCatalogResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hclExemptedCatalogResults[0]["moid"].(string))
	return de
}
