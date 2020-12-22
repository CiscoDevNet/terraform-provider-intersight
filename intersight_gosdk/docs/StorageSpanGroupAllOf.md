# StorageSpanGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.SpanGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.SpanGroup"]
**Disks** | Pointer to [**[]StorageLocalDisk**](StorageLocalDisk.md) |  | [optional] 

## Methods

### NewStorageSpanGroupAllOf

`func NewStorageSpanGroupAllOf(classId string, objectType string, ) *StorageSpanGroupAllOf`

NewStorageSpanGroupAllOf instantiates a new StorageSpanGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpanGroupAllOfWithDefaults

`func NewStorageSpanGroupAllOfWithDefaults() *StorageSpanGroupAllOf`

NewStorageSpanGroupAllOfWithDefaults instantiates a new StorageSpanGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageSpanGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageSpanGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageSpanGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageSpanGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageSpanGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageSpanGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisks

`func (o *StorageSpanGroupAllOf) GetDisks() []StorageLocalDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *StorageSpanGroupAllOf) GetDisksOk() (*[]StorageLocalDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *StorageSpanGroupAllOf) SetDisks(v []StorageLocalDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *StorageSpanGroupAllOf) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *StorageSpanGroupAllOf) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *StorageSpanGroupAllOf) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


