# HyperflexHxapHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapHost"]
**Age** | Pointer to **string** | Denotes age or life time of the Host in nano seconds. | [optional] 
**ChassisVersion** | Pointer to **string** | Chassis version of the Host. | [optional] 
**ClusterUuid** | Pointer to **string** | The UUID of the cluster to which this Host belongs to. | [optional] 
**CpuAllocation** | Pointer to [**NullableVirtualizationCpuAllocation**](virtualization.CpuAllocation.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when host is in failed state. | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**InternalIpAddress** | Pointer to **string** | Internal IP Address of the Host. | [optional] 
**ManagementIpAddress** | Pointer to **string** | Management IP Address of the Host. | [optional] 
**MasterRole** | Pointer to **bool** | Is the role of this host is master in the cluster? true or false. | [optional] 
**MemoryAllocation** | Pointer to [**NullableVirtualizationMemoryAllocation**](virtualization.MemoryAllocation.md) |  | [optional] 
**Version** | Pointer to **string** | Product version of the Host. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapHostAllOf

`func NewHyperflexHxapHostAllOf(classId string, objectType string, ) *HyperflexHxapHostAllOf`

NewHyperflexHxapHostAllOf instantiates a new HyperflexHxapHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostAllOfWithDefaults

`func NewHyperflexHxapHostAllOfWithDefaults() *HyperflexHxapHostAllOf`

NewHyperflexHxapHostAllOfWithDefaults instantiates a new HyperflexHxapHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAge

`func (o *HyperflexHxapHostAllOf) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapHostAllOf) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapHostAllOf) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapHostAllOf) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetChassisVersion

`func (o *HyperflexHxapHostAllOf) GetChassisVersion() string`

GetChassisVersion returns the ChassisVersion field if non-nil, zero value otherwise.

### GetChassisVersionOk

`func (o *HyperflexHxapHostAllOf) GetChassisVersionOk() (*string, bool)`

GetChassisVersionOk returns a tuple with the ChassisVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisVersion

`func (o *HyperflexHxapHostAllOf) SetChassisVersion(v string)`

SetChassisVersion sets ChassisVersion field to given value.

### HasChassisVersion

`func (o *HyperflexHxapHostAllOf) HasChassisVersion() bool`

HasChassisVersion returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexHxapHostAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexHxapHostAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexHxapHostAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexHxapHostAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *HyperflexHxapHostAllOf) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *HyperflexHxapHostAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *HyperflexHxapHostAllOf) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *HyperflexHxapHostAllOf) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *HyperflexHxapHostAllOf) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *HyperflexHxapHostAllOf) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetFailureReason

`func (o *HyperflexHxapHostAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapHostAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapHostAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapHostAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHwPowerState

`func (o *HyperflexHxapHostAllOf) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *HyperflexHxapHostAllOf) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *HyperflexHxapHostAllOf) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *HyperflexHxapHostAllOf) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetInternalIpAddress

`func (o *HyperflexHxapHostAllOf) GetInternalIpAddress() string`

GetInternalIpAddress returns the InternalIpAddress field if non-nil, zero value otherwise.

### GetInternalIpAddressOk

`func (o *HyperflexHxapHostAllOf) GetInternalIpAddressOk() (*string, bool)`

GetInternalIpAddressOk returns a tuple with the InternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIpAddress

`func (o *HyperflexHxapHostAllOf) SetInternalIpAddress(v string)`

SetInternalIpAddress sets InternalIpAddress field to given value.

### HasInternalIpAddress

`func (o *HyperflexHxapHostAllOf) HasInternalIpAddress() bool`

HasInternalIpAddress returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *HyperflexHxapHostAllOf) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *HyperflexHxapHostAllOf) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *HyperflexHxapHostAllOf) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *HyperflexHxapHostAllOf) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMasterRole

`func (o *HyperflexHxapHostAllOf) GetMasterRole() bool`

GetMasterRole returns the MasterRole field if non-nil, zero value otherwise.

### GetMasterRoleOk

`func (o *HyperflexHxapHostAllOf) GetMasterRoleOk() (*bool, bool)`

GetMasterRoleOk returns a tuple with the MasterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRole

`func (o *HyperflexHxapHostAllOf) SetMasterRole(v bool)`

SetMasterRole sets MasterRole field to given value.

### HasMasterRole

`func (o *HyperflexHxapHostAllOf) HasMasterRole() bool`

HasMasterRole returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *HyperflexHxapHostAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *HyperflexHxapHostAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *HyperflexHxapHostAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *HyperflexHxapHostAllOf) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *HyperflexHxapHostAllOf) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *HyperflexHxapHostAllOf) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetVersion

`func (o *HyperflexHxapHostAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHxapHostAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHxapHostAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHxapHostAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapHostAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHostAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHostAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHostAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexHxapHostAllOf) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexHxapHostAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexHxapHostAllOf) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexHxapHostAllOf) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *HyperflexHxapHostAllOf) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexHxapHostAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexHxapHostAllOf) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexHxapHostAllOf) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


