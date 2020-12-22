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

func dataSourceNiatelemetryTenant() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryTenantRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"bfd_if_pol_count": {
				Description: "Number of Bidirectional Forwarding Detection bfdIfPol Model Objects.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"bfd_ifp_count": {
				Description: "Number of objects with Bidirectional Forwarding Detection Interface Policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dhcp_rs_prov_count": {
				Description: "Number of tenants with Dynamic Host Configuration Protocol enabled.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"dn": {
				Description: "Dn for the tenants present.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"fhs_bd_pol_count": {
				Description: "Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_ap_count": {
				Description: "Number of application profiles per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_bd_count": {
				Description: "Number of bridge domains with hardware proxy enabled per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_bd_subnet_count": {
				Description: "Number of bridge domain subnets per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_bdno_arp_count": {
				Description: "Number of bridge domains with ARP flodding disabled per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_cep_count": {
				Description: "Count of number of endpoints per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_rs_bd_to_fhs_count": {
				Description: "Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_rs_bd_to_out_count": {
				Description: "Number of bridge domains connected to layer 3 out per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_site_connp_count": {
				Description: "Number of Multi-Site objects.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"fv_subnet_count": {
				Description: "Number of subnets per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ip_static_route_count": {
				Description: "Number of IP static routes per tenant.",
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
				Description: "Number of L3 out objects for the tenants present.",
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"qos_custom_pol_count": {
				Description: "Number of Quality Of Service Custom Policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"record_type": {
				Description: "Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"record_version": {
				Description: "Version of record being pushed. This determines what was the API version for data available from the device.",
				Type:        schema.TypeString,
				Optional:    true,
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
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
			"ssm_count": {
				Description: "Number of context-level ssm translate policies per tenant.",
				Type:        schema.TypeInt,
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
			"tcam_opt_count": {
				Description: "Number of TCAM optimization enabled per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
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
				Description: "Number of objects with L4 to L7 Services graph.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_backup_pol_count": {
				Description: "Number of objects with Policy Based Routing standby Node. Checks the Policy Based Routing backup policy.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_redirect_dest_count": {
				Description: "Number of objects with Policy Based Routing standby Node. Policy based redirect requires a destination to redirect traffic.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vns_svc_redirect_pol_count": {
				Description: "Number of Policy Based Routing and Policy Based Service Insertion objects. Includes without L4-L7 package.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vrf_count": {
				Description: "Number of Vrfs per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vz_br_cp_count": {
				Description: "Number of Zoning Policy per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vz_rt_cons_count": {
				Description: "Number of Client Contract between End Point Groups per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"vz_rt_prov_count": {
				Description: "Number of Client Contract between End Point Groups per tenant.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}
}

func dataSourceNiatelemetryTenantRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryTenant{}
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
	if v, ok := d.GetOk("dhcp_rs_prov_count"); ok {
		x := int64(v.(int))
		o.SetDhcpRsProvCount(x)
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.SetDn(x)
	}
	if v, ok := d.GetOk("fhs_bd_pol_count"); ok {
		x := int64(v.(int))
		o.SetFhsBdPolCount(x)
	}
	if v, ok := d.GetOk("fv_ap_count"); ok {
		x := int64(v.(int))
		o.SetFvApCount(x)
	}
	if v, ok := d.GetOk("fv_bd_count"); ok {
		x := int64(v.(int))
		o.SetFvBdCount(x)
	}
	if v, ok := d.GetOk("fv_bd_subnet_count"); ok {
		x := int64(v.(int))
		o.SetFvBdSubnetCount(x)
	}
	if v, ok := d.GetOk("fv_bdno_arp_count"); ok {
		x := int64(v.(int))
		o.SetFvBdnoArpCount(x)
	}
	if v, ok := d.GetOk("fv_cep_count"); ok {
		x := int64(v.(int))
		o.SetFvCepCount(x)
	}
	if v, ok := d.GetOk("fv_rs_bd_to_fhs_count"); ok {
		x := int64(v.(int))
		o.SetFvRsBdToFhsCount(x)
	}
	if v, ok := d.GetOk("fv_rs_bd_to_out_count"); ok {
		x := int64(v.(int))
		o.SetFvRsBdToOutCount(x)
	}
	if v, ok := d.GetOk("fv_site_connp_count"); ok {
		x := int64(v.(int))
		o.SetFvSiteConnpCount(x)
	}
	if v, ok := d.GetOk("fv_subnet_count"); ok {
		x := int64(v.(int))
		o.SetFvSubnetCount(x)
	}
	if v, ok := d.GetOk("ip_static_route_count"); ok {
		x := int64(v.(int))
		o.SetIpStaticRouteCount(x)
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
	if v, ok := d.GetOk("qos_custom_pol_count"); ok {
		x := int64(v.(int))
		o.SetQosCustomPolCount(x)
	}
	if v, ok := d.GetOk("record_type"); ok {
		x := (v.(string))
		o.SetRecordType(x)
	}
	if v, ok := d.GetOk("record_version"); ok {
		x := (v.(string))
		o.SetRecordVersion(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("ssm"); ok {
		x := (v.(string))
		o.SetSsm(x)
	}
	if v, ok := d.GetOk("ssm_count"); ok {
		x := int64(v.(int))
		o.SetSsmCount(x)
	}
	if v, ok := d.GetOk("tcam_opt_count"); ok {
		x := int64(v.(int))
		o.SetTcamOptCount(x)
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
	if v, ok := d.GetOk("vz_br_cp_count"); ok {
		x := int64(v.(int))
		o.SetVzBrCpCount(x)
	}
	if v, ok := d.GetOk("vz_rt_cons_count"); ok {
		x := int64(v.(int))
		o.SetVzRtConsCount(x)
	}
	if v, ok := d.GetOk("vz_rt_prov_count"); ok {
		x := int64(v.(int))
		o.SetVzRtProvCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryTenant object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryTenantList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching NiatelemetryTenant: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NiatelemetryTenant list: %s", err.Error())
	}
	var s = &models.NiatelemetryTenantList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NiatelemetryTenant list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for NiatelemetryTenant did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiatelemetryTenant{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("bfd_if_pol_count", (s.GetBfdIfPolCount())); err != nil {
				return diag.Errorf("error occurred while setting property BfdIfPolCount: %s", err.Error())
			}
			if err := d.Set("bfd_ifp_count", (s.GetBfdIfpCount())); err != nil {
				return diag.Errorf("error occurred while setting property BfdIfpCount: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("dhcp_rs_prov_count", (s.GetDhcpRsProvCount())); err != nil {
				return diag.Errorf("error occurred while setting property DhcpRsProvCount: %s", err.Error())
			}
			if err := d.Set("dn", (s.GetDn())); err != nil {
				return diag.Errorf("error occurred while setting property Dn: %s", err.Error())
			}
			if err := d.Set("fhs_bd_pol_count", (s.GetFhsBdPolCount())); err != nil {
				return diag.Errorf("error occurred while setting property FhsBdPolCount: %s", err.Error())
			}
			if err := d.Set("fv_ap_count", (s.GetFvApCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvApCount: %s", err.Error())
			}
			if err := d.Set("fv_bd_count", (s.GetFvBdCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvBdCount: %s", err.Error())
			}
			if err := d.Set("fv_bd_subnet_count", (s.GetFvBdSubnetCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvBdSubnetCount: %s", err.Error())
			}
			if err := d.Set("fv_bdno_arp_count", (s.GetFvBdnoArpCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvBdnoArpCount: %s", err.Error())
			}
			if err := d.Set("fv_cep_count", (s.GetFvCepCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvCepCount: %s", err.Error())
			}
			if err := d.Set("fv_rs_bd_to_fhs_count", (s.GetFvRsBdToFhsCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvRsBdToFhsCount: %s", err.Error())
			}
			if err := d.Set("fv_rs_bd_to_out_count", (s.GetFvRsBdToOutCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvRsBdToOutCount: %s", err.Error())
			}
			if err := d.Set("fv_site_connp_count", (s.GetFvSiteConnpCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvSiteConnpCount: %s", err.Error())
			}
			if err := d.Set("fv_subnet_count", (s.GetFvSubnetCount())); err != nil {
				return diag.Errorf("error occurred while setting property FvSubnetCount: %s", err.Error())
			}
			if err := d.Set("ip_static_route_count", (s.GetIpStaticRouteCount())); err != nil {
				return diag.Errorf("error occurred while setting property IpStaticRouteCount: %s", err.Error())
			}
			if err := d.Set("l3_multicast_count", (s.GetL3MulticastCount())); err != nil {
				return diag.Errorf("error occurred while setting property L3MulticastCount: %s", err.Error())
			}
			if err := d.Set("l3_multicast_ctx_count", (s.GetL3MulticastCtxCount())); err != nil {
				return diag.Errorf("error occurred while setting property L3MulticastCtxCount: %s", err.Error())
			}
			if err := d.Set("l3_multicast_if_count", (s.GetL3MulticastIfCount())); err != nil {
				return diag.Errorf("error occurred while setting property L3MulticastIfCount: %s", err.Error())
			}
			if err := d.Set("l3out_count", (s.GetL3outCount())); err != nil {
				return diag.Errorf("error occurred while setting property L3outCount: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("qos_custom_pol_count", (s.GetQosCustomPolCount())); err != nil {
				return diag.Errorf("error occurred while setting property QosCustomPolCount: %s", err.Error())
			}
			if err := d.Set("record_type", (s.GetRecordType())); err != nil {
				return diag.Errorf("error occurred while setting property RecordType: %s", err.Error())
			}
			if err := d.Set("record_version", (s.GetRecordVersion())); err != nil {
				return diag.Errorf("error occurred while setting property RecordVersion: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("site_name", (s.GetSiteName())); err != nil {
				return diag.Errorf("error occurred while setting property SiteName: %s", err.Error())
			}
			if err := d.Set("ssm", (s.GetSsm())); err != nil {
				return diag.Errorf("error occurred while setting property Ssm: %s", err.Error())
			}
			if err := d.Set("ssm_count", (s.GetSsmCount())); err != nil {
				return diag.Errorf("error occurred while setting property SsmCount: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("tcam_opt_count", (s.GetTcamOptCount())); err != nil {
				return diag.Errorf("error occurred while setting property TcamOptCount: %s", err.Error())
			}
			if err := d.Set("trace_route_ep_count", (s.GetTraceRouteEpCount())); err != nil {
				return diag.Errorf("error occurred while setting property TraceRouteEpCount: %s", err.Error())
			}
			if err := d.Set("trace_route_ep_ext_count", (s.GetTraceRouteEpExtCount())); err != nil {
				return diag.Errorf("error occurred while setting property TraceRouteEpExtCount: %s", err.Error())
			}
			if err := d.Set("trace_route_ext_ep_count", (s.GetTraceRouteExtEpCount())); err != nil {
				return diag.Errorf("error occurred while setting property TraceRouteExtEpCount: %s", err.Error())
			}
			if err := d.Set("trace_route_ext_ext_count", (s.GetTraceRouteExtExtCount())); err != nil {
				return diag.Errorf("error occurred while setting property TraceRouteExtExtCount: %s", err.Error())
			}
			if err := d.Set("vns_abs_graph_count", (s.GetVnsAbsGraphCount())); err != nil {
				return diag.Errorf("error occurred while setting property VnsAbsGraphCount: %s", err.Error())
			}
			if err := d.Set("vns_backup_pol_count", (s.GetVnsBackupPolCount())); err != nil {
				return diag.Errorf("error occurred while setting property VnsBackupPolCount: %s", err.Error())
			}
			if err := d.Set("vns_redirect_dest_count", (s.GetVnsRedirectDestCount())); err != nil {
				return diag.Errorf("error occurred while setting property VnsRedirectDestCount: %s", err.Error())
			}
			if err := d.Set("vns_svc_redirect_pol_count", (s.GetVnsSvcRedirectPolCount())); err != nil {
				return diag.Errorf("error occurred while setting property VnsSvcRedirectPolCount: %s", err.Error())
			}
			if err := d.Set("vrf_count", (s.GetVrfCount())); err != nil {
				return diag.Errorf("error occurred while setting property VrfCount: %s", err.Error())
			}
			if err := d.Set("vz_br_cp_count", (s.GetVzBrCpCount())); err != nil {
				return diag.Errorf("error occurred while setting property VzBrCpCount: %s", err.Error())
			}
			if err := d.Set("vz_rt_cons_count", (s.GetVzRtConsCount())); err != nil {
				return diag.Errorf("error occurred while setting property VzRtConsCount: %s", err.Error())
			}
			if err := d.Set("vz_rt_prov_count", (s.GetVzRtProvCount())); err != nil {
				return diag.Errorf("error occurred while setting property VzRtProvCount: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
