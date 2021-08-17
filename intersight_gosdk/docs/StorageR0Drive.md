# StorageR0Drive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.R0Drive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.R0Drive"]
**DriveSlots** | Pointer to **string** | The set of drive slots where RAID0 virtual drives must be created. | [optional] 
**DriveSlotsList** | Pointer to **string** | The list of drive slots where RAID0 virtual drives must be created (comma seperated). | [optional] [readonly] 
**Enable** | Pointer to **bool** | If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk. | [optional] [default to false]
**VirtualDrivePolicy** | Pointer to [**NullableStorageVirtualDrivePolicy**](storage.VirtualDrivePolicy.md) |  | [optional] 

## Methods

### NewStorageR0Drive

`func NewStorageR0Drive(classId string, objectType string, ) *StorageR0Drive`

NewStorageR0Drive instantiates a new StorageR0Drive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageR0DriveWithDefaults

`func NewStorageR0DriveWithDefaults() *StorageR0Drive`

NewStorageR0DriveWithDefaults instantiates a new StorageR0Drive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageR0Drive) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageR0Drive) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageR0Drive) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageR0Drive) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageR0Drive) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageR0Drive) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriveSlots

`func (o *StorageR0Drive) GetDriveSlots() string`

GetDriveSlots returns the DriveSlots field if non-nil, zero value otherwise.

### GetDriveSlotsOk

`func (o *StorageR0Drive) GetDriveSlotsOk() (*string, bool)`

GetDriveSlotsOk returns a tuple with the DriveSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSlots

`func (o *StorageR0Drive) SetDriveSlots(v string)`

SetDriveSlots sets DriveSlots field to given value.

### HasDriveSlots

`func (o *StorageR0Drive) HasDriveSlots() bool`

HasDriveSlots returns a boolean if a field has been set.

### GetDriveSlotsList

`func (o *StorageR0Drive) GetDriveSlotsList() string`

GetDriveSlotsList returns the DriveSlotsList field if non-nil, zero value otherwise.

### GetDriveSlotsListOk

`func (o *StorageR0Drive) GetDriveSlotsListOk() (*string, bool)`

GetDriveSlotsListOk returns a tuple with the DriveSlotsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSlotsList

`func (o *StorageR0Drive) SetDriveSlotsList(v string)`

SetDriveSlotsList sets DriveSlotsList field to given value.

### HasDriveSlotsList

`func (o *StorageR0Drive) HasDriveSlotsList() bool`

HasDriveSlotsList returns a boolean if a field has been set.

### GetEnable

`func (o *StorageR0Drive) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *StorageR0Drive) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *StorageR0Drive) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *StorageR0Drive) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetVirtualDrivePolicy

`func (o *StorageR0Drive) GetVirtualDrivePolicy() StorageVirtualDrivePolicy`

GetVirtualDrivePolicy returns the VirtualDrivePolicy field if non-nil, zero value otherwise.

### GetVirtualDrivePolicyOk

`func (o *StorageR0Drive) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool)`

GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDrivePolicy

`func (o *StorageR0Drive) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy)`

SetVirtualDrivePolicy sets VirtualDrivePolicy field to given value.

### HasVirtualDrivePolicy

`func (o *StorageR0Drive) HasVirtualDrivePolicy() bool`

HasVirtualDrivePolicy returns a boolean if a field has been set.

### SetVirtualDrivePolicyNil

`func (o *StorageR0Drive) SetVirtualDrivePolicyNil(b bool)`

 SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil

### UnsetVirtualDrivePolicy
`func (o *StorageR0Drive) UnsetVirtualDrivePolicy()`

UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


