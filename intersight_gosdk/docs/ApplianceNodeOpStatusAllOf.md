# ApplianceNodeOpStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeOpStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeOpStatus"]
**ClusterNetworkStatus** | Pointer to [**[]ApplianceNetworkStatus**](ApplianceNetworkStatus.md) |  | [optional] 
**CpuUsage** | Pointer to **float32** | Percentage of CPU currently in use. | [optional] [readonly] 
**MemUsage** | Pointer to **float32** | Percentage of memory currently in use. | [optional] [readonly] 
**NodeHostname** | Pointer to **string** | Hostname of the Intersight Appliance node. | [optional] [readonly] 
**NodeId** | Pointer to **int64** | System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1. | [optional] [readonly] 
**NodeState** | Pointer to **string** | State of the node in terms of its readiness to host Kubernetes pods. * &#x60;Down&#x60; - The node is yet to come up and join as a member of theKubernetes cluster. * &#x60;Preparing&#x60; - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * &#x60;Ready&#x60; - The node is ready to host Kubernetes pods. | [optional] [readonly] [default to "Down"]
**OperationalStatus** | Pointer to **string** | Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * &#x60;Unknown&#x60; - The status of the appliance node is unknown. * &#x60;Operational&#x60; - The appliance node is operational. * &#x60;Impaired&#x60; - The appliance node is impaired. * &#x60;AttentionNeeded&#x60; - The appliance node needs attention. * &#x60;ReadyToJoin&#x60; - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * &#x60;OutOfService&#x60; - The user has taken this node (part of a cluster) to out of service. * &#x60;ReadyForReplacement&#x60; - The cluster node is ready to be replaced. * &#x60;ReplacementInProgress&#x60; - The cluster node replacement is in progress. * &#x60;ReplacementFailed&#x60; - There was a failure during the cluster node replacement. | [optional] [readonly] [default to "Unknown"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**FileSystemOpStatuses** | Pointer to [**[]ApplianceFileSystemOpStatusRelationship**](ApplianceFileSystemOpStatusRelationship.md) | An array of relationships to applianceFileSystemOpStatus resources. | [optional] [readonly] 
**NetworkLinkStatuses** | Pointer to [**[]ApplianceNetworkLinkStatusRelationship**](ApplianceNetworkLinkStatusRelationship.md) | An array of relationships to applianceNetworkLinkStatus resources. | [optional] [readonly] 
**NodeInfo** | Pointer to [**ApplianceNodeInfoRelationship**](ApplianceNodeInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SystemOpStatus** | Pointer to [**ApplianceSystemOpStatusRelationship**](ApplianceSystemOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceNodeOpStatusAllOf

`func NewApplianceNodeOpStatusAllOf(classId string, objectType string, ) *ApplianceNodeOpStatusAllOf`

NewApplianceNodeOpStatusAllOf instantiates a new ApplianceNodeOpStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeOpStatusAllOfWithDefaults

`func NewApplianceNodeOpStatusAllOfWithDefaults() *ApplianceNodeOpStatusAllOf`

NewApplianceNodeOpStatusAllOfWithDefaults instantiates a new ApplianceNodeOpStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeOpStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeOpStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeOpStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeOpStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeOpStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeOpStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterNetworkStatus

`func (o *ApplianceNodeOpStatusAllOf) GetClusterNetworkStatus() []ApplianceNetworkStatus`

GetClusterNetworkStatus returns the ClusterNetworkStatus field if non-nil, zero value otherwise.

### GetClusterNetworkStatusOk

`func (o *ApplianceNodeOpStatusAllOf) GetClusterNetworkStatusOk() (*[]ApplianceNetworkStatus, bool)`

GetClusterNetworkStatusOk returns a tuple with the ClusterNetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetworkStatus

`func (o *ApplianceNodeOpStatusAllOf) SetClusterNetworkStatus(v []ApplianceNetworkStatus)`

SetClusterNetworkStatus sets ClusterNetworkStatus field to given value.

### HasClusterNetworkStatus

`func (o *ApplianceNodeOpStatusAllOf) HasClusterNetworkStatus() bool`

HasClusterNetworkStatus returns a boolean if a field has been set.

### SetClusterNetworkStatusNil

`func (o *ApplianceNodeOpStatusAllOf) SetClusterNetworkStatusNil(b bool)`

 SetClusterNetworkStatusNil sets the value for ClusterNetworkStatus to be an explicit nil

### UnsetClusterNetworkStatus
`func (o *ApplianceNodeOpStatusAllOf) UnsetClusterNetworkStatus()`

UnsetClusterNetworkStatus ensures that no value is present for ClusterNetworkStatus, not even an explicit nil
### GetCpuUsage

`func (o *ApplianceNodeOpStatusAllOf) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ApplianceNodeOpStatusAllOf) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ApplianceNodeOpStatusAllOf) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ApplianceNodeOpStatusAllOf) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemUsage

`func (o *ApplianceNodeOpStatusAllOf) GetMemUsage() float32`

GetMemUsage returns the MemUsage field if non-nil, zero value otherwise.

### GetMemUsageOk

`func (o *ApplianceNodeOpStatusAllOf) GetMemUsageOk() (*float32, bool)`

GetMemUsageOk returns a tuple with the MemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsage

`func (o *ApplianceNodeOpStatusAllOf) SetMemUsage(v float32)`

SetMemUsage sets MemUsage field to given value.

### HasMemUsage

`func (o *ApplianceNodeOpStatusAllOf) HasMemUsage() bool`

HasMemUsage returns a boolean if a field has been set.

### GetNodeHostname

`func (o *ApplianceNodeOpStatusAllOf) GetNodeHostname() string`

GetNodeHostname returns the NodeHostname field if non-nil, zero value otherwise.

### GetNodeHostnameOk

`func (o *ApplianceNodeOpStatusAllOf) GetNodeHostnameOk() (*string, bool)`

GetNodeHostnameOk returns a tuple with the NodeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHostname

`func (o *ApplianceNodeOpStatusAllOf) SetNodeHostname(v string)`

SetNodeHostname sets NodeHostname field to given value.

### HasNodeHostname

`func (o *ApplianceNodeOpStatusAllOf) HasNodeHostname() bool`

HasNodeHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *ApplianceNodeOpStatusAllOf) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ApplianceNodeOpStatusAllOf) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ApplianceNodeOpStatusAllOf) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ApplianceNodeOpStatusAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeState

`func (o *ApplianceNodeOpStatusAllOf) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *ApplianceNodeOpStatusAllOf) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *ApplianceNodeOpStatusAllOf) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *ApplianceNodeOpStatusAllOf) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ApplianceNodeOpStatusAllOf) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ApplianceNodeOpStatusAllOf) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ApplianceNodeOpStatusAllOf) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ApplianceNodeOpStatusAllOf) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceNodeOpStatusAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceNodeOpStatusAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceNodeOpStatusAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceNodeOpStatusAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFileSystemOpStatuses

`func (o *ApplianceNodeOpStatusAllOf) GetFileSystemOpStatuses() []ApplianceFileSystemOpStatusRelationship`

GetFileSystemOpStatuses returns the FileSystemOpStatuses field if non-nil, zero value otherwise.

### GetFileSystemOpStatusesOk

`func (o *ApplianceNodeOpStatusAllOf) GetFileSystemOpStatusesOk() (*[]ApplianceFileSystemOpStatusRelationship, bool)`

GetFileSystemOpStatusesOk returns a tuple with the FileSystemOpStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemOpStatuses

`func (o *ApplianceNodeOpStatusAllOf) SetFileSystemOpStatuses(v []ApplianceFileSystemOpStatusRelationship)`

SetFileSystemOpStatuses sets FileSystemOpStatuses field to given value.

### HasFileSystemOpStatuses

`func (o *ApplianceNodeOpStatusAllOf) HasFileSystemOpStatuses() bool`

HasFileSystemOpStatuses returns a boolean if a field has been set.

### SetFileSystemOpStatusesNil

`func (o *ApplianceNodeOpStatusAllOf) SetFileSystemOpStatusesNil(b bool)`

 SetFileSystemOpStatusesNil sets the value for FileSystemOpStatuses to be an explicit nil

### UnsetFileSystemOpStatuses
`func (o *ApplianceNodeOpStatusAllOf) UnsetFileSystemOpStatuses()`

UnsetFileSystemOpStatuses ensures that no value is present for FileSystemOpStatuses, not even an explicit nil
### GetNetworkLinkStatuses

`func (o *ApplianceNodeOpStatusAllOf) GetNetworkLinkStatuses() []ApplianceNetworkLinkStatusRelationship`

GetNetworkLinkStatuses returns the NetworkLinkStatuses field if non-nil, zero value otherwise.

### GetNetworkLinkStatusesOk

`func (o *ApplianceNodeOpStatusAllOf) GetNetworkLinkStatusesOk() (*[]ApplianceNetworkLinkStatusRelationship, bool)`

GetNetworkLinkStatusesOk returns a tuple with the NetworkLinkStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkStatuses

`func (o *ApplianceNodeOpStatusAllOf) SetNetworkLinkStatuses(v []ApplianceNetworkLinkStatusRelationship)`

SetNetworkLinkStatuses sets NetworkLinkStatuses field to given value.

### HasNetworkLinkStatuses

`func (o *ApplianceNodeOpStatusAllOf) HasNetworkLinkStatuses() bool`

HasNetworkLinkStatuses returns a boolean if a field has been set.

### SetNetworkLinkStatusesNil

`func (o *ApplianceNodeOpStatusAllOf) SetNetworkLinkStatusesNil(b bool)`

 SetNetworkLinkStatusesNil sets the value for NetworkLinkStatuses to be an explicit nil

### UnsetNetworkLinkStatuses
`func (o *ApplianceNodeOpStatusAllOf) UnsetNetworkLinkStatuses()`

UnsetNetworkLinkStatuses ensures that no value is present for NetworkLinkStatuses, not even an explicit nil
### GetNodeInfo

`func (o *ApplianceNodeOpStatusAllOf) GetNodeInfo() ApplianceNodeInfoRelationship`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *ApplianceNodeOpStatusAllOf) GetNodeInfoOk() (*ApplianceNodeInfoRelationship, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *ApplianceNodeOpStatusAllOf) SetNodeInfo(v ApplianceNodeInfoRelationship)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *ApplianceNodeOpStatusAllOf) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceNodeOpStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceNodeOpStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceNodeOpStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceNodeOpStatusAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetSystemOpStatus

`func (o *ApplianceNodeOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship`

GetSystemOpStatus returns the SystemOpStatus field if non-nil, zero value otherwise.

### GetSystemOpStatusOk

`func (o *ApplianceNodeOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool)`

GetSystemOpStatusOk returns a tuple with the SystemOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemOpStatus

`func (o *ApplianceNodeOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship)`

SetSystemOpStatus sets SystemOpStatus field to given value.

### HasSystemOpStatus

`func (o *ApplianceNodeOpStatusAllOf) HasSystemOpStatus() bool`

HasSystemOpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


