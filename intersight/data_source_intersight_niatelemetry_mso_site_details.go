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

func dataSourceNiatelemetryMsoSiteDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryMsoSiteDetailsRead,
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
			"is_cloud_sec_enabled": {
				Description: "Status of cloudSec on Multi-Site Orchestrator site.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"number_of_leafs_per_site_in_mso": {
				Description: "Number of leafs per site in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_pods_per_site_in_mso": {
				Description: "Number of pods per site in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_spines_per_site_in_mso": {
				Description: "Number of spines per site in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
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
			"site_id": {
				Description: "ID of site in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_name": {
				Description: "Name of the site in Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_version": {
				Description: "Version of the controller in the site.",
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

func dataSourceNiatelemetryMsoSiteDetailsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryMsoSiteDetails{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("is_cloud_sec_enabled"); ok {
		x := (v.(string))
		o.SetIsCloudSecEnabled(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("number_of_leafs_per_site_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfLeafsPerSiteInMso(x)
	}
	if v, ok := d.GetOk("number_of_pods_per_site_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfPodsPerSiteInMso(x)
	}
	if v, ok := d.GetOk("number_of_spines_per_site_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfSpinesPerSiteInMso(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("site_id"); ok {
		x := (v.(string))
		o.SetSiteId(x)
	}
	if v, ok := d.GetOk("site_name"); ok {
		x := (v.(string))
		o.SetSiteName(x)
	}
	if v, ok := d.GetOk("site_version"); ok {
		x := (v.(string))
		o.SetSiteVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryMsoSiteDetails object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryMsoSiteDetailsList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching NiatelemetryMsoSiteDetails: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NiatelemetryMsoSiteDetails list: %s", err.Error())
	}
	var s = &models.NiatelemetryMsoSiteDetailsList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NiatelemetryMsoSiteDetails list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NiatelemetryMsoSiteDetails data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NiatelemetryMsoSiteDetails data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiatelemetryMsoSiteDetails{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("is_cloud_sec_enabled", (s.GetIsCloudSecEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property IsCloudSecEnabled: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("number_of_leafs_per_site_in_mso", (s.GetNumberOfLeafsPerSiteInMso())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfLeafsPerSiteInMso: %s", err.Error())
			}
			if err := d.Set("number_of_pods_per_site_in_mso", (s.GetNumberOfPodsPerSiteInMso())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfPodsPerSiteInMso: %s", err.Error())
			}
			if err := d.Set("number_of_spines_per_site_in_mso", (s.GetNumberOfSpinesPerSiteInMso())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfSpinesPerSiteInMso: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}
			if err := d.Set("site_id", (s.GetSiteId())); err != nil {
				return diag.Errorf("error occurred while setting property SiteId: %s", err.Error())
			}
			if err := d.Set("site_name", (s.GetSiteName())); err != nil {
				return diag.Errorf("error occurred while setting property SiteName: %s", err.Error())
			}
			if err := d.Set("site_version", (s.GetSiteVersion())); err != nil {
				return diag.Errorf("error occurred while setting property SiteVersion: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
