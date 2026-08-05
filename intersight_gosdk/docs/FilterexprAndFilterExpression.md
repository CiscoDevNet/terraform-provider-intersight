# FilterexprAndFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "filterexpr.AndFilterExpression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "filterexpr.AndFilterExpression"]
**Fields** | Pointer to [**[]FilterexprFilterExpression**](FilterexprFilterExpression.md) |  | [optional] 

## Methods

### NewFilterexprAndFilterExpression

`func NewFilterexprAndFilterExpression(classId string, objectType string, ) *FilterexprAndFilterExpression`

NewFilterexprAndFilterExpression instantiates a new FilterexprAndFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterexprAndFilterExpressionWithDefaults

`func NewFilterexprAndFilterExpressionWithDefaults() *FilterexprAndFilterExpression`

NewFilterexprAndFilterExpressionWithDefaults instantiates a new FilterexprAndFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FilterexprAndFilterExpression) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FilterexprAndFilterExpression) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FilterexprAndFilterExpression) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FilterexprAndFilterExpression) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FilterexprAndFilterExpression) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FilterexprAndFilterExpression) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFields

`func (o *FilterexprAndFilterExpression) GetFields() []FilterexprFilterExpression`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FilterexprAndFilterExpression) GetFieldsOk() (*[]FilterexprFilterExpression, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FilterexprAndFilterExpression) SetFields(v []FilterexprFilterExpression)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FilterexprAndFilterExpression) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *FilterexprAndFilterExpression) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *FilterexprAndFilterExpression) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


