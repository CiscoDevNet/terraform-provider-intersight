# WorkloadDeferredCleanupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.DeferredCleanupInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.DeferredCleanupInput"]
**Operations** | Pointer to [**[]WorkloadCleanupOperationInput**](WorkloadCleanupOperationInput.md) |  | [optional] 
**DeploymentInput** | Pointer to [**NullableWorkloadDeploymentInputRelationship**](WorkloadDeploymentInputRelationship.md) |  | [optional] 

## Methods

### NewWorkloadDeferredCleanupInput

`func NewWorkloadDeferredCleanupInput(classId string, objectType string, ) *WorkloadDeferredCleanupInput`

NewWorkloadDeferredCleanupInput instantiates a new WorkloadDeferredCleanupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeferredCleanupInputWithDefaults

`func NewWorkloadDeferredCleanupInputWithDefaults() *WorkloadDeferredCleanupInput`

NewWorkloadDeferredCleanupInputWithDefaults instantiates a new WorkloadDeferredCleanupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadDeferredCleanupInput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadDeferredCleanupInput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadDeferredCleanupInput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadDeferredCleanupInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadDeferredCleanupInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadDeferredCleanupInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperations

`func (o *WorkloadDeferredCleanupInput) GetOperations() []WorkloadCleanupOperationInput`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *WorkloadDeferredCleanupInput) GetOperationsOk() (*[]WorkloadCleanupOperationInput, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *WorkloadDeferredCleanupInput) SetOperations(v []WorkloadCleanupOperationInput)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *WorkloadDeferredCleanupInput) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperationsNil

`func (o *WorkloadDeferredCleanupInput) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *WorkloadDeferredCleanupInput) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetDeploymentInput

`func (o *WorkloadDeferredCleanupInput) GetDeploymentInput() WorkloadDeploymentInputRelationship`

GetDeploymentInput returns the DeploymentInput field if non-nil, zero value otherwise.

### GetDeploymentInputOk

`func (o *WorkloadDeferredCleanupInput) GetDeploymentInputOk() (*WorkloadDeploymentInputRelationship, bool)`

GetDeploymentInputOk returns a tuple with the DeploymentInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInput

`func (o *WorkloadDeferredCleanupInput) SetDeploymentInput(v WorkloadDeploymentInputRelationship)`

SetDeploymentInput sets DeploymentInput field to given value.

### HasDeploymentInput

`func (o *WorkloadDeferredCleanupInput) HasDeploymentInput() bool`

HasDeploymentInput returns a boolean if a field has been set.

### SetDeploymentInputNil

`func (o *WorkloadDeferredCleanupInput) SetDeploymentInputNil(b bool)`

 SetDeploymentInputNil sets the value for DeploymentInput to be an explicit nil

### UnsetDeploymentInput
`func (o *WorkloadDeferredCleanupInput) UnsetDeploymentInput()`

UnsetDeploymentInput ensures that no value is present for DeploymentInput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


