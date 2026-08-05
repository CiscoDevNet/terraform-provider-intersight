# FabricSwitchClusterProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchClusterProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchClusterProfile"]
**Action** | Pointer to **string** | The support actions are -- Deploy, Unassign. | [optional] [default to "No-op"]
**ChassisAssignmentMode** | Pointer to **string** | Source of the chassis assigned to the Switch Cluster Profile. Values can be Static or None. Static is used if a chassis is attached directly to a Switch Cluster Profile. None is used if no chassis is attached to a Switch Cluster Profile. Serial pre-assignment is also considered None. * &#x60;Static&#x60; - Chassis is directly assigned to switch cluster profile. * &#x60;None&#x60; - No chassis is assigned to the switch cluster profile. | [optional] [default to "Static"]
**ChassisPreAssignBySerial** | Pointer to **string** | Serial number of the chassis that would be assigned to this pre-assigned switch cluster profile. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It cannot be more than 20 characters. | [optional] 
**ClusterAssignments** | Pointer to [**[]FabricClusterAssignment**](FabricClusterAssignment.md) |  | [optional] 
**ConfigContext** | Pointer to [**NullablePolicyConfigContext**](PolicyConfigContext.md) |  | [optional] 
**DeployStatus** | Pointer to **string** | Deploy status of the switch cluster profile indicating if deployment has been initiated on all the members of the cluster profile. * &#x60;None&#x60; - Switch profiles not deployed on either of the switches. * &#x60;Complete&#x60; - Both switch profiles of the cluster profile are deployed. * &#x60;Partial&#x60; - Only one of the switch profiles of the cluster profile is deployed. | [optional] [readonly] [default to "None"]
**DeployedSwitches** | Pointer to **string** | Values indicating the switches on which the cluster profile has been deployed. 0 indicates that the profile has not been deployed on any switch, 1 indicates that the profile has been deployed on A, 2 indicates that it is deployed on B and 3 indicates that it is deployed on both. * &#x60;None&#x60; - Switch profiles not deployed on either of the fabric interconnects. * &#x60;A&#x60; - Switch profiles deployed only on fabric interconnect A. * &#x60;B&#x60; - Switch profiles deployed only on fabric interconnect B. * &#x60;AB&#x60; - Switch profiles deployed on both fabric interconnect A and B. | [optional] [readonly] [default to "None"]
**OverriddenList** | Pointer to **[]string** |  | [optional] 
**ScheduledChassisAssignment** | Pointer to [**NullableFabricChassisAssignment**](FabricChassisAssignment.md) |  | [optional] 
**TemplateActions** | Pointer to [**[]MotemplateActionEntry**](MotemplateActionEntry.md) |  | [optional] 
**TemplateSyncErrors** | Pointer to [**[]MotemplateSyncError**](MotemplateSyncError.md) |  | [optional] 
**TemplateSyncStatus** | Pointer to **string** | The sync status of the current MO wrt the attached Template MO. * &#x60;None&#x60; - The Enum value represents that the object is not attached to any template. * &#x60;OK&#x60; - The Enum value represents that the object values are in sync with attached template. * &#x60;Scheduled&#x60; - The Enum value represents that the object sync from attached template is scheduled from template. * &#x60;InProgress&#x60; - The Enum value represents that the object sync with the attached template is in progress. * &#x60;OutOfSync&#x60; - The Enum value represents that the object values are not in sync with attached template. | [optional] [readonly] [default to "None"]
**UserLabel** | Pointer to **string** | The user defined label assigned to the switch profile. | [optional] 
**AssignedEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**SwitchProfiles** | Pointer to [**[]FabricSwitchProfileRelationship**](FabricSwitchProfileRelationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfile

`func NewFabricSwitchClusterProfile(classId string, objectType string, ) *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfile instantiates a new FabricSwitchClusterProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileWithDefaults

`func NewFabricSwitchClusterProfileWithDefaults() *FabricSwitchClusterProfile`

NewFabricSwitchClusterProfileWithDefaults instantiates a new FabricSwitchClusterProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchClusterProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchClusterProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *FabricSwitchClusterProfile) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FabricSwitchClusterProfile) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FabricSwitchClusterProfile) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FabricSwitchClusterProfile) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetChassisAssignmentMode

`func (o *FabricSwitchClusterProfile) GetChassisAssignmentMode() string`

GetChassisAssignmentMode returns the ChassisAssignmentMode field if non-nil, zero value otherwise.

### GetChassisAssignmentModeOk

`func (o *FabricSwitchClusterProfile) GetChassisAssignmentModeOk() (*string, bool)`

GetChassisAssignmentModeOk returns a tuple with the ChassisAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisAssignmentMode

`func (o *FabricSwitchClusterProfile) SetChassisAssignmentMode(v string)`

SetChassisAssignmentMode sets ChassisAssignmentMode field to given value.

### HasChassisAssignmentMode

`func (o *FabricSwitchClusterProfile) HasChassisAssignmentMode() bool`

HasChassisAssignmentMode returns a boolean if a field has been set.

### GetChassisPreAssignBySerial

`func (o *FabricSwitchClusterProfile) GetChassisPreAssignBySerial() string`

GetChassisPreAssignBySerial returns the ChassisPreAssignBySerial field if non-nil, zero value otherwise.

### GetChassisPreAssignBySerialOk

`func (o *FabricSwitchClusterProfile) GetChassisPreAssignBySerialOk() (*string, bool)`

GetChassisPreAssignBySerialOk returns a tuple with the ChassisPreAssignBySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisPreAssignBySerial

`func (o *FabricSwitchClusterProfile) SetChassisPreAssignBySerial(v string)`

SetChassisPreAssignBySerial sets ChassisPreAssignBySerial field to given value.

### HasChassisPreAssignBySerial

`func (o *FabricSwitchClusterProfile) HasChassisPreAssignBySerial() bool`

HasChassisPreAssignBySerial returns a boolean if a field has been set.

### GetClusterAssignments

`func (o *FabricSwitchClusterProfile) GetClusterAssignments() []FabricClusterAssignment`

GetClusterAssignments returns the ClusterAssignments field if non-nil, zero value otherwise.

### GetClusterAssignmentsOk

`func (o *FabricSwitchClusterProfile) GetClusterAssignmentsOk() (*[]FabricClusterAssignment, bool)`

GetClusterAssignmentsOk returns a tuple with the ClusterAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAssignments

`func (o *FabricSwitchClusterProfile) SetClusterAssignments(v []FabricClusterAssignment)`

SetClusterAssignments sets ClusterAssignments field to given value.

### HasClusterAssignments

`func (o *FabricSwitchClusterProfile) HasClusterAssignments() bool`

HasClusterAssignments returns a boolean if a field has been set.

### SetClusterAssignmentsNil

`func (o *FabricSwitchClusterProfile) SetClusterAssignmentsNil(b bool)`

 SetClusterAssignmentsNil sets the value for ClusterAssignments to be an explicit nil

### UnsetClusterAssignments
`func (o *FabricSwitchClusterProfile) UnsetClusterAssignments()`

UnsetClusterAssignments ensures that no value is present for ClusterAssignments, not even an explicit nil
### GetConfigContext

`func (o *FabricSwitchClusterProfile) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *FabricSwitchClusterProfile) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *FabricSwitchClusterProfile) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *FabricSwitchClusterProfile) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### SetConfigContextNil

`func (o *FabricSwitchClusterProfile) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *FabricSwitchClusterProfile) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetDeployStatus

`func (o *FabricSwitchClusterProfile) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *FabricSwitchClusterProfile) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *FabricSwitchClusterProfile) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *FabricSwitchClusterProfile) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetDeployedSwitches

`func (o *FabricSwitchClusterProfile) GetDeployedSwitches() string`

GetDeployedSwitches returns the DeployedSwitches field if non-nil, zero value otherwise.

### GetDeployedSwitchesOk

`func (o *FabricSwitchClusterProfile) GetDeployedSwitchesOk() (*string, bool)`

GetDeployedSwitchesOk returns a tuple with the DeployedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedSwitches

`func (o *FabricSwitchClusterProfile) SetDeployedSwitches(v string)`

SetDeployedSwitches sets DeployedSwitches field to given value.

### HasDeployedSwitches

`func (o *FabricSwitchClusterProfile) HasDeployedSwitches() bool`

HasDeployedSwitches returns a boolean if a field has been set.

### GetOverriddenList

`func (o *FabricSwitchClusterProfile) GetOverriddenList() []string`

GetOverriddenList returns the OverriddenList field if non-nil, zero value otherwise.

### GetOverriddenListOk

`func (o *FabricSwitchClusterProfile) GetOverriddenListOk() (*[]string, bool)`

GetOverriddenListOk returns a tuple with the OverriddenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenList

`func (o *FabricSwitchClusterProfile) SetOverriddenList(v []string)`

SetOverriddenList sets OverriddenList field to given value.

### HasOverriddenList

`func (o *FabricSwitchClusterProfile) HasOverriddenList() bool`

HasOverriddenList returns a boolean if a field has been set.

### SetOverriddenListNil

`func (o *FabricSwitchClusterProfile) SetOverriddenListNil(b bool)`

 SetOverriddenListNil sets the value for OverriddenList to be an explicit nil

### UnsetOverriddenList
`func (o *FabricSwitchClusterProfile) UnsetOverriddenList()`

UnsetOverriddenList ensures that no value is present for OverriddenList, not even an explicit nil
### GetScheduledChassisAssignment

`func (o *FabricSwitchClusterProfile) GetScheduledChassisAssignment() FabricChassisAssignment`

GetScheduledChassisAssignment returns the ScheduledChassisAssignment field if non-nil, zero value otherwise.

### GetScheduledChassisAssignmentOk

`func (o *FabricSwitchClusterProfile) GetScheduledChassisAssignmentOk() (*FabricChassisAssignment, bool)`

GetScheduledChassisAssignmentOk returns a tuple with the ScheduledChassisAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledChassisAssignment

`func (o *FabricSwitchClusterProfile) SetScheduledChassisAssignment(v FabricChassisAssignment)`

SetScheduledChassisAssignment sets ScheduledChassisAssignment field to given value.

### HasScheduledChassisAssignment

`func (o *FabricSwitchClusterProfile) HasScheduledChassisAssignment() bool`

HasScheduledChassisAssignment returns a boolean if a field has been set.

### SetScheduledChassisAssignmentNil

`func (o *FabricSwitchClusterProfile) SetScheduledChassisAssignmentNil(b bool)`

 SetScheduledChassisAssignmentNil sets the value for ScheduledChassisAssignment to be an explicit nil

### UnsetScheduledChassisAssignment
`func (o *FabricSwitchClusterProfile) UnsetScheduledChassisAssignment()`

UnsetScheduledChassisAssignment ensures that no value is present for ScheduledChassisAssignment, not even an explicit nil
### GetTemplateActions

`func (o *FabricSwitchClusterProfile) GetTemplateActions() []MotemplateActionEntry`

GetTemplateActions returns the TemplateActions field if non-nil, zero value otherwise.

### GetTemplateActionsOk

`func (o *FabricSwitchClusterProfile) GetTemplateActionsOk() (*[]MotemplateActionEntry, bool)`

GetTemplateActionsOk returns a tuple with the TemplateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateActions

`func (o *FabricSwitchClusterProfile) SetTemplateActions(v []MotemplateActionEntry)`

SetTemplateActions sets TemplateActions field to given value.

### HasTemplateActions

`func (o *FabricSwitchClusterProfile) HasTemplateActions() bool`

HasTemplateActions returns a boolean if a field has been set.

### SetTemplateActionsNil

`func (o *FabricSwitchClusterProfile) SetTemplateActionsNil(b bool)`

 SetTemplateActionsNil sets the value for TemplateActions to be an explicit nil

### UnsetTemplateActions
`func (o *FabricSwitchClusterProfile) UnsetTemplateActions()`

UnsetTemplateActions ensures that no value is present for TemplateActions, not even an explicit nil
### GetTemplateSyncErrors

`func (o *FabricSwitchClusterProfile) GetTemplateSyncErrors() []MotemplateSyncError`

GetTemplateSyncErrors returns the TemplateSyncErrors field if non-nil, zero value otherwise.

### GetTemplateSyncErrorsOk

`func (o *FabricSwitchClusterProfile) GetTemplateSyncErrorsOk() (*[]MotemplateSyncError, bool)`

GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncErrors

`func (o *FabricSwitchClusterProfile) SetTemplateSyncErrors(v []MotemplateSyncError)`

SetTemplateSyncErrors sets TemplateSyncErrors field to given value.

### HasTemplateSyncErrors

`func (o *FabricSwitchClusterProfile) HasTemplateSyncErrors() bool`

HasTemplateSyncErrors returns a boolean if a field has been set.

### SetTemplateSyncErrorsNil

`func (o *FabricSwitchClusterProfile) SetTemplateSyncErrorsNil(b bool)`

 SetTemplateSyncErrorsNil sets the value for TemplateSyncErrors to be an explicit nil

### UnsetTemplateSyncErrors
`func (o *FabricSwitchClusterProfile) UnsetTemplateSyncErrors()`

UnsetTemplateSyncErrors ensures that no value is present for TemplateSyncErrors, not even an explicit nil
### GetTemplateSyncStatus

`func (o *FabricSwitchClusterProfile) GetTemplateSyncStatus() string`

GetTemplateSyncStatus returns the TemplateSyncStatus field if non-nil, zero value otherwise.

### GetTemplateSyncStatusOk

`func (o *FabricSwitchClusterProfile) GetTemplateSyncStatusOk() (*string, bool)`

GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSyncStatus

`func (o *FabricSwitchClusterProfile) SetTemplateSyncStatus(v string)`

SetTemplateSyncStatus sets TemplateSyncStatus field to given value.

### HasTemplateSyncStatus

`func (o *FabricSwitchClusterProfile) HasTemplateSyncStatus() bool`

HasTemplateSyncStatus returns a boolean if a field has been set.

### GetUserLabel

`func (o *FabricSwitchClusterProfile) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *FabricSwitchClusterProfile) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *FabricSwitchClusterProfile) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *FabricSwitchClusterProfile) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetAssignedEntity

`func (o *FabricSwitchClusterProfile) GetAssignedEntity() MoBaseMoRelationship`

GetAssignedEntity returns the AssignedEntity field if non-nil, zero value otherwise.

### GetAssignedEntityOk

`func (o *FabricSwitchClusterProfile) GetAssignedEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedEntityOk returns a tuple with the AssignedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedEntity

`func (o *FabricSwitchClusterProfile) SetAssignedEntity(v MoBaseMoRelationship)`

SetAssignedEntity sets AssignedEntity field to given value.

### HasAssignedEntity

`func (o *FabricSwitchClusterProfile) HasAssignedEntity() bool`

HasAssignedEntity returns a boolean if a field has been set.

### SetAssignedEntityNil

`func (o *FabricSwitchClusterProfile) SetAssignedEntityNil(b bool)`

 SetAssignedEntityNil sets the value for AssignedEntity to be an explicit nil

### UnsetAssignedEntity
`func (o *FabricSwitchClusterProfile) UnsetAssignedEntity()`

UnsetAssignedEntity ensures that no value is present for AssignedEntity, not even an explicit nil
### GetOrganization

`func (o *FabricSwitchClusterProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FabricSwitchClusterProfile) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FabricSwitchClusterProfile) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetSwitchProfiles

`func (o *FabricSwitchClusterProfile) GetSwitchProfiles() []FabricSwitchProfileRelationship`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *FabricSwitchClusterProfile) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *FabricSwitchClusterProfile) SetSwitchProfiles(v []FabricSwitchProfileRelationship)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *FabricSwitchClusterProfile) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### SetSwitchProfilesNil

`func (o *FabricSwitchClusterProfile) SetSwitchProfilesNil(b bool)`

 SetSwitchProfilesNil sets the value for SwitchProfiles to be an explicit nil

### UnsetSwitchProfiles
`func (o *FabricSwitchClusterProfile) UnsetSwitchProfiles()`

UnsetSwitchProfiles ensures that no value is present for SwitchProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


