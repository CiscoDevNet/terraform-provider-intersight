# ServerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "server.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "server.Profile"]
**ConfigChangeContext** | Pointer to [**NullablePolicyConfigChangeContext**](PolicyConfigChangeContext.md) |  | [optional] 
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**DeployStatus** | Pointer to **string** | The status of the server profile indicating if deployment has been initiated on both fabric interconnects or not. * &#x60;None&#x60; - Switch profiles not deployed on either of the switches. * &#x60;Complete&#x60; - Both switch profiles of the cluster profile are deployed. * &#x60;Partial&#x60; - Only one of the switch profiles of the cluster profile is deployed. | [optional] [readonly] [default to "None"]
**DeployedSwitches** | Pointer to **string** | The property which determines if the deployment should be skipped on any of the Fabric Interconnects. It is set based on the state of a fabric interconnect to Intersight before the deployment of the server proile begins. * &#x60;None&#x60; - Server profile configuration not deployed on either of the fabric interconnects. * &#x60;AB&#x60; - Server profile configuration deployed on both fabric interconnects. * &#x60;A&#x60; - Server profile configuration deployed on fabric interconnect A only. * &#x60;B&#x60; - Server profile configuration deployed on fabric interconnect B only. | [optional] [readonly] [default to "None"]
**InternalReservationReferences** | Pointer to [**[]PoolReservationReference**](PoolReservationReference.md) |  | [optional] 
**IsPmcDeployedSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;pmcDeployedSecurePassphrase&#39; property has been set. | [optional] [readonly] [default to false]
**LocationDetails** | Pointer to [**NullableCommGeoLocationDetails**](CommGeoLocationDetails.md) |  | [optional] 
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**PmcDeployedSecurePassphrase** | Pointer to **string** | Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. | [optional] 
**ReservationReferences** | Pointer to [**[]PoolReservationReference**](PoolReservationReference.md) |  | [optional] 
**ServerAssignmentMode** | Pointer to **string** | Source of the server assigned to the Server Profile. Values can be Static, Pool or None. Static is used if a server is attached directly to a Server Profile. Pool is used if a resource pool is attached to a Server Profile. None is used if no server or resource pool is attached to a Server Profile. Slot or Serial pre-assignment is also considered to be None as it is different form of Assign Later. * &#x60;None&#x60; - No server is assigned to the server profile. * &#x60;Static&#x60; - Server is directly assigned to server profile using assign server. * &#x60;Pool&#x60; - Server is assigned from a resource pool. | [optional] [default to "None"]
**ServerPreAssignBySerial** | Pointer to **string** | Serial number of the server that would be assigned to this pre-assigned Server Profile. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It cannot be more than 20 characters. | [optional] 
**ServerPreAssignBySlot** | Pointer to [**NullableServerServerAssignTypeSlot**](ServerServerAssignTypeSlot.md) |  | [optional] 
**StaticUuidAddress** | Pointer to **string** | The UUID address for the server must include UUID prefix xxxxxxxx-xxxx-xxxx along with the UUID suffix of format xxxx-xxxxxxxxxxxx. | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**UserLabel** | Pointer to **string** | User label assigned to the server profile. | [optional] 
**Uuid** | Pointer to **string** | The UUID address that is assigned to the server based on the UUID pool. | [optional] [readonly] 
**AssignedServer** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**AssociatedServer** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**AssociatedServerPool** | Pointer to [**NullableResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ServerConfigChangeDetailRelationship**](ServerConfigChangeDetailRelationship.md) | An array of relationships to serverConfigChangeDetail resources. | [optional] [readonly] 
**LeasedServer** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ResourceLease** | Pointer to [**NullableResourcepoolLeaseRelationship**](ResourcepoolLeaseRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 
**ServerPool** | Pointer to [**NullableResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) |  | [optional] 
**UuidLease** | Pointer to [**NullableUuidpoolUuidLeaseRelationship**](UuidpoolUuidLeaseRelationship.md) |  | [optional] 

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
### GetDeployStatus

`func (o *ServerProfile) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerProfile) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerProfile) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *ServerProfile) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetDeployedSwitches

`func (o *ServerProfile) GetDeployedSwitches() string`

GetDeployedSwitches returns the DeployedSwitches field if non-nil, zero value otherwise.

### GetDeployedSwitchesOk

`func (o *ServerProfile) GetDeployedSwitchesOk() (*string, bool)`

GetDeployedSwitchesOk returns a tuple with the DeployedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSwitches

`func (o *ServerProfile) SetDeployedSwitches(v string)`

SetDeployedSwitches sets DeployedSwitches field to given value.

### HasDeployedSwitches

`func (o *ServerProfile) HasDeployedSwitches() bool`

HasDeployedSwitches returns a boolean if a field has been set.

### GetInternalReservationReferences

`func (o *ServerProfile) GetInternalReservationReferences() []PoolReservationReference`

GetInternalReservationReferences returns the InternalReservationReferences field if non-nil, zero value otherwise.

### GetInternalReservationReferencesOk

`func (o *ServerProfile) GetInternalReservationReferencesOk() (*[]PoolReservationReference, bool)`

GetInternalReservationReferencesOk returns a tuple with the InternalReservationReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalReservationReferences

`func (o *ServerProfile) SetInternalReservationReferences(v []PoolReservationReference)`

SetInternalReservationReferences sets InternalReservationReferences field to given value.

### HasInternalReservationReferences

`func (o *ServerProfile) HasInternalReservationReferences() bool`

HasInternalReservationReferences returns a boolean if a field has been set.

### SetInternalReservationReferencesNil

`func (o *ServerProfile) SetInternalReservationReferencesNil(b bool)`

 SetInternalReservationReferencesNil sets the value for InternalReservationReferences to be an explicit nil

### UnsetInternalReservationReferences
`func (o *ServerProfile) UnsetInternalReservationReferences()`

UnsetInternalReservationReferences ensures that no value is present for InternalReservationReferences, not even an explicit nil
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

### GetLocationDetails

`func (o *ServerProfile) GetLocationDetails() CommGeoLocationDetails`

GetLocationDetails returns the LocationDetails field if non-nil, zero value otherwise.

### GetLocationDetailsOk

`func (o *ServerProfile) GetLocationDetailsOk() (*CommGeoLocationDetails, bool)`

GetLocationDetailsOk returns a tuple with the LocationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDetails

`func (o *ServerProfile) SetLocationDetails(v CommGeoLocationDetails)`

SetLocationDetails sets LocationDetails field to given value.

### HasLocationDetails

`func (o *ServerProfile) HasLocationDetails() bool`

HasLocationDetails returns a boolean if a field has been set.

### SetLocationDetailsNil

`func (o *ServerProfile) SetLocationDetailsNil(b bool)`

 SetLocationDetailsNil sets the value for LocationDetails to be an explicit nil

### UnsetLocationDetails
`func (o *ServerProfile) UnsetLocationDetails()`

UnsetLocationDetails ensures that no value is present for LocationDetails, not even an explicit nil
### GetOverriddenList

`func (o *ServerProfile) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *ServerProfile) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *ServerProfile) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *ServerProfile) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *ServerProfile) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *ServerProfile) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
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

### GetServerPreAssignBySerial

`func (o *ServerProfile) GetServerPreAssignBySerial() string`

GetServerPreAssignBySerial returns the ServerPreAssignBySerial field if non-nil, zero value otherwise.

### GetServerPreAssignBySerialOk

`func (o *ServerProfile) GetServerPreAssignBySerialOk() (*string, bool)`

GetServerPreAssignBySerialOk returns a tuple with the ServerPreAssignBySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPreAssignBySerial

`func (o *ServerProfile) SetServerPreAssignBySerial(v string)`

SetServerPreAssignBySerial sets ServerPreAssignBySerial field to given value.

### HasServerPreAssignBySerial

`func (o *ServerProfile) HasServerPreAssignBySerial() bool`

HasServerPreAssignBySerial returns a boolean if a field has been set.

### GetServerPreAssignBySlot

`func (o *ServerProfile) GetServerPreAssignBySlot() ServerServerAssignTypeSlot`

GetServerPreAssignBySlot returns the ServerPreAssignBySlot field if non-nil, zero value otherwise.

### GetServerPreAssignBySlotOk

`func (o *ServerProfile) GetServerPreAssignBySlotOk() (*ServerServerAssignTypeSlot, bool)`

GetServerPreAssignBySlotOk returns a tuple with the ServerPreAssignBySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPreAssignBySlot

`func (o *ServerProfile) SetServerPreAssignBySlot(v ServerServerAssignTypeSlot)`

SetServerPreAssignBySlot sets ServerPreAssignBySlot field to given value.

### HasServerPreAssignBySlot

`func (o *ServerProfile) HasServerPreAssignBySlot() bool`

HasServerPreAssignBySlot returns a boolean if a field has been set.

### SetServerPreAssignBySlotNil

`func (o *ServerProfile) SetServerPreAssignBySlotNil(b bool)`

 SetServerPreAssignBySlotNil sets the value for ServerPreAssignBySlot to be an explicit nil

### UnsetServerPreAssignBySlot
`func (o *ServerProfile) UnsetServerPreAssignBySlot()`

UnsetServerPreAssignBySlot ensures that no value is present for ServerPreAssignBySlot, not even an explicit nil
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

### GetTemplateActions

`func (o *ServerProfile) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *ServerProfile) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *ServerProfile) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *ServerProfile) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *ServerProfile) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *ServerProfile) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *ServerProfile) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *ServerProfile) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *ServerProfile) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *ServerProfile) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *ServerProfile) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *ServerProfile) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *ServerProfile) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *ServerProfile) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *ServerProfile) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *ServerProfile) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

### GetUserLabel

`func (o *ServerProfile) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ServerProfile) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ServerProfile) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ServerProfile) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

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

### SetAssignedServerNil

`func (o *ServerProfile) SetAssignedServerNil(b bool)`

 SetAssignedServerNil sets the value for AssignedServer to be an explicit nil

### UnsetAssignedServer
`func (o *ServerProfile) UnsetAssignedServer()`

UnsetAssignedServer ensures that no value is present for AssignedServer, not even an explicit nil
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

### SetAssociatedServerNil

`func (o *ServerProfile) SetAssociatedServerNil(b bool)`

 SetAssociatedServerNil sets the value for AssociatedServer to be an explicit nil

### UnsetAssociatedServer
`func (o *ServerProfile) UnsetAssociatedServer()`

UnsetAssociatedServer ensures that no value is present for AssociatedServer, not even an explicit nil
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

### SetAssociatedServerPoolNil

`func (o *ServerProfile) SetAssociatedServerPoolNil(b bool)`

 SetAssociatedServerPoolNil sets the value for AssociatedServerPool to be an explicit nil

### UnsetAssociatedServerPool
`func (o *ServerProfile) UnsetAssociatedServerPool()`

UnsetAssociatedServerPool ensures that no value is present for AssociatedServerPool, not even an explicit nil
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

### SetLeasedServerNil

`func (o *ServerProfile) SetLeasedServerNil(b bool)`

 SetLeasedServerNil sets the value for LeasedServer to be an explicit nil

### UnsetLeasedServer
`func (o *ServerProfile) UnsetLeasedServer()`

UnsetLeasedServer ensures that no value is present for LeasedServer, not even an explicit nil
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

### SetOrganizationNil

`func (o *ServerProfile) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ServerProfile) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
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

### SetResourceLeaseNil

`func (o *ServerProfile) SetResourceLeaseNil(b bool)`

 SetResourceLeaseNil sets the value for ResourceLease to be an explicit nil

### UnsetResourceLease
`func (o *ServerProfile) UnsetResourceLease()`

UnsetResourceLease ensures that no value is present for ResourceLease, not even an explicit nil
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

### SetServerPoolNil

`func (o *ServerProfile) SetServerPoolNil(b bool)`

 SetServerPoolNil sets the value for ServerPool to be an explicit nil

### UnsetServerPool
`func (o *ServerProfile) UnsetServerPool()`

UnsetServerPool ensures that no value is present for ServerPool, not even an explicit nil
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

### SetUuidLeaseNil

`func (o *ServerProfile) SetUuidLeaseNil(b bool)`

 SetUuidLeaseNil sets the value for UuidLease to be an explicit nil

### UnsetUuidLease
`func (o *ServerProfile) UnsetUuidLease()`

UnsetUuidLease ensures that no value is present for UuidLease, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


