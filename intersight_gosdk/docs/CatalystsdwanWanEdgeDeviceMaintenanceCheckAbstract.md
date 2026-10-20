# CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceCheck"]
**Action** | Pointer to **string** | Action requested for this maintenance check. CREATE auto-triggers a refresh, and the server resets it to None when the check reaches a terminal state. * &#x60;None&#x60; - No refresh is currently requested. * &#x60;Refresh&#x60; - Refresh the maintenance activity data for the maintenance check. | [optional] [default to "None"]
**Activities** | Pointer to [**[]CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo**](CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo.md) |  | [optional] 
**Status** | Pointer to **string** | Current status of the maintenance check. * &#x60;None&#x60; - Initial state before processing has started. * &#x60;InProgress&#x60; - The check workflow is running. * &#x60;Completed&#x60; - The check completed successfully for all entries in targetChassis. * &#x60;PartiallyCompleted&#x60; - The check completed for a strict subset of targetChassis; the remainder were skipped due to validation or resolution failures. Per-entry failure details are available on the linked workflowInfo. Consumers may treat this as success or failure based on their own strict/lenient policy. * &#x60;Failed&#x60; - The maintenance check failed before it could produce a usable final activity snapshot for the requested targetChassis set. * &#x60;ValidationFailed&#x60; - No targetChassis entry resolved to a Secure Router known to Catalyst SD-WAN Manager, so the check could not run. If only some entries failed to resolve, status is PartiallyCompleted instead. * &#x60;NoOp&#x60; - No maintenance activity was queried because none of the targetChassis had a Secure Router, or no Secure Router was claimed in Catalyst SD-WAN Manager. | [optional] [readonly] [default to "None"]
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract

`func NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract(classId string, objectType string, ) *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract`

NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstractWithDefaults

`func NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstractWithDefaults() *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract`

NewCatalystsdwanWanEdgeDeviceMaintenanceCheckAbstractWithDefaults instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActivities

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetActivities() []CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetActivitiesOk() (*[]CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetActivities(v []CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivitiesNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetActivitiesNil(b bool)`

 SetActivitiesNil sets the value for Activities to be an explicit nil

### UnsetActivities
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) UnsetActivities()`

UnsetActivities ensures that no value is present for Activities, not even an explicit nil
### GetStatus

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheckAbstract) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


