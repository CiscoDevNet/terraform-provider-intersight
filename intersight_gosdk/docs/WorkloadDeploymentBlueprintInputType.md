# WorkloadDeploymentBlueprintInputType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.DeploymentBlueprintInputType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.DeploymentBlueprintInputType"]
**ExpungedObjects** | Pointer to [**[]WorkloadGeneratedObject**](WorkloadGeneratedObject.md) |  | [optional] 
**GeneratedObjects** | Pointer to [**[]WorkloadGeneratedObject**](WorkloadGeneratedObject.md) |  | [optional] 

## Methods

### NewWorkloadDeploymentBlueprintInputType

`func NewWorkloadDeploymentBlueprintInputType(classId string, objectType string, ) *WorkloadDeploymentBlueprintInputType`

NewWorkloadDeploymentBlueprintInputType instantiates a new WorkloadDeploymentBlueprintInputType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentBlueprintInputTypeWithDefaults

`func NewWorkloadDeploymentBlueprintInputTypeWithDefaults() *WorkloadDeploymentBlueprintInputType`

NewWorkloadDeploymentBlueprintInputTypeWithDefaults instantiates a new WorkloadDeploymentBlueprintInputType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadDeploymentBlueprintInputType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadDeploymentBlueprintInputType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadDeploymentBlueprintInputType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadDeploymentBlueprintInputType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadDeploymentBlueprintInputType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadDeploymentBlueprintInputType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpungedObjects

`func (o *WorkloadDeploymentBlueprintInputType) GetExpungedObjects() []WorkloadGeneratedObject`

GetExpungedObjects returns the ExpungedObjects field if non-nil, zero value otherwise.

### GetExpungedObjectsOk

`func (o *WorkloadDeploymentBlueprintInputType) GetExpungedObjectsOk() (*[]WorkloadGeneratedObject, bool)`

GetExpungedObjectsOk returns a tuple with the ExpungedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpungedObjects

`func (o *WorkloadDeploymentBlueprintInputType) SetExpungedObjects(v []WorkloadGeneratedObject)`

SetExpungedObjects sets ExpungedObjects field to given value.

### HasExpungedObjects

`func (o *WorkloadDeploymentBlueprintInputType) HasExpungedObjects() bool`

HasExpungedObjects returns a boolean if a field has been set.

### SetExpungedObjectsNil

`func (o *WorkloadDeploymentBlueprintInputType) SetExpungedObjectsNil(b bool)`

 SetExpungedObjectsNil sets the value for ExpungedObjects to be an explicit nil

### UnsetExpungedObjects
`func (o *WorkloadDeploymentBlueprintInputType) UnsetExpungedObjects()`

UnsetExpungedObjects ensures that no value is present for ExpungedObjects, not even an explicit nil
### GetGeneratedObjects

`func (o *WorkloadDeploymentBlueprintInputType) GetGeneratedObjects() []WorkloadGeneratedObject`

GetGeneratedObjects returns the GeneratedObjects field if non-nil, zero value otherwise.

### GetGeneratedObjectsOk

`func (o *WorkloadDeploymentBlueprintInputType) GetGeneratedObjectsOk() (*[]WorkloadGeneratedObject, bool)`

GetGeneratedObjectsOk returns a tuple with the GeneratedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedObjects

`func (o *WorkloadDeploymentBlueprintInputType) SetGeneratedObjects(v []WorkloadGeneratedObject)`

SetGeneratedObjects sets GeneratedObjects field to given value.

### HasGeneratedObjects

`func (o *WorkloadDeploymentBlueprintInputType) HasGeneratedObjects() bool`

HasGeneratedObjects returns a boolean if a field has been set.

### SetGeneratedObjectsNil

`func (o *WorkloadDeploymentBlueprintInputType) SetGeneratedObjectsNil(b bool)`

 SetGeneratedObjectsNil sets the value for GeneratedObjects to be an explicit nil

### UnsetGeneratedObjects
`func (o *WorkloadDeploymentBlueprintInputType) UnsetGeneratedObjects()`

UnsetGeneratedObjects ensures that no value is present for GeneratedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


