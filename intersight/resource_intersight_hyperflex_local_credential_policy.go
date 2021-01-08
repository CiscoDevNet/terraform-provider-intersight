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

func resourceHyperflexLocalCredentialPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHyperflexLocalCredentialPolicyCreate,
		ReadContext:   resourceHyperflexLocalCredentialPolicyRead,
		UpdateContext: resourceHyperflexLocalCredentialPolicyUpdate,
		DeleteContext: resourceHyperflexLocalCredentialPolicyDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_profiles": {
				Description: "An array of relationships to hyperflexClusterProfile resources.",
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
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"factory_hypervisor_password": {
				Description: "Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"hxdp_root_pwd": {
				Description: "HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&*! special characters.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_admin": {
				Description: "Hypervisor administrator username must contain only alphanumeric characters.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_admin_pwd": {
				Description: "The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password \"Cisco123\" for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_hxdp_root_pwd_set": {
				Description: "Indicates whether the value of the 'hxdpRootPwd' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"is_hypervisor_admin_pwd_set": {
				Description: "Indicates whether the value of the 'hypervisorAdminPwd' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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

func resourceHyperflexLocalCredentialPolicyCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHyperflexLocalCredentialPolicyWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("hyperflex.LocalCredentialPolicy")

	if v, ok := d.GetOk("cluster_profiles"); ok {
		x := make([]models.HyperflexClusterProfileRelationship, 0)
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
			x = append(x, models.MoMoRefAsHyperflexClusterProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterProfiles(x)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOkExists("factory_hypervisor_password"); ok {
		x := v.(bool)
		o.SetFactoryHypervisorPassword(x)
	}

	if v, ok := d.GetOk("hxdp_root_pwd"); ok {
		x := (v.(string))
		o.SetHxdpRootPwd(x)
	}

	if v, ok := d.GetOk("hypervisor_admin"); ok {
		x := (v.(string))
		o.SetHypervisorAdmin(x)
	}

	if v, ok := d.GetOk("hypervisor_admin_pwd"); ok {
		x := (v.(string))
		o.SetHypervisorAdminPwd(x)
	}

	if v, ok := d.GetOkExists("is_hxdp_root_pwd_set"); ok {
		x := v.(bool)
		o.SetIsHxdpRootPwdSet(x)
	}

	if v, ok := d.GetOkExists("is_hypervisor_admin_pwd_set"); ok {
		x := v.(bool)
		o.SetIsHypervisorAdminPwdSet(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("hyperflex.LocalCredentialPolicy")

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

	r := conn.ApiClient.HyperflexApi.CreateHyperflexLocalCredentialPolicy(conn.ctx).HyperflexLocalCredentialPolicy(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating HyperflexLocalCredentialPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceHyperflexLocalCredentialPolicyRead(c, d, meta)
}

func resourceHyperflexLocalCredentialPolicyRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.HyperflexApi.GetHyperflexLocalCredentialPolicyByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexLocalCredentialPolicy object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching HyperflexLocalCredentialPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("cluster_profiles", flattenListHyperflexClusterProfileRelationship(s.GetClusterProfiles(), d)); err != nil {
		return diag.Errorf("error occurred while setting property ClusterProfiles in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("factory_hypervisor_password", (s.GetFactoryHypervisorPassword())); err != nil {
		return diag.Errorf("error occurred while setting property FactoryHypervisorPassword in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("hypervisor_admin", (s.GetHypervisorAdmin())); err != nil {
		return diag.Errorf("error occurred while setting property HypervisorAdmin in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("is_hxdp_root_pwd_set", (s.GetIsHxdpRootPwdSet())); err != nil {
		return diag.Errorf("error occurred while setting property IsHxdpRootPwdSet in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("is_hypervisor_admin_pwd_set", (s.GetIsHypervisorAdminPwdSet())); err != nil {
		return diag.Errorf("error occurred while setting property IsHypervisorAdminPwdSet in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.GetOrganization(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Organization in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HyperflexLocalCredentialPolicy object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHyperflexLocalCredentialPolicyUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.HyperflexLocalCredentialPolicy{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	o.SetClassId("hyperflex.LocalCredentialPolicy")

	if d.HasChange("cluster_profiles") {
		v := d.Get("cluster_profiles")
		x := make([]models.HyperflexClusterProfileRelationship, 0)
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
			x = append(x, models.MoMoRefAsHyperflexClusterProfileRelationship(o))
		}
		if len(x) > 0 {
			o.SetClusterProfiles(x)
		}
	}

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("factory_hypervisor_password") {
		v := d.Get("factory_hypervisor_password")
		x := (v.(bool))
		o.SetFactoryHypervisorPassword(x)
	}

	if d.HasChange("hxdp_root_pwd") {
		v := d.Get("hxdp_root_pwd")
		x := (v.(string))
		o.SetHxdpRootPwd(x)
	}

	if d.HasChange("hypervisor_admin") {
		v := d.Get("hypervisor_admin")
		x := (v.(string))
		o.SetHypervisorAdmin(x)
	}

	if d.HasChange("hypervisor_admin_pwd") {
		v := d.Get("hypervisor_admin_pwd")
		x := (v.(string))
		o.SetHypervisorAdminPwd(x)
	}

	if d.HasChange("is_hxdp_root_pwd_set") {
		v := d.Get("is_hxdp_root_pwd_set")
		x := (v.(bool))
		o.SetIsHxdpRootPwdSet(x)
	}

	if d.HasChange("is_hypervisor_admin_pwd_set") {
		v := d.Get("is_hypervisor_admin_pwd_set")
		x := (v.(bool))
		o.SetIsHypervisorAdminPwdSet(x)
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

	o.SetObjectType("hyperflex.LocalCredentialPolicy")

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

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexLocalCredentialPolicy(conn.ctx, d.Id()).HyperflexLocalCredentialPolicy(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating HyperflexLocalCredentialPolicy: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceHyperflexLocalCredentialPolicyRead(c, d, meta)
}

func resourceHyperflexLocalCredentialPolicyDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HyperflexApi.DeleteHyperflexLocalCredentialPolicy(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting HyperflexLocalCredentialPolicy object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
