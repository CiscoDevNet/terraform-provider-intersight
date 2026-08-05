# HciClusterOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.ClusterOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.ClusterOperation"]
**ClusterExtId** | Pointer to **string** | The unique identifier of the cluster. | [optional] [readonly] 
**ClusterOpInfos** | Pointer to [**[]HciClusterOpInfo**](HciClusterOpInfo.md) |  | [optional] 
**Name** | Pointer to **string** | The display name of the cluster. | [optional] [readonly] 
**NutanixClusterCheck** | Pointer to [**NullableHciNutanixClusterCheck**](HciNutanixClusterCheck.md) |  | [optional] 
**PendingWorkflowTriggers** | Pointer to [**[]HciPendingWorkflowTrigger**](HciPendingWorkflowTrigger.md) |  | [optional] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**RunningWorkflow** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewHciClusterOperation

`func NewHciClusterOperation(classId string, objectType string, ) *HciClusterOperation`

NewHciClusterOperation instantiates a new HciClusterOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciClusterOperationWithDefaults

`func NewHciClusterOperationWithDefaults() *HciClusterOperation`

NewHciClusterOperationWithDefaults instantiates a new HciClusterOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciClusterOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciClusterOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciClusterOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciClusterOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciClusterOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciClusterOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciClusterOperation) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciClusterOperation) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciClusterOperation) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciClusterOperation) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterOpInfos

`func (o *HciClusterOperation) GetClusterOpInfos() []HciClusterOpInfo`

GetClusterOpInfos returns the ClusterOpInfos field if non-nil, zero value otherwise.

### GetClusterOpInfosOk

`func (o *HciClusterOperation) GetClusterOpInfosOk() (*[]HciClusterOpInfo, bool)`

GetClusterOpInfosOk returns a tuple with the ClusterOpInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterOpInfos

`func (o *HciClusterOperation) SetClusterOpInfos(v []HciClusterOpInfo)`

SetClusterOpInfos sets ClusterOpInfos field to given value.

### HasClusterOpInfos

`func (o *HciClusterOperation) HasClusterOpInfos() bool`

HasClusterOpInfos returns a boolean if a field has been set.

### SetClusterOpInfosNil

`func (o *HciClusterOperation) SetClusterOpInfosNil(b bool)`

 SetClusterOpInfosNil sets the value for ClusterOpInfos to be an explicit nil

### UnsetClusterOpInfos
`func (o *HciClusterOperation) UnsetClusterOpInfos()`

UnsetClusterOpInfos ensures that no value is present for ClusterOpInfos, not even an explicit nil
### GetName

`func (o *HciClusterOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciClusterOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciClusterOperation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciClusterOperation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNutanixClusterCheck

`func (o *HciClusterOperation) GetNutanixClusterCheck() HciNutanixClusterCheck`

GetNutanixClusterCheck returns the NutanixClusterCheck field if non-nil, zero value otherwise.

### GetNutanixClusterCheckOk

`func (o *HciClusterOperation) GetNutanixClusterCheckOk() (*HciNutanixClusterCheck, bool)`

GetNutanixClusterCheckOk returns a tuple with the NutanixClusterCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNutanixClusterCheck

`func (o *HciClusterOperation) SetNutanixClusterCheck(v HciNutanixClusterCheck)`

SetNutanixClusterCheck sets NutanixClusterCheck field to given value.

### HasNutanixClusterCheck

`func (o *HciClusterOperation) HasNutanixClusterCheck() bool`

HasNutanixClusterCheck returns a boolean if a field has been set.

### SetNutanixClusterCheckNil

`func (o *HciClusterOperation) SetNutanixClusterCheckNil(b bool)`

 SetNutanixClusterCheckNil sets the value for NutanixClusterCheck to be an explicit nil

### UnsetNutanixClusterCheck
`func (o *HciClusterOperation) UnsetNutanixClusterCheck()`

UnsetNutanixClusterCheck ensures that no value is present for NutanixClusterCheck, not even an explicit nil
### GetPendingWorkflowTriggers

`func (o *HciClusterOperation) GetPendingWorkflowTriggers() []HciPendingWorkflowTrigger`

GetPendingWorkflowTriggers returns the PendingWorkflowTriggers field if non-nil, zero value otherwise.

### GetPendingWorkflowTriggersOk

`func (o *HciClusterOperation) GetPendingWorkflowTriggersOk() (*[]HciPendingWorkflowTrigger, bool)`

GetPendingWorkflowTriggersOk returns a tuple with the PendingWorkflowTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWorkflowTriggers

`func (o *HciClusterOperation) SetPendingWorkflowTriggers(v []HciPendingWorkflowTrigger)`

SetPendingWorkflowTriggers sets PendingWorkflowTriggers field to given value.

### HasPendingWorkflowTriggers

`func (o *HciClusterOperation) HasPendingWorkflowTriggers() bool`

HasPendingWorkflowTriggers returns a boolean if a field has been set.

### SetPendingWorkflowTriggersNil

`func (o *HciClusterOperation) SetPendingWorkflowTriggersNil(b bool)`

 SetPendingWorkflowTriggersNil sets the value for PendingWorkflowTriggers to be an explicit nil

### UnsetPendingWorkflowTriggers
`func (o *HciClusterOperation) UnsetPendingWorkflowTriggers()`

UnsetPendingWorkflowTriggers ensures that no value is present for PendingWorkflowTriggers, not even an explicit nil
### GetCluster

`func (o *HciClusterOperation) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciClusterOperation) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciClusterOperation) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciClusterOperation) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciClusterOperation) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciClusterOperation) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciClusterOperation) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciClusterOperation) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciClusterOperation) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciClusterOperation) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciClusterOperation) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciClusterOperation) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetRunningWorkflow

`func (o *HciClusterOperation) GetRunningWorkflow() WorkflowWorkflowInfoRelationship`

GetRunningWorkflow returns the RunningWorkflow field if non-nil, zero value otherwise.

### GetRunningWorkflowOk

`func (o *HciClusterOperation) GetRunningWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowOk returns a tuple with the RunningWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflow

`func (o *HciClusterOperation) SetRunningWorkflow(v WorkflowWorkflowInfoRelationship)`

SetRunningWorkflow sets RunningWorkflow field to given value.

### HasRunningWorkflow

`func (o *HciClusterOperation) HasRunningWorkflow() bool`

HasRunningWorkflow returns a boolean if a field has been set.

### SetRunningWorkflowNil

`func (o *HciClusterOperation) SetRunningWorkflowNil(b bool)`

 SetRunningWorkflowNil sets the value for RunningWorkflow to be an explicit nil

### UnsetRunningWorkflow
`func (o *HciClusterOperation) UnsetRunningWorkflow()`

UnsetRunningWorkflow ensures that no value is present for RunningWorkflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


