package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexNode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexNodeRead,
		Schema: map[string]*schema.Schema{
			"build_number": {
				Description: "The build number of the hypervisor running on the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"display_version": {
				Description: "The user-friendly string representation of the hypervisor version of the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"host_name": {
				Description: "The hostname configured for the hypervisor running on the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor": {
				Description: "The type of hypervisor running on the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"lockdown": {
				Description: "The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"model_number": {
				Description: "The model of the host server.",
				Type:        schema.TypeString,
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
			"role": {
				Description: "The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage.\n* `UNKNOWN` - The role of the HyperFlex cluster node is not known.\n* `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster.\n* `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial_number": {
				Description: "The serial of the host server.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "The status of the host. Indicates whether the hypervisor is online.\n* `UNKNOWN` - The host status cannot be determined.\n* `ONLINE` - The host is online and operational.\n* `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster.\n* `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade.\n* `DEGRADED` - The host is degraded and may not be performing in its full operational capacity.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"nr_version": {
				Description: "The version of the hypervisor running on the host.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"build_number": {
						Description: "The build number of the hypervisor running on the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster": {
						Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"cluster_member": {
						Description: "A reference to a assetClusterMember resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"display_version": {
						Description: "The user-friendly string representation of the hypervisor version of the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"host_name": {
						Description: "The hostname configured for the hypervisor running on the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"hypervisor": {
						Description: "The type of hypervisor running on the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"identity": {
						Description: "The unique identifier of the host.",
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
								"links": {
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
											},
											"comments": {
												Description: "Comment for this HyperFlex resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"href": {
												Description: "URI of resource. Target URL for making REST call.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"method": {
												Description: "HTTP verb that this HyperFlex link DT is referring to.\n* `POST` - HTTP verb POST for this task definition.\n* `GET` - HTTP verb GET for this task definition.\n* `PUT` - HTTP verb PUT for this task definition.\n* `DELETE` - HTTP verb DELETE for this task definition.",
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
											"rel": {
												Description: "Relationship of link to this resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
									Computed: true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"uuid": {
									Description: "The unique identifier string of an entity.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"ip": {
						Description: "The hypervisor management address. This can be a FQDN or IP address.",
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
								"address": {
									Description: "The network address as an FQDN or IPv4 address.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"fqdn": {
									Description: "The fully qualified domain name for the network address.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"ip": {
									Description: "The network address as an IPv4 address.",
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
							},
						},
					},
					"lockdown": {
						Description: "The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled.",
						Type:        schema.TypeBool,
						Optional:    true,
						Computed:    true,
					},
					"model_number": {
						Description: "The model of the host server.",
						Type:        schema.TypeString,
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
					"physical_server": {
						Description: "A reference to a computePhysical resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"role": {
						Description: "The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage.\n* `UNKNOWN` - The role of the HyperFlex cluster node is not known.\n* `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster.\n* `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"serial_number": {
						Description: "The serial of the host server.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"status": {
						Description: "The status of the host. Indicates whether the hypervisor is online.\n* `UNKNOWN` - The host status cannot be determined.\n* `ONLINE` - The host is online and operational.\n* `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster.\n* `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade.\n* `DEGRADED` - The host is degraded and may not be performing in its full operational capacity.",
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
					"nr_version": {
						Description: "The version of the hypervisor running on the host.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexNodeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexNode{}
	if v, ok := d.GetOk("build_number"); ok {
		x := (v.(string))
		o.SetBuildNumber(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("display_version"); ok {
		x := (v.(string))
		o.SetDisplayVersion(x)
	}
	if v, ok := d.GetOk("host_name"); ok {
		x := (v.(string))
		o.SetHostName(x)
	}
	if v, ok := d.GetOk("hypervisor"); ok {
		x := (v.(string))
		o.SetHypervisor(x)
	}
	if v, ok := d.GetOk("lockdown"); ok {
		x := (v.(bool))
		o.SetLockdown(x)
	}
	if v, ok := d.GetOk("model_number"); ok {
		x := (v.(string))
		o.SetModelNumber(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("role"); ok {
		x := (v.(string))
		o.SetRole(x)
	}
	if v, ok := d.GetOk("serial_number"); ok {
		x := (v.(string))
		o.SetSerialNumber(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexNode object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexNodeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexNode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexNodeList.GetCount()
	var i int32
	var hyperflexNodeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexNodeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexNode: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexNodeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexNode data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["build_number"] = (s.GetBuildNumber())
				temp["class_id"] = (s.GetClassId())

				temp["cluster"] = flattenMapHyperflexClusterRelationship(s.GetCluster(), d)

				temp["cluster_member"] = flattenMapAssetClusterMemberRelationship(s.GetClusterMember(), d)
				temp["display_version"] = (s.GetDisplayVersion())
				temp["host_name"] = (s.GetHostName())
				temp["hypervisor"] = (s.GetHypervisor())

				temp["identity"] = flattenMapHyperflexHxUuIdDt(s.GetIdentity(), d)

				temp["ip"] = flattenMapHyperflexHxNetworkAddressDt(s.GetIp(), d)
				temp["lockdown"] = (s.GetLockdown())
				temp["model_number"] = (s.GetModelNumber())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())

				temp["physical_server"] = flattenMapComputePhysicalRelationship(s.GetPhysicalServer(), d)
				temp["role"] = (s.GetRole())
				temp["serial_number"] = (s.GetSerialNumber())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["nr_version"] = (s.GetVersion())
				hyperflexNodeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexNodeResults))
	if err := d.Set("results", hyperflexNodeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexNodeResults[0]["moid"].(string))
	return de
}
