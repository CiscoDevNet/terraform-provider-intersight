# WorkflowSolutionInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionInstance"]
**Description** | Pointer to **string** | The description for this solution instance. | [optional] 
**LastStatus** | Pointer to **string** | Last status of the solution instance which will be reverted when an ongoing solution action instance is aborted. * &#x60;NotCreated&#x60; - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * &#x60;Okay&#x60; - The last action on the solution completed and the solution is in Okay state. * &#x60;Decommissioned&#x60; - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**Name** | Pointer to **string** | A name of the solution instance. Name of the solution instance must be unique within a type of Solution definition. | [optional] 
**Status** | Pointer to **string** | Status of the solution instance which controls the actions that can be performed on this instance. * &#x60;NotCreated&#x60; - Solution is not yet created and it is in a draft mode. A solution instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the solution failed and corrective measures need to be taken to bring the solution back to valid state. * &#x60;Okay&#x60; - The last action on the solution completed and the solution is in Okay state. * &#x60;Decommissioned&#x60; - The solution is decommissioned and can be safely deleted. Solution in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**UserId** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SolutionDefinition** | Pointer to [**WorkflowSolutionDefinitionRelationship**](WorkflowSolutionDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionInstanceAllOf

`func NewWorkflowSolutionInstanceAllOf(classId string, objectType string, ) *WorkflowSolutionInstanceAllOf`

NewWorkflowSolutionInstanceAllOf instantiates a new WorkflowSolutionInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionInstanceAllOfWithDefaults

`func NewWorkflowSolutionInstanceAllOfWithDefaults() *WorkflowSolutionInstanceAllOf`

NewWorkflowSolutionInstanceAllOfWithDefaults instantiates a new WorkflowSolutionInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionInstanceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionInstanceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionInstanceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionInstanceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionInstanceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionInstanceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowSolutionInstanceAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSolutionInstanceAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSolutionInstanceAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSolutionInstanceAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastStatus

`func (o *WorkflowSolutionInstanceAllOf) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *WorkflowSolutionInstanceAllOf) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *WorkflowSolutionInstanceAllOf) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *WorkflowSolutionInstanceAllOf) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSolutionInstanceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionInstanceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionInstanceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionInstanceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowSolutionInstanceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowSolutionInstanceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowSolutionInstanceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowSolutionInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowSolutionInstanceAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowSolutionInstanceAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowSolutionInstanceAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowSolutionInstanceAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowSolutionInstanceAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowSolutionInstanceAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowSolutionInstanceAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowSolutionInstanceAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSolutionDefinition

`func (o *WorkflowSolutionInstanceAllOf) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship`

GetSolutionDefinition returns the SolutionDefinition field if non-nil, zero value otherwise.

### GetSolutionDefinitionOk

`func (o *WorkflowSolutionInstanceAllOf) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool)`

GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionDefinition

`func (o *WorkflowSolutionInstanceAllOf) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship)`

SetSolutionDefinition sets SolutionDefinition field to given value.

### HasSolutionDefinition

`func (o *WorkflowSolutionInstanceAllOf) HasSolutionDefinition() bool`

HasSolutionDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


