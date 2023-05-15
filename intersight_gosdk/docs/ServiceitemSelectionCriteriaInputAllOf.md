# ServiceitemSelectionCriteriaInputAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "serviceitem.SelectionCriteriaInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "serviceitem.SelectionCriteriaInput"]
**FilterConditions** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**InputName** | Pointer to **string** | Name of the Policy Input. | [optional] [readonly] 
**InputValue** | Pointer to **interface{}** | The value extracted from the filter conditions. | [optional] [readonly] 

## Methods

### NewServiceitemSelectionCriteriaInputAllOf

`func NewServiceitemSelectionCriteriaInputAllOf(classId string, objectType string, ) *ServiceitemSelectionCriteriaInputAllOf`

NewServiceitemSelectionCriteriaInputAllOf instantiates a new ServiceitemSelectionCriteriaInputAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceitemSelectionCriteriaInputAllOfWithDefaults

`func NewServiceitemSelectionCriteriaInputAllOfWithDefaults() *ServiceitemSelectionCriteriaInputAllOf`

NewServiceitemSelectionCriteriaInputAllOfWithDefaults instantiates a new ServiceitemSelectionCriteriaInputAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilterConditions

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetFilterConditions() []ResourceSelector`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetFilterConditionsOk() (*[]ResourceSelector, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetFilterConditions(v []ResourceSelector)`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *ServiceitemSelectionCriteriaInputAllOf) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### SetFilterConditionsNil

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetFilterConditionsNil(b bool)`

 SetFilterConditionsNil sets the value for FilterConditions to be an explicit nil

### UnsetFilterConditions
`func (o *ServiceitemSelectionCriteriaInputAllOf) UnsetFilterConditions()`

UnsetFilterConditions ensures that no value is present for FilterConditions, not even an explicit nil
### GetInputName

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *ServiceitemSelectionCriteriaInputAllOf) HasInputName() bool`

HasInputName returns a boolean if a field has been set.

### GetInputValue

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetInputValue() interface{}`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ServiceitemSelectionCriteriaInputAllOf) GetInputValueOk() (*interface{}, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetInputValue(v interface{})`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *ServiceitemSelectionCriteriaInputAllOf) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *ServiceitemSelectionCriteriaInputAllOf) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *ServiceitemSelectionCriteriaInputAllOf) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


