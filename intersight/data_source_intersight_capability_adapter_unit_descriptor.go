package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCapabilityAdapterUnitDescriptor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCapabilityAdapterUnitDescriptorRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connectivity_order": {
				Description: "Order in which the ports are connected; sequential or cyclic.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Detailed information about the endpoint.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ethernet_port_speed": {
				Description: "The port speed for ethernet ports in Mbps.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fibre_channel_port_speed": {
				Description: "The port speed for fibre channel ports in Mbps.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fibre_channel_scsi_ioq_limit": {
				Description: "The number of SCSI I/O Queue resources to allocate.",
				Type:        schema.TypeInt,
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
			"num_dce_ports": {
				Description: "Number of Dce Ports for the adaptor.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"prom_card_type": {
				Description: "Prom card type for the adaptor.",
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
				Elem:     &schema.Resource{Schema: resourceCapabilityAdapterUnitDescriptor().Schema},
				Computed: true,
			}},
	}
}

func dataSourceCapabilityAdapterUnitDescriptorRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilityAdapterUnitDescriptor{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connectivity_order"); ok {
		x := (v.(string))
		o.SetConnectivityOrder(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("ethernet_port_speed"); ok {
		x := int64(v.(int))
		o.SetEthernetPortSpeed(x)
	}
	if v, ok := d.GetOk("fibre_channel_port_speed"); ok {
		x := int64(v.(int))
		o.SetFibreChannelPortSpeed(x)
	}
	if v, ok := d.GetOk("fibre_channel_scsi_ioq_limit"); ok {
		x := int64(v.(int))
		o.SetFibreChannelScsiIoqLimit(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("num_dce_ports"); ok {
		x := int64(v.(int))
		o.SetNumDcePorts(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("prom_card_type"); ok {
		x := (v.(string))
		o.SetPromCardType(x)
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
		return diag.Errorf("json marshal of CapabilityAdapterUnitDescriptor object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityAdapterUnitDescriptorList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CapabilityAdapterUnitDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CapabilityAdapterUnitDescriptorList.GetCount()
	var i int32
	var capabilityAdapterUnitDescriptorResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilityAdapterUnitDescriptorList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CapabilityAdapterUnitDescriptor: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CapabilityAdapterUnitDescriptorList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CapabilityAdapterUnitDescriptor data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["capabilities"] = flattenListCapabilityCapabilityRelationship(s.GetCapabilities(), d)
				temp["class_id"] = (s.GetClassId())
				temp["connectivity_order"] = (s.GetConnectivityOrder())
				temp["description"] = (s.GetDescription())
				temp["ethernet_port_speed"] = (s.GetEthernetPortSpeed())
				temp["fibre_channel_port_speed"] = (s.GetFibreChannelPortSpeed())
				temp["fibre_channel_scsi_ioq_limit"] = (s.GetFibreChannelScsiIoqLimit())
				temp["model"] = (s.GetModel())
				temp["moid"] = (s.GetMoid())
				temp["num_dce_ports"] = (s.GetNumDcePorts())
				temp["object_type"] = (s.GetObjectType())
				temp["prom_card_type"] = (s.GetPromCardType())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["vendor"] = (s.GetVendor())
				temp["nr_version"] = (s.GetVersion())
				capabilityAdapterUnitDescriptorResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilityAdapterUnitDescriptorResults))
	if err := d.Set("results", capabilityAdapterUnitDescriptorResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilityAdapterUnitDescriptorResults[0]["moid"].(string))
	return de
}
