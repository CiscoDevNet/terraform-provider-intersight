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

func dataSourceCapabilityEquipmentSlotArray() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilityEquipmentSlotArrayRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"first_index": {
				Description: "First Index information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"height": {
				Description: "Height information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"horizontal_start_offset": {
				Description: "Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_group_separation": {
				Description: "Inline Group Separation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_group_size": {
				Description: "Inline Group Size information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"inline_offset": {
				Description: "Inline Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"location": {
				Description: "Location information for a Switch/Fabric-Interconnect hardware.",
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
			"number_of_slots": {
				Description: "Number of Slots information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"orientation": {
				Description: "Orientation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pid": {
				Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"selector": {
				Description: "Selector information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"sku": {
				Description: "SKU information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slots_per_line": {
				Description: "Slots per line information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"transverse_group_separation": {
				Description: "Transverse Group Separation information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"transverse_group_size": {
				Description: "Transverse Group Size information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"transverse_offset": {
				Description: "Transverse Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"vertical_start_offset": {
				Description: "Vertical Start Offset information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"vid": {
				Description: "VID information for Switch/Fabric-Interconnect.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"width": {
				Description: "Width information for a Switch/Fabric-Interconnect hardware.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceCapabilityEquipmentSlotArray().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCapabilityEquipmentSlotArrayRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilityEquipmentSlotArray{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("first_index"); ok {
		x := v.(float32)
		o.SetFirstIndex(x)
	}
	if v, ok := d.GetOk("height"); ok {
		x := v.(float32)
		o.SetHeight(x)
	}
	if v, ok := d.GetOk("horizontal_start_offset"); ok {
		x := v.(float32)
		o.SetHorizontalStartOffset(x)
	}
	if v, ok := d.GetOk("inline_group_separation"); ok {
		x := v.(float32)
		o.SetInlineGroupSeparation(x)
	}
	if v, ok := d.GetOk("inline_group_size"); ok {
		x := v.(float32)
		o.SetInlineGroupSize(x)
	}
	if v, ok := d.GetOk("inline_offset"); ok {
		x := v.(float32)
		o.SetInlineOffset(x)
	}
	if v, ok := d.GetOk("location"); ok {
		x := (v.(string))
		o.SetLocation(x)
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
	if v, ok := d.GetOk("number_of_slots"); ok {
		x := int64(v.(int))
		o.SetNumberOfSlots(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("orientation"); ok {
		x := (v.(string))
		o.SetOrientation(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("selector"); ok {
		x := (v.(string))
		o.SetSelector(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}
	if v, ok := d.GetOk("slots_per_line"); ok {
		x := int64(v.(int))
		o.SetSlotsPerLine(x)
	}
	if v, ok := d.GetOk("transverse_group_separation"); ok {
		x := v.(float32)
		o.SetTransverseGroupSeparation(x)
	}
	if v, ok := d.GetOk("transverse_group_size"); ok {
		x := v.(float32)
		o.SetTransverseGroupSize(x)
	}
	if v, ok := d.GetOk("transverse_offset"); ok {
		x := v.(float32)
		o.SetTransverseOffset(x)
	}
	if v, ok := d.GetOk("vertical_start_offset"); ok {
		x := v.(float32)
		o.SetVerticalStartOffset(x)
	}
	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}
	if v, ok := d.GetOk("width"); ok {
		x := v.(float32)
		o.SetWidth(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilityEquipmentSlotArray object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentSlotArrayList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of CapabilityEquipmentSlotArray: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of CapabilityEquipmentSlotArray: %s", responseErr.Error())
	}
	count := countResponse.CapabilityEquipmentSlotArrayList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for CapabilityEquipmentSlotArray data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var capabilityEquipmentSlotArrayResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityEquipmentSlotArrayList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching CapabilityEquipmentSlotArray: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching CapabilityEquipmentSlotArray: %s", responseErr.Error())
		}
		results := resMo.CapabilityEquipmentSlotArrayList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["first_index"] = (s.GetFirstIndex())
				temp["height"] = (s.GetHeight())
				temp["horizontal_start_offset"] = (s.GetHorizontalStartOffset())
				temp["inline_group_separation"] = (s.GetInlineGroupSeparation())
				temp["inline_group_size"] = (s.GetInlineGroupSize())
				temp["inline_offset"] = (s.GetInlineOffset())
				temp["location"] = (s.GetLocation())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["number_of_slots"] = (s.GetNumberOfSlots())
				temp["object_type"] = (s.GetObjectType())
				temp["orientation"] = (s.GetOrientation())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["pid"] = (s.GetPid())
				temp["selector"] = (s.GetSelector())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["sku"] = (s.GetSku())
				temp["slots_per_line"] = (s.GetSlotsPerLine())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["transverse_group_separation"] = (s.GetTransverseGroupSeparation())
				temp["transverse_group_size"] = (s.GetTransverseGroupSize())
				temp["transverse_offset"] = (s.GetTransverseOffset())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vertical_start_offset"] = (s.GetVerticalStartOffset())
				temp["vid"] = (s.GetVid())
				temp["width"] = (s.GetWidth())
				capabilityEquipmentSlotArrayResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilityEquipmentSlotArrayResults))
	if err := d.Set("results", capabilityEquipmentSlotArrayResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilityEquipmentSlotArrayResults[0]["moid"].(string))
	return de
}
