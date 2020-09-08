# VirtualizationVmwareHostAllOf

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

### NewVirtualizationVmwareHostAllOf

`func NewVirtualizationVmwareHostAllOf() *VirtualizationVmwareHostAllOf`

NewVirtualizationVmwareHostAllOf instantiates a new VirtualizationVmwareHostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareHostAllOfWithDefaults

`func NewVirtualizationVmwareHostAllOfWithDefaults() *VirtualizationVmwareHostAllOf`

NewVirtualizationVmwareHostAllOfWithDefaults instantiates a new VirtualizationVmwareHostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootTime

`func (o *VirtualizationVmwareHostAllOf) GetBootTime() time.Time`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *VirtualizationVmwareHostAllOf) GetBootTimeOk() (*time.Time, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *VirtualizationVmwareHostAllOf) SetBootTime(v time.Time)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *VirtualizationVmwareHostAllOf) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetConnectionState

`func (o *VirtualizationVmwareHostAllOf) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VirtualizationVmwareHostAllOf) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VirtualizationVmwareHostAllOf) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *VirtualizationVmwareHostAllOf) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetHwPowerState

`func (o *VirtualizationVmwareHostAllOf) GetHwPowerState() string`

GetHwPowerState returns the HwPowerState field if non-nil, zero value otherwise.

### GetHwPowerStateOk

`func (o *VirtualizationVmwareHostAllOf) GetHwPowerStateOk() (*string, bool)`

GetHwPowerStateOk returns a tuple with the HwPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwPowerState

`func (o *VirtualizationVmwareHostAllOf) SetHwPowerState(v string)`

SetHwPowerState sets HwPowerState field to given value.

### HasHwPowerState

`func (o *VirtualizationVmwareHostAllOf) HasHwPowerState() bool`

HasHwPowerState returns a boolean if a field has been set.

### GetNetworkAdapterCount

`func (o *VirtualizationVmwareHostAllOf) GetNetworkAdapterCount() int64`

GetNetworkAdapterCount returns the NetworkAdapterCount field if non-nil, zero value otherwise.

### GetNetworkAdapterCountOk

`func (o *VirtualizationVmwareHostAllOf) GetNetworkAdapterCountOk() (*int64, bool)`

GetNetworkAdapterCountOk returns a tuple with the NetworkAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterCount

`func (o *VirtualizationVmwareHostAllOf) SetNetworkAdapterCount(v int64)`

SetNetworkAdapterCount sets NetworkAdapterCount field to given value.

### HasNetworkAdapterCount

`func (o *VirtualizationVmwareHostAllOf) HasNetworkAdapterCount() bool`

HasNetworkAdapterCount returns a boolean if a field has been set.

### GetResourceConsumed

`func (o *VirtualizationVmwareHostAllOf) GetResourceConsumed() VirtualizationVmwareResourceConsumption`

GetResourceConsumed returns the ResourceConsumed field if non-nil, zero value otherwise.

### GetResourceConsumedOk

`func (o *VirtualizationVmwareHostAllOf) GetResourceConsumedOk() (*VirtualizationVmwareResourceConsumption, bool)`

GetResourceConsumedOk returns a tuple with the ResourceConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConsumed

`func (o *VirtualizationVmwareHostAllOf) SetResourceConsumed(v VirtualizationVmwareResourceConsumption)`

SetResourceConsumed sets ResourceConsumed field to given value.

### HasResourceConsumed

`func (o *VirtualizationVmwareHostAllOf) HasResourceConsumed() bool`

HasResourceConsumed returns a boolean if a field has been set.

### GetStorageAdapterCount

`func (o *VirtualizationVmwareHostAllOf) GetStorageAdapterCount() int64`

GetStorageAdapterCount returns the StorageAdapterCount field if non-nil, zero value otherwise.

### GetStorageAdapterCountOk

`func (o *VirtualizationVmwareHostAllOf) GetStorageAdapterCountOk() (*int64, bool)`

GetStorageAdapterCountOk returns a tuple with the StorageAdapterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAdapterCount

`func (o *VirtualizationVmwareHostAllOf) SetStorageAdapterCount(v int64)`

SetStorageAdapterCount sets StorageAdapterCount field to given value.

### HasStorageAdapterCount

`func (o *VirtualizationVmwareHostAllOf) HasStorageAdapterCount() bool`

HasStorageAdapterCount returns a boolean if a field has been set.

### GetVcenterHostId

`func (o *VirtualizationVmwareHostAllOf) GetVcenterHostId() string`

GetVcenterHostId returns the VcenterHostId field if non-nil, zero value otherwise.

### GetVcenterHostIdOk

`func (o *VirtualizationVmwareHostAllOf) GetVcenterHostIdOk() (*string, bool)`

GetVcenterHostIdOk returns a tuple with the VcenterHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterHostId

`func (o *VirtualizationVmwareHostAllOf) SetVcenterHostId(v string)`

SetVcenterHostId sets VcenterHostId field to given value.

### HasVcenterHostId

`func (o *VirtualizationVmwareHostAllOf) HasVcenterHostId() bool`

HasVcenterHostId returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVmwareHostAllOf) GetCluster() VirtualizationVmwareClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVmwareHostAllOf) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVmwareHostAllOf) SetCluster(v VirtualizationVmwareClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVmwareHostAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareHostAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareHostAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareHostAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareHostAllOf) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatastores

`func (o *VirtualizationVmwareHostAllOf) GetDatastores() []VirtualizationVmwareDatastoreRelationship`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualizationVmwareHostAllOf) GetDatastoresOk() (*[]VirtualizationVmwareDatastoreRelationship, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualizationVmwareHostAllOf) SetDatastores(v []VirtualizationVmwareDatastoreRelationship)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *VirtualizationVmwareHostAllOf) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *VirtualizationVmwareHostAllOf) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *VirtualizationVmwareHostAllOf) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


