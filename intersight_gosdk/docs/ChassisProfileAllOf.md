# ChassisProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "chassis.Profile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "chassis.Profile"]
**ConfigChangeContext** | Pointer to [**NullablePolicyConfigChangeContext**](PolicyConfigChangeContext.md) |  | [optional] 
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**UserLabel** | Pointer to **string** | User label assigned to the chassis profile. | [optional] 
**AssignedChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**AssociatedChassis** | Pointer to [**EquipmentChassisRelationship**](EquipmentChassisRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]ChassisConfigChangeDetailRelationship**](ChassisConfigChangeDetailRelationship.md) | An array of relationships to chassisConfigChangeDetail resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 

## Methods

### NewChassisProfileAllOf

`func NewChassisProfileAllOf(classId string, objectType string, ) *ChassisProfileAllOf`

NewChassisProfileAllOf instantiates a new ChassisProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisProfileAllOfWithDefaults

`func NewChassisProfileAllOfWithDefaults() *ChassisProfileAllOf`

NewChassisProfileAllOfWithDefaults instantiates a new ChassisProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChassisProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChassisProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChassisProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChassisProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChassisProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChassisProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigChangeContext

`func (o *ChassisProfileAllOf) GetConfigChangeContext() PolicyConfigChangeContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *ChassisProfileAllOf) GetConfigChangeContextOk() (*PolicyConfigChangeContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *ChassisProfileAllOf) SetConfigChangeContext(v PolicyConfigChangeContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *ChassisProfileAllOf) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### SetConfigChangeContextNil

`func (o *ChassisProfileAllOf) SetConfigChangeContextNil(b bool)`

 SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil

### UnsetConfigChangeContext
`func (o *ChassisProfileAllOf) UnsetConfigChangeContext()`

UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
### GetConfigChanges

`func (o *ChassisProfileAllOf) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *ChassisProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *ChassisProfileAllOf) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *ChassisProfileAllOf) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### SetConfigChangesNil

`func (o *ChassisProfileAllOf) SetConfigChangesNil(b bool)`

 SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil

### UnsetConfigChanges
`func (o *ChassisProfileAllOf) UnsetConfigChanges()`

UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
### GetUserLabel

`func (o *ChassisProfileAllOf) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ChassisProfileAllOf) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ChassisProfileAllOf) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ChassisProfileAllOf) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetAssignedChassis

`func (o *ChassisProfileAllOf) GetAssignedChassis() EquipmentChassisRelationship`

GetAssignedChassis returns the AssignedChassis field if non-nil, zero value otherwise.

### GetAssignedChassisOk

`func (o *ChassisProfileAllOf) GetAssignedChassisOk() (*EquipmentChassisRelationship, bool)`

GetAssignedChassisOk returns a tuple with the AssignedChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedChassis

`func (o *ChassisProfileAllOf) SetAssignedChassis(v EquipmentChassisRelationship)`

SetAssignedChassis sets AssignedChassis field to given value.

### HasAssignedChassis

`func (o *ChassisProfileAllOf) HasAssignedChassis() bool`

HasAssignedChassis returns a boolean if a field has been set.

### GetAssociatedChassis

`func (o *ChassisProfileAllOf) GetAssociatedChassis() EquipmentChassisRelationship`

GetAssociatedChassis returns the AssociatedChassis field if non-nil, zero value otherwise.

### GetAssociatedChassisOk

`func (o *ChassisProfileAllOf) GetAssociatedChassisOk() (*EquipmentChassisRelationship, bool)`

GetAssociatedChassisOk returns a tuple with the AssociatedChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedChassis

`func (o *ChassisProfileAllOf) SetAssociatedChassis(v EquipmentChassisRelationship)`

SetAssociatedChassis sets AssociatedChassis field to given value.

### HasAssociatedChassis

`func (o *ChassisProfileAllOf) HasAssociatedChassis() bool`

HasAssociatedChassis returns a boolean if a field has been set.

### GetConfigChangeDetails

`func (o *ChassisProfileAllOf) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship`

GetConfigChangeDetails returns the ConfigChangeDetails field if non-nil, zero value otherwise.

### GetConfigChangeDetailsOk

`func (o *ChassisProfileAllOf) GetConfigChangeDetailsOk() (*[]ChassisConfigChangeDetailRelationship, bool)`

GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeDetails

`func (o *ChassisProfileAllOf) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship)`

SetConfigChangeDetails sets ConfigChangeDetails field to given value.

### HasConfigChangeDetails

`func (o *ChassisProfileAllOf) HasConfigChangeDetails() bool`

HasConfigChangeDetails returns a boolean if a field has been set.

### SetConfigChangeDetailsNil

`func (o *ChassisProfileAllOf) SetConfigChangeDetailsNil(b bool)`

 SetConfigChangeDetailsNil sets the value for ConfigChangeDetails to be an explicit nil

### UnsetConfigChangeDetails
`func (o *ChassisProfileAllOf) UnsetConfigChangeDetails()`

UnsetConfigChangeDetails ensures that no value is present for ConfigChangeDetails, not even an explicit nil
### GetOrganization

`func (o *ChassisProfileAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ChassisProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ChassisProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ChassisProfileAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *ChassisProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *ChassisProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *ChassisProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *ChassisProfileAllOf) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *ChassisProfileAllOf) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *ChassisProfileAllOf) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


