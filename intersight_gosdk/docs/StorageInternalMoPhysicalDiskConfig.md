# StorageInternalMoPhysicalDiskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.InternalMoPhysicalDiskConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.InternalMoPhysicalDiskConfig"]
**EncryptionOp** | Pointer to **string** | Physical Disk Encryption operation that is to be set at endpoint. | [optional] [readonly] 
**Slot** | Pointer to **string** | Physical Disk Slot that is to be configured. | [optional] [readonly] 
**State** | Pointer to **string** | Physical Disk State that is to be set at endpoint. | [optional] [readonly] 

## Methods

### NewStorageInternalMoPhysicalDiskConfig

`func NewStorageInternalMoPhysicalDiskConfig(classId string, objectType string, ) *StorageInternalMoPhysicalDiskConfig`

NewStorageInternalMoPhysicalDiskConfig instantiates a new StorageInternalMoPhysicalDiskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageInternalMoPhysicalDiskConfigWithDefaults

`func NewStorageInternalMoPhysicalDiskConfigWithDefaults() *StorageInternalMoPhysicalDiskConfig`

NewStorageInternalMoPhysicalDiskConfigWithDefaults instantiates a new StorageInternalMoPhysicalDiskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageInternalMoPhysicalDiskConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageInternalMoPhysicalDiskConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageInternalMoPhysicalDiskConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageInternalMoPhysicalDiskConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageInternalMoPhysicalDiskConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageInternalMoPhysicalDiskConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEncryptionOp

`func (o *StorageInternalMoPhysicalDiskConfig) GetEncryptionOp() string`

GetEncryptionOp returns the EncryptionOp field if non-nil, zero value otherwise.

### GetEncryptionOpOk

`func (o *StorageInternalMoPhysicalDiskConfig) GetEncryptionOpOk() (*string, bool)`

GetEncryptionOpOk returns a tuple with the EncryptionOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOp

`func (o *StorageInternalMoPhysicalDiskConfig) SetEncryptionOp(v string)`

SetEncryptionOp sets EncryptionOp field to given value.

### HasEncryptionOp

`func (o *StorageInternalMoPhysicalDiskConfig) HasEncryptionOp() bool`

HasEncryptionOp returns a boolean if a field has been set.

### GetSlot

`func (o *StorageInternalMoPhysicalDiskConfig) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageInternalMoPhysicalDiskConfig) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageInternalMoPhysicalDiskConfig) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageInternalMoPhysicalDiskConfig) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetState

`func (o *StorageInternalMoPhysicalDiskConfig) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageInternalMoPhysicalDiskConfig) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageInternalMoPhysicalDiskConfig) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageInternalMoPhysicalDiskConfig) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


