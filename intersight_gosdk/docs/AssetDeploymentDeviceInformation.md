# AssetDeploymentDeviceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeploymentDeviceInformation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeploymentDeviceInformation"]
**ApplicationName** | Pointer to **string** | Application name reported by Cisco Install Base. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of device reported by Cisco Install Base. | [optional] [readonly] 
**DeviceTransactions** | Pointer to [**[]AssetDeviceTransaction**](AssetDeviceTransaction.md) |  | [optional] 
**InstanceId** | Pointer to **string** | Instance number of the device. example \&quot;917280220\&quot;. | [optional] [readonly] 
**ItemType** | Pointer to **string** | Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity. | [optional] [readonly] 
**MlbProductId** | Pointer to **int64** | Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**MlbProductName** | Pointer to **string** | Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle. | [optional] [readonly] 
**OldDeviceId** | Pointer to **string** | Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device. | [optional] [readonly] 
**OldDeviceStatusDescription** | Pointer to **string** | Description of status of old Cisco device, which got replaced by the new device. | [optional] [readonly] 
**OldDeviceStatusId** | Pointer to **int32** | Status ID of old Cisco device, which got replaced by the new device. * &#x60;0&#x60; - A default value to catch cases where device status is not correctly detected. * &#x60;10000&#x60; - Device is installed successfully. * &#x60;1010041&#x60; - Device is currently in Return Material Authorization process. * &#x60;10005&#x60; - Device is replaced successfully with another device. * &#x60;10007&#x60; - Device is returned succcessfuly. * &#x60;10009&#x60; - Device is terminated at customer&#39;s end. | [optional] [readonly] [default to 0]
**OldInstanceId** | Pointer to **string** | Instance number of the old device, which got replaced by the new device. | [optional] [readonly] 

## Methods

### NewAssetDeploymentDeviceInformation

`func NewAssetDeploymentDeviceInformation(classId string, objectType string, ) *AssetDeploymentDeviceInformation`

NewAssetDeploymentDeviceInformation instantiates a new AssetDeploymentDeviceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentDeviceInformationWithDefaults

`func NewAssetDeploymentDeviceInformationWithDefaults() *AssetDeploymentDeviceInformation`

NewAssetDeploymentDeviceInformationWithDefaults instantiates a new AssetDeploymentDeviceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeploymentDeviceInformation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeploymentDeviceInformation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeploymentDeviceInformation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeploymentDeviceInformation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeploymentDeviceInformation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeploymentDeviceInformation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApplicationName

`func (o *AssetDeploymentDeviceInformation) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AssetDeploymentDeviceInformation) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AssetDeploymentDeviceInformation) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AssetDeploymentDeviceInformation) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *AssetDeploymentDeviceInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetDeploymentDeviceInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetDeploymentDeviceInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetDeploymentDeviceInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceTransactions

`func (o *AssetDeploymentDeviceInformation) GetDeviceTransactions() []AssetDeviceTransaction`

GetDeviceTransactions returns the DeviceTransactions field if non-nil, zero value otherwise.

### GetDeviceTransactionsOk

`func (o *AssetDeploymentDeviceInformation) GetDeviceTransactionsOk() (*[]AssetDeviceTransaction, bool)`

GetDeviceTransactionsOk returns a tuple with the DeviceTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTransactions

`func (o *AssetDeploymentDeviceInformation) SetDeviceTransactions(v []AssetDeviceTransaction)`

SetDeviceTransactions sets DeviceTransactions field to given value.

### HasDeviceTransactions

`func (o *AssetDeploymentDeviceInformation) HasDeviceTransactions() bool`

HasDeviceTransactions returns a boolean if a field has been set.

### SetDeviceTransactionsNil

`func (o *AssetDeploymentDeviceInformation) SetDeviceTransactionsNil(b bool)`

 SetDeviceTransactionsNil sets the value for DeviceTransactions to be an explicit nil

### UnsetDeviceTransactions
`func (o *AssetDeploymentDeviceInformation) UnsetDeviceTransactions()`

UnsetDeviceTransactions ensures that no value is present for DeviceTransactions, not even an explicit nil
### GetInstanceId

`func (o *AssetDeploymentDeviceInformation) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AssetDeploymentDeviceInformation) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AssetDeploymentDeviceInformation) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AssetDeploymentDeviceInformation) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeploymentDeviceInformation) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeploymentDeviceInformation) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeploymentDeviceInformation) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeploymentDeviceInformation) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMlbProductId

`func (o *AssetDeploymentDeviceInformation) GetMlbProductId() int64`

GetMlbProductId returns the MlbProductId field if non-nil, zero value otherwise.

### GetMlbProductIdOk

`func (o *AssetDeploymentDeviceInformation) GetMlbProductIdOk() (*int64, bool)`

GetMlbProductIdOk returns a tuple with the MlbProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductId

`func (o *AssetDeploymentDeviceInformation) SetMlbProductId(v int64)`

SetMlbProductId sets MlbProductId field to given value.

### HasMlbProductId

`func (o *AssetDeploymentDeviceInformation) HasMlbProductId() bool`

HasMlbProductId returns a boolean if a field has been set.

### GetMlbProductName

`func (o *AssetDeploymentDeviceInformation) GetMlbProductName() string`

GetMlbProductName returns the MlbProductName field if non-nil, zero value otherwise.

### GetMlbProductNameOk

`func (o *AssetDeploymentDeviceInformation) GetMlbProductNameOk() (*string, bool)`

GetMlbProductNameOk returns a tuple with the MlbProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductName

`func (o *AssetDeploymentDeviceInformation) SetMlbProductName(v string)`

SetMlbProductName sets MlbProductName field to given value.

### HasMlbProductName

`func (o *AssetDeploymentDeviceInformation) HasMlbProductName() bool`

HasMlbProductName returns a boolean if a field has been set.

### GetOldDeviceId

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceId() string`

GetOldDeviceId returns the OldDeviceId field if non-nil, zero value otherwise.

### GetOldDeviceIdOk

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceIdOk() (*string, bool)`

GetOldDeviceIdOk returns a tuple with the OldDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceId

`func (o *AssetDeploymentDeviceInformation) SetOldDeviceId(v string)`

SetOldDeviceId sets OldDeviceId field to given value.

### HasOldDeviceId

`func (o *AssetDeploymentDeviceInformation) HasOldDeviceId() bool`

HasOldDeviceId returns a boolean if a field has been set.

### GetOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceStatusDescription() string`

GetOldDeviceStatusDescription returns the OldDeviceStatusDescription field if non-nil, zero value otherwise.

### GetOldDeviceStatusDescriptionOk

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceStatusDescriptionOk() (*string, bool)`

GetOldDeviceStatusDescriptionOk returns a tuple with the OldDeviceStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformation) SetOldDeviceStatusDescription(v string)`

SetOldDeviceStatusDescription sets OldDeviceStatusDescription field to given value.

### HasOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformation) HasOldDeviceStatusDescription() bool`

HasOldDeviceStatusDescription returns a boolean if a field has been set.

### GetOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceStatusId() int32`

GetOldDeviceStatusId returns the OldDeviceStatusId field if non-nil, zero value otherwise.

### GetOldDeviceStatusIdOk

`func (o *AssetDeploymentDeviceInformation) GetOldDeviceStatusIdOk() (*int32, bool)`

GetOldDeviceStatusIdOk returns a tuple with the OldDeviceStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformation) SetOldDeviceStatusId(v int32)`

SetOldDeviceStatusId sets OldDeviceStatusId field to given value.

### HasOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformation) HasOldDeviceStatusId() bool`

HasOldDeviceStatusId returns a boolean if a field has been set.

### GetOldInstanceId

`func (o *AssetDeploymentDeviceInformation) GetOldInstanceId() string`

GetOldInstanceId returns the OldInstanceId field if non-nil, zero value otherwise.

### GetOldInstanceIdOk

`func (o *AssetDeploymentDeviceInformation) GetOldInstanceIdOk() (*string, bool)`

GetOldInstanceIdOk returns a tuple with the OldInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldInstanceId

`func (o *AssetDeploymentDeviceInformation) SetOldInstanceId(v string)`

SetOldInstanceId sets OldInstanceId field to given value.

### HasOldInstanceId

`func (o *AssetDeploymentDeviceInformation) HasOldInstanceId() bool`

HasOldInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


