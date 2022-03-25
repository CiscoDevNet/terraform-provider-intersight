# WorkflowServiceItemInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemInstance"]
**Description** | Pointer to **string** | The description for this service item instance. | [optional] 
**LastStatus** | Pointer to **string** | Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**Name** | Pointer to **string** | A name of the service item instance. Name of the service item instance must be unique within a type of Service item definition. | [optional] 
**Status** | Pointer to **string** | Status of the service item instance which controls the actions that can be performed on this instance. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**UserId** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemDefinition** | Pointer to [**WorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemInstanceAllOf

`func NewWorkflowServiceItemInstanceAllOf(classId string, objectType string, ) *WorkflowServiceItemInstanceAllOf`

NewWorkflowServiceItemInstanceAllOf instantiates a new WorkflowServiceItemInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemInstanceAllOfWithDefaults

`func NewWorkflowServiceItemInstanceAllOfWithDefaults() *WorkflowServiceItemInstanceAllOf`

NewWorkflowServiceItemInstanceAllOfWithDefaults instantiates a new WorkflowServiceItemInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemInstanceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemInstanceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemInstanceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemInstanceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemInstanceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemInstanceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowServiceItemInstanceAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemInstanceAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemInstanceAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemInstanceAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastStatus

`func (o *WorkflowServiceItemInstanceAllOf) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *WorkflowServiceItemInstanceAllOf) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *WorkflowServiceItemInstanceAllOf) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *WorkflowServiceItemInstanceAllOf) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemInstanceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemInstanceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemInstanceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemInstanceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowServiceItemInstanceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowServiceItemInstanceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowServiceItemInstanceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowServiceItemInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *WorkflowServiceItemInstanceAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowServiceItemInstanceAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowServiceItemInstanceAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowServiceItemInstanceAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowServiceItemInstanceAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowServiceItemInstanceAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowServiceItemInstanceAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowServiceItemInstanceAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServiceItemDefinition

`func (o *WorkflowServiceItemInstanceAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemInstanceAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemInstanceAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemInstanceAllOf) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


