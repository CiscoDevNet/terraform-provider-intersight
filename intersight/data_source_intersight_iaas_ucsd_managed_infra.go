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

func dataSourceIaasUcsdManagedInfra() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIaasUcsdManagedInfraRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"advanced_catalog_count": {
				Description: "Total advanced catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"bm_catalog_count": {
				Description: "Total bare metal catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"container_catalog_count": {
				Description: "Total service container catalogs in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"esxi_host_count": {
				Description: "Total ESXi hosts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"external_group_count": {
				Description: "Total external (Ldap) groups in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"guid": {
				Description: "A reference to a iaasUcsdInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"hyperv_host_count": {
				Description: "Total HyperV hosts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"local_group_count": {
				Description: "Total local groups in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"standard_catalog_count": {
				Description: "Total standard catalogs in UCSD.",
				Type:        schema.TypeInt,
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
			"user_count": {
				Description: "Total user accounts in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"vdc_count": {
				Description: "Total virtual datacenters in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"vm_count": {
				Description: "Total Virtual machines in UCSD.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceIaasUcsdManagedInfraRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.IaasUcsdManagedInfra{}
	if v, ok := d.GetOk("advanced_catalog_count"); ok {
		x := int64(v.(int))
		o.SetAdvancedCatalogCount(x)
	}
	if v, ok := d.GetOk("bm_catalog_count"); ok {
		x := int64(v.(int))
		o.SetBmCatalogCount(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("container_catalog_count"); ok {
		x := int64(v.(int))
		o.SetContainerCatalogCount(x)
	}
	if v, ok := d.GetOk("esxi_host_count"); ok {
		x := int64(v.(int))
		o.SetEsxiHostCount(x)
	}
	if v, ok := d.GetOk("external_group_count"); ok {
		x := int64(v.(int))
		o.SetExternalGroupCount(x)
	}
	if v, ok := d.GetOk("hyperv_host_count"); ok {
		x := int64(v.(int))
		o.SetHypervHostCount(x)
	}
	if v, ok := d.GetOk("local_group_count"); ok {
		x := int64(v.(int))
		o.SetLocalGroupCount(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("standard_catalog_count"); ok {
		x := int64(v.(int))
		o.SetStandardCatalogCount(x)
	}
	if v, ok := d.GetOk("user_count"); ok {
		x := int64(v.(int))
		o.SetUserCount(x)
	}
	if v, ok := d.GetOk("vdc_count"); ok {
		x := int64(v.(int))
		o.SetVdcCount(x)
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.SetVmCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of IaasUcsdManagedInfra object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.IaasApi.GetIaasUcsdManagedInfraList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching IaasUcsdManagedInfra: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for IaasUcsdManagedInfra list: %s", err.Error())
	}
	var s = &models.IaasUcsdManagedInfraList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to IaasUcsdManagedInfra list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for IaasUcsdManagedInfra data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for IaasUcsdManagedInfra data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.IaasUcsdManagedInfra{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("advanced_catalog_count", (s.GetAdvancedCatalogCount())); err != nil {
				return diag.Errorf("error occurred while setting property AdvancedCatalogCount: %s", err.Error())
			}
			if err := d.Set("bm_catalog_count", (s.GetBmCatalogCount())); err != nil {
				return diag.Errorf("error occurred while setting property BmCatalogCount: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("container_catalog_count", (s.GetContainerCatalogCount())); err != nil {
				return diag.Errorf("error occurred while setting property ContainerCatalogCount: %s", err.Error())
			}
			if err := d.Set("esxi_host_count", (s.GetEsxiHostCount())); err != nil {
				return diag.Errorf("error occurred while setting property EsxiHostCount: %s", err.Error())
			}
			if err := d.Set("external_group_count", (s.GetExternalGroupCount())); err != nil {
				return diag.Errorf("error occurred while setting property ExternalGroupCount: %s", err.Error())
			}

			if err := d.Set("guid", flattenMapIaasUcsdInfoRelationship(s.GetGuid(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Guid: %s", err.Error())
			}
			if err := d.Set("hyperv_host_count", (s.GetHypervHostCount())); err != nil {
				return diag.Errorf("error occurred while setting property HypervHostCount: %s", err.Error())
			}
			if err := d.Set("local_group_count", (s.GetLocalGroupCount())); err != nil {
				return diag.Errorf("error occurred while setting property LocalGroupCount: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("standard_catalog_count", (s.GetStandardCatalogCount())); err != nil {
				return diag.Errorf("error occurred while setting property StandardCatalogCount: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("user_count", (s.GetUserCount())); err != nil {
				return diag.Errorf("error occurred while setting property UserCount: %s", err.Error())
			}
			if err := d.Set("vdc_count", (s.GetVdcCount())); err != nil {
				return diag.Errorf("error occurred while setting property VdcCount: %s", err.Error())
			}
			if err := d.Set("vm_count", (s.GetVmCount())); err != nil {
				return diag.Errorf("error occurred while setting property VmCount: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
