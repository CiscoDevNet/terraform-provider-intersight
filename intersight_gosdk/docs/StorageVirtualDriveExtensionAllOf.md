# StorageVirtualDriveExtensionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootable** | Pointer to **string** | The ability to boot from the virtual drive. | [optional] [readonly] 
**ContainerId** | Pointer to **int64** | The container id of the virtual drive. | [optional] [readonly] 
**DriveState** | Pointer to **string** | The state of the virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Virtual drive. | [optional] [readonly] 
**OperDeviceId** | Pointer to **string** | The operational device id of the virtual drive. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID assigned to the virtual drive. | [optional] [readonly] 
**VendorUuid** | Pointer to **string** | The UUID value of the vendor assigned to the virtual drive. | [optional] [readonly] 
**VirtualDriveDn** | Pointer to **string** | The distinguished name of the virtual drive for which the extended data is provided. | [optional] [readonly] 
**VirtualDriveId** | Pointer to **string** | The Id of the virtual drive. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](storage.Controller.Relationship.md) |  | [optional] 
**VirtualDrive** | Pointer to [**StorageVirtualDriveRelationship**](storage.VirtualDrive.Relationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveExtensionAllOf

`func NewStorageVirtualDriveExtensionAllOf() *StorageVirtualDriveExtensionAllOf`

NewStorageVirtualDriveExtensionAllOf instantiates a new StorageVirtualDriveExtensionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveExtensionAllOfWithDefaults

`func NewStorageVirtualDriveExtensionAllOfWithDefaults() *StorageVirtualDriveExtensionAllOf`

NewStorageVirtualDriveExtensionAllOfWithDefaults instantiates a new StorageVirtualDriveExtensionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootable

`func (o *StorageVirtualDriveExtensionAllOf) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDriveExtensionAllOf) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDriveExtensionAllOf) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDriveExtensionAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetContainerId

`func (o *StorageVirtualDriveExtensionAllOf) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *StorageVirtualDriveExtensionAllOf) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *StorageVirtualDriveExtensionAllOf) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *StorageVirtualDriveExtensionAllOf) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDriveExtensionAllOf) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDriveExtensionAllOf) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDriveExtensionAllOf) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDriveExtensionAllOf) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveExtensionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveExtensionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveExtensionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveExtensionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperDeviceId

`func (o *StorageVirtualDriveExtensionAllOf) GetOperDeviceId() string`

GetOperDeviceId returns the OperDeviceId field if non-nil, zero value otherwise.

### GetOperDeviceIdOk

`func (o *StorageVirtualDriveExtensionAllOf) GetOperDeviceIdOk() (*string, bool)`

GetOperDeviceIdOk returns a tuple with the OperDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperDeviceId

`func (o *StorageVirtualDriveExtensionAllOf) SetOperDeviceId(v string)`

SetOperDeviceId sets OperDeviceId field to given value.

### HasOperDeviceId

`func (o *StorageVirtualDriveExtensionAllOf) HasOperDeviceId() bool`

HasOperDeviceId returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDriveExtensionAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDriveExtensionAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDriveExtensionAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDriveExtensionAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDriveExtensionAllOf) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDriveExtensionAllOf) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDriveExtensionAllOf) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDriveExtensionAllOf) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveDn

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveDn() string`

GetVirtualDriveDn returns the VirtualDriveDn field if non-nil, zero value otherwise.

### GetVirtualDriveDnOk

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveDnOk() (*string, bool)`

GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveDn

`func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDriveDn(v string)`

SetVirtualDriveDn sets VirtualDriveDn field to given value.

### HasVirtualDriveDn

`func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDriveDn() bool`

HasVirtualDriveDn returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDriveExtensionAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDriveExtensionAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageVirtualDriveExtensionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDriveExtensionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDriveExtensionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDriveExtensionAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDriveExtensionAllOf) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDriveExtensionAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDriveExtensionAllOf) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDriveExtensionAllOf) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDrive() StorageVirtualDriveRelationship`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDrive(v StorageVirtualDriveRelationship)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


