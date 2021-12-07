# VirtualizationIweVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweVirtualMachine"]
**AffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**Age** | Pointer to **string** | Denotes age or life time of the VM in nano seconds. | [optional] 
**AntiAffinitySelectors** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**Disks** | Pointer to [**[]VirtualizationVmDisk**](VirtualizationVmDisk.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when VM is in failed state. | [optional] 
**Interfaces** | Pointer to [**[]VirtualizationVmInterface**](VirtualizationVmInterface.md) |  | [optional] 
**Labels** | Pointer to [**[]InfraMetaData**](InfraMetaData.md) |  | [optional] 
**NetworkCount** | Pointer to **int64** | Number network interfaces associated with the virtual machine. | [optional] 
**StartTime** | Pointer to **string** | Denotes the VM start timestamp. | [optional] 
**Status** | Pointer to **string** | Status of virtual machine. * &#x60;Unknown&#x60; - Virtual machine state is not available. * &#x60;Running&#x60; - Virtual machine is running normally. * &#x60;Stopped&#x60; - Virtual machine has been stopped. * &#x60;WaitForLaunch&#x60; - Virtual machine is wating to be launched. * &#x60;Paused&#x60; - Virtual machine is currently paused. * &#x60;ImportInProgress&#x60; - Virtual machine image is being imported into the platform. * &#x60;ImportFailed&#x60; - Virtual machine image import operation failed. * &#x60;DiskCloneInProgress&#x60; - Disk clone operation for the virtual machine is in progress. * &#x60;DiskCloneFailed&#x60; - Disk clone operation for the virtual machine failed. * &#x60;DiskCreateInProgress&#x60; - Disk create operation for the virtual machine is in progress. * &#x60;DiskCreateFailed&#x60; - Disk create operation for the virtual machine failed. * &#x60;Processing&#x60; - Virtual machine is being created. * &#x60;UnSchedulable&#x60; - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications. * &#x60;Failed&#x60; - Some virtual machine operation has failed. More information is available as part of the results of the operation. * &#x60;&#x60; - Virtual machine status is not available. | [optional] [default to "Unknown"]
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationIweHostRelationship**](VirtualizationIweHostRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweVirtualMachineAllOf

`func NewVirtualizationIweVirtualMachineAllOf(classId string, objectType string, ) *VirtualizationIweVirtualMachineAllOf`

NewVirtualizationIweVirtualMachineAllOf instantiates a new VirtualizationIweVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweVirtualMachineAllOfWithDefaults

`func NewVirtualizationIweVirtualMachineAllOfWithDefaults() *VirtualizationIweVirtualMachineAllOf`

NewVirtualizationIweVirtualMachineAllOfWithDefaults instantiates a new VirtualizationIweVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### SetAffinitySelectorsNil

`func (o *VirtualizationIweVirtualMachineAllOf) SetAffinitySelectorsNil(b bool)`

 SetAffinitySelectorsNil sets the value for AffinitySelectors to be an explicit nil

### UnsetAffinitySelectors
`func (o *VirtualizationIweVirtualMachineAllOf) UnsetAffinitySelectors()`

UnsetAffinitySelectors ensures that no value is present for AffinitySelectors, not even an explicit nil
### GetAge

`func (o *VirtualizationIweVirtualMachineAllOf) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VirtualizationIweVirtualMachineAllOf) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *VirtualizationIweVirtualMachineAllOf) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAntiAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *VirtualizationIweVirtualMachineAllOf) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### SetAntiAffinitySelectorsNil

`func (o *VirtualizationIweVirtualMachineAllOf) SetAntiAffinitySelectorsNil(b bool)`

 SetAntiAffinitySelectorsNil sets the value for AntiAffinitySelectors to be an explicit nil

### UnsetAntiAffinitySelectors
`func (o *VirtualizationIweVirtualMachineAllOf) UnsetAntiAffinitySelectors()`

UnsetAntiAffinitySelectors ensures that no value is present for AntiAffinitySelectors, not even an explicit nil
### GetDisks

`func (o *VirtualizationIweVirtualMachineAllOf) GetDisks() []VirtualizationVmDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetDisksOk() (*[]VirtualizationVmDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VirtualizationIweVirtualMachineAllOf) SetDisks(v []VirtualizationVmDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VirtualizationIweVirtualMachineAllOf) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *VirtualizationIweVirtualMachineAllOf) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *VirtualizationIweVirtualMachineAllOf) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetFailureReason

`func (o *VirtualizationIweVirtualMachineAllOf) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VirtualizationIweVirtualMachineAllOf) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VirtualizationIweVirtualMachineAllOf) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetInterfaces

`func (o *VirtualizationIweVirtualMachineAllOf) GetInterfaces() []VirtualizationVmInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetInterfacesOk() (*[]VirtualizationVmInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VirtualizationIweVirtualMachineAllOf) SetInterfaces(v []VirtualizationVmInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *VirtualizationIweVirtualMachineAllOf) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *VirtualizationIweVirtualMachineAllOf) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *VirtualizationIweVirtualMachineAllOf) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetLabels

`func (o *VirtualizationIweVirtualMachineAllOf) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualizationIweVirtualMachineAllOf) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualizationIweVirtualMachineAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *VirtualizationIweVirtualMachineAllOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *VirtualizationIweVirtualMachineAllOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetNetworkCount

`func (o *VirtualizationIweVirtualMachineAllOf) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *VirtualizationIweVirtualMachineAllOf) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *VirtualizationIweVirtualMachineAllOf) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetStartTime

`func (o *VirtualizationIweVirtualMachineAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VirtualizationIweVirtualMachineAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VirtualizationIweVirtualMachineAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationIweVirtualMachineAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationIweVirtualMachineAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationIweVirtualMachineAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationIweVirtualMachineAllOf) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweVirtualMachineAllOf) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationIweVirtualMachineAllOf) GetHost() VirtualizationIweHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationIweVirtualMachineAllOf) GetHostOk() (*VirtualizationIweHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationIweVirtualMachineAllOf) SetHost(v VirtualizationIweHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationIweVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


