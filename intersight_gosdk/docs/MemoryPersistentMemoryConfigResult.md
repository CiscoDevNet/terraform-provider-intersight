# MemoryPersistentMemoryConfigResult

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

### NewMemoryPersistentMemoryConfigResult

`func NewMemoryPersistentMemoryConfigResult() *MemoryPersistentMemoryConfigResult`

NewMemoryPersistentMemoryConfigResult instantiates a new MemoryPersistentMemoryConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryConfigResultWithDefaults

`func NewMemoryPersistentMemoryConfigResultWithDefaults() *MemoryPersistentMemoryConfigResult`

NewMemoryPersistentMemoryConfigResultWithDefaults instantiates a new MemoryPersistentMemoryConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDesc() string`

GetConfigErrorDesc returns the ConfigErrorDesc field if non-nil, zero value otherwise.

### GetConfigErrorDescOk

`func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDescOk() (*string, bool)`

GetConfigErrorDescOk returns a tuple with the ConfigErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResult) SetConfigErrorDesc(v string)`

SetConfigErrorDesc sets ConfigErrorDesc field to given value.

### HasConfigErrorDesc

`func (o *MemoryPersistentMemoryConfigResult) HasConfigErrorDesc() bool`

HasConfigErrorDesc returns a boolean if a field has been set.

### GetConfigResult

`func (o *MemoryPersistentMemoryConfigResult) GetConfigResult() string`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *MemoryPersistentMemoryConfigResult) GetConfigResultOk() (*string, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *MemoryPersistentMemoryConfigResult) SetConfigResult(v string)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *MemoryPersistentMemoryConfigResult) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNo() int64`

GetConfigSequenceNo returns the ConfigSequenceNo field if non-nil, zero value otherwise.

### GetConfigSequenceNoOk

`func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNoOk() (*int64, bool)`

GetConfigSequenceNoOk returns a tuple with the ConfigSequenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResult) SetConfigSequenceNo(v int64)`

SetConfigSequenceNo sets ConfigSequenceNo field to given value.

### HasConfigSequenceNo

`func (o *MemoryPersistentMemoryConfigResult) HasConfigSequenceNo() bool`

HasConfigSequenceNo returns a boolean if a field has been set.

### GetConfigState

`func (o *MemoryPersistentMemoryConfigResult) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *MemoryPersistentMemoryConfigResult) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *MemoryPersistentMemoryConfigResult) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *MemoryPersistentMemoryConfigResult) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigResult) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship`

GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field if non-nil, zero value otherwise.

### GetMemoryPersistentMemoryConfigurationOk

`func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool)`

GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResult) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship)`

SetMemoryPersistentMemoryConfiguration sets MemoryPersistentMemoryConfiguration field to given value.

### HasMemoryPersistentMemoryConfiguration

`func (o *MemoryPersistentMemoryConfigResult) HasMemoryPersistentMemoryConfiguration() bool`

HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.

### GetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResults() []MemoryPersistentMemoryNamespaceConfigResultRelationship`

GetPersistentMemoryNamespaceConfigResults returns the PersistentMemoryNamespaceConfigResults field if non-nil, zero value otherwise.

### GetPersistentMemoryNamespaceConfigResultsOk

`func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResultsOk() (*[]MemoryPersistentMemoryNamespaceConfigResultRelationship, bool)`

GetPersistentMemoryNamespaceConfigResultsOk returns a tuple with the PersistentMemoryNamespaceConfigResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResult) SetPersistentMemoryNamespaceConfigResults(v []MemoryPersistentMemoryNamespaceConfigResultRelationship)`

SetPersistentMemoryNamespaceConfigResults sets PersistentMemoryNamespaceConfigResults field to given value.

### HasPersistentMemoryNamespaceConfigResults

`func (o *MemoryPersistentMemoryConfigResult) HasPersistentMemoryNamespaceConfigResults() bool`

HasPersistentMemoryNamespaceConfigResults returns a boolean if a field has been set.

### SetPersistentMemoryNamespaceConfigResultsNil

`func (o *MemoryPersistentMemoryConfigResult) SetPersistentMemoryNamespaceConfigResultsNil(b bool)`

 SetPersistentMemoryNamespaceConfigResultsNil sets the value for PersistentMemoryNamespaceConfigResults to be an explicit nil

### UnsetPersistentMemoryNamespaceConfigResults
`func (o *MemoryPersistentMemoryConfigResult) UnsetPersistentMemoryNamespaceConfigResults()`

UnsetPersistentMemoryNamespaceConfigResults ensures that no value is present for PersistentMemoryNamespaceConfigResults, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryConfigResult) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


