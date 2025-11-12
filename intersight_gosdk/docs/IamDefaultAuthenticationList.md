# IamDefaultAuthenticationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;iam.DefaultAuthentication&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]IamDefaultAuthentication**](IamDefaultAuthentication.md) | The array of &#39;iam.DefaultAuthentication&#39; resources matching the request. | [optional] 

## Methods

### NewIamDefaultAuthenticationList

`func NewIamDefaultAuthenticationList() *IamDefaultAuthenticationList`

NewIamDefaultAuthenticationList instantiates a new IamDefaultAuthenticationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamDefaultAuthenticationListWithDefaults

`func NewIamDefaultAuthenticationListWithDefaults() *IamDefaultAuthenticationList`

NewIamDefaultAuthenticationListWithDefaults instantiates a new IamDefaultAuthenticationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IamDefaultAuthenticationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IamDefaultAuthenticationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IamDefaultAuthenticationList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IamDefaultAuthenticationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *IamDefaultAuthenticationList) GetResults() []IamDefaultAuthentication`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IamDefaultAuthenticationList) GetResultsOk() (*[]IamDefaultAuthentication, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IamDefaultAuthenticationList) SetResults(v []IamDefaultAuthentication)`

SetResults sets Results field to given value.

### HasResults

`func (o *IamDefaultAuthenticationList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *IamDefaultAuthenticationList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *IamDefaultAuthenticationList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


