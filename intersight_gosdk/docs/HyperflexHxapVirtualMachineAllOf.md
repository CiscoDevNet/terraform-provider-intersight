# HyperflexHxapVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapVirtualMachine"]
**AffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**Age** | Pointer to **string** | Denotes age or life time of the VM in nano seconds. | [optional] 
**AntiAffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**Disks** | Pointer to [**[]HyperflexVmDisk**](HyperflexVmDisk.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when VM is in failed state. | [optional] 
**Interfaces** | Pointer to [**[]HyperflexVmInterface**](HyperflexVmInterface.md) |  | [optional] 
**Labels** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**NetworkCount** | Pointer to **int64** | Number network interfaces associated with the virtual machine. | [optional] 
**StartTime** | Pointer to **string** | Denotes the VM start timestamp. | [optional] 
**Status** | Pointer to **string** | Status of virtual machine. * &#x60;Unknown&#x60; - Virtual machine state is not available. * &#x60;Running&#x60; - Virtual machine is running normally. * &#x60;Stopped&#x60; - Virtual machine has been stopped. * &#x60;WaitForLaunch&#x60; - Virtual machine is wating to be launched. * &#x60;Paused&#x60; - Virtual machine is currently paused. * &#x60;ImportInProgress&#x60; - Virtual machine image is being imported into the platform. * &#x60;ImportFailed&#x60; - Virtual machine image import operation failed. * &#x60;DiskCloneInProgress&#x60; - Disk clone operation for the virtual machine is in progress. * &#x60;DiskCloneFailed&#x60; - Disk clone operation for the virtual machine failed. * &#x60;Processing&#x60; - Virtual machine is being created. * &#x60;UnSchedulable&#x60; - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications. * &#x60;Failed&#x60; - Some virtual machine operation has failed. More information is available as part of the results of the operation. * &#x60;&#x60; - Virtual machine status is not available. | [optional] [default to "Unknown"]
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**Host** | Pointer to [**HyperflexHxapHostRelationship**](HyperflexHxapHostRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapVirtualMachineAllOf

`func NewHyperflexHxapVirtualMachineAllOf(classId string, objectType string, ) *HyperflexHxapVirtualMachineAllOf`

NewHyperflexHxapVirtualMachineAllOf instantiates a new HyperflexHxapVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapVirtualMachineAllOfWithDefaults

`func NewHyperflexHxapVirtualMachineAllOfWithDefaults() *HyperflexHxapVirtualMachineAllOf`

NewHyperflexHxapVirtualMachineAllOfWithDefaults instantiates a new HyperflexHxapVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### SetAffinitySelectorsNil

`func (o *HyperflexHxapVirtualMachineAllOf) SetAffinitySelectorsNil(b bool)`

 SetAffinitySelectorsNil sets the value for AffinitySelectors to be an explicit nil

### UnsetAffinitySelectors
`func (o *HyperflexHxapVirtualMachineAllOf) UnsetAffinitySelectors()`

UnsetAffinitySelectors ensures that no value is present for AffinitySelectors, not even an explicit nil
### GetAge

`func (o *HyperflexHxapVirtualMachineAllOf) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapVirtualMachineAllOf) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapVirtualMachineAllOf) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachineAllOf) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### SetAntiAffinitySelectorsNil

`func (o *HyperflexHxapVirtualMachineAllOf) SetAntiAffinitySelectorsNil(b bool)`

 SetAntiAffinitySelectorsNil sets the value for AntiAffinitySelectors to be an explicit nil

### UnsetAntiAffinitySelectors
`func (o *HyperflexHxapVirtualMachineAllOf) UnsetAntiAffinitySelectors()`

UnsetAntiAffinitySelectors ensures that no value is present for AntiAffinitySelectors, not even an explicit nil
### GetDisks

`func (o *HyperflexHxapVirtualMachineAllOf) GetDisks() []HyperflexVmDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetDisksOk() (*[]HyperflexVmDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HyperflexHxapVirtualMachineAllOf) SetDisks(v []HyperflexVmDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HyperflexHxapVirtualMachineAllOf) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *HyperflexHxapVirtualMachineAllOf) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *HyperflexHxapVirtualMachineAllOf) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetFailureReason

`func (o *HyperflexHxapVirtualMachineAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapVirtualMachineAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapVirtualMachineAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetInterfaces

`func (o *HyperflexHxapVirtualMachineAllOf) GetInterfaces() []HyperflexVmInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetInterfacesOk() (*[]HyperflexVmInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *HyperflexHxapVirtualMachineAllOf) SetInterfaces(v []HyperflexVmInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *HyperflexHxapVirtualMachineAllOf) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *HyperflexHxapVirtualMachineAllOf) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *HyperflexHxapVirtualMachineAllOf) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetLabels

`func (o *HyperflexHxapVirtualMachineAllOf) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HyperflexHxapVirtualMachineAllOf) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HyperflexHxapVirtualMachineAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *HyperflexHxapVirtualMachineAllOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *HyperflexHxapVirtualMachineAllOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetNetworkCount

`func (o *HyperflexHxapVirtualMachineAllOf) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *HyperflexHxapVirtualMachineAllOf) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *HyperflexHxapVirtualMachineAllOf) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperflexHxapVirtualMachineAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexHxapVirtualMachineAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexHxapVirtualMachineAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxapVirtualMachineAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxapVirtualMachineAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxapVirtualMachineAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapVirtualMachineAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapVirtualMachineAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapVirtualMachineAllOf) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapVirtualMachineAllOf) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapVirtualMachineAllOf) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


