# StorageNvmePhysicalDiskStateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmePhysicalDiskState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmePhysicalDiskState"]
**Slot** | Pointer to **string** | Physical Disk Slot that is to be configured. | [optional] [readonly] 
**State** | Pointer to **string** | Physical Disk State that is to be set at endpoint. | [optional] [readonly] 

## Methods

### NewStorageNvmePhysicalDiskStateAllOf

`func NewStorageNvmePhysicalDiskStateAllOf(classId string, objectType string, ) *StorageNvmePhysicalDiskStateAllOf`

NewStorageNvmePhysicalDiskStateAllOf instantiates a new StorageNvmePhysicalDiskStateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmePhysicalDiskStateAllOfWithDefaults

`func NewStorageNvmePhysicalDiskStateAllOfWithDefaults() *StorageNvmePhysicalDiskStateAllOf`

NewStorageNvmePhysicalDiskStateAllOfWithDefaults instantiates a new StorageNvmePhysicalDiskStateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmePhysicalDiskStateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmePhysicalDiskStateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmePhysicalDiskStateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmePhysicalDiskStateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmePhysicalDiskStateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmePhysicalDiskStateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlot

`func (o *StorageNvmePhysicalDiskStateAllOf) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *StorageNvmePhysicalDiskStateAllOf) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *StorageNvmePhysicalDiskStateAllOf) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *StorageNvmePhysicalDiskStateAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetState

`func (o *StorageNvmePhysicalDiskStateAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNvmePhysicalDiskStateAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNvmePhysicalDiskStateAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNvmePhysicalDiskStateAllOf) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


