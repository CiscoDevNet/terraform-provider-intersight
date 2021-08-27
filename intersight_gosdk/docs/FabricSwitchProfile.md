# FabricSwitchProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SwitchProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SwitchProfile"]
**ConfigChanges** | Pointer to [**NullablePolicyConfigChange**](PolicyConfigChange.md) |  | [optional] 
**AssignedSwitch** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**AssociatedSwitch** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**ConfigChangeDetails** | Pointer to [**[]FabricConfigChangeDetailRelationship**](FabricConfigChangeDetailRelationship.md) | An array of relationships to fabricConfigChangeDetail resources. | [optional] [readonly] 
**ConfigResult** | Pointer to [**FabricConfigResultRelationship**](FabricConfigResultRelationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 
**SwitchClusterProfile** | Pointer to [**FabricSwitchClusterProfileRelationship**](FabricSwitchClusterProfileRelationship.md) |  | [optional] 

## Methods

### NewFabricSwitchProfile

`func NewFabricSwitchProfile(classId string, objectType string, ) *FabricSwitchProfile`

NewFabricSwitchProfile instantiates a new FabricSwitchProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchProfileWithDefaults

`func NewFabricSwitchProfileWithDefaults() *FabricSwitchProfile`

NewFabricSwitchProfileWithDefaults instantiates a new FabricSwitchProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSwitchProfile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchProfile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchProfile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSwitchProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigChanges

`func (o *FabricSwitchProfile) GetConfigChanges() PolicyConfigChange`

GetConfigChanges returns the ConfigChanges field if non-nil, zero value otherwise.

### GetConfigChangesOk

`func (o *FabricSwitchProfile) GetConfigChangesOk() (*PolicyConfigChange, bool)`

GetConfigChangesOk returns a tuple with the ConfigChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChanges

`func (o *FabricSwitchProfile) SetConfigChanges(v PolicyConfigChange)`

SetConfigChanges sets ConfigChanges field to given value.

### HasConfigChanges

`func (o *FabricSwitchProfile) HasConfigChanges() bool`

HasConfigChanges returns a boolean if a field has been set.

### SetConfigChangesNil

`func (o *FabricSwitchProfile) SetConfigChangesNil(b bool)`

 SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil

### UnsetConfigChanges
`func (o *FabricSwitchProfile) UnsetConfigChanges()`

UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
### GetAssignedSwitch

`func (o *FabricSwitchProfile) GetAssignedSwitch() NetworkElementRelationship`

GetAssignedSwitch returns the AssignedSwitch field if non-nil, zero value otherwise.

### GetAssignedSwitchOk

`func (o *FabricSwitchProfile) GetAssignedSwitchOk() (*NetworkElementRelationship, bool)`

GetAssignedSwitchOk returns a tuple with the AssignedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSwitch

`func (o *FabricSwitchProfile) SetAssignedSwitch(v NetworkElementRelationship)`

SetAssignedSwitch sets AssignedSwitch field to given value.

### HasAssignedSwitch

`func (o *FabricSwitchProfile) HasAssignedSwitch() bool`

HasAssignedSwitch returns a boolean if a field has been set.

### GetAssociatedSwitch

`func (o *FabricSwitchProfile) GetAssociatedSwitch() NetworkElementRelationship`

GetAssociatedSwitch returns the AssociatedSwitch field if non-nil, zero value otherwise.

### GetAssociatedSwitchOk

`func (o *FabricSwitchProfile) GetAssociatedSwitchOk() (*NetworkElementRelationship, bool)`

GetAssociatedSwitchOk returns a tuple with the AssociatedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedSwitch

`func (o *FabricSwitchProfile) SetAssociatedSwitch(v NetworkElementRelationship)`

SetAssociatedSwitch sets AssociatedSwitch field to given value.

### HasAssociatedSwitch

`func (o *FabricSwitchProfile) HasAssociatedSwitch() bool`

HasAssociatedSwitch returns a boolean if a field has been set.

### GetConfigChangeDetails

`func (o *FabricSwitchProfile) GetConfigChangeDetails() []FabricConfigChangeDetailRelationship`

GetConfigChangeDetails returns the ConfigChangeDetails field if non-nil, zero value otherwise.

### GetConfigChangeDetailsOk

`func (o *FabricSwitchProfile) GetConfigChangeDetailsOk() (*[]FabricConfigChangeDetailRelationship, bool)`

GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeDetails

`func (o *FabricSwitchProfile) SetConfigChangeDetails(v []FabricConfigChangeDetailRelationship)`

SetConfigChangeDetails sets ConfigChangeDetails field to given value.

### HasConfigChangeDetails

`func (o *FabricSwitchProfile) HasConfigChangeDetails() bool`

HasConfigChangeDetails returns a boolean if a field has been set.

### SetConfigChangeDetailsNil

`func (o *FabricSwitchProfile) SetConfigChangeDetailsNil(b bool)`

 SetConfigChangeDetailsNil sets the value for ConfigChangeDetails to be an explicit nil

### UnsetConfigChangeDetails
`func (o *FabricSwitchProfile) UnsetConfigChangeDetails()`

UnsetConfigChangeDetails ensures that no value is present for ConfigChangeDetails, not even an explicit nil
### GetConfigResult

`func (o *FabricSwitchProfile) GetConfigResult() FabricConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *FabricSwitchProfile) GetConfigResultOk() (*FabricConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *FabricSwitchProfile) SetConfigResult(v FabricConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *FabricSwitchProfile) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *FabricSwitchProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *FabricSwitchProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *FabricSwitchProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *FabricSwitchProfile) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *FabricSwitchProfile) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *FabricSwitchProfile) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil
### GetSwitchClusterProfile

`func (o *FabricSwitchProfile) GetSwitchClusterProfile() FabricSwitchClusterProfileRelationship`

GetSwitchClusterProfile returns the SwitchClusterProfile field if non-nil, zero value otherwise.

### GetSwitchClusterProfileOk

`func (o *FabricSwitchProfile) GetSwitchClusterProfileOk() (*FabricSwitchClusterProfileRelationship, bool)`

GetSwitchClusterProfileOk returns a tuple with the SwitchClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchClusterProfile

`func (o *FabricSwitchProfile) SetSwitchClusterProfile(v FabricSwitchClusterProfileRelationship)`

SetSwitchClusterProfile sets SwitchClusterProfile field to given value.

### HasSwitchClusterProfile

`func (o *FabricSwitchProfile) HasSwitchClusterProfile() bool`

HasSwitchClusterProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


