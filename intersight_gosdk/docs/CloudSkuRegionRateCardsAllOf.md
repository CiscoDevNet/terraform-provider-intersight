# CloudSkuRegionRateCardsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.SkuRegionRateCards"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.SkuRegionRateCards"]
**Currency** | Pointer to **string** | The currency code used for the price. For e.g. USD or EUR. * &#x60;USD&#x60; - The currency code for United states dollar. * &#x60;EUR&#x60; - The currency code for European Union. | [optional] [default to "USD"]
**CustomAttributes** | Pointer to [**[]CloudCustomAttributes**](CloudCustomAttributes.md) |  | [optional] 
**DistributionType** | Pointer to **string** | The OS distribution running on this instance type. | [optional] 
**IsUserDefined** | Pointer to **bool** | Applicable for private cloud where price is set by the user. | [optional] 
**PlatformType** | Pointer to **string** | The platformType for this SKU. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;CAPIC&#x60; - An Application Policy Infrastructure Controller cloud instance. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;IMCRack&#x60; - A standalone UCS M6 and above server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;IWE&#x60; - An Intersight Workload Engine. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - A Cisco Intersight Connected Virtual Appliance. * &#x60;IntersightAssist&#x60; - A Cisco Intersight Assist. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NexusDevice&#x60; - A generic platform type to support Nexus Network Device. This can also be extended to support all network devices later on. * &#x60;ACISwitch&#x60; - A platform type to support ACI Switches. * &#x60;NexusSwitch&#x60; - A platform type to support Cisco Nexus Switches. * &#x60;MDSDevice&#x60; - A platform type to support MDS devices. * &#x60;UCSC890&#x60; - A standalone Cisco UCSC890 server. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;NetAppActiveIqUnifiedManager&#x60; - A NetApp Active IQ Unified Manager. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft Hyper-V system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;NewRelic&#x60; - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * &#x60;ServiceNow&#x60; - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * &#x60;ReadHatOpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;CloudFoundry&#x60; - An open source cloud platform on which developers can build, deploy, run and scale applications. * &#x60;MicrosoftAzureApplicationInsights&#x60; - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * &#x60;OpenStack&#x60; - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;MySqlServer&#x60; - An instance of either Oracle MySQL Database or the open source MariaDB. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;AmazonWebService&#x60; - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * &#x60;AmazonWebServiceBilling&#x60; - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * &#x60;MicrosoftAzureServicePrincipal&#x60; - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * &#x60;MicrosoftAzureEnterpriseAgreement&#x60; - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * &#x60;DellCompellent&#x60; - A Dell Compellent storage system. * &#x60;HPE3Par&#x60; - A HPE 3PAR storage system. * &#x60;RedHatEnterpriseVirtualization&#x60; - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * &#x60;NutanixAcropolis&#x60; - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * &#x60;HPEOneView&#x60; - A HPE Oneview management system that manages compute, storage, and networking. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;HitachiVirtualStoragePlatform&#x60; - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. * &#x60;TerraformCloud&#x60; - A Terraform Cloud account. * &#x60;TerraformAgent&#x60; - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * &#x60;CustomTarget&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * &#x60;AnsibleEndpoint&#x60; - An external endpoint added as Target that can be accessed through Ansible in Intersight Cloud Orchestrator automation workflow. * &#x60;HTTPEndpoint&#x60; - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic, Bearer Token. * &#x60;SSHEndpoint&#x60; - An external endpoint added as Target that can be accessed through SSH in Intersight Cloud Orchestrator automation workflow. * &#x60;CiscoCatalyst&#x60; - A Cisco Catalyst networking switch device. * &#x60;PowerShellEndpoint&#x60; - A Windows machine on which PowerShell scripts can be executed remotely. | [optional] [default to ""]
**Price** | Pointer to **float32** | Price of the instance type. | [optional] 
**RegionId** | Pointer to **string** | The regionId associated with this rate card. | [optional] 
**ServiceCategory** | Pointer to **string** | Indicates if this sku belongs to Compute, Storage, Database or Network category. * &#x60;Compute&#x60; - Compute service offered by cloud provider. * &#x60;Storage&#x60; - Storage service offered by cloud provider. * &#x60;Database&#x60; - Database service offered by cloud provider. * &#x60;Network&#x60; - Network service offered by cloud provider. | [optional] [default to "Compute"]
**SkuName** | Pointer to **string** | The sku name associated with this rate card. | [optional] 
**Unit** | Pointer to **string** | The billing unit to use for computing the price. For e.g. when serviceCategory is Compute the unit will be \&quot;Hrs\&quot;, for Storage it will be \&quot;GB-Mo\&quot;. | [optional] 
**ValidFrom** | Pointer to **int64** | The epoch start time from which the price will be applied. | [optional] 
**ValidTo** | Pointer to **int64** | The epoch end time of the current price. | [optional] 
**Region** | Pointer to [**CloudRegionsRelationship**](CloudRegionsRelationship.md) |  | [optional] 
**Sku** | Pointer to [**CloudBaseSkuRelationship**](CloudBaseSkuRelationship.md) |  | [optional] 

## Methods

### NewCloudSkuRegionRateCardsAllOf

`func NewCloudSkuRegionRateCardsAllOf(classId string, objectType string, ) *CloudSkuRegionRateCardsAllOf`

NewCloudSkuRegionRateCardsAllOf instantiates a new CloudSkuRegionRateCardsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkuRegionRateCardsAllOfWithDefaults

`func NewCloudSkuRegionRateCardsAllOfWithDefaults() *CloudSkuRegionRateCardsAllOf`

NewCloudSkuRegionRateCardsAllOfWithDefaults instantiates a new CloudSkuRegionRateCardsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSkuRegionRateCardsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSkuRegionRateCardsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSkuRegionRateCardsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSkuRegionRateCardsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSkuRegionRateCardsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSkuRegionRateCardsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrency

`func (o *CloudSkuRegionRateCardsAllOf) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudSkuRegionRateCardsAllOf) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudSkuRegionRateCardsAllOf) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudSkuRegionRateCardsAllOf) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *CloudSkuRegionRateCardsAllOf) GetCustomAttributes() []CloudCustomAttributes`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *CloudSkuRegionRateCardsAllOf) GetCustomAttributesOk() (*[]CloudCustomAttributes, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *CloudSkuRegionRateCardsAllOf) SetCustomAttributes(v []CloudCustomAttributes)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *CloudSkuRegionRateCardsAllOf) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *CloudSkuRegionRateCardsAllOf) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *CloudSkuRegionRateCardsAllOf) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetDistributionType

`func (o *CloudSkuRegionRateCardsAllOf) GetDistributionType() string`

GetDistributionType returns the DistributionType field if non-nil, zero value otherwise.

### GetDistributionTypeOk

`func (o *CloudSkuRegionRateCardsAllOf) GetDistributionTypeOk() (*string, bool)`

GetDistributionTypeOk returns a tuple with the DistributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionType

`func (o *CloudSkuRegionRateCardsAllOf) SetDistributionType(v string)`

SetDistributionType sets DistributionType field to given value.

### HasDistributionType

`func (o *CloudSkuRegionRateCardsAllOf) HasDistributionType() bool`

HasDistributionType returns a boolean if a field has been set.

### GetIsUserDefined

`func (o *CloudSkuRegionRateCardsAllOf) GetIsUserDefined() bool`

GetIsUserDefined returns the IsUserDefined field if non-nil, zero value otherwise.

### GetIsUserDefinedOk

`func (o *CloudSkuRegionRateCardsAllOf) GetIsUserDefinedOk() (*bool, bool)`

GetIsUserDefinedOk returns a tuple with the IsUserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserDefined

`func (o *CloudSkuRegionRateCardsAllOf) SetIsUserDefined(v bool)`

SetIsUserDefined sets IsUserDefined field to given value.

### HasIsUserDefined

`func (o *CloudSkuRegionRateCardsAllOf) HasIsUserDefined() bool`

HasIsUserDefined returns a boolean if a field has been set.

### GetPlatformType

`func (o *CloudSkuRegionRateCardsAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *CloudSkuRegionRateCardsAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *CloudSkuRegionRateCardsAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *CloudSkuRegionRateCardsAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPrice

`func (o *CloudSkuRegionRateCardsAllOf) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CloudSkuRegionRateCardsAllOf) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CloudSkuRegionRateCardsAllOf) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CloudSkuRegionRateCardsAllOf) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudSkuRegionRateCardsAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudSkuRegionRateCardsAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudSkuRegionRateCardsAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudSkuRegionRateCardsAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetServiceCategory

`func (o *CloudSkuRegionRateCardsAllOf) GetServiceCategory() string`

GetServiceCategory returns the ServiceCategory field if non-nil, zero value otherwise.

### GetServiceCategoryOk

`func (o *CloudSkuRegionRateCardsAllOf) GetServiceCategoryOk() (*string, bool)`

GetServiceCategoryOk returns a tuple with the ServiceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCategory

`func (o *CloudSkuRegionRateCardsAllOf) SetServiceCategory(v string)`

SetServiceCategory sets ServiceCategory field to given value.

### HasServiceCategory

`func (o *CloudSkuRegionRateCardsAllOf) HasServiceCategory() bool`

HasServiceCategory returns a boolean if a field has been set.

### GetSkuName

`func (o *CloudSkuRegionRateCardsAllOf) GetSkuName() string`

GetSkuName returns the SkuName field if non-nil, zero value otherwise.

### GetSkuNameOk

`func (o *CloudSkuRegionRateCardsAllOf) GetSkuNameOk() (*string, bool)`

GetSkuNameOk returns a tuple with the SkuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuName

`func (o *CloudSkuRegionRateCardsAllOf) SetSkuName(v string)`

SetSkuName sets SkuName field to given value.

### HasSkuName

`func (o *CloudSkuRegionRateCardsAllOf) HasSkuName() bool`

HasSkuName returns a boolean if a field has been set.

### GetUnit

`func (o *CloudSkuRegionRateCardsAllOf) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CloudSkuRegionRateCardsAllOf) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CloudSkuRegionRateCardsAllOf) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CloudSkuRegionRateCardsAllOf) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValidFrom

`func (o *CloudSkuRegionRateCardsAllOf) GetValidFrom() int64`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CloudSkuRegionRateCardsAllOf) GetValidFromOk() (*int64, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CloudSkuRegionRateCardsAllOf) SetValidFrom(v int64)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *CloudSkuRegionRateCardsAllOf) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *CloudSkuRegionRateCardsAllOf) GetValidTo() int64`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CloudSkuRegionRateCardsAllOf) GetValidToOk() (*int64, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CloudSkuRegionRateCardsAllOf) SetValidTo(v int64)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *CloudSkuRegionRateCardsAllOf) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetRegion

`func (o *CloudSkuRegionRateCardsAllOf) GetRegion() CloudRegionsRelationship`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudSkuRegionRateCardsAllOf) GetRegionOk() (*CloudRegionsRelationship, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudSkuRegionRateCardsAllOf) SetRegion(v CloudRegionsRelationship)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudSkuRegionRateCardsAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSku

`func (o *CloudSkuRegionRateCardsAllOf) GetSku() CloudBaseSkuRelationship`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CloudSkuRegionRateCardsAllOf) GetSkuOk() (*CloudBaseSkuRelationship, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CloudSkuRegionRateCardsAllOf) SetSku(v CloudBaseSkuRelationship)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CloudSkuRegionRateCardsAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


