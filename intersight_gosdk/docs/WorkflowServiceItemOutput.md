# WorkflowServiceItemOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemOutput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemOutput"]
**Name** | Pointer to **string** | Output name which is used in the output definition of the service item. | [optional] 
**Output** | Pointer to **interface{}** | Service item output for a service item instance and the format is specified by output definition of the service item definition. | [optional] 
**ServiceItemInstance** | Pointer to [**WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemOutput

`func NewWorkflowServiceItemOutput(classId string, objectType string, ) *WorkflowServiceItemOutput`

NewWorkflowServiceItemOutput instantiates a new WorkflowServiceItemOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemOutputWithDefaults

`func NewWorkflowServiceItemOutputWithDefaults() *WorkflowServiceItemOutput`

NewWorkflowServiceItemOutputWithDefaults instantiates a new WorkflowServiceItemOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemOutput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemOutput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemOutput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemOutput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemOutput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemOutput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *WorkflowServiceItemOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowServiceItemOutput) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowServiceItemOutput) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowServiceItemOutput) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowServiceItemOutput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowServiceItemOutput) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowServiceItemOutput) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetServiceItemInstance

`func (o *WorkflowServiceItemOutput) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowServiceItemOutput) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowServiceItemOutput) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowServiceItemOutput) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


