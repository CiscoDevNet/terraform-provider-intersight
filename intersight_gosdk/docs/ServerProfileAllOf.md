# ServerProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigChanges** | Pointer to [**PolicyConfigChange**](policy.ConfigChange.md) |  | [optional] 
**IsPmcDeployedSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;pmcDeployedSecurePassphrase&#39; property has been set. | [optional] [readonly] 
**PmcDeployedSecurePassphrase** | Pointer to **string** | Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. | [optional] 
**TargetPlatform** | Pointer to **string** | The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;Standalone&#x60; - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * &#x60;FIAttached&#x60; - Servers which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "Standalone"]
**AssignedServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 
**AssociatedServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ServerConfigChangeDetailRelationship**](server.ConfigChangeDetail.Relationship.md) | An array of relationships to serverConfigChangeDetail resources. | [optional] [readonly] 
**ConfigResult** | Pointer to [**ServerConfigResultRelationship**](server.ConfigResult.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 

## Methods

### NewServerProfileAllOf

`func NewServerProfileAllOf() *ServerProfileAllOf`

NewServerProfileAllOf instantiates a new ServerProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProfileAllOfWithDefaults

`func NewServerProfileAllOfWithDefaults() *ServerProfileAllOf`

NewServerProfileAllOfWithDefaults instantiates a new ServerProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigChanges

`func (o *ServerProfileAllOf) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *ServerProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *ServerProfileAllOf) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *ServerProfileAllOf) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### GetIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSet() bool`

GetIsPmcDeployedSecurePassphraseSet returns the IsPmcDeployedSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsPmcDeployedSecurePassphraseSetOk

`func (o *ServerProfileAllOf) GetIsPmcDeployedSecurePassphraseSetOk() (*bool, bool)`

GetIsPmcDeployedSecurePassphraseSetOk returns a tuple with the IsPmcDeployedSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfileAllOf) SetIsPmcDeployedSecurePassphraseSet(v bool)`

SetIsPmcDeployedSecurePassphraseSet sets IsPmcDeployedSecurePassphraseSet field to given value.

### HasIsPmcDeployedSecurePassphraseSet

`func (o *ServerProfileAllOf) HasIsPmcDeployedSecurePassphraseSet() bool`

HasIsPmcDeployedSecurePassphraseSet returns a boolean if a field has been set.

### GetPmcDeployedSecurePassphrase

`func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphrase() string`

GetPmcDeployedSecurePassphrase returns the PmcDeployedSecurePassphrase field if non-nil, zero value otherwise.

### GetPmcDeployedSecurePassphraseOk

`func (o *ServerProfileAllOf) GetPmcDeployedSecurePassphraseOk() (*string, bool)`

GetPmcDeployedSecurePassphraseOk returns a tuple with the PmcDeployedSecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmcDeployedSecurePassphrase

`func (o *ServerProfileAllOf) SetPmcDeployedSecurePassphrase(v string)`

SetPmcDeployedSecurePassphrase sets PmcDeployedSecurePassphrase field to given value.

### HasPmcDeployedSecurePassphrase

`func (o *ServerProfileAllOf) HasPmcDeployedSecurePassphrase() bool`

HasPmcDeployedSecurePassphrase returns a boolean if a field has been set.

### GetTargetPlatform

`func (o *ServerProfileAllOf) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *ServerProfileAllOf) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *ServerProfileAllOf) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *ServerProfileAllOf) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

### GetAssignedServer

`func (o *ServerProfileAllOf) GetAssignedServer() ComputePhysicalRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *ServerProfileAllOf) GetAssignedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *ServerProfileAllOf) SetAssignedServer(v ComputePhysicalRelationship)`

SetAssignedServer sets AssignedServer field to given value.

### HasAssignedServer

`func (o *ServerProfileAllOf) HasAssignedServer() bool`

HasAssignedServer returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *ServerProfileAllOf) GetAssociatedServer() ComputePhysicalRelationship`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *ServerProfileAllOf) GetAssociatedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *ServerProfileAllOf) SetAssociatedServer(v ComputePhysicalRelationship)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *ServerProfileAllOf) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetConfigChangeDetails

`func (o *ServerProfileAllOf) GetConfigChangeDetails() []ServerConfigChangeDetailRelationship`

GetConfigChangeDetails returns the ConfigChangeDetails field if non-nil, zero value otherwise.

### GetConfigChangeDetailsOk

`func (o *ServerProfileAllOf) GetConfigChangeDetailsOk() (*[]ServerConfigChangeDetailRelationship, bool)`

GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeDetails

`func (o *ServerProfileAllOf) SetConfigChangeDetails(v []ServerConfigChangeDetailRelationship)`

SetConfigChangeDetails sets ConfigChangeDetails field to given value.

### HasConfigChangeDetails

`func (o *ServerProfileAllOf) HasConfigChangeDetails() bool`

HasConfigChangeDetails returns a boolean if a field has been set.

### SetConfigChangeDetailsNil

`func (o *ServerProfileAllOf) SetConfigChangeDetailsNil(b bool)`

 SetConfigChangeDetailsNil sets the value for ConfigChangeDetails to be an explicit nil

### UnsetConfigChangeDetails
`func (o *ServerProfileAllOf) UnsetConfigChangeDetails()`

UnsetConfigChangeDetails ensures that no value is present for ConfigChangeDetails, not even an explicit nil
### GetConfigResult

`func (o *ServerProfileAllOf) GetConfigResult() ServerConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ServerProfileAllOf) GetConfigResultOk() (*ServerConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ServerProfileAllOf) SetConfigResult(v ServerConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ServerProfileAllOf) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetOrganization

`func (o *ServerProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ServerProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ServerProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ServerProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *ServerProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *ServerProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *ServerProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *ServerProfileAllOf) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *ServerProfileAllOf) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *ServerProfileAllOf) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


