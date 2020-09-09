# StoragePureHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostGroupName** | Pointer to **string** | Name of host group where the host belongs to. Empty if host is not part of any HostGroup. | [optional] [readonly] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**HostGroup** | Pointer to [**StoragePureHostGroupRelationship**](storage.PureHostGroup.Relationship.md) |  | [optional] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](storage.PureProtectionGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePureHost

`func NewStoragePureHost() *StoragePureHost`

NewStoragePureHost instantiates a new StoragePureHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureHostWithDefaults

`func NewStoragePureHostWithDefaults() *StoragePureHost`

NewStoragePureHostWithDefaults instantiates a new StoragePureHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostGroupName

`func (o *StoragePureHost) GetHostGroupName() string`

GetHostGroupName returns the HostGroupName field if non-nil, zero value otherwise.

### GetHostGroupNameOk

`func (o *StoragePureHost) GetHostGroupNameOk() (*string, bool)`

GetHostGroupNameOk returns a tuple with the HostGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupName

`func (o *StoragePureHost) SetHostGroupName(v string)`

SetHostGroupName sets HostGroupName field to given value.

### HasHostGroupName

`func (o *StoragePureHost) HasHostGroupName() bool`

HasHostGroupName returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureHost) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureHost) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureHost) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureHost) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHostGroup

`func (o *StoragePureHost) GetHostGroup() StoragePureHostGroupRelationship`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *StoragePureHost) GetHostGroupOk() (*StoragePureHostGroupRelationship, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *StoragePureHost) SetHostGroup(v StoragePureHostGroupRelationship)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *StoragePureHost) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.

### GetProtectionGroup

`func (o *StoragePureHost) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureHost) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureHost) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureHost) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureHost) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureHost) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureHost) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureHost) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


