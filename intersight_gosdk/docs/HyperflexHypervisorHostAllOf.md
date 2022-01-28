# HyperflexHypervisorHostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HypervisorHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HypervisorHost"]
**ConfiguredMemory** | Pointer to **int64** | Memory configured for controller virtual machine. | [optional] [readonly] 
**ControllerVmDisplayVersion** | Pointer to **string** | The display version of HyperFlex software running on the controller VM. | [optional] [readonly] 
**ControllerVmHxdpDisplayVersion** | Pointer to **string** | Platform storage software product display version running on controller VM. | [optional] [readonly] 
**ControllerVmHxdpVersion** | Pointer to **string** | Platform storage software product version running on controller VM. | [optional] [readonly] 
**ControllerVmUuid** | Pointer to **string** | The UUID of the controller VM which belongs to this host. | [optional] [readonly] 
**ControllerVmVersion** | Pointer to **string** | The version of HyperFlex software running on the controller VM. | [optional] [readonly] 
**DataIp** | Pointer to [**NullableNetworkHyperFlexNetworkAddress**](NetworkHyperFlexNetworkAddress.md) |  | [optional] 
**HostStatus** | Pointer to **string** | The status of the HyperFlex host. * &#x60;UNKNOWN&#x60; - Current status of the HyperFlex host is unknown. * &#x60;ONLINE&#x60; - The HyperFlex host is online. * &#x60;OFFLINE&#x60; - The HyperFlex host is offline. * &#x60;INMAINTENANCE&#x60; - The HyperFlex host is in maintenance mode. * &#x60;DEGRADED&#x60; - Current status of the HyperFlex virtual machine is degraded. | [optional] [readonly] [default to "UNKNOWN"]
**Hypervisor** | Pointer to **string** | The hypervisor type of the host. | [optional] [readonly] 
**Ip** | Pointer to [**NullableNetworkHyperFlexNetworkAddress**](NetworkHyperFlexNetworkAddress.md) |  | [optional] 
**Lockdown** | Pointer to **bool** | Flag indicating whether the HyperFlex host is in lockdown mode. | [optional] [readonly] 
**MgmtIp** | Pointer to [**NullableNetworkHyperFlexNetworkAddress**](NetworkHyperFlexNetworkAddress.md) |  | [optional] 
**OsVersion** | Pointer to **string** | The operation system version of the controller VM. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of the HyperFlex host. * &#x60;UNKNOWN&#x60; - The role of the HyperFlex host is unknown. * &#x60;STORAGE&#x60; - The HyperFlex host&#39;s role is storage. * &#x60;COMPUTE&#x60; - The HyperFlex host&#39;s role is compute. | [optional] [readonly] [default to "UNKNOWN"]
**TemplateVersion** | Pointer to **string** | The controller virtual machine template version. | [optional] [readonly] 
**VirtualCpus** | Pointer to **int32** | Configured number of virtual CPUs for Controller virtual machine. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**Node** | Pointer to [**HyperflexNodeRelationship**](HyperflexNodeRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHypervisorHostAllOf

`func NewHyperflexHypervisorHostAllOf(classId string, objectType string, ) *HyperflexHypervisorHostAllOf`

NewHyperflexHypervisorHostAllOf instantiates a new HyperflexHypervisorHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHypervisorHostAllOfWithDefaults

`func NewHyperflexHypervisorHostAllOfWithDefaults() *HyperflexHypervisorHostAllOf`

NewHyperflexHypervisorHostAllOfWithDefaults instantiates a new HyperflexHypervisorHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHypervisorHostAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHypervisorHostAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHypervisorHostAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHypervisorHostAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHypervisorHostAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHypervisorHostAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfiguredMemory

`func (o *HyperflexHypervisorHostAllOf) GetConfiguredMemory() int64`

GetConfiguredMemory returns the ConfiguredMemory field if non-nil, zero value otherwise.

### GetConfiguredMemoryOk

`func (o *HyperflexHypervisorHostAllOf) GetConfiguredMemoryOk() (*int64, bool)`

GetConfiguredMemoryOk returns a tuple with the ConfiguredMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredMemory

`func (o *HyperflexHypervisorHostAllOf) SetConfiguredMemory(v int64)`

SetConfiguredMemory sets ConfiguredMemory field to given value.

### HasConfiguredMemory

`func (o *HyperflexHypervisorHostAllOf) HasConfiguredMemory() bool`

HasConfiguredMemory returns a boolean if a field has been set.

### GetControllerVmDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmDisplayVersion() string`

GetControllerVmDisplayVersion returns the ControllerVmDisplayVersion field if non-nil, zero value otherwise.

### GetControllerVmDisplayVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmDisplayVersionOk() (*string, bool)`

GetControllerVmDisplayVersionOk returns a tuple with the ControllerVmDisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) SetControllerVmDisplayVersion(v string)`

SetControllerVmDisplayVersion sets ControllerVmDisplayVersion field to given value.

### HasControllerVmDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) HasControllerVmDisplayVersion() bool`

HasControllerVmDisplayVersion returns a boolean if a field has been set.

### GetControllerVmHxdpDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmHxdpDisplayVersion() string`

GetControllerVmHxdpDisplayVersion returns the ControllerVmHxdpDisplayVersion field if non-nil, zero value otherwise.

### GetControllerVmHxdpDisplayVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmHxdpDisplayVersionOk() (*string, bool)`

GetControllerVmHxdpDisplayVersionOk returns a tuple with the ControllerVmHxdpDisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmHxdpDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) SetControllerVmHxdpDisplayVersion(v string)`

SetControllerVmHxdpDisplayVersion sets ControllerVmHxdpDisplayVersion field to given value.

### HasControllerVmHxdpDisplayVersion

`func (o *HyperflexHypervisorHostAllOf) HasControllerVmHxdpDisplayVersion() bool`

HasControllerVmHxdpDisplayVersion returns a boolean if a field has been set.

### GetControllerVmHxdpVersion

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmHxdpVersion() string`

GetControllerVmHxdpVersion returns the ControllerVmHxdpVersion field if non-nil, zero value otherwise.

### GetControllerVmHxdpVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmHxdpVersionOk() (*string, bool)`

GetControllerVmHxdpVersionOk returns a tuple with the ControllerVmHxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmHxdpVersion

`func (o *HyperflexHypervisorHostAllOf) SetControllerVmHxdpVersion(v string)`

SetControllerVmHxdpVersion sets ControllerVmHxdpVersion field to given value.

### HasControllerVmHxdpVersion

`func (o *HyperflexHypervisorHostAllOf) HasControllerVmHxdpVersion() bool`

HasControllerVmHxdpVersion returns a boolean if a field has been set.

### GetControllerVmUuid

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmUuid() string`

GetControllerVmUuid returns the ControllerVmUuid field if non-nil, zero value otherwise.

### GetControllerVmUuidOk

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmUuidOk() (*string, bool)`

GetControllerVmUuidOk returns a tuple with the ControllerVmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmUuid

`func (o *HyperflexHypervisorHostAllOf) SetControllerVmUuid(v string)`

SetControllerVmUuid sets ControllerVmUuid field to given value.

### HasControllerVmUuid

`func (o *HyperflexHypervisorHostAllOf) HasControllerVmUuid() bool`

HasControllerVmUuid returns a boolean if a field has been set.

### GetControllerVmVersion

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmVersion() string`

GetControllerVmVersion returns the ControllerVmVersion field if non-nil, zero value otherwise.

### GetControllerVmVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetControllerVmVersionOk() (*string, bool)`

GetControllerVmVersionOk returns a tuple with the ControllerVmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVmVersion

`func (o *HyperflexHypervisorHostAllOf) SetControllerVmVersion(v string)`

SetControllerVmVersion sets ControllerVmVersion field to given value.

### HasControllerVmVersion

`func (o *HyperflexHypervisorHostAllOf) HasControllerVmVersion() bool`

HasControllerVmVersion returns a boolean if a field has been set.

### GetDataIp

`func (o *HyperflexHypervisorHostAllOf) GetDataIp() NetworkHyperFlexNetworkAddress`

GetDataIp returns the DataIp field if non-nil, zero value otherwise.

### GetDataIpOk

`func (o *HyperflexHypervisorHostAllOf) GetDataIpOk() (*NetworkHyperFlexNetworkAddress, bool)`

GetDataIpOk returns a tuple with the DataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIp

`func (o *HyperflexHypervisorHostAllOf) SetDataIp(v NetworkHyperFlexNetworkAddress)`

SetDataIp sets DataIp field to given value.

### HasDataIp

`func (o *HyperflexHypervisorHostAllOf) HasDataIp() bool`

HasDataIp returns a boolean if a field has been set.

### SetDataIpNil

`func (o *HyperflexHypervisorHostAllOf) SetDataIpNil(b bool)`

 SetDataIpNil sets the value for DataIp to be an explicit nil

### UnsetDataIp
`func (o *HyperflexHypervisorHostAllOf) UnsetDataIp()`

UnsetDataIp ensures that no value is present for DataIp, not even an explicit nil
### GetHostStatus

`func (o *HyperflexHypervisorHostAllOf) GetHostStatus() string`

GetHostStatus returns the HostStatus field if non-nil, zero value otherwise.

### GetHostStatusOk

`func (o *HyperflexHypervisorHostAllOf) GetHostStatusOk() (*string, bool)`

GetHostStatusOk returns a tuple with the HostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostStatus

`func (o *HyperflexHypervisorHostAllOf) SetHostStatus(v string)`

SetHostStatus sets HostStatus field to given value.

### HasHostStatus

`func (o *HyperflexHypervisorHostAllOf) HasHostStatus() bool`

HasHostStatus returns a boolean if a field has been set.

### GetHypervisor

`func (o *HyperflexHypervisorHostAllOf) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *HyperflexHypervisorHostAllOf) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *HyperflexHypervisorHostAllOf) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *HyperflexHypervisorHostAllOf) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetIp

`func (o *HyperflexHypervisorHostAllOf) GetIp() NetworkHyperFlexNetworkAddress`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HyperflexHypervisorHostAllOf) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HyperflexHypervisorHostAllOf) SetIp(v NetworkHyperFlexNetworkAddress)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HyperflexHypervisorHostAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HyperflexHypervisorHostAllOf) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HyperflexHypervisorHostAllOf) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetLockdown

`func (o *HyperflexHypervisorHostAllOf) GetLockdown() bool`

GetLockdown returns the Lockdown field if non-nil, zero value otherwise.

### GetLockdownOk

`func (o *HyperflexHypervisorHostAllOf) GetLockdownOk() (*bool, bool)`

GetLockdownOk returns a tuple with the Lockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockdown

`func (o *HyperflexHypervisorHostAllOf) SetLockdown(v bool)`

SetLockdown sets Lockdown field to given value.

### HasLockdown

`func (o *HyperflexHypervisorHostAllOf) HasLockdown() bool`

HasLockdown returns a boolean if a field has been set.

### GetMgmtIp

`func (o *HyperflexHypervisorHostAllOf) GetMgmtIp() NetworkHyperFlexNetworkAddress`

GetMgmtIp returns the MgmtIp field if non-nil, zero value otherwise.

### GetMgmtIpOk

`func (o *HyperflexHypervisorHostAllOf) GetMgmtIpOk() (*NetworkHyperFlexNetworkAddress, bool)`

GetMgmtIpOk returns a tuple with the MgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIp

`func (o *HyperflexHypervisorHostAllOf) SetMgmtIp(v NetworkHyperFlexNetworkAddress)`

SetMgmtIp sets MgmtIp field to given value.

### HasMgmtIp

`func (o *HyperflexHypervisorHostAllOf) HasMgmtIp() bool`

HasMgmtIp returns a boolean if a field has been set.

### SetMgmtIpNil

`func (o *HyperflexHypervisorHostAllOf) SetMgmtIpNil(b bool)`

 SetMgmtIpNil sets the value for MgmtIp to be an explicit nil

### UnsetMgmtIp
`func (o *HyperflexHypervisorHostAllOf) UnsetMgmtIp()`

UnsetMgmtIp ensures that no value is present for MgmtIp, not even an explicit nil
### GetOsVersion

`func (o *HyperflexHypervisorHostAllOf) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *HyperflexHypervisorHostAllOf) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *HyperflexHypervisorHostAllOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetRole

`func (o *HyperflexHypervisorHostAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HyperflexHypervisorHostAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HyperflexHypervisorHostAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *HyperflexHypervisorHostAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *HyperflexHypervisorHostAllOf) GetTemplateVersion() string`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *HyperflexHypervisorHostAllOf) GetTemplateVersionOk() (*string, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *HyperflexHypervisorHostAllOf) SetTemplateVersion(v string)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *HyperflexHypervisorHostAllOf) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetVirtualCpus

`func (o *HyperflexHypervisorHostAllOf) GetVirtualCpus() int32`

GetVirtualCpus returns the VirtualCpus field if non-nil, zero value otherwise.

### GetVirtualCpusOk

`func (o *HyperflexHypervisorHostAllOf) GetVirtualCpusOk() (*int32, bool)`

GetVirtualCpusOk returns a tuple with the VirtualCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCpus

`func (o *HyperflexHypervisorHostAllOf) SetVirtualCpus(v int32)`

SetVirtualCpus sets VirtualCpus field to given value.

### HasVirtualCpus

`func (o *HyperflexHypervisorHostAllOf) HasVirtualCpus() bool`

HasVirtualCpus returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHypervisorHostAllOf) GetCluster() HyperflexClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHypervisorHostAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHypervisorHostAllOf) SetCluster(v HyperflexClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHypervisorHostAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNode

`func (o *HyperflexHypervisorHostAllOf) GetNode() HyperflexNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HyperflexHypervisorHostAllOf) GetNodeOk() (*HyperflexNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HyperflexHypervisorHostAllOf) SetNode(v HyperflexNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HyperflexHypervisorHostAllOf) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


