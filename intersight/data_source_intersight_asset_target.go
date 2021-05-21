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

func dataSourceAssetTarget() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAssetTargetRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"claimed_by_user_name": {
				Description: "The name or email id of the user who claimed the target.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connector_version": {
				Description: "The Device Connector version for target types which are managed by via embedded Device Connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
			"external_ip_address": {
				Description: "ExternalIpAddress is applicable for targets which are managed via an Intersight Device Connector. The value is the IP Address of the target as seen from Intersight. It is either the IP Address of the managed target's interface which has a route to the internet or a NAT IP Addresss when the target is deployed in a private network.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"management_location": {
				Description: "The location from which Intersight manages the target.\n* `Unknown` - The management mechanism is not detected. Unknown is used as a default by the implementation and is not an allowed user input.\n* `Intersight` - Management of a target is performed directly from Intersight. Network connections are established from Intersight to a management interface of the Target and authenticated using user provided credentials.\n* `IntersightAssist` - Management of a target is performed via a selected Intersight Assist. Network connections are established from the Intersight Assist to a management interface of the Target.\n* `DeviceConnector` - An Intersight Device Connector running within the Target establishes a connection to Intersight and management of the target is performed via this connection.",
				Type:        schema.TypeString,
				Optional:    true,
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
				Description: "A user provided name for the managed target.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"read_only": {
				Description: "For targets which are managed by an embedded Intersight Device Connector, this field indicates that an administrator of the target has disabled management operations of the Device Connector and only monitoring is permitted.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target.\n* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.\n* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.\n* `NotConnected` - Intersight is unable to establish a connection to the target.\n* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.\n* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.\n* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status_error_reason": {
				Description: "StatusErrorReason provides additional context for the Status.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"target_type": {
				Description: "The type of the managed target. For example a UCS Server or VMware Vcenter target.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `IMCRack` - A standalone UCS M6 and above server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.\n* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.\n* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.\n* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceAssetTarget().Schema},
				Computed: true,
			}},
	}
}

func dataSourceAssetTargetRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.AssetTarget{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("claimed_by_user_name"); ok {
		x := (v.(string))
		o.SetClaimedByUserName(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connector_version"); ok {
		x := (v.(string))
		o.SetConnectorVersion(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("external_ip_address"); ok {
		x := (v.(string))
		o.SetExternalIpAddress(x)
	}
	if v, ok := d.GetOk("management_location"); ok {
		x := (v.(string))
		o.SetManagementLocation(x)
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
	if v, ok := d.GetOk("read_only"); ok {
		x := (v.(bool))
		o.SetReadOnly(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("status_error_reason"); ok {
		x := (v.(string))
		o.SetStatusErrorReason(x)
	}
	if v, ok := d.GetOk("target_type"); ok {
		x := (v.(string))
		o.SetTargetType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of AssetTarget object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.AssetApi.GetAssetTargetList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of AssetTarget: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of AssetTarget: %s", responseErr.Error())
	}
	count := countResponse.AssetTargetList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for AssetTarget data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var assetTargetResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.AssetApi.GetAssetTargetList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching AssetTarget: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching AssetTarget: %s", responseErr.Error())
		}
		results := resMo.AssetTargetList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})

				temp["account"] = flattenMapIamAccountRelationship(s.GetAccount(), d)
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["assist"] = flattenMapAssetTargetRelationship(s.GetAssist(), d)
				temp["claimed_by_user_name"] = (s.GetClaimedByUserName())
				temp["class_id"] = (s.GetClassId())

				temp["connections"] = flattenListAssetConnection(s.GetConnections(), d)
				temp["connector_version"] = (s.GetConnectorVersion())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["external_ip_address"] = (s.GetExternalIpAddress())
				temp["ip_address"] = (s.GetIpAddress())
				temp["management_location"] = (s.GetManagementLocation())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["product_id"] = (s.GetProductId())
				temp["read_only"] = (s.GetReadOnly())

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)

				temp["services"] = flattenListAssetService(s.GetServices(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["status"] = (s.GetStatus())
				temp["status_error_reason"] = (s.GetStatusErrorReason())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_id"] = (s.GetTargetId())
				temp["target_type"] = (s.GetTargetType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				assetTargetResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(assetTargetResults))
	if err := d.Set("results", assetTargetResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(assetTargetResults[0]["moid"].(string))
	return de
}
