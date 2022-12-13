# NiaapiPuvScriptDownloaderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;niaapi.PuvScriptDownloader&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]NiaapiPuvScriptDownloader**](NiaapiPuvScriptDownloader.md) | The array of &#39;niaapi.PuvScriptDownloader&#39; resources matching the request. | [optional] 

## Methods

### NewNiaapiPuvScriptDownloaderList

`func NewNiaapiPuvScriptDownloaderList() *NiaapiPuvScriptDownloaderList`

NewNiaapiPuvScriptDownloaderList instantiates a new NiaapiPuvScriptDownloaderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiPuvScriptDownloaderListWithDefaults

`func NewNiaapiPuvScriptDownloaderListWithDefaults() *NiaapiPuvScriptDownloaderList`

NewNiaapiPuvScriptDownloaderListWithDefaults instantiates a new NiaapiPuvScriptDownloaderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NiaapiPuvScriptDownloaderList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NiaapiPuvScriptDownloaderList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NiaapiPuvScriptDownloaderList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NiaapiPuvScriptDownloaderList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *NiaapiPuvScriptDownloaderList) GetResults() []NiaapiPuvScriptDownloader`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NiaapiPuvScriptDownloaderList) GetResultsOk() (*[]NiaapiPuvScriptDownloader, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NiaapiPuvScriptDownloaderList) SetResults(v []NiaapiPuvScriptDownloader)`

SetResults sets Results field to given value.

### HasResults

`func (o *NiaapiPuvScriptDownloaderList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NiaapiPuvScriptDownloaderList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NiaapiPuvScriptDownloaderList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


