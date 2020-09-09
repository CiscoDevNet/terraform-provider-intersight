package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNiatelemetryTenant() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNiatelemetryTenantRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bfd_if_pol_count": {
				Description: "Bidirectional Forwarding Detection bfdIfPol Model Object count scale.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bfd_ifp_count": {
				Description: "Bidirectional Forwarding Detection Interface Policy count scale.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"dn": {
				Description: "Dn for the tenants present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fhs_bd_pol_count": {
				Description: "First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_rs_bd_to_fhs_count": {
				Description: "First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_site_connp_count": {
				Description: "Multi-Site scale for fvSiteConnp Model Object.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"l3_multicast_count": {
				Description: "Number of layer 3 multicasts.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"l3_multicast_ctx_count": {
				Description: "Number of layer 3 multicast CtxP.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"l3_multicast_if_count": {
				Description: "Number of layer 3 multicast IfP.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"l3out_count": {
				Description: "L3 out scale for the tenants present.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"site_name": {
				Description: "The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ssm": {
				Description: "SSM property feature usage.",
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
			"trace_route_ep_count": {
				Description: "Number of ITrace route endpoint per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"trace_route_ep_ext_count": {
				Description: "Number of ITrace endpoint external routes per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"trace_route_ext_ep_count": {
				Description: "Number of ITrace external endpoint routes per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"trace_route_ext_ext_count": {
				Description: "Number of ITrace external routes per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_abs_graph_count": {
				Description: "L4 to L7 Services graph count scale.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_backup_pol_count": {
				Description: "Policy Based Routing standby Node count scale. Checks the Policy Based Routing backup policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_redirect_dest_count": {
				Description: "Policy Based Routing standby Node count scale. Policy based redirect requires a destination to redirect traffic.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_svc_redirect_pol_count": {
				Description: "Policy Based Routing and Policy Based Service Insertion count scale. Includes without L4-L7 package.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vrf_count": {
				Description: "Vrf scale count per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}

func dataSourceNiatelemetryTenantRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewNiatelemetryTenantWithDefaults()
	if v, ok := d.GetOk("bfd_if_pol_count"); ok {
		x := int64(v.(int))
		o.SetBfdIfPolCount(x)
	}
	if v, ok := d.GetOk("bfd_ifp_count"); ok {
		x := int64(v.(int))
		o.SetBfdIfpCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fhs_bd_pol_count"); ok {
		x := int64(v.(int))
		o.SetFhsBdPolCount(x)
	}
	if v, ok := d.GetOk("fv_rs_bd_to_fhs_count"); ok {
		x := int64(v.(int))
		o.SetFvRsBdToFhsCount(x)
	}
	if v, ok := d.GetOk("fv_site_connp_count"); ok {
		x := int64(v.(int))
		o.SetFvSiteConnpCount(x)
	}
	if v, ok := d.GetOk("l3_multicast_count"); ok {
		x := int64(v.(int))
		o.SetL3MulticastCount(x)
	}
	if v, ok := d.GetOk("l3_multicast_ctx_count"); ok {
		x := int64(v.(int))
		o.SetL3MulticastCtxCount(x)
	}
	if v, ok := d.GetOk("l3_multicast_if_count"); ok {
		x := int64(v.(int))
		o.SetL3MulticastIfCount(x)
	}
	if v, ok := d.GetOk("l3out_count"); ok {
		x := int64(v.(int))
		o.SetL3outCount(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("ssm"); ok {
		x := (v.(string))
		o.SetSsm(x)
	}
	if v, ok := d.GetOk("trace_route_ep_count"); ok {
		x := int64(v.(int))
		o.SetTraceRouteEpCount(x)
	}
	if v, ok := d.GetOk("trace_route_ep_ext_count"); ok {
		x := int64(v.(int))
		o.SetTraceRouteEpExtCount(x)
	}
	if v, ok := d.GetOk("trace_route_ext_ep_count"); ok {
		x := int64(v.(int))
		o.SetTraceRouteExtEpCount(x)
	}
	if v, ok := d.GetOk("trace_route_ext_ext_count"); ok {
		x := int64(v.(int))
		o.SetTraceRouteExtExtCount(x)
	}
	if v, ok := d.GetOk("vns_abs_graph_count"); ok {
		x := int64(v.(int))
		o.SetVnsAbsGraphCount(x)
	}
	if v, ok := d.GetOk("vns_backup_pol_count"); ok {
		x := int64(v.(int))
		o.SetVnsBackupPolCount(x)
	}
	if v, ok := d.GetOk("vns_redirect_dest_count"); ok {
		x := int64(v.(int))
		o.SetVnsRedirectDestCount(x)
	}
	if v, ok := d.GetOk("vns_svc_redirect_pol_count"); ok {
		x := int64(v.(int))
		o.SetVnsSvcRedirectPolCount(x)
	}
	if v, ok := d.GetOk("vrf_count"); ok {
		x := int64(v.(int))
		o.SetVrfCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.NiatelemetryApi.GetNiatelemetryTenantList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.NiatelemetryTenantList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to NiatelemetryTenant: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewNiatelemetryTenantWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("bfd_if_pol_count", (s.GetBfdIfPolCount())); err != nil {
				return fmt.Errorf("error occurred while setting property BfdIfPolCount: %+v", err)
			}
			if err := d.Set("bfd_ifp_count", (s.GetBfdIfpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property BfdIfpCount: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return fmt.Errorf("error occurred while setting property Dn: %+v", err)
			}
			if err := d.Set("fhs_bd_pol_count", (s.GetFhsBdPolCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FhsBdPolCount: %+v", err)
			}
			if err := d.Set("fv_rs_bd_to_fhs_count", (s.GetFvRsBdToFhsCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FvRsBdToFhsCount: %+v", err)
			}
			if err := d.Set("fv_site_connp_count", (s.GetFvSiteConnpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property FvSiteConnpCount: %+v", err)
			}
			if err := d.Set("l3_multicast_count", (s.GetL3MulticastCount())); err != nil {
				return fmt.Errorf("error occurred while setting property L3MulticastCount: %+v", err)
			}
			if err := d.Set("l3_multicast_ctx_count", (s.GetL3MulticastCtxCount())); err != nil {
				return fmt.Errorf("error occurred while setting property L3MulticastCtxCount: %+v", err)
			}
			if err := d.Set("l3_multicast_if_count", (s.GetL3MulticastIfCount())); err != nil {
				return fmt.Errorf("error occurred while setting property L3MulticastIfCount: %+v", err)
			}
			if err := d.Set("l3out_count", (s.GetL3outCount())); err != nil {
				return fmt.Errorf("error occurred while setting property L3outCount: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property RegisteredDevice: %+v", err)
			}
			if err := d.Set("site_name", (s.GetSiteName())); err != nil {
				return fmt.Errorf("error occurred while setting property SiteName: %+v", err)
			}
			if err := d.Set("ssm", (s.GetSsm())); err != nil {
				return fmt.Errorf("error occurred while setting property Ssm: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("trace_route_ep_count", (s.GetTraceRouteEpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TraceRouteEpCount: %+v", err)
			}
			if err := d.Set("trace_route_ep_ext_count", (s.GetTraceRouteEpExtCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TraceRouteEpExtCount: %+v", err)
			}
			if err := d.Set("trace_route_ext_ep_count", (s.GetTraceRouteExtEpCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TraceRouteExtEpCount: %+v", err)
			}
			if err := d.Set("trace_route_ext_ext_count", (s.GetTraceRouteExtExtCount())); err != nil {
				return fmt.Errorf("error occurred while setting property TraceRouteExtExtCount: %+v", err)
			}
			if err := d.Set("vns_abs_graph_count", (s.GetVnsAbsGraphCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VnsAbsGraphCount: %+v", err)
			}
			if err := d.Set("vns_backup_pol_count", (s.GetVnsBackupPolCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VnsBackupPolCount: %+v", err)
			}
			if err := d.Set("vns_redirect_dest_count", (s.GetVnsRedirectDestCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VnsRedirectDestCount: %+v", err)
			}
			if err := d.Set("vns_svc_redirect_pol_count", (s.GetVnsSvcRedirectPolCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VnsSvcRedirectPolCount: %+v", err)
			}
			if err := d.Set("vrf_count", (s.GetVrfCount())); err != nil {
				return fmt.Errorf("error occurred while setting property VrfCount: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
