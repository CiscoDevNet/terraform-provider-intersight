# HciEsxiVmNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.EsxiVmNic"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.EsxiVmNic"]
**AdapterType** | Pointer to **string** | The adapter type of the NIC. Possible values are &#39;E1000&#39;, &#39;E1000E&#39;, &#39;VMXNET&#39;, &#39;VMXNET2&#39;, &#39;VMXNET3&#39;, &#39;PCNET32&#39;, &#39;SRIOV&#39;. | [optional] [readonly] 
**IpAddresses** | Pointer to [**[]HciIpAddress**](HciIpAddress.md) |  | [optional] 
**IsConnected** | Pointer to **bool** | Indicates if the NIC is connected. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address of the NIC. | [optional] [readonly] 
**NicExtId** | Pointer to **string** | The unique identifier of the NIC. | [optional] [readonly] 
**PortgroupName** | Pointer to **string** | The name of the port group. | [optional] [readonly] 
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vm** | Pointer to [**NullableHciEsxiVmRelationship**](HciEsxiVmRelationship.md) |  | [optional] 

## Methods

### NewHciEsxiVmNic

`func NewHciEsxiVmNic(classId string, objectType string, ) *HciEsxiVmNic`

NewHciEsxiVmNic instantiates a new HciEsxiVmNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEsxiVmNicWithDefaults

`func NewHciEsxiVmNicWithDefaults() *HciEsxiVmNic`

NewHciEsxiVmNicWithDefaults instantiates a new HciEsxiVmNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEsxiVmNic) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEsxiVmNic) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEsxiVmNic) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEsxiVmNic) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEsxiVmNic) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEsxiVmNic) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterType

`func (o *HciEsxiVmNic) GetAdapterType() string`

GetAdapterType returns the AdapterType field if non-nil, zero value otherwise.

### GetAdapterTypeOk

`func (o *HciEsxiVmNic) GetAdapterTypeOk() (*string, bool)`

GetAdapterTypeOk returns a tuple with the AdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterType

`func (o *HciEsxiVmNic) SetAdapterType(v string)`

SetAdapterType sets AdapterType field to given value.

### HasAdapterType

`func (o *HciEsxiVmNic) HasAdapterType() bool`

HasAdapterType returns a boolean if a field has been set.

### GetIpAddresses

`func (o *HciEsxiVmNic) GetIpAddresses() []HciIpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *HciEsxiVmNic) GetIpAddressesOk() (*[]HciIpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *HciEsxiVmNic) SetIpAddresses(v []HciIpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *HciEsxiVmNic) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *HciEsxiVmNic) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *HciEsxiVmNic) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetIsConnected

`func (o *HciEsxiVmNic) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *HciEsxiVmNic) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *HciEsxiVmNic) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *HciEsxiVmNic) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetMacAddress

`func (o *HciEsxiVmNic) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HciEsxiVmNic) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HciEsxiVmNic) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HciEsxiVmNic) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNicExtId

`func (o *HciEsxiVmNic) GetNicExtId() string`

GetNicExtId returns the NicExtId field if non-nil, zero value otherwise.

### GetNicExtIdOk

`func (o *HciEsxiVmNic) GetNicExtIdOk() (*string, bool)`

GetNicExtIdOk returns a tuple with the NicExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicExtId

`func (o *HciEsxiVmNic) SetNicExtId(v string)`

SetNicExtId sets NicExtId field to given value.

### HasNicExtId

`func (o *HciEsxiVmNic) HasNicExtId() bool`

HasNicExtId returns a boolean if a field has been set.

### GetPortgroupName

`func (o *HciEsxiVmNic) GetPortgroupName() string`

GetPortgroupName returns the PortgroupName field if non-nil, zero value otherwise.

### GetPortgroupNameOk

`func (o *HciEsxiVmNic) GetPortgroupNameOk() (*string, bool)`

GetPortgroupNameOk returns a tuple with the PortgroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortgroupName

`func (o *HciEsxiVmNic) SetPortgroupName(v string)`

SetPortgroupName sets PortgroupName field to given value.

### HasPortgroupName

`func (o *HciEsxiVmNic) HasPortgroupName() bool`

HasPortgroupName returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciEsxiVmNic) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciEsxiVmNic) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciEsxiVmNic) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciEsxiVmNic) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HciEsxiVmNic) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciEsxiVmNic) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciEsxiVmNic) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciEsxiVmNic) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciEsxiVmNic) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciEsxiVmNic) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVm

`func (o *HciEsxiVmNic) GetVm() HciEsxiVmRelationship`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HciEsxiVmNic) GetVmOk() (*HciEsxiVmRelationship, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HciEsxiVmNic) SetVm(v HciEsxiVmRelationship)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HciEsxiVmNic) HasVm() bool`

HasVm returns a boolean if a field has been set.

### SetVmNil

`func (o *HciEsxiVmNic) SetVmNil(b bool)`

 SetVmNil sets the value for Vm to be an explicit nil

### UnsetVm
`func (o *HciEsxiVmNic) UnsetVm()`

UnsetVm ensures that no value is present for Vm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


