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

func resourceLicenseLicenseReservationOp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLicenseLicenseReservationOpCreate,
		ReadContext:   resourceLicenseLicenseReservationOpRead,
		UpdateContext: resourceLicenseLicenseReservationOpUpdate,
		DeleteContext: resourceLicenseLicenseReservationOpDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"account": {
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
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"auth_code": {
				Description: "Revervation code used to install the license.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auth_code_installed": {
				Description: "Flag to indicate whether authorization code is installed.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"confirm_code": {
				Description: "Confirm code used to complete the license update on smart license account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"generate_request_code": {
				Description: "Trigger the generation of request code for specific license reservation.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"generate_return_code": {
				Description: "Trigger the generation of return code for specific license reservation.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"request_code": {
				Description: "Revervation code used to generate authorization code from CSSM.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"return_code": {
				Description: "Return code used to return the reserved license to smart license account.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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

func resourceLicenseLicenseReservationOpCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewLicenseLicenseReservationOpWithDefaults()
	if v, ok := d.GetOk("account"); ok {
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
			o.SetAccount(x)
		}
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("auth_code"); ok {
		x := (v.(string))
		o.SetAuthCode(x)
	}

	if v, ok := d.GetOkExists("auth_code_installed"); ok {
		x := v.(bool)
		o.SetAuthCodeInstalled(x)
	}

	o.SetClassId("license.LicenseReservationOp")

	if v, ok := d.GetOk("confirm_code"); ok {
		x := (v.(string))
		o.SetConfirmCode(x)
	}

	if v, ok := d.GetOkExists("generate_request_code"); ok {
		x := v.(bool)
		o.SetGenerateRequestCode(x)
	}

	if v, ok := d.GetOkExists("generate_return_code"); ok {
		x := v.(bool)
		o.SetGenerateReturnCode(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("license.LicenseReservationOp")

	if v, ok := d.GetOk("request_code"); ok {
		x := (v.(string))
		o.SetRequestCode(x)
	}

	if v, ok := d.GetOk("return_code"); ok {
		x := (v.(string))
		o.SetReturnCode(x)
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

	r := conn.ApiClient.LicenseApi.CreateLicenseLicenseReservationOp(conn.ctx).LicenseLicenseReservationOp(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("failed while creating LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return resourceLicenseLicenseReservationOpRead(c, d, meta)
}

func resourceLicenseLicenseReservationOpRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.LicenseApi.GetLicenseLicenseReservationOpByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "LicenseLicenseReservationOp object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		return diag.Errorf("error occurred while fetching LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Account in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("auth_code", (s.GetAuthCode())); err != nil {
		return diag.Errorf("error occurred while setting property AuthCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("auth_code_installed", (s.GetAuthCodeInstalled())); err != nil {
		return diag.Errorf("error occurred while setting property AuthCodeInstalled in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("confirm_code", (s.GetConfirmCode())); err != nil {
		return diag.Errorf("error occurred while setting property ConfirmCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("generate_request_code", (s.GetGenerateRequestCode())); err != nil {
		return diag.Errorf("error occurred while setting property GenerateRequestCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("generate_return_code", (s.GetGenerateReturnCode())); err != nil {
		return diag.Errorf("error occurred while setting property GenerateReturnCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("request_code", (s.GetRequestCode())); err != nil {
		return diag.Errorf("error occurred while setting property RequestCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("return_code", (s.GetReturnCode())); err != nil {
		return diag.Errorf("error occurred while setting property ReturnCode in LicenseLicenseReservationOp object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in LicenseLicenseReservationOp object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceLicenseLicenseReservationOpUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.LicenseLicenseReservationOp{}
	if d.HasChange("account") {
		v := d.Get("account")
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
			o.SetAccount(x)
		}
	}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("auth_code") {
		v := d.Get("auth_code")
		x := (v.(string))
		o.SetAuthCode(x)
	}

	if d.HasChange("auth_code_installed") {
		v := d.Get("auth_code_installed")
		x := (v.(bool))
		o.SetAuthCodeInstalled(x)
	}

	o.SetClassId("license.LicenseReservationOp")

	if d.HasChange("confirm_code") {
		v := d.Get("confirm_code")
		x := (v.(string))
		o.SetConfirmCode(x)
	}

	if d.HasChange("generate_request_code") {
		v := d.Get("generate_request_code")
		x := (v.(bool))
		o.SetGenerateRequestCode(x)
	}

	if d.HasChange("generate_return_code") {
		v := d.Get("generate_return_code")
		x := (v.(bool))
		o.SetGenerateReturnCode(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	o.SetObjectType("license.LicenseReservationOp")

	if d.HasChange("request_code") {
		v := d.Get("request_code")
		x := (v.(string))
		o.SetRequestCode(x)
	}

	if d.HasChange("return_code") {
		v := d.Get("return_code")
		x := (v.(string))
		o.SetReturnCode(x)
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

	r := conn.ApiClient.LicenseApi.UpdateLicenseLicenseReservationOp(conn.ctx, d.Id()).LicenseLicenseReservationOp(*o)
	result, _, responseErr := r.Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while updating LicenseLicenseReservationOp: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return resourceLicenseLicenseReservationOpRead(c, d, meta)
}

func resourceLicenseLicenseReservationOpDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	var warning = diag.Diagnostic{Severity: diag.Warning, Summary: "LicenseLicenseReservationOp does not allow delete functionality"}
	de = append(de, warning)
	return de
}
