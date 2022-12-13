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
**ReservationReferences** | Pointer to [**[]PoolReservationReference**](PoolReservationReference.md) |  | [optional] 
**ServerAssignmentMode** | Pointer to **string** | Source of the server assigned to the server profile. Values can be Static, Pool or None. Static is used if a server is attached directly to server profile. Pool is used if a resource pool is attached to server profile. None is used if no server or resource pool is attached to server profile. * &#x60;None&#x60; - No server is assigned to the server profile. * &#x60;Static&#x60; - Server is directly assigned to server profile using assign server. * &#x60;Pool&#x60; - Server is assigned from a resource pool. | [optional] [default to "None"]
**StaticUuidAddress** | Pointer to **string** | The UUID address for the server must include UUID prefix xxxxxxxx-xxxx-xxxx along with the UUID suffix of format xxxx-xxxxxxxxxxxx. | [optional] 
**Uuid** | Pointer to **string** | The UUID address that is assigned to the server based on the UUID pool. | [optional] [readonly] 
**AssignedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**AssociatedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**AssociatedServerPool** | Pointer to [**ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ServerConfigChangeDetailRelationship**](ServerConfigChangeDetailRelationship.md) | An array of relationships to serverConfigChangeDetail resources. | [optional] [readonly] 
**LeasedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ResourceLease** | Pointer to [**ResourcepoolLeaseRelationship**](ResourcepoolLeaseRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 
**ServerPool** | Pointer to [**ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**UuidLease** | Pointer to [**UuidpoolUuidLeaseRelationship**](UuidpoolUuidLeaseRelationship.md) |  | [optional] 

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

### GetReservationReferences

`func (o *ServerProfile) GetReservationReferences() []PoolReservationReference`

GetReservationReferences returns the ReservationReferences field if non-nil, zero value otherwise.

### GetReservationReferencesOk

`func (o *ServerProfile) GetReservationReferencesOk() (*[]PoolReservationReference, bool)`

GetReservationReferencesOk returns a tuple with the ReservationReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationReferences

`func (o *ServerProfile) SetReservationReferences(v []PoolReservationReference)`

SetReservationReferences sets ReservationReferences field to given value.

### HasReservationReferences

`func (o *ServerProfile) HasReservationReferences() bool`

HasReservationReferences returns a boolean if a field has been set.

### SetReservationReferencesNil

`func (o *ServerProfile) SetReservationReferencesNil(b bool)`

 SetReservationReferencesNil sets the value for ReservationReferences to be an explicit nil

### UnsetReservationReferences
`func (o *ServerProfile) UnsetReservationReferences()`

UnsetReservationReferences ensures that no value is present for ReservationReferences, not even an explicit nil
### GetServerAssignmentMode

`func (o *ServerProfile) GetServerAssignmentMode() string`

GetServerAssignmentMode returns the ServerAssignmentMode field if non-nil, zero value otherwise.

### GetServerAssignmentModeOk

`func (o *ServerProfile) GetServerAssignmentModeOk() (*string, bool)`

GetServerAssignmentModeOk returns a tuple with the ServerAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssignmentMode

`func (o *ServerProfile) SetServerAssignmentMode(v string)`

SetServerAssignmentMode sets ServerAssignmentMode field to given value.

### HasServerAssignmentMode

`func (o *ServerProfile) HasServerAssignmentMode() bool`

HasServerAssignmentMode returns a boolean if a field has been set.

### GetStaticUuidAddress

`func (o *ServerProfile) GetStaticUuidAddress() string`

GetStaticUuidAddress returns the StaticUuidAddress field if non-nil, zero value otherwise.

### GetStaticUuidAddressOk

`func (o *ServerProfile) GetStaticUuidAddressOk() (*string, bool)`

GetStaticUuidAddressOk returns a tuple with the StaticUuidAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticUuidAddress

`func (o *ServerProfile) SetStaticUuidAddress(v string)`

SetStaticUuidAddress sets StaticUuidAddress field to given value.

### HasStaticUuidAddress

`func (o *ServerProfile) HasStaticUuidAddress() bool`

HasStaticUuidAddress returns a boolean if a field has been set.

### GetUuid

`func (o *ServerProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServerProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServerProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServerProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

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

### GetAssociatedServerPool

`func (o *ServerProfile) GetAssociatedServerPool() ResourcepoolPoolRelationship`

GetAssociatedServerPool returns the AssociatedServerPool field if non-nil, zero value otherwise.

### GetAssociatedServerPoolOk

`func (o *ServerProfile) GetAssociatedServerPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetAssociatedServerPoolOk returns a tuple with the AssociatedServerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServerPool

`func (o *ServerProfile) SetAssociatedServerPool(v ResourcepoolPoolRelationship)`

SetAssociatedServerPool sets AssociatedServerPool field to given value.

### HasAssociatedServerPool

`func (o *ServerProfile) HasAssociatedServerPool() bool`

HasAssociatedServerPool returns a boolean if a field has been set.

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
### GetLeasedServer

`func (o *ServerProfile) GetLeasedServer() ComputePhysicalRelationship`

GetLeasedServer returns the LeasedServer field if non-nil, zero value otherwise.

### GetLeasedServerOk

`func (o *ServerProfile) GetLeasedServerOk() (*ComputePhysicalRelationship, bool)`

GetLeasedServerOk returns a tuple with the LeasedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasedServer

`func (o *ServerProfile) SetLeasedServer(v ComputePhysicalRelationship)`

SetLeasedServer sets LeasedServer field to given value.

### HasLeasedServer

`func (o *ServerProfile) HasLeasedServer() bool`

HasLeasedServer returns a boolean if a field has been set.

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

### GetResourceLease

`func (o *ServerProfile) GetResourceLease() ResourcepoolLeaseRelationship`

GetResourceLease returns the ResourceLease field if non-nil, zero value otherwise.

### GetResourceLeaseOk

`func (o *ServerProfile) GetResourceLeaseOk() (*ResourcepoolLeaseRelationship, bool)`

GetResourceLeaseOk returns a tuple with the ResourceLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLease

`func (o *ServerProfile) SetResourceLease(v ResourcepoolLeaseRelationship)`

SetResourceLease sets ResourceLease field to given value.

### HasResourceLease

`func (o *ServerProfile) HasResourceLease() bool`

HasResourceLease returns a boolean if a field has been set.

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
### GetServerPool

`func (o *ServerProfile) GetServerPool() ResourcepoolPoolRelationship`

GetServerPool returns the ServerPool field if non-nil, zero value otherwise.

### GetServerPoolOk

`func (o *ServerProfile) GetServerPoolOk() (*ResourcepoolPoolRelationship, bool)`

GetServerPoolOk returns a tuple with the ServerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPool

`func (o *ServerProfile) SetServerPool(v ResourcepoolPoolRelationship)`

SetServerPool sets ServerPool field to given value.

### HasServerPool

`func (o *ServerProfile) HasServerPool() bool`

HasServerPool returns a boolean if a field has been set.

### GetUuidLease

`func (o *ServerProfile) GetUuidLease() UuidpoolUuidLeaseRelationship`

GetUuidLease returns the UuidLease field if non-nil, zero value otherwise.

### GetUuidLeaseOk

`func (o *ServerProfile) GetUuidLeaseOk() (*UuidpoolUuidLeaseRelationship, bool)`

GetUuidLeaseOk returns a tuple with the UuidLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidLease

`func (o *ServerProfile) SetUuidLease(v UuidpoolUuidLeaseRelationship)`

SetUuidLease sets UuidLease field to given value.

### HasUuidLease

`func (o *ServerProfile) HasUuidLease() bool`

HasUuidLease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


