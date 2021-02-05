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

func dataSourceStorageNetAppIpInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageNetAppIpInterfaceRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"array_controller": {
				Description: "A reference to a storageNetAppNode resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "IP interface is enabled or not.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"home_node": {
				Description: "Name of home node of IP interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"home_port": {
				Description: "Name of home port of IP interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ip_address": {
				Description: "IP address of inteface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ip_family": {
				Description: "IP address family of inteface.\n* `IPv4` - IPv4 Address type.\n* `IPv6` - IPv6 Address type.",
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
			"name": {
				Description: "Name of IP interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"net_app_ethernet_port": {
				Description: "A reference to a storageNetAppEthernetPort resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"netmask": {
				Description: "Netmask of Interface.",
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
			"service_policy_name": {
				Description: "Services of IP interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"service_policy_uuid": {
				Description: "Services of IP interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"state": {
				Description: "State of IP interface.\n* `down` - An inactive port is listed as Down.\n* `up` - An active port is listed as Up.\n* `present` - An active port is listed as present.",
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
			"tenant": {
				Description: "A reference to a storageNetAppStorageVm resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"uuid": {
				Description: "Uuid of  NetApp IP Interface.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func dataSourceStorageNetAppIpInterfaceRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.StorageNetAppIpInterface{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("enabled"); ok {
		x := (v.(string))
		o.SetEnabled(x)
	}
	if v, ok := d.GetOk("home_node"); ok {
		x := (v.(string))
		o.SetHomeNode(x)
	}
	if v, ok := d.GetOk("home_port"); ok {
		x := (v.(string))
		o.SetHomePort(x)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		x := (v.(string))
		o.SetIpAddress(x)
	}
	if v, ok := d.GetOk("ip_family"); ok {
		x := (v.(string))
		o.SetIpFamily(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("netmask"); ok {
		x := (v.(string))
		o.SetNetmask(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("service_policy_name"); ok {
		x := (v.(string))
		o.SetServicePolicyName(x)
	}
	if v, ok := d.GetOk("service_policy_uuid"); ok {
		x := (v.(string))
		o.SetServicePolicyUuid(x)
	}
	if v, ok := d.GetOk("state"); ok {
		x := (v.(string))
		o.SetState(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of StorageNetAppIpInterface object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.StorageApi.GetStorageNetAppIpInterfaceList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching StorageNetAppIpInterface: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for StorageNetAppIpInterface list: %s", err.Error())
	}
	var s = &models.StorageNetAppIpInterfaceList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to StorageNetAppIpInterface list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for StorageNetAppIpInterface data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for StorageNetAppIpInterface data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.StorageNetAppIpInterface{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}

			if err := d.Set("array_controller", flattenMapStorageNetAppNodeRelationship(s.GetArrayController(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ArrayController: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("enabled", (s.GetEnabled())); err != nil {
				return diag.Errorf("error occurred while setting property Enabled: %s", err.Error())
			}
			if err := d.Set("home_node", (s.GetHomeNode())); err != nil {
				return diag.Errorf("error occurred while setting property HomeNode: %s", err.Error())
			}
			if err := d.Set("home_port", (s.GetHomePort())); err != nil {
				return diag.Errorf("error occurred while setting property HomePort: %s", err.Error())
			}
			if err := d.Set("ip_address", (s.GetIpAddress())); err != nil {
				return diag.Errorf("error occurred while setting property IpAddress: %s", err.Error())
			}
			if err := d.Set("ip_family", (s.GetIpFamily())); err != nil {
				return diag.Errorf("error occurred while setting property IpFamily: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return diag.Errorf("error occurred while setting property Name: %s", err.Error())
			}

			if err := d.Set("net_app_ethernet_port", flattenMapStorageNetAppEthernetPortRelationship(s.GetNetAppEthernetPort(), d)); err != nil {
				return diag.Errorf("error occurred while setting property NetAppEthernetPort: %s", err.Error())
			}
			if err := d.Set("netmask", (s.GetNetmask())); err != nil {
				return diag.Errorf("error occurred while setting property Netmask: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("service_policy_name", (s.GetServicePolicyName())); err != nil {
				return diag.Errorf("error occurred while setting property ServicePolicyName: %s", err.Error())
			}
			if err := d.Set("service_policy_uuid", (s.GetServicePolicyUuid())); err != nil {
				return diag.Errorf("error occurred while setting property ServicePolicyUuid: %s", err.Error())
			}
			if err := d.Set("services", (s.GetServices())); err != nil {
				return diag.Errorf("error occurred while setting property Services: %s", err.Error())
			}
			if err := d.Set("state", (s.GetState())); err != nil {
				return diag.Errorf("error occurred while setting property State: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("tenant", flattenMapStorageNetAppStorageVmRelationship(s.GetTenant(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tenant: %s", err.Error())
			}
			if err := d.Set("uuid", (s.GetUuid())); err != nil {
				return diag.Errorf("error occurred while setting property Uuid: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
