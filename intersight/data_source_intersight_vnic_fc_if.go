package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicFcIf() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicFcIfRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
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
			},
			"name": {
				Description: "Name of the virtual fibre channel interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"order": {
				Description: "The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"persistent_bindings": {
				Description: "Enables retention of LUN ID associations in memory until they are manually cleared.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"static_wwpn_address": {
				Description: "The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.\nAllowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF.\nTo ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters.\n* `fc-initiator` - The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems.\n* `fc-nvme-initiator` - Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems.\n* `fc-nvme-target` - Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver.\n* `fc-target` - Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vif_id": {
				Description: "This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"wwpn": {
				Description: "The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"wwpn_address_type": {
				Description: "Type of allocation selected to assign a WWPN address to the vhba.\n* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.\n* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicFcIf().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicFcIfRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicFcIf{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("order"); ok {
		x := int64(v.(int))
		o.SetOrder(x)
	}
	if v, ok := d.GetOk("persistent_bindings"); ok {
		x := (v.(bool))
		o.SetPersistentBindings(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("static_wwpn_address"); ok {
		x := (v.(string))
		o.SetStaticWwpnAddress(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("vif_id"); ok {
		x := int64(v.(int))
		o.SetVifId(x)
	}
	if v, ok := d.GetOk("wwpn"); ok {
		x := (v.(string))
		o.SetWwpn(x)
	}
	if v, ok := d.GetOk("wwpn_address_type"); ok {
		x := (v.(string))
		o.SetWwpnAddressType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicFcIf object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcIfList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of VnicFcIf: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of VnicFcIf: %s", responseErr.Error())
	}
	count := countResponse.VnicFcIfList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for VnicFcIf data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var vnicFcIfResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicFcIfList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching VnicFcIf: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching VnicFcIf: %s", responseErr.Error())
		}
		results := resMo.VnicFcIfList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["fc_adapter_policy"] = flattenMapVnicFcAdapterPolicyRelationship(s.GetFcAdapterPolicy(), d)

				temp["fc_network_policy"] = flattenMapVnicFcNetworkPolicyRelationship(s.GetFcNetworkPolicy(), d)

				temp["fc_qos_policy"] = flattenMapVnicFcQosPolicyRelationship(s.GetFcQosPolicy(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["order"] = (s.GetOrder())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["persistent_bindings"] = (s.GetPersistentBindings())

				temp["placement"] = flattenMapVnicPlacementSettings(s.GetPlacement(), d)

				temp["profile"] = flattenMapPolicyAbstractConfigProfileRelationship(s.GetProfile(), d)

				temp["san_connectivity_policy"] = flattenMapVnicSanConnectivityPolicyRelationship(s.GetSanConnectivityPolicy(), d)

				temp["scp_vhba"] = flattenMapVnicFcIfRelationship(s.GetScpVhba(), d)
				temp["shared_scope"] = (s.GetSharedScope())

				temp["sp_vhbas"] = flattenListVnicFcIfRelationship(s.GetSpVhbas(), d)
				temp["static_wwpn_address"] = (s.GetStaticWwpnAddress())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["type"] = (s.GetType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vif_id"] = (s.GetVifId())
				temp["wwpn"] = (s.GetWwpn())
				temp["wwpn_address_type"] = (s.GetWwpnAddressType())

				temp["wwpn_lease"] = flattenMapFcpoolLeaseRelationship(s.GetWwpnLease(), d)

				temp["wwpn_pool"] = flattenMapFcpoolPoolRelationship(s.GetWwpnPool(), d)
				vnicFcIfResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicFcIfResults))
	if err := d.Set("results", vnicFcIfResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicFcIfResults[0]["moid"].(string))
	return de
}
