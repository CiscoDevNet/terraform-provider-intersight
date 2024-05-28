# WorkflowCatalogServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CatalogServiceRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CatalogServiceRequest"]
**Description** | Pointer to **string** | The description for this catalog service request which provides information on request from the user. | [optional] 
**EndTime** | Pointer to **time.Time** | The time at which the service request completed or stopped. | [optional] [readonly] 
**Input** | Pointer to **interface{}** | Inputs for a catalog service request instance creation, format is specified by input definition of the catalog item definition. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the catalog service request instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Messages** | Pointer to [**[]ServicerequestMessage**](ServicerequestMessage.md) |  | [optional] 
**Name** | Pointer to **string** | A name of the catalog service request instance. | [optional] 
**Operation** | Pointer to [**NullableWorkflowBaseOperation**](WorkflowBaseOperation.md) |  | [optional] 
**SelectionCriteriaInputs** | Pointer to [**[]ServiceitemSelectionCriteriaInput**](ServiceitemSelectionCriteriaInput.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the catalog service request instance which determines the actions that are allowed on this instance. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**UserIdOrEmail** | Pointer to **string** | The user identifier who invoked the request to create the catalog service request. | [optional] [readonly] 
**CatalogItemDefinition** | Pointer to [**NullableWorkflowCatalogItemDefinitionRelationship**](WorkflowCatalogItemDefinitionRelationship.md) |  | [optional] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**NullableIamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemActionInstances** | Pointer to [**[]WorkflowServiceItemActionInstanceRelationship**](WorkflowServiceItemActionInstanceRelationship.md) | An array of relationships to workflowServiceItemActionInstance resources. | [optional] [readonly] 
**ServiceItemInstances** | Pointer to [**[]WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) | An array of relationships to workflowServiceItemInstance resources. | [optional] [readonly] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewWorkflowCatalogServiceRequest

`func NewWorkflowCatalogServiceRequest(classId string, objectType string, ) *WorkflowCatalogServiceRequest`

NewWorkflowCatalogServiceRequest instantiates a new WorkflowCatalogServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCatalogServiceRequestWithDefaults

`func NewWorkflowCatalogServiceRequestWithDefaults() *WorkflowCatalogServiceRequest`

NewWorkflowCatalogServiceRequestWithDefaults instantiates a new WorkflowCatalogServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCatalogServiceRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCatalogServiceRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCatalogServiceRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCatalogServiceRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCatalogServiceRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCatalogServiceRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowCatalogServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCatalogServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCatalogServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCatalogServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowCatalogServiceRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowCatalogServiceRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowCatalogServiceRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowCatalogServiceRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowCatalogServiceRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowCatalogServiceRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowCatalogServiceRequest) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowCatalogServiceRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowCatalogServiceRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowCatalogServiceRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLabel

`func (o *WorkflowCatalogServiceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowCatalogServiceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowCatalogServiceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowCatalogServiceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMessages

`func (o *WorkflowCatalogServiceRequest) GetMessages() []ServicerequestMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *WorkflowCatalogServiceRequest) GetMessagesOk() (*[]ServicerequestMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *WorkflowCatalogServiceRequest) SetMessages(v []ServicerequestMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *WorkflowCatalogServiceRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *WorkflowCatalogServiceRequest) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *WorkflowCatalogServiceRequest) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *WorkflowCatalogServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCatalogServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCatalogServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowCatalogServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *WorkflowCatalogServiceRequest) GetOperation() WorkflowBaseOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *WorkflowCatalogServiceRequest) GetOperationOk() (*WorkflowBaseOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *WorkflowCatalogServiceRequest) SetOperation(v WorkflowBaseOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *WorkflowCatalogServiceRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *WorkflowCatalogServiceRequest) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *WorkflowCatalogServiceRequest) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequest) GetSelectionCriteriaInputs() []ServiceitemSelectionCriteriaInput`

GetSelectionCriteriaInputs returns the SelectionCriteriaInputs field if non-nil, zero value otherwise.

### GetSelectionCriteriaInputsOk

`func (o *WorkflowCatalogServiceRequest) GetSelectionCriteriaInputsOk() (*[]ServiceitemSelectionCriteriaInput, bool)`

GetSelectionCriteriaInputsOk returns a tuple with the SelectionCriteriaInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequest) SetSelectionCriteriaInputs(v []ServiceitemSelectionCriteriaInput)`

SetSelectionCriteriaInputs sets SelectionCriteriaInputs field to given value.

### HasSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequest) HasSelectionCriteriaInputs() bool`

HasSelectionCriteriaInputs returns a boolean if a field has been set.

### SetSelectionCriteriaInputsNil

`func (o *WorkflowCatalogServiceRequest) SetSelectionCriteriaInputsNil(b bool)`

 SetSelectionCriteriaInputsNil sets the value for SelectionCriteriaInputs to be an explicit nil

### UnsetSelectionCriteriaInputs
`func (o *WorkflowCatalogServiceRequest) UnsetSelectionCriteriaInputs()`

UnsetSelectionCriteriaInputs ensures that no value is present for SelectionCriteriaInputs, not even an explicit nil
### GetStatus

`func (o *WorkflowCatalogServiceRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowCatalogServiceRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowCatalogServiceRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowCatalogServiceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *WorkflowCatalogServiceRequest) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowCatalogServiceRequest) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowCatalogServiceRequest) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowCatalogServiceRequest) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequest) GetCatalogItemDefinition() WorkflowCatalogItemDefinitionRelationship`

GetCatalogItemDefinition returns the CatalogItemDefinition field if non-nil, zero value otherwise.

### GetCatalogItemDefinitionOk

`func (o *WorkflowCatalogServiceRequest) GetCatalogItemDefinitionOk() (*WorkflowCatalogItemDefinitionRelationship, bool)`

GetCatalogItemDefinitionOk returns a tuple with the CatalogItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequest) SetCatalogItemDefinition(v WorkflowCatalogItemDefinitionRelationship)`

SetCatalogItemDefinition sets CatalogItemDefinition field to given value.

### HasCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequest) HasCatalogItemDefinition() bool`

HasCatalogItemDefinition returns a boolean if a field has been set.

### SetCatalogItemDefinitionNil

`func (o *WorkflowCatalogServiceRequest) SetCatalogItemDefinitionNil(b bool)`

 SetCatalogItemDefinitionNil sets the value for CatalogItemDefinition to be an explicit nil

### UnsetCatalogItemDefinition
`func (o *WorkflowCatalogServiceRequest) UnsetCatalogItemDefinition()`

UnsetCatalogItemDefinition ensures that no value is present for CatalogItemDefinition, not even an explicit nil
### GetIdp

`func (o *WorkflowCatalogServiceRequest) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *WorkflowCatalogServiceRequest) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *WorkflowCatalogServiceRequest) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *WorkflowCatalogServiceRequest) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *WorkflowCatalogServiceRequest) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *WorkflowCatalogServiceRequest) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetIdpReference

`func (o *WorkflowCatalogServiceRequest) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *WorkflowCatalogServiceRequest) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *WorkflowCatalogServiceRequest) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *WorkflowCatalogServiceRequest) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

### SetIdpReferenceNil

`func (o *WorkflowCatalogServiceRequest) SetIdpReferenceNil(b bool)`

 SetIdpReferenceNil sets the value for IdpReference to be an explicit nil

### UnsetIdpReference
`func (o *WorkflowCatalogServiceRequest) UnsetIdpReference()`

UnsetIdpReference ensures that no value is present for IdpReference, not even an explicit nil
### GetOrganization

`func (o *WorkflowCatalogServiceRequest) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowCatalogServiceRequest) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowCatalogServiceRequest) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowCatalogServiceRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WorkflowCatalogServiceRequest) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WorkflowCatalogServiceRequest) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequest) GetServiceItemActionInstances() []WorkflowServiceItemActionInstanceRelationship`

GetServiceItemActionInstances returns the ServiceItemActionInstances field if non-nil, zero value otherwise.

### GetServiceItemActionInstancesOk

`func (o *WorkflowCatalogServiceRequest) GetServiceItemActionInstancesOk() (*[]WorkflowServiceItemActionInstanceRelationship, bool)`

GetServiceItemActionInstancesOk returns a tuple with the ServiceItemActionInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequest) SetServiceItemActionInstances(v []WorkflowServiceItemActionInstanceRelationship)`

SetServiceItemActionInstances sets ServiceItemActionInstances field to given value.

### HasServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequest) HasServiceItemActionInstances() bool`

HasServiceItemActionInstances returns a boolean if a field has been set.

### SetServiceItemActionInstancesNil

`func (o *WorkflowCatalogServiceRequest) SetServiceItemActionInstancesNil(b bool)`

 SetServiceItemActionInstancesNil sets the value for ServiceItemActionInstances to be an explicit nil

### UnsetServiceItemActionInstances
`func (o *WorkflowCatalogServiceRequest) UnsetServiceItemActionInstances()`

UnsetServiceItemActionInstances ensures that no value is present for ServiceItemActionInstances, not even an explicit nil
### GetServiceItemInstances

`func (o *WorkflowCatalogServiceRequest) GetServiceItemInstances() []WorkflowServiceItemInstanceRelationship`

GetServiceItemInstances returns the ServiceItemInstances field if non-nil, zero value otherwise.

### GetServiceItemInstancesOk

`func (o *WorkflowCatalogServiceRequest) GetServiceItemInstancesOk() (*[]WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstancesOk returns a tuple with the ServiceItemInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstances

`func (o *WorkflowCatalogServiceRequest) SetServiceItemInstances(v []WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstances sets ServiceItemInstances field to given value.

### HasServiceItemInstances

`func (o *WorkflowCatalogServiceRequest) HasServiceItemInstances() bool`

HasServiceItemInstances returns a boolean if a field has been set.

### SetServiceItemInstancesNil

`func (o *WorkflowCatalogServiceRequest) SetServiceItemInstancesNil(b bool)`

 SetServiceItemInstancesNil sets the value for ServiceItemInstances to be an explicit nil

### UnsetServiceItemInstances
`func (o *WorkflowCatalogServiceRequest) UnsetServiceItemInstances()`

UnsetServiceItemInstances ensures that no value is present for ServiceItemInstances, not even an explicit nil
### GetUser

`func (o *WorkflowCatalogServiceRequest) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowCatalogServiceRequest) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowCatalogServiceRequest) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowCatalogServiceRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *WorkflowCatalogServiceRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *WorkflowCatalogServiceRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


