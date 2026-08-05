# ChassisProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.Profile"]
**ChassisAssignmentMode** | Pointer to **string** | Source of the chassis assigned to the Chassis Profile. Values can be Static or None. Static is used if a chassis is attached directly to a Chassis Profile. None is used if no chassis is attached to a Chassis Profile. Slot or Serial pre-assignment is also considered to be None as it is different form of Assign Later. * &#x60;Static&#x60; - Chassis is directly assigned to chassis profile using assign chassis. * &#x60;None&#x60; - No chassis is assigned to the chassis profile. | [optional] [default to "Static"]
**ChassisPreAssignBySerial** | Pointer to **string** | Serial number of the chassis that would be assigned to this pre-assigned Chassis Profile. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It cannot be more than 20 characters. | [optional] 
**ChassisReservation** | Pointer to [**ResourcepoolReservationReference**](ResourcepoolReservationReference.md) |  | [optional] 
**ConfigChangeContext** | Pointer to [**NullablePolicyConfigChangeContext**](PolicyConfigChangeContext.md) |  | [optional] 
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**ReservationReferences** | Pointer to [**[]PoolReservationReference**](PoolReservationReference.md) |  | [optional] 
**ScheduledChassisAssignment** | Pointer to [**NullableChassisChassisAssignment**](ChassisChassisAssignment.md) |  | [optional] 
**UserLabel** | Pointer to **string** | User label assigned to the chassis profile. | [optional] 
**AssignedChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**AssociatedChassis** | Pointer to [**NullableEquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ChassisConfigChangeDetailRelationship**](ChassisConfigChangeDetailRelationship.md) | An array of relationships to chassisConfigChangeDetail resources. | [optional] [readonly] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 

## Methods

### NewChassisProfile

`func NewChassisProfile(classId string, objectType string, ) *ChassisProfile`

NewChassisProfile instantiates a new ChassisProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisProfileWithDefaults

`func NewChassisProfileWithDefaults() *ChassisProfile`

NewChassisProfileWithDefaults instantiates a new ChassisProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChassisAssignmentMode

`func (o *ChassisProfile) GetChassisAssignmentMode() string`

GetChassisAssignmentMode returns the ChassisAssignmentMode field if non-nil, zero value otherwise.

### GetChassisAssignmentModeOk

`func (o *ChassisProfile) GetChassisAssignmentModeOk() (*string, bool)`

GetChassisAssignmentModeOk returns a tuple with the ChassisAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisAssignmentMode

`func (o *ChassisProfile) SetChassisAssignmentMode(v string)`

SetChassisAssignmentMode sets ChassisAssignmentMode field to given value.

### HasChassisAssignmentMode

`func (o *ChassisProfile) HasChassisAssignmentMode() bool`

HasChassisAssignmentMode returns a boolean if a field has been set.

### GetChassisPreAssignBySerial

`func (o *ChassisProfile) GetChassisPreAssignBySerial() string`

GetChassisPreAssignBySerial returns the ChassisPreAssignBySerial field if non-nil, zero value otherwise.

### GetChassisPreAssignBySerialOk

`func (o *ChassisProfile) GetChassisPreAssignBySerialOk() (*string, bool)`

GetChassisPreAssignBySerialOk returns a tuple with the ChassisPreAssignBySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisPreAssignBySerial

`func (o *ChassisProfile) SetChassisPreAssignBySerial(v string)`

SetChassisPreAssignBySerial sets ChassisPreAssignBySerial field to given value.

### HasChassisPreAssignBySerial

`func (o *ChassisProfile) HasChassisPreAssignBySerial() bool`

HasChassisPreAssignBySerial returns a boolean if a field has been set.

### GetChassisReservation

`func (o *ChassisProfile) GetChassisReservation() ResourcepoolReservationReference`

GetChassisReservation returns the ChassisReservation field if non-nil, zero value otherwise.

### GetChassisReservationOk

`func (o *ChassisProfile) GetChassisReservationOk() (*ResourcepoolReservationReference, bool)`

GetChassisReservationOk returns a tuple with the ChassisReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisReservation

`func (o *ChassisProfile) SetChassisReservation(v ResourcepoolReservationReference)`

SetChassisReservation sets ChassisReservation field to given value.

### HasChassisReservation

`func (o *ChassisProfile) HasChassisReservation() bool`

HasChassisReservation returns a boolean if a field has been set.

### GetConfigChangeContext

`func (o *ChassisProfile) GetConfigChangeContext() PolicyConfigChangeContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *ChassisProfile) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *ChassisProfile) SetConfigChangeContext(v PolicyConfigChangeContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *ChassisProfile) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### SetConfigChangeContextNil

`func (o *ChassisProfile) SetConfigChangeContextNil(b bool)`

 SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil

### UnsetConfigChangeContext
`func (o *ChassisProfile) UnsetConfigChangeContext()`

UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
### GetConfigChanges

`func (o *ChassisProfile) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *ChassisProfile) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *ChassisProfile) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *ChassisProfile) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### SetConfigChangesNil

`func (o *ChassisProfile) SetConfigChangesNil(b bool)`

 SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil

### UnsetConfigChanges
`func (o *ChassisProfile) UnsetConfigChanges()`

UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
### GetReservationReferences

`func (o *ChassisProfile) GetReservationReferences() []PoolReservationReference`

GetReservationReferences returns the ReservationReferences field if non-nil, zero value otherwise.

### GetReservationReferencesOk

`func (o *ChassisProfile) GetReservationReferencesOk() (*[]PoolReservationReference, bool)`

GetReservationReferencesOk returns a tuple with the ReservationReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationReferences

`func (o *ChassisProfile) SetReservationReferences(v []PoolReservationReference)`

SetReservationReferences sets ReservationReferences field to given value.

### HasReservationReferences

`func (o *ChassisProfile) HasReservationReferences() bool`

HasReservationReferences returns a boolean if a field has been set.

### SetReservationReferencesNil

`func (o *ChassisProfile) SetReservationReferencesNil(b bool)`

 SetReservationReferencesNil sets the value for ReservationReferences to be an explicit nil

### UnsetReservationReferences
`func (o *ChassisProfile) UnsetReservationReferences()`

UnsetReservationReferences ensures that no value is present for ReservationReferences, not even an explicit nil
### GetScheduledChassisAssignment

`func (o *ChassisProfile) GetScheduledChassisAssignment() ChassisChassisAssignment`

GetScheduledChassisAssignment returns the ScheduledChassisAssignment field if non-nil, zero value otherwise.

### GetScheduledChassisAssignmentOk

`func (o *ChassisProfile) GetScheduledChassisAssignmentOk() (*ChassisChassisAssignment, bool)`

GetScheduledChassisAssignmentOk returns a tuple with the ScheduledChassisAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledChassisAssignment

`func (o *ChassisProfile) SetScheduledChassisAssignment(v ChassisChassisAssignment)`

SetScheduledChassisAssignment sets ScheduledChassisAssignment field to given value.

### HasScheduledChassisAssignment

`func (o *ChassisProfile) HasScheduledChassisAssignment() bool`

HasScheduledChassisAssignment returns a boolean if a field has been set.

### SetScheduledChassisAssignmentNil

`func (o *ChassisProfile) SetScheduledChassisAssignmentNil(b bool)`

 SetScheduledChassisAssignmentNil sets the value for ScheduledChassisAssignment to be an explicit nil

### UnsetScheduledChassisAssignment
`func (o *ChassisProfile) UnsetScheduledChassisAssignment()`

UnsetScheduledChassisAssignment ensures that no value is present for ScheduledChassisAssignment, not even an explicit nil
### GetUserLabel

`func (o *ChassisProfile) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ChassisProfile) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ChassisProfile) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ChassisProfile) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetAssignedChassis

`func (o *ChassisProfile) GetAssignedChassis() EquipmentChassisRelationship`

GetAssignedChassis returns the AssignedChassis field if non-nil, zero value otherwise.

### GetAssignedChassisOk

`func (o *ChassisProfile) GetAssignedChassisOk() (*EquipmentChassisRelationship, bool)`

GetAssignedChassisOk returns a tuple with the AssignedChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedChassis

`func (o *ChassisProfile) SetAssignedChassis(v EquipmentChassisRelationship)`

SetAssignedChassis sets AssignedChassis field to given value.

### HasAssignedChassis

`func (o *ChassisProfile) HasAssignedChassis() bool`

HasAssignedChassis returns a boolean if a field has been set.

### SetAssignedChassisNil

`func (o *ChassisProfile) SetAssignedChassisNil(b bool)`

 SetAssignedChassisNil sets the value for AssignedChassis to be an explicit nil

### UnsetAssignedChassis
`func (o *ChassisProfile) UnsetAssignedChassis()`

UnsetAssignedChassis ensures that no value is present for AssignedChassis, not even an explicit nil
### GetAssociatedChassis

`func (o *ChassisProfile) GetAssociatedChassis() EquipmentChassisRelationship`

GetAssociatedChassis returns the AssociatedChassis field if non-nil, zero value otherwise.

### GetAssociatedChassisOk

`func (o *ChassisProfile) GetAssociatedChassisOk() (*EquipmentChassisRelationship, bool)`

GetAssociatedChassisOk returns a tuple with the AssociatedChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedChassis

`func (o *ChassisProfile) SetAssociatedChassis(v EquipmentChassisRelationship)`

SetAssociatedChassis sets AssociatedChassis field to given value.

### HasAssociatedChassis

`func (o *ChassisProfile) HasAssociatedChassis() bool`

HasAssociatedChassis returns a boolean if a field has been set.

### SetAssociatedChassisNil

`func (o *ChassisProfile) SetAssociatedChassisNil(b bool)`

 SetAssociatedChassisNil sets the value for AssociatedChassis to be an explicit nil

### UnsetAssociatedChassis
`func (o *ChassisProfile) UnsetAssociatedChassis()`

UnsetAssociatedChassis ensures that no value is present for AssociatedChassis, not even an explicit nil
### GetConfigChangeDetails

`func (o *ChassisProfile) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship`

GetConfigChangeDetails returns the ConfigChangeDetails field if non-nil, zero value otherwise.

### GetConfigChangeDetailsOk

`func (o *ChassisProfile) GetConfigChangeDetailsOk() (*[]ChassisConfigChangeDetailRelationship, bool)`

GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeDetails

`func (o *ChassisProfile) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship)`

SetConfigChangeDetails sets ConfigChangeDetails field to given value.

### HasConfigChangeDetails

`func (o *ChassisProfile) HasConfigChangeDetails() bool`

HasConfigChangeDetails returns a boolean if a field has been set.

### SetConfigChangeDetailsNil

`func (o *ChassisProfile) SetConfigChangeDetailsNil(b bool)`

 SetConfigChangeDetailsNil sets the value for ConfigChangeDetails to be an explicit nil

### UnsetConfigChangeDetails
`func (o *ChassisProfile) UnsetConfigChangeDetails()`

UnsetConfigChangeDetails ensures that no value is present for ConfigChangeDetails, not even an explicit nil
### GetOrganization

`func (o *ChassisProfile) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisProfile) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ChassisProfile) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ChassisProfile) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetRunningWorkflows

`func (o *ChassisProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *ChassisProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *ChassisProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *ChassisProfile) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *ChassisProfile) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *ChassisProfile) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


