package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVnicEthIf() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVnicEthIfRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"failover_enabled": {
				Description: "Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"iscsi_ip_v4_address_allocation_type": {
				Description: "Static/Pool/DHCP Type of IP address allocated to the vNIC. It is derived from iSCSI boot policy IP Address type.\n* `None` - Type indicates that there is no IP associated to an vnic.\n* `DHCP` - The IP address is assigned using DHCP, if available.\n* `Static` - Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.\n* `Pool` - An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"iscsi_ipv4_address": {
				Description: "IP address associated to the vNIC.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mac_address": {
				Description: "The MAC address that is assigned to the vNIC based on the MAC pool that has been assigned to the LAN Connectivity Policy.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mac_address_type": {
				Description: "Type of allocation selected to assign a MAC address for the vnic.\n* `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.\n* `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.",
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
				Description: "Name of the virtual ethernet interface.",
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
			"standby_vif_id": {
				Description: "The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"static_mac_address": {
				Description: "The MAC address must be in hexadecimal format xx:xx:xx:xx:xx:xx.\nTo ensure uniqueness of MACs in the LAN fabric, you are strongly encouraged to use the\nfollowing MAC prefix 00:25:B5:xx:xx:xx.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vif_id": {
				Description: "The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceVnicEthIf().Schema},
				Computed: true,
			}},
	}
}

func dataSourceVnicEthIfRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.VnicEthIf{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("failover_enabled"); ok {
		x := (v.(bool))
		o.SetFailoverEnabled(x)
	}
	if v, ok := d.GetOk("iscsi_ip_v4_address_allocation_type"); ok {
		x := (v.(string))
		o.SetIscsiIpV4AddressAllocationType(x)
	}
	if v, ok := d.GetOk("iscsi_ipv4_address"); ok {
		x := (v.(string))
		o.SetIscsiIpv4Address(x)
	}
	if v, ok := d.GetOk("mac_address"); ok {
		x := (v.(string))
		o.SetMacAddress(x)
	}
	if v, ok := d.GetOk("mac_address_type"); ok {
		x := (v.(string))
		o.SetMacAddressType(x)
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
	if v, ok := d.GetOk("standby_vif_id"); ok {
		x := int64(v.(int))
		o.SetStandbyVifId(x)
	}
	if v, ok := d.GetOk("static_mac_address"); ok {
		x := (v.(string))
		o.SetStaticMacAddress(x)
	}
	if v, ok := d.GetOk("vif_id"); ok {
		x := int64(v.(int))
		o.SetVifId(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of VnicEthIf object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthIfList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of VnicEthIf: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.VnicEthIfList.GetCount()
	var i int32
	var vnicEthIfResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.VnicApi.GetVnicEthIfList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching VnicEthIf: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.VnicEthIfList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for VnicEthIf data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["cdn"] = flattenMapVnicCdn(s.GetCdn(), d)
				temp["class_id"] = (s.GetClassId())

				temp["eth_adapter_policy"] = flattenMapVnicEthAdapterPolicyRelationship(s.GetEthAdapterPolicy(), d)

				temp["eth_network_policy"] = flattenMapVnicEthNetworkPolicyRelationship(s.GetEthNetworkPolicy(), d)

				temp["eth_qos_policy"] = flattenMapVnicEthQosPolicyRelationship(s.GetEthQosPolicy(), d)

				temp["fabric_eth_network_control_policy"] = flattenMapFabricEthNetworkControlPolicyRelationship(s.GetFabricEthNetworkControlPolicy(), d)

				temp["fabric_eth_network_group_policy"] = flattenListFabricEthNetworkGroupPolicyRelationship(s.GetFabricEthNetworkGroupPolicy(), d)
				temp["failover_enabled"] = (s.GetFailoverEnabled())

				temp["ip_lease"] = flattenMapIppoolIpLeaseRelationship(s.GetIpLease(), d)

				temp["iscsi_boot_policy"] = flattenMapVnicIscsiBootPolicyRelationship(s.GetIscsiBootPolicy(), d)
				temp["iscsi_ip_v4_address_allocation_type"] = (s.GetIscsiIpV4AddressAllocationType())

				temp["iscsi_ip_v4_config"] = flattenMapIppoolIpV4Config(s.GetIscsiIpV4Config(), d)
				temp["iscsi_ipv4_address"] = (s.GetIscsiIpv4Address())

				temp["lan_connectivity_policy"] = flattenMapVnicLanConnectivityPolicyRelationship(s.GetLanConnectivityPolicy(), d)

				temp["lcp_vnic"] = flattenMapVnicEthIfRelationship(s.GetLcpVnic(), d)
				temp["mac_address"] = (s.GetMacAddress())
				temp["mac_address_type"] = (s.GetMacAddressType())

				temp["mac_lease"] = flattenMapMacpoolLeaseRelationship(s.GetMacLease(), d)

				temp["mac_pool"] = flattenMapMacpoolPoolRelationship(s.GetMacPool(), d)
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["order"] = (s.GetOrder())

				temp["placement"] = flattenMapVnicPlacementSettings(s.GetPlacement(), d)

				temp["profile"] = flattenMapPolicyAbstractConfigProfileRelationship(s.GetProfile(), d)

				temp["sp_vnics"] = flattenListVnicEthIfRelationship(s.GetSpVnics(), d)
				temp["standby_vif_id"] = (s.GetStandbyVifId())
				temp["static_mac_address"] = (s.GetStaticMacAddress())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["usnic_settings"] = flattenMapVnicUsnicSettings(s.GetUsnicSettings(), d)
				temp["vif_id"] = (s.GetVifId())

				temp["vmq_settings"] = flattenMapVnicVmqSettings(s.GetVmqSettings(), d)
				vnicEthIfResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(vnicEthIfResults))
	if err := d.Set("results", vnicEthIfResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(vnicEthIfResults[0]["moid"].(string))
	return de
}
