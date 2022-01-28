# HyperflexHypervisorVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HypervisorVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HypervisorVirtualMachine"]
**ConnectionState** | Pointer to **string** | The connectivity state of the HyperFlex virtual machine. | [optional] [readonly] 
**GuestOsState** | Pointer to **string** | Guest operating system state of the HyperFlex virtual machine. | [optional] [readonly] 
**HostUuid** | Pointer to **string** | Host UUID of the HyperFlex virtual machine. | [optional] [readonly] 
**Ip** | Pointer to [**NullableNetworkHyperFlexNetworkAddress**](NetworkHyperFlexNetworkAddress.md) |  | [optional] 
**Path** | Pointer to **string** | Directory path where virtual machine is stored. | [optional] [readonly] 
**PlatformInstanceId** | Pointer to **string** | The instance id of platform which a virtual machine is running on. | [optional] [readonly] 
**StorageProvisionedInBytes** | Pointer to **int64** | Total provisioned storage to the HyperFlex virtual machine in bytes. | [optional] [readonly] 
**StorageUsedInBytes** | Pointer to **int64** | Storage used by HyperFlex virtual machine in bytes. | [optional] [readonly] 
**Template** | Pointer to **bool** | Flag indicating whether or not this virtual machine is a template. Apply to the ESXi platform only. | [optional] [readonly] 
**VmInstanceUuid** | Pointer to **string** | The instance UUID of a virtual machine. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Host** | Pointer to [**HyperflexHypervisorHostRelationship**](HyperflexHypervisorHostRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHypervisorVirtualMachine

`func NewHyperflexHypervisorVirtualMachine(classId string, objectType string, ) *HyperflexHypervisorVirtualMachine`

NewHyperflexHypervisorVirtualMachine instantiates a new HyperflexHypervisorVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHypervisorVirtualMachineWithDefaults

`func NewHyperflexHypervisorVirtualMachineWithDefaults() *HyperflexHypervisorVirtualMachine`

NewHyperflexHypervisorVirtualMachineWithDefaults instantiates a new HyperflexHypervisorVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHypervisorVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHypervisorVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHypervisorVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHypervisorVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHypervisorVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHypervisorVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionState

`func (o *HyperflexHypervisorVirtualMachine) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *HyperflexHypervisorVirtualMachine) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *HyperflexHypervisorVirtualMachine) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *HyperflexHypervisorVirtualMachine) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetGuestOsState

`func (o *HyperflexHypervisorVirtualMachine) GetGuestOsState() string`

GetGuestOsState returns the GuestOsState field if non-nil, zero value otherwise.

### GetGuestOsStateOk

`func (o *HyperflexHypervisorVirtualMachine) GetGuestOsStateOk() (*string, bool)`

GetGuestOsStateOk returns a tuple with the GuestOsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsState

`func (o *HyperflexHypervisorVirtualMachine) SetGuestOsState(v string)`

SetGuestOsState sets GuestOsState field to given value.

### HasGuestOsState

`func (o *HyperflexHypervisorVirtualMachine) HasGuestOsState() bool`

HasGuestOsState returns a boolean if a field has been set.

### GetHostUuid

`func (o *HyperflexHypervisorVirtualMachine) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *HyperflexHypervisorVirtualMachine) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *HyperflexHypervisorVirtualMachine) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *HyperflexHypervisorVirtualMachine) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetIp

`func (o *HyperflexHypervisorVirtualMachine) GetIp() NetworkHyperFlexNetworkAddress`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexHypervisorVirtualMachine) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexHypervisorVirtualMachine) SetIp(v NetworkHyperFlexNetworkAddress)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexHypervisorVirtualMachine) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HyperflexHypervisorVirtualMachine) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HyperflexHypervisorVirtualMachine) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPath

`func (o *HyperflexHypervisorVirtualMachine) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HyperflexHypervisorVirtualMachine) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HyperflexHypervisorVirtualMachine) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HyperflexHypervisorVirtualMachine) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachine) GetPlatformInstanceId() string`

GetPlatformInstanceId returns the PlatformInstanceId field if non-nil, zero value otherwise.

### GetPlatformInstanceIdOk

`func (o *HyperflexHypervisorVirtualMachine) GetPlatformInstanceIdOk() (*string, bool)`

GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachine) SetPlatformInstanceId(v string)`

SetPlatformInstanceId sets PlatformInstanceId field to given value.

### HasPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachine) HasPlatformInstanceId() bool`

HasPlatformInstanceId returns a boolean if a field has been set.

### GetStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachine) GetStorageProvisionedInBytes() int64`

GetStorageProvisionedInBytes returns the StorageProvisionedInBytes field if non-nil, zero value otherwise.

### GetStorageProvisionedInBytesOk

`func (o *HyperflexHypervisorVirtualMachine) GetStorageProvisionedInBytesOk() (*int64, bool)`

GetStorageProvisionedInBytesOk returns a tuple with the StorageProvisionedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachine) SetStorageProvisionedInBytes(v int64)`

SetStorageProvisionedInBytes sets StorageProvisionedInBytes field to given value.

### HasStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachine) HasStorageProvisionedInBytes() bool`

HasStorageProvisionedInBytes returns a boolean if a field has been set.

### GetStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachine) GetStorageUsedInBytes() int64`

GetStorageUsedInBytes returns the StorageUsedInBytes field if non-nil, zero value otherwise.

### GetStorageUsedInBytesOk

`func (o *HyperflexHypervisorVirtualMachine) GetStorageUsedInBytesOk() (*int64, bool)`

GetStorageUsedInBytesOk returns a tuple with the StorageUsedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachine) SetStorageUsedInBytes(v int64)`

SetStorageUsedInBytes sets StorageUsedInBytes field to given value.

### HasStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachine) HasStorageUsedInBytes() bool`

HasStorageUsedInBytes returns a boolean if a field has been set.

### GetTemplate

`func (o *HyperflexHypervisorVirtualMachine) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *HyperflexHypervisorVirtualMachine) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *HyperflexHypervisorVirtualMachine) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *HyperflexHypervisorVirtualMachine) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachine) GetVmInstanceUuid() string`

GetVmInstanceUuid returns the VmInstanceUuid field if non-nil, zero value otherwise.

### GetVmInstanceUuidOk

`func (o *HyperflexHypervisorVirtualMachine) GetVmInstanceUuidOk() (*string, bool)`

GetVmInstanceUuidOk returns a tuple with the VmInstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachine) SetVmInstanceUuid(v string)`

SetVmInstanceUuid sets VmInstanceUuid field to given value.

### HasVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachine) HasVmInstanceUuid() bool`

HasVmInstanceUuid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHypervisorVirtualMachine) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHypervisorVirtualMachine) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHypervisorVirtualMachine) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHypervisorVirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHypervisorVirtualMachine) GetHost() HyperflexHypervisorHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHypervisorVirtualMachine) GetHostOk() (*HyperflexHypervisorHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHypervisorVirtualMachine) SetHost(v HyperflexHypervisorHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHypervisorVirtualMachine) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


