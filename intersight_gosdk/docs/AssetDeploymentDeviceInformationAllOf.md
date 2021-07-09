# AssetDeploymentDeviceInformationAllOf

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

### NewAssetDeploymentDeviceInformationAllOf

`func NewAssetDeploymentDeviceInformationAllOf(classId string, objectType string, ) *AssetDeploymentDeviceInformationAllOf`

NewAssetDeploymentDeviceInformationAllOf instantiates a new AssetDeploymentDeviceInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeploymentDeviceInformationAllOfWithDefaults

`func NewAssetDeploymentDeviceInformationAllOfWithDefaults() *AssetDeploymentDeviceInformationAllOf`

NewAssetDeploymentDeviceInformationAllOfWithDefaults instantiates a new AssetDeploymentDeviceInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeploymentDeviceInformationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeploymentDeviceInformationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeploymentDeviceInformationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeploymentDeviceInformationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApplicationName

`func (o *AssetDeploymentDeviceInformationAllOf) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AssetDeploymentDeviceInformationAllOf) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AssetDeploymentDeviceInformationAllOf) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *AssetDeploymentDeviceInformationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetDeploymentDeviceInformationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetDeploymentDeviceInformationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceTransactions

`func (o *AssetDeploymentDeviceInformationAllOf) GetDeviceTransactions() []AssetDeviceTransaction`

GetDeviceTransactions returns the DeviceTransactions field if non-nil, zero value otherwise.

### GetDeviceTransactionsOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetDeviceTransactionsOk() (*[]AssetDeviceTransaction, bool)`

GetDeviceTransactionsOk returns a tuple with the DeviceTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTransactions

`func (o *AssetDeploymentDeviceInformationAllOf) SetDeviceTransactions(v []AssetDeviceTransaction)`

SetDeviceTransactions sets DeviceTransactions field to given value.

### HasDeviceTransactions

`func (o *AssetDeploymentDeviceInformationAllOf) HasDeviceTransactions() bool`

HasDeviceTransactions returns a boolean if a field has been set.

### SetDeviceTransactionsNil

`func (o *AssetDeploymentDeviceInformationAllOf) SetDeviceTransactionsNil(b bool)`

 SetDeviceTransactionsNil sets the value for DeviceTransactions to be an explicit nil

### UnsetDeviceTransactions
`func (o *AssetDeploymentDeviceInformationAllOf) UnsetDeviceTransactions()`

UnsetDeviceTransactions ensures that no value is present for DeviceTransactions, not even an explicit nil
### GetInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetItemType

`func (o *AssetDeploymentDeviceInformationAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AssetDeploymentDeviceInformationAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *AssetDeploymentDeviceInformationAllOf) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMlbProductId

`func (o *AssetDeploymentDeviceInformationAllOf) GetMlbProductId() int64`

GetMlbProductId returns the MlbProductId field if non-nil, zero value otherwise.

### GetMlbProductIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetMlbProductIdOk() (*int64, bool)`

GetMlbProductIdOk returns a tuple with the MlbProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductId

`func (o *AssetDeploymentDeviceInformationAllOf) SetMlbProductId(v int64)`

SetMlbProductId sets MlbProductId field to given value.

### HasMlbProductId

`func (o *AssetDeploymentDeviceInformationAllOf) HasMlbProductId() bool`

HasMlbProductId returns a boolean if a field has been set.

### GetMlbProductName

`func (o *AssetDeploymentDeviceInformationAllOf) GetMlbProductName() string`

GetMlbProductName returns the MlbProductName field if non-nil, zero value otherwise.

### GetMlbProductNameOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetMlbProductNameOk() (*string, bool)`

GetMlbProductNameOk returns a tuple with the MlbProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlbProductName

`func (o *AssetDeploymentDeviceInformationAllOf) SetMlbProductName(v string)`

SetMlbProductName sets MlbProductName field to given value.

### HasMlbProductName

`func (o *AssetDeploymentDeviceInformationAllOf) HasMlbProductName() bool`

HasMlbProductName returns a boolean if a field has been set.

### GetOldDeviceId

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceId() string`

GetOldDeviceId returns the OldDeviceId field if non-nil, zero value otherwise.

### GetOldDeviceIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceIdOk() (*string, bool)`

GetOldDeviceIdOk returns a tuple with the OldDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceId

`func (o *AssetDeploymentDeviceInformationAllOf) SetOldDeviceId(v string)`

SetOldDeviceId sets OldDeviceId field to given value.

### HasOldDeviceId

`func (o *AssetDeploymentDeviceInformationAllOf) HasOldDeviceId() bool`

HasOldDeviceId returns a boolean if a field has been set.

### GetOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceStatusDescription() string`

GetOldDeviceStatusDescription returns the OldDeviceStatusDescription field if non-nil, zero value otherwise.

### GetOldDeviceStatusDescriptionOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceStatusDescriptionOk() (*string, bool)`

GetOldDeviceStatusDescriptionOk returns a tuple with the OldDeviceStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformationAllOf) SetOldDeviceStatusDescription(v string)`

SetOldDeviceStatusDescription sets OldDeviceStatusDescription field to given value.

### HasOldDeviceStatusDescription

`func (o *AssetDeploymentDeviceInformationAllOf) HasOldDeviceStatusDescription() bool`

HasOldDeviceStatusDescription returns a boolean if a field has been set.

### GetOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceStatusId() int32`

GetOldDeviceStatusId returns the OldDeviceStatusId field if non-nil, zero value otherwise.

### GetOldDeviceStatusIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldDeviceStatusIdOk() (*int32, bool)`

GetOldDeviceStatusIdOk returns a tuple with the OldDeviceStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformationAllOf) SetOldDeviceStatusId(v int32)`

SetOldDeviceStatusId sets OldDeviceStatusId field to given value.

### HasOldDeviceStatusId

`func (o *AssetDeploymentDeviceInformationAllOf) HasOldDeviceStatusId() bool`

HasOldDeviceStatusId returns a boolean if a field has been set.

### GetOldInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldInstanceId() string`

GetOldInstanceId returns the OldInstanceId field if non-nil, zero value otherwise.

### GetOldInstanceIdOk

`func (o *AssetDeploymentDeviceInformationAllOf) GetOldInstanceIdOk() (*string, bool)`

GetOldInstanceIdOk returns a tuple with the OldInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) SetOldInstanceId(v string)`

SetOldInstanceId sets OldInstanceId field to given value.

### HasOldInstanceId

`func (o *AssetDeploymentDeviceInformationAllOf) HasOldInstanceId() bool`

HasOldInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


