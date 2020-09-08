# StoragePhysicalDiskUsage

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

### NewStoragePhysicalDiskUsage

`func NewStoragePhysicalDiskUsage() *StoragePhysicalDiskUsage`

NewStoragePhysicalDiskUsage instantiates a new StoragePhysicalDiskUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhysicalDiskUsageWithDefaults

`func NewStoragePhysicalDiskUsageWithDefaults() *StoragePhysicalDiskUsage`

NewStoragePhysicalDiskUsageWithDefaults instantiates a new StoragePhysicalDiskUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfBlocks

`func (o *StoragePhysicalDiskUsage) GetNumberOfBlocks() string`

GetNumberOfBlocks returns the NumberOfBlocks field if non-nil, zero value otherwise.

### GetNumberOfBlocksOk

`func (o *StoragePhysicalDiskUsage) GetNumberOfBlocksOk() (*string, bool)`

GetNumberOfBlocksOk returns a tuple with the NumberOfBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBlocks

`func (o *StoragePhysicalDiskUsage) SetNumberOfBlocks(v string)`

SetNumberOfBlocks sets NumberOfBlocks field to given value.

### HasNumberOfBlocks

`func (o *StoragePhysicalDiskUsage) HasNumberOfBlocks() bool`

HasNumberOfBlocks returns a boolean if a field has been set.

### GetPhysicalDrive

`func (o *StoragePhysicalDiskUsage) GetPhysicalDrive() string`

GetPhysicalDrive returns the PhysicalDrive field if non-nil, zero value otherwise.

### GetPhysicalDriveOk

`func (o *StoragePhysicalDiskUsage) GetPhysicalDriveOk() (*string, bool)`

GetPhysicalDriveOk returns a tuple with the PhysicalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDrive

`func (o *StoragePhysicalDiskUsage) SetPhysicalDrive(v string)`

SetPhysicalDrive sets PhysicalDrive field to given value.

### HasPhysicalDrive

`func (o *StoragePhysicalDiskUsage) HasPhysicalDrive() bool`

HasPhysicalDrive returns a boolean if a field has been set.

### GetSpan

`func (o *StoragePhysicalDiskUsage) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *StoragePhysicalDiskUsage) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *StoragePhysicalDiskUsage) SetSpan(v string)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *StoragePhysicalDiskUsage) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetStartingBlock

`func (o *StoragePhysicalDiskUsage) GetStartingBlock() string`

GetStartingBlock returns the StartingBlock field if non-nil, zero value otherwise.

### GetStartingBlockOk

`func (o *StoragePhysicalDiskUsage) GetStartingBlockOk() (*string, bool)`

GetStartingBlockOk returns a tuple with the StartingBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBlock

`func (o *StoragePhysicalDiskUsage) SetStartingBlock(v string)`

SetStartingBlock sets StartingBlock field to given value.

### HasStartingBlock

`func (o *StoragePhysicalDiskUsage) HasStartingBlock() bool`

HasStartingBlock returns a boolean if a field has been set.

### GetState

`func (o *StoragePhysicalDiskUsage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StoragePhysicalDiskUsage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StoragePhysicalDiskUsage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StoragePhysicalDiskUsage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StoragePhysicalDiskUsage) GetVirtualDrive() string`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StoragePhysicalDiskUsage) GetVirtualDriveOk() (*string, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StoragePhysicalDiskUsage) SetVirtualDrive(v string)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StoragePhysicalDiskUsage) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsage) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StoragePhysicalDiskUsage) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsage) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StoragePhysicalDiskUsage) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StoragePhysicalDiskUsage) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePhysicalDiskUsage) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePhysicalDiskUsage) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePhysicalDiskUsage) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


