# MemoryPersistentMemoryConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigErrorDesc** | Pointer to **string** | Error in the result of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigResult** | Pointer to **string** | Result of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigSequenceNo** | Pointer to **int64** | Sequence number of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | State of a previously applied Persistent Memory configuration on a server. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**MemoryPersistentMemoryConfiguration** | Pointer to [**MemoryPersistentMemoryConfigurationRelationship**](memory.PersistentMemoryConfiguration.Relationship.md) |  | [optional] 
**PersistentMemoryNamespaceConfigResults** | Pointer to [**[]MemoryPersistentMemoryNamespaceConfigResultRelationship**](memory.PersistentMemoryNamespaceConfigResult.Relationship.md) | An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryConfigResultAllOf

`func NewMemoryPersistentMemoryConfigResultAllOf() *MemoryPersistentMemoryConfigResultAllOf`

NewMemoryPersistentMemoryConfigResultAllOf instantiates a new MemoryPersistentMemoryConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryConfigResultAllOfWithDefaults

`func NewMemoryPersistentMemoryConfigResultAllOfWithDefaults() *MemoryPersistentMemoryConfigResultAllOf`

NewMemoryPersistentMemoryConfigResultAllOfWithDefaults instantiates a new MemoryPersistentMemoryConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigErrorDesc() string`

GetConfigErrorDesc returns the ConfigErrorDesc field if non-nil, zero value otherwise.

### GetConfigErrorDescOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigErrorDescOk() (*string, bool)`

GetConfigErrorDescOk returns a tuple with the ConfigErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetConfigErrorDesc(v string)`

SetConfigErrorDesc sets ConfigErrorDesc field to given value.

### HasConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasConfigErrorDesc() bool`

HasConfigErrorDesc returns a boolean if a field has been set.

### GetConfigResult

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigResult() string`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigResultOk() (*string, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetConfigResult(v string)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigSequenceNo() int64`

GetConfigSequenceNo returns the ConfigSequenceNo field if non-nil, zero value otherwise.

### GetConfigSequenceNoOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigSequenceNoOk() (*int64, bool)`

GetConfigSequenceNoOk returns a tuple with the ConfigSequenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetConfigSequenceNo(v int64)`

SetConfigSequenceNo sets ConfigSequenceNo field to given value.

### HasConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasConfigSequenceNo() bool`

HasConfigSequenceNo returns a boolean if a field has been set.

### GetConfigState

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetPersistentMemoryNamespaceConfigResults() []MemoryPersistentMemoryNamespaceConfigResultRelationship`

GetPersistentMemoryNamespaceConfigResults returns the PersistentMemoryNamespaceConfigResults field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespaceConfigResultsOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetPersistentMemoryNamespaceConfigResultsOk() (*[]MemoryPersistentMemoryNamespaceConfigResultRelationship, bool)`

GetPersistentMemoryNamespaceConfigResultsOk returns a tuple with the PersistentMemoryNamespaceConfigResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetPersistentMemoryNamespaceConfigResults(v []MemoryPersistentMemoryNamespaceConfigResultRelationship)`

SetPersistentMemoryNamespaceConfigResults sets PersistentMemoryNamespaceConfigResults field to given value.

### HasPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasPersistentMemoryNamespaceConfigResults() bool`

HasPersistentMemoryNamespaceConfigResults returns a boolean if a field has been set.

### SetPersistentMemoryNamespaceConfigResultsNil

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetPersistentMemoryNamespaceConfigResultsNil(b bool)`

 SetPersistentMemoryNamespaceConfigResultsNil sets the value for PersistentMemoryNamespaceConfigResults to be an explicit nil

### UnsetPersistentMemoryNamespaceConfigResults
`func (o *MemoryPersistentMemoryConfigResultAllOf) UnsetPersistentMemoryNamespaceConfigResults()`

UnsetPersistentMemoryNamespaceConfigResults ensures that no value is present for PersistentMemoryNamespaceConfigResults, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryConfigResultAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResultAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


