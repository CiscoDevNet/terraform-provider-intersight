# WorkloadCleanupOperationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.CleanupOperationInput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.CleanupOperationInput"]
**BlueprintRefName** | Pointer to **string** | The reference name for the blueprint that produced the generated object. | [optional] [readonly] 
**GeneratedObjectName** | Pointer to **string** | The generated object metadata name targeted by the deferred cleanup operation. | [optional] [readonly] 
**Objects** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Operation** | Pointer to **string** | The cleanup operation to be performed after workload instances converge on the current deployment input. * &#x60;Update&#x60; - The operation is to update the object via an HTTP PATCH operation. * &#x60;Delete&#x60; - The operation is to immediately delete the object via an HTTP DELETE operation. * &#x60;MarkDeleted&#x60; - The object is marked as deleted. This adds a tag to the object indicating that it should be deleted as part of the object cleanup task during workload deployment. | [optional] [readonly] [default to "Update"]

## Methods

### NewWorkloadCleanupOperationInput

`func NewWorkloadCleanupOperationInput(classId string, objectType string, ) *WorkloadCleanupOperationInput`

NewWorkloadCleanupOperationInput instantiates a new WorkloadCleanupOperationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadCleanupOperationInputWithDefaults

`func NewWorkloadCleanupOperationInputWithDefaults() *WorkloadCleanupOperationInput`

NewWorkloadCleanupOperationInputWithDefaults instantiates a new WorkloadCleanupOperationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadCleanupOperationInput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadCleanupOperationInput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadCleanupOperationInput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadCleanupOperationInput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadCleanupOperationInput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadCleanupOperationInput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintRefName

`func (o *WorkloadCleanupOperationInput) GetBlueprintRefName() string`

GetBlueprintRefName returns the BlueprintRefName field if non-nil, zero value otherwise.

### GetBlueprintRefNameOk

`func (o *WorkloadCleanupOperationInput) GetBlueprintRefNameOk() (*string, bool)`

GetBlueprintRefNameOk returns a tuple with the BlueprintRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintRefName

`func (o *WorkloadCleanupOperationInput) SetBlueprintRefName(v string)`

SetBlueprintRefName sets BlueprintRefName field to given value.

### HasBlueprintRefName

`func (o *WorkloadCleanupOperationInput) HasBlueprintRefName() bool`

HasBlueprintRefName returns a boolean if a field has been set.

### GetGeneratedObjectName

`func (o *WorkloadCleanupOperationInput) GetGeneratedObjectName() string`

GetGeneratedObjectName returns the GeneratedObjectName field if non-nil, zero value otherwise.

### GetGeneratedObjectNameOk

`func (o *WorkloadCleanupOperationInput) GetGeneratedObjectNameOk() (*string, bool)`

GetGeneratedObjectNameOk returns a tuple with the GeneratedObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedObjectName

`func (o *WorkloadCleanupOperationInput) SetGeneratedObjectName(v string)`

SetGeneratedObjectName sets GeneratedObjectName field to given value.

### HasGeneratedObjectName

`func (o *WorkloadCleanupOperationInput) HasGeneratedObjectName() bool`

HasGeneratedObjectName returns a boolean if a field has been set.

### GetObjects

`func (o *WorkloadCleanupOperationInput) GetObjects() []MoMoRef`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *WorkloadCleanupOperationInput) GetObjectsOk() (*[]MoMoRef, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *WorkloadCleanupOperationInput) SetObjects(v []MoMoRef)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *WorkloadCleanupOperationInput) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *WorkloadCleanupOperationInput) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *WorkloadCleanupOperationInput) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOperation

`func (o *WorkloadCleanupOperationInput) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *WorkloadCleanupOperationInput) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *WorkloadCleanupOperationInput) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *WorkloadCleanupOperationInput) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


