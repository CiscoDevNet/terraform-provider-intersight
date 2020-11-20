# WorkflowBuildTaskMetaOwnerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.BuildTaskMetaOwner"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.BuildTaskMetaOwner"]
**Service** | Pointer to **string** | The microservice owner responsible for the tasks. | [optional] [readonly] 
**WorkflowTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkflowBuildTaskMetaOwnerAllOf

`func NewWorkflowBuildTaskMetaOwnerAllOf(classId string, objectType string, ) *WorkflowBuildTaskMetaOwnerAllOf`

NewWorkflowBuildTaskMetaOwnerAllOf instantiates a new WorkflowBuildTaskMetaOwnerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBuildTaskMetaOwnerAllOfWithDefaults

`func NewWorkflowBuildTaskMetaOwnerAllOfWithDefaults() *WorkflowBuildTaskMetaOwnerAllOf`

NewWorkflowBuildTaskMetaOwnerAllOfWithDefaults instantiates a new WorkflowBuildTaskMetaOwnerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBuildTaskMetaOwnerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBuildTaskMetaOwnerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetService

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkflowBuildTaskMetaOwnerAllOf) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *WorkflowBuildTaskMetaOwnerAllOf) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetWorkflowTypes() []string`

GetWorkflowTypes returns the WorkflowTypes field if non-nil, zero value otherwise.

### GetWorkflowTypesOk

`func (o *WorkflowBuildTaskMetaOwnerAllOf) GetWorkflowTypesOk() (*[]string, bool)`

GetWorkflowTypesOk returns a tuple with the WorkflowTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwnerAllOf) SetWorkflowTypes(v []string)`

SetWorkflowTypes sets WorkflowTypes field to given value.

### HasWorkflowTypes

`func (o *WorkflowBuildTaskMetaOwnerAllOf) HasWorkflowTypes() bool`

HasWorkflowTypes returns a boolean if a field has been set.

### SetWorkflowTypesNil

`func (o *WorkflowBuildTaskMetaOwnerAllOf) SetWorkflowTypesNil(b bool)`

 SetWorkflowTypesNil sets the value for WorkflowTypes to be an explicit nil

### UnsetWorkflowTypes
`func (o *WorkflowBuildTaskMetaOwnerAllOf) UnsetWorkflowTypes()`

UnsetWorkflowTypes ensures that no value is present for WorkflowTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


