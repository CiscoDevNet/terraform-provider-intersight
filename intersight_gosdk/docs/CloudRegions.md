# CloudRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.Regions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.Regions"]
**AlternateNames** | Pointer to **[]string** |  | [optional] 
**DefaultZone** | Pointer to **string** | The default zone for this region. | [optional] 
**Group** | Pointer to **string** | Property to identify regions in same category which shares credentials. For e.g. AWS Govcloud Vs AWS Global Vs AWS China. | [optional] 
**IsActive** | Pointer to **bool** | Flag to indicate of this region is active or not. | [optional] [default to true]
**IsBillingOnly** | Pointer to **bool** | Property to pick regions for orchestration. | [optional] 
**Name** | Pointer to **string** | The display name of the region. | [optional] 
**PlatformType** | Pointer to **string** | The platform type for this region. For e.g. AmazonWebService. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;IMCRack&#x60; - A standalone UCS M6 and above server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;UCSC890&#x60; - A standalone Cisco UCSC890 server. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;ReadHatOpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;AmazonWebService&#x60; - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * &#x60;AmazonWebServiceBilling&#x60; - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * &#x60;DellCompellent&#x60; - A Dell Compellent storage system. * &#x60;HPE3Par&#x60; - A HPE 3PAR storage system. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * &#x60;HPEOneView&#x60; - A HPE Oneview management system that manages compute, storage, and networking. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. * &#x60;TerraformCloud&#x60; - A Terraform Cloud account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * &#x60;CustomTarget&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;HTTPEndpoint&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. | [optional] [default to ""]
**RegionEndPoint** | Pointer to **string** | HTTP endpoint of the region. For example https://ec2.us-east-2.amazonaws.com. | [optional] 
**RegionId** | Pointer to **string** | The region Id which is assigned by the cloud provider. For e.g. us-east-1. | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 
**Target** | Pointer to [**AssetTargetRelationship**](asset.Target.Relationship.md) |  | [optional] 

## Methods

### NewCloudRegions

`func NewCloudRegions(classId string, objectType string, ) *CloudRegions`

NewCloudRegions instantiates a new CloudRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegionsWithDefaults

`func NewCloudRegionsWithDefaults() *CloudRegions`

NewCloudRegionsWithDefaults instantiates a new CloudRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudRegions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudRegions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudRegions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudRegions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudRegions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudRegions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlternateNames

`func (o *CloudRegions) GetAlternateNames() []string`

GetAlternateNames returns the AlternateNames field if non-nil, zero value otherwise.

### GetAlternateNamesOk

`func (o *CloudRegions) GetAlternateNamesOk() (*[]string, bool)`

GetAlternateNamesOk returns a tuple with the AlternateNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateNames

`func (o *CloudRegions) SetAlternateNames(v []string)`

SetAlternateNames sets AlternateNames field to given value.

### HasAlternateNames

`func (o *CloudRegions) HasAlternateNames() bool`

HasAlternateNames returns a boolean if a field has been set.

### SetAlternateNamesNil

`func (o *CloudRegions) SetAlternateNamesNil(b bool)`

 SetAlternateNamesNil sets the value for AlternateNames to be an explicit nil

### UnsetAlternateNames
`func (o *CloudRegions) UnsetAlternateNames()`

UnsetAlternateNames ensures that no value is present for AlternateNames, not even an explicit nil
### GetDefaultZone

`func (o *CloudRegions) GetDefaultZone() string`

GetDefaultZone returns the DefaultZone field if non-nil, zero value otherwise.

### GetDefaultZoneOk

`func (o *CloudRegions) GetDefaultZoneOk() (*string, bool)`

GetDefaultZoneOk returns a tuple with the DefaultZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultZone

`func (o *CloudRegions) SetDefaultZone(v string)`

SetDefaultZone sets DefaultZone field to given value.

### HasDefaultZone

`func (o *CloudRegions) HasDefaultZone() bool`

HasDefaultZone returns a boolean if a field has been set.

### GetGroup

`func (o *CloudRegions) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CloudRegions) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CloudRegions) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CloudRegions) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIsActive

`func (o *CloudRegions) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CloudRegions) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CloudRegions) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CloudRegions) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsBillingOnly

`func (o *CloudRegions) GetIsBillingOnly() bool`

GetIsBillingOnly returns the IsBillingOnly field if non-nil, zero value otherwise.

### GetIsBillingOnlyOk

`func (o *CloudRegions) GetIsBillingOnlyOk() (*bool, bool)`

GetIsBillingOnlyOk returns a tuple with the IsBillingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingOnly

`func (o *CloudRegions) SetIsBillingOnly(v bool)`

SetIsBillingOnly sets IsBillingOnly field to given value.

### HasIsBillingOnly

`func (o *CloudRegions) HasIsBillingOnly() bool`

HasIsBillingOnly returns a boolean if a field has been set.

### GetName

`func (o *CloudRegions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRegions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRegions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRegions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *CloudRegions) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CloudRegions) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CloudRegions) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CloudRegions) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetRegionEndPoint

`func (o *CloudRegions) GetRegionEndPoint() string`

GetRegionEndPoint returns the RegionEndPoint field if non-nil, zero value otherwise.

### GetRegionEndPointOk

`func (o *CloudRegions) GetRegionEndPointOk() (*string, bool)`

GetRegionEndPointOk returns a tuple with the RegionEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionEndPoint

`func (o *CloudRegions) SetRegionEndPoint(v string)`

SetRegionEndPoint sets RegionEndPoint field to given value.

### HasRegionEndPoint

`func (o *CloudRegions) HasRegionEndPoint() bool`

HasRegionEndPoint returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudRegions) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudRegions) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudRegions) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudRegions) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetZones

`func (o *CloudRegions) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CloudRegions) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CloudRegions) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CloudRegions) HasZones() bool`

HasZones returns a boolean if a field has been set.

### SetZonesNil

`func (o *CloudRegions) SetZonesNil(b bool)`

 SetZonesNil sets the value for Zones to be an explicit nil

### UnsetZones
`func (o *CloudRegions) UnsetZones()`

UnsetZones ensures that no value is present for Zones, not even an explicit nil
### GetTarget

`func (o *CloudRegions) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudRegions) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudRegions) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudRegions) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


