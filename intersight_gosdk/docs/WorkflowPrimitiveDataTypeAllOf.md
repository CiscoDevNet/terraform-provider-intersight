# WorkflowPrimitiveDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.PrimitiveDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.PrimitiveDataType"]
**Properties** | Pointer to [**NullableWorkflowPrimitiveDataProperty**](WorkflowPrimitiveDataProperty.md) |  | [optional] 

## Methods

### NewWorkflowPrimitiveDataTypeAllOf

`func NewWorkflowPrimitiveDataTypeAllOf(classId string, objectType string, ) *WorkflowPrimitiveDataTypeAllOf`

NewWorkflowPrimitiveDataTypeAllOf instantiates a new WorkflowPrimitiveDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPrimitiveDataTypeAllOfWithDefaults

`func NewWorkflowPrimitiveDataTypeAllOfWithDefaults() *WorkflowPrimitiveDataTypeAllOf`

NewWorkflowPrimitiveDataTypeAllOfWithDefaults instantiates a new WorkflowPrimitiveDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPrimitiveDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPrimitiveDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPrimitiveDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPrimitiveDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPrimitiveDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPrimitiveDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowPrimitiveDataTypeAllOf) GetProperties() WorkflowPrimitiveDataProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowPrimitiveDataTypeAllOf) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowPrimitiveDataTypeAllOf) SetProperties(v WorkflowPrimitiveDataProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowPrimitiveDataTypeAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowPrimitiveDataTypeAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowPrimitiveDataTypeAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


