# HciClusterOpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ClusterOpInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ClusterOpInfo"]
**LastRunTime** | Pointer to **time.Time** | The timestamp when the operation was last started. Set when transitioning to \&quot;InProgress\&quot; or \&quot;Scheduled\&quot; states. | [optional] [readonly] 
**Status** | Pointer to **string** | The current state of the operation. Possible values: - NotStarted: Operation never executed. - Scheduled: Workflow trigger requested (async mode), awaiting callback. - InProgress: Workflow is executing. - Completed: Operation finished successfully. - Failed: Operation failed or timed out. | [optional] [readonly] [default to "NotStarted"]
**WorkflowInfoMoid** | Pointer to **string** | The MoID of the workflow.WorkflowInfo object tracking this operation&#39;s execution. Empty for operations in \&quot;Scheduled\&quot; or \&quot;NotStarted\&quot; states. | [optional] [readonly] 
**WorkflowType** | Pointer to **string** | The workflow type identifier for this operation. Examples: \&quot;hci.NutanixClusterCheck\&quot;, \&quot;hci.ClusterUpgrade\&quot;. | [optional] [readonly] 

## Methods

### NewHciClusterOpInfo

`func NewHciClusterOpInfo(classId string, objectType string, ) *HciClusterOpInfo`

NewHciClusterOpInfo instantiates a new HciClusterOpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciClusterOpInfoWithDefaults

`func NewHciClusterOpInfoWithDefaults() *HciClusterOpInfo`

NewHciClusterOpInfoWithDefaults instantiates a new HciClusterOpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciClusterOpInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciClusterOpInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciClusterOpInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciClusterOpInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciClusterOpInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciClusterOpInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastRunTime

`func (o *HciClusterOpInfo) GetLastRunTime() time.Time`

GetLastRunTime returns the LastRunTime field if non-nil, zero value otherwise.

### GetLastRunTimeOk

`func (o *HciClusterOpInfo) GetLastRunTimeOk() (*time.Time, bool)`

GetLastRunTimeOk returns a tuple with the LastRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunTime

`func (o *HciClusterOpInfo) SetLastRunTime(v time.Time)`

SetLastRunTime sets LastRunTime field to given value.

### HasLastRunTime

`func (o *HciClusterOpInfo) HasLastRunTime() bool`

HasLastRunTime returns a boolean if a field has been set.

### GetStatus

`func (o *HciClusterOpInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HciClusterOpInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HciClusterOpInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HciClusterOpInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowInfoMoid

`func (o *HciClusterOpInfo) GetWorkflowInfoMoid() string`

GetWorkflowInfoMoid returns the WorkflowInfoMoid field if non-nil, zero value otherwise.

### GetWorkflowInfoMoidOk

`func (o *HciClusterOpInfo) GetWorkflowInfoMoidOk() (*string, bool)`

GetWorkflowInfoMoidOk returns a tuple with the WorkflowInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfoMoid

`func (o *HciClusterOpInfo) SetWorkflowInfoMoid(v string)`

SetWorkflowInfoMoid sets WorkflowInfoMoid field to given value.

### HasWorkflowInfoMoid

`func (o *HciClusterOpInfo) HasWorkflowInfoMoid() bool`

HasWorkflowInfoMoid returns a boolean if a field has been set.

### GetWorkflowType

`func (o *HciClusterOpInfo) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *HciClusterOpInfo) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *HciClusterOpInfo) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *HciClusterOpInfo) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


