# ServiceitemSelectionCriteriaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "serviceitem.SelectionCriteriaInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "serviceitem.SelectionCriteriaInput"]
**FilterConditions** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 
**InputName** | Pointer to **string** | Name of the Policy Input. | [optional] [readonly] 
**InputValue** | Pointer to **interface{}** | The value extracted from the filter conditions. | [optional] [readonly] 

## Methods

### NewServiceitemSelectionCriteriaInput

`func NewServiceitemSelectionCriteriaInput(classId string, objectType string, ) *ServiceitemSelectionCriteriaInput`

NewServiceitemSelectionCriteriaInput instantiates a new ServiceitemSelectionCriteriaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceitemSelectionCriteriaInputWithDefaults

`func NewServiceitemSelectionCriteriaInputWithDefaults() *ServiceitemSelectionCriteriaInput`

NewServiceitemSelectionCriteriaInputWithDefaults instantiates a new ServiceitemSelectionCriteriaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServiceitemSelectionCriteriaInput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServiceitemSelectionCriteriaInput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServiceitemSelectionCriteriaInput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServiceitemSelectionCriteriaInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServiceitemSelectionCriteriaInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServiceitemSelectionCriteriaInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilterConditions

`func (o *ServiceitemSelectionCriteriaInput) GetFilterConditions() []ResourceSelector`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *ServiceitemSelectionCriteriaInput) GetFilterConditionsOk() (*[]ResourceSelector, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *ServiceitemSelectionCriteriaInput) SetFilterConditions(v []ResourceSelector)`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *ServiceitemSelectionCriteriaInput) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### SetFilterConditionsNil

`func (o *ServiceitemSelectionCriteriaInput) SetFilterConditionsNil(b bool)`

 SetFilterConditionsNil sets the value for FilterConditions to be an explicit nil

### UnsetFilterConditions
`func (o *ServiceitemSelectionCriteriaInput) UnsetFilterConditions()`

UnsetFilterConditions ensures that no value is present for FilterConditions, not even an explicit nil
### GetInputName

`func (o *ServiceitemSelectionCriteriaInput) GetInputName() string`

GetInputName returns the InputName field if non-nil, zero value otherwise.

### GetInputNameOk

`func (o *ServiceitemSelectionCriteriaInput) GetInputNameOk() (*string, bool)`

GetInputNameOk returns a tuple with the InputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputName

`func (o *ServiceitemSelectionCriteriaInput) SetInputName(v string)`

SetInputName sets InputName field to given value.

### HasInputName

`func (o *ServiceitemSelectionCriteriaInput) HasInputName() bool`

HasInputName returns a boolean if a field has been set.

### GetInputValue

`func (o *ServiceitemSelectionCriteriaInput) GetInputValue() interface{}`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ServiceitemSelectionCriteriaInput) GetInputValueOk() (*interface{}, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ServiceitemSelectionCriteriaInput) SetInputValue(v interface{})`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *ServiceitemSelectionCriteriaInput) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *ServiceitemSelectionCriteriaInput) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *ServiceitemSelectionCriteriaInput) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


