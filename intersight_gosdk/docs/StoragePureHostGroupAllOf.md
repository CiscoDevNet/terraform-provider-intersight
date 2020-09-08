# StoragePureHostGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostNames** | Pointer to **[]string** |  | [optional] 
**Array** | Pointer to [**StoragePureArrayRelationship**](storage.PureArray.Relationship.md) |  | [optional] 
**Hosts** | Pointer to [**[]StoragePureHostRelationship**](storage.PureHost.Relationship.md) | An array of relationships to storagePureHost resources. | [optional] [readonly] 
**ProtectionGroup** | Pointer to [**StoragePureProtectionGroupRelationship**](storage.PureProtectionGroup.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePureHostGroupAllOf

`func NewStoragePureHostGroupAllOf() *StoragePureHostGroupAllOf`

NewStoragePureHostGroupAllOf instantiates a new StoragePureHostGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureHostGroupAllOfWithDefaults

`func NewStoragePureHostGroupAllOfWithDefaults() *StoragePureHostGroupAllOf`

NewStoragePureHostGroupAllOfWithDefaults instantiates a new StoragePureHostGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostNames

`func (o *StoragePureHostGroupAllOf) GetHostNames() []string`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *StoragePureHostGroupAllOf) GetHostNamesOk() (*[]string, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *StoragePureHostGroupAllOf) SetHostNames(v []string)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *StoragePureHostGroupAllOf) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureHostGroupAllOf) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureHostGroupAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureHostGroupAllOf) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureHostGroupAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetHosts

`func (o *StoragePureHostGroupAllOf) GetHosts() []StoragePureHostRelationship`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StoragePureHostGroupAllOf) GetHostsOk() (*[]StoragePureHostRelationship, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StoragePureHostGroupAllOf) SetHosts(v []StoragePureHostRelationship)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StoragePureHostGroupAllOf) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *StoragePureHostGroupAllOf) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *StoragePureHostGroupAllOf) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetProtectionGroup

`func (o *StoragePureHostGroupAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship`

GetProtectionGroup returns the ProtectionGroup field if non-nil, zero value otherwise.

### GetProtectionGroupOk

`func (o *StoragePureHostGroupAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool)`

GetProtectionGroupOk returns a tuple with the ProtectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroup

`func (o *StoragePureHostGroupAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship)`

SetProtectionGroup sets ProtectionGroup field to given value.

### HasProtectionGroup

`func (o *StoragePureHostGroupAllOf) HasProtectionGroup() bool`

HasProtectionGroup returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePureHostGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureHostGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureHostGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureHostGroupAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


