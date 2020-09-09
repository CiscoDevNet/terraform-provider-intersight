# PoolAbstractBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeBlockCount** | Pointer to **int64** | Free IDs that can be allocated in this block. | [optional] [readonly] 
**NextIdAllocator** | Pointer to **int64** | Moving counter to allocate the next identifier. | [optional] [readonly] 

## Methods

### NewPoolAbstractBlock

`func NewPoolAbstractBlock() *PoolAbstractBlock`

NewPoolAbstractBlock instantiates a new PoolAbstractBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractBlockWithDefaults

`func NewPoolAbstractBlockWithDefaults() *PoolAbstractBlock`

NewPoolAbstractBlockWithDefaults instantiates a new PoolAbstractBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


