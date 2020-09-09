# MemoryPersistentMemoryRegionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeCapacity** | Pointer to **string** | Free capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Region. | [optional] [readonly] 
**InterleavedSetId** | Pointer to **string** | ID of the Interleaved Set formed for this Persistent Memory Region. | [optional] [readonly] 
**LocaterIds** | Pointer to **string** | Set of locator IDs that are included in the Persistent Memory Region. | [optional] [readonly] 
**PersistentMemoryType** | Pointer to **string** | Persistent Memory type of the Persistent Memory Region. | [optional] [readonly] 
**RegionId** | Pointer to **string** | ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Region. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**PersistentMemoryNamespaces** | Pointer to [**[]MemoryPersistentMemoryNamespaceRelationship**](memory.PersistentMemoryNamespace.Relationship.md) | An array of relationships to memoryPersistentMemoryNamespace resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryRegionAllOf

`func NewMemoryPersistentMemoryRegionAllOf() *MemoryPersistentMemoryRegionAllOf`

NewMemoryPersistentMemoryRegionAllOf instantiates a new MemoryPersistentMemoryRegionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryRegionAllOfWithDefaults

`func NewMemoryPersistentMemoryRegionAllOfWithDefaults() *MemoryPersistentMemoryRegionAllOf`

NewMemoryPersistentMemoryRegionAllOfWithDefaults instantiates a new MemoryPersistentMemoryRegionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) GetFreeCapacity() string`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetFreeCapacityOk() (*string, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) SetFreeCapacity(v string)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryRegionAllOf) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryRegionAllOf) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryRegionAllOf) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetInterleavedSetId

`func (o *MemoryPersistentMemoryRegionAllOf) GetInterleavedSetId() string`

GetInterleavedSetId returns the InterleavedSetId field if non-nil, zero value otherwise.

### GetInterleavedSetIdOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetInterleavedSetIdOk() (*string, bool)`

GetInterleavedSetIdOk returns a tuple with the InterleavedSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedSetId

`func (o *MemoryPersistentMemoryRegionAllOf) SetInterleavedSetId(v string)`

SetInterleavedSetId sets InterleavedSetId field to given value.

### HasInterleavedSetId

`func (o *MemoryPersistentMemoryRegionAllOf) HasInterleavedSetId() bool`

HasInterleavedSetId returns a boolean if a field has been set.

### GetLocaterIds

`func (o *MemoryPersistentMemoryRegionAllOf) GetLocaterIds() string`

GetLocaterIds returns the LocaterIds field if non-nil, zero value otherwise.

### GetLocaterIdsOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetLocaterIdsOk() (*string, bool)`

GetLocaterIdsOk returns a tuple with the LocaterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaterIds

`func (o *MemoryPersistentMemoryRegionAllOf) SetLocaterIds(v string)`

SetLocaterIds sets LocaterIds field to given value.

### HasLocaterIds

`func (o *MemoryPersistentMemoryRegionAllOf) HasLocaterIds() bool`

HasLocaterIds returns a boolean if a field has been set.

### GetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryType() string`

GetPersistentMemoryType returns the PersistentMemoryType field if non-nil, zero value otherwise.

### GetPersistentMemoryTypeOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryTypeOk() (*string, bool)`

GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionAllOf) SetPersistentMemoryType(v string)`

SetPersistentMemoryType sets PersistentMemoryType field to given value.

### HasPersistentMemoryType

`func (o *MemoryPersistentMemoryRegionAllOf) HasPersistentMemoryType() bool`

HasPersistentMemoryType returns a boolean if a field has been set.

### GetRegionId

`func (o *MemoryPersistentMemoryRegionAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *MemoryPersistentMemoryRegionAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *MemoryPersistentMemoryRegionAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryRegionAllOf) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryRegionAllOf) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryRegionAllOf) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryRegionAllOf) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryRegionAllOf) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryRegionAllOf) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryRegionAllOf) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegionAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionAllOf) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionAllOf) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegionAllOf) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryNamespaces() []MemoryPersistentMemoryNamespaceRelationship`

GetPersistentMemoryNamespaces returns the PersistentMemoryNamespaces field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespacesOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryNamespacesOk() (*[]MemoryPersistentMemoryNamespaceRelationship, bool)`

GetPersistentMemoryNamespacesOk returns a tuple with the PersistentMemoryNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionAllOf) SetPersistentMemoryNamespaces(v []MemoryPersistentMemoryNamespaceRelationship)`

SetPersistentMemoryNamespaces sets PersistentMemoryNamespaces field to given value.

### HasPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegionAllOf) HasPersistentMemoryNamespaces() bool`

HasPersistentMemoryNamespaces returns a boolean if a field has been set.

### SetPersistentMemoryNamespacesNil

`func (o *MemoryPersistentMemoryRegionAllOf) SetPersistentMemoryNamespacesNil(b bool)`

 SetPersistentMemoryNamespacesNil sets the value for PersistentMemoryNamespaces to be an explicit nil

### UnsetPersistentMemoryNamespaces
`func (o *MemoryPersistentMemoryRegionAllOf) UnsetPersistentMemoryNamespaces()`

UnsetPersistentMemoryNamespaces ensures that no value is present for PersistentMemoryNamespaces, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryRegionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryRegionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryRegionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryRegionAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


