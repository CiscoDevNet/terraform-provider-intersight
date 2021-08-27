# ChassisProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.Profile"]
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**TargetPlatform** | Pointer to **string** | The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * &#x60;FIAttached&#x60; - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. | [optional] [default to "FIAttached"]
**AssignedChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**AssociatedChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ChassisConfigChangeDetailRelationship**](ChassisConfigChangeDetailRelationship.md) | An array of relationships to chassisConfigChangeDetail resources. | [optional] [readonly] 
**ConfigResult** | Pointer to [**ChassisConfigResultRelationship**](ChassisConfigResultRelationship.md) |  | [optional] 
**IomProfiles** | Pointer to [**[]ChassisIomProfileRelationship**](ChassisIomProfileRelationship.md) | An array of relationships to chassisIomProfile resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
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
### GetTargetPlatform

`func (o *ChassisProfile) GetTargetPlatform() string`

GetTargetPlatform returns the TargetPlatform field if non-nil, zero value otherwise.

### GetTargetPlatformOk

`func (o *ChassisProfile) GetTargetPlatformOk() (*string, bool)`

GetTargetPlatformOk returns a tuple with the TargetPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatform

`func (o *ChassisProfile) SetTargetPlatform(v string)`

SetTargetPlatform sets TargetPlatform field to given value.

### HasTargetPlatform

`func (o *ChassisProfile) HasTargetPlatform() bool`

HasTargetPlatform returns a boolean if a field has been set.

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
### GetConfigResult

`func (o *ChassisProfile) GetConfigResult() ChassisConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *ChassisProfile) GetConfigResultOk() (*ChassisConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *ChassisProfile) SetConfigResult(v ChassisConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *ChassisProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetIomProfiles

`func (o *ChassisProfile) GetIomProfiles() []ChassisIomProfileRelationship`

GetIomProfiles returns the IomProfiles field if non-nil, zero value otherwise.

### GetIomProfilesOk

`func (o *ChassisProfile) GetIomProfilesOk() (*[]ChassisIomProfileRelationship, bool)`

GetIomProfilesOk returns a tuple with the IomProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomProfiles

`func (o *ChassisProfile) SetIomProfiles(v []ChassisIomProfileRelationship)`

SetIomProfiles sets IomProfiles field to given value.

### HasIomProfiles

`func (o *ChassisProfile) HasIomProfiles() bool`

HasIomProfiles returns a boolean if a field has been set.

### SetIomProfilesNil

`func (o *ChassisProfile) SetIomProfilesNil(b bool)`

 SetIomProfilesNil sets the value for IomProfiles to be an explicit nil

### UnsetIomProfiles
`func (o *ChassisProfile) UnsetIomProfiles()`

UnsetIomProfiles ensures that no value is present for IomProfiles, not even an explicit nil
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


