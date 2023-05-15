# WorkflowServiceItemInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemInstance"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemInstance"]
**Description** | Pointer to **string** | The description for this service item instance. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the resource. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LastStatus** | Pointer to **string** | Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**Name** | Pointer to **string** | A name of the service item instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**ResourcelifecycleStatus** | Pointer to **string** | Lifecycle state of service item instance. * &#x60;Creating&#x60; - The service item is not yet created and creation action is in progress. * &#x60;Created&#x60; - The service item is created. * &#x60;Decommissioning&#x60; - The service item is not yet decommissioned and decommission action is in progress. * &#x60;Decommissioned&#x60; - The service item is decommisioned. * &#x60;Deleting&#x60; - The service item is not yet deleted and deletion action is in progress. * &#x60;Deleted&#x60; - The service item is deleted. * &#x60;Failed&#x60; - The service item action is failed to perform the operation. | [optional] [readonly] [default to "Creating"]
**Status** | Pointer to **string** | Status of the service item instance which controls the actions that can be performed on this instance. * &#x60;NotCreated&#x60; - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * &#x60;InProgress&#x60; - An action is in progress and until that action has reached a final state, another action cannot be started. * &#x60;Failed&#x60; - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * &#x60;Okay&#x60; - The last action on the service item instance completed and the service item instance is in Okay state. * &#x60;Decommissioned&#x60; - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. | [optional] [readonly] [default to "NotCreated"]
**UserIdOrEmail** | Pointer to **string** | The user identifier which indicates the user that started this workflow. | [optional] [readonly] 
**CatalogServiceRequest** | Pointer to [**[]WorkflowCatalogServiceRequestRelationship**](WorkflowCatalogServiceRequestRelationship.md) | An array of relationships to workflowCatalogServiceRequest resources. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**IamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ServiceItemDefinition** | Pointer to [**WorkflowServiceItemDefinitionRelationship**](WorkflowServiceItemDefinitionRelationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

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

### GetLabel

`func (o *WorkflowServiceItemInstanceAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemInstanceAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemInstanceAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemInstanceAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

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

### GetResourcelifecycleStatus

`func (o *WorkflowServiceItemInstanceAllOf) GetResourcelifecycleStatus() string`

GetResourcelifecycleStatus returns the ResourcelifecycleStatus field if non-nil, zero value otherwise.

### GetResourcelifecycleStatusOk

`func (o *WorkflowServiceItemInstanceAllOf) GetResourcelifecycleStatusOk() (*string, bool)`

GetResourcelifecycleStatusOk returns a tuple with the ResourcelifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcelifecycleStatus

`func (o *WorkflowServiceItemInstanceAllOf) SetResourcelifecycleStatus(v string)`

SetResourcelifecycleStatus sets ResourcelifecycleStatus field to given value.

### HasResourcelifecycleStatus

`func (o *WorkflowServiceItemInstanceAllOf) HasResourcelifecycleStatus() bool`

HasResourcelifecycleStatus returns a boolean if a field has been set.

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

### GetUserIdOrEmail

`func (o *WorkflowServiceItemInstanceAllOf) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowServiceItemInstanceAllOf) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowServiceItemInstanceAllOf) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowServiceItemInstanceAllOf) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetCatalogServiceRequest

`func (o *WorkflowServiceItemInstanceAllOf) GetCatalogServiceRequest() []WorkflowCatalogServiceRequestRelationship`

GetCatalogServiceRequest returns the CatalogServiceRequest field if non-nil, zero value otherwise.

### GetCatalogServiceRequestOk

`func (o *WorkflowServiceItemInstanceAllOf) GetCatalogServiceRequestOk() (*[]WorkflowCatalogServiceRequestRelationship, bool)`

GetCatalogServiceRequestOk returns a tuple with the CatalogServiceRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogServiceRequest

`func (o *WorkflowServiceItemInstanceAllOf) SetCatalogServiceRequest(v []WorkflowCatalogServiceRequestRelationship)`

SetCatalogServiceRequest sets CatalogServiceRequest field to given value.

### HasCatalogServiceRequest

`func (o *WorkflowServiceItemInstanceAllOf) HasCatalogServiceRequest() bool`

HasCatalogServiceRequest returns a boolean if a field has been set.

### SetCatalogServiceRequestNil

`func (o *WorkflowServiceItemInstanceAllOf) SetCatalogServiceRequestNil(b bool)`

 SetCatalogServiceRequestNil sets the value for CatalogServiceRequest to be an explicit nil

### UnsetCatalogServiceRequest
`func (o *WorkflowServiceItemInstanceAllOf) UnsetCatalogServiceRequest()`

UnsetCatalogServiceRequest ensures that no value is present for CatalogServiceRequest, not even an explicit nil
### GetIdp

`func (o *WorkflowServiceItemInstanceAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *WorkflowServiceItemInstanceAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *WorkflowServiceItemInstanceAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *WorkflowServiceItemInstanceAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetIdpReference

`func (o *WorkflowServiceItemInstanceAllOf) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *WorkflowServiceItemInstanceAllOf) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *WorkflowServiceItemInstanceAllOf) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *WorkflowServiceItemInstanceAllOf) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

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

### GetUser

`func (o *WorkflowServiceItemInstanceAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowServiceItemInstanceAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowServiceItemInstanceAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowServiceItemInstanceAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


