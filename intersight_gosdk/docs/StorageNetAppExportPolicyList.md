# StorageNetAppExportPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;storage.NetAppExportPolicy&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]StorageNetAppExportPolicy**](StorageNetAppExportPolicy.md) | The array of &#39;storage.NetAppExportPolicy&#39; resources matching the request. | [optional] 

## Methods

### NewStorageNetAppExportPolicyList

`func NewStorageNetAppExportPolicyList() *StorageNetAppExportPolicyList`

NewStorageNetAppExportPolicyList instantiates a new StorageNetAppExportPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppExportPolicyListWithDefaults

`func NewStorageNetAppExportPolicyListWithDefaults() *StorageNetAppExportPolicyList`

NewStorageNetAppExportPolicyListWithDefaults instantiates a new StorageNetAppExportPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StorageNetAppExportPolicyList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StorageNetAppExportPolicyList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StorageNetAppExportPolicyList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *StorageNetAppExportPolicyList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *StorageNetAppExportPolicyList) GetResults() []StorageNetAppExportPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StorageNetAppExportPolicyList) GetResultsOk() (*[]StorageNetAppExportPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StorageNetAppExportPolicyList) SetResults(v []StorageNetAppExportPolicy)`

SetResults sets Results field to given value.

### HasResults

`func (o *StorageNetAppExportPolicyList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *StorageNetAppExportPolicyList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *StorageNetAppExportPolicyList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


