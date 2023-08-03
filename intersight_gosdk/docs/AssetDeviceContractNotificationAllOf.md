# AssetDeviceContractNotificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceContractNotification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceContractNotification"]
**Contract** | Pointer to [**NullableAssetContractInformation**](AssetContractInformation.md) |  | [optional] 
**ContractStatus** | Pointer to **string** | Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered. * &#x60;Unknown&#x60; - The device&#39;s contract status cannot be determined. * &#x60;Not Covered&#x60; - The Cisco device does not have a valid support contract. * &#x60;Active&#x60; - The Cisco device is covered under a active support contract. * &#x60;Expiring Soon&#x60; - The contract for this Cisco device is going to expire in the next 30 days. | [optional] [default to "Unknown"]
**ContractStatusReason** | Pointer to **string** | Reason for contract status. In case of Not Covered, reason is either Terminated or Expired. * &#x60;&#x60; - There is no reason for the specified contract status. * &#x60;Line Item Expired&#x60; - The Cisco device does not have a valid support contract, it has expired. * &#x60;Line Item Terminated&#x60; - The Cisco device does not have a valid support contract, it has been terminated. | [optional] [default to ""]
**ContractUpdatedTime** | Pointer to **time.Time** | Date and time indicating when the contract data is last refreshed. | [optional] 
**CoveredProductLineEndDate** | Pointer to **string** | End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API. | [optional] 
**DeviceId** | Pointer to **string** | Unique identifier of the Cisco device. | [optional] 
**EndCustomer** | Pointer to [**NullableAssetCustomerInformation**](AssetCustomerInformation.md) |  | [optional] 
**EndUserGlobalUltimate** | Pointer to [**NullableAssetGlobalUltimate**](AssetGlobalUltimate.md) |  | [optional] 
**IsValid** | Pointer to **bool** | Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs. | [optional] 
**ItemType** | Pointer to **string** | Item type of this specific Cisco device. example \&quot;Chassis\&quot;. | [optional] 
**LastDateOfSupport** | Pointer to **time.Time** | The last date of hardware support for this device. | [optional] 
**MaintenancePurchaseOrderNumber** | Pointer to **string** | Maintenance purchase order number for the Cisco device. | [optional] 
**MaintenanceSalesOrderNumber** | Pointer to **string** | Maintenance sales order number for the Cisco device. | [optional] 
**Product** | Pointer to [**NullableAssetProductInformation**](AssetProductInformation.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number for the Cisco device. It is a unique number assigned for every purchase. | [optional] 
**ResellerGlobalUltimate** | Pointer to [**NullableAssetGlobalUltimate**](AssetGlobalUltimate.md) |  | [optional] 
**SalesOrderNumber** | Pointer to **string** | Sales order number for the Cisco device. It is a unique number assigned for every sale. | [optional] 
**ServiceDescription** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] 
**ServiceEndDate** | Pointer to **time.Time** | End date for the Cisco service contract that covers this Cisco device. | [optional] 
**ServiceLevel** | Pointer to **string** | The type of service contract that covers the Cisco device. | [optional] 
**ServiceSku** | Pointer to **string** | The SKU of the service contract that covers the Cisco device. | [optional] 
**ServiceStartDate** | Pointer to **time.Time** | Start date for the Cisco service contract that covers this Cisco device. | [optional] 
**StateContract** | Pointer to **string** | Internal property used for triggering and tracking actions for contract information. * &#x60;Update&#x60; - Sn2Info/Contract information needs to be updated. * &#x60;OK&#x60; - Sn2Info/Contract information was fetched succcessfuly and updated. * &#x60;Failed&#x60; - Sn2Info/Contract information was not available  or failed while fetching. * &#x60;Retry&#x60; - Sn2Info/Contract information update failed and will be retried later. | [optional] [default to "Update"]
**StateSn2Info** | Pointer to **string** | Internal property used for triggering and tracking actions for sn2info information. * &#x60;Update&#x60; - Sn2Info/Contract information needs to be updated. * &#x60;OK&#x60; - Sn2Info/Contract information was fetched succcessfuly and updated. * &#x60;Failed&#x60; - Sn2Info/Contract information was not available  or failed while fetching. * &#x60;Retry&#x60; - Sn2Info/Contract information update failed and will be retried later. | [optional] [default to "Update"]
**WarrantyEndDate** | Pointer to **string** | End date for the warranty that covers the Cisco device. | [optional] 
**WarrantyType** | Pointer to **string** | Type of warranty that covers the Cisco device. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAssetDeviceContractNotificationAllOf

`func NewAssetDeviceContractNotificationAllOf(classId string, objectType string, ) *AssetDeviceContractNotificationAllOf`

NewAssetDeviceContractNotificationAllOf instantiates a new AssetDeviceContractNotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceContractNotificationAllOfWithDefaults

`func NewAssetDeviceContractNotificationAllOfWithDefaults() *AssetDeviceContractNotificationAllOf`

NewAssetDeviceContractNotificationAllOfWithDefaults instantiates a new AssetDeviceContractNotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceContractNotificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceContractNotificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceContractNotificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceContractNotificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceContractNotificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceContractNotificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContract

`func (o *AssetDeviceContractNotificationAllOf) GetContract() AssetContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AssetDeviceContractNotificationAllOf) GetContractOk() (*AssetContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AssetDeviceContractNotificationAllOf) SetContract(v AssetContractInformation)`

SetContract sets Contract field to given value.

### HasContract

`func (o *AssetDeviceContractNotificationAllOf) HasContract() bool`

HasContract returns a boolean if a field has been set.

### SetContractNil

`func (o *AssetDeviceContractNotificationAllOf) SetContractNil(b bool)`

 SetContractNil sets the value for Contract to be an explicit nil

### UnsetContract
`func (o *AssetDeviceContractNotificationAllOf) UnsetContract()`

UnsetContract ensures that no value is present for Contract, not even an explicit nil
### GetContractStatus

`func (o *AssetDeviceContractNotificationAllOf) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *AssetDeviceContractNotificationAllOf) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *AssetDeviceContractNotificationAllOf) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *AssetDeviceContractNotificationAllOf) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractStatusReason

`func (o *AssetDeviceContractNotificationAllOf) GetContractStatusReason() string`

GetContractStatusReason returns the ContractStatusReason field if non-nil, zero value otherwise.

### GetContractStatusReasonOk

`func (o *AssetDeviceContractNotificationAllOf) GetContractStatusReasonOk() (*string, bool)`

GetContractStatusReasonOk returns a tuple with the ContractStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatusReason

`func (o *AssetDeviceContractNotificationAllOf) SetContractStatusReason(v string)`

SetContractStatusReason sets ContractStatusReason field to given value.

### HasContractStatusReason

`func (o *AssetDeviceContractNotificationAllOf) HasContractStatusReason() bool`

HasContractStatusReason returns a boolean if a field has been set.

### GetContractUpdatedTime

`func (o *AssetDeviceContractNotificationAllOf) GetContractUpdatedTime() time.Time`

GetContractUpdatedTime returns the ContractUpdatedTime field if non-nil, zero value otherwise.

### GetContractUpdatedTimeOk

`func (o *AssetDeviceContractNotificationAllOf) GetContractUpdatedTimeOk() (*time.Time, bool)`

GetContractUpdatedTimeOk returns a tuple with the ContractUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractUpdatedTime

`func (o *AssetDeviceContractNotificationAllOf) SetContractUpdatedTime(v time.Time)`

SetContractUpdatedTime sets ContractUpdatedTime field to given value.

### HasContractUpdatedTime

`func (o *AssetDeviceContractNotificationAllOf) HasContractUpdatedTime() bool`

HasContractUpdatedTime returns a boolean if a field has been set.

### GetCoveredProductLineEndDate

`func (o *AssetDeviceContractNotificationAllOf) GetCoveredProductLineEndDate() string`

GetCoveredProductLineEndDate returns the CoveredProductLineEndDate field if non-nil, zero value otherwise.

### GetCoveredProductLineEndDateOk

`func (o *AssetDeviceContractNotificationAllOf) GetCoveredProductLineEndDateOk() (*string, bool)`

GetCoveredProductLineEndDateOk returns a tuple with the CoveredProductLineEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredProductLineEndDate

`func (o *AssetDeviceContractNotificationAllOf) SetCoveredProductLineEndDate(v string)`

SetCoveredProductLineEndDate sets CoveredProductLineEndDate field to given value.

### HasCoveredProductLineEndDate

`func (o *AssetDeviceContractNotificationAllOf) HasCoveredProductLineEndDate() bool`

HasCoveredProductLineEndDate returns a boolean if a field has been set.

### GetDeviceId

`func (o *AssetDeviceContractNotificationAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AssetDeviceContractNotificationAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AssetDeviceContractNotificationAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AssetDeviceContractNotificationAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEndCustomer

`func (o *AssetDeviceContractNotificationAllOf) GetEndCustomer() AssetCustomerInformation`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *AssetDeviceContractNotificationAllOf) GetEndCustomerOk() (*AssetCustomerInformation, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *AssetDeviceContractNotificationAllOf) SetEndCustomer(v AssetCustomerInformation)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *AssetDeviceContractNotificationAllOf) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### SetEndCustomerNil

`func (o *AssetDeviceContractNotificationAllOf) SetEndCustomerNil(b bool)`

 SetEndCustomerNil sets the value for EndCustomer to be an explicit nil

### UnsetEndCustomer
`func (o *AssetDeviceContractNotificationAllOf) UnsetEndCustomer()`

UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
### GetEndUserGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) GetEndUserGlobalUltimate() AssetGlobalUltimate`

GetEndUserGlobalUltimate returns the EndUserGlobalUltimate field if non-nil, zero value otherwise.

### GetEndUserGlobalUltimateOk

`func (o *AssetDeviceContractNotificationAllOf) GetEndUserGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetEndUserGlobalUltimateOk returns a tuple with the EndUserGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) SetEndUserGlobalUltimate(v AssetGlobalUltimate)`

SetEndUserGlobalUltimate sets EndUserGlobalUltimate field to given value.

### HasEndUserGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) HasEndUserGlobalUltimate() bool`

HasEndUserGlobalUltimate returns a boolean if a field has been set.

### SetEndUserGlobalUltimateNil

`func (o *AssetDeviceContractNotificationAllOf) SetEndUserGlobalUltimateNil(b bool)`

 SetEndUserGlobalUltimateNil sets the value for EndUserGlobalUltimate to be an explicit nil

### UnsetEndUserGlobalUltimate
`func (o *AssetDeviceContractNotificationAllOf) UnsetEndUserGlobalUltimate()`

UnsetEndUserGlobalUltimate ensures that no value is present for EndUserGlobalUltimate, not even an explicit nil
### GetIsValid

`func (o *AssetDeviceContractNotificationAllOf) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AssetDeviceContractNotificationAllOf) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AssetDeviceContractNotificationAllOf) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AssetDeviceContractNotificationAllOf) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeviceContractNotificationAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeviceContractNotificationAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeviceContractNotificationAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeviceContractNotificationAllOf) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetLastDateOfSupport

`func (o *AssetDeviceContractNotificationAllOf) GetLastDateOfSupport() time.Time`

GetLastDateOfSupport returns the LastDateOfSupport field if non-nil, zero value otherwise.

### GetLastDateOfSupportOk

`func (o *AssetDeviceContractNotificationAllOf) GetLastDateOfSupportOk() (*time.Time, bool)`

GetLastDateOfSupportOk returns a tuple with the LastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateOfSupport

`func (o *AssetDeviceContractNotificationAllOf) SetLastDateOfSupport(v time.Time)`

SetLastDateOfSupport sets LastDateOfSupport field to given value.

### HasLastDateOfSupport

`func (o *AssetDeviceContractNotificationAllOf) HasLastDateOfSupport() bool`

HasLastDateOfSupport returns a boolean if a field has been set.

### GetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) GetMaintenancePurchaseOrderNumber() string`

GetMaintenancePurchaseOrderNumber returns the MaintenancePurchaseOrderNumber field if non-nil, zero value otherwise.

### GetMaintenancePurchaseOrderNumberOk

`func (o *AssetDeviceContractNotificationAllOf) GetMaintenancePurchaseOrderNumberOk() (*string, bool)`

GetMaintenancePurchaseOrderNumberOk returns a tuple with the MaintenancePurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) SetMaintenancePurchaseOrderNumber(v string)`

SetMaintenancePurchaseOrderNumber sets MaintenancePurchaseOrderNumber field to given value.

### HasMaintenancePurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) HasMaintenancePurchaseOrderNumber() bool`

HasMaintenancePurchaseOrderNumber returns a boolean if a field has been set.

### GetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) GetMaintenanceSalesOrderNumber() string`

GetMaintenanceSalesOrderNumber returns the MaintenanceSalesOrderNumber field if non-nil, zero value otherwise.

### GetMaintenanceSalesOrderNumberOk

`func (o *AssetDeviceContractNotificationAllOf) GetMaintenanceSalesOrderNumberOk() (*string, bool)`

GetMaintenanceSalesOrderNumberOk returns a tuple with the MaintenanceSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) SetMaintenanceSalesOrderNumber(v string)`

SetMaintenanceSalesOrderNumber sets MaintenanceSalesOrderNumber field to given value.

### HasMaintenanceSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) HasMaintenanceSalesOrderNumber() bool`

HasMaintenanceSalesOrderNumber returns a boolean if a field has been set.

### GetProduct

`func (o *AssetDeviceContractNotificationAllOf) GetProduct() AssetProductInformation`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AssetDeviceContractNotificationAllOf) GetProductOk() (*AssetProductInformation, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AssetDeviceContractNotificationAllOf) SetProduct(v AssetProductInformation)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AssetDeviceContractNotificationAllOf) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *AssetDeviceContractNotificationAllOf) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *AssetDeviceContractNotificationAllOf) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetPurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *AssetDeviceContractNotificationAllOf) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetResellerGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) GetResellerGlobalUltimate() AssetGlobalUltimate`

GetResellerGlobalUltimate returns the ResellerGlobalUltimate field if non-nil, zero value otherwise.

### GetResellerGlobalUltimateOk

`func (o *AssetDeviceContractNotificationAllOf) GetResellerGlobalUltimateOk() (*AssetGlobalUltimate, bool)`

GetResellerGlobalUltimateOk returns a tuple with the ResellerGlobalUltimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) SetResellerGlobalUltimate(v AssetGlobalUltimate)`

SetResellerGlobalUltimate sets ResellerGlobalUltimate field to given value.

### HasResellerGlobalUltimate

`func (o *AssetDeviceContractNotificationAllOf) HasResellerGlobalUltimate() bool`

HasResellerGlobalUltimate returns a boolean if a field has been set.

### SetResellerGlobalUltimateNil

`func (o *AssetDeviceContractNotificationAllOf) SetResellerGlobalUltimateNil(b bool)`

 SetResellerGlobalUltimateNil sets the value for ResellerGlobalUltimate to be an explicit nil

### UnsetResellerGlobalUltimate
`func (o *AssetDeviceContractNotificationAllOf) UnsetResellerGlobalUltimate()`

UnsetResellerGlobalUltimate ensures that no value is present for ResellerGlobalUltimate, not even an explicit nil
### GetSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) GetSalesOrderNumber() string`

GetSalesOrderNumber returns the SalesOrderNumber field if non-nil, zero value otherwise.

### GetSalesOrderNumberOk

`func (o *AssetDeviceContractNotificationAllOf) GetSalesOrderNumberOk() (*string, bool)`

GetSalesOrderNumberOk returns a tuple with the SalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) SetSalesOrderNumber(v string)`

SetSalesOrderNumber sets SalesOrderNumber field to given value.

### HasSalesOrderNumber

`func (o *AssetDeviceContractNotificationAllOf) HasSalesOrderNumber() bool`

HasSalesOrderNumber returns a boolean if a field has been set.

### GetServiceDescription

`func (o *AssetDeviceContractNotificationAllOf) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *AssetDeviceContractNotificationAllOf) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *AssetDeviceContractNotificationAllOf) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.

### HasServiceDescription

`func (o *AssetDeviceContractNotificationAllOf) HasServiceDescription() bool`

HasServiceDescription returns a boolean if a field has been set.

### GetServiceEndDate

`func (o *AssetDeviceContractNotificationAllOf) GetServiceEndDate() time.Time`

GetServiceEndDate returns the ServiceEndDate field if non-nil, zero value otherwise.

### GetServiceEndDateOk

`func (o *AssetDeviceContractNotificationAllOf) GetServiceEndDateOk() (*time.Time, bool)`

GetServiceEndDateOk returns a tuple with the ServiceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndDate

`func (o *AssetDeviceContractNotificationAllOf) SetServiceEndDate(v time.Time)`

SetServiceEndDate sets ServiceEndDate field to given value.

### HasServiceEndDate

`func (o *AssetDeviceContractNotificationAllOf) HasServiceEndDate() bool`

HasServiceEndDate returns a boolean if a field has been set.

### GetServiceLevel

`func (o *AssetDeviceContractNotificationAllOf) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *AssetDeviceContractNotificationAllOf) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *AssetDeviceContractNotificationAllOf) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *AssetDeviceContractNotificationAllOf) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetServiceSku

`func (o *AssetDeviceContractNotificationAllOf) GetServiceSku() string`

GetServiceSku returns the ServiceSku field if non-nil, zero value otherwise.

### GetServiceSkuOk

`func (o *AssetDeviceContractNotificationAllOf) GetServiceSkuOk() (*string, bool)`

GetServiceSkuOk returns a tuple with the ServiceSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSku

`func (o *AssetDeviceContractNotificationAllOf) SetServiceSku(v string)`

SetServiceSku sets ServiceSku field to given value.

### HasServiceSku

`func (o *AssetDeviceContractNotificationAllOf) HasServiceSku() bool`

HasServiceSku returns a boolean if a field has been set.

### GetServiceStartDate

`func (o *AssetDeviceContractNotificationAllOf) GetServiceStartDate() time.Time`

GetServiceStartDate returns the ServiceStartDate field if non-nil, zero value otherwise.

### GetServiceStartDateOk

`func (o *AssetDeviceContractNotificationAllOf) GetServiceStartDateOk() (*time.Time, bool)`

GetServiceStartDateOk returns a tuple with the ServiceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartDate

`func (o *AssetDeviceContractNotificationAllOf) SetServiceStartDate(v time.Time)`

SetServiceStartDate sets ServiceStartDate field to given value.

### HasServiceStartDate

`func (o *AssetDeviceContractNotificationAllOf) HasServiceStartDate() bool`

HasServiceStartDate returns a boolean if a field has been set.

### GetStateContract

`func (o *AssetDeviceContractNotificationAllOf) GetStateContract() string`

GetStateContract returns the StateContract field if non-nil, zero value otherwise.

### GetStateContractOk

`func (o *AssetDeviceContractNotificationAllOf) GetStateContractOk() (*string, bool)`

GetStateContractOk returns a tuple with the StateContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateContract

`func (o *AssetDeviceContractNotificationAllOf) SetStateContract(v string)`

SetStateContract sets StateContract field to given value.

### HasStateContract

`func (o *AssetDeviceContractNotificationAllOf) HasStateContract() bool`

HasStateContract returns a boolean if a field has been set.

### GetStateSn2Info

`func (o *AssetDeviceContractNotificationAllOf) GetStateSn2Info() string`

GetStateSn2Info returns the StateSn2Info field if non-nil, zero value otherwise.

### GetStateSn2InfoOk

`func (o *AssetDeviceContractNotificationAllOf) GetStateSn2InfoOk() (*string, bool)`

GetStateSn2InfoOk returns a tuple with the StateSn2Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateSn2Info

`func (o *AssetDeviceContractNotificationAllOf) SetStateSn2Info(v string)`

SetStateSn2Info sets StateSn2Info field to given value.

### HasStateSn2Info

`func (o *AssetDeviceContractNotificationAllOf) HasStateSn2Info() bool`

HasStateSn2Info returns a boolean if a field has been set.

### GetWarrantyEndDate

`func (o *AssetDeviceContractNotificationAllOf) GetWarrantyEndDate() string`

GetWarrantyEndDate returns the WarrantyEndDate field if non-nil, zero value otherwise.

### GetWarrantyEndDateOk

`func (o *AssetDeviceContractNotificationAllOf) GetWarrantyEndDateOk() (*string, bool)`

GetWarrantyEndDateOk returns a tuple with the WarrantyEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEndDate

`func (o *AssetDeviceContractNotificationAllOf) SetWarrantyEndDate(v string)`

SetWarrantyEndDate sets WarrantyEndDate field to given value.

### HasWarrantyEndDate

`func (o *AssetDeviceContractNotificationAllOf) HasWarrantyEndDate() bool`

HasWarrantyEndDate returns a boolean if a field has been set.

### GetWarrantyType

`func (o *AssetDeviceContractNotificationAllOf) GetWarrantyType() string`

GetWarrantyType returns the WarrantyType field if non-nil, zero value otherwise.

### GetWarrantyTypeOk

`func (o *AssetDeviceContractNotificationAllOf) GetWarrantyTypeOk() (*string, bool)`

GetWarrantyTypeOk returns a tuple with the WarrantyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyType

`func (o *AssetDeviceContractNotificationAllOf) SetWarrantyType(v string)`

SetWarrantyType sets WarrantyType field to given value.

### HasWarrantyType

`func (o *AssetDeviceContractNotificationAllOf) HasWarrantyType() bool`

HasWarrantyType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AssetDeviceContractNotificationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AssetDeviceContractNotificationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AssetDeviceContractNotificationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AssetDeviceContractNotificationAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


