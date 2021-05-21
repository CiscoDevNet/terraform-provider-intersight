# VirtualizationVmwareVirtualNetworkInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVirtualNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVirtualNetworkInterface"]
**AdapterType** | Pointer to **string** | Type of virtual ethernet adapter for virtual network interface. | [optional] 
**ConnectAtPowerOn** | Pointer to **bool** | Connect or not to connect the device when the virtual machine starts. | [optional] 
**Connected** | Pointer to **bool** | Device is currently connected or not. Valid only while the virtual machine is running. | [optional] 
**Key** | Pointer to **int64** | The internally assigned key of this virtual network interface. This entity is not manipulated by users. | [optional] 
**MacAddress** | Pointer to **string** | MAC address assigned to virtual network interface. | [optional] 
**MacAddressType** | Pointer to **string** | MAC address type for the mac address assigned to virtual network interface. * &#x60;manual&#x60; - Statically assigned MAC address. * &#x60;generated&#x60; - Automatically generated MAC address. * &#x60;assigned&#x60; - MAC address assigned by VCenter to the virtual network interface card. | [optional] [default to "manual"]
**NetworkType** | Pointer to **string** | Type of network for virtual network interface. It can be either standard or distributed. | [optional] 
**VmIdentity** | Pointer to **string** | Identity of the virtual machine where the virtual network interface is created. | [optional] 
**Network** | Pointer to [**VirtualizationBaseNetworkRelationship**](virtualization.BaseNetwork.Relationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**VirtualizationVmwareVirtualMachineRelationship**](virtualization.VmwareVirtualMachine.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareVirtualNetworkInterfaceAllOf

`func NewVirtualizationVmwareVirtualNetworkInterfaceAllOf(classId string, objectType string, ) *VirtualizationVmwareVirtualNetworkInterfaceAllOf`

NewVirtualizationVmwareVirtualNetworkInterfaceAllOf instantiates a new VirtualizationVmwareVirtualNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualNetworkInterfaceAllOfWithDefaults

`func NewVirtualizationVmwareVirtualNetworkInterfaceAllOfWithDefaults() *VirtualizationVmwareVirtualNetworkInterfaceAllOf`

NewVirtualizationVmwareVirtualNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetAdapterType() string`

GetAdapterType returns the AdapterType field if non-nil, zero value otherwise.

### GetAdapterTypeOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetAdapterTypeOk() (*string, bool)`

GetAdapterTypeOk returns a tuple with the AdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetAdapterType(v string)`

SetAdapterType sets AdapterType field to given value.

### HasAdapterType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasAdapterType() bool`

HasAdapterType returns a boolean if a field has been set.

### GetConnectAtPowerOn

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetConnectAtPowerOn() bool`

GetConnectAtPowerOn returns the ConnectAtPowerOn field if non-nil, zero value otherwise.

### GetConnectAtPowerOnOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetConnectAtPowerOnOk() (*bool, bool)`

GetConnectAtPowerOnOk returns a tuple with the ConnectAtPowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAtPowerOn

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetConnectAtPowerOn(v bool)`

SetConnectAtPowerOn sets ConnectAtPowerOn field to given value.

### HasConnectAtPowerOn

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasConnectAtPowerOn() bool`

HasConnectAtPowerOn returns a boolean if a field has been set.

### GetConnected

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetKey

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetKey(v int64)`

SetKey sets Key field to given value.

### HasKey

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacAddressType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

### GetNetworkType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetVmIdentity

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetVmIdentity() string`

GetVmIdentity returns the VmIdentity field if non-nil, zero value otherwise.

### GetVmIdentityOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetVmIdentityOk() (*string, bool)`

GetVmIdentityOk returns a tuple with the VmIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIdentity

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetVmIdentity(v string)`

SetVmIdentity sets VmIdentity field to given value.

### HasVmIdentity

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasVmIdentity() bool`

HasVmIdentity returns a boolean if a field has been set.

### GetNetwork

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetNetwork() VirtualizationBaseNetworkRelationship`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetNetworkOk() (*VirtualizationBaseNetworkRelationship, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetNetwork(v VirtualizationBaseNetworkRelationship)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VirtualizationVmwareVirtualNetworkInterfaceAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


