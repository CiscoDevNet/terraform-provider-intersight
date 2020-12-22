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

func resourceHclHyperflexSoftwareCompatibilityInfo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHclHyperflexSoftwareCompatibilityInfoCreate,
		ReadContext:   resourceHclHyperflexSoftwareCompatibilityInfoRead,
		UpdateContext: resourceHclHyperflexSoftwareCompatibilityInfoUpdate,
		DeleteContext: resourceHclHyperflexSoftwareCompatibilityInfoDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"app_catalog": {
				Description: "A reference to a hyperflexAppCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"constraints": {
				Type:     schema.TypeList,
				Optional: true,
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
						"constraint_name": {
							Description: "Name or key of the applicable compatibility constraint.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"constraint_value": {
							Description: "Value of the applicable compatibility constraint. Could either be a string value or a regex.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
			},
			"hxdp_version": {
				Description: "HXDP component software version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_type": {
				Description: "Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ESXi",
			},
			"hypervisor_version": {
				Description: "Hypervisor component software version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"server_fw_version": {
				Description: "UCS Server Firmware component software version.",
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
		},
	}
}

func resourceHclHyperflexSoftwareCompatibilityInfoCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewHclHyperflexSoftwareCompatibilityInfoWithDefaults()
	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("app_catalog"); ok {
		p := make([]models.HyperflexAppCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsHyperflexAppCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAppCatalog(x)
		}
	}

	o.SetClassId("hcl.HyperflexSoftwareCompatibilityInfo")

	if v, ok := d.GetOk("constraints"); ok {
		x := make([]models.HclConstraint, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewHclConstraintWithDefaults()
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
			o.SetClassId("hcl.Constraint")
			if v, ok := l["constraint_name"]; ok {
				{
					x := (v.(string))
					o.SetConstraintName(x)
				}
			}
			if v, ok := l["constraint_value"]; ok {
				{
					x := (v.(string))
					o.SetConstraintValue(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetConstraints(x)
		}
	}

	if v, ok := d.GetOk("hxdp_version"); ok {
		x := (v.(string))
		o.SetHxdpVersion(x)
	}

	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}

	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("hcl.HyperflexSoftwareCompatibilityInfo")

	if v, ok := d.GetOk("server_fw_version"); ok {
		x := (v.(string))
		o.SetServerFwVersion(x)
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

	r := conn.ApiClient.HclApi.CreateHclHyperflexSoftwareCompatibilityInfo(conn.ctx).HclHyperflexSoftwareCompatibilityInfo(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating HclHyperflexSoftwareCompatibilityInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceHclHyperflexSoftwareCompatibilityInfoRead(c, d, meta)
}

func resourceHclHyperflexSoftwareCompatibilityInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.HclApi.GetHclHyperflexSoftwareCompatibilityInfoByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HclHyperflexSoftwareCompatibilityInfo object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching HclHyperflexSoftwareCompatibilityInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("app_catalog", flattenMapHyperflexAppCatalogRelationship(s.GetAppCatalog(), d)); err != nil {
		return diag.Errorf("error occurred while setting property AppCatalog in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("constraints", flattenListHclConstraint(s.GetConstraints(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Constraints in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("hxdp_version", (s.GetHxdpVersion())); err != nil {
		return diag.Errorf("error occurred while setting property HxdpVersion in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("hypervisor_type", (s.GetHypervisorType())); err != nil {
		return diag.Errorf("error occurred while setting property HypervisorType in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("hypervisor_version", (s.GetHypervisorVersion())); err != nil {
		return diag.Errorf("error occurred while setting property HypervisorVersion in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("server_fw_version", (s.GetServerFwVersion())); err != nil {
		return diag.Errorf("error occurred while setting property ServerFwVersion in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HclHyperflexSoftwareCompatibilityInfo object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHclHyperflexSoftwareCompatibilityInfoUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.HclHyperflexSoftwareCompatibilityInfo{}
	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("app_catalog") {
		v := d.Get("app_catalog")
		p := make([]models.HyperflexAppCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsHyperflexAppCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetAppCatalog(x)
		}
	}

	o.SetClassId("hcl.HyperflexSoftwareCompatibilityInfo")

	if d.HasChange("constraints") {
		v := d.Get("constraints")
		x := make([]models.HclConstraint, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.HclConstraint{}
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
			o.SetClassId("hcl.Constraint")
			if v, ok := l["constraint_name"]; ok {
				{
					x := (v.(string))
					o.SetConstraintName(x)
				}
			}
			if v, ok := l["constraint_value"]; ok {
				{
					x := (v.(string))
					o.SetConstraintValue(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetConstraints(x)
		}
	}

	if d.HasChange("hxdp_version") {
		v := d.Get("hxdp_version")
		x := (v.(string))
		o.SetHxdpVersion(x)
	}

	if d.HasChange("hypervisor_type") {
		v := d.Get("hypervisor_type")
		x := (v.(string))
		o.SetHypervisorType(x)
	}

	if d.HasChange("hypervisor_version") {
		v := d.Get("hypervisor_version")
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("hcl.HyperflexSoftwareCompatibilityInfo")

	if d.HasChange("server_fw_version") {
		v := d.Get("server_fw_version")
		x := (v.(string))
		o.SetServerFwVersion(x)
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

	r := conn.ApiClient.HclApi.UpdateHclHyperflexSoftwareCompatibilityInfo(conn.ctx, d.Id()).HclHyperflexSoftwareCompatibilityInfo(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating HclHyperflexSoftwareCompatibilityInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceHclHyperflexSoftwareCompatibilityInfoRead(c, d, meta)
}

func resourceHclHyperflexSoftwareCompatibilityInfoDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HclApi.DeleteHclHyperflexSoftwareCompatibilityInfo(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr.Error() != "" {
		return diag.Errorf("error occurred while deleting HclHyperflexSoftwareCompatibilityInfo object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
	}
	return de
}
