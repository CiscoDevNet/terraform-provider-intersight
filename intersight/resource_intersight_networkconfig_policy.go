package intersight

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkconfigPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkconfigPolicyCreate,
		ReadContext:   resourceNetworkconfigPolicyRead,
		UpdateContext: resourceNetworkconfigPolicyUpdate,
		DeleteContext: resourceNetworkconfigPolicyDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"alternate_ipv4dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"alternate_ipv6dns_server": {
				Description: "IP address of the secondary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"appliance_account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dynamic_dns_domain": {
				Description: "The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enable_dynamic_dns": {
				Description: "If enabled, updates the resource records to the DNS from Cisco IMC.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv4dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6": {
				Description: "If enabled, allows to configure IPv6 properties.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"enable_ipv6dns_from_dhcp": {
				Description: "If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
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
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				ForceNew:   true,
			},
			"preferred_ipv4dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"preferred_ipv6dns_server": {
				Description: "IP address of the primary DNS server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
				Optional:    true,
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
							Computed:    true,
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
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
		},
	}
}

func resourceNetworkconfigPolicyCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewNetworkconfigPolicyWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("alternate_ipv4dns_server"); ok {
		x := (v.(string))
		o.SetAlternateIpv4dnsServer(x)
	}

	if v, ok := d.GetOk("alternate_ipv6dns_server"); ok {
		x := (v.(string))
		o.SetAlternateIpv6dnsServer(x)
	}

	if v, ok := d.GetOk("appliance_account"); ok {
		p := make([]models.IamAccountRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetApplianceAccount(x)
		}
	}

	o.SetClassId("networkconfig.Policy")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("dynamic_dns_domain"); ok {
		x := (v.(string))
		o.SetDynamicDnsDomain(x)
	}

	if v, ok := d.GetOkExists("enable_dynamic_dns"); ok {
		x := v.(bool)
		o.SetEnableDynamicDns(x)
	}

	if v, ok := d.GetOkExists("enable_ipv4dns_from_dhcp"); ok {
		x := v.(bool)
		o.SetEnableIpv4dnsFromDhcp(x)
	}

	if v, ok := d.GetOkExists("enable_ipv6"); ok {
		x := v.(bool)
		o.SetEnableIpv6(x)
	}

	if v, ok := d.GetOkExists("enable_ipv6dns_from_dhcp"); ok {
		x := v.(bool)
		o.SetEnableIpv6dnsFromDhcp(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("networkconfig.Policy")

	if v, ok := d.GetOk("organization"); ok {
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if v, ok := d.GetOk("preferred_ipv4dns_server"); ok {
		x := (v.(string))
		o.SetPreferredIpv4dnsServer(x)
	}

	if v, ok := d.GetOk("preferred_ipv6dns_server"); ok {
		x := (v.(string))
		o.SetPreferredIpv6dnsServer(x)
	}

	if v, ok := d.GetOk("profiles"); ok {
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	r := conn.ApiClient.NetworkconfigApi.CreateNetworkconfigPolicy(conn.ctx).NetworkconfigPolicy(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("failed while creating NetworkconfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceNetworkconfigPolicyRead(c, d, meta)...)
}
func detachNetworkconfigPolicyProfiles(d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkconfigPolicy{}
	o.SetClassId("networkconfig.Policy")
	o.SetObjectType("networkconfig.Policy")
	o.SetProfiles([]models.PolicyAbstractConfigProfileRelationship{})

	r := conn.ApiClient.NetworkconfigApi.UpdateNetworkconfigPolicy(conn.ctx, d.Id()).NetworkconfigPolicy(*o)
	_, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while detaching profile/profiles: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	return de
}

func resourceNetworkconfigPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.NetworkconfigApi.GetNetworkconfigPolicyByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "NetworkconfigPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching NetworkconfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("alternate_ipv4dns_server", (s.GetAlternateIpv4dnsServer())); err != nil {
		return diag.Errorf("error occurred while setting property AlternateIpv4dnsServer in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("alternate_ipv6dns_server", (s.GetAlternateIpv6dnsServer())); err != nil {
		return diag.Errorf("error occurred while setting property AlternateIpv6dnsServer in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("appliance_account", flattenMapIamAccountRelationship(s.GetApplianceAccount(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ApplianceAccount in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("dynamic_dns_domain", (s.GetDynamicDnsDomain())); err != nil {
		return diag.Errorf("error occurred while setting property DynamicDnsDomain in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("enable_dynamic_dns", (s.GetEnableDynamicDns())); err != nil {
		return diag.Errorf("error occurred while setting property EnableDynamicDns in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("enable_ipv4dns_from_dhcp", (s.GetEnableIpv4dnsFromDhcp())); err != nil {
		return diag.Errorf("error occurred while setting property EnableIpv4dnsFromDhcp in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("enable_ipv6", (s.GetEnableIpv6())); err != nil {
		return diag.Errorf("error occurred while setting property EnableIpv6 in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("enable_ipv6dns_from_dhcp", (s.GetEnableIpv6dnsFromDhcp())); err != nil {
		return diag.Errorf("error occurred while setting property EnableIpv6dnsFromDhcp in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("preferred_ipv4dns_server", (s.GetPreferredIpv4dnsServer())); err != nil {
		return diag.Errorf("error occurred while setting property PreferredIpv4dnsServer in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("preferred_ipv6dns_server", (s.GetPreferredIpv6dnsServer())); err != nil {
		return diag.Errorf("error occurred while setting property PreferredIpv6dnsServer in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.GetProfiles(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Profiles in NetworkconfigPolicy object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in NetworkconfigPolicy object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceNetworkconfigPolicyUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NetworkconfigPolicy{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("alternate_ipv4dns_server") {
		v := d.Get("alternate_ipv4dns_server")
		x := (v.(string))
		o.SetAlternateIpv4dnsServer(x)
	}

	if d.HasChange("alternate_ipv6dns_server") {
		v := d.Get("alternate_ipv6dns_server")
		x := (v.(string))
		o.SetAlternateIpv6dnsServer(x)
	}

	if d.HasChange("appliance_account") {
		v := d.Get("appliance_account")
		p := make([]models.IamAccountRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsIamAccountRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetApplianceAccount(x)
		}
	}

	o.SetClassId("networkconfig.Policy")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("dynamic_dns_domain") {
		v := d.Get("dynamic_dns_domain")
		x := (v.(string))
		o.SetDynamicDnsDomain(x)
	}

	if d.HasChange("enable_dynamic_dns") {
		v := d.Get("enable_dynamic_dns")
		x := (v.(bool))
		o.SetEnableDynamicDns(x)
	}

	if d.HasChange("enable_ipv4dns_from_dhcp") {
		v := d.Get("enable_ipv4dns_from_dhcp")
		x := (v.(bool))
		o.SetEnableIpv4dnsFromDhcp(x)
	}

	if d.HasChange("enable_ipv6") {
		v := d.Get("enable_ipv6")
		x := (v.(bool))
		o.SetEnableIpv6(x)
	}

	if d.HasChange("enable_ipv6dns_from_dhcp") {
		v := d.Get("enable_ipv6dns_from_dhcp")
		x := (v.(bool))
		o.SetEnableIpv6dnsFromDhcp(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("networkconfig.Policy")

	if d.HasChange("organization") {
		v := d.Get("organization")
		p := make([]models.OrganizationOrganizationRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOrganizationOrganizationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetOrganization(x)
		}
	}

	if d.HasChange("preferred_ipv4dns_server") {
		v := d.Get("preferred_ipv4dns_server")
		x := (v.(string))
		o.SetPreferredIpv4dnsServer(x)
	}

	if d.HasChange("preferred_ipv6dns_server") {
		v := d.Get("preferred_ipv6dns_server")
		x := (v.(string))
		o.SetPreferredIpv6dnsServer(x)
	}

	if d.HasChange("profiles") {
		v := d.Get("profiles")
		x := make([]models.PolicyAbstractConfigProfileRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsPolicyAbstractConfigProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetProfiles(x)
		}
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	r := conn.ApiClient.NetworkconfigApi.UpdateNetworkconfigPolicy(conn.ctx, d.Id()).NetworkconfigPolicy(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while updating NetworkconfigPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceNetworkconfigPolicyRead(c, d, meta)...)
}

func resourceNetworkconfigPolicyDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	if p, ok := d.GetOk("profiles"); ok {
		if len(p.([]interface{})) > 0 {
			e := detachNetworkconfigPolicyProfiles(d, meta)
			if e.HasError() {
				return e
			}
		}
	}
	p := conn.ApiClient.NetworkconfigApi.DeleteNetworkconfigPolicy(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		deleteErr := deleteErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while deleting NetworkconfigPolicy object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
