# WorkflowEnvironmentVariableReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.EnvironmentVariableReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.EnvironmentVariableReference"]
**CatalogMoid** | Pointer to **string** | Environment variable catalog moid. | [optional] [readonly] 
**Moid** | Pointer to **string** | Environment variable moid. | [optional] [readonly] 
**Name** | Pointer to **string** | Environment variable name. | [optional] [readonly] 

## Methods

### NewWorkflowEnvironmentVariableReference

`func NewWorkflowEnvironmentVariableReference(classId string, objectType string, ) *WorkflowEnvironmentVariableReference`

NewWorkflowEnvironmentVariableReference instantiates a new WorkflowEnvironmentVariableReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEnvironmentVariableReferenceWithDefaults

`func NewWorkflowEnvironmentVariableReferenceWithDefaults() *WorkflowEnvironmentVariableReference`

NewWorkflowEnvironmentVariableReferenceWithDefaults instantiates a new WorkflowEnvironmentVariableReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowEnvironmentVariableReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowEnvironmentVariableReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowEnvironmentVariableReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowEnvironmentVariableReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowEnvironmentVariableReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowEnvironmentVariableReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowEnvironmentVariableReference) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowEnvironmentVariableReference) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowEnvironmentVariableReference) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowEnvironmentVariableReference) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetMoid

`func (o *WorkflowEnvironmentVariableReference) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *WorkflowEnvironmentVariableReference) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *WorkflowEnvironmentVariableReference) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *WorkflowEnvironmentVariableReference) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetName

`func (o *WorkflowEnvironmentVariableReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowEnvironmentVariableReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowEnvironmentVariableReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowEnvironmentVariableReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


