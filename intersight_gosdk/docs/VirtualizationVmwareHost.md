# VirtualizationVmwareHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootTime** | Pointer to [**time.Time**](time.Time.md) | The time when this host booted up. | [optional] 
**ConnectionState** | Pointer to **string** | Indicates if the host is connected to the vCenter. Values are connected, not connected. | [optional] 
**HwPowerState** | Pointer to **string** | Is the host Powered-up or Powered-down. * &#x60;Unknown&#x60; - The entity&#39;s power state is unknown. * &#x60;PoweredOn&#x60; - The entity is powered on. * &#x60;PoweredOff&#x60; - The entity is powered down. * &#x60;StandBy&#x60; - The entity is in standby mode. * &#x60;Paused&#x60; - The entity is in pause state. | [optional] [default to "Unknown"]
**NetworkAdapterCount** | Pointer to **int64** | The count of all network adapters attached to this host. | [optional] 
**ResourceConsumed** | Pointer to [**VirtualizationVmwareResourceConsumption**](virtualization.VmwareResourceConsumption.md) |  | [optional] 
**StorageAdapterCount** | Pointer to **int64** | The count of all storage adapters attached to this host. | [optional] 
**VcenterHostId** | Pointer to **string** | The identity of this host within vCenter (optional). | [optional] 
**Cluster** | Pointer to [**VirtualizationVmwareClusterRelationship**](virtualization.VmwareCluster.Relationship.md) |  | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](virtualization.VmwareDatacenter.Relationship.md) |  | [optional] 
**Datastores** | Pointer to [**[]VirtualizationVmwareDatastoreRelationship**](virtualization.VmwareDatastore.Relationship.md) | An array of relationships to virtualizationVmwareDatastore resources. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareHost

`func NewVirtualizationVmwareHost() *VirtualizationVmwareHost`

NewVirtualizationVmwareHost instantiates a new VirtualizationVmwareHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareHostWithDefaults

`func NewVirtualizationVmwareHostWithDefaults() *VirtualizationVmwareHost`

NewVirtualizationVmwareHostWithDefaults instantiates a new VirtualizationVmwareHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


