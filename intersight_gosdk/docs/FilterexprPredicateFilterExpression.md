# FilterexprPredicateFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "filterexpr.PredicateFilterExpression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "filterexpr.PredicateFilterExpression"]
**Operator** | Pointer to **string** | Comparison or set operator (for example eq, ne, contains, in); semantics are defined by the consuming service. | [optional] 
**Property** | Pointer to **string** | Property path for OData or the target query DSL (e.g. dotted paths and array keys when supported). | [optional] 
**PropertyReference** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Value** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFilterexprPredicateFilterExpression

`func NewFilterexprPredicateFilterExpression(classId string, objectType string, ) *FilterexprPredicateFilterExpression`

NewFilterexprPredicateFilterExpression instantiates a new FilterexprPredicateFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterexprPredicateFilterExpressionWithDefaults

`func NewFilterexprPredicateFilterExpressionWithDefaults() *FilterexprPredicateFilterExpression`

NewFilterexprPredicateFilterExpressionWithDefaults instantiates a new FilterexprPredicateFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FilterexprPredicateFilterExpression) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FilterexprPredicateFilterExpression) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FilterexprPredicateFilterExpression) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FilterexprPredicateFilterExpression) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FilterexprPredicateFilterExpression) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FilterexprPredicateFilterExpression) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperator

`func (o *FilterexprPredicateFilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterexprPredicateFilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterexprPredicateFilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *FilterexprPredicateFilterExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetProperty

`func (o *FilterexprPredicateFilterExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *FilterexprPredicateFilterExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *FilterexprPredicateFilterExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *FilterexprPredicateFilterExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetPropertyReference

`func (o *FilterexprPredicateFilterExpression) GetPropertyReference() []MoMoRef`

GetPropertyReference returns the PropertyReference field if non-nil, zero value otherwise.

### GetPropertyReferenceOk

`func (o *FilterexprPredicateFilterExpression) GetPropertyReferenceOk() (*[]MoMoRef, bool)`

GetPropertyReferenceOk returns a tuple with the PropertyReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyReference

`func (o *FilterexprPredicateFilterExpression) SetPropertyReference(v []MoMoRef)`

SetPropertyReference sets PropertyReference field to given value.

### HasPropertyReference

`func (o *FilterexprPredicateFilterExpression) HasPropertyReference() bool`

HasPropertyReference returns a boolean if a field has been set.

### SetPropertyReferenceNil

`func (o *FilterexprPredicateFilterExpression) SetPropertyReferenceNil(b bool)`

 SetPropertyReferenceNil sets the value for PropertyReference to be an explicit nil

### UnsetPropertyReference
`func (o *FilterexprPredicateFilterExpression) UnsetPropertyReference()`

UnsetPropertyReference ensures that no value is present for PropertyReference, not even an explicit nil
### GetValue

`func (o *FilterexprPredicateFilterExpression) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterexprPredicateFilterExpression) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterexprPredicateFilterExpression) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FilterexprPredicateFilterExpression) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FilterexprPredicateFilterExpression) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FilterexprPredicateFilterExpression) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


