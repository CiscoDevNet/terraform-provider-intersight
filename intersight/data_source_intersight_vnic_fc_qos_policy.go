package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicFcQosPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicFcQosPolicyRead,
		Schema: map[string]*schema.Schema{
			"burst": {
				Description: "The burst traffic, in bytes, allowed on the vNIC.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cos": {
				Description: "Class of Service to be associated to the traffic on the virtual interface.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"max_data_field_size": {
				Description: "The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.",
				Type:        schema.TypeInt,
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
			"priority": {
				Description: "The priortity matching the System QoS specified in the fabric profile.\n* `Best Effort` - QoS Priority for Best-effort traffic.\n* `FC` - QoS Priority for FC traffic.\n* `Platinum` - QoS Priority for Platinum traffic.\n* `Gold` - QoS Priority for Gold traffic.\n* `Silver` - QoS Priority for Silver traffic.\n* `Bronze` - QoS Priority for Bronze traffic.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rate_limit": {
				Description: "The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicFcQosPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicFcQosPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicFcQosPolicy{}
	if v, ok := d.GetOk("burst"); ok {
		x := int64(v.(int))
		o.SetBurst(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cos"); ok {
		x := int64(v.(int))
		o.SetCos(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("max_data_field_size"); ok {
		x := int64(v.(int))
		o.SetMaxDataFieldSize(x)
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
	if v, ok := d.GetOk("priority"); ok {
		x := (v.(string))
		o.SetPriority(x)
	}
	if v, ok := d.GetOk("rate_limit"); ok {
		x := int64(v.(int))
		o.SetRateLimit(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicFcQosPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcQosPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicFcQosPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicFcQosPolicyList.GetCount()
	var i int32
	var vnicFcQosPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcQosPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicFcQosPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicFcQosPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicFcQosPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["burst"] = (s.GetBurst())
				temp["class_id"] = (s.GetClassId())
				temp["cos"] = (s.GetCos())
				temp["description"] = (s.GetDescription())
				temp["max_data_field_size"] = (s.GetMaxDataFieldSize())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["priority"] = (s.GetPriority())
				temp["rate_limit"] = (s.GetRateLimit())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				vnicFcQosPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicFcQosPolicyResults))
	if err := d.Set("results", vnicFcQosPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicFcQosPolicyResults[0]["moid"].(string))
	return de
}
