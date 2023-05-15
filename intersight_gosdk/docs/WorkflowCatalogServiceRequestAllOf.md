# WorkflowCatalogServiceRequestAllOf

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
**CatalogItemDefinition** | Pointer to [**WorkflowCatalogItemDefinitionRelationship**](WorkflowCatalogItemDefinitionRelationship.md) |  | [optional] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**IamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemActionInstances** | Pointer to [**[]WorkflowServiceItemActionInstanceRelationship**](WorkflowServiceItemActionInstanceRelationship.md) | An array of relationships to workflowServiceItemActionInstance resources. | [optional] [readonly] 
**ServiceItemInstances** | Pointer to [**[]WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) | An array of relationships to workflowServiceItemInstance resources. | [optional] [readonly] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewWorkflowCatalogServiceRequestAllOf

`func NewWorkflowCatalogServiceRequestAllOf(classId string, objectType string, ) *WorkflowCatalogServiceRequestAllOf`

NewWorkflowCatalogServiceRequestAllOf instantiates a new WorkflowCatalogServiceRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCatalogServiceRequestAllOfWithDefaults

`func NewWorkflowCatalogServiceRequestAllOfWithDefaults() *WorkflowCatalogServiceRequestAllOf`

NewWorkflowCatalogServiceRequestAllOfWithDefaults instantiates a new WorkflowCatalogServiceRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCatalogServiceRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCatalogServiceRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCatalogServiceRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCatalogServiceRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowCatalogServiceRequestAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCatalogServiceRequestAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCatalogServiceRequestAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowCatalogServiceRequestAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowCatalogServiceRequestAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowCatalogServiceRequestAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowCatalogServiceRequestAllOf) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowCatalogServiceRequestAllOf) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowCatalogServiceRequestAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetLabel

`func (o *WorkflowCatalogServiceRequestAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowCatalogServiceRequestAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowCatalogServiceRequestAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMessages

`func (o *WorkflowCatalogServiceRequestAllOf) GetMessages() []ServicerequestMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetMessagesOk() (*[]ServicerequestMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *WorkflowCatalogServiceRequestAllOf) SetMessages(v []ServicerequestMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *WorkflowCatalogServiceRequestAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *WorkflowCatalogServiceRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCatalogServiceRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowCatalogServiceRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *WorkflowCatalogServiceRequestAllOf) GetOperation() WorkflowBaseOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetOperationOk() (*WorkflowBaseOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *WorkflowCatalogServiceRequestAllOf) SetOperation(v WorkflowBaseOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *WorkflowCatalogServiceRequestAllOf) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequestAllOf) GetSelectionCriteriaInputs() []ServiceitemSelectionCriteriaInput`

GetSelectionCriteriaInputs returns the SelectionCriteriaInputs field if non-nil, zero value otherwise.

### GetSelectionCriteriaInputsOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetSelectionCriteriaInputsOk() (*[]ServiceitemSelectionCriteriaInput, bool)`

GetSelectionCriteriaInputsOk returns a tuple with the SelectionCriteriaInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequestAllOf) SetSelectionCriteriaInputs(v []ServiceitemSelectionCriteriaInput)`

SetSelectionCriteriaInputs sets SelectionCriteriaInputs field to given value.

### HasSelectionCriteriaInputs

`func (o *WorkflowCatalogServiceRequestAllOf) HasSelectionCriteriaInputs() bool`

HasSelectionCriteriaInputs returns a boolean if a field has been set.

### SetSelectionCriteriaInputsNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetSelectionCriteriaInputsNil(b bool)`

 SetSelectionCriteriaInputsNil sets the value for SelectionCriteriaInputs to be an explicit nil

### UnsetSelectionCriteriaInputs
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetSelectionCriteriaInputs()`

UnsetSelectionCriteriaInputs ensures that no value is present for SelectionCriteriaInputs, not even an explicit nil
### GetStatus

`func (o *WorkflowCatalogServiceRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowCatalogServiceRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowCatalogServiceRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *WorkflowCatalogServiceRequestAllOf) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowCatalogServiceRequestAllOf) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowCatalogServiceRequestAllOf) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequestAllOf) GetCatalogItemDefinition() WorkflowCatalogItemDefinitionRelationship`

GetCatalogItemDefinition returns the CatalogItemDefinition field if non-nil, zero value otherwise.

### GetCatalogItemDefinitionOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetCatalogItemDefinitionOk() (*WorkflowCatalogItemDefinitionRelationship, bool)`

GetCatalogItemDefinitionOk returns a tuple with the CatalogItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequestAllOf) SetCatalogItemDefinition(v WorkflowCatalogItemDefinitionRelationship)`

SetCatalogItemDefinition sets CatalogItemDefinition field to given value.

### HasCatalogItemDefinition

`func (o *WorkflowCatalogServiceRequestAllOf) HasCatalogItemDefinition() bool`

HasCatalogItemDefinition returns a boolean if a field has been set.

### GetIdp

`func (o *WorkflowCatalogServiceRequestAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *WorkflowCatalogServiceRequestAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *WorkflowCatalogServiceRequestAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpReference

`func (o *WorkflowCatalogServiceRequestAllOf) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *WorkflowCatalogServiceRequestAllOf) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *WorkflowCatalogServiceRequestAllOf) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowCatalogServiceRequestAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowCatalogServiceRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowCatalogServiceRequestAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemActionInstances() []WorkflowServiceItemActionInstanceRelationship`

GetServiceItemActionInstances returns the ServiceItemActionInstances field if non-nil, zero value otherwise.

### GetServiceItemActionInstancesOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemActionInstancesOk() (*[]WorkflowServiceItemActionInstanceRelationship, bool)`

GetServiceItemActionInstancesOk returns a tuple with the ServiceItemActionInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemActionInstances(v []WorkflowServiceItemActionInstanceRelationship)`

SetServiceItemActionInstances sets ServiceItemActionInstances field to given value.

### HasServiceItemActionInstances

`func (o *WorkflowCatalogServiceRequestAllOf) HasServiceItemActionInstances() bool`

HasServiceItemActionInstances returns a boolean if a field has been set.

### SetServiceItemActionInstancesNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemActionInstancesNil(b bool)`

 SetServiceItemActionInstancesNil sets the value for ServiceItemActionInstances to be an explicit nil

### UnsetServiceItemActionInstances
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetServiceItemActionInstances()`

UnsetServiceItemActionInstances ensures that no value is present for ServiceItemActionInstances, not even an explicit nil
### GetServiceItemInstances

`func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemInstances() []WorkflowServiceItemInstanceRelationship`

GetServiceItemInstances returns the ServiceItemInstances field if non-nil, zero value otherwise.

### GetServiceItemInstancesOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetServiceItemInstancesOk() (*[]WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstancesOk returns a tuple with the ServiceItemInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstances

`func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemInstances(v []WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstances sets ServiceItemInstances field to given value.

### HasServiceItemInstances

`func (o *WorkflowCatalogServiceRequestAllOf) HasServiceItemInstances() bool`

HasServiceItemInstances returns a boolean if a field has been set.

### SetServiceItemInstancesNil

`func (o *WorkflowCatalogServiceRequestAllOf) SetServiceItemInstancesNil(b bool)`

 SetServiceItemInstancesNil sets the value for ServiceItemInstances to be an explicit nil

### UnsetServiceItemInstances
`func (o *WorkflowCatalogServiceRequestAllOf) UnsetServiceItemInstances()`

UnsetServiceItemInstances ensures that no value is present for ServiceItemInstances, not even an explicit nil
### GetUser

`func (o *WorkflowCatalogServiceRequestAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowCatalogServiceRequestAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowCatalogServiceRequestAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowCatalogServiceRequestAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


