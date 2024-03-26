# HyperflexClusterHealthCheckExecutionSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ClusterHealthCheckExecutionSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ClusterHealthCheckExecutionSnapshot"]
**ExecutionContext** | Pointer to **string** | The execution context of the HyperFlex health checks. * &#x60;UNKNOWN&#x60; - The current context of HyperFlex health check execution is unknown. * &#x60;WORKFLOW&#x60; - The HyperFlex health check execution is initiated through an orchestration workflow. * &#x60;SCHEDULED&#x60; - The HyperFlex health check execution is through a scheduled run. | [optional] [default to "UNKNOWN"]
**Timestamp** | Pointer to **time.Time** | Timestamp of the last health check execution on the HyperFlex cluster. | [optional] 
**HxCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterHealthCheckExecutionSnapshotAllOf

`func NewHyperflexClusterHealthCheckExecutionSnapshotAllOf(classId string, objectType string, ) *HyperflexClusterHealthCheckExecutionSnapshotAllOf`

NewHyperflexClusterHealthCheckExecutionSnapshotAllOf instantiates a new HyperflexClusterHealthCheckExecutionSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterHealthCheckExecutionSnapshotAllOfWithDefaults

`func NewHyperflexClusterHealthCheckExecutionSnapshotAllOfWithDefaults() *HyperflexClusterHealthCheckExecutionSnapshotAllOf`

NewHyperflexClusterHealthCheckExecutionSnapshotAllOfWithDefaults instantiates a new HyperflexClusterHealthCheckExecutionSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetExecutionContext() string`

GetExecutionContext returns the ExecutionContext field if non-nil, zero value otherwise.

### GetExecutionContextOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetExecutionContextOk() (*string, bool)`

GetExecutionContextOk returns a tuple with the ExecutionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetExecutionContext(v string)`

SetExecutionContext sets ExecutionContext field to given value.

### HasExecutionContext

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasExecutionContext() bool`

HasExecutionContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *HyperflexClusterHealthCheckExecutionSnapshotAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


