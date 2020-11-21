# WorkflowBuildTaskMetaOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.BuildTaskMetaOwner"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.BuildTaskMetaOwner"]
**Service** | Pointer to **string** | The microservice owner responsible for the tasks. | [optional] [readonly] 
**WorkflowTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowBuildTaskMetaOwner

`func NewWorkflowBuildTaskMetaOwner(classId string, objectType string, ) *WorkflowBuildTaskMetaOwner`

NewWorkflowBuildTaskMetaOwner instantiates a new WorkflowBuildTaskMetaOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBuildTaskMetaOwnerWithDefaults

`func NewWorkflowBuildTaskMetaOwnerWithDefaults() *WorkflowBuildTaskMetaOwner`

NewWorkflowBuildTaskMetaOwnerWithDefaults instantiates a new WorkflowBuildTaskMetaOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBuildTaskMetaOwner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBuildTaskMetaOwner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBuildTaskMetaOwner) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBuildTaskMetaOwner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBuildTaskMetaOwner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBuildTaskMetaOwner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetService

`func (o *WorkflowBuildTaskMetaOwner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkflowBuildTaskMetaOwner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkflowBuildTaskMetaOwner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *WorkflowBuildTaskMetaOwner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwner) GetWorkflowTypes() []string`

GetWorkflowTypes returns the WorkflowTypes field if non-nil, zero value otherwise.

### GetWorkflowTypesOk

`func (o *WorkflowBuildTaskMetaOwner) GetWorkflowTypesOk() (*[]string, bool)`

GetWorkflowTypesOk returns a tuple with the WorkflowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwner) SetWorkflowTypes(v []string)`

SetWorkflowTypes sets WorkflowTypes field to given value.

### HasWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwner) HasWorkflowTypes() bool`

HasWorkflowTypes returns a boolean if a field has been set.

### SetWorkflowTypesNil

`func (o *WorkflowBuildTaskMetaOwner) SetWorkflowTypesNil(b bool)`

 SetWorkflowTypesNil sets the value for WorkflowTypes to be an explicit nil

### UnsetWorkflowTypes
`func (o *WorkflowBuildTaskMetaOwner) UnsetWorkflowTypes()`

UnsetWorkflowTypes ensures that no value is present for WorkflowTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


