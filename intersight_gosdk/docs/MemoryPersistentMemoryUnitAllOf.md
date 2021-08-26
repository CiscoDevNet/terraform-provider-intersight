# MemoryPersistentMemoryUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryUnit"]
**AppDirectCapacity** | Pointer to **string** | AppDirect capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**CountStatus** | Pointer to **string** | Count status of the Persistent Memory Module on a server. | [optional] [readonly] 
**FirmwareVersion** | Pointer to **string** | Firmware version of the firware running on the Persistent Memory Module on a server. | [optional] [readonly] 
**FrozenStatus** | Pointer to **string** | Frozen status of the Persistent Memory Module on a server. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Module on a server. | [optional] [readonly] 
**LockStatus** | Pointer to **string** | Lock status of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryCapacity** | Pointer to **string** | Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**MemoryId** | Pointer to **int64** | ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**PersistentMemoryCapacity** | Pointer to **string** | Persistent Memory capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**ReservedCapacity** | Pointer to **string** | Reserved capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**SecurityStatus** | Pointer to **string** | Security status of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Module on a server. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Module on a server. | [optional] [readonly] 
**Uid** | Pointer to **string** | UID of the Persistent Memory Module on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryArray** | Pointer to [**MemoryArrayRelationship**](MemoryArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryUnitAllOf

`func NewMemoryPersistentMemoryUnitAllOf(classId string, objectType string, ) *MemoryPersistentMemoryUnitAllOf`

NewMemoryPersistentMemoryUnitAllOf instantiates a new MemoryPersistentMemoryUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryUnitAllOfWithDefaults

`func NewMemoryPersistentMemoryUnitAllOfWithDefaults() *MemoryPersistentMemoryUnitAllOf`

NewMemoryPersistentMemoryUnitAllOfWithDefaults instantiates a new MemoryPersistentMemoryUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) GetAppDirectCapacity() string`

GetAppDirectCapacity returns the AppDirectCapacity field if non-nil, zero value otherwise.

### GetAppDirectCapacityOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetAppDirectCapacityOk() (*string, bool)`

GetAppDirectCapacityOk returns a tuple with the AppDirectCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) SetAppDirectCapacity(v string)`

SetAppDirectCapacity sets AppDirectCapacity field to given value.

### HasAppDirectCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) HasAppDirectCapacity() bool`

HasAppDirectCapacity returns a boolean if a field has been set.

### GetCountStatus

`func (o *MemoryPersistentMemoryUnitAllOf) GetCountStatus() string`

GetCountStatus returns the CountStatus field if non-nil, zero value otherwise.

### GetCountStatusOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetCountStatusOk() (*string, bool)`

GetCountStatusOk returns a tuple with the CountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountStatus

`func (o *MemoryPersistentMemoryUnitAllOf) SetCountStatus(v string)`

SetCountStatus sets CountStatus field to given value.

### HasCountStatus

`func (o *MemoryPersistentMemoryUnitAllOf) HasCountStatus() bool`

HasCountStatus returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *MemoryPersistentMemoryUnitAllOf) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *MemoryPersistentMemoryUnitAllOf) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *MemoryPersistentMemoryUnitAllOf) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFrozenStatus

`func (o *MemoryPersistentMemoryUnitAllOf) GetFrozenStatus() string`

GetFrozenStatus returns the FrozenStatus field if non-nil, zero value otherwise.

### GetFrozenStatusOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetFrozenStatusOk() (*string, bool)`

GetFrozenStatusOk returns a tuple with the FrozenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenStatus

`func (o *MemoryPersistentMemoryUnitAllOf) SetFrozenStatus(v string)`

SetFrozenStatus sets FrozenStatus field to given value.

### HasFrozenStatus

`func (o *MemoryPersistentMemoryUnitAllOf) HasFrozenStatus() bool`

HasFrozenStatus returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryUnitAllOf) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryUnitAllOf) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryUnitAllOf) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetLockStatus

`func (o *MemoryPersistentMemoryUnitAllOf) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *MemoryPersistentMemoryUnitAllOf) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *MemoryPersistentMemoryUnitAllOf) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryCapacity() string`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryCapacityOk() (*string, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) SetMemoryCapacity(v string)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryId() int64`

GetMemoryId returns the MemoryId field if non-nil, zero value otherwise.

### GetMemoryIdOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryIdOk() (*int64, bool)`

GetMemoryIdOk returns a tuple with the MemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) SetMemoryId(v int64)`

SetMemoryId sets MemoryId field to given value.

### HasMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) HasMemoryId() bool`

HasMemoryId returns a boolean if a field has been set.

### GetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) GetPersistentMemoryCapacity() string`

GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field if non-nil, zero value otherwise.

### GetPersistentMemoryCapacityOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetPersistentMemoryCapacityOk() (*string, bool)`

GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) SetPersistentMemoryCapacity(v string)`

SetPersistentMemoryCapacity sets PersistentMemoryCapacity field to given value.

### HasPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) HasPersistentMemoryCapacity() bool`

HasPersistentMemoryCapacity returns a boolean if a field has been set.

### GetReservedCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) GetReservedCapacity() string`

GetReservedCapacity returns the ReservedCapacity field if non-nil, zero value otherwise.

### GetReservedCapacityOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetReservedCapacityOk() (*string, bool)`

GetReservedCapacityOk returns a tuple with the ReservedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) SetReservedCapacity(v string)`

SetReservedCapacity sets ReservedCapacity field to given value.

### HasReservedCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) HasReservedCapacity() bool`

HasReservedCapacity returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *MemoryPersistentMemoryUnitAllOf) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *MemoryPersistentMemoryUnitAllOf) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *MemoryPersistentMemoryUnitAllOf) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryUnitAllOf) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryUnitAllOf) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryUnitAllOf) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryUnitAllOf) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryUnitAllOf) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetUid

`func (o *MemoryPersistentMemoryUnitAllOf) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MemoryPersistentMemoryUnitAllOf) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MemoryPersistentMemoryUnitAllOf) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryUnitAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryArray

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryArray() MemoryArrayRelationship`

GetMemoryArray returns the MemoryArray field if non-nil, zero value otherwise.

### GetMemoryArrayOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetMemoryArrayOk() (*MemoryArrayRelationship, bool)`

GetMemoryArrayOk returns a tuple with the MemoryArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryArray

`func (o *MemoryPersistentMemoryUnitAllOf) SetMemoryArray(v MemoryArrayRelationship)`

SetMemoryArray sets MemoryArray field to given value.

### HasMemoryArray

`func (o *MemoryPersistentMemoryUnitAllOf) HasMemoryArray() bool`

HasMemoryArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryPersistentMemoryUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


