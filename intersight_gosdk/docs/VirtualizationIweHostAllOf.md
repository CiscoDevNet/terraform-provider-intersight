# VirtualizationIweHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweHost"]
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
**StorageCapacity** | Pointer to [**NullableVirtualizationStorageCapacity**](VirtualizationStorageCapacity.md) |  | [optional] 
**StorageVmPowerState** | Pointer to **string** | Is the Storage Controller VM on the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**Version** | Pointer to **string** | Product version of the Host. | [optional] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) |  | [optional] 
**PhysicalServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweHostAllOf

`func NewVirtualizationIweHostAllOf(classId string, objectType string, ) *VirtualizationIweHostAllOf`

NewVirtualizationIweHostAllOf instantiates a new VirtualizationIweHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweHostAllOfWithDefaults

`func NewVirtualizationIweHostAllOfWithDefaults() *VirtualizationIweHostAllOf`

NewVirtualizationIweHostAllOfWithDefaults instantiates a new VirtualizationIweHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAge

`func (o *VirtualizationIweHostAllOf) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VirtualizationIweHostAllOf) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VirtualizationIweHostAllOf) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *VirtualizationIweHostAllOf) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetChassisVersion

`func (o *VirtualizationIweHostAllOf) GetChassisVersion() string`

GetChassisVersion returns the ChassisVersion field if non-nil, zero value otherwise.

### GetChassisVersionOk

`func (o *VirtualizationIweHostAllOf) GetChassisVersionOk() (*string, bool)`

GetChassisVersionOk returns a tuple with the ChassisVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisVersion

`func (o *VirtualizationIweHostAllOf) SetChassisVersion(v string)`

SetChassisVersion sets ChassisVersion field to given value.

### HasChassisVersion

`func (o *VirtualizationIweHostAllOf) HasChassisVersion() bool`

HasChassisVersion returns a boolean if a field has been set.

### GetClusterUuid

`func (o *VirtualizationIweHostAllOf) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *VirtualizationIweHostAllOf) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *VirtualizationIweHostAllOf) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *VirtualizationIweHostAllOf) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *VirtualizationIweHostAllOf) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VirtualizationIweHostAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VirtualizationIweHostAllOf) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VirtualizationIweHostAllOf) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *VirtualizationIweHostAllOf) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *VirtualizationIweHostAllOf) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetFailureReason

`func (o *VirtualizationIweHostAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VirtualizationIweHostAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VirtualizationIweHostAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VirtualizationIweHostAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHwPowerState

`func (o *VirtualizationIweHostAllOf) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *VirtualizationIweHostAllOf) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *VirtualizationIweHostAllOf) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *VirtualizationIweHostAllOf) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetInternalIpAddress

`func (o *VirtualizationIweHostAllOf) GetInternalIpAddress() string`

GetInternalIpAddress returns the InternalIpAddress field if non-nil, zero value otherwise.

### GetInternalIpAddressOk

`func (o *VirtualizationIweHostAllOf) GetInternalIpAddressOk() (*string, bool)`

GetInternalIpAddressOk returns a tuple with the InternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIpAddress

`func (o *VirtualizationIweHostAllOf) SetInternalIpAddress(v string)`

SetInternalIpAddress sets InternalIpAddress field to given value.

### HasInternalIpAddress

`func (o *VirtualizationIweHostAllOf) HasInternalIpAddress() bool`

HasInternalIpAddress returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *VirtualizationIweHostAllOf) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *VirtualizationIweHostAllOf) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *VirtualizationIweHostAllOf) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *VirtualizationIweHostAllOf) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMasterRole

`func (o *VirtualizationIweHostAllOf) GetMasterRole() bool`

GetMasterRole returns the MasterRole field if non-nil, zero value otherwise.

### GetMasterRoleOk

`func (o *VirtualizationIweHostAllOf) GetMasterRoleOk() (*bool, bool)`

GetMasterRoleOk returns a tuple with the MasterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRole

`func (o *VirtualizationIweHostAllOf) SetMasterRole(v bool)`

SetMasterRole sets MasterRole field to given value.

### HasMasterRole

`func (o *VirtualizationIweHostAllOf) HasMasterRole() bool`

HasMasterRole returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VirtualizationIweHostAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VirtualizationIweHostAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VirtualizationIweHostAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VirtualizationIweHostAllOf) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *VirtualizationIweHostAllOf) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *VirtualizationIweHostAllOf) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetStorageCapacity

`func (o *VirtualizationIweHostAllOf) GetStorageCapacity() VirtualizationStorageCapacity`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *VirtualizationIweHostAllOf) GetStorageCapacityOk() (*VirtualizationStorageCapacity, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *VirtualizationIweHostAllOf) SetStorageCapacity(v VirtualizationStorageCapacity)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *VirtualizationIweHostAllOf) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### SetStorageCapacityNil

`func (o *VirtualizationIweHostAllOf) SetStorageCapacityNil(b bool)`

 SetStorageCapacityNil sets the value for StorageCapacity to be an explicit nil

### UnsetStorageCapacity
`func (o *VirtualizationIweHostAllOf) UnsetStorageCapacity()`

UnsetStorageCapacity ensures that no value is present for StorageCapacity, not even an explicit nil
### GetStorageVmPowerState

`func (o *VirtualizationIweHostAllOf) GetStorageVmPowerState() string`

GetStorageVmPowerState returns the StorageVmPowerState field if non-nil, zero value otherwise.

### GetStorageVmPowerStateOk

`func (o *VirtualizationIweHostAllOf) GetStorageVmPowerStateOk() (*string, bool)`

GetStorageVmPowerStateOk returns a tuple with the StorageVmPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVmPowerState

`func (o *VirtualizationIweHostAllOf) SetStorageVmPowerState(v string)`

SetStorageVmPowerState sets StorageVmPowerState field to given value.

### HasStorageVmPowerState

`func (o *VirtualizationIweHostAllOf) HasStorageVmPowerState() bool`

HasStorageVmPowerState returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualizationIweHostAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualizationIweHostAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualizationIweHostAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualizationIweHostAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweHostAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweHostAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweHostAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweHostAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterMember

`func (o *VirtualizationIweHostAllOf) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *VirtualizationIweHostAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *VirtualizationIweHostAllOf) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *VirtualizationIweHostAllOf) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetPhysicalServer

`func (o *VirtualizationIweHostAllOf) GetPhysicalServer() ComputePhysicalRelationship`

GetPhysicalServer returns the PhysicalServer field if non-nil, zero value otherwise.

### GetPhysicalServerOk

`func (o *VirtualizationIweHostAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool)`

GetPhysicalServerOk returns a tuple with the PhysicalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalServer

`func (o *VirtualizationIweHostAllOf) SetPhysicalServer(v ComputePhysicalRelationship)`

SetPhysicalServer sets PhysicalServer field to given value.

### HasPhysicalServer

`func (o *VirtualizationIweHostAllOf) HasPhysicalServer() bool`

HasPhysicalServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


