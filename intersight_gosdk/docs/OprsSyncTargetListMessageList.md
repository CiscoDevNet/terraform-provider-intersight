# OprsSyncTargetListMessageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;oprs.SyncTargetListMessage&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]OprsSyncTargetListMessage**](OprsSyncTargetListMessage.md) | The array of &#39;oprs.SyncTargetListMessage&#39; resources matching the request. | [optional] 

## Methods

### NewOprsSyncTargetListMessageList

`func NewOprsSyncTargetListMessageList() *OprsSyncTargetListMessageList`

NewOprsSyncTargetListMessageList instantiates a new OprsSyncTargetListMessageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOprsSyncTargetListMessageListWithDefaults

`func NewOprsSyncTargetListMessageListWithDefaults() *OprsSyncTargetListMessageList`

NewOprsSyncTargetListMessageListWithDefaults instantiates a new OprsSyncTargetListMessageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OprsSyncTargetListMessageList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OprsSyncTargetListMessageList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OprsSyncTargetListMessageList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OprsSyncTargetListMessageList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *OprsSyncTargetListMessageList) GetResults() []OprsSyncTargetListMessage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OprsSyncTargetListMessageList) GetResultsOk() (*[]OprsSyncTargetListMessage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OprsSyncTargetListMessageList) SetResults(v []OprsSyncTargetListMessage)`

SetResults sets Results field to given value.

### HasResults

`func (o *OprsSyncTargetListMessageList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *OprsSyncTargetListMessageList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *OprsSyncTargetListMessageList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


