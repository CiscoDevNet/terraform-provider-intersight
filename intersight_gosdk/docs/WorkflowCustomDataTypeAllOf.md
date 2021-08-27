# WorkflowCustomDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CustomDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CustomDataType"]
**Properties** | Pointer to [**NullableWorkflowCustomDataProperty**](WorkflowCustomDataProperty.md) |  | [optional] 

## Methods

### NewWorkflowCustomDataTypeAllOf

`func NewWorkflowCustomDataTypeAllOf(classId string, objectType string, ) *WorkflowCustomDataTypeAllOf`

NewWorkflowCustomDataTypeAllOf instantiates a new WorkflowCustomDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCustomDataTypeAllOfWithDefaults

`func NewWorkflowCustomDataTypeAllOfWithDefaults() *WorkflowCustomDataTypeAllOf`

NewWorkflowCustomDataTypeAllOfWithDefaults instantiates a new WorkflowCustomDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCustomDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCustomDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCustomDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCustomDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCustomDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCustomDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowCustomDataTypeAllOf) GetProperties() WorkflowCustomDataProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowCustomDataTypeAllOf) GetPropertiesOk() (*WorkflowCustomDataProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowCustomDataTypeAllOf) SetProperties(v WorkflowCustomDataProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowCustomDataTypeAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowCustomDataTypeAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowCustomDataTypeAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


