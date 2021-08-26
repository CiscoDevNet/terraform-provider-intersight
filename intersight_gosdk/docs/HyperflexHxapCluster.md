# HyperflexHxapCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapCluster"]
**ConfiguredCpuOverSubFactor** | Pointer to **float64** | CPU oversubscription factor configured on the cluster. | [optional] 
**CpuAllocation** | Pointer to [**NullableVirtualizationCpuAllocation**](VirtualizationCpuAllocation.md) |  | [optional] 
**CurrentCpuOverSubFactor** | Pointer to **float64** | Current oversubscription factor of the cluster. | [optional] 
**DatacenterName** | Pointer to **string** | Datacenter to which the cluster belongs. | [optional] 
**FailureReason** | Pointer to **string** | Reason for the failure when cluster is in failed state. | [optional] 
**HypervisorBuild** | Pointer to **string** | Hypervisor version of HyperFlex compute cluster along with build number. | [optional] 
**ManagementIpAddress** | Pointer to **string** | Management IP Address of the cluster. | [optional] 
**MemoryAllocation** | Pointer to [**NullableVirtualizationMemoryAllocation**](VirtualizationMemoryAllocation.md) |  | [optional] 
**HxCluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapCluster

`func NewHyperflexHxapCluster(classId string, objectType string, ) *HyperflexHxapCluster`

NewHyperflexHxapCluster instantiates a new HyperflexHxapCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapClusterWithDefaults

`func NewHyperflexHxapClusterWithDefaults() *HyperflexHxapCluster`

NewHyperflexHxapClusterWithDefaults instantiates a new HyperflexHxapCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfiguredCpuOverSubFactor

`func (o *HyperflexHxapCluster) GetConfiguredCpuOverSubFactor() float64`

GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field if non-nil, zero value otherwise.

### GetConfiguredCpuOverSubFactorOk

`func (o *HyperflexHxapCluster) GetConfiguredCpuOverSubFactorOk() (*float64, bool)`

GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredCpuOverSubFactor

`func (o *HyperflexHxapCluster) SetConfiguredCpuOverSubFactor(v float64)`

SetConfiguredCpuOverSubFactor sets ConfiguredCpuOverSubFactor field to given value.

### HasConfiguredCpuOverSubFactor

`func (o *HyperflexHxapCluster) HasConfiguredCpuOverSubFactor() bool`

HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *HyperflexHxapCluster) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *HyperflexHxapCluster) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *HyperflexHxapCluster) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *HyperflexHxapCluster) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *HyperflexHxapCluster) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *HyperflexHxapCluster) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetCurrentCpuOverSubFactor

`func (o *HyperflexHxapCluster) GetCurrentCpuOverSubFactor() float64`

GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field if non-nil, zero value otherwise.

### GetCurrentCpuOverSubFactorOk

`func (o *HyperflexHxapCluster) GetCurrentCpuOverSubFactorOk() (*float64, bool)`

GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuOverSubFactor

`func (o *HyperflexHxapCluster) SetCurrentCpuOverSubFactor(v float64)`

SetCurrentCpuOverSubFactor sets CurrentCpuOverSubFactor field to given value.

### HasCurrentCpuOverSubFactor

`func (o *HyperflexHxapCluster) HasCurrentCpuOverSubFactor() bool`

HasCurrentCpuOverSubFactor returns a boolean if a field has been set.

### GetDatacenterName

`func (o *HyperflexHxapCluster) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *HyperflexHxapCluster) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *HyperflexHxapCluster) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *HyperflexHxapCluster) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetFailureReason

`func (o *HyperflexHxapCluster) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapCluster) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapCluster) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapCluster) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHypervisorBuild

`func (o *HyperflexHxapCluster) GetHypervisorBuild() string`

GetHypervisorBuild returns the HypervisorBuild field if non-nil, zero value otherwise.

### GetHypervisorBuildOk

`func (o *HyperflexHxapCluster) GetHypervisorBuildOk() (*string, bool)`

GetHypervisorBuildOk returns a tuple with the HypervisorBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorBuild

`func (o *HyperflexHxapCluster) SetHypervisorBuild(v string)`

SetHypervisorBuild sets HypervisorBuild field to given value.

### HasHypervisorBuild

`func (o *HyperflexHxapCluster) HasHypervisorBuild() bool`

HasHypervisorBuild returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *HyperflexHxapCluster) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *HyperflexHxapCluster) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *HyperflexHxapCluster) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *HyperflexHxapCluster) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *HyperflexHxapCluster) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *HyperflexHxapCluster) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *HyperflexHxapCluster) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *HyperflexHxapCluster) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *HyperflexHxapCluster) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *HyperflexHxapCluster) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetHxCluster

`func (o *HyperflexHxapCluster) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexHxapCluster) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexHxapCluster) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexHxapCluster) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHxapCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHxapCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHxapCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHxapCluster) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


