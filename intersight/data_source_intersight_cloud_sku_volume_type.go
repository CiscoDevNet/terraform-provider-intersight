package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudSkuVolumeType() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudSkuVolumeTypeRead,
		Schema: map[string]*schema.Schema{
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Any additional description for the instance type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"iops_unit": {
				Description: "The units to measure IOPS.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_active": {
				Description: "Flag to indicate if this SKU is active or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_auto_discovered": {
				Description: "Flag to indicate if SKU is discovered during inventory collection.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_bootable": {
				Description: "Indicates if this volume can be used as a boot volume.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_default": {
				Description: "Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"is_provisioned_iops": {
				Description: "Indicates if this volume type supports provisioned IOPS.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"max_iops": {
				Description: "The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_read_iops": {
				Description: "The maximum read IOPS that this volume type supports.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_read_throughput": {
				Description: "The maximum read throughput limit for this volume type.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_throughput": {
				Description: "The maximum throughput limit for this volume type.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_volume_size": {
				Description: "The maximum storage size that this volume type supports.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_write_iops": {
				Description: "The maximum write IOPS that this volume type supports.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"max_write_throughput": {
				Description: "The maximum write throughput limit for this volume type.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"min_volume_size": {
				Description: "The minimum storage size that this volume type supports.",
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The display name for instance type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"platform_type": {
				Description: "The platformType for this SKU.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A Dynatrace controller that monitors applications.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"service_category": {
				Description: "Indicates if this sku belongs to Compute, Storage, Database or Network category.\n* `Compute` - Compute service offered by cloud provider.\n* `Storage` - Storage service offered by cloud provider.\n* `Database` - Database service offered by cloud provider.\n* `Network` - Network service offered by cloud provider.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"service_family": {
				Description: "Property to identify the family of service that the sku belongs to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"service_name": {
				Description: "Any display name for the ServiceCategory if available.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"throughput_unit": {
				Description: "The units for measuring throughput.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "The volume type like gp2 or st1.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"volume_size_unit": {
				Description: "The units for measuring volume size.",
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
					"custom_attributes": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"attribute_name": {
									Description: "The name of an attribute. If used as a key-value pair then this field represents the key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"attribute_type": {
									Description: "The data type for attributeValue. For e.g. string, int, float.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"attribute_value": {
									Description: "The attribute value. If used as a key-value pair then this field represents the value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"description": {
						Description: "Any additional description for the instance type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"iops_unit": {
						Description: "The units to measure IOPS.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"is_active": {
						Description: "Flag to indicate if this SKU is active or not.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_auto_discovered": {
						Description: "Flag to indicate if SKU is discovered during inventory collection.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_bootable": {
						Description: "Indicates if this volume can be used as a boot volume.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_default": {
						Description: "Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"is_provisioned_iops": {
						Description: "Indicates if this volume type supports provisioned IOPS.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"max_iops": {
						Description: "The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_read_iops": {
						Description: "The maximum read IOPS that this volume type supports.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_read_throughput": {
						Description: "The maximum read throughput limit for this volume type.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_throughput": {
						Description: "The maximum throughput limit for this volume type.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_volume_size": {
						Description: "The maximum storage size that this volume type supports.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_write_iops": {
						Description: "The maximum write IOPS that this volume type supports.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"max_write_throughput": {
						Description: "The maximum write throughput limit for this volume type.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"min_volume_size": {
						Description: "The minimum storage size that this volume type supports.",
						Type:        schema.TypeFloat,
						Optional:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "The display name for instance type.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"platform_type": {
						Description: "The platformType for this SKU.\n* `` - The device reported an empty or unrecognized platform type.\n* `APIC` - An Application Policy Infrastructure Controller cluster.\n* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.\n* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).\n* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.\n* `IMC` - A standalone UCS Server Integrated Management Controller.\n* `IMCM4` - A standalone UCS M4 Server.\n* `IMCM5` - A standalone UCS M5 server.\n* `UCSIOM` - An UCS Chassis IO module.\n* `HX` - A HyperFlex storage controller.\n* `HyperFlexAP` - A HyperFlex Application Platform.\n* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.\n* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.\n* `IntersightAssist` - A Cisco Intersight Assist.\n* `PureStorageFlashArray` - A Pure Storage FlashArray device.\n* `NetAppOntap` - A NetApp ONTAP storage system.\n* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.\n* `EmcScaleIo` - An EMC ScaleIO storage system.\n* `EmcVmax` - An EMC VMAX storage system.\n* `EmcVplex` - An EMC VPLEX storage system.\n* `EmcXtremIo` - An EMC XtremIO storage system.\n* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.\n* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.\n* `AppDynamics` - An AppDynamics controller that monitors applications.\n* `Dynatrace` - A Dynatrace controller that monitors applications.\n* `MicrosoftSqlServer` - A Microsoft SQL database server.\n* `Kubernetes` - A Kubernetes cluster that runs containerized applications.\n* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.\n* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.\n* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.\n* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.\n* `DellCompellent` - A Dell Compellent storage system.\n* `HPE3Par` - A HPE 3PAR storage system.\n* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.\n* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.\n* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.\n* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.\n* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.\n* `IMCBlade` - An Intersight managed UCS Blade Server.\n* `TerraformCloud` - A Terraform Cloud account.\n* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.\n* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.\n* `CiscoCatalyst` - A Cisco Catalyst networking switch device.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"service_category": {
						Description: "Indicates if this sku belongs to Compute, Storage, Database or Network category.\n* `Compute` - Compute service offered by cloud provider.\n* `Storage` - Storage service offered by cloud provider.\n* `Database` - Database service offered by cloud provider.\n* `Network` - Network service offered by cloud provider.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"service_family": {
						Description: "Property to identify the family of service that the sku belongs to.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"service_name": {
						Description: "Any display name for the ServiceCategory if available.",
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
					"target": {
						Description: "A reference to a assetTarget resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
					"throughput_unit": {
						Description: "The units for measuring throughput.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "The volume type like gp2 or st1.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"volume_size_unit": {
						Description: "The units for measuring volume size.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceCloudSkuVolumeTypeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CloudSkuVolumeType{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("iops_unit"); ok {
		x := (v.(string))
		o.SetIopsUnit(x)
	}
	if v, ok := d.GetOk("is_active"); ok {
		x := (v.(bool))
		o.SetIsActive(x)
	}
	if v, ok := d.GetOk("is_auto_discovered"); ok {
		x := (v.(bool))
		o.SetIsAutoDiscovered(x)
	}
	if v, ok := d.GetOk("is_bootable"); ok {
		x := (v.(bool))
		o.SetIsBootable(x)
	}
	if v, ok := d.GetOk("is_default"); ok {
		x := (v.(bool))
		o.SetIsDefault(x)
	}
	if v, ok := d.GetOk("is_provisioned_iops"); ok {
		x := (v.(bool))
		o.SetIsProvisionedIops(x)
	}
	if v, ok := d.GetOk("max_iops"); ok {
		x := v.(float64)
		o.SetMaxIops(x)
	}
	if v, ok := d.GetOk("max_read_iops"); ok {
		x := v.(float64)
		o.SetMaxReadIops(x)
	}
	if v, ok := d.GetOk("max_read_throughput"); ok {
		x := v.(float64)
		o.SetMaxReadThroughput(x)
	}
	if v, ok := d.GetOk("max_throughput"); ok {
		x := v.(float64)
		o.SetMaxThroughput(x)
	}
	if v, ok := d.GetOk("max_volume_size"); ok {
		x := v.(float64)
		o.SetMaxVolumeSize(x)
	}
	if v, ok := d.GetOk("max_write_iops"); ok {
		x := v.(float64)
		o.SetMaxWriteIops(x)
	}
	if v, ok := d.GetOk("max_write_throughput"); ok {
		x := v.(float64)
		o.SetMaxWriteThroughput(x)
	}
	if v, ok := d.GetOk("min_volume_size"); ok {
		x := v.(float64)
		o.SetMinVolumeSize(x)
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
	if v, ok := d.GetOk("platform_type"); ok {
		x := (v.(string))
		o.SetPlatformType(x)
	}
	if v, ok := d.GetOk("service_category"); ok {
		x := (v.(string))
		o.SetServiceCategory(x)
	}
	if v, ok := d.GetOk("service_family"); ok {
		x := (v.(string))
		o.SetServiceFamily(x)
	}
	if v, ok := d.GetOk("service_name"); ok {
		x := (v.(string))
		o.SetServiceName(x)
	}
	if v, ok := d.GetOk("throughput_unit"); ok {
		x := (v.(string))
		o.SetThroughputUnit(x)
	}
	if v, ok := d.GetOk("type"); ok {
		x := (v.(string))
		o.SetType(x)
	}
	if v, ok := d.GetOk("volume_size_unit"); ok {
		x := (v.(string))
		o.SetVolumeSizeUnit(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CloudSkuVolumeType object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CloudApi.GetCloudSkuVolumeTypeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of CloudSkuVolumeType: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.CloudSkuVolumeTypeList.GetCount()
	var i int32
	var cloudSkuVolumeTypeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CloudApi.GetCloudSkuVolumeTypeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching CloudSkuVolumeType: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.CloudSkuVolumeTypeList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for CloudSkuVolumeType data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())

				temp["custom_attributes"] = flattenListCloudCustomAttributes(s.GetCustomAttributes(), d)
				temp["description"] = (s.GetDescription())
				temp["iops_unit"] = (s.GetIopsUnit())
				temp["is_active"] = (s.GetIsActive())
				temp["is_auto_discovered"] = (s.GetIsAutoDiscovered())
				temp["is_bootable"] = (s.GetIsBootable())
				temp["is_default"] = (s.GetIsDefault())
				temp["is_provisioned_iops"] = (s.GetIsProvisionedIops())
				temp["max_iops"] = (s.GetMaxIops())
				temp["max_read_iops"] = (s.GetMaxReadIops())
				temp["max_read_throughput"] = (s.GetMaxReadThroughput())
				temp["max_throughput"] = (s.GetMaxThroughput())
				temp["max_volume_size"] = (s.GetMaxVolumeSize())
				temp["max_write_iops"] = (s.GetMaxWriteIops())
				temp["max_write_throughput"] = (s.GetMaxWriteThroughput())
				temp["min_volume_size"] = (s.GetMinVolumeSize())
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["object_type"] = (s.GetObjectType())
				temp["platform_type"] = (s.GetPlatformType())
				temp["service_category"] = (s.GetServiceCategory())
				temp["service_family"] = (s.GetServiceFamily())
				temp["service_name"] = (s.GetServiceName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["target"] = flattenMapAssetTargetRelationship(s.GetTarget(), d)
				temp["throughput_unit"] = (s.GetThroughputUnit())
				temp["type"] = (s.GetType())
				temp["volume_size_unit"] = (s.GetVolumeSizeUnit())
				cloudSkuVolumeTypeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(cloudSkuVolumeTypeResults))
	if err := d.Set("results", cloudSkuVolumeTypeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(cloudSkuVolumeTypeResults[0]["moid"].(string))
	return de
}
