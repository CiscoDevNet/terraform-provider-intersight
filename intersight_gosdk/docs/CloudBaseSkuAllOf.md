# CloudBaseSkuAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CustomAttributes** | Pointer to [**[]CloudCustomAttributes**](CloudCustomAttributes.md) |  | [optional] 
**Description** | Pointer to **string** | Any additional description for the instance type. | [optional] 
**IsActive** | Pointer to **bool** | Flag to indicate if this SKU is active or not. | [optional] [default to true]
**IsAutoDiscovered** | Pointer to **bool** | Flag to indicate if SKU is discovered during inventory collection. | [optional] 
**Name** | Pointer to **string** | The display name for instance type. | [optional] 
**PlatformType** | Pointer to **string** | The platformType for this SKU. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;IMCRack&#x60; - A standalone UCS M6 and above server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;UCSC890&#x60; - A standalone Cisco UCSC890 server. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;ReadHatOpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;AmazonWebService&#x60; - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * &#x60;AmazonWebServiceBilling&#x60; - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * &#x60;DellCompellent&#x60; - A Dell Compellent storage system. * &#x60;HPE3Par&#x60; - A HPE 3PAR storage system. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * &#x60;HPEOneView&#x60; - A HPE Oneview management system that manages compute, storage, and networking. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. * &#x60;TerraformCloud&#x60; - A Terraform Cloud account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * &#x60;CustomTarget&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;HTTPEndpoint&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. | [optional] [default to ""]
**ServiceCategory** | Pointer to **string** | Indicates if this sku belongs to Compute, Storage, Database or Network category. * &#x60;Compute&#x60; - Compute service offered by cloud provider. * &#x60;Storage&#x60; - Storage service offered by cloud provider. * &#x60;Database&#x60; - Database service offered by cloud provider. * &#x60;Network&#x60; - Network service offered by cloud provider. | [optional] [default to "Compute"]
**ServiceFamily** | Pointer to **string** | Property to identify the family of service that the sku belongs to. | [optional] 
**ServiceName** | Pointer to **string** | Any display name for the ServiceCategory if available. | [optional] 
**Target** | Pointer to [**AssetTargetRelationship**](asset.Target.Relationship.md) |  | [optional] 

## Methods

### NewCloudBaseSkuAllOf

`func NewCloudBaseSkuAllOf(classId string, objectType string, ) *CloudBaseSkuAllOf`

NewCloudBaseSkuAllOf instantiates a new CloudBaseSkuAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseSkuAllOfWithDefaults

`func NewCloudBaseSkuAllOfWithDefaults() *CloudBaseSkuAllOf`

NewCloudBaseSkuAllOfWithDefaults instantiates a new CloudBaseSkuAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseSkuAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseSkuAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseSkuAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseSkuAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseSkuAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseSkuAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCustomAttributes

`func (o *CloudBaseSkuAllOf) GetCustomAttributes() []CloudCustomAttributes`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *CloudBaseSkuAllOf) GetCustomAttributesOk() (*[]CloudCustomAttributes, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *CloudBaseSkuAllOf) SetCustomAttributes(v []CloudCustomAttributes)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *CloudBaseSkuAllOf) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *CloudBaseSkuAllOf) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *CloudBaseSkuAllOf) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetDescription

`func (o *CloudBaseSkuAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudBaseSkuAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudBaseSkuAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudBaseSkuAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *CloudBaseSkuAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CloudBaseSkuAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CloudBaseSkuAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CloudBaseSkuAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsAutoDiscovered

`func (o *CloudBaseSkuAllOf) GetIsAutoDiscovered() bool`

GetIsAutoDiscovered returns the IsAutoDiscovered field if non-nil, zero value otherwise.

### GetIsAutoDiscoveredOk

`func (o *CloudBaseSkuAllOf) GetIsAutoDiscoveredOk() (*bool, bool)`

GetIsAutoDiscoveredOk returns a tuple with the IsAutoDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDiscovered

`func (o *CloudBaseSkuAllOf) SetIsAutoDiscovered(v bool)`

SetIsAutoDiscovered sets IsAutoDiscovered field to given value.

### HasIsAutoDiscovered

`func (o *CloudBaseSkuAllOf) HasIsAutoDiscovered() bool`

HasIsAutoDiscovered returns a boolean if a field has been set.

### GetName

`func (o *CloudBaseSkuAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBaseSkuAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBaseSkuAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBaseSkuAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *CloudBaseSkuAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CloudBaseSkuAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CloudBaseSkuAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CloudBaseSkuAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetServiceCategory

`func (o *CloudBaseSkuAllOf) GetServiceCategory() string`

GetServiceCategory returns the ServiceCategory field if non-nil, zero value otherwise.

### GetServiceCategoryOk

`func (o *CloudBaseSkuAllOf) GetServiceCategoryOk() (*string, bool)`

GetServiceCategoryOk returns a tuple with the ServiceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCategory

`func (o *CloudBaseSkuAllOf) SetServiceCategory(v string)`

SetServiceCategory sets ServiceCategory field to given value.

### HasServiceCategory

`func (o *CloudBaseSkuAllOf) HasServiceCategory() bool`

HasServiceCategory returns a boolean if a field has been set.

### GetServiceFamily

`func (o *CloudBaseSkuAllOf) GetServiceFamily() string`

GetServiceFamily returns the ServiceFamily field if non-nil, zero value otherwise.

### GetServiceFamilyOk

`func (o *CloudBaseSkuAllOf) GetServiceFamilyOk() (*string, bool)`

GetServiceFamilyOk returns a tuple with the ServiceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFamily

`func (o *CloudBaseSkuAllOf) SetServiceFamily(v string)`

SetServiceFamily sets ServiceFamily field to given value.

### HasServiceFamily

`func (o *CloudBaseSkuAllOf) HasServiceFamily() bool`

HasServiceFamily returns a boolean if a field has been set.

### GetServiceName

`func (o *CloudBaseSkuAllOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CloudBaseSkuAllOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CloudBaseSkuAllOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CloudBaseSkuAllOf) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetTarget

`func (o *CloudBaseSkuAllOf) GetTarget() AssetTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudBaseSkuAllOf) GetTargetOk() (*AssetTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudBaseSkuAllOf) SetTarget(v AssetTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudBaseSkuAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


