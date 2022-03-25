# WorkflowSolutionOutputAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionOutput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionOutput"]
**Name** | Pointer to **string** | Output name which is used in the output definition of the solution. | [optional] 
**Output** | Pointer to **interface{}** | Solution output for a solution instance and the format is specified by output definition of the solution definition. | [optional] 
**UpgradedMoid** | Pointer to **string** | Stores the upgraded Moid for help during future lookups. | [optional] [readonly] 
**SolutionInstance** | Pointer to [**WorkflowSolutionInstanceRelationship**](WorkflowSolutionInstanceRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionOutputAllOf

`func NewWorkflowSolutionOutputAllOf(classId string, objectType string, ) *WorkflowSolutionOutputAllOf`

NewWorkflowSolutionOutputAllOf instantiates a new WorkflowSolutionOutputAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionOutputAllOfWithDefaults

`func NewWorkflowSolutionOutputAllOfWithDefaults() *WorkflowSolutionOutputAllOf`

NewWorkflowSolutionOutputAllOfWithDefaults instantiates a new WorkflowSolutionOutputAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionOutputAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionOutputAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionOutputAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionOutputAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionOutputAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionOutputAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *WorkflowSolutionOutputAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionOutputAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionOutputAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionOutputAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowSolutionOutputAllOf) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowSolutionOutputAllOf) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowSolutionOutputAllOf) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowSolutionOutputAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowSolutionOutputAllOf) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowSolutionOutputAllOf) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetUpgradedMoid

`func (o *WorkflowSolutionOutputAllOf) GetUpgradedMoid() string`

GetUpgradedMoid returns the UpgradedMoid field if non-nil, zero value otherwise.

### GetUpgradedMoidOk

`func (o *WorkflowSolutionOutputAllOf) GetUpgradedMoidOk() (*string, bool)`

GetUpgradedMoidOk returns a tuple with the UpgradedMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedMoid

`func (o *WorkflowSolutionOutputAllOf) SetUpgradedMoid(v string)`

SetUpgradedMoid sets UpgradedMoid field to given value.

### HasUpgradedMoid

`func (o *WorkflowSolutionOutputAllOf) HasUpgradedMoid() bool`

HasUpgradedMoid returns a boolean if a field has been set.

### GetSolutionInstance

`func (o *WorkflowSolutionOutputAllOf) GetSolutionInstance() WorkflowSolutionInstanceRelationship`

GetSolutionInstance returns the SolutionInstance field if non-nil, zero value otherwise.

### GetSolutionInstanceOk

`func (o *WorkflowSolutionOutputAllOf) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool)`

GetSolutionInstanceOk returns a tuple with the SolutionInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionInstance

`func (o *WorkflowSolutionOutputAllOf) SetSolutionInstance(v WorkflowSolutionInstanceRelationship)`

SetSolutionInstance sets SolutionInstance field to given value.

### HasSolutionInstance

`func (o *WorkflowSolutionOutputAllOf) HasSolutionInstance() bool`

HasSolutionInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


