# AssetDeviceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceInformation"]
**ApplicationName** | Pointer to **string** | Application name reported by Cisco Install Base. | [optional] [readonly] 
**DeviceTransactions** | Pointer to [**[]AssetDeviceTransaction**](AssetDeviceTransaction.md) |  | [optional] 
**EndCustomer** | Pointer to [**NullableAssetCustomerInformation**](AssetCustomerInformation.md) |  | [optional] 
**InstanceId** | Pointer to **string** | Instance number of the device. example \&quot;917280220\&quot;. | [optional] [readonly] 
**ItemType** | Pointer to **string** | Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity. | [optional] [readonly] 
**MlbOfferType** | Pointer to **string** | Identifier for consumption based pricing. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**MlbProductId** | Pointer to **int64** | Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**MlbProductName** | Pointer to **string** | Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**OldDeviceId** | Pointer to **string** | Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device. | [optional] [readonly] 
**OldDeviceStatusDescription** | Pointer to **string** | Description of status of old Cisco device, which got replaced by the new device. | [optional] [readonly] 
**OldDeviceStatusId** | Pointer to **int32** | Status ID of old Cisco device, which got replaced by the new device. * &#x60;0&#x60; - A default value to catch cases where device status is not correctly detected. * &#x60;10000&#x60; - Device is installed successfully. * &#x60;1010041&#x60; - Device is currently in Return Material Authorization process. * &#x60;10005&#x60; - Device is replaced successfully with another device. * &#x60;10007&#x60; - Device is returned succcessfuly. * &#x60;10009&#x60; - Device is terminated at customer&#39;s end. | [optional] [readonly] [default to 0]
**OldInstanceId** | Pointer to **string** | Instance number of the old device, which got replaced by the new device. | [optional] [readonly] 
**ProductFamily** | Pointer to **string** | Product Family is the field used to identify the hypervisor type. example \&quot;ESXi\&quot;. | [optional] [readonly] 
**ProductType** | Pointer to **string** | Product type helps to determine if device has to be billed using consumption metering. example \&quot;SERVER\&quot;. | [optional] [readonly] 
**UnitOfMeasure** | Pointer to **string** | Unit of Measure is flag used to identify the type of metric being pushed. example - \&quot;STORAGE\&quot; for hardware metrics , \&quot;VM\&quot; - for hypervisor metrics. * &#x60;None&#x60; - A default value to catch cases where unit of measure is not correctly detected. * &#x60;STORAGE&#x60; - The metric type of the device is a storage metric. * &#x60;NODE&#x60; - The metric type of the device is a hardware metric. * &#x60;VM&#x60; - The metric type of the device is a hypervisor metric. | [optional] [readonly] [default to "None"]

## Methods

### NewAssetDeviceInformation

`func NewAssetDeviceInformation(classId string, objectType string, ) *AssetDeviceInformation`

NewAssetDeviceInformation instantiates a new AssetDeviceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceInformationWithDefaults

`func NewAssetDeviceInformationWithDefaults() *AssetDeviceInformation`

NewAssetDeviceInformationWithDefaults instantiates a new AssetDeviceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceInformation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceInformation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceInformation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceInformation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceInformation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceInformation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApplicationName

`func (o *AssetDeviceInformation) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AssetDeviceInformation) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AssetDeviceInformation) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AssetDeviceInformation) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetDeviceTransactions

`func (o *AssetDeviceInformation) GetDeviceTransactions() []AssetDeviceTransaction`

GetDeviceTransactions returns the DeviceTransactions field if non-nil, zero value otherwise.

### GetDeviceTransactionsOk

`func (o *AssetDeviceInformation) GetDeviceTransactionsOk() (*[]AssetDeviceTransaction, bool)`

GetDeviceTransactionsOk returns a tuple with the DeviceTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTransactions

`func (o *AssetDeviceInformation) SetDeviceTransactions(v []AssetDeviceTransaction)`

SetDeviceTransactions sets DeviceTransactions field to given value.

### HasDeviceTransactions

`func (o *AssetDeviceInformation) HasDeviceTransactions() bool`

HasDeviceTransactions returns a boolean if a field has been set.

### SetDeviceTransactionsNil

`func (o *AssetDeviceInformation) SetDeviceTransactionsNil(b bool)`

 SetDeviceTransactionsNil sets the value for DeviceTransactions to be an explicit nil

### UnsetDeviceTransactions
`func (o *AssetDeviceInformation) UnsetDeviceTransactions()`

UnsetDeviceTransactions ensures that no value is present for DeviceTransactions, not even an explicit nil
### GetEndCustomer

`func (o *AssetDeviceInformation) GetEndCustomer() AssetCustomerInformation`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *AssetDeviceInformation) GetEndCustomerOk() (*AssetCustomerInformation, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *AssetDeviceInformation) SetEndCustomer(v AssetCustomerInformation)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *AssetDeviceInformation) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### SetEndCustomerNil

`func (o *AssetDeviceInformation) SetEndCustomerNil(b bool)`

 SetEndCustomerNil sets the value for EndCustomer to be an explicit nil

### UnsetEndCustomer
`func (o *AssetDeviceInformation) UnsetEndCustomer()`

UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
### GetInstanceId

`func (o *AssetDeviceInformation) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AssetDeviceInformation) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AssetDeviceInformation) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AssetDeviceInformation) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeviceInformation) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeviceInformation) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeviceInformation) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeviceInformation) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMlbOfferType

`func (o *AssetDeviceInformation) GetMlbOfferType() string`

GetMlbOfferType returns the MlbOfferType field if non-nil, zero value otherwise.

### GetMlbOfferTypeOk

`func (o *AssetDeviceInformation) GetMlbOfferTypeOk() (*string, bool)`

GetMlbOfferTypeOk returns a tuple with the MlbOfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbOfferType

`func (o *AssetDeviceInformation) SetMlbOfferType(v string)`

SetMlbOfferType sets MlbOfferType field to given value.

### HasMlbOfferType

`func (o *AssetDeviceInformation) HasMlbOfferType() bool`

HasMlbOfferType returns a boolean if a field has been set.

### GetMlbProductId

`func (o *AssetDeviceInformation) GetMlbProductId() int64`

GetMlbProductId returns the MlbProductId field if non-nil, zero value otherwise.

### GetMlbProductIdOk

`func (o *AssetDeviceInformation) GetMlbProductIdOk() (*int64, bool)`

GetMlbProductIdOk returns a tuple with the MlbProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductId

`func (o *AssetDeviceInformation) SetMlbProductId(v int64)`

SetMlbProductId sets MlbProductId field to given value.

### HasMlbProductId

`func (o *AssetDeviceInformation) HasMlbProductId() bool`

HasMlbProductId returns a boolean if a field has been set.

### GetMlbProductName

`func (o *AssetDeviceInformation) GetMlbProductName() string`

GetMlbProductName returns the MlbProductName field if non-nil, zero value otherwise.

### GetMlbProductNameOk

`func (o *AssetDeviceInformation) GetMlbProductNameOk() (*string, bool)`

GetMlbProductNameOk returns a tuple with the MlbProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductName

`func (o *AssetDeviceInformation) SetMlbProductName(v string)`

SetMlbProductName sets MlbProductName field to given value.

### HasMlbProductName

`func (o *AssetDeviceInformation) HasMlbProductName() bool`

HasMlbProductName returns a boolean if a field has been set.

### GetOldDeviceId

`func (o *AssetDeviceInformation) GetOldDeviceId() string`

GetOldDeviceId returns the OldDeviceId field if non-nil, zero value otherwise.

### GetOldDeviceIdOk

`func (o *AssetDeviceInformation) GetOldDeviceIdOk() (*string, bool)`

GetOldDeviceIdOk returns a tuple with the OldDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceId

`func (o *AssetDeviceInformation) SetOldDeviceId(v string)`

SetOldDeviceId sets OldDeviceId field to given value.

### HasOldDeviceId

`func (o *AssetDeviceInformation) HasOldDeviceId() bool`

HasOldDeviceId returns a boolean if a field has been set.

### GetOldDeviceStatusDescription

`func (o *AssetDeviceInformation) GetOldDeviceStatusDescription() string`

GetOldDeviceStatusDescription returns the OldDeviceStatusDescription field if non-nil, zero value otherwise.

### GetOldDeviceStatusDescriptionOk

`func (o *AssetDeviceInformation) GetOldDeviceStatusDescriptionOk() (*string, bool)`

GetOldDeviceStatusDescriptionOk returns a tuple with the OldDeviceStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusDescription

`func (o *AssetDeviceInformation) SetOldDeviceStatusDescription(v string)`

SetOldDeviceStatusDescription sets OldDeviceStatusDescription field to given value.

### HasOldDeviceStatusDescription

`func (o *AssetDeviceInformation) HasOldDeviceStatusDescription() bool`

HasOldDeviceStatusDescription returns a boolean if a field has been set.

### GetOldDeviceStatusId

`func (o *AssetDeviceInformation) GetOldDeviceStatusId() int32`

GetOldDeviceStatusId returns the OldDeviceStatusId field if non-nil, zero value otherwise.

### GetOldDeviceStatusIdOk

`func (o *AssetDeviceInformation) GetOldDeviceStatusIdOk() (*int32, bool)`

GetOldDeviceStatusIdOk returns a tuple with the OldDeviceStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusId

`func (o *AssetDeviceInformation) SetOldDeviceStatusId(v int32)`

SetOldDeviceStatusId sets OldDeviceStatusId field to given value.

### HasOldDeviceStatusId

`func (o *AssetDeviceInformation) HasOldDeviceStatusId() bool`

HasOldDeviceStatusId returns a boolean if a field has been set.

### GetOldInstanceId

`func (o *AssetDeviceInformation) GetOldInstanceId() string`

GetOldInstanceId returns the OldInstanceId field if non-nil, zero value otherwise.

### GetOldInstanceIdOk

`func (o *AssetDeviceInformation) GetOldInstanceIdOk() (*string, bool)`

GetOldInstanceIdOk returns a tuple with the OldInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldInstanceId

`func (o *AssetDeviceInformation) SetOldInstanceId(v string)`

SetOldInstanceId sets OldInstanceId field to given value.

### HasOldInstanceId

`func (o *AssetDeviceInformation) HasOldInstanceId() bool`

HasOldInstanceId returns a boolean if a field has been set.

### GetProductFamily

`func (o *AssetDeviceInformation) GetProductFamily() string`

GetProductFamily returns the ProductFamily field if non-nil, zero value otherwise.

### GetProductFamilyOk

`func (o *AssetDeviceInformation) GetProductFamilyOk() (*string, bool)`

GetProductFamilyOk returns a tuple with the ProductFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFamily

`func (o *AssetDeviceInformation) SetProductFamily(v string)`

SetProductFamily sets ProductFamily field to given value.

### HasProductFamily

`func (o *AssetDeviceInformation) HasProductFamily() bool`

HasProductFamily returns a boolean if a field has been set.

### GetProductType

`func (o *AssetDeviceInformation) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AssetDeviceInformation) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AssetDeviceInformation) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AssetDeviceInformation) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *AssetDeviceInformation) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *AssetDeviceInformation) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *AssetDeviceInformation) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *AssetDeviceInformation) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


