# WorkflowResourceSelectorCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ResourceSelectorCriteria"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ResourceSelectorCriteria"]
**Parameters** | Pointer to **interface{}** | The runtime properties and the value can be stored in this property. | [optional] 
**ResourceSelectionCriteria** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkflowResourceSelectorCriteria

`func NewWorkflowResourceSelectorCriteria(classId string, objectType string, ) *WorkflowResourceSelectorCriteria`

NewWorkflowResourceSelectorCriteria instantiates a new WorkflowResourceSelectorCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowResourceSelectorCriteriaWithDefaults

`func NewWorkflowResourceSelectorCriteriaWithDefaults() *WorkflowResourceSelectorCriteria`

NewWorkflowResourceSelectorCriteriaWithDefaults instantiates a new WorkflowResourceSelectorCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowResourceSelectorCriteria) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowResourceSelectorCriteria) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowResourceSelectorCriteria) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowResourceSelectorCriteria) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowResourceSelectorCriteria) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowResourceSelectorCriteria) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetParameters

`func (o *WorkflowResourceSelectorCriteria) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WorkflowResourceSelectorCriteria) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WorkflowResourceSelectorCriteria) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WorkflowResourceSelectorCriteria) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *WorkflowResourceSelectorCriteria) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *WorkflowResourceSelectorCriteria) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetResourceSelectionCriteria

`func (o *WorkflowResourceSelectorCriteria) GetResourceSelectionCriteria() MoMoRef`

GetResourceSelectionCriteria returns the ResourceSelectionCriteria field if non-nil, zero value otherwise.

### GetResourceSelectionCriteriaOk

`func (o *WorkflowResourceSelectorCriteria) GetResourceSelectionCriteriaOk() (*MoMoRef, bool)`

GetResourceSelectionCriteriaOk returns a tuple with the ResourceSelectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelectionCriteria

`func (o *WorkflowResourceSelectorCriteria) SetResourceSelectionCriteria(v MoMoRef)`

SetResourceSelectionCriteria sets ResourceSelectionCriteria field to given value.

### HasResourceSelectionCriteria

`func (o *WorkflowResourceSelectorCriteria) HasResourceSelectionCriteria() bool`

HasResourceSelectionCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


