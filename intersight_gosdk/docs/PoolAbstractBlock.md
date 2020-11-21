# PoolAbstractBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**FreeBlockCount** | Pointer to **int64** | Free IDs that can be allocated in this block. | [optional] [readonly] 
**NextIdAllocator** | Pointer to **int64** | Moving counter to allocate the next identifier. | [optional] [readonly] 

## Methods

### NewPoolAbstractBlock

`func NewPoolAbstractBlock(classId string, objectType string, ) *PoolAbstractBlock`

NewPoolAbstractBlock instantiates a new PoolAbstractBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractBlockWithDefaults

`func NewPoolAbstractBlockWithDefaults() *PoolAbstractBlock`

NewPoolAbstractBlockWithDefaults instantiates a new PoolAbstractBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractBlock) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractBlock) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractBlock) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFreeBlockCount

`func (o *PoolAbstractBlock) GetFreeBlockCount() int64`

GetFreeBlockCount returns the FreeBlockCount field if non-nil, zero value otherwise.

### GetFreeBlockCountOk

`func (o *PoolAbstractBlock) GetFreeBlockCountOk() (*int64, bool)`

GetFreeBlockCountOk returns a tuple with the FreeBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBlockCount

`func (o *PoolAbstractBlock) SetFreeBlockCount(v int64)`

SetFreeBlockCount sets FreeBlockCount field to given value.

### HasFreeBlockCount

`func (o *PoolAbstractBlock) HasFreeBlockCount() bool`

HasFreeBlockCount returns a boolean if a field has been set.

### GetNextIdAllocator

`func (o *PoolAbstractBlock) GetNextIdAllocator() int64`

GetNextIdAllocator returns the NextIdAllocator field if non-nil, zero value otherwise.

### GetNextIdAllocatorOk

`func (o *PoolAbstractBlock) GetNextIdAllocatorOk() (*int64, bool)`

GetNextIdAllocatorOk returns a tuple with the NextIdAllocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextIdAllocator

`func (o *PoolAbstractBlock) SetNextIdAllocator(v int64)`

SetNextIdAllocator sets NextIdAllocator field to given value.

### HasNextIdAllocator

`func (o *PoolAbstractBlock) HasNextIdAllocator() bool`

HasNextIdAllocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


