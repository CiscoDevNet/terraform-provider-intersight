# HyperflexVcenterConfigPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;hyperflex.VcenterConfigPolicy&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]MoTagKeySummary**](MoTagKeySummary.md) |  | [optional] 

## Methods

### NewHyperflexVcenterConfigPolicyResponse

`func NewHyperflexVcenterConfigPolicyResponse(objectType string, ) *HyperflexVcenterConfigPolicyResponse`

NewHyperflexVcenterConfigPolicyResponse instantiates a new HyperflexVcenterConfigPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVcenterConfigPolicyResponseWithDefaults

`func NewHyperflexVcenterConfigPolicyResponseWithDefaults() *HyperflexVcenterConfigPolicyResponse`

NewHyperflexVcenterConfigPolicyResponseWithDefaults instantiates a new HyperflexVcenterConfigPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *HyperflexVcenterConfigPolicyResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVcenterConfigPolicyResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVcenterConfigPolicyResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *HyperflexVcenterConfigPolicyResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HyperflexVcenterConfigPolicyResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HyperflexVcenterConfigPolicyResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HyperflexVcenterConfigPolicyResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *HyperflexVcenterConfigPolicyResponse) GetResults() []MoTagKeySummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HyperflexVcenterConfigPolicyResponse) GetResultsOk() (*[]MoTagKeySummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HyperflexVcenterConfigPolicyResponse) SetResults(v []MoTagKeySummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *HyperflexVcenterConfigPolicyResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *HyperflexVcenterConfigPolicyResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *HyperflexVcenterConfigPolicyResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


