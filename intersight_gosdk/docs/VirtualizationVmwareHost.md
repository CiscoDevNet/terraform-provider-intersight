# VirtualizationVmwareHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareHost"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareHost"]
**BootTime** | Pointer to **time.Time** | The time when this host booted up. | [optional] 
**ConnectionState** | Pointer to **string** | Indicates if the host is connected to the vCenter. Values are connected, not connected. | [optional] 
**DcInvPath** | Pointer to **string** | This field stores the inventory path of a datacenter. Used in host maintenance action. | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweringOn&#x60; - The entity is powering on. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweringOff&#x60; - The entity is powering off. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. * &#x60;Rebooting&#x60; - The entity reboot is in progress. * &#x60;&#x60; - The entity&#39;s power state is not available. | [optional] [default to "Unknown"]
**IsSshEnabled** | Pointer to **bool** | True if SSH is enabled in the host, false otherwise. | [optional] 
**NetworkAdapterCount** | Pointer to **int64** | The count of all network adapters attached to this host. | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**QuarantineMode** | Pointer to **bool** | Indicates if the host is in quarantine mode. Will be set to True, when in quarantine mode. | [optional] 
**ResourceConsumed** | Pointer to [**NullableVirtualizationVmwareResourceConsumption**](VirtualizationVmwareResourceConsumption.md) |  | [optional] 
**StorageAdapterCount** | Pointer to **int64** | The count of all storage adapters attached to this host. | [optional] 
**TimeZone** | Pointer to **string** | Time zone this host is in. | [optional] 
**VcenterHostId** | Pointer to **string** | The identity of this host within vCenter (optional). | [optional] 
**Cluster** | Pointer to [**NullableVirtualizationVmwareClusterRelationship**](VirtualizationVmwareClusterRelationship.md) |  | [optional] 
**Datacenter** | Pointer to [**NullableVirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 
**Datastores** | Pointer to [**[]VirtualizationVmwareDatastoreRelationship**](VirtualizationVmwareDatastoreRelationship.md) | An array of relationships to virtualizationVmwareDatastore resources. | [optional] [readonly] 
**DistributedNetworks** | Pointer to [**[]VirtualizationVmwareDistributedNetworkRelationship**](VirtualizationVmwareDistributedNetworkRelationship.md) | An array of relationships to virtualizationVmwareDistributedNetwork resources. | [optional] [readonly] 
**DistributedSwitches** | Pointer to [**[]VirtualizationVmwareDistributedSwitchRelationship**](VirtualizationVmwareDistributedSwitchRelationship.md) | An array of relationships to virtualizationVmwareDistributedSwitch resources. | [optional] [readonly] 
**HyperFlexNode** | Pointer to [**NullableHyperflexNodeRelationship**](HyperflexNodeRelationship.md) |  | [optional] 
**Server** | Pointer to [**NullableComputePhysicalSummaryRelationship**](ComputePhysicalSummaryRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareHost

`func NewVirtualizationVmwareHost(classId string, objectType string, ) *VirtualizationVmwareHost`

NewVirtualizationVmwareHost instantiates a new VirtualizationVmwareHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareHostWithDefaults

`func NewVirtualizationVmwareHostWithDefaults() *VirtualizationVmwareHost`

NewVirtualizationVmwareHostWithDefaults instantiates a new VirtualizationVmwareHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareHost) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareHost) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareHost) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareHost) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareHost) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareHost) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootTime

`func (o *VirtualizationVmwareHost) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationVmwareHost) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationVmwareHost) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationVmwareHost) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetConnectionState

`func (o *VirtualizationVmwareHost) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VirtualizationVmwareHost) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VirtualizationVmwareHost) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *VirtualizationVmwareHost) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetDcInvPath

`func (o *VirtualizationVmwareHost) GetDcInvPath() string`

GetDcInvPath returns the DcInvPath field if non-nil, zero value otherwise.

### GetDcInvPathOk

`func (o *VirtualizationVmwareHost) GetDcInvPathOk() (*string, bool)`

GetDcInvPathOk returns a tuple with the DcInvPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcInvPath

`func (o *VirtualizationVmwareHost) SetDcInvPath(v string)`

SetDcInvPath sets DcInvPath field to given value.

### HasDcInvPath

`func (o *VirtualizationVmwareHost) HasDcInvPath() bool`

HasDcInvPath returns a boolean if a field has been set.

### GetDnsServers

`func (o *VirtualizationVmwareHost) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VirtualizationVmwareHost) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VirtualizationVmwareHost) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *VirtualizationVmwareHost) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *VirtualizationVmwareHost) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *VirtualizationVmwareHost) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetHwPowerState

`func (o *VirtualizationVmwareHost) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *VirtualizationVmwareHost) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *VirtualizationVmwareHost) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *VirtualizationVmwareHost) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetIsSshEnabled

`func (o *VirtualizationVmwareHost) GetIsSshEnabled() bool`

GetIsSshEnabled returns the IsSshEnabled field if non-nil, zero value otherwise.

### GetIsSshEnabledOk

`func (o *VirtualizationVmwareHost) GetIsSshEnabledOk() (*bool, bool)`

GetIsSshEnabledOk returns a tuple with the IsSshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSshEnabled

`func (o *VirtualizationVmwareHost) SetIsSshEnabled(v bool)`

SetIsSshEnabled sets IsSshEnabled field to given value.

### HasIsSshEnabled

`func (o *VirtualizationVmwareHost) HasIsSshEnabled() bool`

HasIsSshEnabled returns a boolean if a field has been set.

### GetNetworkAdapterCount

`func (o *VirtualizationVmwareHost) GetNetworkAdapterCount() int64`

GetNetworkAdapterCount returns the NetworkAdapterCount field if non-nil, zero value otherwise.

### GetNetworkAdapterCountOk

`func (o *VirtualizationVmwareHost) GetNetworkAdapterCountOk() (*int64, bool)`

GetNetworkAdapterCountOk returns a tuple with the NetworkAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterCount

`func (o *VirtualizationVmwareHost) SetNetworkAdapterCount(v int64)`

SetNetworkAdapterCount sets NetworkAdapterCount field to given value.

### HasNetworkAdapterCount

`func (o *VirtualizationVmwareHost) HasNetworkAdapterCount() bool`

HasNetworkAdapterCount returns a boolean if a field has been set.

### GetNtpServers

`func (o *VirtualizationVmwareHost) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *VirtualizationVmwareHost) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *VirtualizationVmwareHost) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *VirtualizationVmwareHost) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *VirtualizationVmwareHost) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *VirtualizationVmwareHost) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetQuarantineMode

`func (o *VirtualizationVmwareHost) GetQuarantineMode() bool`

GetQuarantineMode returns the QuarantineMode field if non-nil, zero value otherwise.

### GetQuarantineModeOk

`func (o *VirtualizationVmwareHost) GetQuarantineModeOk() (*bool, bool)`

GetQuarantineModeOk returns a tuple with the QuarantineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineMode

`func (o *VirtualizationVmwareHost) SetQuarantineMode(v bool)`

SetQuarantineMode sets QuarantineMode field to given value.

### HasQuarantineMode

`func (o *VirtualizationVmwareHost) HasQuarantineMode() bool`

HasQuarantineMode returns a boolean if a field has been set.

### GetResourceConsumed

`func (o *VirtualizationVmwareHost) GetResourceConsumed() VirtualizationVmwareResourceConsumption`

GetResourceConsumed returns the ResourceConsumed field if non-nil, zero value otherwise.

### GetResourceConsumedOk

`func (o *VirtualizationVmwareHost) GetResourceConsumedOk() (*VirtualizationVmwareResourceConsumption, bool)`

GetResourceConsumedOk returns a tuple with the ResourceConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConsumed

`func (o *VirtualizationVmwareHost) SetResourceConsumed(v VirtualizationVmwareResourceConsumption)`

SetResourceConsumed sets ResourceConsumed field to given value.

### HasResourceConsumed

`func (o *VirtualizationVmwareHost) HasResourceConsumed() bool`

HasResourceConsumed returns a boolean if a field has been set.

### SetResourceConsumedNil

`func (o *VirtualizationVmwareHost) SetResourceConsumedNil(b bool)`

 SetResourceConsumedNil sets the value for ResourceConsumed to be an explicit nil

### UnsetResourceConsumed
`func (o *VirtualizationVmwareHost) UnsetResourceConsumed()`

UnsetResourceConsumed ensures that no value is present for ResourceConsumed, not even an explicit nil
### GetStorageAdapterCount

`func (o *VirtualizationVmwareHost) GetStorageAdapterCount() int64`

GetStorageAdapterCount returns the StorageAdapterCount field if non-nil, zero value otherwise.

### GetStorageAdapterCountOk

`func (o *VirtualizationVmwareHost) GetStorageAdapterCountOk() (*int64, bool)`

GetStorageAdapterCountOk returns a tuple with the StorageAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAdapterCount

`func (o *VirtualizationVmwareHost) SetStorageAdapterCount(v int64)`

SetStorageAdapterCount sets StorageAdapterCount field to given value.

### HasStorageAdapterCount

`func (o *VirtualizationVmwareHost) HasStorageAdapterCount() bool`

HasStorageAdapterCount returns a boolean if a field has been set.

### GetTimeZone

`func (o *VirtualizationVmwareHost) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *VirtualizationVmwareHost) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *VirtualizationVmwareHost) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *VirtualizationVmwareHost) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetVcenterHostId

`func (o *VirtualizationVmwareHost) GetVcenterHostId() string`

GetVcenterHostId returns the VcenterHostId field if non-nil, zero value otherwise.

### GetVcenterHostIdOk

`func (o *VirtualizationVmwareHost) GetVcenterHostIdOk() (*string, bool)`

GetVcenterHostIdOk returns a tuple with the VcenterHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterHostId

`func (o *VirtualizationVmwareHost) SetVcenterHostId(v string)`

SetVcenterHostId sets VcenterHostId field to given value.

### HasVcenterHostId

`func (o *VirtualizationVmwareHost) HasVcenterHostId() bool`

HasVcenterHostId returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareHost) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareHost) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareHost) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareHost) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *VirtualizationVmwareHost) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *VirtualizationVmwareHost) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDatacenter

`func (o *VirtualizationVmwareHost) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareHost) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareHost) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareHost) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### SetDatacenterNil

`func (o *VirtualizationVmwareHost) SetDatacenterNil(b bool)`

 SetDatacenterNil sets the value for Datacenter to be an explicit nil

### UnsetDatacenter
`func (o *VirtualizationVmwareHost) UnsetDatacenter()`

UnsetDatacenter ensures that no value is present for Datacenter, not even an explicit nil
### GetDatastores

`func (o *VirtualizationVmwareHost) GetDatastores() []VirtualizationVmwareDatastoreRelationship`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualizationVmwareHost) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualizationVmwareHost) SetDatastores(v []VirtualizationVmwareDatastoreRelationship)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VirtualizationVmwareHost) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *VirtualizationVmwareHost) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *VirtualizationVmwareHost) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetDistributedNetworks

`func (o *VirtualizationVmwareHost) GetDistributedNetworks() []VirtualizationVmwareDistributedNetworkRelationship`

GetDistributedNetworks returns the DistributedNetworks field if non-nil, zero value otherwise.

### GetDistributedNetworksOk

`func (o *VirtualizationVmwareHost) GetDistributedNetworksOk() (*[]VirtualizationVmwareDistributedNetworkRelationship, bool)`

GetDistributedNetworksOk returns a tuple with the DistributedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedNetworks

`func (o *VirtualizationVmwareHost) SetDistributedNetworks(v []VirtualizationVmwareDistributedNetworkRelationship)`

SetDistributedNetworks sets DistributedNetworks field to given value.

### HasDistributedNetworks

`func (o *VirtualizationVmwareHost) HasDistributedNetworks() bool`

HasDistributedNetworks returns a boolean if a field has been set.

### SetDistributedNetworksNil

`func (o *VirtualizationVmwareHost) SetDistributedNetworksNil(b bool)`

 SetDistributedNetworksNil sets the value for DistributedNetworks to be an explicit nil

### UnsetDistributedNetworks
`func (o *VirtualizationVmwareHost) UnsetDistributedNetworks()`

UnsetDistributedNetworks ensures that no value is present for DistributedNetworks, not even an explicit nil
### GetDistributedSwitches

`func (o *VirtualizationVmwareHost) GetDistributedSwitches() []VirtualizationVmwareDistributedSwitchRelationship`

GetDistributedSwitches returns the DistributedSwitches field if non-nil, zero value otherwise.

### GetDistributedSwitchesOk

`func (o *VirtualizationVmwareHost) GetDistributedSwitchesOk() (*[]VirtualizationVmwareDistributedSwitchRelationship, bool)`

GetDistributedSwitchesOk returns a tuple with the DistributedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedSwitches

`func (o *VirtualizationVmwareHost) SetDistributedSwitches(v []VirtualizationVmwareDistributedSwitchRelationship)`

SetDistributedSwitches sets DistributedSwitches field to given value.

### HasDistributedSwitches

`func (o *VirtualizationVmwareHost) HasDistributedSwitches() bool`

HasDistributedSwitches returns a boolean if a field has been set.

### SetDistributedSwitchesNil

`func (o *VirtualizationVmwareHost) SetDistributedSwitchesNil(b bool)`

 SetDistributedSwitchesNil sets the value for DistributedSwitches to be an explicit nil

### UnsetDistributedSwitches
`func (o *VirtualizationVmwareHost) UnsetDistributedSwitches()`

UnsetDistributedSwitches ensures that no value is present for DistributedSwitches, not even an explicit nil
### GetHyperFlexNode

`func (o *VirtualizationVmwareHost) GetHyperFlexNode() HyperflexNodeRelationship`

GetHyperFlexNode returns the HyperFlexNode field if non-nil, zero value otherwise.

### GetHyperFlexNodeOk

`func (o *VirtualizationVmwareHost) GetHyperFlexNodeOk() (*HyperflexNodeRelationship, bool)`

GetHyperFlexNodeOk returns a tuple with the HyperFlexNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperFlexNode

`func (o *VirtualizationVmwareHost) SetHyperFlexNode(v HyperflexNodeRelationship)`

SetHyperFlexNode sets HyperFlexNode field to given value.

### HasHyperFlexNode

`func (o *VirtualizationVmwareHost) HasHyperFlexNode() bool`

HasHyperFlexNode returns a boolean if a field has been set.

### SetHyperFlexNodeNil

`func (o *VirtualizationVmwareHost) SetHyperFlexNodeNil(b bool)`

 SetHyperFlexNodeNil sets the value for HyperFlexNode to be an explicit nil

### UnsetHyperFlexNode
`func (o *VirtualizationVmwareHost) UnsetHyperFlexNode()`

UnsetHyperFlexNode ensures that no value is present for HyperFlexNode, not even an explicit nil
### GetServer

`func (o *VirtualizationVmwareHost) GetServer() ComputePhysicalSummaryRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VirtualizationVmwareHost) GetServerOk() (*ComputePhysicalSummaryRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VirtualizationVmwareHost) SetServer(v ComputePhysicalSummaryRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *VirtualizationVmwareHost) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *VirtualizationVmwareHost) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *VirtualizationVmwareHost) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


