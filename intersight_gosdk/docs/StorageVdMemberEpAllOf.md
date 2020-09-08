# StorageVdMemberEpAllOf

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

### NewStorageVdMemberEpAllOf

`func NewStorageVdMemberEpAllOf() *StorageVdMemberEpAllOf`

NewStorageVdMemberEpAllOf instantiates a new StorageVdMemberEpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVdMemberEpAllOfWithDefaults

`func NewStorageVdMemberEpAllOfWithDefaults() *StorageVdMemberEpAllOf`

NewStorageVdMemberEpAllOfWithDefaults instantiates a new StorageVdMemberEpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperQualifierReason

`func (o *StorageVdMemberEpAllOf) GetOperQualifierReason() string`

GetOperQualifierReason returns the OperQualifierReason field if non-nil, zero value otherwise.

### GetOperQualifierReasonOk

`func (o *StorageVdMemberEpAllOf) GetOperQualifierReasonOk() (*string, bool)`

GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperQualifierReason

`func (o *StorageVdMemberEpAllOf) SetOperQualifierReason(v string)`

SetOperQualifierReason sets OperQualifierReason field to given value.

### HasOperQualifierReason

`func (o *StorageVdMemberEpAllOf) HasOperQualifierReason() bool`

HasOperQualifierReason returns a boolean if a field has been set.

### GetPresence

`func (o *StorageVdMemberEpAllOf) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *StorageVdMemberEpAllOf) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *StorageVdMemberEpAllOf) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *StorageVdMemberEpAllOf) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetRole

`func (o *StorageVdMemberEpAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *StorageVdMemberEpAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *StorageVdMemberEpAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *StorageVdMemberEpAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSpanId

`func (o *StorageVdMemberEpAllOf) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *StorageVdMemberEpAllOf) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *StorageVdMemberEpAllOf) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *StorageVdMemberEpAllOf) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetVdMemberEpId

`func (o *StorageVdMemberEpAllOf) GetVdMemberEpId() int64`

GetVdMemberEpId returns the VdMemberEpId field if non-nil, zero value otherwise.

### GetVdMemberEpIdOk

`func (o *StorageVdMemberEpAllOf) GetVdMemberEpIdOk() (*int64, bool)`

GetVdMemberEpIdOk returns a tuple with the VdMemberEpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdMemberEpId

`func (o *StorageVdMemberEpAllOf) SetVdMemberEpId(v int64)`

SetVdMemberEpId sets VdMemberEpId field to given value.

### HasVdMemberEpId

`func (o *StorageVdMemberEpAllOf) HasVdMemberEpId() bool`

HasVdMemberEpId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *StorageVdMemberEpAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *StorageVdMemberEpAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *StorageVdMemberEpAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *StorageVdMemberEpAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageVdMemberEpAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageVdMemberEpAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageVdMemberEpAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageVdMemberEpAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetStorageVirtualDrive

`func (o *StorageVdMemberEpAllOf) GetStorageVirtualDrive() StorageVirtualDriveRelationship`

GetStorageVirtualDrive returns the StorageVirtualDrive field if non-nil, zero value otherwise.

### GetStorageVirtualDriveOk

`func (o *StorageVdMemberEpAllOf) GetStorageVirtualDriveOk() (*StorageVirtualDriveRelationship, bool)`

GetStorageVirtualDriveOk returns a tuple with the StorageVirtualDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVirtualDrive

`func (o *StorageVdMemberEpAllOf) SetStorageVirtualDrive(v StorageVirtualDriveRelationship)`

SetStorageVirtualDrive sets StorageVirtualDrive field to given value.

### HasStorageVirtualDrive

`func (o *StorageVdMemberEpAllOf) HasStorageVirtualDrive() bool`

HasStorageVirtualDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


