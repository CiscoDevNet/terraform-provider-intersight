# VnicRoceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassOfService** | Pointer to **int32** | The Class of Service for RoCE on this virtual interface. * &#x60;5&#x60; - RDMA CoS Service Level 5. * &#x60;1&#x60; - RDMA CoS Service Level 1. * &#x60;2&#x60; - RDMA CoS Service Level 2. * &#x60;4&#x60; - RDMA CoS Service Level 4. * &#x60;6&#x60; - RDMA CoS Service Level 6. | [optional] [default to 5]
**Enabled** | Pointer to **bool** | If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface. | [optional] 
**MemoryRegions** | Pointer to **int64** | The number of memory regions per adapter. Recommended value &#x3D; integer power of 2. | [optional] 
**QueuePairs** | Pointer to **int64** | The number of queue pairs per adapter. Recommended value &#x3D; integer power of 2. | [optional] 
**ResourceGroups** | Pointer to **int64** | The number of resource groups per adapter. Recommended value &#x3D; be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance. | [optional] 
**Version** | Pointer to **int32** | Configures RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoceV1 is supported onn Cisco VIC models 13xx and only RoceV2 is supported on models 14xx. * &#x60;1&#x60; - RDMA over Converged Ethernet Protocol Version 1. * &#x60;2&#x60; - RDMA over Converged Ethernet Protocol Version 2. | [optional] [default to 1]

## Methods

### NewVnicRoceSettings

`func NewVnicRoceSettings() *VnicRoceSettings`

NewVnicRoceSettings instantiates a new VnicRoceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicRoceSettingsWithDefaults

`func NewVnicRoceSettingsWithDefaults() *VnicRoceSettings`

NewVnicRoceSettingsWithDefaults instantiates a new VnicRoceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassOfService

`func (o *VnicRoceSettings) GetClassOfService() int32`

GetClassOfService returns the ClassOfService field if non-nil, zero value otherwise.

### GetClassOfServiceOk

`func (o *VnicRoceSettings) GetClassOfServiceOk() (*int32, bool)`

GetClassOfServiceOk returns a tuple with the ClassOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassOfService

`func (o *VnicRoceSettings) SetClassOfService(v int32)`

SetClassOfService sets ClassOfService field to given value.

### HasClassOfService

`func (o *VnicRoceSettings) HasClassOfService() bool`

HasClassOfService returns a boolean if a field has been set.

### GetEnabled

`func (o *VnicRoceSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicRoceSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicRoceSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicRoceSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMemoryRegions

`func (o *VnicRoceSettings) GetMemoryRegions() int64`

GetMemoryRegions returns the MemoryRegions field if non-nil, zero value otherwise.

### GetMemoryRegionsOk

`func (o *VnicRoceSettings) GetMemoryRegionsOk() (*int64, bool)`

GetMemoryRegionsOk returns a tuple with the MemoryRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRegions

`func (o *VnicRoceSettings) SetMemoryRegions(v int64)`

SetMemoryRegions sets MemoryRegions field to given value.

### HasMemoryRegions

`func (o *VnicRoceSettings) HasMemoryRegions() bool`

HasMemoryRegions returns a boolean if a field has been set.

### GetQueuePairs

`func (o *VnicRoceSettings) GetQueuePairs() int64`

GetQueuePairs returns the QueuePairs field if non-nil, zero value otherwise.

### GetQueuePairsOk

`func (o *VnicRoceSettings) GetQueuePairsOk() (*int64, bool)`

GetQueuePairsOk returns a tuple with the QueuePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePairs

`func (o *VnicRoceSettings) SetQueuePairs(v int64)`

SetQueuePairs sets QueuePairs field to given value.

### HasQueuePairs

`func (o *VnicRoceSettings) HasQueuePairs() bool`

HasQueuePairs returns a boolean if a field has been set.

### GetResourceGroups

`func (o *VnicRoceSettings) GetResourceGroups() int64`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *VnicRoceSettings) GetResourceGroupsOk() (*int64, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *VnicRoceSettings) SetResourceGroups(v int64)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *VnicRoceSettings) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### GetVersion

`func (o *VnicRoceSettings) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VnicRoceSettings) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VnicRoceSettings) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VnicRoceSettings) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


