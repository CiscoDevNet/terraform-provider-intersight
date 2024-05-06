# StorageNvmeDedicatedHotSpareConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeDedicatedHotSpareConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeDedicatedHotSpareConfiguration"]
**IsNewVd** | Pointer to **bool** | This defines if the vd does not exists at endpoint for specific storage controller per drive group. Only if it&#39;s false we will create dedicated hot spares for the existing vds. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed. | [optional] [readonly] 
**Slot** | Pointer to **string** | Physical Disk Slot that is used as dedicated hot spare. | [optional] [readonly] 
**VolumeDn** | Pointer to **string** | The volume dn of the dedicated hot spare, this will be unique for each dedicated hot spare. | [optional] [readonly] 

## Methods

### NewStorageNvmeDedicatedHotSpareConfiguration

`func NewStorageNvmeDedicatedHotSpareConfiguration(classId string, objectType string, ) *StorageNvmeDedicatedHotSpareConfiguration`

NewStorageNvmeDedicatedHotSpareConfiguration instantiates a new StorageNvmeDedicatedHotSpareConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeDedicatedHotSpareConfigurationWithDefaults

`func NewStorageNvmeDedicatedHotSpareConfigurationWithDefaults() *StorageNvmeDedicatedHotSpareConfiguration`

NewStorageNvmeDedicatedHotSpareConfigurationWithDefaults instantiates a new StorageNvmeDedicatedHotSpareConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetIsNewVd() bool`

GetIsNewVd returns the IsNewVd field if non-nil, zero value otherwise.

### GetIsNewVdOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetIsNewVdOk() (*bool, bool)`

GetIsNewVdOk returns a tuple with the IsNewVd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetIsNewVd(v bool)`

SetIsNewVd sets IsNewVd field to given value.

### HasIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfiguration) HasIsNewVd() bool`

HasIsNewVd returns a boolean if a field has been set.

### GetName

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNvmeDedicatedHotSpareConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlot

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageNvmeDedicatedHotSpareConfiguration) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetVolumeDn() string`

GetVolumeDn returns the VolumeDn field if non-nil, zero value otherwise.

### GetVolumeDnOk

`func (o *StorageNvmeDedicatedHotSpareConfiguration) GetVolumeDnOk() (*string, bool)`

GetVolumeDnOk returns a tuple with the VolumeDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfiguration) SetVolumeDn(v string)`

SetVolumeDn sets VolumeDn field to given value.

### HasVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfiguration) HasVolumeDn() bool`

HasVolumeDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


