# StorageVirtualDriveExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.VirtualDriveExtension"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.VirtualDriveExtension"]
**Bootable** | Pointer to **string** | The ability to boot from the virtual drive. | [optional] [readonly] 
**ContainerId** | Pointer to **int64** | The container id of the virtual drive. | [optional] [readonly] 
**DriveState** | Pointer to **string** | The state of the virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Virtual drive. | [optional] [readonly] 
**OperDeviceId** | Pointer to **string** | The operational device id of the virtual drive. | [optional] [readonly] 
**Uuid** | Pointer to **string** | The UUID assigned to the virtual drive. | [optional] [readonly] 
**VendorUuid** | Pointer to **string** | The UUID value of the vendor assigned to the virtual drive. | [optional] [readonly] 
**VirtualDriveDn** | Pointer to **string** | The distinguished name of the virtual drive for which the extended data is provided. | [optional] [readonly] 
**VirtualDriveId** | Pointer to **string** | The Id of the virtual drive. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**StorageController** | Pointer to [**StorageControllerRelationship**](StorageControllerRelationship.md) |  | [optional] 
**VirtualDrive** | Pointer to [**StorageVirtualDriveRelationship**](StorageVirtualDriveRelationship.md) |  | [optional] 

## Methods

### NewStorageVirtualDriveExtension

`func NewStorageVirtualDriveExtension(classId string, objectType string, ) *StorageVirtualDriveExtension`

NewStorageVirtualDriveExtension instantiates a new StorageVirtualDriveExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVirtualDriveExtensionWithDefaults

`func NewStorageVirtualDriveExtensionWithDefaults() *StorageVirtualDriveExtension`

NewStorageVirtualDriveExtensionWithDefaults instantiates a new StorageVirtualDriveExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageVirtualDriveExtension) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageVirtualDriveExtension) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageVirtualDriveExtension) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageVirtualDriveExtension) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageVirtualDriveExtension) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageVirtualDriveExtension) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootable

`func (o *StorageVirtualDriveExtension) GetBootable() string`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageVirtualDriveExtension) GetBootableOk() (*string, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageVirtualDriveExtension) SetBootable(v string)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageVirtualDriveExtension) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetContainerId

`func (o *StorageVirtualDriveExtension) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *StorageVirtualDriveExtension) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *StorageVirtualDriveExtension) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *StorageVirtualDriveExtension) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetDriveState

`func (o *StorageVirtualDriveExtension) GetDriveState() string`

GetDriveState returns the DriveState field if non-nil, zero value otherwise.

### GetDriveStateOk

`func (o *StorageVirtualDriveExtension) GetDriveStateOk() (*string, bool)`

GetDriveStateOk returns a tuple with the DriveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveState

`func (o *StorageVirtualDriveExtension) SetDriveState(v string)`

SetDriveState sets DriveState field to given value.

### HasDriveState

`func (o *StorageVirtualDriveExtension) HasDriveState() bool`

HasDriveState returns a boolean if a field has been set.

### GetName

`func (o *StorageVirtualDriveExtension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVirtualDriveExtension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVirtualDriveExtension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVirtualDriveExtension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperDeviceId

`func (o *StorageVirtualDriveExtension) GetOperDeviceId() string`

GetOperDeviceId returns the OperDeviceId field if non-nil, zero value otherwise.

### GetOperDeviceIdOk

`func (o *StorageVirtualDriveExtension) GetOperDeviceIdOk() (*string, bool)`

GetOperDeviceIdOk returns a tuple with the OperDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperDeviceId

`func (o *StorageVirtualDriveExtension) SetOperDeviceId(v string)`

SetOperDeviceId sets OperDeviceId field to given value.

### HasOperDeviceId

`func (o *StorageVirtualDriveExtension) HasOperDeviceId() bool`

HasOperDeviceId returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVirtualDriveExtension) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVirtualDriveExtension) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVirtualDriveExtension) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVirtualDriveExtension) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorUuid

`func (o *StorageVirtualDriveExtension) GetVendorUuid() string`

GetVendorUuid returns the VendorUuid field if non-nil, zero value otherwise.

### GetVendorUuidOk

`func (o *StorageVirtualDriveExtension) GetVendorUuidOk() (*string, bool)`

GetVendorUuidOk returns a tuple with the VendorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUuid

`func (o *StorageVirtualDriveExtension) SetVendorUuid(v string)`

SetVendorUuid sets VendorUuid field to given value.

### HasVendorUuid

`func (o *StorageVirtualDriveExtension) HasVendorUuid() bool`

HasVendorUuid returns a boolean if a field has been set.

### GetVirtualDriveDn

`func (o *StorageVirtualDriveExtension) GetVirtualDriveDn() string`

GetVirtualDriveDn returns the VirtualDriveDn field if non-nil, zero value otherwise.

### GetVirtualDriveDnOk

`func (o *StorageVirtualDriveExtension) GetVirtualDriveDnOk() (*string, bool)`

GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveDn

`func (o *StorageVirtualDriveExtension) SetVirtualDriveDn(v string)`

SetVirtualDriveDn sets VirtualDriveDn field to given value.

### HasVirtualDriveDn

`func (o *StorageVirtualDriveExtension) HasVirtualDriveDn() bool`

HasVirtualDriveDn returns a boolean if a field has been set.

### GetVirtualDriveId

`func (o *StorageVirtualDriveExtension) GetVirtualDriveId() string`

GetVirtualDriveId returns the VirtualDriveId field if non-nil, zero value otherwise.

### GetVirtualDriveIdOk

`func (o *StorageVirtualDriveExtension) GetVirtualDriveIdOk() (*string, bool)`

GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDriveId

`func (o *StorageVirtualDriveExtension) SetVirtualDriveId(v string)`

SetVirtualDriveId sets VirtualDriveId field to given value.

### HasVirtualDriveId

`func (o *StorageVirtualDriveExtension) HasVirtualDriveId() bool`

HasVirtualDriveId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVirtualDriveExtension) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVirtualDriveExtension) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageVirtualDriveExtension) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVirtualDriveExtension) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVirtualDriveExtension) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVirtualDriveExtension) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageController

`func (o *StorageVirtualDriveExtension) GetStorageController() StorageControllerRelationship`

GetStorageController returns the StorageController field if non-nil, zero value otherwise.

### GetStorageControllerOk

`func (o *StorageVirtualDriveExtension) GetStorageControllerOk() (*StorageControllerRelationship, bool)`

GetStorageControllerOk returns a tuple with the StorageController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageController

`func (o *StorageVirtualDriveExtension) SetStorageController(v StorageControllerRelationship)`

SetStorageController sets StorageController field to given value.

### HasStorageController

`func (o *StorageVirtualDriveExtension) HasStorageController() bool`

HasStorageController returns a boolean if a field has been set.

### GetVirtualDrive

`func (o *StorageVirtualDriveExtension) GetVirtualDrive() StorageVirtualDriveRelationship`

GetVirtualDrive returns the VirtualDrive field if non-nil, zero value otherwise.

### GetVirtualDriveOk

`func (o *StorageVirtualDriveExtension) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetVirtualDriveOk returns a tuple with the VirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrive

`func (o *StorageVirtualDriveExtension) SetVirtualDrive(v StorageVirtualDriveRelationship)`

SetVirtualDrive sets VirtualDrive field to given value.

### HasVirtualDrive

`func (o *StorageVirtualDriveExtension) HasVirtualDrive() bool`

HasVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


