# VirtualizationNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.NetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.NetworkInterface"]
**AdaptorType** | Pointer to **string** | Virtual machine network adaptor type. * &#x60;Unknown&#x60; - The type of the network adaptor type is unknown. * &#x60;E1000&#x60; - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * &#x60;SRIOV&#x60; - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * &#x60;VMXNET2&#x60; - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * &#x60;VMXNET3&#x60; - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. | [optional] [default to "Unknown"]
**Bridge** | Pointer to **string** | Virtual machine network bridge name. | [optional] 
**ConnectAtPowerOn** | Pointer to **bool** | Connect the adaptor at virtual machine power on. | [optional] 
**DirectPathIo** | Pointer to **bool** | Enable the direct path I/O. | [optional] 
**MacAddress** | Pointer to **string** | Virtual machine network mac address. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


