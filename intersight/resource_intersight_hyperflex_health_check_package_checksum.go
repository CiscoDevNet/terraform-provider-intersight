package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHyperflexHealthCheckPackageChecksum() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHyperflexHealthCheckPackageChecksumCreate,
		ReadContext:   resourceHyperflexHealthCheckPackageChecksumRead,
		UpdateContext: resourceHyperflexHealthCheckPackageChecksumUpdate,
		DeleteContext: resourceHyperflexHealthCheckPackageChecksumDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CustomizeTagDiff,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ancestors": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"checksum": {
				Description: "SHA512 checksum of the health check package.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "hyperflex.HealthCheckPackageChecksum",
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
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the health check Debian package.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "hyperflex.HealthCheckPackageChecksum",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"package_name": {
				Description: "HyperFlex health check Debian package file name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.MoRef",
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
						},
					},
				},
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
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
			"timestamp": {
				Description: "Timestamp of last update of HyperFlex health check package checksum.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nr_version": {
				Description: "HyperFlex health check Debian Package Version.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
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
										Default:     "mo.MoRef",
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
									},
								},
							},
						},
						"timestamp": {
							Description: "The time this versioned Managed Object was created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceHyperflexHealthCheckPackageChecksumCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewHyperflexHealthCheckPackageChecksumWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("checksum"); ok {
		x := (v.(string))
		o.SetChecksum(x)
	}

	o.SetClassId("hyperflex.HealthCheckPackageChecksum")

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("hyperflex.HealthCheckPackageChecksum")

	if v, ok := d.GetOk("package_name"); ok {
		x := (v.(string))
		o.SetPackageName(x)
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

	if v, ok := d.GetOk("timestamp"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetTimestamp(x)
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.HyperflexApi.CreateHyperflexHealthCheckPackageChecksum(conn.ctx).HyperflexHealthCheckPackageChecksum(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating HyperflexHealthCheckPackageChecksum: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating HyperflexHealthCheckPackageChecksum: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceHyperflexHealthCheckPackageChecksumRead(c, d, meta)...)
}

func resourceHyperflexHealthCheckPackageChecksumRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	r := conn.ApiClient.HyperflexApi.GetHyperflexHealthCheckPackageChecksumByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexHealthCheckPackageChecksum object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexHealthCheckPackageChecksum: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching HyperflexHealthCheckPackageChecksum: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("checksum", (s.GetChecksum())); err != nil {
		return diag.Errorf("error occurred while setting property Checksum in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("package_name", (s.GetPackageName())); err != nil {
		return diag.Errorf("error occurred while setting property PackageName in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("timestamp", (s.GetTimestamp()).String()); err != nil {
		return diag.Errorf("error occurred while setting property Timestamp in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("nr_version", (s.GetVersion())); err != nil {
		return diag.Errorf("error occurred while setting property Version in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in HyperflexHealthCheckPackageChecksum object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceHyperflexHealthCheckPackageChecksumUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexHealthCheckPackageChecksum{}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("checksum") {
		v := d.Get("checksum")
		x := (v.(string))
		o.SetChecksum(x)
	}

	o.SetClassId("hyperflex.HealthCheckPackageChecksum")

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

	o.SetObjectType("hyperflex.HealthCheckPackageChecksum")

	if d.HasChange("package_name") {
		v := d.Get("package_name")
		x := (v.(string))
		o.SetPackageName(x)
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
		o.SetTags(x)
	}

	if d.HasChange("timestamp") {
		v := d.Get("timestamp")
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetTimestamp(x)
	}

	if d.HasChange("nr_version") {
		v := d.Get("nr_version")
		x := (v.(string))
		o.SetVersion(x)
	}

	r := conn.ApiClient.HyperflexApi.UpdateHyperflexHealthCheckPackageChecksum(conn.ctx, d.Id()).HyperflexHealthCheckPackageChecksum(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating HyperflexHealthCheckPackageChecksum: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating HyperflexHealthCheckPackageChecksum: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceHyperflexHealthCheckPackageChecksumRead(c, d, meta)...)
}

func resourceHyperflexHealthCheckPackageChecksumDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.HyperflexApi.DeleteHyperflexHealthCheckPackageChecksum(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "HyperflexHealthCheckPackageChecksumDelete: HyperflexHealthCheckPackageChecksum object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting HyperflexHealthCheckPackageChecksum object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting HyperflexHealthCheckPackageChecksum object: %s", deleteErr.Error())
	}
	return de
}
