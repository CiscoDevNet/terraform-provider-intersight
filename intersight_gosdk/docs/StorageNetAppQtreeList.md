# StorageNetAppQtreeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;storage.NetAppQtree&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]StorageNetAppQtree**](StorageNetAppQtree.md) | The array of &#39;storage.NetAppQtree&#39; resources matching the request. | [optional] 

## Methods

### NewStorageNetAppQtreeList

`func NewStorageNetAppQtreeList() *StorageNetAppQtreeList`

NewStorageNetAppQtreeList instantiates a new StorageNetAppQtreeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppQtreeListWithDefaults

`func NewStorageNetAppQtreeListWithDefaults() *StorageNetAppQtreeList`

NewStorageNetAppQtreeListWithDefaults instantiates a new StorageNetAppQtreeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StorageNetAppQtreeList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StorageNetAppQtreeList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StorageNetAppQtreeList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *StorageNetAppQtreeList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *StorageNetAppQtreeList) GetResults() []StorageNetAppQtree`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StorageNetAppQtreeList) GetResultsOk() (*[]StorageNetAppQtree, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StorageNetAppQtreeList) SetResults(v []StorageNetAppQtree)`

SetResults sets Results field to given value.

### HasResults

`func (o *StorageNetAppQtreeList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *StorageNetAppQtreeList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *StorageNetAppQtreeList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


