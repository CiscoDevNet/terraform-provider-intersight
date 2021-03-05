package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicFcAdapterPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicFcAdapterPolicyRead,
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
			"error_detection_timeout": {
				Description: "Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"io_throttle_count": {
				Description: "The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"lun_count": {
				Description: "The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"lun_queue_depth": {
				Description: "The number of commands that the HBA can send and receive in a single transmission per LUN.",
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
			"resource_allocation_timeout": {
				Description: "Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicFcAdapterPolicy().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicFcAdapterPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicFcAdapterPolicy{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("error_detection_timeout"); ok {
		x := int64(v.(int))
		o.SetErrorDetectionTimeout(x)
	}
	if v, ok := d.GetOk("io_throttle_count"); ok {
		x := int64(v.(int))
		o.SetIoThrottleCount(x)
	}
	if v, ok := d.GetOk("lun_count"); ok {
		x := int64(v.(int))
		o.SetLunCount(x)
	}
	if v, ok := d.GetOk("lun_queue_depth"); ok {
		x := int64(v.(int))
		o.SetLunQueueDepth(x)
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
	if v, ok := d.GetOk("resource_allocation_timeout"); ok {
		x := int64(v.(int))
		o.SetResourceAllocationTimeout(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicFcAdapterPolicy object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcAdapterPolicyList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicFcAdapterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicFcAdapterPolicyList.GetCount()
	var i int32
	var vnicFcAdapterPolicyResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcAdapterPolicyList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicFcAdapterPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicFcAdapterPolicyList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicFcAdapterPolicy data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["description"] = (s.GetDescription())
				temp["error_detection_timeout"] = (s.GetErrorDetectionTimeout())

				temp["error_recovery_settings"] = flattenMapVnicFcErrorRecoverySettings(s.GetErrorRecoverySettings(), d)

				temp["flogi_settings"] = flattenMapVnicFlogiSettings(s.GetFlogiSettings(), d)

				temp["interrupt_settings"] = flattenMapVnicFcInterruptSettings(s.GetInterruptSettings(), d)
				temp["io_throttle_count"] = (s.GetIoThrottleCount())
				temp["lun_count"] = (s.GetLunCount())
				temp["lun_queue_depth"] = (s.GetLunQueueDepth())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())

				temp["organization"] = flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)

				temp["plogi_settings"] = flattenMapVnicPlogiSettings(s.GetPlogiSettings(), d)
				temp["resource_allocation_timeout"] = (s.GetResourceAllocationTimeout())

				temp["rx_queue_settings"] = flattenMapVnicFcQueueSettings(s.GetRxQueueSettings(), d)

				temp["scsi_queue_settings"] = flattenMapVnicScsiQueueSettings(s.GetScsiQueueSettings(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["tx_queue_settings"] = flattenMapVnicFcQueueSettings(s.GetTxQueueSettings(), d)
				vnicFcAdapterPolicyResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicFcAdapterPolicyResults))
	if err := d.Set("results", vnicFcAdapterPolicyResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicFcAdapterPolicyResults[0]["moid"].(string))
	return de
}
