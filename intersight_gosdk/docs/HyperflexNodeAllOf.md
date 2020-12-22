# HyperflexNodeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Node"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Node"]
**BuildNumber** | Pointer to **string** | The build number of the hypervisor running on the host. | [optional] [readonly] 
**DisplayVersion** | Pointer to **string** | The user-friendly string representation of the hypervisor version of the host. | [optional] [readonly] 
**HostName** | Pointer to **string** | The hostname configured for the hypervisor running on the host. | [optional] [readonly] 
**Hypervisor** | Pointer to **string** | The type of hypervisor running on the host. | [optional] [readonly] 
**Identity** | Pointer to [**NullableHyperflexHxUuIdDt**](hyperflex.HxUuIdDt.md) |  | [optional] 
**Ip** | Pointer to [**NullableHyperflexHxNetworkAddressDt**](hyperflex.HxNetworkAddressDt.md) |  | [optional] 
**Lockdown** | Pointer to **bool** | The admin state of lockdown mode on the host. If &#39;true&#39;, lockdown mode is enabled. | [optional] [readonly] 
**ModelNumber** | Pointer to **string** | The model of the host server. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage. * &#x60;UNKNOWN&#x60; - The role of the HyperFlex cluster node is not known. * &#x60;STORAGE&#x60; - The HyperFlex cluster node provides both storage and compute resources for the cluster. * &#x60;COMPUTE&#x60; - The HyperFlex cluster node provides compute resources for the cluster. | [optional] [readonly] [default to "UNKNOWN"]
**SerialNumber** | Pointer to **string** | The serial of the host server. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the host. Indicates whether the hypervisor is online. * &#x60;UNKNOWN&#x60; - The host status cannot be determined. * &#x60;ONLINE&#x60; - The host is online and operational. * &#x60;OFFLINE&#x60; - The host is offline and is currently not participating in the HyperFlex cluster. * &#x60;INMAINTENANCE&#x60; - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade. * &#x60;DEGRADED&#x60; - The host is degraded and may not be performing in its full operational capacity. | [optional] [readonly] [default to "UNKNOWN"]
**Version** | Pointer to **string** | The version of the hypervisor running on the host. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeAllOf

`func NewHyperflexNodeAllOf(classId string, objectType string, ) *HyperflexNodeAllOf`

NewHyperflexNodeAllOf instantiates a new HyperflexNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeAllOfWithDefaults

`func NewHyperflexNodeAllOfWithDefaults() *HyperflexNodeAllOf`

NewHyperflexNodeAllOfWithDefaults instantiates a new HyperflexNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNodeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNodeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBuildNumber

`func (o *HyperflexNodeAllOf) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *HyperflexNodeAllOf) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *HyperflexNodeAllOf) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *HyperflexNodeAllOf) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *HyperflexNodeAllOf) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *HyperflexNodeAllOf) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *HyperflexNodeAllOf) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *HyperflexNodeAllOf) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetHostName

`func (o *HyperflexNodeAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexNodeAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexNodeAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexNodeAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHypervisor

`func (o *HyperflexNodeAllOf) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HyperflexNodeAllOf) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HyperflexNodeAllOf) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *HyperflexNodeAllOf) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetIdentity

`func (o *HyperflexNodeAllOf) GetIdentity() HyperflexHxUuIdDt`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HyperflexNodeAllOf) GetIdentityOk() (*HyperflexHxUuIdDt, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HyperflexNodeAllOf) SetIdentity(v HyperflexHxUuIdDt)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HyperflexNodeAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *HyperflexNodeAllOf) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *HyperflexNodeAllOf) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetIp

`func (o *HyperflexNodeAllOf) GetIp() HyperflexHxNetworkAddressDt`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexNodeAllOf) GetIpOk() (*HyperflexHxNetworkAddressDt, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexNodeAllOf) SetIp(v HyperflexHxNetworkAddressDt)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexNodeAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HyperflexNodeAllOf) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HyperflexNodeAllOf) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLockdown

`func (o *HyperflexNodeAllOf) GetLockdown() bool`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *HyperflexNodeAllOf) GetLockdownOk() (*bool, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *HyperflexNodeAllOf) SetLockdown(v bool)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *HyperflexNodeAllOf) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetModelNumber

`func (o *HyperflexNodeAllOf) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *HyperflexNodeAllOf) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *HyperflexNodeAllOf) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *HyperflexNodeAllOf) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetRole

`func (o *HyperflexNodeAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HyperflexNodeAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HyperflexNodeAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *HyperflexNodeAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HyperflexNodeAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HyperflexNodeAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HyperflexNodeAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HyperflexNodeAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexNodeAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexNodeAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexNodeAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexNodeAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexNodeAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexNodeAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexNodeAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexNodeAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexNodeAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexNodeAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexNodeAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexNodeAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexNodeAllOf) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexNodeAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexNodeAllOf) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexNodeAllOf) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *HyperflexNodeAllOf) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexNodeAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexNodeAllOf) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexNodeAllOf) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


