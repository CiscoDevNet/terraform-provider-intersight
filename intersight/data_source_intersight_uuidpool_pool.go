package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUuidpoolPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUuidpoolPoolRead,
		Schema: map[string]*schema.Schema{
			"assigned": {
				Description: "Number of IDs that are currently assigned.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"assignment_order": {
				Description: "Assignment order decides the order in which the next identifier is allocated.\n* `sequential` - Identifiers are assigned in a sequential order.\n* `default` - Assignment order is decided by the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
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
			"prefix": {
				Description: "The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": {
				Description: "Total number of identifiers in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceUuidpoolPool().Schema},
				Computed: true,
			}},
	}
}

func dataSourceUuidpoolPoolRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.UuidpoolPool{}
	if v, ok := d.GetOk("assigned"); ok {
		x := int64(v.(int))
		o.SetAssigned(x)
	}
	if v, ok := d.GetOk("assignment_order"); ok {
		x := (v.(string))
		o.SetAssignmentOrder(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
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
	if v, ok := d.GetOk("prefix"); ok {
		x := (v.(string))
		o.SetPrefix(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of UuidpoolPool object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.UuidpoolApi.GetUuidpoolPoolList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of UuidpoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.UuidpoolPoolList.GetCount()
	var i int32
	var uuidpoolPoolResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.UuidpoolApi.GetUuidpoolPoolList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching UuidpoolPool: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.UuidpoolPoolList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for UuidpoolPool data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["assigned"] = (s.GetAssigned())
				temp["assignment_order"] = (s.GetAssignmentOrder())

				temp["block_heads"] = flattenListUuidpoolBlockRelationship(s.GetBlockHeads(), d)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)
				temp["prefix"] = (s.GetPrefix())
				temp["size"] = (s.GetSize())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["uuid_suffix_blocks"] = flattenListUuidpoolUuidBlock(s.GetUuidSuffixBlocks(), d)
				uuidpoolPoolResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(uuidpoolPoolResults))
	if err := d.Set("results", uuidpoolPoolResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(uuidpoolPoolResults[0]["moid"].(string))
	return de
}
