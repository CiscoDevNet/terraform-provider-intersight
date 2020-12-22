# StorageEnclosureDiskListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;storage.EnclosureDisk&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]StorageEnclosureDisk**](StorageEnclosureDisk.md) | The array of &#39;storage.EnclosureDisk&#39; resources matching the request. | [optional] 

## Methods

### NewStorageEnclosureDiskListAllOf

`func NewStorageEnclosureDiskListAllOf() *StorageEnclosureDiskListAllOf`

NewStorageEnclosureDiskListAllOf instantiates a new StorageEnclosureDiskListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnclosureDiskListAllOfWithDefaults

`func NewStorageEnclosureDiskListAllOfWithDefaults() *StorageEnclosureDiskListAllOf`

NewStorageEnclosureDiskListAllOfWithDefaults instantiates a new StorageEnclosureDiskListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StorageEnclosureDiskListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StorageEnclosureDiskListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StorageEnclosureDiskListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *StorageEnclosureDiskListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *StorageEnclosureDiskListAllOf) GetResults() []StorageEnclosureDisk`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StorageEnclosureDiskListAllOf) GetResultsOk() (*[]StorageEnclosureDisk, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StorageEnclosureDiskListAllOf) SetResults(v []StorageEnclosureDisk)`

SetResults sets Results field to given value.

### HasResults

`func (o *StorageEnclosureDiskListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *StorageEnclosureDiskListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *StorageEnclosureDiskListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


