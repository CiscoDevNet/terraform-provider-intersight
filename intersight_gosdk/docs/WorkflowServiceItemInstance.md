# WorkflowServiceItemInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemInstance"]
**Description** | Pointer to **string** | The description for this service item instance. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the resource. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LastStatus** | Pointer to **string** | Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**MigrationHistory** | Pointer to [**[]WorkflowMigrationHistory**](WorkflowMigrationHistory.md) |  | [optional] 
**Name** | Pointer to **string** | A name of the service item instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**ResourcelifecycleStatus** | Pointer to **string** | Lifecycle state of service item instance. * &#x60;Creating&#x60; - The service item is not yet created and creation action is in progress. * &#x60;Created&#x60; - The service item is created. * &#x60;Decommissioning&#x60; - The service item is not yet decommissioned and decommission action is in progress. * &#x60;Decommissioned&#x60; - The service item is decommisioned. * &#x60;Deleting&#x60; - The service item is not yet deleted and deletion action is in progress. * &#x60;Deleted&#x60; - The service item is deleted. * &#x60;Failed&#x60; - The service item action is failed to perform the operation. | [optional] [readonly] [default to "Creating"]
**Status** | Pointer to **string** | Status of the service item instance which controls the actions that can be performed on this instance. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**UserIdOrEmail** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**CatalogServiceRequest** | Pointer to [**[]WorkflowCatalogServiceRequestRelationship**](WorkflowCatalogServiceRequestRelationship.md) | An array of relationships to workflowCatalogServiceRequest resources. | [optional] [readonly] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**NullableIamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemDefinition** | Pointer to [**NullableWorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemInstance

`func NewWorkflowServiceItemInstance(classId string, objectType string, ) *WorkflowServiceItemInstance`

NewWorkflowServiceItemInstance instantiates a new WorkflowServiceItemInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemInstanceWithDefaults

`func NewWorkflowServiceItemInstanceWithDefaults() *WorkflowServiceItemInstance`

NewWorkflowServiceItemInstanceWithDefaults instantiates a new WorkflowServiceItemInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemInstance) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemInstance) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemInstance) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemInstance) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemInstance) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemInstance) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowServiceItemInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowServiceItemInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemInstance) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemInstance) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastStatus

`func (o *WorkflowServiceItemInstance) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *WorkflowServiceItemInstance) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *WorkflowServiceItemInstance) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *WorkflowServiceItemInstance) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetMigrationHistory

`func (o *WorkflowServiceItemInstance) GetMigrationHistory() []WorkflowMigrationHistory`

GetMigrationHistory returns the MigrationHistory field if non-nil, zero value otherwise.

### GetMigrationHistoryOk

`func (o *WorkflowServiceItemInstance) GetMigrationHistoryOk() (*[]WorkflowMigrationHistory, bool)`

GetMigrationHistoryOk returns a tuple with the MigrationHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHistory

`func (o *WorkflowServiceItemInstance) SetMigrationHistory(v []WorkflowMigrationHistory)`

SetMigrationHistory sets MigrationHistory field to given value.

### HasMigrationHistory

`func (o *WorkflowServiceItemInstance) HasMigrationHistory() bool`

HasMigrationHistory returns a boolean if a field has been set.

### SetMigrationHistoryNil

`func (o *WorkflowServiceItemInstance) SetMigrationHistoryNil(b bool)`

 SetMigrationHistoryNil sets the value for MigrationHistory to be an explicit nil

### UnsetMigrationHistory
`func (o *WorkflowServiceItemInstance) UnsetMigrationHistory()`

UnsetMigrationHistory ensures that no value is present for MigrationHistory, not even an explicit nil
### GetName

`func (o *WorkflowServiceItemInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourcelifecycleStatus

`func (o *WorkflowServiceItemInstance) GetResourcelifecycleStatus() string`

GetResourcelifecycleStatus returns the ResourcelifecycleStatus field if non-nil, zero value otherwise.

### GetResourcelifecycleStatusOk

`func (o *WorkflowServiceItemInstance) GetResourcelifecycleStatusOk() (*string, bool)`

GetResourcelifecycleStatusOk returns a tuple with the ResourcelifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcelifecycleStatus

`func (o *WorkflowServiceItemInstance) SetResourcelifecycleStatus(v string)`

SetResourcelifecycleStatus sets ResourcelifecycleStatus field to given value.

### HasResourcelifecycleStatus

`func (o *WorkflowServiceItemInstance) HasResourcelifecycleStatus() bool`

HasResourcelifecycleStatus returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowServiceItemInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowServiceItemInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowServiceItemInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowServiceItemInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *WorkflowServiceItemInstance) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowServiceItemInstance) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowServiceItemInstance) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowServiceItemInstance) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetCatalogServiceRequest

`func (o *WorkflowServiceItemInstance) GetCatalogServiceRequest() []WorkflowCatalogServiceRequestRelationship`

GetCatalogServiceRequest returns the CatalogServiceRequest field if non-nil, zero value otherwise.

### GetCatalogServiceRequestOk

`func (o *WorkflowServiceItemInstance) GetCatalogServiceRequestOk() (*[]WorkflowCatalogServiceRequestRelationship, bool)`

GetCatalogServiceRequestOk returns a tuple with the CatalogServiceRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogServiceRequest

`func (o *WorkflowServiceItemInstance) SetCatalogServiceRequest(v []WorkflowCatalogServiceRequestRelationship)`

SetCatalogServiceRequest sets CatalogServiceRequest field to given value.

### HasCatalogServiceRequest

`func (o *WorkflowServiceItemInstance) HasCatalogServiceRequest() bool`

HasCatalogServiceRequest returns a boolean if a field has been set.

### SetCatalogServiceRequestNil

`func (o *WorkflowServiceItemInstance) SetCatalogServiceRequestNil(b bool)`

 SetCatalogServiceRequestNil sets the value for CatalogServiceRequest to be an explicit nil

### UnsetCatalogServiceRequest
`func (o *WorkflowServiceItemInstance) UnsetCatalogServiceRequest()`

UnsetCatalogServiceRequest ensures that no value is present for CatalogServiceRequest, not even an explicit nil
### GetIdp

`func (o *WorkflowServiceItemInstance) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *WorkflowServiceItemInstance) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *WorkflowServiceItemInstance) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *WorkflowServiceItemInstance) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *WorkflowServiceItemInstance) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *WorkflowServiceItemInstance) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetIdpReference

`func (o *WorkflowServiceItemInstance) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *WorkflowServiceItemInstance) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *WorkflowServiceItemInstance) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *WorkflowServiceItemInstance) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

### SetIdpReferenceNil

`func (o *WorkflowServiceItemInstance) SetIdpReferenceNil(b bool)`

 SetIdpReferenceNil sets the value for IdpReference to be an explicit nil

### UnsetIdpReference
`func (o *WorkflowServiceItemInstance) UnsetIdpReference()`

UnsetIdpReference ensures that no value is present for IdpReference, not even an explicit nil
### GetOrganization

`func (o *WorkflowServiceItemInstance) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowServiceItemInstance) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowServiceItemInstance) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowServiceItemInstance) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WorkflowServiceItemInstance) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WorkflowServiceItemInstance) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetServiceItemDefinition

`func (o *WorkflowServiceItemInstance) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship`

GetServiceItemDefinition returns the ServiceItemDefinition field if non-nil, zero value otherwise.

### GetServiceItemDefinitionOk

`func (o *WorkflowServiceItemInstance) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool)`

GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemDefinition

`func (o *WorkflowServiceItemInstance) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship)`

SetServiceItemDefinition sets ServiceItemDefinition field to given value.

### HasServiceItemDefinition

`func (o *WorkflowServiceItemInstance) HasServiceItemDefinition() bool`

HasServiceItemDefinition returns a boolean if a field has been set.

### SetServiceItemDefinitionNil

`func (o *WorkflowServiceItemInstance) SetServiceItemDefinitionNil(b bool)`

 SetServiceItemDefinitionNil sets the value for ServiceItemDefinition to be an explicit nil

### UnsetServiceItemDefinition
`func (o *WorkflowServiceItemInstance) UnsetServiceItemDefinition()`

UnsetServiceItemDefinition ensures that no value is present for ServiceItemDefinition, not even an explicit nil
### GetUser

`func (o *WorkflowServiceItemInstance) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowServiceItemInstance) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowServiceItemInstance) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowServiceItemInstance) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *WorkflowServiceItemInstance) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *WorkflowServiceItemInstance) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


