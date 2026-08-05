# HciPendingWorkflowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.PendingWorkflowTrigger"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.PendingWorkflowTrigger"]
**TriggerTime** | Pointer to **time.Time** | The timestamp when the workflow trigger request was sent. Used for timeout detection (default timeout: 20 minutes). | [optional] [readonly] 
**WorkflowType** | Pointer to **string** | The workflow type identifier for the pending operation. Example: \&quot;hci.NutanixClusterCheck\&quot;. | [optional] [readonly] 

## Methods

### NewHciPendingWorkflowTrigger

`func NewHciPendingWorkflowTrigger(classId string, objectType string, ) *HciPendingWorkflowTrigger`

NewHciPendingWorkflowTrigger instantiates a new HciPendingWorkflowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciPendingWorkflowTriggerWithDefaults

`func NewHciPendingWorkflowTriggerWithDefaults() *HciPendingWorkflowTrigger`

NewHciPendingWorkflowTriggerWithDefaults instantiates a new HciPendingWorkflowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciPendingWorkflowTrigger) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciPendingWorkflowTrigger) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciPendingWorkflowTrigger) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciPendingWorkflowTrigger) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciPendingWorkflowTrigger) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciPendingWorkflowTrigger) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTriggerTime

`func (o *HciPendingWorkflowTrigger) GetTriggerTime() time.Time`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *HciPendingWorkflowTrigger) GetTriggerTimeOk() (*time.Time, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *HciPendingWorkflowTrigger) SetTriggerTime(v time.Time)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *HciPendingWorkflowTrigger) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetWorkflowType

`func (o *HciPendingWorkflowTrigger) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *HciPendingWorkflowTrigger) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *HciPendingWorkflowTrigger) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *HciPendingWorkflowTrigger) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


