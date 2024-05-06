# FabricSwitchProfileTemplateListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;fabric.SwitchProfileTemplate&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]FabricSwitchProfileTemplate**](FabricSwitchProfileTemplate.md) | The array of &#39;fabric.SwitchProfileTemplate&#39; resources matching the request. | [optional] 

## Methods

### NewFabricSwitchProfileTemplateListAllOf

`func NewFabricSwitchProfileTemplateListAllOf() *FabricSwitchProfileTemplateListAllOf`

NewFabricSwitchProfileTemplateListAllOf instantiates a new FabricSwitchProfileTemplateListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchProfileTemplateListAllOfWithDefaults

`func NewFabricSwitchProfileTemplateListAllOfWithDefaults() *FabricSwitchProfileTemplateListAllOf`

NewFabricSwitchProfileTemplateListAllOfWithDefaults instantiates a new FabricSwitchProfileTemplateListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FabricSwitchProfileTemplateListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FabricSwitchProfileTemplateListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FabricSwitchProfileTemplateListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FabricSwitchProfileTemplateListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *FabricSwitchProfileTemplateListAllOf) GetResults() []FabricSwitchProfileTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FabricSwitchProfileTemplateListAllOf) GetResultsOk() (*[]FabricSwitchProfileTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FabricSwitchProfileTemplateListAllOf) SetResults(v []FabricSwitchProfileTemplate)`

SetResults sets Results field to given value.

### HasResults

`func (o *FabricSwitchProfileTemplateListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *FabricSwitchProfileTemplateListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *FabricSwitchProfileTemplateListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


