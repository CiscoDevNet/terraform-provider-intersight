# AssetDeviceContractInformationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to [**AssetContractInformation**](asset.ContractInformation.md) |  | [optional] 
**ContractStatus** | Pointer to **string** | Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered. * &#x60;Not Covered&#x60; - The Cisco device does not have a valid support contract. * &#x60;Active&#x60; - The Cisco device is covered under a active support contract. * &#x60;Expiring Soon&#x60; - The contract for this Cisco device is going to expire in the next 30 days. | [optional] [readonly] [default to "Not Covered"]
**CoveredProductLineEndDate** | Pointer to **string** | End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future. * &#x60;None&#x60; - A default value to catch cases where device type is not correctly detected. * &#x60;CiscoUcsServer&#x60; - A device of type server. It includes Cisco IMC and UCS Managed servers. * &#x60;CiscoUcsFI&#x60; - A device of type Fabric Interconnect. It includes the various types of Cisco Fabric Interconnects supported by Cisco Intersight. * &#x60;CiscoUcsChassis&#x60; - A device of type Chassis. It includes various UCS chassis supported by Cisco Intersight. | [optional] [readonly] [default to "None"]
**EndCustomer** | Pointer to [**AssetCustomerInformation**](asset.CustomerInformation.md) |  | [optional] 
**EndUserGlobalUltimate** | Pointer to [**AssetGlobalUltimate**](asset.GlobalUltimate.md) |  | [optional] 
**IsValid** | Pointer to **bool** | Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs. | [optional] [readonly] 
**ItemType** | Pointer to **string** | Item type of this specific Cisco device. example \&quot;Chassis\&quot;. | [optional] [readonly] 
**MaintenancePurchaseOrderNumber** | Pointer to **string** | Maintenance purchase order number for the Cisco device. | [optional] [readonly] 
**MaintenanceSalesOrderNumber** | Pointer to **string** | Maintenance sales order number for the Cisco device. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type of the Cisco device. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [readonly] [default to ""]
**Product** | Pointer to [**AssetProductInformation**](asset.ProductInformation.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number for the Cisco device. It is a unique number assigned for every purchase. | [optional] [readonly] 
**ResellerGlobalUltimate** | Pointer to [**AssetGlobalUltimate**](asset.GlobalUltimate.md) |  | [optional] 
**SalesOrderNumber** | Pointer to **string** | Sales order number for the Cisco device. It is a unique number assigned for every sale. | [optional] [readonly] 
**ServiceDescription** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceEndDate** | Pointer to [**time.Time**](time.Time.md) | End date for the Cisco service contract that covers this Cisco device. | [optional] [readonly] 
**ServiceLevel** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceSku** | Pointer to **string** | The SKU of the service contract that covers the Cisco device. | [optional] [readonly] 
**ServiceStartDate** | Pointer to [**time.Time**](time.Time.md) | Start date for the Cisco service contract that covers this Cisco device. | [optional] [readonly] 
**StateContract** | Pointer to **string** | Internal property used for triggering and tracking actions for contract information. * &#x60;Update&#x60; - Sn2Info/Contract information needs to be updated. * &#x60;OK&#x60; - Sn2Info/Contract information was fetched succcessfuly and updated. * &#x60;Failed&#x60; - Sn2Info/Contract information was not available  or failed while fetching. * &#x60;Retry&#x60; - Sn2Info/Contract information update failed and will be retried later. | [optional] [default to "Update"]
**WarrantyEndDate** | Pointer to **string** | End date for the warranty that covers the Cisco device. | [optional] [readonly] 
**WarrantyType** | Pointer to **string** | Type of warranty that covers the Cisco device. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeviceContractInformationAllOf

`func NewAssetDeviceContractInformationAllOf() *AssetDeviceContractInformationAllOf`

NewAssetDeviceContractInformationAllOf instantiates a new AssetDeviceContractInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceContractInformationAllOfWithDefaults

`func NewAssetDeviceContractInformationAllOfWithDefaults() *AssetDeviceContractInformationAllOf`

NewAssetDeviceContractInformationAllOfWithDefaults instantiates a new AssetDeviceContractInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *AssetDeviceContractInformationAllOf) GetContract() AssetContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AssetDeviceContractInformationAllOf) GetContractOk() (*AssetContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AssetDeviceContractInformationAllOf) SetContract(v AssetContractInformation)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AssetDeviceContractInformationAllOf) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetContractStatus

`func (o *AssetDeviceContractInformationAllOf) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *AssetDeviceContractInformationAllOf) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *AssetDeviceContractInformationAllOf) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *AssetDeviceContractInformationAllOf) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetCoveredProductLineEndDate

`func (o *AssetDeviceContractInformationAllOf) GetCoveredProductLineEndDate() string`

GetCoveredProductLineEndDate returns the CoveredProductLineEndDate field if non-nil, zero value otherwise.

### GetCoveredProductLineEndDateOk

`func (o *AssetDeviceContractInformationAllOf) GetCoveredProductLineEndDateOk() (*string, bool)`

GetCoveredProductLineEndDateOk returns a tuple with the CoveredProductLineEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredProductLineEndDate

`func (o *AssetDeviceContractInformationAllOf) SetCoveredProductLineEndDate(v string)`

SetCoveredProductLineEndDate sets CoveredProductLineEndDate field to given value.

### HasCoveredProductLineEndDate

`func (o *AssetDeviceContractInformationAllOf) HasCoveredProductLineEndDate() bool`

HasCoveredProductLineEndDate returns a boolean if a field has been set.

### GetDeviceId

`func (o *AssetDeviceContractInformationAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetDeviceContractInformationAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetDeviceContractInformationAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetDeviceContractInformationAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *AssetDeviceContractInformationAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AssetDeviceContractInformationAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AssetDeviceContractInformationAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AssetDeviceContractInformationAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetEndCustomer

`func (o *AssetDeviceContractInformationAllOf) GetEndCustomer() AssetCustomerInformation`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *AssetDeviceContractInformationAllOf) GetEndCustomerOk() (*AssetCustomerInformation, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *AssetDeviceContractInformationAllOf) SetEndCustomer(v AssetCustomerInformation)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *AssetDeviceContractInformationAllOf) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### GetEndUserGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) GetEndUserGlobalUltimate() AssetGlobalUltimate`

GetEndUserGlobalUltimate returns the EndUserGlobalUltimate field if non-nil, zero value otherwise.

### GetEndUserGlobalUltimateOk

`func (o *AssetDeviceContractInformationAllOf) GetEndUserGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetEndUserGlobalUltimateOk returns a tuple with the EndUserGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) SetEndUserGlobalUltimate(v AssetGlobalUltimate)`

SetEndUserGlobalUltimate sets EndUserGlobalUltimate field to given value.

### HasEndUserGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) HasEndUserGlobalUltimate() bool`

HasEndUserGlobalUltimate returns a boolean if a field has been set.

### GetIsValid

`func (o *AssetDeviceContractInformationAllOf) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AssetDeviceContractInformationAllOf) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AssetDeviceContractInformationAllOf) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AssetDeviceContractInformationAllOf) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeviceContractInformationAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeviceContractInformationAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeviceContractInformationAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeviceContractInformationAllOf) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) GetMaintenancePurchaseOrderNumber() string`

GetMaintenancePurchaseOrderNumber returns the MaintenancePurchaseOrderNumber field if non-nil, zero value otherwise.

### GetMaintenancePurchaseOrderNumberOk

`func (o *AssetDeviceContractInformationAllOf) GetMaintenancePurchaseOrderNumberOk() (*string, bool)`

GetMaintenancePurchaseOrderNumberOk returns a tuple with the MaintenancePurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) SetMaintenancePurchaseOrderNumber(v string)`

SetMaintenancePurchaseOrderNumber sets MaintenancePurchaseOrderNumber field to given value.

### HasMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) HasMaintenancePurchaseOrderNumber() bool`

HasMaintenancePurchaseOrderNumber returns a boolean if a field has been set.

### GetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) GetMaintenanceSalesOrderNumber() string`

GetMaintenanceSalesOrderNumber returns the MaintenanceSalesOrderNumber field if non-nil, zero value otherwise.

### GetMaintenanceSalesOrderNumberOk

`func (o *AssetDeviceContractInformationAllOf) GetMaintenanceSalesOrderNumberOk() (*string, bool)`

GetMaintenanceSalesOrderNumberOk returns a tuple with the MaintenanceSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) SetMaintenanceSalesOrderNumber(v string)`

SetMaintenanceSalesOrderNumber sets MaintenanceSalesOrderNumber field to given value.

### HasMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) HasMaintenanceSalesOrderNumber() bool`

HasMaintenanceSalesOrderNumber returns a boolean if a field has been set.

### GetPlatformType

`func (o *AssetDeviceContractInformationAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AssetDeviceContractInformationAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AssetDeviceContractInformationAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AssetDeviceContractInformationAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetProduct

`func (o *AssetDeviceContractInformationAllOf) GetProduct() AssetProductInformation`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AssetDeviceContractInformationAllOf) GetProductOk() (*AssetProductInformation, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AssetDeviceContractInformationAllOf) SetProduct(v AssetProductInformation)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AssetDeviceContractInformationAllOf) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *AssetDeviceContractInformationAllOf) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *AssetDeviceContractInformationAllOf) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetResellerGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) GetResellerGlobalUltimate() AssetGlobalUltimate`

GetResellerGlobalUltimate returns the ResellerGlobalUltimate field if non-nil, zero value otherwise.

### GetResellerGlobalUltimateOk

`func (o *AssetDeviceContractInformationAllOf) GetResellerGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetResellerGlobalUltimateOk returns a tuple with the ResellerGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) SetResellerGlobalUltimate(v AssetGlobalUltimate)`

SetResellerGlobalUltimate sets ResellerGlobalUltimate field to given value.

### HasResellerGlobalUltimate

`func (o *AssetDeviceContractInformationAllOf) HasResellerGlobalUltimate() bool`

HasResellerGlobalUltimate returns a boolean if a field has been set.

### GetSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) GetSalesOrderNumber() string`

GetSalesOrderNumber returns the SalesOrderNumber field if non-nil, zero value otherwise.

### GetSalesOrderNumberOk

`func (o *AssetDeviceContractInformationAllOf) GetSalesOrderNumberOk() (*string, bool)`

GetSalesOrderNumberOk returns a tuple with the SalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) SetSalesOrderNumber(v string)`

SetSalesOrderNumber sets SalesOrderNumber field to given value.

### HasSalesOrderNumber

`func (o *AssetDeviceContractInformationAllOf) HasSalesOrderNumber() bool`

HasSalesOrderNumber returns a boolean if a field has been set.

### GetServiceDescription

`func (o *AssetDeviceContractInformationAllOf) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *AssetDeviceContractInformationAllOf) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *AssetDeviceContractInformationAllOf) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.

### HasServiceDescription

`func (o *AssetDeviceContractInformationAllOf) HasServiceDescription() bool`

HasServiceDescription returns a boolean if a field has been set.

### GetServiceEndDate

`func (o *AssetDeviceContractInformationAllOf) GetServiceEndDate() time.Time`

GetServiceEndDate returns the ServiceEndDate field if non-nil, zero value otherwise.

### GetServiceEndDateOk

`func (o *AssetDeviceContractInformationAllOf) GetServiceEndDateOk() (*time.Time, bool)`

GetServiceEndDateOk returns a tuple with the ServiceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndDate

`func (o *AssetDeviceContractInformationAllOf) SetServiceEndDate(v time.Time)`

SetServiceEndDate sets ServiceEndDate field to given value.

### HasServiceEndDate

`func (o *AssetDeviceContractInformationAllOf) HasServiceEndDate() bool`

HasServiceEndDate returns a boolean if a field has been set.

### GetServiceLevel

`func (o *AssetDeviceContractInformationAllOf) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *AssetDeviceContractInformationAllOf) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *AssetDeviceContractInformationAllOf) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *AssetDeviceContractInformationAllOf) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetServiceSku

`func (o *AssetDeviceContractInformationAllOf) GetServiceSku() string`

GetServiceSku returns the ServiceSku field if non-nil, zero value otherwise.

### GetServiceSkuOk

`func (o *AssetDeviceContractInformationAllOf) GetServiceSkuOk() (*string, bool)`

GetServiceSkuOk returns a tuple with the ServiceSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSku

`func (o *AssetDeviceContractInformationAllOf) SetServiceSku(v string)`

SetServiceSku sets ServiceSku field to given value.

### HasServiceSku

`func (o *AssetDeviceContractInformationAllOf) HasServiceSku() bool`

HasServiceSku returns a boolean if a field has been set.

### GetServiceStartDate

`func (o *AssetDeviceContractInformationAllOf) GetServiceStartDate() time.Time`

GetServiceStartDate returns the ServiceStartDate field if non-nil, zero value otherwise.

### GetServiceStartDateOk

`func (o *AssetDeviceContractInformationAllOf) GetServiceStartDateOk() (*time.Time, bool)`

GetServiceStartDateOk returns a tuple with the ServiceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartDate

`func (o *AssetDeviceContractInformationAllOf) SetServiceStartDate(v time.Time)`

SetServiceStartDate sets ServiceStartDate field to given value.

### HasServiceStartDate

`func (o *AssetDeviceContractInformationAllOf) HasServiceStartDate() bool`

HasServiceStartDate returns a boolean if a field has been set.

### GetStateContract

`func (o *AssetDeviceContractInformationAllOf) GetStateContract() string`

GetStateContract returns the StateContract field if non-nil, zero value otherwise.

### GetStateContractOk

`func (o *AssetDeviceContractInformationAllOf) GetStateContractOk() (*string, bool)`

GetStateContractOk returns a tuple with the StateContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateContract

`func (o *AssetDeviceContractInformationAllOf) SetStateContract(v string)`

SetStateContract sets StateContract field to given value.

### HasStateContract

`func (o *AssetDeviceContractInformationAllOf) HasStateContract() bool`

HasStateContract returns a boolean if a field has been set.

### GetWarrantyEndDate

`func (o *AssetDeviceContractInformationAllOf) GetWarrantyEndDate() string`

GetWarrantyEndDate returns the WarrantyEndDate field if non-nil, zero value otherwise.

### GetWarrantyEndDateOk

`func (o *AssetDeviceContractInformationAllOf) GetWarrantyEndDateOk() (*string, bool)`

GetWarrantyEndDateOk returns a tuple with the WarrantyEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEndDate

`func (o *AssetDeviceContractInformationAllOf) SetWarrantyEndDate(v string)`

SetWarrantyEndDate sets WarrantyEndDate field to given value.

### HasWarrantyEndDate

`func (o *AssetDeviceContractInformationAllOf) HasWarrantyEndDate() bool`

HasWarrantyEndDate returns a boolean if a field has been set.

### GetWarrantyType

`func (o *AssetDeviceContractInformationAllOf) GetWarrantyType() string`

GetWarrantyType returns the WarrantyType field if non-nil, zero value otherwise.

### GetWarrantyTypeOk

`func (o *AssetDeviceContractInformationAllOf) GetWarrantyTypeOk() (*string, bool)`

GetWarrantyTypeOk returns a tuple with the WarrantyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyType

`func (o *AssetDeviceContractInformationAllOf) SetWarrantyType(v string)`

SetWarrantyType sets WarrantyType field to given value.

### HasWarrantyType

`func (o *AssetDeviceContractInformationAllOf) HasWarrantyType() bool`

HasWarrantyType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetDeviceContractInformationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetDeviceContractInformationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetDeviceContractInformationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetDeviceContractInformationAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


