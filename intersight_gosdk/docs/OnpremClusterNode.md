# OnpremClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.ClusterNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.ClusterNode"]
**BootTime** | Pointer to **time.Time** | Start time of the device connector in the Intersight Appliance node. | [optional] [readonly] 
**ClusterMode** | Pointer to **string** | Mode of the appliance cluster. | [optional] [readonly] 
**CpuCount** | Pointer to **int64** | Number of CPUs configured for the Intersight Appliance node. | [optional] [readonly] 
**DeploymentType** | Pointer to **string** | Deployment type of the Intersight Appliance node. | [optional] [readonly] 
**Disks** | Pointer to [**[]OnpremResourceInfo**](OnpremResourceInfo.md) |  | [optional] 
**Hostname** | Pointer to **string** | The hostname or FQDN of the Intersight Appliance node. | [optional] [readonly] 
**Memory** | Pointer to [**NullableOnpremResourceInfo**](OnpremResourceInfo.md) |  | [optional] 
**NodeId** | Pointer to **int64** | Identifier of the Intersight Appliance node (one based). | [optional] [readonly] 
**PingErrorNodes** | Pointer to **[]string** |  | [optional] 
**PingOk** | Pointer to **bool** | Indicates if the node can ping other nodes in the Intersight Appliance cluster. The Ping operation is a high level application specific status check operation, not an ICMP ping between the hosts. | [optional] [readonly] 
**PrimaryNode** | Pointer to **bool** | Indicates if this node is the primary node. | [optional] [readonly] 
**RsyncErrorNodes** | Pointer to **[]string** |  | [optional] 
**RsyncOk** | Pointer to **bool** | Indicates if this node can rsync to other nodes. | [optional] [readonly] 
**Version** | Pointer to **string** | Current version of the device connector in the Intersight Appliance node. | [optional] [readonly] 
**VirtualEnvType** | Pointer to **string** | Virtual Env type of the Intersight Appliance node (ESXi, Hyper-V or KVM). | [optional] [readonly] 

## Methods

### NewOnpremClusterNode

`func NewOnpremClusterNode(classId string, objectType string, ) *OnpremClusterNode`

NewOnpremClusterNode instantiates a new OnpremClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremClusterNodeWithDefaults

`func NewOnpremClusterNodeWithDefaults() *OnpremClusterNode`

NewOnpremClusterNodeWithDefaults instantiates a new OnpremClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremClusterNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremClusterNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremClusterNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremClusterNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremClusterNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremClusterNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootTime

`func (o *OnpremClusterNode) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *OnpremClusterNode) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *OnpremClusterNode) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *OnpremClusterNode) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetClusterMode

`func (o *OnpremClusterNode) GetClusterMode() string`

GetClusterMode returns the ClusterMode field if non-nil, zero value otherwise.

### GetClusterModeOk

`func (o *OnpremClusterNode) GetClusterModeOk() (*string, bool)`

GetClusterModeOk returns a tuple with the ClusterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMode

`func (o *OnpremClusterNode) SetClusterMode(v string)`

SetClusterMode sets ClusterMode field to given value.

### HasClusterMode

`func (o *OnpremClusterNode) HasClusterMode() bool`

HasClusterMode returns a boolean if a field has been set.

### GetCpuCount

`func (o *OnpremClusterNode) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *OnpremClusterNode) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *OnpremClusterNode) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *OnpremClusterNode) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetDeploymentType

`func (o *OnpremClusterNode) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *OnpremClusterNode) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *OnpremClusterNode) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *OnpremClusterNode) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDisks

`func (o *OnpremClusterNode) GetDisks() []OnpremResourceInfo`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *OnpremClusterNode) GetDisksOk() (*[]OnpremResourceInfo, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *OnpremClusterNode) SetDisks(v []OnpremResourceInfo)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *OnpremClusterNode) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *OnpremClusterNode) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *OnpremClusterNode) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetHostname

`func (o *OnpremClusterNode) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OnpremClusterNode) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OnpremClusterNode) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *OnpremClusterNode) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMemory

`func (o *OnpremClusterNode) GetMemory() OnpremResourceInfo`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *OnpremClusterNode) GetMemoryOk() (*OnpremResourceInfo, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *OnpremClusterNode) SetMemory(v OnpremResourceInfo)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *OnpremClusterNode) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *OnpremClusterNode) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *OnpremClusterNode) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetNodeId

`func (o *OnpremClusterNode) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *OnpremClusterNode) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *OnpremClusterNode) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *OnpremClusterNode) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPingErrorNodes

`func (o *OnpremClusterNode) GetPingErrorNodes() []string`

GetPingErrorNodes returns the PingErrorNodes field if non-nil, zero value otherwise.

### GetPingErrorNodesOk

`func (o *OnpremClusterNode) GetPingErrorNodesOk() (*[]string, bool)`

GetPingErrorNodesOk returns a tuple with the PingErrorNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingErrorNodes

`func (o *OnpremClusterNode) SetPingErrorNodes(v []string)`

SetPingErrorNodes sets PingErrorNodes field to given value.

### HasPingErrorNodes

`func (o *OnpremClusterNode) HasPingErrorNodes() bool`

HasPingErrorNodes returns a boolean if a field has been set.

### SetPingErrorNodesNil

`func (o *OnpremClusterNode) SetPingErrorNodesNil(b bool)`

 SetPingErrorNodesNil sets the value for PingErrorNodes to be an explicit nil

### UnsetPingErrorNodes
`func (o *OnpremClusterNode) UnsetPingErrorNodes()`

UnsetPingErrorNodes ensures that no value is present for PingErrorNodes, not even an explicit nil
### GetPingOk

`func (o *OnpremClusterNode) GetPingOk() bool`

GetPingOk returns the PingOk field if non-nil, zero value otherwise.

### GetPingOkOk

`func (o *OnpremClusterNode) GetPingOkOk() (*bool, bool)`

GetPingOkOk returns a tuple with the PingOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOk

`func (o *OnpremClusterNode) SetPingOk(v bool)`

SetPingOk sets PingOk field to given value.

### HasPingOk

`func (o *OnpremClusterNode) HasPingOk() bool`

HasPingOk returns a boolean if a field has been set.

### GetPrimaryNode

`func (o *OnpremClusterNode) GetPrimaryNode() bool`

GetPrimaryNode returns the PrimaryNode field if non-nil, zero value otherwise.

### GetPrimaryNodeOk

`func (o *OnpremClusterNode) GetPrimaryNodeOk() (*bool, bool)`

GetPrimaryNodeOk returns a tuple with the PrimaryNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNode

`func (o *OnpremClusterNode) SetPrimaryNode(v bool)`

SetPrimaryNode sets PrimaryNode field to given value.

### HasPrimaryNode

`func (o *OnpremClusterNode) HasPrimaryNode() bool`

HasPrimaryNode returns a boolean if a field has been set.

### GetRsyncErrorNodes

`func (o *OnpremClusterNode) GetRsyncErrorNodes() []string`

GetRsyncErrorNodes returns the RsyncErrorNodes field if non-nil, zero value otherwise.

### GetRsyncErrorNodesOk

`func (o *OnpremClusterNode) GetRsyncErrorNodesOk() (*[]string, bool)`

GetRsyncErrorNodesOk returns a tuple with the RsyncErrorNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsyncErrorNodes

`func (o *OnpremClusterNode) SetRsyncErrorNodes(v []string)`

SetRsyncErrorNodes sets RsyncErrorNodes field to given value.

### HasRsyncErrorNodes

`func (o *OnpremClusterNode) HasRsyncErrorNodes() bool`

HasRsyncErrorNodes returns a boolean if a field has been set.

### SetRsyncErrorNodesNil

`func (o *OnpremClusterNode) SetRsyncErrorNodesNil(b bool)`

 SetRsyncErrorNodesNil sets the value for RsyncErrorNodes to be an explicit nil

### UnsetRsyncErrorNodes
`func (o *OnpremClusterNode) UnsetRsyncErrorNodes()`

UnsetRsyncErrorNodes ensures that no value is present for RsyncErrorNodes, not even an explicit nil
### GetRsyncOk

`func (o *OnpremClusterNode) GetRsyncOk() bool`

GetRsyncOk returns the RsyncOk field if non-nil, zero value otherwise.

### GetRsyncOkOk

`func (o *OnpremClusterNode) GetRsyncOkOk() (*bool, bool)`

GetRsyncOkOk returns a tuple with the RsyncOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsyncOk

`func (o *OnpremClusterNode) SetRsyncOk(v bool)`

SetRsyncOk sets RsyncOk field to given value.

### HasRsyncOk

`func (o *OnpremClusterNode) HasRsyncOk() bool`

HasRsyncOk returns a boolean if a field has been set.

### GetVersion

`func (o *OnpremClusterNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OnpremClusterNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OnpremClusterNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OnpremClusterNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVirtualEnvType

`func (o *OnpremClusterNode) GetVirtualEnvType() string`

GetVirtualEnvType returns the VirtualEnvType field if non-nil, zero value otherwise.

### GetVirtualEnvTypeOk

`func (o *OnpremClusterNode) GetVirtualEnvTypeOk() (*string, bool)`

GetVirtualEnvTypeOk returns a tuple with the VirtualEnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEnvType

`func (o *OnpremClusterNode) SetVirtualEnvType(v string)`

SetVirtualEnvType sets VirtualEnvType field to given value.

### HasVirtualEnvType

`func (o *OnpremClusterNode) HasVirtualEnvType() bool`

HasVirtualEnvType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


