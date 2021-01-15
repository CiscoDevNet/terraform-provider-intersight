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

func dataSourceNiatelemetryNexusDashboards() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNiatelemetryNexusDashboardsRead,
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
			},
			"cluster_name": {
				Description: "Nexus Dashboard can onboard multiple APIC clusters/sites.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_cluster_healthy": {
				Description: "Health of Nexus Dashboard cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nd_cluster_size": {
				Description: "Number of nodes in Nexus Dashboard cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"nd_type": {
				Description: "Node type in Nexus Dashboard cluster.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"nd_version": {
				Description: "Version running on Nexus Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"number_of_apps": {
				Description: "Number of applications installed in the Nexus Dashboard.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_sites_serviced": {
				Description: "Number of sites serviced by ND.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
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

func dataSourceNiatelemetryNexusDashboardsRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.NiatelemetryNexusDashboards{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_name"); ok {
		x := (v.(string))
		o.SetClusterName(x)
	}
	if v, ok := d.GetOk("is_cluster_healthy"); ok {
		x := (v.(string))
		o.SetIsClusterHealthy(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("nd_cluster_size"); ok {
		x := int64(v.(int))
		o.SetNdClusterSize(x)
	}
	if v, ok := d.GetOk("nd_type"); ok {
		x := (v.(string))
		o.SetNdType(x)
	}
	if v, ok := d.GetOk("nd_version"); ok {
		x := (v.(string))
		o.SetNdVersion(x)
	}
	if v, ok := d.GetOk("number_of_apps"); ok {
		x := int64(v.(int))
		o.SetNumberOfApps(x)
	}
	if v, ok := d.GetOk("number_of_sites_serviced"); ok {
		x := int64(v.(int))
		o.SetNumberOfSitesServiced(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNexusDashboards object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNexusDashboardsList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching NiatelemetryNexusDashboards: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for NiatelemetryNexusDashboards list: %s", err.Error())
	}
	var s = &models.NiatelemetryNexusDashboardsList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to NiatelemetryNexusDashboards list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for NiatelemetryNexusDashboards data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for NiatelemetryNexusDashboards data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.NiatelemetryNexusDashboards{}
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
			if err := d.Set("cluster_name", (s.GetClusterName())); err != nil {
				return diag.Errorf("error occurred while setting property ClusterName: %s", err.Error())
			}
			if err := d.Set("is_cluster_healthy", (s.GetIsClusterHealthy())); err != nil {
				return diag.Errorf("error occurred while setting property IsClusterHealthy: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("nd_cluster_size", (s.GetNdClusterSize())); err != nil {
				return diag.Errorf("error occurred while setting property NdClusterSize: %s", err.Error())
			}
			if err := d.Set("nd_type", (s.GetNdType())); err != nil {
				return diag.Errorf("error occurred while setting property NdType: %s", err.Error())
			}
			if err := d.Set("nd_version", (s.GetNdVersion())); err != nil {
				return diag.Errorf("error occurred while setting property NdVersion: %s", err.Error())
			}
			if err := d.Set("number_of_apps", (s.GetNumberOfApps())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfApps: %s", err.Error())
			}
			if err := d.Set("number_of_sites_serviced", (s.GetNumberOfSitesServiced())); err != nil {
				return diag.Errorf("error occurred while setting property NumberOfSitesServiced: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)); err != nil {
				return diag.Errorf("error occurred while setting property RegisteredDevice: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
