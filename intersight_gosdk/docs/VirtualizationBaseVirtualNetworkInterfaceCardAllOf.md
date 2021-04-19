# VirtualizationBaseVirtualNetworkInterfaceCardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.HxapVirtualMachineNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "hyperflex.HxapVirtualMachineNetworkInterface"]
**AdapterType** | Pointer to **string** | Type or model of the virtual network interface card. * &#x60;Unknown&#x60; - The type of the network adaptor type is unknown. * &#x60;E1000&#x60; - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * &#x60;SRIOV&#x60; - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * &#x60;VMXNET2&#x60; - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * &#x60;VMXNET3&#x60; - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * &#x60;E1000E&#x60; - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * &#x60;NE2K_PCI&#x60; - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip&#39;s DMA logic will use to store received packets or to get received packets. * &#x60;PCnet&#x60; - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * &#x60;RTL8139&#x60; - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * &#x60;VirtIO&#x60; - VirtIO is a standardized interface which allows virtual machines access to simplified \&quot;virtual\&quot; devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \&quot;emulated\&quot; devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * &#x60;&#x60; - Default network adaptor type supported by the hypervisor. | [optional] [default to "Unknown"]
**MacAddress** | Pointer to **string** | MAC address assigned to the virtual network interface card. | [optional] 
**Name** | Pointer to **string** | Name of the virtual network interface card. | [optional] 

## Methods

### NewVirtualizationBaseVirtualNetworkInterfaceCardAllOf

`func NewVirtualizationBaseVirtualNetworkInterfaceCardAllOf(classId string, objectType string, ) *VirtualizationBaseVirtualNetworkInterfaceCardAllOf`

NewVirtualizationBaseVirtualNetworkInterfaceCardAllOf instantiates a new VirtualizationBaseVirtualNetworkInterfaceCardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualNetworkInterfaceCardAllOfWithDefaults

`func NewVirtualizationBaseVirtualNetworkInterfaceCardAllOfWithDefaults() *VirtualizationBaseVirtualNetworkInterfaceCardAllOf`

NewVirtualizationBaseVirtualNetworkInterfaceCardAllOfWithDefaults instantiates a new VirtualizationBaseVirtualNetworkInterfaceCardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterType

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetAdapterType() string`

GetAdapterType returns the AdapterType field if non-nil, zero value otherwise.

### GetAdapterTypeOk

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetAdapterTypeOk() (*string, bool)`

GetAdapterTypeOk returns a tuple with the AdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterType

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetAdapterType(v string)`

SetAdapterType sets AdapterType field to given value.

### HasAdapterType

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasAdapterType() bool`

HasAdapterType returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


