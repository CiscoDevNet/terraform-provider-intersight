# HyperflexNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Node"]
**BuildNumber** | Pointer to **string** | The build number of the hypervisor running on the host. | [optional] [readonly] 
**DisplayVersion** | Pointer to **string** | The user-friendly string representation of the hypervisor version of the host. | [optional] [readonly] 
**EmptySlotsList** | Pointer to **[]string** |  | [optional] 
**HostName** | Pointer to **string** | The hostname configured for the hypervisor running on the host. | [optional] [readonly] 
**HxdpDataIp** | Pointer to [**NullableHyperflexHxNetworkAddressDt**](HyperflexHxNetworkAddressDt.md) |  | [optional] 
**HxdpMmgtIp** | Pointer to [**NullableHyperflexHxNetworkAddressDt**](HyperflexHxNetworkAddressDt.md) |  | [optional] 
**Hypervisor** | Pointer to **string** | The type of hypervisor running on the host. | [optional] [readonly] 
**HypervisorDataIp** | Pointer to [**NullableHyperflexHxNetworkAddressDt**](HyperflexHxNetworkAddressDt.md) |  | [optional] 
**Identity** | Pointer to [**NullableHyperflexHxUuIdDt**](HyperflexHxUuIdDt.md) |  | [optional] 
**Ip** | Pointer to [**NullableHyperflexHxNetworkAddressDt**](HyperflexHxNetworkAddressDt.md) |  | [optional] 
**Lockdown** | Pointer to **bool** | The admin state of lockdown mode on the host. If &#39;true&#39;, lockdown mode is enabled. | [optional] [readonly] 
**ModelNumber** | Pointer to **string** | The model of the host server. | [optional] [readonly] 
**NodeMaintenanceMode** | Pointer to **string** | The status of maintenance mode on the HyperFlex node. * &#x60;Unknown&#x60; - The maintenance mode status could not be determined. * &#x60;InMaintenanceMode&#x60; - The node has maintenance mode enabled. The node has been temporarily been relinquished from the cluster to allow for maintenance operations. * &#x60;NotInMaintenanceMode&#x60; - The node does not have maintenance mode enabled. | [optional] [readonly] [default to "Unknown"]
**NodeStatus** | Pointer to **string** | The operational status of the HyperFlex node. * &#x60;Unknown&#x60; - The default operational status of a HyperFlex node. * &#x60;Invalid&#x60; - The status of the node cannot be determined by the storage platform. * &#x60;Ready&#x60; - The platform node has been acknowledged by the cluster. * &#x60;Unpublished&#x60; - The node is not yet added to the storage cluster. * &#x60;Deleted&#x60; - The node has been removed from the cluster. * &#x60;Blocked&#x60; - The node is blocked from being added to the cluster. * &#x60;Blacklisted&#x60; - The deprecated value for &#39;Blocked&#39;. It is included to maintain backwards compatibility with clusters running a HyperFlex Data Platform version older than 5.0(1a). * &#x60;Allowed&#x60; - The node is allowed to be added to the cluster. * &#x60;Whitelisted&#x60; - The deprecated value for &#39;Allowed&#39;. It is included to maintain backwards compatibility with clusters running a HyperFlex Data Platform version older than 5.0(1a). * &#x60;InMaintenance&#x60; - The node is in maintenance mode. It has been temporarily relinquished from the cluster to allow for maintenance operations such as software upgrades. * &#x60;Online&#x60; - The node is participating in the storage cluster and is available for storage operations. * &#x60;Offline&#x60; - The node is part of the storage cluster, but is not available for storage operations. | [optional] [readonly] [default to "Unknown"]
**NodeUuid** | Pointer to **string** | The unique identifier of the HyperFlex node. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage. * &#x60;UNKNOWN&#x60; - The role of the HyperFlex cluster node is not known. * &#x60;STORAGE&#x60; - The HyperFlex cluster node provides both storage and compute resources for the cluster. * &#x60;COMPUTE&#x60; - The HyperFlex cluster node provides compute resources for the cluster. | [optional] [readonly] [default to "UNKNOWN"]
**SerialNumber** | Pointer to **string** | The serial of the host server. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the host. Indicates whether the hypervisor is online. * &#x60;UNKNOWN&#x60; - The host status cannot be determined. * &#x60;ONLINE&#x60; - The host is online and operational. * &#x60;OFFLINE&#x60; - The host is offline and is currently not participating in the HyperFlex cluster. * &#x60;INMAINTENANCE&#x60; - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade. * &#x60;DEGRADED&#x60; - The host is degraded and may not be performing in its full operational capacity. | [optional] [readonly] [default to "UNKNOWN"]
**Version** | Pointer to **string** | The version of the hypervisor running on the host. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) |  | [optional] 
**Drives** | Pointer to [**[]HyperflexDriveRelationship**](HyperflexDriveRelationship.md) | An array of relationships to hyperflexDrive resources. | [optional] [readonly] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewHyperflexNode

`func NewHyperflexNode(classId string, objectType string, ) *HyperflexNode`

NewHyperflexNode instantiates a new HyperflexNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeWithDefaults

`func NewHyperflexNodeWithDefaults() *HyperflexNode`

NewHyperflexNodeWithDefaults instantiates a new HyperflexNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBuildNumber

`func (o *HyperflexNode) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *HyperflexNode) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *HyperflexNode) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *HyperflexNode) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *HyperflexNode) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *HyperflexNode) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *HyperflexNode) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *HyperflexNode) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetEmptySlotsList

`func (o *HyperflexNode) GetEmptySlotsList() []string`

GetEmptySlotsList returns the EmptySlotsList field if non-nil, zero value otherwise.

### GetEmptySlotsListOk

`func (o *HyperflexNode) GetEmptySlotsListOk() (*[]string, bool)`

GetEmptySlotsListOk returns a tuple with the EmptySlotsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySlotsList

`func (o *HyperflexNode) SetEmptySlotsList(v []string)`

SetEmptySlotsList sets EmptySlotsList field to given value.

### HasEmptySlotsList

`func (o *HyperflexNode) HasEmptySlotsList() bool`

HasEmptySlotsList returns a boolean if a field has been set.

### SetEmptySlotsListNil

`func (o *HyperflexNode) SetEmptySlotsListNil(b bool)`

 SetEmptySlotsListNil sets the value for EmptySlotsList to be an explicit nil

### UnsetEmptySlotsList
`func (o *HyperflexNode) UnsetEmptySlotsList()`

UnsetEmptySlotsList ensures that no value is present for EmptySlotsList, not even an explicit nil
### GetHostName

`func (o *HyperflexNode) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexNode) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexNode) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexNode) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHxdpDataIp

`func (o *HyperflexNode) GetHxdpDataIp() HyperflexHxNetworkAddressDt`

GetHxdpDataIp returns the HxdpDataIp field if non-nil, zero value otherwise.

### GetHxdpDataIpOk

`func (o *HyperflexNode) GetHxdpDataIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetHxdpDataIpOk returns a tuple with the HxdpDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpDataIp

`func (o *HyperflexNode) SetHxdpDataIp(v HyperflexHxNetworkAddressDt)`

SetHxdpDataIp sets HxdpDataIp field to given value.

### HasHxdpDataIp

`func (o *HyperflexNode) HasHxdpDataIp() bool`

HasHxdpDataIp returns a boolean if a field has been set.

### SetHxdpDataIpNil

`func (o *HyperflexNode) SetHxdpDataIpNil(b bool)`

 SetHxdpDataIpNil sets the value for HxdpDataIp to be an explicit nil

### UnsetHxdpDataIp
`func (o *HyperflexNode) UnsetHxdpDataIp()`

UnsetHxdpDataIp ensures that no value is present for HxdpDataIp, not even an explicit nil
### GetHxdpMmgtIp

`func (o *HyperflexNode) GetHxdpMmgtIp() HyperflexHxNetworkAddressDt`

GetHxdpMmgtIp returns the HxdpMmgtIp field if non-nil, zero value otherwise.

### GetHxdpMmgtIpOk

`func (o *HyperflexNode) GetHxdpMmgtIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetHxdpMmgtIpOk returns a tuple with the HxdpMmgtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpMmgtIp

`func (o *HyperflexNode) SetHxdpMmgtIp(v HyperflexHxNetworkAddressDt)`

SetHxdpMmgtIp sets HxdpMmgtIp field to given value.

### HasHxdpMmgtIp

`func (o *HyperflexNode) HasHxdpMmgtIp() bool`

HasHxdpMmgtIp returns a boolean if a field has been set.

### SetHxdpMmgtIpNil

`func (o *HyperflexNode) SetHxdpMmgtIpNil(b bool)`

 SetHxdpMmgtIpNil sets the value for HxdpMmgtIp to be an explicit nil

### UnsetHxdpMmgtIp
`func (o *HyperflexNode) UnsetHxdpMmgtIp()`

UnsetHxdpMmgtIp ensures that no value is present for HxdpMmgtIp, not even an explicit nil
### GetHypervisor

`func (o *HyperflexNode) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HyperflexNode) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HyperflexNode) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *HyperflexNode) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetHypervisorDataIp

`func (o *HyperflexNode) GetHypervisorDataIp() HyperflexHxNetworkAddressDt`

GetHypervisorDataIp returns the HypervisorDataIp field if non-nil, zero value otherwise.

### GetHypervisorDataIpOk

`func (o *HyperflexNode) GetHypervisorDataIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorDataIp

`func (o *HyperflexNode) SetHypervisorDataIp(v HyperflexHxNetworkAddressDt)`

SetHypervisorDataIp sets HypervisorDataIp field to given value.

### HasHypervisorDataIp

`func (o *HyperflexNode) HasHypervisorDataIp() bool`

HasHypervisorDataIp returns a boolean if a field has been set.

### SetHypervisorDataIpNil

`func (o *HyperflexNode) SetHypervisorDataIpNil(b bool)`

 SetHypervisorDataIpNil sets the value for HypervisorDataIp to be an explicit nil

### UnsetHypervisorDataIp
`func (o *HyperflexNode) UnsetHypervisorDataIp()`

UnsetHypervisorDataIp ensures that no value is present for HypervisorDataIp, not even an explicit nil
### GetIdentity

`func (o *HyperflexNode) GetIdentity() HyperflexHxUuIdDt`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HyperflexNode) GetIdentityOk() (*HyperflexHxUuIdDt, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HyperflexNode) SetIdentity(v HyperflexHxUuIdDt)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HyperflexNode) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *HyperflexNode) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *HyperflexNode) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetIp

`func (o *HyperflexNode) GetIp() HyperflexHxNetworkAddressDt`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexNode) GetIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexNode) SetIp(v HyperflexHxNetworkAddressDt)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexNode) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HyperflexNode) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HyperflexNode) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLockdown

`func (o *HyperflexNode) GetLockdown() bool`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *HyperflexNode) GetLockdownOk() (*bool, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *HyperflexNode) SetLockdown(v bool)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *HyperflexNode) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetModelNumber

`func (o *HyperflexNode) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *HyperflexNode) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *HyperflexNode) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *HyperflexNode) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetNodeMaintenanceMode

`func (o *HyperflexNode) GetNodeMaintenanceMode() string`

GetNodeMaintenanceMode returns the NodeMaintenanceMode field if non-nil, zero value otherwise.

### GetNodeMaintenanceModeOk

`func (o *HyperflexNode) GetNodeMaintenanceModeOk() (*string, bool)`

GetNodeMaintenanceModeOk returns a tuple with the NodeMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMaintenanceMode

`func (o *HyperflexNode) SetNodeMaintenanceMode(v string)`

SetNodeMaintenanceMode sets NodeMaintenanceMode field to given value.

### HasNodeMaintenanceMode

`func (o *HyperflexNode) HasNodeMaintenanceMode() bool`

HasNodeMaintenanceMode returns a boolean if a field has been set.

### GetNodeStatus

`func (o *HyperflexNode) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *HyperflexNode) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *HyperflexNode) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *HyperflexNode) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetNodeUuid

`func (o *HyperflexNode) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *HyperflexNode) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *HyperflexNode) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *HyperflexNode) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetRole

`func (o *HyperflexNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HyperflexNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HyperflexNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *HyperflexNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HyperflexNode) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HyperflexNode) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HyperflexNode) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HyperflexNode) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexNode) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexNode) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexNode) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexNode) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexNode) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexNode) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexNode) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexNode) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetDrives

`func (o *HyperflexNode) GetDrives() []HyperflexDriveRelationship`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *HyperflexNode) GetDrivesOk() (*[]HyperflexDriveRelationship, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *HyperflexNode) SetDrives(v []HyperflexDriveRelationship)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *HyperflexNode) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrivesNil

`func (o *HyperflexNode) SetDrivesNil(b bool)`

 SetDrivesNil sets the value for Drives to be an explicit nil

### UnsetDrives
`func (o *HyperflexNode) UnsetDrives()`

UnsetDrives ensures that no value is present for Drives, not even an explicit nil
### GetPhysicalServer

`func (o *HyperflexNode) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexNode) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexNode) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexNode) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


