# WorkloadBaseBlueprintReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Blueprint** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Input** | Pointer to **interface{}** | The input data for the referred blueprint. All required inputs of the blueprint must have a value provided. | [optional] 
**Name** | Pointer to **string** | The name for the referred blueprint. This name must be unique within the workload definition. | [optional] 
**RefName** | Pointer to **string** | The reference name for the blueprint which is derived by the system. | [optional] [readonly] 
**ResourceConstraint** | Pointer to [**NullableWorkloadResourceConstraint**](WorkloadResourceConstraint.md) |  | [optional] 

## Methods

### NewWorkloadBaseBlueprintReference

`func NewWorkloadBaseBlueprintReference(classId string, objectType string, ) *WorkloadBaseBlueprintReference`

NewWorkloadBaseBlueprintReference instantiates a new WorkloadBaseBlueprintReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBaseBlueprintReferenceWithDefaults

`func NewWorkloadBaseBlueprintReferenceWithDefaults() *WorkloadBaseBlueprintReference`

NewWorkloadBaseBlueprintReferenceWithDefaults instantiates a new WorkloadBaseBlueprintReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadBaseBlueprintReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadBaseBlueprintReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadBaseBlueprintReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadBaseBlueprintReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadBaseBlueprintReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadBaseBlueprintReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprint

`func (o *WorkloadBaseBlueprintReference) GetBlueprint() MoMoRef`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *WorkloadBaseBlueprintReference) GetBlueprintOk() (*MoMoRef, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *WorkloadBaseBlueprintReference) SetBlueprint(v MoMoRef)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *WorkloadBaseBlueprintReference) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetInput

`func (o *WorkloadBaseBlueprintReference) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkloadBaseBlueprintReference) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkloadBaseBlueprintReference) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkloadBaseBlueprintReference) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkloadBaseBlueprintReference) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkloadBaseBlueprintReference) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *WorkloadBaseBlueprintReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadBaseBlueprintReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadBaseBlueprintReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadBaseBlueprintReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefName

`func (o *WorkloadBaseBlueprintReference) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *WorkloadBaseBlueprintReference) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *WorkloadBaseBlueprintReference) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *WorkloadBaseBlueprintReference) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetResourceConstraint

`func (o *WorkloadBaseBlueprintReference) GetResourceConstraint() WorkloadResourceConstraint`

GetResourceConstraint returns the ResourceConstraint field if non-nil, zero value otherwise.

### GetResourceConstraintOk

`func (o *WorkloadBaseBlueprintReference) GetResourceConstraintOk() (*WorkloadResourceConstraint, bool)`

GetResourceConstraintOk returns a tuple with the ResourceConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConstraint

`func (o *WorkloadBaseBlueprintReference) SetResourceConstraint(v WorkloadResourceConstraint)`

SetResourceConstraint sets ResourceConstraint field to given value.

### HasResourceConstraint

`func (o *WorkloadBaseBlueprintReference) HasResourceConstraint() bool`

HasResourceConstraint returns a boolean if a field has been set.

### SetResourceConstraintNil

`func (o *WorkloadBaseBlueprintReference) SetResourceConstraintNil(b bool)`

 SetResourceConstraintNil sets the value for ResourceConstraint to be an explicit nil

### UnsetResourceConstraint
`func (o *WorkloadBaseBlueprintReference) UnsetResourceConstraint()`

UnsetResourceConstraint ensures that no value is present for ResourceConstraint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


