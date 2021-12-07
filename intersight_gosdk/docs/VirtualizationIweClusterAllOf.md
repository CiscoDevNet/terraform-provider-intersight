# VirtualizationIweClusterAllOf

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
**HypervisorVersion** | Pointer to **string** | The version of hypervisor running on this cluster. | [optional] [readonly] 
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

### NewVirtualizationIweClusterAllOf

`func NewVirtualizationIweClusterAllOf(classId string, objectType string, ) *VirtualizationIweClusterAllOf`

NewVirtualizationIweClusterAllOf instantiates a new VirtualizationIweClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweClusterAllOfWithDefaults

`func NewVirtualizationIweClusterAllOfWithDefaults() *VirtualizationIweClusterAllOf`

NewVirtualizationIweClusterAllOfWithDefaults instantiates a new VirtualizationIweClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacityRunway

`func (o *VirtualizationIweClusterAllOf) GetCapacityRunway() int64`

GetCapacityRunway returns the CapacityRunway field if non-nil, zero value otherwise.

### GetCapacityRunwayOk

`func (o *VirtualizationIweClusterAllOf) GetCapacityRunwayOk() (*int64, bool)`

GetCapacityRunwayOk returns a tuple with the CapacityRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRunway

`func (o *VirtualizationIweClusterAllOf) SetCapacityRunway(v int64)`

SetCapacityRunway sets CapacityRunway field to given value.

### HasCapacityRunway

`func (o *VirtualizationIweClusterAllOf) HasCapacityRunway() bool`

HasCapacityRunway returns a boolean if a field has been set.

### GetClusterName

`func (o *VirtualizationIweClusterAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VirtualizationIweClusterAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VirtualizationIweClusterAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VirtualizationIweClusterAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetComputeNodeCount

`func (o *VirtualizationIweClusterAllOf) GetComputeNodeCount() int64`

GetComputeNodeCount returns the ComputeNodeCount field if non-nil, zero value otherwise.

### GetComputeNodeCountOk

`func (o *VirtualizationIweClusterAllOf) GetComputeNodeCountOk() (*int64, bool)`

GetComputeNodeCountOk returns a tuple with the ComputeNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodeCount

`func (o *VirtualizationIweClusterAllOf) SetComputeNodeCount(v int64)`

SetComputeNodeCount sets ComputeNodeCount field to given value.

### HasComputeNodeCount

`func (o *VirtualizationIweClusterAllOf) HasComputeNodeCount() bool`

HasComputeNodeCount returns a boolean if a field has been set.

### GetConfiguredCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) GetConfiguredCpuOverSubFactor() float64`

GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field if non-nil, zero value otherwise.

### GetConfiguredCpuOverSubFactorOk

`func (o *VirtualizationIweClusterAllOf) GetConfiguredCpuOverSubFactorOk() (*float64, bool)`

GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) SetConfiguredCpuOverSubFactor(v float64)`

SetConfiguredCpuOverSubFactor sets ConfiguredCpuOverSubFactor field to given value.

### HasConfiguredCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) HasConfiguredCpuOverSubFactor() bool`

HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.

### GetConvergedNodeCount

`func (o *VirtualizationIweClusterAllOf) GetConvergedNodeCount() int64`

GetConvergedNodeCount returns the ConvergedNodeCount field if non-nil, zero value otherwise.

### GetConvergedNodeCountOk

`func (o *VirtualizationIweClusterAllOf) GetConvergedNodeCountOk() (*int64, bool)`

GetConvergedNodeCountOk returns a tuple with the ConvergedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergedNodeCount

`func (o *VirtualizationIweClusterAllOf) SetConvergedNodeCount(v int64)`

SetConvergedNodeCount sets ConvergedNodeCount field to given value.

### HasConvergedNodeCount

`func (o *VirtualizationIweClusterAllOf) HasConvergedNodeCount() bool`

HasConvergedNodeCount returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *VirtualizationIweClusterAllOf) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VirtualizationIweClusterAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VirtualizationIweClusterAllOf) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VirtualizationIweClusterAllOf) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *VirtualizationIweClusterAllOf) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *VirtualizationIweClusterAllOf) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetCurrentCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) GetCurrentCpuOverSubFactor() float64`

GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field if non-nil, zero value otherwise.

### GetCurrentCpuOverSubFactorOk

`func (o *VirtualizationIweClusterAllOf) GetCurrentCpuOverSubFactorOk() (*float64, bool)`

GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) SetCurrentCpuOverSubFactor(v float64)`

SetCurrentCpuOverSubFactor sets CurrentCpuOverSubFactor field to given value.

### HasCurrentCpuOverSubFactor

`func (o *VirtualizationIweClusterAllOf) HasCurrentCpuOverSubFactor() bool`

HasCurrentCpuOverSubFactor returns a boolean if a field has been set.

### GetDatacenterName

`func (o *VirtualizationIweClusterAllOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *VirtualizationIweClusterAllOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *VirtualizationIweClusterAllOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *VirtualizationIweClusterAllOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetDeploymentType

`func (o *VirtualizationIweClusterAllOf) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *VirtualizationIweClusterAllOf) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *VirtualizationIweClusterAllOf) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *VirtualizationIweClusterAllOf) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetDriveType

`func (o *VirtualizationIweClusterAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *VirtualizationIweClusterAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *VirtualizationIweClusterAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *VirtualizationIweClusterAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetFailureReason

`func (o *VirtualizationIweClusterAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VirtualizationIweClusterAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VirtualizationIweClusterAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VirtualizationIweClusterAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHxVersion

`func (o *VirtualizationIweClusterAllOf) GetHxVersion() string`

GetHxVersion returns the HxVersion field if non-nil, zero value otherwise.

### GetHxVersionOk

`func (o *VirtualizationIweClusterAllOf) GetHxVersionOk() (*string, bool)`

GetHxVersionOk returns a tuple with the HxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxVersion

`func (o *VirtualizationIweClusterAllOf) SetHxVersion(v string)`

SetHxVersion sets HxVersion field to given value.

### HasHxVersion

`func (o *VirtualizationIweClusterAllOf) HasHxVersion() bool`

HasHxVersion returns a boolean if a field has been set.

### GetHypervisorBuild

`func (o *VirtualizationIweClusterAllOf) GetHypervisorBuild() string`

GetHypervisorBuild returns the HypervisorBuild field if non-nil, zero value otherwise.

### GetHypervisorBuildOk

`func (o *VirtualizationIweClusterAllOf) GetHypervisorBuildOk() (*string, bool)`

GetHypervisorBuildOk returns a tuple with the HypervisorBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorBuild

`func (o *VirtualizationIweClusterAllOf) SetHypervisorBuild(v string)`

SetHypervisorBuild sets HypervisorBuild field to given value.

### HasHypervisorBuild

`func (o *VirtualizationIweClusterAllOf) HasHypervisorBuild() bool`

HasHypervisorBuild returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *VirtualizationIweClusterAllOf) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *VirtualizationIweClusterAllOf) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *VirtualizationIweClusterAllOf) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *VirtualizationIweClusterAllOf) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *VirtualizationIweClusterAllOf) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *VirtualizationIweClusterAllOf) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *VirtualizationIweClusterAllOf) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *VirtualizationIweClusterAllOf) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VirtualizationIweClusterAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VirtualizationIweClusterAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VirtualizationIweClusterAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VirtualizationIweClusterAllOf) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *VirtualizationIweClusterAllOf) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *VirtualizationIweClusterAllOf) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetStorageCapacity

`func (o *VirtualizationIweClusterAllOf) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *VirtualizationIweClusterAllOf) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *VirtualizationIweClusterAllOf) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *VirtualizationIweClusterAllOf) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### GetStorageNodeCount

`func (o *VirtualizationIweClusterAllOf) GetStorageNodeCount() int64`

GetStorageNodeCount returns the StorageNodeCount field if non-nil, zero value otherwise.

### GetStorageNodeCountOk

`func (o *VirtualizationIweClusterAllOf) GetStorageNodeCountOk() (*int64, bool)`

GetStorageNodeCountOk returns a tuple with the StorageNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNodeCount

`func (o *VirtualizationIweClusterAllOf) SetStorageNodeCount(v int64)`

SetStorageNodeCount sets StorageNodeCount field to given value.

### HasStorageNodeCount

`func (o *VirtualizationIweClusterAllOf) HasStorageNodeCount() bool`

HasStorageNodeCount returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *VirtualizationIweClusterAllOf) GetStorageUtilization() float32`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *VirtualizationIweClusterAllOf) GetStorageUtilizationOk() (*float32, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *VirtualizationIweClusterAllOf) SetStorageUtilization(v float32)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *VirtualizationIweClusterAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetUtilizationPercentage

`func (o *VirtualizationIweClusterAllOf) GetUtilizationPercentage() float32`

GetUtilizationPercentage returns the UtilizationPercentage field if non-nil, zero value otherwise.

### GetUtilizationPercentageOk

`func (o *VirtualizationIweClusterAllOf) GetUtilizationPercentageOk() (*float32, bool)`

GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercentage

`func (o *VirtualizationIweClusterAllOf) SetUtilizationPercentage(v float32)`

SetUtilizationPercentage sets UtilizationPercentage field to given value.

### HasUtilizationPercentage

`func (o *VirtualizationIweClusterAllOf) HasUtilizationPercentage() bool`

HasUtilizationPercentage returns a boolean if a field has been set.

### GetUtilizationTrendPercentage

`func (o *VirtualizationIweClusterAllOf) GetUtilizationTrendPercentage() float32`

GetUtilizationTrendPercentage returns the UtilizationTrendPercentage field if non-nil, zero value otherwise.

### GetUtilizationTrendPercentageOk

`func (o *VirtualizationIweClusterAllOf) GetUtilizationTrendPercentageOk() (*float32, bool)`

GetUtilizationTrendPercentageOk returns a tuple with the UtilizationTrendPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationTrendPercentage

`func (o *VirtualizationIweClusterAllOf) SetUtilizationTrendPercentage(v float32)`

SetUtilizationTrendPercentage sets UtilizationTrendPercentage field to given value.

### HasUtilizationTrendPercentage

`func (o *VirtualizationIweClusterAllOf) HasUtilizationTrendPercentage() bool`

HasUtilizationTrendPercentage returns a boolean if a field has been set.

### GetAssociatedProfile

`func (o *VirtualizationIweClusterAllOf) GetAssociatedProfile() PolicyAbstractProfileRelationship`

GetAssociatedProfile returns the AssociatedProfile field if non-nil, zero value otherwise.

### GetAssociatedProfileOk

`func (o *VirtualizationIweClusterAllOf) GetAssociatedProfileOk() (*PolicyAbstractProfileRelationship, bool)`

GetAssociatedProfileOk returns a tuple with the AssociatedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedProfile

`func (o *VirtualizationIweClusterAllOf) SetAssociatedProfile(v PolicyAbstractProfileRelationship)`

SetAssociatedProfile sets AssociatedProfile field to given value.

### HasAssociatedProfile

`func (o *VirtualizationIweClusterAllOf) HasAssociatedProfile() bool`

HasAssociatedProfile returns a boolean if a field has been set.

### GetHxCluster

`func (o *VirtualizationIweClusterAllOf) GetHxCluster() StorageBaseClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *VirtualizationIweClusterAllOf) GetHxClusterOk() (*StorageBaseClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *VirtualizationIweClusterAllOf) SetHxCluster(v StorageBaseClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *VirtualizationIweClusterAllOf) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationIweClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationIweClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationIweClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationIweClusterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


