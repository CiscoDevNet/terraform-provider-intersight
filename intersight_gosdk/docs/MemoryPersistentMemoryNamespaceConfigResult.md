# MemoryPersistentMemoryNamespaceConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigStatus** | Pointer to **string** | Status of the Persistent Memory Namespace needed to be configured. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of a Persistent Memory Namespace that needed to be configured. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID in which the Persistent Memory Namespace needed to be configured. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID in which the Persistent Memory Namespace needed to be configured. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryConfigResult** | Pointer to [**MemoryPersistentMemoryConfigResultRelationship**](memory.PersistentMemoryConfigResult.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryNamespaceConfigResult

`func NewMemoryPersistentMemoryNamespaceConfigResult() *MemoryPersistentMemoryNamespaceConfigResult`

NewMemoryPersistentMemoryNamespaceConfigResult instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults

`func NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults() *MemoryPersistentMemoryNamespaceConfigResult`

NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigStatus

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatus() string`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatusOk() (*string, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetConfigStatus(v string)`

SetConfigStatus sets ConfigStatus field to given value.

### HasConfigStatus

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasConfigStatus() bool`

HasConfigStatus returns a boolean if a field has been set.

### GetName

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResult() MemoryPersistentMemoryConfigResultRelationship`

GetMemoryPersistentMemoryConfigResult returns the MemoryPersistentMemoryConfigResult field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigResultOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResultOk() (*MemoryPersistentMemoryConfigResultRelationship, bool)`

GetMemoryPersistentMemoryConfigResultOk returns a tuple with the MemoryPersistentMemoryConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetMemoryPersistentMemoryConfigResult(v MemoryPersistentMemoryConfigResultRelationship)`

SetMemoryPersistentMemoryConfigResult sets MemoryPersistentMemoryConfigResult field to given value.

### HasMemoryPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasMemoryPersistentMemoryConfigResult() bool`

HasMemoryPersistentMemoryConfigResult returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryNamespaceConfigResult) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


