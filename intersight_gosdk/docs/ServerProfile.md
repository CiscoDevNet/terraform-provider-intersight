# ServerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.Profile"]
**ConfigChangeContext** | Pointer to [**NullablePolicyConfigChangeContext**](PolicyConfigChangeContext.md) |  | [optional] 
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**IsPmcDeployedSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;pmcDeployedSecurePassphrase&#39; property has been set. | [optional] [readonly] [default to false]
**PmcDeployedSecurePassphrase** | Pointer to **string** | Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. | [optional] 
**AssignedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**AssociatedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ServerConfigChangeDetailRelationship**](ServerConfigChangeDetailRelationship.md) | An array of relationships to serverConfigChangeDetail resources. | [optional] [readonly] 
**ConfigResult** | Pointer to [**ServerConfigResultRelationship**](ServerConfigResultRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 

## Methods

### NewServerProfile

`func NewServerProfile(classId string, objectType string, ) *ServerProfile`

NewServerProfile instantiates a new ServerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProfileWithDefaults

`func NewServerProfileWithDefaults() *ServerProfile`

NewServerProfileWithDefaults instantiates a new ServerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServerProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServerProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigChangeContext

`func (o *ServerProfile) GetConfigChangeContext() PolicyConfigChangeContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *ServerProfile) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *ServerProfile) SetConfigChangeContext(v PolicyConfigChangeContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *ServerProfile) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### SetConfigChangeContextNil

`func (o *ServerProfile) SetConfigChangeContextNil(b bool)`

 SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil

### UnsetConfigChangeContext
`func (o *ServerProfile) UnsetConfigChangeContext()`

UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
### GetConfigChanges

`func (o *ServerProfile) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *ServerProfile) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *ServerProfile) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *ServerProfile) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### SetConfigChangesNil

`func (o *ServerProfile) SetConfigChangesNil(b bool)`

 SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil

### UnsetConfigChanges
`func (o *ServerProfile) UnsetConfigChanges()`

UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
### GetIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfile) GetIsPmcDeployedSecurePassphraseSet() bool`

GetIsPmcDeployedSecurePassphraseSet returns the IsPmcDeployedSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsPmcDeployedSecurePassphraseSetOk

`func (o *ServerProfile) GetIsPmcDeployedSecurePassphraseSetOk() (*bool, bool)`

GetIsPmcDeployedSecurePassphraseSetOk returns a tuple with the IsPmcDeployedSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfile) SetIsPmcDeployedSecurePassphraseSet(v bool)`

SetIsPmcDeployedSecurePassphraseSet sets IsPmcDeployedSecurePassphraseSet field to given value.

### HasIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfile) HasIsPmcDeployedSecurePassphraseSet() bool`

HasIsPmcDeployedSecurePassphraseSet returns a boolean if a field has been set.

### GetPmcDeployedSecurePassphrase

`func (o *ServerProfile) GetPmcDeployedSecurePassphrase() string`

GetPmcDeployedSecurePassphrase returns the PmcDeployedSecurePassphrase field if non-nil, zero value otherwise.

### GetPmcDeployedSecurePassphraseOk

`func (o *ServerProfile) GetPmcDeployedSecurePassphraseOk() (*string, bool)`

GetPmcDeployedSecurePassphraseOk returns a tuple with the PmcDeployedSecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmcDeployedSecurePassphrase

`func (o *ServerProfile) SetPmcDeployedSecurePassphrase(v string)`

SetPmcDeployedSecurePassphrase sets PmcDeployedSecurePassphrase field to given value.

### HasPmcDeployedSecurePassphrase

`func (o *ServerProfile) HasPmcDeployedSecurePassphrase() bool`

HasPmcDeployedSecurePassphrase returns a boolean if a field has been set.

### GetAssignedServer

`func (o *ServerProfile) GetAssignedServer() ComputePhysicalRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *ServerProfile) GetAssignedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *ServerProfile) SetAssignedServer(v ComputePhysicalRelationship)`

SetAssignedServer sets AssignedServer field to given value.

### HasAssignedServer

`func (o *ServerProfile) HasAssignedServer() bool`

HasAssignedServer returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *ServerProfile) GetAssociatedServer() ComputePhysicalRelationship`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *ServerProfile) GetAssociatedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *ServerProfile) SetAssociatedServer(v ComputePhysicalRelationship)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *ServerProfile) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetConfigChangeDetails

`func (o *ServerProfile) GetConfigChangeDetails() []ServerConfigChangeDetailRelationship`

GetConfigChangeDetails returns the ConfigChangeDetails field if non-nil, zero value otherwise.

### GetConfigChangeDetailsOk

`func (o *ServerProfile) GetConfigChangeDetailsOk() (*[]ServerConfigChangeDetailRelationship, bool)`

GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeDetails

`func (o *ServerProfile) SetConfigChangeDetails(v []ServerConfigChangeDetailRelationship)`

SetConfigChangeDetails sets ConfigChangeDetails field to given value.

### HasConfigChangeDetails

`func (o *ServerProfile) HasConfigChangeDetails() bool`

HasConfigChangeDetails returns a boolean if a field has been set.

### SetConfigChangeDetailsNil

`func (o *ServerProfile) SetConfigChangeDetailsNil(b bool)`

 SetConfigChangeDetailsNil sets the value for ConfigChangeDetails to be an explicit nil

### UnsetConfigChangeDetails
`func (o *ServerProfile) UnsetConfigChangeDetails()`

UnsetConfigChangeDetails ensures that no value is present for ConfigChangeDetails, not even an explicit nil
### GetConfigResult

`func (o *ServerProfile) GetConfigResult() ServerConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ServerProfile) GetConfigResultOk() (*ServerConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ServerProfile) SetConfigResult(v ServerConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ServerProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetOrganization

`func (o *ServerProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServerProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServerProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServerProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *ServerProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *ServerProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *ServerProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *ServerProfile) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *ServerProfile) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *ServerProfile) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


