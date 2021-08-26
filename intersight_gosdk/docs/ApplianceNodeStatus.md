# ApplianceNodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeStatus"]
**CpuUsage** | Pointer to **float32** | Percentage of CPU currently in use. | [optional] [readonly] 
**MemUsage** | Pointer to **float32** | Percentage of memory currently in use. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. | [optional] [readonly] 
**NodeState** | Pointer to **string** | State of the node in terms of its readiness to host Kubernetes pods. * &#x60;Down&#x60; - The node is yet to come up and join as a member of theKubernetes cluster. * &#x60;Preparing&#x60; - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * &#x60;Ready&#x60; - The node is ready to host Kubernetes pods. | [optional] [readonly] [default to "Down"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - Operational status of the Intersight Appliance entity is Unknown. * &#x60;Operational&#x60; - Operational status of the Intersight Appliance entity is Operational. * &#x60;Impaired&#x60; - Operational status of the Intersight Appliance entity is Impaired. * &#x60;AttentionNeeded&#x60; - Operational status of the Intersight Appliance entity is AttentionNeeded. | [optional] [readonly] [default to "Unknown"]
**StatusChecks** | Pointer to [**[]ApplianceStatusCheck**](ApplianceStatusCheck.md) |  | [optional] 
**FileSystemStatuses** | Pointer to [**[]ApplianceFileSystemStatusRelationship**](ApplianceFileSystemStatusRelationship.md) | An array of relationships to applianceFileSystemStatus resources. | [optional] [readonly] 
**NodeInfo** | Pointer to [**ApplianceNodeInfoRelationship**](ApplianceNodeInfoRelationship.md) |  | [optional] 
**SystemStatus** | Pointer to [**ApplianceSystemStatusRelationship**](ApplianceSystemStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceNodeStatus

`func NewApplianceNodeStatus(classId string, objectType string, ) *ApplianceNodeStatus`

NewApplianceNodeStatus instantiates a new ApplianceNodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeStatusWithDefaults

`func NewApplianceNodeStatusWithDefaults() *ApplianceNodeStatus`

NewApplianceNodeStatusWithDefaults instantiates a new ApplianceNodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuUsage

`func (o *ApplianceNodeStatus) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ApplianceNodeStatus) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ApplianceNodeStatus) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ApplianceNodeStatus) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemUsage

`func (o *ApplianceNodeStatus) GetMemUsage() float32`

GetMemUsage returns the MemUsage field if non-nil, zero value otherwise.

### GetMemUsageOk

`func (o *ApplianceNodeStatus) GetMemUsageOk() (*float32, bool)`

GetMemUsageOk returns a tuple with the MemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsage

`func (o *ApplianceNodeStatus) SetMemUsage(v float32)`

SetMemUsage sets MemUsage field to given value.

### HasMemUsage

`func (o *ApplianceNodeStatus) HasMemUsage() bool`

HasMemUsage returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeStatus) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeStatus) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeStatus) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeStatus) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeState

`func (o *ApplianceNodeStatus) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *ApplianceNodeStatus) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *ApplianceNodeStatus) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *ApplianceNodeStatus) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceNodeStatus) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceNodeStatus) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceNodeStatus) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceNodeStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetStatusChecks

`func (o *ApplianceNodeStatus) GetStatusChecks() []ApplianceStatusCheck`

GetStatusChecks returns the StatusChecks field if non-nil, zero value otherwise.

### GetStatusChecksOk

`func (o *ApplianceNodeStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool)`

GetStatusChecksOk returns a tuple with the StatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecks

`func (o *ApplianceNodeStatus) SetStatusChecks(v []ApplianceStatusCheck)`

SetStatusChecks sets StatusChecks field to given value.

### HasStatusChecks

`func (o *ApplianceNodeStatus) HasStatusChecks() bool`

HasStatusChecks returns a boolean if a field has been set.

### SetStatusChecksNil

`func (o *ApplianceNodeStatus) SetStatusChecksNil(b bool)`

 SetStatusChecksNil sets the value for StatusChecks to be an explicit nil

### UnsetStatusChecks
`func (o *ApplianceNodeStatus) UnsetStatusChecks()`

UnsetStatusChecks ensures that no value is present for StatusChecks, not even an explicit nil
### GetFileSystemStatuses

`func (o *ApplianceNodeStatus) GetFileSystemStatuses() []ApplianceFileSystemStatusRelationship`

GetFileSystemStatuses returns the FileSystemStatuses field if non-nil, zero value otherwise.

### GetFileSystemStatusesOk

`func (o *ApplianceNodeStatus) GetFileSystemStatusesOk() (*[]ApplianceFileSystemStatusRelationship, bool)`

GetFileSystemStatusesOk returns a tuple with the FileSystemStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemStatuses

`func (o *ApplianceNodeStatus) SetFileSystemStatuses(v []ApplianceFileSystemStatusRelationship)`

SetFileSystemStatuses sets FileSystemStatuses field to given value.

### HasFileSystemStatuses

`func (o *ApplianceNodeStatus) HasFileSystemStatuses() bool`

HasFileSystemStatuses returns a boolean if a field has been set.

### SetFileSystemStatusesNil

`func (o *ApplianceNodeStatus) SetFileSystemStatusesNil(b bool)`

 SetFileSystemStatusesNil sets the value for FileSystemStatuses to be an explicit nil

### UnsetFileSystemStatuses
`func (o *ApplianceNodeStatus) UnsetFileSystemStatuses()`

UnsetFileSystemStatuses ensures that no value is present for FileSystemStatuses, not even an explicit nil
### GetNodeInfo

`func (o *ApplianceNodeStatus) GetNodeInfo() ApplianceNodeInfoRelationship`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *ApplianceNodeStatus) GetNodeInfoOk() (*ApplianceNodeInfoRelationship, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *ApplianceNodeStatus) SetNodeInfo(v ApplianceNodeInfoRelationship)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *ApplianceNodeStatus) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetSystemStatus

`func (o *ApplianceNodeStatus) GetSystemStatus() ApplianceSystemStatusRelationship`

GetSystemStatus returns the SystemStatus field if non-nil, zero value otherwise.

### GetSystemStatusOk

`func (o *ApplianceNodeStatus) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool)`

GetSystemStatusOk returns a tuple with the SystemStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemStatus

`func (o *ApplianceNodeStatus) SetSystemStatus(v ApplianceSystemStatusRelationship)`

SetSystemStatus sets SystemStatus field to given value.

### HasSystemStatus

`func (o *ApplianceNodeStatus) HasSystemStatus() bool`

HasSystemStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


