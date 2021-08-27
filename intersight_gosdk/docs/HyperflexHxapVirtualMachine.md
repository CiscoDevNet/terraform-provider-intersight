# HyperflexHxapVirtualMachine

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

### NewHyperflexHxapVirtualMachine

`func NewHyperflexHxapVirtualMachine(classId string, objectType string, ) *HyperflexHxapVirtualMachine`

NewHyperflexHxapVirtualMachine instantiates a new HyperflexHxapVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapVirtualMachineWithDefaults

`func NewHyperflexHxapVirtualMachineWithDefaults() *HyperflexHxapVirtualMachine`

NewHyperflexHxapVirtualMachineWithDefaults instantiates a new HyperflexHxapVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachine) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### SetAffinitySelectorsNil

`func (o *HyperflexHxapVirtualMachine) SetAffinitySelectorsNil(b bool)`

 SetAffinitySelectorsNil sets the value for AffinitySelectors to be an explicit nil

### UnsetAffinitySelectors
`func (o *HyperflexHxapVirtualMachine) UnsetAffinitySelectors()`

UnsetAffinitySelectors ensures that no value is present for AffinitySelectors, not even an explicit nil
### GetAge

`func (o *HyperflexHxapVirtualMachine) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapVirtualMachine) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapVirtualMachine) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapVirtualMachine) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### SetAntiAffinitySelectorsNil

`func (o *HyperflexHxapVirtualMachine) SetAntiAffinitySelectorsNil(b bool)`

 SetAntiAffinitySelectorsNil sets the value for AntiAffinitySelectors to be an explicit nil

### UnsetAntiAffinitySelectors
`func (o *HyperflexHxapVirtualMachine) UnsetAntiAffinitySelectors()`

UnsetAntiAffinitySelectors ensures that no value is present for AntiAffinitySelectors, not even an explicit nil
### GetDisks

`func (o *HyperflexHxapVirtualMachine) GetDisks() []HyperflexVmDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HyperflexHxapVirtualMachine) GetDisksOk() (*[]HyperflexVmDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HyperflexHxapVirtualMachine) SetDisks(v []HyperflexVmDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HyperflexHxapVirtualMachine) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *HyperflexHxapVirtualMachine) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *HyperflexHxapVirtualMachine) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetFailureReason

`func (o *HyperflexHxapVirtualMachine) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapVirtualMachine) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapVirtualMachine) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapVirtualMachine) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetInterfaces

`func (o *HyperflexHxapVirtualMachine) GetInterfaces() []HyperflexVmInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *HyperflexHxapVirtualMachine) GetInterfacesOk() (*[]HyperflexVmInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *HyperflexHxapVirtualMachine) SetInterfaces(v []HyperflexVmInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *HyperflexHxapVirtualMachine) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *HyperflexHxapVirtualMachine) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *HyperflexHxapVirtualMachine) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetLabels

`func (o *HyperflexHxapVirtualMachine) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HyperflexHxapVirtualMachine) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HyperflexHxapVirtualMachine) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HyperflexHxapVirtualMachine) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *HyperflexHxapVirtualMachine) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *HyperflexHxapVirtualMachine) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetNetworkCount

`func (o *HyperflexHxapVirtualMachine) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *HyperflexHxapVirtualMachine) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *HyperflexHxapVirtualMachine) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *HyperflexHxapVirtualMachine) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperflexHxapVirtualMachine) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexHxapVirtualMachine) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexHxapVirtualMachine) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexHxapVirtualMachine) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxapVirtualMachine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxapVirtualMachine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxapVirtualMachine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxapVirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapVirtualMachine) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapVirtualMachine) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapVirtualMachine) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapVirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapVirtualMachine) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapVirtualMachine) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapVirtualMachine) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapVirtualMachine) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


