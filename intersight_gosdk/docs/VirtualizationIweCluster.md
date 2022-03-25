# VirtualizationIweCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweCluster"]
**CapacityRunway** | Pointer to **int64** | The number of days remaining before the cluster&#39;s storage utilization reaches the recommended capacity limit of 76%. Default value is math.MaxInt32 to indicate that the capacity runway is \&quot;Unknown\&quot; for a cluster that is not connected or with not sufficient data. | [optional] [readonly] [default to 2147483647]
**ClusterName** | Pointer to **string** | The name of this HyperFlex cluster. | [optional] [readonly] 
**ComputeNodeCount** | Pointer to **int64** | The number of compute nodes that belong to this cluster. | [optional] [readonly] 
**ConfiguredCpuOverSubFactor** | Pointer to **float64** | CPU oversubscription factor configured on the cluster. | [optional] 
**ConvergedNodeCount** | Pointer to **int64** | The number of converged nodes that belong to this cluster. | [optional] [readonly] 
**CpuAllocation** | Pointer to [**NullableVirtualizationCpuAllocation**](VirtualizationCpuAllocation.md) |  | [optional] 
**CurrentCpuOverSubFactor** | Pointer to **float64** | Current oversubscription factor of the cluster. | [optional] 
**DatacenterName** | Pointer to **string** | Datacenter to which the cluster belongs. | [optional] 
**DeploymentType** | Pointer to **string** | The deployment type of the HyperFlex cluster. The cluster can have one of the following configurations: 1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site. 2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites. 3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes. If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined, the deployment type is set as &#39;NA&#39; (not available). * &#x60;NA&#x60; - The deployment type of the cluster is not available. * &#x60;Datacenter&#x60; - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes on the same site. * &#x60;Stretched Cluster&#x60; - The deployment type of a cluster consisting of UCS Fabric Interconnect-attached nodes across different sites. * &#x60;Edge&#x60; - The deployment type of a cluster consisting of 2 or more standalone nodes. | [optional] [readonly] [default to "NA"]
**DriveType** | Pointer to **string** | The type of the drives used for storage in this cluster. * &#x60;NA&#x60; - The drive type of the cluster is not available. * &#x60;All-Flash&#x60; - Indicates that this cluster contains flash drives only. * &#x60;Hybrid&#x60; - Indicates that this cluster contains both flash and hard disk drives. | [optional] [readonly] [default to "NA"]
**FailureReason** | Pointer to **string** | Reason for the failure when cluster is in failed state. | [optional] 
**HxVersion** | Pointer to **string** | The HyperFlex Data or Application Platform version of this cluster. | [optional] [readonly] 
**HypervisorBuild** | Pointer to **string** | Hypervisor version of HyperFlex compute cluster along with build number. | [optional] 
**ManagementIpAddress** | Pointer to **string** | Management IP Address of the cluster. | [optional] 
**MemoryAllocation** | Pointer to [**NullableVirtualizationMemoryAllocation**](VirtualizationMemoryAllocation.md) |  | [optional] 
**StorageCapacity** | Pointer to **int64** | The storage capacity in this cluster. | [optional] [readonly] 
**StorageNodeCount** | Pointer to **int64** | The number of storage nodes that belong to this cluster. | [optional] [readonly] 
**StorageUtilization** | Pointer to **float32** | The storage utilization is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationPercentage** | Pointer to **float32** | The storage utilization percentage is computed based on total capacity and current capacity utilization. | [optional] [readonly] 
**UtilizationTrendPercentage** | Pointer to **float32** | The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data. | [optional] [readonly] 
**AssociatedProfile** | Pointer to [**PolicyAbstractProfileRelationship**](PolicyAbstractProfileRelationship.md) |  | [optional] 
**HxCluster** | Pointer to [**StorageBaseClusterRelationship**](StorageBaseClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweCluster

`func NewVirtualizationIweCluster(classId string, objectType string, ) *VirtualizationIweCluster`

NewVirtualizationIweCluster instantiates a new VirtualizationIweCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweClusterWithDefaults

`func NewVirtualizationIweClusterWithDefaults() *VirtualizationIweCluster`

NewVirtualizationIweClusterWithDefaults instantiates a new VirtualizationIweCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacityRunway

`func (o *VirtualizationIweCluster) GetCapacityRunway() int64`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *VirtualizationIweCluster) GetCapacityRunwayOk() (*int64, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *VirtualizationIweCluster) SetCapacityRunway(v int64)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *VirtualizationIweCluster) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetClusterName

`func (o *VirtualizationIweCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VirtualizationIweCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VirtualizationIweCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VirtualizationIweCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetComputeNodeCount

`func (o *VirtualizationIweCluster) GetComputeNodeCount() int64`

GetComputeNodeCount returns the ComputeNodeCount field if non-nil, zero value otherwise.

### GetComputeNodeCountOk

`func (o *VirtualizationIweCluster) GetComputeNodeCountOk() (*int64, bool)`

GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodeCount

`func (o *VirtualizationIweCluster) SetComputeNodeCount(v int64)`

SetComputeNodeCount sets ComputeNodeCount field to given value.

### HasComputeNodeCount

`func (o *VirtualizationIweCluster) HasComputeNodeCount() bool`

HasComputeNodeCount returns a boolean if a field has been set.

### GetConfiguredCpuOverSubFactor

`func (o *VirtualizationIweCluster) GetConfiguredCpuOverSubFactor() float64`

GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field if non-nil, zero value otherwise.

### GetConfiguredCpuOverSubFactorOk

`func (o *VirtualizationIweCluster) GetConfiguredCpuOverSubFactorOk() (*float64, bool)`

GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredCpuOverSubFactor

`func (o *VirtualizationIweCluster) SetConfiguredCpuOverSubFactor(v float64)`

SetConfiguredCpuOverSubFactor sets ConfiguredCpuOverSubFactor field to given value.

### HasConfiguredCpuOverSubFactor

`func (o *VirtualizationIweCluster) HasConfiguredCpuOverSubFactor() bool`

HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.

### GetConvergedNodeCount

`func (o *VirtualizationIweCluster) GetConvergedNodeCount() int64`

GetConvergedNodeCount returns the ConvergedNodeCount field if non-nil, zero value otherwise.

### GetConvergedNodeCountOk

`func (o *VirtualizationIweCluster) GetConvergedNodeCountOk() (*int64, bool)`

GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergedNodeCount

`func (o *VirtualizationIweCluster) SetConvergedNodeCount(v int64)`

SetConvergedNodeCount sets ConvergedNodeCount field to given value.

### HasConvergedNodeCount

`func (o *VirtualizationIweCluster) HasConvergedNodeCount() bool`

HasConvergedNodeCount returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *VirtualizationIweCluster) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VirtualizationIweCluster) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VirtualizationIweCluster) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VirtualizationIweCluster) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *VirtualizationIweCluster) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *VirtualizationIweCluster) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetCurrentCpuOverSubFactor

`func (o *VirtualizationIweCluster) GetCurrentCpuOverSubFactor() float64`

GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field if non-nil, zero value otherwise.

### GetCurrentCpuOverSubFactorOk

`func (o *VirtualizationIweCluster) GetCurrentCpuOverSubFactorOk() (*float64, bool)`

GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuOverSubFactor

`func (o *VirtualizationIweCluster) SetCurrentCpuOverSubFactor(v float64)`

SetCurrentCpuOverSubFactor sets CurrentCpuOverSubFactor field to given value.

### HasCurrentCpuOverSubFactor

`func (o *VirtualizationIweCluster) HasCurrentCpuOverSubFactor() bool`

HasCurrentCpuOverSubFactor returns a boolean if a field has been set.

### GetDatacenterName

`func (o *VirtualizationIweCluster) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *VirtualizationIweCluster) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *VirtualizationIweCluster) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *VirtualizationIweCluster) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetDeploymentType

`func (o *VirtualizationIweCluster) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *VirtualizationIweCluster) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *VirtualizationIweCluster) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *VirtualizationIweCluster) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDriveType

`func (o *VirtualizationIweCluster) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *VirtualizationIweCluster) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *VirtualizationIweCluster) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *VirtualizationIweCluster) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetFailureReason

`func (o *VirtualizationIweCluster) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VirtualizationIweCluster) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VirtualizationIweCluster) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VirtualizationIweCluster) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHxVersion

`func (o *VirtualizationIweCluster) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *VirtualizationIweCluster) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *VirtualizationIweCluster) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *VirtualizationIweCluster) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHypervisorBuild

`func (o *VirtualizationIweCluster) GetHypervisorBuild() string`

GetHypervisorBuild returns the HypervisorBuild field if non-nil, zero value otherwise.

### GetHypervisorBuildOk

`func (o *VirtualizationIweCluster) GetHypervisorBuildOk() (*string, bool)`

GetHypervisorBuildOk returns a tuple with the HypervisorBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorBuild

`func (o *VirtualizationIweCluster) SetHypervisorBuild(v string)`

SetHypervisorBuild sets HypervisorBuild field to given value.

### HasHypervisorBuild

`func (o *VirtualizationIweCluster) HasHypervisorBuild() bool`

HasHypervisorBuild returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *VirtualizationIweCluster) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *VirtualizationIweCluster) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *VirtualizationIweCluster) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *VirtualizationIweCluster) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VirtualizationIweCluster) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VirtualizationIweCluster) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VirtualizationIweCluster) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VirtualizationIweCluster) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *VirtualizationIweCluster) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *VirtualizationIweCluster) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetStorageCapacity

`func (o *VirtualizationIweCluster) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *VirtualizationIweCluster) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *VirtualizationIweCluster) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *VirtualizationIweCluster) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### GetStorageNodeCount

`func (o *VirtualizationIweCluster) GetStorageNodeCount() int64`

GetStorageNodeCount returns the StorageNodeCount field if non-nil, zero value otherwise.

### GetStorageNodeCountOk

`func (o *VirtualizationIweCluster) GetStorageNodeCountOk() (*int64, bool)`

GetStorageNodeCountOk returns a tuple with the StorageNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNodeCount

`func (o *VirtualizationIweCluster) SetStorageNodeCount(v int64)`

SetStorageNodeCount sets StorageNodeCount field to given value.

### HasStorageNodeCount

`func (o *VirtualizationIweCluster) HasStorageNodeCount() bool`

HasStorageNodeCount returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *VirtualizationIweCluster) GetStorageUtilization() float32`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *VirtualizationIweCluster) GetStorageUtilizationOk() (*float32, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *VirtualizationIweCluster) SetStorageUtilization(v float32)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *VirtualizationIweCluster) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetUtilizationPercentage

`func (o *VirtualizationIweCluster) GetUtilizationPercentage() float32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *VirtualizationIweCluster) GetUtilizationPercentageOk() (*float32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *VirtualizationIweCluster) SetUtilizationPercentage(v float32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.

### HasUtilizationPercentage

`func (o *VirtualizationIweCluster) HasUtilizationPercentage() bool`

HasUtilizationPercentage returns a boolean if a field has been set.

### GetUtilizationTrendPercentage

`func (o *VirtualizationIweCluster) GetUtilizationTrendPercentage() float32`

GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field if non-nil, zero value otherwise.

### GetUtilizationTrendPercentageOk

`func (o *VirtualizationIweCluster) GetUtilizationTrendPercentageOk() (*float32, bool)`

GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTrendPercentage

`func (o *VirtualizationIweCluster) SetUtilizationTrendPercentage(v float32)`

SetUtilizationTrendPercentage sets UtilizationTrendPercentage field to given value.

### HasUtilizationTrendPercentage

`func (o *VirtualizationIweCluster) HasUtilizationTrendPercentage() bool`

HasUtilizationTrendPercentage returns a boolean if a field has been set.

### GetAssociatedProfile

`func (o *VirtualizationIweCluster) GetAssociatedProfile() PolicyAbstractProfileRelationship`

GetAssociatedProfile returns the AssociatedProfile field if non-nil, zero value otherwise.

### GetAssociatedProfileOk

`func (o *VirtualizationIweCluster) GetAssociatedProfileOk() (*PolicyAbstractProfileRelationship, bool)`

GetAssociatedProfileOk returns a tuple with the AssociatedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedProfile

`func (o *VirtualizationIweCluster) SetAssociatedProfile(v PolicyAbstractProfileRelationship)`

SetAssociatedProfile sets AssociatedProfile field to given value.

### HasAssociatedProfile

`func (o *VirtualizationIweCluster) HasAssociatedProfile() bool`

HasAssociatedProfile returns a boolean if a field has been set.

### GetHxCluster

`func (o *VirtualizationIweCluster) GetHxCluster() StorageBaseClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *VirtualizationIweCluster) GetHxClusterOk() (*StorageBaseClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *VirtualizationIweCluster) SetHxCluster(v StorageBaseClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *VirtualizationIweCluster) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationIweCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationIweCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationIweCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationIweCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


