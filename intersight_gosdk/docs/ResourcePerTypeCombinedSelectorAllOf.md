# ResourcePerTypeCombinedSelectorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.PerTypeCombinedSelector"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.PerTypeCombinedSelector"]
**CombinedSelector** | Pointer to **string** | A single filter expression created by OR&#39;ing the $filter criteria of the &#39;selectors&#39;. Used to efficiently maintain the membership of the Group. | [optional] [readonly] 
**EmptyFilter** | Pointer to **bool** | If true, then resources are added using just object type without filter. | [optional] [readonly] 
**SelectorObjectType** | Pointer to **string** | The ObjectType on which the selectors are defined. Used to efficiently query resource groups for a given ObjectType. | [optional] [readonly] 

## Methods

### NewResourcePerTypeCombinedSelectorAllOf

`func NewResourcePerTypeCombinedSelectorAllOf(classId string, objectType string, ) *ResourcePerTypeCombinedSelectorAllOf`

NewResourcePerTypeCombinedSelectorAllOf instantiates a new ResourcePerTypeCombinedSelectorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePerTypeCombinedSelectorAllOfWithDefaults

`func NewResourcePerTypeCombinedSelectorAllOfWithDefaults() *ResourcePerTypeCombinedSelectorAllOf`

NewResourcePerTypeCombinedSelectorAllOfWithDefaults instantiates a new ResourcePerTypeCombinedSelectorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcePerTypeCombinedSelectorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcePerTypeCombinedSelectorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCombinedSelector

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetCombinedSelector() string`

GetCombinedSelector returns the CombinedSelector field if non-nil, zero value otherwise.

### GetCombinedSelectorOk

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetCombinedSelectorOk() (*string, bool)`

GetCombinedSelectorOk returns a tuple with the CombinedSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedSelector

`func (o *ResourcePerTypeCombinedSelectorAllOf) SetCombinedSelector(v string)`

SetCombinedSelector sets CombinedSelector field to given value.

### HasCombinedSelector

`func (o *ResourcePerTypeCombinedSelectorAllOf) HasCombinedSelector() bool`

HasCombinedSelector returns a boolean if a field has been set.

### GetEmptyFilter

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetEmptyFilter() bool`

GetEmptyFilter returns the EmptyFilter field if non-nil, zero value otherwise.

### GetEmptyFilterOk

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetEmptyFilterOk() (*bool, bool)`

GetEmptyFilterOk returns a tuple with the EmptyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyFilter

`func (o *ResourcePerTypeCombinedSelectorAllOf) SetEmptyFilter(v bool)`

SetEmptyFilter sets EmptyFilter field to given value.

### HasEmptyFilter

`func (o *ResourcePerTypeCombinedSelectorAllOf) HasEmptyFilter() bool`

HasEmptyFilter returns a boolean if a field has been set.

### GetSelectorObjectType

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetSelectorObjectType() string`

GetSelectorObjectType returns the SelectorObjectType field if non-nil, zero value otherwise.

### GetSelectorObjectTypeOk

`func (o *ResourcePerTypeCombinedSelectorAllOf) GetSelectorObjectTypeOk() (*string, bool)`

GetSelectorObjectTypeOk returns a tuple with the SelectorObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorObjectType

`func (o *ResourcePerTypeCombinedSelectorAllOf) SetSelectorObjectType(v string)`

SetSelectorObjectType sets SelectorObjectType field to given value.

### HasSelectorObjectType

`func (o *ResourcePerTypeCombinedSelectorAllOf) HasSelectorObjectType() bool`

HasSelectorObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


