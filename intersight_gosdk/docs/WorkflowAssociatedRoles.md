# WorkflowAssociatedRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Moid** | Pointer to **string** | Stores the identifier of the task definition for which the required roles are cached in the workflow definition. In the case of sub workflow tasks, this property stores the identifier of the workflow that is wrapped in the sub workflow task. | [optional] [readonly] 
**Roles** | Pointer to **[]string** |  | [optional] 
**TaskNames** | Pointer to **[]string** |  | [optional] 
**WorkflowRoles** | Pointer to [**[]WorkflowAssociatedRoles**](workflow.AssociatedRoles.md) |  | [optional] 

## Methods

### NewWorkflowAssociatedRoles

`func NewWorkflowAssociatedRoles() *WorkflowAssociatedRoles`

NewWorkflowAssociatedRoles instantiates a new WorkflowAssociatedRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAssociatedRolesWithDefaults

`func NewWorkflowAssociatedRolesWithDefaults() *WorkflowAssociatedRoles`

NewWorkflowAssociatedRolesWithDefaults instantiates a new WorkflowAssociatedRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoid

`func (o *WorkflowAssociatedRoles) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *WorkflowAssociatedRoles) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *WorkflowAssociatedRoles) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *WorkflowAssociatedRoles) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetRoles

`func (o *WorkflowAssociatedRoles) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *WorkflowAssociatedRoles) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *WorkflowAssociatedRoles) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *WorkflowAssociatedRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTaskNames

`func (o *WorkflowAssociatedRoles) GetTaskNames() []string`

GetTaskNames returns the TaskNames field if non-nil, zero value otherwise.

### GetTaskNamesOk

`func (o *WorkflowAssociatedRoles) GetTaskNamesOk() (*[]string, bool)`

GetTaskNamesOk returns a tuple with the TaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskNames

`func (o *WorkflowAssociatedRoles) SetTaskNames(v []string)`

SetTaskNames sets TaskNames field to given value.

### HasTaskNames

`func (o *WorkflowAssociatedRoles) HasTaskNames() bool`

HasTaskNames returns a boolean if a field has been set.

### GetWorkflowRoles

`func (o *WorkflowAssociatedRoles) GetWorkflowRoles() []WorkflowAssociatedRoles`

GetWorkflowRoles returns the WorkflowRoles field if non-nil, zero value otherwise.

### GetWorkflowRolesOk

`func (o *WorkflowAssociatedRoles) GetWorkflowRolesOk() (*[]WorkflowAssociatedRoles, bool)`

GetWorkflowRolesOk returns a tuple with the WorkflowRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRoles

`func (o *WorkflowAssociatedRoles) SetWorkflowRoles(v []WorkflowAssociatedRoles)`

SetWorkflowRoles sets WorkflowRoles field to given value.

### HasWorkflowRoles

`func (o *WorkflowAssociatedRoles) HasWorkflowRoles() bool`

HasWorkflowRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


