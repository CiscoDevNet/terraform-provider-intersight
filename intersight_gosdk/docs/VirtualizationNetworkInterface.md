# VirtualizationNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.NetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.NetworkInterface"]
**AdaptorType** | Pointer to **string** | Virtual machine network adaptor type. * &#x60;Unknown&#x60; - The type of the network adaptor type is unknown. * &#x60;E1000&#x60; - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * &#x60;SRIOV&#x60; - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * &#x60;VMXNET2&#x60; - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * &#x60;VMXNET3&#x60; - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * &#x60;E1000E&#x60; - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * &#x60;NE2K_PCI&#x60; - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip&#39;s DMA logic will use to store received packets or to get received packets. * &#x60;PCnet&#x60; - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * &#x60;RTL8139&#x60; - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * &#x60;VirtIO&#x60; - VirtIO is a standardized interface which allows virtual machines access to simplified \&quot;virtual\&quot; devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \&quot;emulated\&quot; devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * &#x60;&#x60; - Default network adaptor type supported by the hypervisor. | [optional] [default to "Unknown"]
**Bridge** | Pointer to **string** | Virtual machine network bridge name. | [optional] 
**ConnectAtPowerOn** | Pointer to **bool** | Connect the adaptor at virtual machine power on. | [optional] 
**DirectPathIo** | Pointer to **bool** | Enable the direct path I/O. | [optional] 
**IpForwardingEnabled** | Pointer to **bool** | Set to true, if IP forwarding is enabled on the NIC. | [optional] 
**Ipv6Address** | Pointer to **bool** | Set to true, if IPv6 address should be allocated for the NIC. | [optional] 
**MacAddress** | Pointer to **string** | Virtual machine network mac address. | [optional] 
**Name** | Pointer to **string** | Name of the network interface. This may be different from guest operating system assigned. | [optional] 
**NetworkId** | Pointer to **string** | Identity of the network to which this network interface belongs. | [optional] 
**NicId** | Pointer to **string** | Identity of the network interface. | [optional] 
**Order** | Pointer to **int64** | Order of the NIC attachment to the VM. | [optional] 
**PrivateIpAllocationMode** | Pointer to **string** | Allocation mode for NIC addresses e.g. DHCP or static. * &#x60;DHCP&#x60; - Dynamic IP address allocation using DHCP protocol. * &#x60;STATIC_IP&#x60; - Assign fixed / static IPs to resources for use. * &#x60;IPAM_CALLOUT&#x60; - Use callout scripts to query cloud IP allocation tools to assign network parameters. * &#x60;PREALLOCATE_IP&#x60; - Allows the cloud infrastructure IP allocation to be dynamically provided before the server boots up. | [optional] [default to "DHCP"]
**PublicIpAllocate** | Pointer to **bool** | Set to true, if public IP should be allocated for the NIC. | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**StaticIpAddress** | Pointer to [**[]VirtualizationIpAddressInfo**](VirtualizationIpAddressInfo.md) |  | [optional] 
**SubnetId** | Pointer to **string** | Subnet identifier for the NIC. | [optional] 

## Methods

### NewVirtualizationNetworkInterface

`func NewVirtualizationNetworkInterface(classId string, objectType string, ) *VirtualizationNetworkInterface`

NewVirtualizationNetworkInterface instantiates a new VirtualizationNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationNetworkInterfaceWithDefaults

`func NewVirtualizationNetworkInterfaceWithDefaults() *VirtualizationNetworkInterface`

NewVirtualizationNetworkInterfaceWithDefaults instantiates a new VirtualizationNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdaptorType

`func (o *VirtualizationNetworkInterface) GetAdaptorType() string`

GetAdaptorType returns the AdaptorType field if non-nil, zero value otherwise.

### GetAdaptorTypeOk

`func (o *VirtualizationNetworkInterface) GetAdaptorTypeOk() (*string, bool)`

GetAdaptorTypeOk returns a tuple with the AdaptorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptorType

`func (o *VirtualizationNetworkInterface) SetAdaptorType(v string)`

SetAdaptorType sets AdaptorType field to given value.

### HasAdaptorType

`func (o *VirtualizationNetworkInterface) HasAdaptorType() bool`

HasAdaptorType returns a boolean if a field has been set.

### GetBridge

`func (o *VirtualizationNetworkInterface) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *VirtualizationNetworkInterface) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *VirtualizationNetworkInterface) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *VirtualizationNetworkInterface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetConnectAtPowerOn

`func (o *VirtualizationNetworkInterface) GetConnectAtPowerOn() bool`

GetConnectAtPowerOn returns the ConnectAtPowerOn field if non-nil, zero value otherwise.

### GetConnectAtPowerOnOk

`func (o *VirtualizationNetworkInterface) GetConnectAtPowerOnOk() (*bool, bool)`

GetConnectAtPowerOnOk returns a tuple with the ConnectAtPowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAtPowerOn

`func (o *VirtualizationNetworkInterface) SetConnectAtPowerOn(v bool)`

SetConnectAtPowerOn sets ConnectAtPowerOn field to given value.

### HasConnectAtPowerOn

`func (o *VirtualizationNetworkInterface) HasConnectAtPowerOn() bool`

HasConnectAtPowerOn returns a boolean if a field has been set.

### GetDirectPathIo

`func (o *VirtualizationNetworkInterface) GetDirectPathIo() bool`

GetDirectPathIo returns the DirectPathIo field if non-nil, zero value otherwise.

### GetDirectPathIoOk

`func (o *VirtualizationNetworkInterface) GetDirectPathIoOk() (*bool, bool)`

GetDirectPathIoOk returns a tuple with the DirectPathIo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPathIo

`func (o *VirtualizationNetworkInterface) SetDirectPathIo(v bool)`

SetDirectPathIo sets DirectPathIo field to given value.

### HasDirectPathIo

`func (o *VirtualizationNetworkInterface) HasDirectPathIo() bool`

HasDirectPathIo returns a boolean if a field has been set.

### GetIpForwardingEnabled

`func (o *VirtualizationNetworkInterface) GetIpForwardingEnabled() bool`

GetIpForwardingEnabled returns the IpForwardingEnabled field if non-nil, zero value otherwise.

### GetIpForwardingEnabledOk

`func (o *VirtualizationNetworkInterface) GetIpForwardingEnabledOk() (*bool, bool)`

GetIpForwardingEnabledOk returns a tuple with the IpForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpForwardingEnabled

`func (o *VirtualizationNetworkInterface) SetIpForwardingEnabled(v bool)`

SetIpForwardingEnabled sets IpForwardingEnabled field to given value.

### HasIpForwardingEnabled

`func (o *VirtualizationNetworkInterface) HasIpForwardingEnabled() bool`

HasIpForwardingEnabled returns a boolean if a field has been set.

### GetIpv6Address

`func (o *VirtualizationNetworkInterface) GetIpv6Address() bool`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *VirtualizationNetworkInterface) GetIpv6AddressOk() (*bool, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *VirtualizationNetworkInterface) SetIpv6Address(v bool)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *VirtualizationNetworkInterface) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *VirtualizationNetworkInterface) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VirtualizationNetworkInterface) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VirtualizationNetworkInterface) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VirtualizationNetworkInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNicId

`func (o *VirtualizationNetworkInterface) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *VirtualizationNetworkInterface) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *VirtualizationNetworkInterface) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *VirtualizationNetworkInterface) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetOrder

`func (o *VirtualizationNetworkInterface) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VirtualizationNetworkInterface) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VirtualizationNetworkInterface) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VirtualizationNetworkInterface) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPrivateIpAllocationMode

`func (o *VirtualizationNetworkInterface) GetPrivateIpAllocationMode() string`

GetPrivateIpAllocationMode returns the PrivateIpAllocationMode field if non-nil, zero value otherwise.

### GetPrivateIpAllocationModeOk

`func (o *VirtualizationNetworkInterface) GetPrivateIpAllocationModeOk() (*string, bool)`

GetPrivateIpAllocationModeOk returns a tuple with the PrivateIpAllocationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAllocationMode

`func (o *VirtualizationNetworkInterface) SetPrivateIpAllocationMode(v string)`

SetPrivateIpAllocationMode sets PrivateIpAllocationMode field to given value.

### HasPrivateIpAllocationMode

`func (o *VirtualizationNetworkInterface) HasPrivateIpAllocationMode() bool`

HasPrivateIpAllocationMode returns a boolean if a field has been set.

### GetPublicIpAllocate

`func (o *VirtualizationNetworkInterface) GetPublicIpAllocate() bool`

GetPublicIpAllocate returns the PublicIpAllocate field if non-nil, zero value otherwise.

### GetPublicIpAllocateOk

`func (o *VirtualizationNetworkInterface) GetPublicIpAllocateOk() (*bool, bool)`

GetPublicIpAllocateOk returns a tuple with the PublicIpAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAllocate

`func (o *VirtualizationNetworkInterface) SetPublicIpAllocate(v bool)`

SetPublicIpAllocate sets PublicIpAllocate field to given value.

### HasPublicIpAllocate

`func (o *VirtualizationNetworkInterface) HasPublicIpAllocate() bool`

HasPublicIpAllocate returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *VirtualizationNetworkInterface) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *VirtualizationNetworkInterface) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *VirtualizationNetworkInterface) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *VirtualizationNetworkInterface) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *VirtualizationNetworkInterface) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *VirtualizationNetworkInterface) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetStaticIpAddress

`func (o *VirtualizationNetworkInterface) GetStaticIpAddress() []VirtualizationIpAddressInfo`

GetStaticIpAddress returns the StaticIpAddress field if non-nil, zero value otherwise.

### GetStaticIpAddressOk

`func (o *VirtualizationNetworkInterface) GetStaticIpAddressOk() (*[]VirtualizationIpAddressInfo, bool)`

GetStaticIpAddressOk returns a tuple with the StaticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpAddress

`func (o *VirtualizationNetworkInterface) SetStaticIpAddress(v []VirtualizationIpAddressInfo)`

SetStaticIpAddress sets StaticIpAddress field to given value.

### HasStaticIpAddress

`func (o *VirtualizationNetworkInterface) HasStaticIpAddress() bool`

HasStaticIpAddress returns a boolean if a field has been set.

### SetStaticIpAddressNil

`func (o *VirtualizationNetworkInterface) SetStaticIpAddressNil(b bool)`

 SetStaticIpAddressNil sets the value for StaticIpAddress to be an explicit nil

### UnsetStaticIpAddress
`func (o *VirtualizationNetworkInterface) UnsetStaticIpAddress()`

UnsetStaticIpAddress ensures that no value is present for StaticIpAddress, not even an explicit nil
### GetSubnetId

`func (o *VirtualizationNetworkInterface) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VirtualizationNetworkInterface) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VirtualizationNetworkInterface) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *VirtualizationNetworkInterface) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


