package intersight

import (
	"context"
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
			"number_of_schemas_in_mso": {
				Description: "Number of total schemas in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_sites_in_mso": {
				Description: "Number of sites in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_sites_serviced": {
				Description: "Number of sites serviced by ND.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"number_of_tenants_in_mso": {
				Description: "Number of total tenants in Multi-Site Orchestrator.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"type_of_site_in_mso": {
				Description: "Type of site added to Multi-Site Orchestrator.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
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
					"number_of_schemas_in_mso": {
						Description: "Number of total schemas in Multi-Site Orchestrator.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_sites_in_mso": {
						Description: "Number of sites in Multi-Site Orchestrator.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_sites_serviced": {
						Description: "Number of sites serviced by ND.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"number_of_tenants_in_mso": {
						Description: "Number of total tenants in Multi-Site Orchestrator.",
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
					"type_of_site_in_mso": {
						Description: "Type of site added to Multi-Site Orchestrator.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
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
	if v, ok := d.GetOk("number_of_schemas_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfSchemasInMso(x)
	}
	if v, ok := d.GetOk("number_of_sites_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfSitesInMso(x)
	}
	if v, ok := d.GetOk("number_of_sites_serviced"); ok {
		x := int64(v.(int))
		o.SetNumberOfSitesServiced(x)
	}
	if v, ok := d.GetOk("number_of_tenants_in_mso"); ok {
		x := int64(v.(int))
		o.SetNumberOfTenantsInMso(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("type_of_site_in_mso"); ok {
		x := (v.(string))
		o.SetTypeOfSiteInMso(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of NiatelemetryNexusDashboards object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNexusDashboardsList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of NiatelemetryNexusDashboards: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.NiatelemetryNexusDashboardsList.GetCount()
	var i int32
	var niatelemetryNexusDashboardsResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.NiatelemetryApi.GetNiatelemetryNexusDashboardsList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching NiatelemetryNexusDashboards: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.NiatelemetryNexusDashboardsList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for NiatelemetryNexusDashboards data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["cluster_name"] = (s.GetClusterName())
				temp["is_cluster_healthy"] = (s.GetIsClusterHealthy())
				temp["moid"] = (s.GetMoid())
				temp["nd_cluster_size"] = (s.GetNdClusterSize())
				temp["nd_type"] = (s.GetNdType())
				temp["nd_version"] = (s.GetNdVersion())
				temp["number_of_apps"] = (s.GetNumberOfApps())
				temp["number_of_schemas_in_mso"] = (s.GetNumberOfSchemasInMso())
				temp["number_of_sites_in_mso"] = (s.GetNumberOfSitesInMso())
				temp["number_of_sites_serviced"] = (s.GetNumberOfSitesServiced())
				temp["number_of_tenants_in_mso"] = (s.GetNumberOfTenantsInMso())
				temp["object_type"] = (s.GetObjectType())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type_of_site_in_mso"] = (s.GetTypeOfSiteInMso())
				niatelemetryNexusDashboardsResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(niatelemetryNexusDashboardsResults))
	if err := d.Set("results", niatelemetryNexusDashboardsResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(niatelemetryNexusDashboardsResults[0]["moid"].(string))
	return de
}
