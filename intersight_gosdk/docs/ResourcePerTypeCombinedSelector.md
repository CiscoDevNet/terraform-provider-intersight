# ResourcePerTypeCombinedSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombinedSelector** | Pointer to **string** | A single filter expression created by OR&#39;ing the $filter criteria of the &#39;selectors&#39;. Used to efficiently maintain the membership of the Group. | [optional] [readonly] 
**EmptyFilter** | Pointer to **bool** | If true, then resources are added using just object type without filter. | [optional] [readonly] 
**SelectorObjectType** | Pointer to **string** | The ObjectType on which the selectors are defined. Used to efficiently query resource groups for a given ObjectType. | [optional] [readonly] 

## Methods

### NewResourcePerTypeCombinedSelector

`func NewResourcePerTypeCombinedSelector() *ResourcePerTypeCombinedSelector`

NewResourcePerTypeCombinedSelector instantiates a new ResourcePerTypeCombinedSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePerTypeCombinedSelectorWithDefaults

`func NewResourcePerTypeCombinedSelectorWithDefaults() *ResourcePerTypeCombinedSelector`

NewResourcePerTypeCombinedSelectorWithDefaults instantiates a new ResourcePerTypeCombinedSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinedSelector

`func (o *ResourcePerTypeCombinedSelector) GetCombinedSelector() string`

GetCombinedSelector returns the CombinedSelector field if non-nil, zero value otherwise.

### GetCombinedSelectorOk

`func (o *ResourcePerTypeCombinedSelector) GetCombinedSelectorOk() (*string, bool)`

GetCombinedSelectorOk returns a tuple with the CombinedSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedSelector

`func (o *ResourcePerTypeCombinedSelector) SetCombinedSelector(v string)`

SetCombinedSelector sets CombinedSelector field to given value.

### HasCombinedSelector

`func (o *ResourcePerTypeCombinedSelector) HasCombinedSelector() bool`

HasCombinedSelector returns a boolean if a field has been set.

### GetEmptyFilter

`func (o *ResourcePerTypeCombinedSelector) GetEmptyFilter() bool`

GetEmptyFilter returns the EmptyFilter field if non-nil, zero value otherwise.

### GetEmptyFilterOk

`func (o *ResourcePerTypeCombinedSelector) GetEmptyFilterOk() (*bool, bool)`

GetEmptyFilterOk returns a tuple with the EmptyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyFilter

`func (o *ResourcePerTypeCombinedSelector) SetEmptyFilter(v bool)`

SetEmptyFilter sets EmptyFilter field to given value.

### HasEmptyFilter

`func (o *ResourcePerTypeCombinedSelector) HasEmptyFilter() bool`

HasEmptyFilter returns a boolean if a field has been set.

### GetSelectorObjectType

`func (o *ResourcePerTypeCombinedSelector) GetSelectorObjectType() string`

GetSelectorObjectType returns the SelectorObjectType field if non-nil, zero value otherwise.

### GetSelectorObjectTypeOk

`func (o *ResourcePerTypeCombinedSelector) GetSelectorObjectTypeOk() (*string, bool)`

GetSelectorObjectTypeOk returns a tuple with the SelectorObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorObjectType

`func (o *ResourcePerTypeCombinedSelector) SetSelectorObjectType(v string)`

SetSelectorObjectType sets SelectorObjectType field to given value.

### HasSelectorObjectType

`func (o *ResourcePerTypeCombinedSelector) HasSelectorObjectType() bool`

HasSelectorObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


