---
subcategory: "cloud"
layout: "intersight"
page_title: "Intersight: intersight_cloud_sku_volume_type"
description: |-
  Stores information about the volume types.
---

# Data Source: intersight_cloud_sku_volume_type
Stores information about the volume types.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_cloud_sku_volume_type.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Any additional description for the instance type. 
* `iops_unit`:(string) The units to measure IOPS. 
* `is_active`:(bool) Flag to indicate if this SKU is active or not. 
* `is_auto_discovered`:(bool) Flag to indicate if SKU is discovered during inventory collection. 
* `is_bootable`:(bool) Indicates if this volume can be used as a boot volume. 
* `is_default`:(bool) Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified. 
* `is_provisioned_iops`:(bool) Indicates if this volume type supports provisioned IOPS. 
* `max_iops`:(float) The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value. 
* `max_read_iops`:(float) The maximum read IOPS that this volume type supports. 
* `max_read_throughput`:(float) The maximum read throughput limit for this volume type. 
* `max_throughput`:(float) The maximum throughput limit for this volume type. 
* `max_volume_size`:(float) The maximum storage size that this volume type supports. 
* `max_write_iops`:(float) The maximum write IOPS that this volume type supports. 
* `max_write_throughput`:(float) The maximum write throughput limit for this volume type. 
* `min_volume_size`:(float) The minimum storage size that this volume type supports. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The display name for instance type. 
* `platform_type`:(string) The platformType for this SKU.* `` - The device reported an empty or unrecognized platform type.* `APIC` - An Application Policy Infrastructure Controller cluster.* `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.* `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).* `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.* `IMC` - A standalone UCS Server Integrated Management Controller.* `IMCM4` - A standalone UCS M4 Server.* `IMCM5` - A standalone UCS M5 server.* `UCSIOM` - An UCS Chassis IO module.* `HX` - A HyperFlex storage controller.* `HyperFlexAP` - A HyperFlex Application Platform.* `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.* `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance.* `IntersightAssist` - A Cisco Intersight Assist.* `PureStorageFlashArray` - A Pure Storage FlashArray device.* `NetAppOntap` - A NetApp ONTAP storage system.* `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager.* `EmcScaleIo` - An EMC ScaleIO storage system.* `EmcVmax` - An EMC VMAX storage system.* `EmcVplex` - An EMC VPLEX storage system.* `EmcXtremIo` - An EMC XtremIO storage system.* `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines.* `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines.* `AppDynamics` - An AppDynamics controller that monitors applications.* `Dynatrace` - A Dynatrace controller that monitors applications.* `MicrosoftSqlServer` - A Microsoft SQL database server.* `Kubernetes` - A Kubernetes cluster that runs containerized applications.* `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.* `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket.* `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.* `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.* `DellCompellent` - A Dell Compellent storage system.* `HPE3Par` - A HPE 3PAR storage system.* `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.* `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.* `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking.* `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.* `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.* `IMCBlade` - An Intersight managed UCS Blade Server.* `TerraformCloud` - A Terraform Cloud account.* `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.* `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.* `CiscoCatalyst` - A Cisco Catalyst networking switch device. 
* `service_category`:(string) Indicates if this sku belongs to Compute, Storage, Database or Network category.* `Compute` - Compute service offered by cloud provider.* `Storage` - Storage service offered by cloud provider.* `Database` - Database service offered by cloud provider.* `Network` - Network service offered by cloud provider. 
* `service_family`:(string) Property to identify the family of service that the sku belongs to. 
* `service_name`:(string) Any display name for the ServiceCategory if available. 
* `throughput_unit`:(string) The units for measuring throughput. 
* `type`:(string) The volume type like gp2 or st1. 
* `volume_size_unit`:(string) The units for measuring volume size. 
 