# ConvergedinfraSwitchComplianceDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;convergedinfra.SwitchComplianceDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]ConvergedinfraSwitchComplianceDetails**](ConvergedinfraSwitchComplianceDetails.md) | The array of &#39;convergedinfra.SwitchComplianceDetails&#39; resources matching the request. | [optional] 

## Methods

### NewConvergedinfraSwitchComplianceDetailsList

`func NewConvergedinfraSwitchComplianceDetailsList() *ConvergedinfraSwitchComplianceDetailsList`

NewConvergedinfraSwitchComplianceDetailsList instantiates a new ConvergedinfraSwitchComplianceDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraSwitchComplianceDetailsListWithDefaults

`func NewConvergedinfraSwitchComplianceDetailsListWithDefaults() *ConvergedinfraSwitchComplianceDetailsList`

NewConvergedinfraSwitchComplianceDetailsListWithDefaults instantiates a new ConvergedinfraSwitchComplianceDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ConvergedinfraSwitchComplianceDetailsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConvergedinfraSwitchComplianceDetailsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConvergedinfraSwitchComplianceDetailsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ConvergedinfraSwitchComplianceDetailsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ConvergedinfraSwitchComplianceDetailsList) GetResults() []ConvergedinfraSwitchComplianceDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConvergedinfraSwitchComplianceDetailsList) GetResultsOk() (*[]ConvergedinfraSwitchComplianceDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConvergedinfraSwitchComplianceDetailsList) SetResults(v []ConvergedinfraSwitchComplianceDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConvergedinfraSwitchComplianceDetailsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *ConvergedinfraSwitchComplianceDetailsList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ConvergedinfraSwitchComplianceDetailsList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


