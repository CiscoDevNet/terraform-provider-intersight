# StorageM2VirtualDriveConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.M2VirtualDriveConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.M2VirtualDriveConfig"]
**ControllerSlot** | Pointer to **string** | Select the M.2 RAID controller slot on which the virtual drive is to be created. For example, MSTOR-RAID-1 means that the virtual drive is to be created on the M.2 RAID controller in the first slot. Select MSTOR-RAID-2 only when there is a M.2 RAID controller in the second slot. * &#x60;MSTOR-RAID-1&#x60; - Virtual drive  will be created on the M.2 RAID controller in the first slot. * &#x60;MSTOR-RAID-2&#x60; - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available. * &#x60;MSTOR-RAID-1,MSTOR-RAID-2&#x60; - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available. | [optional] [default to "MSTOR-RAID-1"]
**Enable** | Pointer to **bool** | If enabled, this will create a virtual drive on the M.2 RAID controller. | [optional] [default to false]

## Methods

### NewStorageM2VirtualDriveConfigAllOf

`func NewStorageM2VirtualDriveConfigAllOf(classId string, objectType string, ) *StorageM2VirtualDriveConfigAllOf`

NewStorageM2VirtualDriveConfigAllOf instantiates a new StorageM2VirtualDriveConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageM2VirtualDriveConfigAllOfWithDefaults

`func NewStorageM2VirtualDriveConfigAllOfWithDefaults() *StorageM2VirtualDriveConfigAllOf`

NewStorageM2VirtualDriveConfigAllOfWithDefaults instantiates a new StorageM2VirtualDriveConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageM2VirtualDriveConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageM2VirtualDriveConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageM2VirtualDriveConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageM2VirtualDriveConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageM2VirtualDriveConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageM2VirtualDriveConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerSlot

`func (o *StorageM2VirtualDriveConfigAllOf) GetControllerSlot() string`

GetControllerSlot returns the ControllerSlot field if non-nil, zero value otherwise.

### GetControllerSlotOk

`func (o *StorageM2VirtualDriveConfigAllOf) GetControllerSlotOk() (*string, bool)`

GetControllerSlotOk returns a tuple with the ControllerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerSlot

`func (o *StorageM2VirtualDriveConfigAllOf) SetControllerSlot(v string)`

SetControllerSlot sets ControllerSlot field to given value.

### HasControllerSlot

`func (o *StorageM2VirtualDriveConfigAllOf) HasControllerSlot() bool`

HasControllerSlot returns a boolean if a field has been set.

### GetEnable

`func (o *StorageM2VirtualDriveConfigAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *StorageM2VirtualDriveConfigAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *StorageM2VirtualDriveConfigAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *StorageM2VirtualDriveConfigAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


