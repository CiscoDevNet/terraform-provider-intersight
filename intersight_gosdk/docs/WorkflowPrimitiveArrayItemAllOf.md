# WorkflowPrimitiveArrayItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.PrimitiveArrayItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.PrimitiveArrayItem"]
**Properties** | Pointer to [**NullableWorkflowPrimitiveDataProperty**](WorkflowPrimitiveDataProperty.md) |  | [optional] 

## Methods

### NewWorkflowPrimitiveArrayItemAllOf

`func NewWorkflowPrimitiveArrayItemAllOf(classId string, objectType string, ) *WorkflowPrimitiveArrayItemAllOf`

NewWorkflowPrimitiveArrayItemAllOf instantiates a new WorkflowPrimitiveArrayItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPrimitiveArrayItemAllOfWithDefaults

`func NewWorkflowPrimitiveArrayItemAllOfWithDefaults() *WorkflowPrimitiveArrayItemAllOf`

NewWorkflowPrimitiveArrayItemAllOfWithDefaults instantiates a new WorkflowPrimitiveArrayItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPrimitiveArrayItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPrimitiveArrayItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPrimitiveArrayItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPrimitiveArrayItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPrimitiveArrayItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPrimitiveArrayItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowPrimitiveArrayItemAllOf) GetProperties() WorkflowPrimitiveDataProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowPrimitiveArrayItemAllOf) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowPrimitiveArrayItemAllOf) SetProperties(v WorkflowPrimitiveDataProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowPrimitiveArrayItemAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowPrimitiveArrayItemAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowPrimitiveArrayItemAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


