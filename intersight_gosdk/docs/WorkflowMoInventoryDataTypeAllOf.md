# WorkflowMoInventoryDataTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.MoInventoryDataType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.MoInventoryDataType"]
**Properties** | Pointer to [**[]WorkflowMoInventoryProperty**](WorkflowMoInventoryProperty.md) |  | [optional] 

## Methods

### NewWorkflowMoInventoryDataTypeAllOf

`func NewWorkflowMoInventoryDataTypeAllOf(classId string, objectType string, ) *WorkflowMoInventoryDataTypeAllOf`

NewWorkflowMoInventoryDataTypeAllOf instantiates a new WorkflowMoInventoryDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMoInventoryDataTypeAllOfWithDefaults

`func NewWorkflowMoInventoryDataTypeAllOfWithDefaults() *WorkflowMoInventoryDataTypeAllOf`

NewWorkflowMoInventoryDataTypeAllOfWithDefaults instantiates a new WorkflowMoInventoryDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMoInventoryDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMoInventoryDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMoInventoryDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMoInventoryDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMoInventoryDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMoInventoryDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *WorkflowMoInventoryDataTypeAllOf) GetProperties() []WorkflowMoInventoryProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowMoInventoryDataTypeAllOf) GetPropertiesOk() (*[]WorkflowMoInventoryProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowMoInventoryDataTypeAllOf) SetProperties(v []WorkflowMoInventoryProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowMoInventoryDataTypeAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowMoInventoryDataTypeAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowMoInventoryDataTypeAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


