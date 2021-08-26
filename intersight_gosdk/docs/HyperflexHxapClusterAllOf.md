# HyperflexHxapClusterAllOf

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

### NewHyperflexHxapClusterAllOf

`func NewHyperflexHxapClusterAllOf(classId string, objectType string, ) *HyperflexHxapClusterAllOf`

NewHyperflexHxapClusterAllOf instantiates a new HyperflexHxapClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapClusterAllOfWithDefaults

`func NewHyperflexHxapClusterAllOfWithDefaults() *HyperflexHxapClusterAllOf`

NewHyperflexHxapClusterAllOfWithDefaults instantiates a new HyperflexHxapClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapClusterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapClusterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapClusterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapClusterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapClusterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapClusterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfiguredCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) GetConfiguredCpuOverSubFactor() float64`

GetConfiguredCpuOverSubFactor returns the ConfiguredCpuOverSubFactor field if non-nil, zero value otherwise.

### GetConfiguredCpuOverSubFactorOk

`func (o *HyperflexHxapClusterAllOf) GetConfiguredCpuOverSubFactorOk() (*float64, bool)`

GetConfiguredCpuOverSubFactorOk returns a tuple with the ConfiguredCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) SetConfiguredCpuOverSubFactor(v float64)`

SetConfiguredCpuOverSubFactor sets ConfiguredCpuOverSubFactor field to given value.

### HasConfiguredCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) HasConfiguredCpuOverSubFactor() bool`

HasConfiguredCpuOverSubFactor returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *HyperflexHxapClusterAllOf) GetCpuAllocation() VirtualizationCpuAllocation`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *HyperflexHxapClusterAllOf) GetCpuAllocationOk() (*VirtualizationCpuAllocation, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *HyperflexHxapClusterAllOf) SetCpuAllocation(v VirtualizationCpuAllocation)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *HyperflexHxapClusterAllOf) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### SetCpuAllocationNil

`func (o *HyperflexHxapClusterAllOf) SetCpuAllocationNil(b bool)`

 SetCpuAllocationNil sets the value for CpuAllocation to be an explicit nil

### UnsetCpuAllocation
`func (o *HyperflexHxapClusterAllOf) UnsetCpuAllocation()`

UnsetCpuAllocation ensures that no value is present for CpuAllocation, not even an explicit nil
### GetCurrentCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) GetCurrentCpuOverSubFactor() float64`

GetCurrentCpuOverSubFactor returns the CurrentCpuOverSubFactor field if non-nil, zero value otherwise.

### GetCurrentCpuOverSubFactorOk

`func (o *HyperflexHxapClusterAllOf) GetCurrentCpuOverSubFactorOk() (*float64, bool)`

GetCurrentCpuOverSubFactorOk returns a tuple with the CurrentCpuOverSubFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) SetCurrentCpuOverSubFactor(v float64)`

SetCurrentCpuOverSubFactor sets CurrentCpuOverSubFactor field to given value.

### HasCurrentCpuOverSubFactor

`func (o *HyperflexHxapClusterAllOf) HasCurrentCpuOverSubFactor() bool`

HasCurrentCpuOverSubFactor returns a boolean if a field has been set.

### GetDatacenterName

`func (o *HyperflexHxapClusterAllOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *HyperflexHxapClusterAllOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *HyperflexHxapClusterAllOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *HyperflexHxapClusterAllOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetFailureReason

`func (o *HyperflexHxapClusterAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapClusterAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapClusterAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapClusterAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHypervisorBuild

`func (o *HyperflexHxapClusterAllOf) GetHypervisorBuild() string`

GetHypervisorBuild returns the HypervisorBuild field if non-nil, zero value otherwise.

### GetHypervisorBuildOk

`func (o *HyperflexHxapClusterAllOf) GetHypervisorBuildOk() (*string, bool)`

GetHypervisorBuildOk returns a tuple with the HypervisorBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorBuild

`func (o *HyperflexHxapClusterAllOf) SetHypervisorBuild(v string)`

SetHypervisorBuild sets HypervisorBuild field to given value.

### HasHypervisorBuild

`func (o *HyperflexHxapClusterAllOf) HasHypervisorBuild() bool`

HasHypervisorBuild returns a boolean if a field has been set.

### GetManagementIpAddress

`func (o *HyperflexHxapClusterAllOf) GetManagementIpAddress() string`

GetManagementIpAddress returns the ManagementIpAddress field if non-nil, zero value otherwise.

### GetManagementIpAddressOk

`func (o *HyperflexHxapClusterAllOf) GetManagementIpAddressOk() (*string, bool)`

GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIpAddress

`func (o *HyperflexHxapClusterAllOf) SetManagementIpAddress(v string)`

SetManagementIpAddress sets ManagementIpAddress field to given value.

### HasManagementIpAddress

`func (o *HyperflexHxapClusterAllOf) HasManagementIpAddress() bool`

HasManagementIpAddress returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *HyperflexHxapClusterAllOf) GetMemoryAllocation() VirtualizationMemoryAllocation`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *HyperflexHxapClusterAllOf) GetMemoryAllocationOk() (*VirtualizationMemoryAllocation, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *HyperflexHxapClusterAllOf) SetMemoryAllocation(v VirtualizationMemoryAllocation)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *HyperflexHxapClusterAllOf) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.

### SetMemoryAllocationNil

`func (o *HyperflexHxapClusterAllOf) SetMemoryAllocationNil(b bool)`

 SetMemoryAllocationNil sets the value for MemoryAllocation to be an explicit nil

### UnsetMemoryAllocation
`func (o *HyperflexHxapClusterAllOf) UnsetMemoryAllocation()`

UnsetMemoryAllocation ensures that no value is present for MemoryAllocation, not even an explicit nil
### GetHxCluster

`func (o *HyperflexHxapClusterAllOf) GetHxCluster() HyperflexClusterRelationship`

GetHxCluster returns the HxCluster field if non-nil, zero value otherwise.

### GetHxClusterOk

`func (o *HyperflexHxapClusterAllOf) GetHxClusterOk() (*HyperflexClusterRelationship, bool)`

GetHxClusterOk returns a tuple with the HxCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxCluster

`func (o *HyperflexHxapClusterAllOf) SetHxCluster(v HyperflexClusterRelationship)`

SetHxCluster sets HxCluster field to given value.

### HasHxCluster

`func (o *HyperflexHxapClusterAllOf) HasHxCluster() bool`

HasHxCluster returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHxapClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHxapClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHxapClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHxapClusterAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


