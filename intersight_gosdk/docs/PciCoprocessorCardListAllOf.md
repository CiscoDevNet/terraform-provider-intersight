# PciCoprocessorCardListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;pci.CoprocessorCard&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]PciCoprocessorCard**](pci.CoprocessorCard.md) | The array of &#39;pci.CoprocessorCard&#39; resources matching the request. | [optional] 

## Methods

### NewPciCoprocessorCardListAllOf

`func NewPciCoprocessorCardListAllOf() *PciCoprocessorCardListAllOf`

NewPciCoprocessorCardListAllOf instantiates a new PciCoprocessorCardListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciCoprocessorCardListAllOfWithDefaults

`func NewPciCoprocessorCardListAllOfWithDefaults() *PciCoprocessorCardListAllOf`

NewPciCoprocessorCardListAllOfWithDefaults instantiates a new PciCoprocessorCardListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PciCoprocessorCardListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PciCoprocessorCardListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PciCoprocessorCardListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PciCoprocessorCardListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PciCoprocessorCardListAllOf) GetResults() []PciCoprocessorCard`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PciCoprocessorCardListAllOf) GetResultsOk() (*[]PciCoprocessorCard, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PciCoprocessorCardListAllOf) SetResults(v []PciCoprocessorCard)`

SetResults sets Results field to given value.

### HasResults

`func (o *PciCoprocessorCardListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *PciCoprocessorCardListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *PciCoprocessorCardListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


