# VirtualizationBaseHostPciDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareHostGpu"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareHostGpu"]
**Bus** | Pointer to **int64** | The bus ID of this PCI device. | [optional] [readonly] 
**DeviceId** | Pointer to **int64** | The device ID of this PCI device. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | The device name of this PCI device. | [optional] [readonly] 
**Function** | Pointer to **int64** | The function ID of this PCI device. | [optional] [readonly] 
**Identity** | Pointer to **string** | The internally generated identity of this PCI device. | [optional] 
**PassthroughActive** | Pointer to **bool** | Indicates if passthrough is active for this PCI device (meaning enabled + rebooted). | [optional] [readonly] 
**PassthroughEnabled** | Pointer to **bool** | Indicates if passthrough feature is enabled for this PCI device. PCI passthrough feature enables you to access and manage hardware devices from a virtual machine. When PCI passthrough is configured, the PCI devices function as if they were physically attached to the guest operating system. | [optional] [readonly] 
**PciClassId** | Pointer to **int64** | The class ID of this PCI device. | [optional] [readonly] 
**PciId** | Pointer to **string** | The ID of this PCI device. | [optional] [readonly] 
**Slot** | Pointer to **int64** | The slot ID of this PCI device. | [optional] [readonly] 
**SubDeviceId** | Pointer to **int64** | The sub device ID of this PCI device. | [optional] [readonly] 
**SubVendorId** | Pointer to **int64** | The sub vendor ID of this PCI device. | [optional] [readonly] 
**VendorId** | Pointer to **int64** | The vendor ID of this PCI device. | [optional] [readonly] 
**VendorName** | Pointer to **string** | The vendor name of this PCI device. | [optional] [readonly] 
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](VirtualizationBaseClusterRelationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationBaseHostRelationship**](VirtualizationBaseHostRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationBaseHostPciDeviceAllOf

`func NewVirtualizationBaseHostPciDeviceAllOf(classId string, objectType string, ) *VirtualizationBaseHostPciDeviceAllOf`

NewVirtualizationBaseHostPciDeviceAllOf instantiates a new VirtualizationBaseHostPciDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseHostPciDeviceAllOfWithDefaults

`func NewVirtualizationBaseHostPciDeviceAllOfWithDefaults() *VirtualizationBaseHostPciDeviceAllOf`

NewVirtualizationBaseHostPciDeviceAllOfWithDefaults instantiates a new VirtualizationBaseHostPciDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBus

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetBus() int64`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetBusOk() (*int64, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetBus(v int64)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetFunction

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetFunction(v int64)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetIdentity

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetPassthroughActive

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPassthroughActive() bool`

GetPassthroughActive returns the PassthroughActive field if non-nil, zero value otherwise.

### GetPassthroughActiveOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPassthroughActiveOk() (*bool, bool)`

GetPassthroughActiveOk returns a tuple with the PassthroughActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughActive

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetPassthroughActive(v bool)`

SetPassthroughActive sets PassthroughActive field to given value.

### HasPassthroughActive

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasPassthroughActive() bool`

HasPassthroughActive returns a boolean if a field has been set.

### GetPassthroughEnabled

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPassthroughEnabled() bool`

GetPassthroughEnabled returns the PassthroughEnabled field if non-nil, zero value otherwise.

### GetPassthroughEnabledOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPassthroughEnabledOk() (*bool, bool)`

GetPassthroughEnabledOk returns a tuple with the PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughEnabled

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetPassthroughEnabled(v bool)`

SetPassthroughEnabled sets PassthroughEnabled field to given value.

### HasPassthroughEnabled

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasPassthroughEnabled() bool`

HasPassthroughEnabled returns a boolean if a field has been set.

### GetPciClassId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPciClassId() int64`

GetPciClassId returns the PciClassId field if non-nil, zero value otherwise.

### GetPciClassIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPciClassIdOk() (*int64, bool)`

GetPciClassIdOk returns a tuple with the PciClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciClassId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetPciClassId(v int64)`

SetPciClassId sets PciClassId field to given value.

### HasPciClassId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasPciClassId() bool`

HasPciClassId returns a boolean if a field has been set.

### GetPciId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPciId() string`

GetPciId returns the PciId field if non-nil, zero value otherwise.

### GetPciIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetPciIdOk() (*string, bool)`

GetPciIdOk returns a tuple with the PciId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetPciId(v string)`

SetPciId sets PciId field to given value.

### HasPciId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasPciId() bool`

HasPciId returns a boolean if a field has been set.

### GetSlot

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSlot() int64`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSlotOk() (*int64, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetSlot(v int64)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetSubDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSubDeviceId() int64`

GetSubDeviceId returns the SubDeviceId field if non-nil, zero value otherwise.

### GetSubDeviceIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSubDeviceIdOk() (*int64, bool)`

GetSubDeviceIdOk returns a tuple with the SubDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetSubDeviceId(v int64)`

SetSubDeviceId sets SubDeviceId field to given value.

### HasSubDeviceId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasSubDeviceId() bool`

HasSubDeviceId returns a boolean if a field has been set.

### GetSubVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSubVendorId() int64`

GetSubVendorId returns the SubVendorId field if non-nil, zero value otherwise.

### GetSubVendorIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetSubVendorIdOk() (*int64, bool)`

GetSubVendorIdOk returns a tuple with the SubVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetSubVendorId(v int64)`

SetSubVendorId sets SubVendorId field to given value.

### HasSubVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasSubVendorId() bool`

HasSubVendorId returns a boolean if a field has been set.

### GetVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorName

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetHost() VirtualizationBaseHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationBaseHostPciDeviceAllOf) GetHostOk() (*VirtualizationBaseHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationBaseHostPciDeviceAllOf) SetHost(v VirtualizationBaseHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationBaseHostPciDeviceAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


