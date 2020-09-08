# TechsupportmanagementTechSupportStatusList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;techsupportmanagement.TechSupportStatus&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]TechsupportmanagementTechSupportStatus**](techsupportmanagement.TechSupportStatus.md) | The array of &#39;techsupportmanagement.TechSupportStatus&#39; resources matching the request. | [optional] 

## Methods

### NewTechsupportmanagementTechSupportStatusList

`func NewTechsupportmanagementTechSupportStatusList() *TechsupportmanagementTechSupportStatusList`

NewTechsupportmanagementTechSupportStatusList instantiates a new TechsupportmanagementTechSupportStatusList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportStatusListWithDefaults

`func NewTechsupportmanagementTechSupportStatusListWithDefaults() *TechsupportmanagementTechSupportStatusList`

NewTechsupportmanagementTechSupportStatusListWithDefaults instantiates a new TechsupportmanagementTechSupportStatusList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TechsupportmanagementTechSupportStatusList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TechsupportmanagementTechSupportStatusList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TechsupportmanagementTechSupportStatusList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TechsupportmanagementTechSupportStatusList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *TechsupportmanagementTechSupportStatusList) GetResults() []TechsupportmanagementTechSupportStatus`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TechsupportmanagementTechSupportStatusList) GetResultsOk() (*[]TechsupportmanagementTechSupportStatus, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TechsupportmanagementTechSupportStatusList) SetResults(v []TechsupportmanagementTechSupportStatus)`

SetResults sets Results field to given value.

### HasResults

`func (o *TechsupportmanagementTechSupportStatusList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *TechsupportmanagementTechSupportStatusList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *TechsupportmanagementTechSupportStatusList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


