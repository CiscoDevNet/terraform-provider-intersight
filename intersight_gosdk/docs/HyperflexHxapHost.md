# HyperflexHxapHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapHost"]
**Age** | Pointer to **string** | Denotes age or life time of the Host in nano seconds. | [optional] 
**ChassisVersion** | Pointer to **string** | Chassis version of the Host. | [optional] 
**ClusterUuid** | Pointer to **string** | The UUID of the cluster to which this Host belongs to. | [optional] 
**CpuAllocation** | Pointer to [**NullableVirtualizationCpuAllocation**](VirtualizationCpuAllocation.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when host is in failed state. | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**InternalIpAddress** | Pointer to **string** | Internal IP Address of the Host. | [optional] 
**ManagementIpAddress** | Pointer to **string** | Management IP Address of the Host. | [optional] 
**MasterRole** | Pointer to **bool** | Is the role of this host is master in the cluster? true or false. | [optional] 
**MemoryAllocation** | Pointer to [**NullableVirtualizationMemoryAllocation**](VirtualizationMemoryAllocation.md) |  | [optional] 
**StorageVmPowerState** | Pointer to **string** | Is the Storage Controller VM on the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**Version** | Pointer to **string** | Product version of the Host. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapHost

`func NewHyperflexHxapHost(classId string, objectType string, ) *HyperflexHxapHost`

NewHyperflexHxapHost instantiates a new HyperflexHxapHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostWithDefaults

`func NewHyperflexHxapHostWithDefaults() *HyperflexHxapHost`

NewHyperflexHxapHostWithDefaults instantiates a new HyperflexHxapHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapHost) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapHost) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapHost) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapHost) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapHost) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapHost) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAge

`func (o *HyperflexHxapHost) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapHost) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapHost) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapHost) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetChassisVersion

`func (o *HyperflexHxapHost) GetChassisVersion() string`

GetChassisVersion returns the ChassisVersion field if non-nil, zero value otherwise.

### GetChassisVersionOk

`func (o *HyperflexHxapHost) GetChassisVersionOk() (*string, bool)`

GetChassisVersionOk returns a tuple with the ChassisVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisVersion

`func (o *HyperflexHxapHost) SetChassisVersion(v string)`

SetChassisVersion sets ChassisVersion field to given value.

### HasChassisVersion

`func (o *HyperflexHxapHost) HasChassisVersion() bool`

HasChassisVersion returns a boolean if a field has been set.

### GetClusterUuid

`func (o *HyperflexHxapHost) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *HyperflexHxapHost) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *HyperflexHxapHost) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *HyperflexHxapHost) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *HyperflexHxapHost) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *HyperflexHxapHost) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *HyperflexHxapHost) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *HyperflexHxapHost) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *HyperflexHxapHost) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *HyperflexHxapHost) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetFailureReason

`func (o *HyperflexHxapHost) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapHost) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapHost) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapHost) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHwPowerState

`func (o *HyperflexHxapHost) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *HyperflexHxapHost) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *HyperflexHxapHost) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *HyperflexHxapHost) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetInternalIpAddress

`func (o *HyperflexHxapHost) GetInternalIpAddress() string`

GetInternalIpAddress returns the InternalIpAddress field if non-nil, zero value otherwise.

### GetInternalIpAddressOk

`func (o *HyperflexHxapHost) GetInternalIpAddressOk() (*string, bool)`

GetInternalIpAddressOk returns a tuple with the InternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIpAddress

`func (o *HyperflexHxapHost) SetInternalIpAddress(v string)`

SetInternalIpAddress sets InternalIpAddress field to given value.

### HasInternalIpAddress

`func (o *HyperflexHxapHost) HasInternalIpAddress() bool`

HasInternalIpAddress returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *HyperflexHxapHost) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *HyperflexHxapHost) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *HyperflexHxapHost) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *HyperflexHxapHost) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMasterRole

`func (o *HyperflexHxapHost) GetMasterRole() bool`

GetMasterRole returns the MasterRole field if non-nil, zero value otherwise.

### GetMasterRoleOk

`func (o *HyperflexHxapHost) GetMasterRoleOk() (*bool, bool)`

GetMasterRoleOk returns a tuple with the MasterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRole

`func (o *HyperflexHxapHost) SetMasterRole(v bool)`

SetMasterRole sets MasterRole field to given value.

### HasMasterRole

`func (o *HyperflexHxapHost) HasMasterRole() bool`

HasMasterRole returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *HyperflexHxapHost) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *HyperflexHxapHost) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *HyperflexHxapHost) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *HyperflexHxapHost) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *HyperflexHxapHost) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *HyperflexHxapHost) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetStorageVmPowerState

`func (o *HyperflexHxapHost) GetStorageVmPowerState() string`

GetStorageVmPowerState returns the StorageVmPowerState field if non-nil, zero value otherwise.

### GetStorageVmPowerStateOk

`func (o *HyperflexHxapHost) GetStorageVmPowerStateOk() (*string, bool)`

GetStorageVmPowerStateOk returns a tuple with the StorageVmPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVmPowerState

`func (o *HyperflexHxapHost) SetStorageVmPowerState(v string)`

SetStorageVmPowerState sets StorageVmPowerState field to given value.

### HasStorageVmPowerState

`func (o *HyperflexHxapHost) HasStorageVmPowerState() bool`

HasStorageVmPowerState returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexHxapHost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHxapHost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHxapHost) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHxapHost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapHost) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHost) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHost) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHost) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *HyperflexHxapHost) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *HyperflexHxapHost) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *HyperflexHxapHost) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *HyperflexHxapHost) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *HyperflexHxapHost) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *HyperflexHxapHost) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *HyperflexHxapHost) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *HyperflexHxapHost) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


