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
**MacAddress** | Pointer to **string** | Virtual machine network mac address. | [optional] 
**Name** | Pointer to **string** | Name of the network interface. This may be different from guest operating assigned. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


