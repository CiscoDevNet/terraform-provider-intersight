package intersight

import (
	"context"
	"encoding/json"
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
			"additional_properties": {
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
		},
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
	resMo, _, responseErr := conn.ApiClient.HclApi.GetHclExemptedCatalogList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching HclExemptedCatalog: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HclExemptedCatalog list: %s", err.Error())
	}
	var s = &models.HclExemptedCatalogList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HclExemptedCatalog list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HclExemptedCatalog data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HclExemptedCatalog data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HclExemptedCatalog{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("comments", (s.GetComments())); err != nil {
				return diag.Errorf("error occurred while setting property Comments: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("os_vendor", (s.GetOsVendor())); err != nil {
				return diag.Errorf("error occurred while setting property OsVendor: %s", err.Error())
			}
			if err := d.Set("os_version", (s.GetOsVersion())); err != nil {
				return diag.Errorf("error occurred while setting property OsVersion: %s", err.Error())
			}
			if err := d.Set("processor_name", (s.GetProcessorName())); err != nil {
				return diag.Errorf("error occurred while setting property ProcessorName: %s", err.Error())
			}
			if err := d.Set("product_models", (s.GetProductModels())); err != nil {
				return diag.Errorf("error occurred while setting property ProductModels: %s", err.Error())
			}
			if err := d.Set("product_type", (s.GetProductType())); err != nil {
				return diag.Errorf("error occurred while setting property ProductType: %s", err.Error())
			}
			if err := d.Set("server_pid", (s.GetServerPid())); err != nil {
				return diag.Errorf("error occurred while setting property ServerPid: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("ucs_version", (s.GetUcsVersion())); err != nil {
				return diag.Errorf("error occurred while setting property UcsVersion: %s", err.Error())
			}
			if err := d.Set("version_type", (s.GetVersionType())); err != nil {
				return diag.Errorf("error occurred while setting property VersionType: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
