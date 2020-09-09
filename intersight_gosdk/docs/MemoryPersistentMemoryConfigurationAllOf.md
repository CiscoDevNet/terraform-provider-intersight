# MemoryPersistentMemoryConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryCapacity** | Pointer to **string** | Memory capacity in GiB of a Persistent Memory configuration on a server. | [optional] [readonly] 
**NumOfModules** | Pointer to **string** | Number of Persistent Memory Modules on a server. | [optional] [readonly] 
**NumOfRegions** | Pointer to **string** | Number of Persistent Memory Regions on a server. | [optional] [readonly] 
**PersistentMemoryCapacity** | Pointer to **string** | Persistent memory capacity in GiB of a Persistent Memory configuration on a server. | [optional] [readonly] 
**ReservedCapacity** | Pointer to **string** | Reserved capacity in GiB of a Persistent Memory configuration on a server. | [optional] [readonly] 
**SecurityState** | Pointer to **string** | Collective security state of all Persistent Memory modules on a server. | [optional] [readonly] 
**TotalCapacity** | Pointer to **string** | Total capacity in GiB of a Persistent Memory configuration on a server. | [optional] [readonly] 
**ComputeBoard** | Pointer to [**ComputeBoardRelationship**](compute.Board.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**PersistentMemoryConfigResult** | Pointer to [**MemoryPersistentMemoryConfigResultRelationship**](memory.PersistentMemoryConfigResult.Relationship.md) |  | [optional] 
**PersistentMemoryRegions** | Pointer to [**[]MemoryPersistentMemoryRegionRelationship**](memory.PersistentMemoryRegion.Relationship.md) | An array of relationships to memoryPersistentMemoryRegion resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewMemoryPersistentMemoryConfigurationAllOf

`func NewMemoryPersistentMemoryConfigurationAllOf() *MemoryPersistentMemoryConfigurationAllOf`

NewMemoryPersistentMemoryConfigurationAllOf instantiates a new MemoryPersistentMemoryConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryConfigurationAllOfWithDefaults

`func NewMemoryPersistentMemoryConfigurationAllOfWithDefaults() *MemoryPersistentMemoryConfigurationAllOf`

NewMemoryPersistentMemoryConfigurationAllOfWithDefaults instantiates a new MemoryPersistentMemoryConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetMemoryCapacity() string`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetMemoryCapacityOk() (*string, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetMemoryCapacity(v string)`

SetMemoryCapacity sets MemoryCapacity field to given value.

### HasMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasMemoryCapacity() bool`

HasMemoryCapacity returns a boolean if a field has been set.

### GetNumOfModules

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetNumOfModules() string`

GetNumOfModules returns the NumOfModules field if non-nil, zero value otherwise.

### GetNumOfModulesOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetNumOfModulesOk() (*string, bool)`

GetNumOfModulesOk returns a tuple with the NumOfModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfModules

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetNumOfModules(v string)`

SetNumOfModules sets NumOfModules field to given value.

### HasNumOfModules

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasNumOfModules() bool`

HasNumOfModules returns a boolean if a field has been set.

### GetNumOfRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetNumOfRegions() string`

GetNumOfRegions returns the NumOfRegions field if non-nil, zero value otherwise.

### GetNumOfRegionsOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetNumOfRegionsOk() (*string, bool)`

GetNumOfRegionsOk returns a tuple with the NumOfRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetNumOfRegions(v string)`

SetNumOfRegions sets NumOfRegions field to given value.

### HasNumOfRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasNumOfRegions() bool`

HasNumOfRegions returns a boolean if a field has been set.

### GetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryCapacity() string`

GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field if non-nil, zero value otherwise.

### GetPersistentMemoryCapacityOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryCapacityOk() (*string, bool)`

GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetPersistentMemoryCapacity(v string)`

SetPersistentMemoryCapacity sets PersistentMemoryCapacity field to given value.

### HasPersistentMemoryCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasPersistentMemoryCapacity() bool`

HasPersistentMemoryCapacity returns a boolean if a field has been set.

### GetReservedCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetReservedCapacity() string`

GetReservedCapacity returns the ReservedCapacity field if non-nil, zero value otherwise.

### GetReservedCapacityOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetReservedCapacityOk() (*string, bool)`

GetReservedCapacityOk returns a tuple with the ReservedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetReservedCapacity(v string)`

SetReservedCapacity sets ReservedCapacity field to given value.

### HasReservedCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasReservedCapacity() bool`

HasReservedCapacity returns a boolean if a field has been set.

### GetSecurityState

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetSecurityState() string`

GetSecurityState returns the SecurityState field if non-nil, zero value otherwise.

### GetSecurityStateOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetSecurityStateOk() (*string, bool)`

GetSecurityStateOk returns a tuple with the SecurityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityState

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetSecurityState(v string)`

SetSecurityState sets SecurityState field to given value.

### HasSecurityState

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasSecurityState() bool`

HasSecurityState returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetTotalCapacity() string`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetTotalCapacityOk() (*string, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetTotalCapacity(v string)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetComputeBoard

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetComputeBoard() ComputeBoardRelationship`

GetComputeBoard returns the ComputeBoard field if non-nil, zero value otherwise.

### GetComputeBoardOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool)`

GetComputeBoardOk returns a tuple with the ComputeBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBoard

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetComputeBoard(v ComputeBoardRelationship)`

SetComputeBoard sets ComputeBoard field to given value.

### HasComputeBoard

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasComputeBoard() bool`

HasComputeBoard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryConfigResult() MemoryPersistentMemoryConfigResultRelationship`

GetPersistentMemoryConfigResult returns the PersistentMemoryConfigResult field if non-nil, zero value otherwise.

### GetPersistentMemoryConfigResultOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryConfigResultOk() (*MemoryPersistentMemoryConfigResultRelationship, bool)`

GetPersistentMemoryConfigResultOk returns a tuple with the PersistentMemoryConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetPersistentMemoryConfigResult(v MemoryPersistentMemoryConfigResultRelationship)`

SetPersistentMemoryConfigResult sets PersistentMemoryConfigResult field to given value.

### HasPersistentMemoryConfigResult

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasPersistentMemoryConfigResult() bool`

HasPersistentMemoryConfigResult returns a boolean if a field has been set.

### GetPersistentMemoryRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryRegions() []MemoryPersistentMemoryRegionRelationship`

GetPersistentMemoryRegions returns the PersistentMemoryRegions field if non-nil, zero value otherwise.

### GetPersistentMemoryRegionsOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetPersistentMemoryRegionsOk() (*[]MemoryPersistentMemoryRegionRelationship, bool)`

GetPersistentMemoryRegionsOk returns a tuple with the PersistentMemoryRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentMemoryRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetPersistentMemoryRegions(v []MemoryPersistentMemoryRegionRelationship)`

SetPersistentMemoryRegions sets PersistentMemoryRegions field to given value.

### HasPersistentMemoryRegions

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasPersistentMemoryRegions() bool`

HasPersistentMemoryRegions returns a boolean if a field has been set.

### SetPersistentMemoryRegionsNil

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetPersistentMemoryRegionsNil(b bool)`

 SetPersistentMemoryRegionsNil sets the value for PersistentMemoryRegions to be an explicit nil

### UnsetPersistentMemoryRegions
`func (o *MemoryPersistentMemoryConfigurationAllOf) UnsetPersistentMemoryRegions()`

UnsetPersistentMemoryRegions ensures that no value is present for PersistentMemoryRegions, not even an explicit nil
### GetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *MemoryPersistentMemoryConfigurationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *MemoryPersistentMemoryConfigurationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *MemoryPersistentMemoryConfigurationAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


