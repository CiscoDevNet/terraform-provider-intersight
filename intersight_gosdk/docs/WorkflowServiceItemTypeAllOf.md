# WorkflowServiceItemTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemType"]
**Description** | Pointer to **string** | The description of this service item. | [optional] 
**Name** | Pointer to **string** | The name of the service item as defined by the user. | [optional] 
**ServiceItemDefinition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**ServiceItemInputDefinition** | Pointer to [**[]WorkflowServiceItemInputDefinitionType**](WorkflowServiceItemInputDefinitionType.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemTypeAllOf

`func NewWorkflowServiceItemTypeAllOf(classId string, objectType string, ) *WorkflowServiceItemTypeAllOf`

NewWorkflowServiceItemTypeAllOf instantiates a new WorkflowServiceItemTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemTypeAllOfWithDefaults

`func NewWorkflowServiceItemTypeAllOfWithDefaults() *WorkflowServiceItemTypeAllOf`

NewWorkflowServiceItemTypeAllOfWithDefaults instantiates a new WorkflowServiceItemTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowServiceItemTypeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemTypeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemTypeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemTypeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemTypeAllOf) GetServiceItemDefinition() MoMoRef`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemTypeAllOf) GetServiceItemDefinitionOk() (*MoMoRef, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemTypeAllOf) SetServiceItemDefinition(v MoMoRef)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemTypeAllOf) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### GetServiceItemInputDefinition

`func (o *WorkflowServiceItemTypeAllOf) GetServiceItemInputDefinition() []WorkflowServiceItemInputDefinitionType`

GetServiceItemInputDefinition returns the ServiceItemInputDefinition field if non-nil, zero value otherwise.

### GetServiceItemInputDefinitionOk

`func (o *WorkflowServiceItemTypeAllOf) GetServiceItemInputDefinitionOk() (*[]WorkflowServiceItemInputDefinitionType, bool)`

GetServiceItemInputDefinitionOk returns a tuple with the ServiceItemInputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInputDefinition

`func (o *WorkflowServiceItemTypeAllOf) SetServiceItemInputDefinition(v []WorkflowServiceItemInputDefinitionType)`

SetServiceItemInputDefinition sets ServiceItemInputDefinition field to given value.

### HasServiceItemInputDefinition

`func (o *WorkflowServiceItemTypeAllOf) HasServiceItemInputDefinition() bool`

HasServiceItemInputDefinition returns a boolean if a field has been set.

### SetServiceItemInputDefinitionNil

`func (o *WorkflowServiceItemTypeAllOf) SetServiceItemInputDefinitionNil(b bool)`

 SetServiceItemInputDefinitionNil sets the value for ServiceItemInputDefinition to be an explicit nil

### UnsetServiceItemInputDefinition
`func (o *WorkflowServiceItemTypeAllOf) UnsetServiceItemInputDefinition()`

UnsetServiceItemInputDefinition ensures that no value is present for ServiceItemInputDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


