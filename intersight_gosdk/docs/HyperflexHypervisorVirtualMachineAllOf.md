# HyperflexHypervisorVirtualMachineAllOf

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

### NewHyperflexHypervisorVirtualMachineAllOf

`func NewHyperflexHypervisorVirtualMachineAllOf(classId string, objectType string, ) *HyperflexHypervisorVirtualMachineAllOf`

NewHyperflexHypervisorVirtualMachineAllOf instantiates a new HyperflexHypervisorVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHypervisorVirtualMachineAllOfWithDefaults

`func NewHyperflexHypervisorVirtualMachineAllOfWithDefaults() *HyperflexHypervisorVirtualMachineAllOf`

NewHyperflexHypervisorVirtualMachineAllOfWithDefaults instantiates a new HyperflexHypervisorVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionState

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetGuestOsState

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetGuestOsState() string`

GetGuestOsState returns the GuestOsState field if non-nil, zero value otherwise.

### GetGuestOsStateOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetGuestOsStateOk() (*string, bool)`

GetGuestOsStateOk returns a tuple with the GuestOsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsState

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetGuestOsState(v string)`

SetGuestOsState sets GuestOsState field to given value.

### HasGuestOsState

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasGuestOsState() bool`

HasGuestOsState returns a boolean if a field has been set.

### GetHostUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetIp

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetIp() NetworkHyperFlexNetworkAddress`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetIp(v NetworkHyperFlexNetworkAddress)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HyperflexHypervisorVirtualMachineAllOf) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPath

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetPlatformInstanceId() string`

GetPlatformInstanceId returns the PlatformInstanceId field if non-nil, zero value otherwise.

### GetPlatformInstanceIdOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetPlatformInstanceIdOk() (*string, bool)`

GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetPlatformInstanceId(v string)`

SetPlatformInstanceId sets PlatformInstanceId field to given value.

### HasPlatformInstanceId

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasPlatformInstanceId() bool`

HasPlatformInstanceId returns a boolean if a field has been set.

### GetStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageProvisionedInBytes() int64`

GetStorageProvisionedInBytes returns the StorageProvisionedInBytes field if non-nil, zero value otherwise.

### GetStorageProvisionedInBytesOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageProvisionedInBytesOk() (*int64, bool)`

GetStorageProvisionedInBytesOk returns a tuple with the StorageProvisionedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetStorageProvisionedInBytes(v int64)`

SetStorageProvisionedInBytes sets StorageProvisionedInBytes field to given value.

### HasStorageProvisionedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasStorageProvisionedInBytes() bool`

HasStorageProvisionedInBytes returns a boolean if a field has been set.

### GetStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageUsedInBytes() int64`

GetStorageUsedInBytes returns the StorageUsedInBytes field if non-nil, zero value otherwise.

### GetStorageUsedInBytesOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetStorageUsedInBytesOk() (*int64, bool)`

GetStorageUsedInBytesOk returns a tuple with the StorageUsedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetStorageUsedInBytes(v int64)`

SetStorageUsedInBytes sets StorageUsedInBytes field to given value.

### HasStorageUsedInBytes

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasStorageUsedInBytes() bool`

HasStorageUsedInBytes returns a boolean if a field has been set.

### GetTemplate

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetVmInstanceUuid() string`

GetVmInstanceUuid returns the VmInstanceUuid field if non-nil, zero value otherwise.

### GetVmInstanceUuidOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetVmInstanceUuidOk() (*string, bool)`

GetVmInstanceUuidOk returns a tuple with the VmInstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetVmInstanceUuid(v string)`

SetVmInstanceUuid sets VmInstanceUuid field to given value.

### HasVmInstanceUuid

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasVmInstanceUuid() bool`

HasVmInstanceUuid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetHost() HyperflexHypervisorHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHypervisorVirtualMachineAllOf) GetHostOk() (*HyperflexHypervisorHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHypervisorVirtualMachineAllOf) SetHost(v HyperflexHypervisorHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHypervisorVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


