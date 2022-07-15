# ConvergedinfraAdapterComplianceDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;convergedinfra.AdapterComplianceDetails&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]ConvergedinfraAdapterComplianceDetails**](ConvergedinfraAdapterComplianceDetails.md) | The array of &#39;convergedinfra.AdapterComplianceDetails&#39; resources matching the request. | [optional] 

## Methods

### NewConvergedinfraAdapterComplianceDetailsList

`func NewConvergedinfraAdapterComplianceDetailsList() *ConvergedinfraAdapterComplianceDetailsList`

NewConvergedinfraAdapterComplianceDetailsList instantiates a new ConvergedinfraAdapterComplianceDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraAdapterComplianceDetailsListWithDefaults

`func NewConvergedinfraAdapterComplianceDetailsListWithDefaults() *ConvergedinfraAdapterComplianceDetailsList`

NewConvergedinfraAdapterComplianceDetailsListWithDefaults instantiates a new ConvergedinfraAdapterComplianceDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ConvergedinfraAdapterComplianceDetailsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConvergedinfraAdapterComplianceDetailsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConvergedinfraAdapterComplianceDetailsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ConvergedinfraAdapterComplianceDetailsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *ConvergedinfraAdapterComplianceDetailsList) GetResults() []ConvergedinfraAdapterComplianceDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ConvergedinfraAdapterComplianceDetailsList) GetResultsOk() (*[]ConvergedinfraAdapterComplianceDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ConvergedinfraAdapterComplianceDetailsList) SetResults(v []ConvergedinfraAdapterComplianceDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *ConvergedinfraAdapterComplianceDetailsList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *ConvergedinfraAdapterComplianceDetailsList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *ConvergedinfraAdapterComplianceDetailsList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


