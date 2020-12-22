# VnicRoceSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.RoceSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.RoceSettings"]
**ClassOfService** | Pointer to **int32** | The Class of Service for RoCE on this virtual interface. * &#x60;5&#x60; - RDMA CoS Service Level 5. * &#x60;1&#x60; - RDMA CoS Service Level 1. * &#x60;2&#x60; - RDMA CoS Service Level 2. * &#x60;4&#x60; - RDMA CoS Service Level 4. * &#x60;6&#x60; - RDMA CoS Service Level 6. | [optional] [default to 5]
**Enabled** | Pointer to **bool** | If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface. | [optional] 
**MemoryRegions** | Pointer to **int64** | The number of memory regions per adapter. Recommended value &#x3D; integer power of 2. | [optional] [default to 131072]
**QueuePairs** | Pointer to **int64** | The number of queue pairs per adapter. Recommended value &#x3D; integer power of 2. | [optional] [default to 256]
**ResourceGroups** | Pointer to **int64** | The number of resource groups per adapter. Recommended value &#x3D; be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance. | [optional] [default to 4]
**Version** | Pointer to **int32** | Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters. * &#x60;1&#x60; - RDMA over Converged Ethernet Protocol Version 1. * &#x60;2&#x60; - RDMA over Converged Ethernet Protocol Version 2. | [optional] [default to 1]

## Methods

### NewVnicRoceSettingsAllOf

`func NewVnicRoceSettingsAllOf(classId string, objectType string, ) *VnicRoceSettingsAllOf`

NewVnicRoceSettingsAllOf instantiates a new VnicRoceSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicRoceSettingsAllOfWithDefaults

`func NewVnicRoceSettingsAllOfWithDefaults() *VnicRoceSettingsAllOf`

NewVnicRoceSettingsAllOfWithDefaults instantiates a new VnicRoceSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicRoceSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicRoceSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicRoceSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicRoceSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicRoceSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicRoceSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClassOfService

`func (o *VnicRoceSettingsAllOf) GetClassOfService() int32`

GetClassOfService returns the ClassOfService field if non-nil, zero value otherwise.

### GetClassOfServiceOk

`func (o *VnicRoceSettingsAllOf) GetClassOfServiceOk() (*int32, bool)`

GetClassOfServiceOk returns a tuple with the ClassOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassOfService

`func (o *VnicRoceSettingsAllOf) SetClassOfService(v int32)`

SetClassOfService sets ClassOfService field to given value.

### HasClassOfService

`func (o *VnicRoceSettingsAllOf) HasClassOfService() bool`

HasClassOfService returns a boolean if a field has been set.

### GetEnabled

`func (o *VnicRoceSettingsAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicRoceSettingsAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicRoceSettingsAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicRoceSettingsAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMemoryRegions

`func (o *VnicRoceSettingsAllOf) GetMemoryRegions() int64`

GetMemoryRegions returns the MemoryRegions field if non-nil, zero value otherwise.

### GetMemoryRegionsOk

`func (o *VnicRoceSettingsAllOf) GetMemoryRegionsOk() (*int64, bool)`

GetMemoryRegionsOk returns a tuple with the MemoryRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRegions

`func (o *VnicRoceSettingsAllOf) SetMemoryRegions(v int64)`

SetMemoryRegions sets MemoryRegions field to given value.

### HasMemoryRegions

`func (o *VnicRoceSettingsAllOf) HasMemoryRegions() bool`

HasMemoryRegions returns a boolean if a field has been set.

### GetQueuePairs

`func (o *VnicRoceSettingsAllOf) GetQueuePairs() int64`

GetQueuePairs returns the QueuePairs field if non-nil, zero value otherwise.

### GetQueuePairsOk

`func (o *VnicRoceSettingsAllOf) GetQueuePairsOk() (*int64, bool)`

GetQueuePairsOk returns a tuple with the QueuePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePairs

`func (o *VnicRoceSettingsAllOf) SetQueuePairs(v int64)`

SetQueuePairs sets QueuePairs field to given value.

### HasQueuePairs

`func (o *VnicRoceSettingsAllOf) HasQueuePairs() bool`

HasQueuePairs returns a boolean if a field has been set.

### GetResourceGroups

`func (o *VnicRoceSettingsAllOf) GetResourceGroups() int64`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *VnicRoceSettingsAllOf) GetResourceGroupsOk() (*int64, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *VnicRoceSettingsAllOf) SetResourceGroups(v int64)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *VnicRoceSettingsAllOf) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### GetVersion

`func (o *VnicRoceSettingsAllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VnicRoceSettingsAllOf) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VnicRoceSettingsAllOf) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VnicRoceSettingsAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


