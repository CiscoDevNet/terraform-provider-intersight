# StorageNvmeDedicatedHotSpareConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeDedicatedHotSpareConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeDedicatedHotSpareConfiguration"]
**IsNewVd** | Pointer to **bool** | This defines if the vd does not exists at endpoint for specific storage controller per drive group. Only if it&#39;s false we will create dedicated hot spares for the existing vds. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. | [optional] [readonly] 
**Slot** | Pointer to **string** | Physical Disk Slot that is used as dedicated hot spare. | [optional] [readonly] 
**VolumeDn** | Pointer to **string** | The volume dn of the dedicated hot spare, this will be unique for each dedicated hot spare. | [optional] [readonly] 

## Methods

### NewStorageNvmeDedicatedHotSpareConfigurationAllOf

`func NewStorageNvmeDedicatedHotSpareConfigurationAllOf(classId string, objectType string, ) *StorageNvmeDedicatedHotSpareConfigurationAllOf`

NewStorageNvmeDedicatedHotSpareConfigurationAllOf instantiates a new StorageNvmeDedicatedHotSpareConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeDedicatedHotSpareConfigurationAllOfWithDefaults

`func NewStorageNvmeDedicatedHotSpareConfigurationAllOfWithDefaults() *StorageNvmeDedicatedHotSpareConfigurationAllOf`

NewStorageNvmeDedicatedHotSpareConfigurationAllOfWithDefaults instantiates a new StorageNvmeDedicatedHotSpareConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetIsNewVd() bool`

GetIsNewVd returns the IsNewVd field if non-nil, zero value otherwise.

### GetIsNewVdOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetIsNewVdOk() (*bool, bool)`

GetIsNewVdOk returns a tuple with the IsNewVd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetIsNewVd(v bool)`

SetIsNewVd sets IsNewVd field to given value.

### HasIsNewVd

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) HasIsNewVd() bool`

HasIsNewVd returns a boolean if a field has been set.

### GetName

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlot

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetVolumeDn() string`

GetVolumeDn returns the VolumeDn field if non-nil, zero value otherwise.

### GetVolumeDnOk

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) GetVolumeDnOk() (*string, bool)`

GetVolumeDnOk returns a tuple with the VolumeDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) SetVolumeDn(v string)`

SetVolumeDn sets VolumeDn field to given value.

### HasVolumeDn

`func (o *StorageNvmeDedicatedHotSpareConfigurationAllOf) HasVolumeDn() bool`

HasVolumeDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


