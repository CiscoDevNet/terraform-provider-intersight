# WorkflowServiceItemType

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

### NewWorkflowServiceItemType

`func NewWorkflowServiceItemType(classId string, objectType string, ) *WorkflowServiceItemType`

NewWorkflowServiceItemType instantiates a new WorkflowServiceItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemTypeWithDefaults

`func NewWorkflowServiceItemTypeWithDefaults() *WorkflowServiceItemType`

NewWorkflowServiceItemTypeWithDefaults instantiates a new WorkflowServiceItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowServiceItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemType) GetServiceItemDefinition() MoMoRef`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemType) GetServiceItemDefinitionOk() (*MoMoRef, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemType) SetServiceItemDefinition(v MoMoRef)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemType) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### GetServiceItemInputDefinition

`func (o *WorkflowServiceItemType) GetServiceItemInputDefinition() []WorkflowServiceItemInputDefinitionType`

GetServiceItemInputDefinition returns the ServiceItemInputDefinition field if non-nil, zero value otherwise.

### GetServiceItemInputDefinitionOk

`func (o *WorkflowServiceItemType) GetServiceItemInputDefinitionOk() (*[]WorkflowServiceItemInputDefinitionType, bool)`

GetServiceItemInputDefinitionOk returns a tuple with the ServiceItemInputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInputDefinition

`func (o *WorkflowServiceItemType) SetServiceItemInputDefinition(v []WorkflowServiceItemInputDefinitionType)`

SetServiceItemInputDefinition sets ServiceItemInputDefinition field to given value.

### HasServiceItemInputDefinition

`func (o *WorkflowServiceItemType) HasServiceItemInputDefinition() bool`

HasServiceItemInputDefinition returns a boolean if a field has been set.

### SetServiceItemInputDefinitionNil

`func (o *WorkflowServiceItemType) SetServiceItemInputDefinitionNil(b bool)`

 SetServiceItemInputDefinitionNil sets the value for ServiceItemInputDefinition to be an explicit nil

### UnsetServiceItemInputDefinition
`func (o *WorkflowServiceItemType) UnsetServiceItemInputDefinition()`

UnsetServiceItemInputDefinition ensures that no value is present for ServiceItemInputDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


