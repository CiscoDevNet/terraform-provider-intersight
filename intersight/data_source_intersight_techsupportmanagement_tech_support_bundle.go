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

func dataSourceTechsupportmanagementTechSupportBundle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTechsupportmanagementTechSupportBundleRead,
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
			"device_identifier": {
				Description: "The device identifier used to uniquely identify an individual device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_type": {
				Description: "The device type obtained from the inventory.",
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
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pid": {
				Description: "Product identification of the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"platform_type": {
				Description: "The platform type of the device.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `IMCRack` - A standalone UCS M6 and above server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation.\n* `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications.\n* `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications.\n* `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial": {
				Description: "Serial number of the device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceTechsupportmanagementTechSupportBundle().Schema},
				Computed: true,
			}},
	}
}

func dataSourceTechsupportmanagementTechSupportBundleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.TechsupportmanagementTechSupportBundle{}
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
	if v, ok := d.GetOk("device_identifier"); ok {
		x := (v.(string))
		o.SetDeviceIdentifier(x)
	}
	if v, ok := d.GetOk("device_type"); ok {
		x := (v.(string))
		o.SetDeviceType(x)
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
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of TechsupportmanagementTechSupportBundle object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.TechsupportmanagementApi.GetTechsupportmanagementTechSupportBundleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of TechsupportmanagementTechSupportBundle: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of TechsupportmanagementTechSupportBundle: %s", responseErr.Error())
	}
	count := countResponse.TechsupportmanagementTechSupportBundleList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for TechsupportmanagementTechSupportBundle data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var techsupportmanagementTechSupportBundleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.TechsupportmanagementApi.GetTechsupportmanagementTechSupportBundleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching TechsupportmanagementTechSupportBundle: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching TechsupportmanagementTechSupportBundle: %s", responseErr.Error())
		}
		results := resMo.TechsupportmanagementTechSupportBundleList.GetResults()
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
				temp["device_identifier"] = (s.GetDeviceIdentifier())

				temp["device_registration"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDeviceRegistration(), d)
				temp["device_type"] = (s.GetDeviceType())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["pid"] = (s.GetPid())

				temp["platform_param"] = flattenMapConnectorPlatformParamBase(s.GetPlatformParam(), d)
				temp["platform_type"] = (s.GetPlatformType())
				temp["serial"] = (s.GetSerial())
				temp["shared_scope"] = (s.GetSharedScope())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["target_resource"] = flattenMapMoBaseMoRelationship(s.GetTargetResource(), d)

				temp["tech_support_status"] = flattenMapTechsupportmanagementTechSupportStatusRelationship(s.GetTechSupportStatus(), d)

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				techsupportmanagementTechSupportBundleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(techsupportmanagementTechSupportBundleResults))
	if err := d.Set("results", techsupportmanagementTechSupportBundleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(techsupportmanagementTechSupportBundleResults[0]["moid"].(string))
	return de
}
