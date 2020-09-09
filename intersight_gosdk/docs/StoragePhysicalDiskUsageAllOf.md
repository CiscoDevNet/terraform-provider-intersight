# StoragePhysicalDiskUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfBlocks** | Pointer to **string** | The number of blocks that are a part of the virtual drive. | [optional] [readonly] 
**PhysicalDrive** | Pointer to **string** | The physical disk for which the usage is reported. | [optional] [readonly] 
**Span** | Pointer to **string** | The span of the physical disk. | [optional] [readonly] 
**StartingBlock** | Pointer to **string** | The starting block id of the virtual drive within the physical drive. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the physical disk usage. | [optional] [readonly] 
**VirtualDrive** | Pointer to **string** | The virtual drive corresponding to the physical disk. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewStoragePhysicalDiskUsageAllOf

`func NewStoragePhysicalDiskUsageAllOf() *StoragePhysicalDiskUsageAllOf`

NewStoragePhysicalDiskUsageAllOf instantiates a new StoragePhysicalDiskUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskUsageAllOfWithDefaults

`func NewStoragePhysicalDiskUsageAllOfWithDefaults() *StoragePhysicalDiskUsageAllOf`

NewStoragePhysicalDiskUsageAllOfWithDefaults instantiates a new StoragePhysicalDiskUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfBlocks

`func (o *StoragePhysicalDiskUsageAllOf) GetNumberOfBlocks() string`

GetNumberOfBlocks returns the NumberOfBlocks field if non-nil, zero value otherwise.

### GetNumberOfBlocksOk

`func (o *StoragePhysicalDiskUsageAllOf) GetNumberOfBlocksOk() (*string, bool)`

GetNumberOfBlocksOk returns a tuple with the NumberOfBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBlocks

`func (o *StoragePhysicalDiskUsageAllOf) SetNumberOfBlocks(v string)`

SetNumberOfBlocks sets NumberOfBlocks field to given value.

### HasNumberOfBlocks

`func (o *StoragePhysicalDiskUsageAllOf) HasNumberOfBlocks() bool`

HasNumberOfBlocks returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StoragePhysicalDiskUsageAllOf) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StoragePhysicalDiskUsageAllOf) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StoragePhysicalDiskUsageAllOf) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StoragePhysicalDiskUsageAllOf) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetSpan

`func (o *StoragePhysicalDiskUsageAllOf) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *StoragePhysicalDiskUsageAllOf) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *StoragePhysicalDiskUsageAllOf) SetSpan(v string)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *StoragePhysicalDiskUsageAllOf) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetStartingBlock

`func (o *StoragePhysicalDiskUsageAllOf) GetStartingBlock() string`

GetStartingBlock returns the StartingBlock field if non-nil, zero value otherwise.

### GetStartingBlockOk

`func (o *StoragePhysicalDiskUsageAllOf) GetStartingBlockOk() (*string, bool)`

GetStartingBlockOk returns a tuple with the StartingBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBlock

`func (o *StoragePhysicalDiskUsageAllOf) SetStartingBlock(v string)`

SetStartingBlock sets StartingBlock field to given value.

### HasStartingBlock

`func (o *StoragePhysicalDiskUsageAllOf) HasStartingBlock() bool`

HasStartingBlock returns a boolean if a field has been set.

### GetState

`func (o *StoragePhysicalDiskUsageAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StoragePhysicalDiskUsageAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StoragePhysicalDiskUsageAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StoragePhysicalDiskUsageAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StoragePhysicalDiskUsageAllOf) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StoragePhysicalDiskUsageAllOf) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StoragePhysicalDiskUsageAllOf) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StoragePhysicalDiskUsageAllOf) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskUsageAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsageAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePhysicalDiskUsageAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskUsageAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskUsageAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskUsageAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


