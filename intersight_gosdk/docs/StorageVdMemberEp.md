# StorageVdMemberEp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperQualifierReason** | Pointer to **string** | For certain states, indicates the reason why the operState is in that state. | [optional] [readonly] 
**Presence** | Pointer to **string** | The presence state of the local disk. | [optional] [readonly] 
**Role** | Pointer to **string** | Role of the disk normal or hot-spare, used by virtual-drive. | [optional] [readonly] 
**SpanId** | Pointer to **string** | The span id number of the virtual drive. | [optional] [readonly] 
**VdMemberEpId** | Pointer to **int64** | The local disk slot number as id. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**StorageVirtualDrive** | Pointer to [**StorageVirtualDriveRelationship**](storage.VirtualDrive.Relationship.md) |  | [optional] 

## Methods

### NewStorageVdMemberEp

`func NewStorageVdMemberEp() *StorageVdMemberEp`

NewStorageVdMemberEp instantiates a new StorageVdMemberEp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVdMemberEpWithDefaults

`func NewStorageVdMemberEpWithDefaults() *StorageVdMemberEp`

NewStorageVdMemberEpWithDefaults instantiates a new StorageVdMemberEp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperQualifierReason

`func (o *StorageVdMemberEp) GetOperQualifierReason() string`

GetOperQualifierReason returns the OperQualifierReason field if non-nil, zero value otherwise.

### GetOperQualifierReasonOk

`func (o *StorageVdMemberEp) GetOperQualifierReasonOk() (*string, bool)`

GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifierReason

`func (o *StorageVdMemberEp) SetOperQualifierReason(v string)`

SetOperQualifierReason sets OperQualifierReason field to given value.

### HasOperQualifierReason

`func (o *StorageVdMemberEp) HasOperQualifierReason() bool`

HasOperQualifierReason returns a boolean if a field has been set.

### GetPresence

`func (o *StorageVdMemberEp) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageVdMemberEp) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageVdMemberEp) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageVdMemberEp) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRole

`func (o *StorageVdMemberEp) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *StorageVdMemberEp) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *StorageVdMemberEp) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *StorageVdMemberEp) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSpanId

`func (o *StorageVdMemberEp) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *StorageVdMemberEp) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *StorageVdMemberEp) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *StorageVdMemberEp) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetVdMemberEpId

`func (o *StorageVdMemberEp) GetVdMemberEpId() int64`

GetVdMemberEpId returns the VdMemberEpId field if non-nil, zero value otherwise.

### GetVdMemberEpIdOk

`func (o *StorageVdMemberEp) GetVdMemberEpIdOk() (*int64, bool)`

GetVdMemberEpIdOk returns a tuple with the VdMemberEpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdMemberEpId

`func (o *StorageVdMemberEp) SetVdMemberEpId(v int64)`

SetVdMemberEpId sets VdMemberEpId field to given value.

### HasVdMemberEpId

`func (o *StorageVdMemberEp) HasVdMemberEpId() bool`

HasVdMemberEpId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVdMemberEp) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVdMemberEp) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVdMemberEp) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVdMemberEp) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageVdMemberEp) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVdMemberEp) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVdMemberEp) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVdMemberEp) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageVirtualDrive

`func (o *StorageVdMemberEp) GetStorageVirtualDrive() StorageVirtualDriveRelationship`

GetStorageVirtualDrive returns the StorageVirtualDrive field if non-nil, zero value otherwise.

### GetStorageVirtualDriveOk

`func (o *StorageVdMemberEp) GetStorageVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetStorageVirtualDriveOk returns a tuple with the StorageVirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualDrive

`func (o *StorageVdMemberEp) SetStorageVirtualDrive(v StorageVirtualDriveRelationship)`

SetStorageVirtualDrive sets StorageVirtualDrive field to given value.

### HasStorageVirtualDrive

`func (o *StorageVdMemberEp) HasStorageVirtualDrive() bool`

HasStorageVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


