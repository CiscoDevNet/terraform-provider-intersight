# WorkflowMoReferenceDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoReferenceDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoReferenceDataType"]
**Properties** | Pointer to [**[]WorkflowMoReferenceProperty**](WorkflowMoReferenceProperty.md) |  | [optional] 

## Methods

### NewWorkflowMoReferenceDataTypeAllOf

`func NewWorkflowMoReferenceDataTypeAllOf(classId string, objectType string, ) *WorkflowMoReferenceDataTypeAllOf`

NewWorkflowMoReferenceDataTypeAllOf instantiates a new WorkflowMoReferenceDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoReferenceDataTypeAllOfWithDefaults

`func NewWorkflowMoReferenceDataTypeAllOfWithDefaults() *WorkflowMoReferenceDataTypeAllOf`

NewWorkflowMoReferenceDataTypeAllOfWithDefaults instantiates a new WorkflowMoReferenceDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoReferenceDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoReferenceDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoReferenceDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoReferenceDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoReferenceDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoReferenceDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowMoReferenceDataTypeAllOf) GetProperties() []WorkflowMoReferenceProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowMoReferenceDataTypeAllOf) GetPropertiesOk() (*[]WorkflowMoReferenceProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowMoReferenceDataTypeAllOf) SetProperties(v []WorkflowMoReferenceProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowMoReferenceDataTypeAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowMoReferenceDataTypeAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowMoReferenceDataTypeAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


