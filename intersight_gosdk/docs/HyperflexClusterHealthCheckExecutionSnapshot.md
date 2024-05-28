# HyperflexClusterHealthCheckExecutionSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterHealthCheckExecutionSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterHealthCheckExecutionSnapshot"]
**ExecutionContext** | Pointer to **string** | The execution context of the HyperFlex health checks. * &#x60;UNKNOWN&#x60; - The current context of HyperFlex health check execution is unknown. * &#x60;WORKFLOW&#x60; - The HyperFlex health check execution is initiated through an orchestration workflow. * &#x60;SCHEDULED&#x60; - The HyperFlex health check execution is through a scheduled run. | [optional] [default to "UNKNOWN"]
**Timestamp** | Pointer to **time.Time** | Timestamp of the last health check execution on the HyperFlex cluster. | [optional] 
**HxCluster** | Pointer to [**NullableHyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterHealthCheckExecutionSnapshot

`func NewHyperflexClusterHealthCheckExecutionSnapshot(classId string, objectType string, ) *HyperflexClusterHealthCheckExecutionSnapshot`

NewHyperflexClusterHealthCheckExecutionSnapshot instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults

`func NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults() *HyperflexClusterHealthCheckExecutionSnapshot`

NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetExecutionContext() string`

GetExecutionContext returns the ExecutionContext field if non-nil, zero value otherwise.

### GetExecutionContextOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetExecutionContextOk() (*string, bool)`

GetExecutionContextOk returns a tuple with the ExecutionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetExecutionContext(v string)`

SetExecutionContext sets ExecutionContext field to given value.

### HasExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasExecutionContext() bool`

HasExecutionContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### SetHxClusterNil

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetHxClusterNil(b bool)`

 SetHxClusterNil sets the value for HxCluster to be an explicit nil

### UnsetHxCluster
`func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetHxCluster()`

UnsetHxCluster ensures that no value is present for HxCluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


