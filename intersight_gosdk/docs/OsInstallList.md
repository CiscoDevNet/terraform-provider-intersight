# OsInstallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;os.Install&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]OsInstall**](OsInstall.md) | The array of &#39;os.Install&#39; resources matching the request. | [optional] 

## Methods

### NewOsInstallList

`func NewOsInstallList() *OsInstallList`

NewOsInstallList instantiates a new OsInstallList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsInstallListWithDefaults

`func NewOsInstallListWithDefaults() *OsInstallList`

NewOsInstallListWithDefaults instantiates a new OsInstallList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OsInstallList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OsInstallList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OsInstallList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OsInstallList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *OsInstallList) GetResults() []OsInstall`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OsInstallList) GetResultsOk() (*[]OsInstall, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OsInstallList) SetResults(v []OsInstall)`

SetResults sets Results field to given value.

### HasResults

`func (o *OsInstallList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *OsInstallList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *OsInstallList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


