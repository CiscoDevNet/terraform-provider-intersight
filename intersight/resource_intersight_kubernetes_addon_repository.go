package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKubernetesAddonRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKubernetesAddonRepositoryCreate,
		ReadContext:   resourceKubernetesAddonRepositoryRead,
		UpdateContext: resourceKubernetesAddonRepositoryUpdate,
		DeleteContext: resourceKubernetesAddonRepositoryDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
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
							Computed:    true,
						},
					},
				},
			},
			"ca_cert": {
				Description: "CA certificate for the addon registry if it is not a well known CA.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Default:     "x509.Certificate",
						},
						"issuer": {
							Description: "The X.509 distinguished name of the issuer of this certificate.",
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
										Default:     "pkix.DistinguishedName",
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"locality": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "pkix.DistinguishedName",
									},
									"organization": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"organizational_unit": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"state": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
								},
							},
						},
						"not_after": {
							Description: "The date on which the certificate's validity period ends.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"not_before": {
							Description: "The date on which the certificate's validity period begins.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "x509.Certificate",
						},
						"pem_certificate": {
							Description: "The base64 encoded certificate in PEM format.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sha256_fingerprint": {
							Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"signature_algorithm": {
							Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"subject": {
							Description: "The X.509 distinguished name of the subject of this certificate.",
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
										Default:     "pkix.DistinguishedName",
									},
									"common_name": {
										Description: "A required component that identifies a person or an object.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"country": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"locality": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "pkix.DistinguishedName",
									},
									"organization": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"organizational_unit": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"state": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
								},
							},
						},
					},
				},
			},
			"catalog": {
				Description: "A reference to a workflowCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Computed:    true,
						},
					},
				},
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "kubernetes.AddonRepository",
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
			"insecure_skip_verification": {
				Description: "Allow connecting to http registries or https registries which do not have certificate signed by a well known CA.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"is_token_set": {
				Description: "Indicates whether the value of the 'token' property has been set.",
				Type:        schema.TypeBool,
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
				Description: "Name of the addon repository or registry.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "kubernetes.AddonRepository",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
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
							Computed:    true,
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
							Computed:    true,
						},
					},
				},
			},
			"registered_device": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
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
							Computed:    true,
						},
					},
				},
			},
			"repository_url": {
				Description: "URL for the repository where the addon is hosted.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"username": {
				Description: "Username to authenticate to the addon registry.",
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
										Computed:    true,
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
										Computed:    true,
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

func resourceKubernetesAddonRepositoryCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewKubernetesAddonRepositoryWithDefaults()
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(x) > 0 {
			o.SetAncestors(x)
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {
		p := make([]models.X509Certificate, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewX509CertificateWithDefaults()
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
			o.SetClassId("x509.Certificate")
			if v, ok := l["issuer"]; ok {
				{
					p := make([]models.PkixDistinguishedName, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewPkixDistinguishedNameWithDefaults()
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
						o.SetClassId("pkix.DistinguishedName")
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.SetCommonName(x)
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetCountry(x)
								}
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetLocality(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganization(x)
								}
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganizationalUnit(x)
								}
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetState(x)
								}
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetIssuer(x)
					}
				}
			}
			if v, ok := l["not_after"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetNotAfter(x)
				}
			}
			if v, ok := l["not_before"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetNotBefore(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["pem_certificate"]; ok {
				{
					x := (v.(string))
					o.SetPemCertificate(x)
				}
			}
			if v, ok := l["sha256_fingerprint"]; ok {
				{
					x := (v.(string))
					o.SetSha256Fingerprint(x)
				}
			}
			if v, ok := l["signature_algorithm"]; ok {
				{
					x := (v.(string))
					o.SetSignatureAlgorithm(x)
				}
			}
			if v, ok := l["subject"]; ok {
				{
					p := make([]models.PkixDistinguishedName, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewPkixDistinguishedNameWithDefaults()
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
						o.SetClassId("pkix.DistinguishedName")
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.SetCommonName(x)
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetCountry(x)
								}
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetLocality(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganization(x)
								}
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganizationalUnit(x)
								}
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetState(x)
								}
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetSubject(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCaCert(x)
		}
	}

	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("kubernetes.AddonRepository")

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOkExists("insecure_skip_verification"); ok {
		x := v.(bool)
		o.SetInsecureSkipVerification(x)
	}

	if v, ok := d.GetOkExists("is_password_set"); ok {
		x := v.(bool)
		o.SetIsPasswordSet(x)
	}

	if v, ok := d.GetOkExists("is_token_set"); ok {
		x := v.(bool)
		o.SetIsTokenSet(x)
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

	o.SetObjectType("kubernetes.AddonRepository")

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetOwners(x)
		}
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(x) > 0 {
			o.SetPermissionResources(x)
		}
	}

	if v, ok := d.GetOk("registered_device"); ok {
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if v, ok := d.GetOk("repository_url"); ok {
		x := (v.(string))
		o.SetRepositoryUrl(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
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

	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoVersionContextWithDefaults()
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
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
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
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["ref_mo"]; ok {
				{
					p := make([]models.MoMoRef, 0, 1)
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
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetRefMo(x)
					}
				}
			}
			if v, ok := l["timestamp"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetTimestamp(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			if v, ok := l["version_type"]; ok {
				{
					x := (v.(string))
					o.SetVersionType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	r := conn.ApiClient.KubernetesApi.CreateKubernetesAddonRepository(conn.ctx).KubernetesAddonRepository(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating KubernetesAddonRepository: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating KubernetesAddonRepository: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceKubernetesAddonRepositoryRead(c, d, meta)...)
}

func resourceKubernetesAddonRepositoryRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	r := conn.ApiClient.KubernetesApi.GetKubernetesAddonRepositoryByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "KubernetesAddonRepository object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching KubernetesAddonRepository: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching KubernetesAddonRepository: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("ca_cert", flattenMapX509Certificate(s.GetCaCert(), d)); err != nil {
		return diag.Errorf("error occurred while setting property CaCert in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("catalog", flattenMapWorkflowCatalogRelationship(s.GetCatalog(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Catalog in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("insecure_skip_verification", (s.GetInsecureSkipVerification())); err != nil {
		return diag.Errorf("error occurred while setting property InsecureSkipVerification in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("is_password_set", (s.GetIsPasswordSet())); err != nil {
		return diag.Errorf("error occurred while setting property IsPasswordSet in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("is_token_set", (s.GetIsTokenSet())); err != nil {
		return diag.Errorf("error occurred while setting property IsTokenSet in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
		return diag.Errorf("error occurred while setting property RegisteredDevice in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("repository_url", (s.GetRepositoryUrl())); err != nil {
		return diag.Errorf("error occurred while setting property RepositoryUrl in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("username", (s.GetUsername())); err != nil {
		return diag.Errorf("error occurred while setting property Username in KubernetesAddonRepository object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in KubernetesAddonRepository object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceKubernetesAddonRepositoryUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.KubernetesAddonRepository{}
	if d.HasChange("account_moid") {
		v := d.Get("account_moid")
		x := (v.(string))
		o.SetAccountMoid(x)
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

	if d.HasChange("ancestors") {
		v := d.Get("ancestors")
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(x) > 0 {
			o.SetAncestors(x)
		}
	}

	if d.HasChange("ca_cert") {
		v := d.Get("ca_cert")
		p := make([]models.X509Certificate, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.X509Certificate{}
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
			o.SetClassId("x509.Certificate")
			if v, ok := l["issuer"]; ok {
				{
					p := make([]models.PkixDistinguishedName, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewPkixDistinguishedNameWithDefaults()
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
						o.SetClassId("pkix.DistinguishedName")
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.SetCommonName(x)
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetCountry(x)
								}
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetLocality(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganization(x)
								}
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganizationalUnit(x)
								}
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetState(x)
								}
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetIssuer(x)
					}
				}
			}
			if v, ok := l["not_after"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetNotAfter(x)
				}
			}
			if v, ok := l["not_before"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetNotBefore(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["pem_certificate"]; ok {
				{
					x := (v.(string))
					o.SetPemCertificate(x)
				}
			}
			if v, ok := l["sha256_fingerprint"]; ok {
				{
					x := (v.(string))
					o.SetSha256Fingerprint(x)
				}
			}
			if v, ok := l["signature_algorithm"]; ok {
				{
					x := (v.(string))
					o.SetSignatureAlgorithm(x)
				}
			}
			if v, ok := l["subject"]; ok {
				{
					p := make([]models.PkixDistinguishedName, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewPkixDistinguishedNameWithDefaults()
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
						o.SetClassId("pkix.DistinguishedName")
						if v, ok := l["common_name"]; ok {
							{
								x := (v.(string))
								o.SetCommonName(x)
							}
						}
						if v, ok := l["country"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetCountry(x)
								}
							}
						}
						if v, ok := l["locality"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetLocality(x)
								}
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["organization"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganization(x)
								}
							}
						}
						if v, ok := l["organizational_unit"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetOrganizationalUnit(x)
								}
							}
						}
						if v, ok := l["state"]; ok {
							{
								x := make([]string, 0)
								y := reflect.ValueOf(v)
								for i := 0; i < y.Len(); i++ {
									x = append(x, y.Index(i).Interface().(string))
								}
								if len(x) > 0 {
									o.SetState(x)
								}
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetSubject(x)
					}
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCaCert(x)
		}
	}

	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.WorkflowCatalogRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsWorkflowCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("kubernetes.AddonRepository")

	if d.HasChange("create_time") {
		v := d.Get("create_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}

	if d.HasChange("domain_group_moid") {
		v := d.Get("domain_group_moid")
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if d.HasChange("insecure_skip_verification") {
		v := d.Get("insecure_skip_verification")
		x := (v.(bool))
		o.SetInsecureSkipVerification(x)
	}

	if d.HasChange("is_password_set") {
		v := d.Get("is_password_set")
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}

	if d.HasChange("is_token_set") {
		v := d.Get("is_token_set")
		x := (v.(bool))
		o.SetIsTokenSet(x)
	}

	if d.HasChange("mod_time") {
		v := d.Get("mod_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
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

	o.SetObjectType("kubernetes.AddonRepository")

	if d.HasChange("owners") {
		v := d.Get("owners")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		if len(x) > 0 {
			o.SetOwners(x)
		}
	}

	if d.HasChange("parent") {
		v := d.Get("parent")
		p := make([]models.MoBaseMoRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if d.HasChange("permission_resources") {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
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
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(x) > 0 {
			o.SetPermissionResources(x)
		}
	}

	if d.HasChange("registered_device") {
		v := d.Get("registered_device")
		p := make([]models.AssetDeviceRegistrationRelationship, 0, 1)
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
			p = append(p, models.MoMoRefAsAssetDeviceRegistrationRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetRegisteredDevice(x)
		}
	}

	if d.HasChange("repository_url") {
		v := d.Get("repository_url")
		x := (v.(string))
		o.SetRepositoryUrl(x)
	}

	if d.HasChange("shared_scope") {
		v := d.Get("shared_scope")
		x := (v.(string))
		o.SetSharedScope(x)
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

	if d.HasChange("username") {
		v := d.Get("username")
		x := (v.(string))
		o.SetUsername(x)
	}

	if d.HasChange("version_context") {
		v := d.Get("version_context")
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
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
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
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
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["ref_mo"]; ok {
				{
					p := make([]models.MoMoRef, 0, 1)
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
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetRefMo(x)
					}
				}
			}
			if v, ok := l["timestamp"]; ok {
				{
					x, _ := time.Parse(v.(string), time.RFC1123)
					o.SetTimestamp(x)
				}
			}
			if v, ok := l["nr_version"]; ok {
				{
					x := (v.(string))
					o.SetVersion(x)
				}
			}
			if v, ok := l["version_type"]; ok {
				{
					x := (v.(string))
					o.SetVersionType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	r := conn.ApiClient.KubernetesApi.UpdateKubernetesAddonRepository(conn.ctx, d.Id()).KubernetesAddonRepository(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating KubernetesAddonRepository: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating KubernetesAddonRepository: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceKubernetesAddonRepositoryRead(c, d, meta)...)
}

func resourceKubernetesAddonRepositoryDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.KubernetesApi.DeleteKubernetesAddonRepository(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting KubernetesAddonRepository object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting KubernetesAddonRepository object: %s", deleteErr.Error())
	}
	return de
}
