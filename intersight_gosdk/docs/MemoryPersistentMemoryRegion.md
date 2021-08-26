# MemoryPersistentMemoryRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PersistentMemoryRegion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PersistentMemoryRegion"]
**FreeCapacity** | Pointer to **string** | Free capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**HealthState** | Pointer to **string** | Health state of the Persistent Memory Region. | [optional] [readonly] 
**InterleavedSetId** | Pointer to **string** | ID of the Interleaved Set formed for this Persistent Memory Region. | [optional] [readonly] 
**LocaterIds** | Pointer to **string** | Set of locator IDs that are included in the Persistent Memory Region. | [optional] [readonly] 
**PersistentMemoryType** | Pointer to **string** | Persistent Memory type of the Persistent Memory Region. | [optional] [readonly] 
**RegionId** | Pointer to **string** | ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketId** | Pointer to **string** | Socket ID of the Persistent Memory Region. | [optional] [readonly] 
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the Persistent Memory Region. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of the Persistent Memory Region. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**MemoryPersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](MemoryPersistentMemoryConfigurationRelationship.md) |  | [optional] 
**PersistentMemoryNamespaces** | Pointer to [**[]MemoryPersistentMemoryNamespaceRelationship**](MemoryPersistentMemoryNamespaceRelationship.md) | An array of relationships to memoryPersistentMemoryNamespace resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryRegion

`func NewMemoryPersistentMemoryRegion(classId string, objectType string, ) *MemoryPersistentMemoryRegion`

NewMemoryPersistentMemoryRegion instantiates a new MemoryPersistentMemoryRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryRegionWithDefaults

`func NewMemoryPersistentMemoryRegionWithDefaults() *MemoryPersistentMemoryRegion`

NewMemoryPersistentMemoryRegionWithDefaults instantiates a new MemoryPersistentMemoryRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPersistentMemoryRegion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPersistentMemoryRegion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPersistentMemoryRegion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPersistentMemoryRegion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPersistentMemoryRegion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPersistentMemoryRegion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFreeCapacity

`func (o *MemoryPersistentMemoryRegion) GetFreeCapacity() string`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *MemoryPersistentMemoryRegion) GetFreeCapacityOk() (*string, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *MemoryPersistentMemoryRegion) SetFreeCapacity(v string)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *MemoryPersistentMemoryRegion) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetHealthState

`func (o *MemoryPersistentMemoryRegion) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemoryPersistentMemoryRegion) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemoryPersistentMemoryRegion) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *MemoryPersistentMemoryRegion) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetInterleavedSetId

`func (o *MemoryPersistentMemoryRegion) GetInterleavedSetId() string`

GetInterleavedSetId returns the InterleavedSetId field if non-nil, zero value otherwise.

### GetInterleavedSetIdOk

`func (o *MemoryPersistentMemoryRegion) GetInterleavedSetIdOk() (*string, bool)`

GetInterleavedSetIdOk returns a tuple with the InterleavedSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedSetId

`func (o *MemoryPersistentMemoryRegion) SetInterleavedSetId(v string)`

SetInterleavedSetId sets InterleavedSetId field to given value.

### HasInterleavedSetId

`func (o *MemoryPersistentMemoryRegion) HasInterleavedSetId() bool`

HasInterleavedSetId returns a boolean if a field has been set.

### GetLocaterIds

`func (o *MemoryPersistentMemoryRegion) GetLocaterIds() string`

GetLocaterIds returns the LocaterIds field if non-nil, zero value otherwise.

### GetLocaterIdsOk

`func (o *MemoryPersistentMemoryRegion) GetLocaterIdsOk() (*string, bool)`

GetLocaterIdsOk returns a tuple with the LocaterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaterIds

`func (o *MemoryPersistentMemoryRegion) SetLocaterIds(v string)`

SetLocaterIds sets LocaterIds field to given value.

### HasLocaterIds

`func (o *MemoryPersistentMemoryRegion) HasLocaterIds() bool`

HasLocaterIds returns a boolean if a field has been set.

### GetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryType() string`

GetPersistentMemoryType returns the PersistentMemoryType field if non-nil, zero value otherwise.

### GetPersistentMemoryTypeOk

`func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryTypeOk() (*string, bool)`

GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryType

`func (o *MemoryPersistentMemoryRegion) SetPersistentMemoryType(v string)`

SetPersistentMemoryType sets PersistentMemoryType field to given value.

### HasPersistentMemoryType

`func (o *MemoryPersistentMemoryRegion) HasPersistentMemoryType() bool`

HasPersistentMemoryType returns a boolean if a field has been set.

### GetRegionId

`func (o *MemoryPersistentMemoryRegion) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *MemoryPersistentMemoryRegion) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *MemoryPersistentMemoryRegion) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *MemoryPersistentMemoryRegion) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryRegion) GetSocketId() string`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryRegion) GetSocketIdOk() (*string, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryRegion) SetSocketId(v string)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryRegion) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryRegion) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryRegion) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryRegion) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryRegion) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryRegion) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryRegion) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryRegion) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryRegion) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegion) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryRegion) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegion) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryRegion) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegion) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryRegion) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegion) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryRegion) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryNamespaces() []MemoryPersistentMemoryNamespaceRelationship`

GetPersistentMemoryNamespaces returns the PersistentMemoryNamespaces field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespacesOk

`func (o *MemoryPersistentMemoryRegion) GetPersistentMemoryNamespacesOk() (*[]MemoryPersistentMemoryNamespaceRelationship, bool)`

GetPersistentMemoryNamespacesOk returns a tuple with the PersistentMemoryNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegion) SetPersistentMemoryNamespaces(v []MemoryPersistentMemoryNamespaceRelationship)`

SetPersistentMemoryNamespaces sets PersistentMemoryNamespaces field to given value.

### HasPersistentMemoryNamespaces

`func (o *MemoryPersistentMemoryRegion) HasPersistentMemoryNamespaces() bool`

HasPersistentMemoryNamespaces returns a boolean if a field has been set.

### SetPersistentMemoryNamespacesNil

`func (o *MemoryPersistentMemoryRegion) SetPersistentMemoryNamespacesNil(b bool)`

 SetPersistentMemoryNamespacesNil sets the value for PersistentMemoryNamespaces to be an explicit nil

### UnsetPersistentMemoryNamespaces
`func (o *MemoryPersistentMemoryRegion) UnsetPersistentMemoryNamespaces()`

UnsetPersistentMemoryNamespaces ensures that no value is present for PersistentMemoryNamespaces, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryRegion) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryRegion) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryRegion) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryRegion) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


