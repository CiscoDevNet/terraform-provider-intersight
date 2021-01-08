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

func dataSourceHyperflexNodeProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexNodeProfileRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"assigned_server": {
				Description: "A reference to a computeRackUnit resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_profile": {
				Description: "A reference to a hyperflexClusterProfile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"description": {
				Description: "Description of the profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hxdp_data_ip": {
				Description: "IP address for storage data network (Controller VM interface).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hxdp_mgmt_ip": {
				Description: "IP address for HyperFlex management network.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_control_ip": {
				Description: "IP address for hypervisor control such as VM migration or pod management.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_data_ip": {
				Description: "IP address for storage data network (Hypervisor interface).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_mgmt_ip": {
				Description: "IP address for Hypervisor management network.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete profile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"src_template": {
				Description: "A reference to a policyAbstractProfile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
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
			"type": {
				Description: "Defines the type of the profile. Accepted value is instance.\n* `instance` - The profile defines the configuration for a specific instance of a target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceHyperflexNodeProfileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexNodeProfile{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("hxdp_data_ip"); ok {
		x := (v.(string))
		o.SetHxdpDataIp(x)
	}
	if v, ok := d.GetOk("hxdp_mgmt_ip"); ok {
		x := (v.(string))
		o.SetHxdpMgmtIp(x)
	}
	if v, ok := d.GetOk("hypervisor_control_ip"); ok {
		x := (v.(string))
		o.SetHypervisorControlIp(x)
	}
	if v, ok := d.GetOk("hypervisor_data_ip"); ok {
		x := (v.(string))
		o.SetHypervisorDataIp(x)
	}
	if v, ok := d.GetOk("hypervisor_mgmt_ip"); ok {
		x := (v.(string))
		o.SetHypervisorMgmtIp(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexNodeProfile object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexNodeProfileList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexNodeProfile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexNodeProfile list: %s", err.Error())
	}
	var s = &models.HyperflexNodeProfileList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexNodeProfile list: %s", err.Error())
	}
	result := s.GetResults()
	if result == nil {
		return diag.Errorf("your query for HyperflexNodeProfile did not return results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexNodeProfile{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("assigned_server", flattenMapComputeRackUnitRelationship(s.GetAssignedServer(), d)); err != nil {
				return diag.Errorf("error occurred while setting property AssignedServer: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster_profile", flattenMapHyperflexClusterProfileRelationship(s.GetClusterProfile(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ClusterProfile: %s", err.Error())
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return diag.Errorf("error occurred while setting property Description: %s", err.Error())
			}
			if err := d.Set("hxdp_data_ip", (s.GetHxdpDataIp())); err != nil {
				return diag.Errorf("error occurred while setting property HxdpDataIp: %s", err.Error())
			}
			if err := d.Set("hxdp_mgmt_ip", (s.GetHxdpMgmtIp())); err != nil {
				return diag.Errorf("error occurred while setting property HxdpMgmtIp: %s", err.Error())
			}
			if err := d.Set("hypervisor_control_ip", (s.GetHypervisorControlIp())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorControlIp: %s", err.Error())
			}
			if err := d.Set("hypervisor_data_ip", (s.GetHypervisorDataIp())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorDataIp: %s", err.Error())
			}
			if err := d.Set("hypervisor_mgmt_ip", (s.GetHypervisorMgmtIp())); err != nil {
				return diag.Errorf("error occurred while setting property HypervisorMgmtIp: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}

			if err := d.Set("src_template", flattenMapPolicyAbstractProfileRelationship(s.GetSrcTemplate(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SrcTemplate: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("type", (s.GetType())); err != nil {
				return diag.Errorf("error occurred while setting property Type: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
