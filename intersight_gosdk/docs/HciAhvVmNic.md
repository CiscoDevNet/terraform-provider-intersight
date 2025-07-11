# HciAhvVmNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.AhvVmNic"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.AhvVmNic"]
**IpAddresses** | Pointer to [**[]HciIpAddress**](HciIpAddress.md) |  | [optional] 
**IsConnected** | Pointer to **bool** | Indicates if the NIC is connected. | [optional] [readonly] 
**LearnedIpAddresses** | Pointer to [**[]HciIpAddress**](HciIpAddress.md) |  | [optional] 
**MacAddress** | Pointer to **string** | The MAC address of the NIC. | [optional] [readonly] 
**Model** | Pointer to **string** | The model name of emulated NIC. | [optional] [readonly] 
**NetworkFunctionNicType** | Pointer to **string** | The type of this network function NIC. | [optional] [readonly] 
**NicExtId** | Pointer to **string** | The unique identifier of the NIC. | [optional] [readonly] 
**NicType** | Pointer to **string** | The type of the NIC. Possible values are &#39;NORMAL_NIC&#39;, &#39;DIRECT_NIC&#39;, &#39;NETWORK_FUNCTION_NIC&#39;, &#39;SPAN_DESTINATION_NIC&#39;. | [optional] [readonly] 
**NumQueues** | Pointer to **int32** | The number of Tx/Rx queue pairs for this NIC. | [optional] [readonly] 
**ShouldAllowUnknownMacs** | Pointer to **bool** | Indicates whether an unknown unicast traffic is forwarded to this NIC or not, only for the NICs on the overlay subnets. | [optional] [readonly] 
**TrunkedVlans** | Pointer to **[]int32** |  | [optional] 
**VlanId** | Pointer to **[]int32** |  | [optional] 
**VlanMode** | Pointer to **string** | By default, all the virtual NICs are created in ACCESS mode, which permits only one VLAN per virtual network. TRUNKED mode allows multiple VLANs on a single VM NIC for network-aware user VMs. | [optional] [readonly] 
**VmExtId** | Pointer to **string** | The unique identifier of the VM. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vm** | Pointer to [**NullableHciAhvVmRelationship**](HciAhvVmRelationship.md) |  | [optional] 

## Methods

### NewHciAhvVmNic

`func NewHciAhvVmNic(classId string, objectType string, ) *HciAhvVmNic`

NewHciAhvVmNic instantiates a new HciAhvVmNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciAhvVmNicWithDefaults

`func NewHciAhvVmNicWithDefaults() *HciAhvVmNic`

NewHciAhvVmNicWithDefaults instantiates a new HciAhvVmNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciAhvVmNic) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciAhvVmNic) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciAhvVmNic) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciAhvVmNic) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciAhvVmNic) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciAhvVmNic) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddresses

`func (o *HciAhvVmNic) GetIpAddresses() []HciIpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *HciAhvVmNic) GetIpAddressesOk() (*[]HciIpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *HciAhvVmNic) SetIpAddresses(v []HciIpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *HciAhvVmNic) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *HciAhvVmNic) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *HciAhvVmNic) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetIsConnected

`func (o *HciAhvVmNic) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *HciAhvVmNic) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *HciAhvVmNic) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *HciAhvVmNic) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetLearnedIpAddresses

`func (o *HciAhvVmNic) GetLearnedIpAddresses() []HciIpAddress`

GetLearnedIpAddresses returns the LearnedIpAddresses field if non-nil, zero value otherwise.

### GetLearnedIpAddressesOk

`func (o *HciAhvVmNic) GetLearnedIpAddressesOk() (*[]HciIpAddress, bool)`

GetLearnedIpAddressesOk returns a tuple with the LearnedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnedIpAddresses

`func (o *HciAhvVmNic) SetLearnedIpAddresses(v []HciIpAddress)`

SetLearnedIpAddresses sets LearnedIpAddresses field to given value.

### HasLearnedIpAddresses

`func (o *HciAhvVmNic) HasLearnedIpAddresses() bool`

HasLearnedIpAddresses returns a boolean if a field has been set.

### SetLearnedIpAddressesNil

`func (o *HciAhvVmNic) SetLearnedIpAddressesNil(b bool)`

 SetLearnedIpAddressesNil sets the value for LearnedIpAddresses to be an explicit nil

### UnsetLearnedIpAddresses
`func (o *HciAhvVmNic) UnsetLearnedIpAddresses()`

UnsetLearnedIpAddresses ensures that no value is present for LearnedIpAddresses, not even an explicit nil
### GetMacAddress

`func (o *HciAhvVmNic) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HciAhvVmNic) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HciAhvVmNic) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HciAhvVmNic) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *HciAhvVmNic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HciAhvVmNic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HciAhvVmNic) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HciAhvVmNic) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetworkFunctionNicType

`func (o *HciAhvVmNic) GetNetworkFunctionNicType() string`

GetNetworkFunctionNicType returns the NetworkFunctionNicType field if non-nil, zero value otherwise.

### GetNetworkFunctionNicTypeOk

`func (o *HciAhvVmNic) GetNetworkFunctionNicTypeOk() (*string, bool)`

GetNetworkFunctionNicTypeOk returns a tuple with the NetworkFunctionNicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionNicType

`func (o *HciAhvVmNic) SetNetworkFunctionNicType(v string)`

SetNetworkFunctionNicType sets NetworkFunctionNicType field to given value.

### HasNetworkFunctionNicType

`func (o *HciAhvVmNic) HasNetworkFunctionNicType() bool`

HasNetworkFunctionNicType returns a boolean if a field has been set.

### GetNicExtId

`func (o *HciAhvVmNic) GetNicExtId() string`

GetNicExtId returns the NicExtId field if non-nil, zero value otherwise.

### GetNicExtIdOk

`func (o *HciAhvVmNic) GetNicExtIdOk() (*string, bool)`

GetNicExtIdOk returns a tuple with the NicExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicExtId

`func (o *HciAhvVmNic) SetNicExtId(v string)`

SetNicExtId sets NicExtId field to given value.

### HasNicExtId

`func (o *HciAhvVmNic) HasNicExtId() bool`

HasNicExtId returns a boolean if a field has been set.

### GetNicType

`func (o *HciAhvVmNic) GetNicType() string`

GetNicType returns the NicType field if non-nil, zero value otherwise.

### GetNicTypeOk

`func (o *HciAhvVmNic) GetNicTypeOk() (*string, bool)`

GetNicTypeOk returns a tuple with the NicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicType

`func (o *HciAhvVmNic) SetNicType(v string)`

SetNicType sets NicType field to given value.

### HasNicType

`func (o *HciAhvVmNic) HasNicType() bool`

HasNicType returns a boolean if a field has been set.

### GetNumQueues

`func (o *HciAhvVmNic) GetNumQueues() int32`

GetNumQueues returns the NumQueues field if non-nil, zero value otherwise.

### GetNumQueuesOk

`func (o *HciAhvVmNic) GetNumQueuesOk() (*int32, bool)`

GetNumQueuesOk returns a tuple with the NumQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueues

`func (o *HciAhvVmNic) SetNumQueues(v int32)`

SetNumQueues sets NumQueues field to given value.

### HasNumQueues

`func (o *HciAhvVmNic) HasNumQueues() bool`

HasNumQueues returns a boolean if a field has been set.

### GetShouldAllowUnknownMacs

`func (o *HciAhvVmNic) GetShouldAllowUnknownMacs() bool`

GetShouldAllowUnknownMacs returns the ShouldAllowUnknownMacs field if non-nil, zero value otherwise.

### GetShouldAllowUnknownMacsOk

`func (o *HciAhvVmNic) GetShouldAllowUnknownMacsOk() (*bool, bool)`

GetShouldAllowUnknownMacsOk returns a tuple with the ShouldAllowUnknownMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAllowUnknownMacs

`func (o *HciAhvVmNic) SetShouldAllowUnknownMacs(v bool)`

SetShouldAllowUnknownMacs sets ShouldAllowUnknownMacs field to given value.

### HasShouldAllowUnknownMacs

`func (o *HciAhvVmNic) HasShouldAllowUnknownMacs() bool`

HasShouldAllowUnknownMacs returns a boolean if a field has been set.

### GetTrunkedVlans

`func (o *HciAhvVmNic) GetTrunkedVlans() []int32`

GetTrunkedVlans returns the TrunkedVlans field if non-nil, zero value otherwise.

### GetTrunkedVlansOk

`func (o *HciAhvVmNic) GetTrunkedVlansOk() (*[]int32, bool)`

GetTrunkedVlansOk returns a tuple with the TrunkedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkedVlans

`func (o *HciAhvVmNic) SetTrunkedVlans(v []int32)`

SetTrunkedVlans sets TrunkedVlans field to given value.

### HasTrunkedVlans

`func (o *HciAhvVmNic) HasTrunkedVlans() bool`

HasTrunkedVlans returns a boolean if a field has been set.

### SetTrunkedVlansNil

`func (o *HciAhvVmNic) SetTrunkedVlansNil(b bool)`

 SetTrunkedVlansNil sets the value for TrunkedVlans to be an explicit nil

### UnsetTrunkedVlans
`func (o *HciAhvVmNic) UnsetTrunkedVlans()`

UnsetTrunkedVlans ensures that no value is present for TrunkedVlans, not even an explicit nil
### GetVlanId

`func (o *HciAhvVmNic) GetVlanId() []int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *HciAhvVmNic) GetVlanIdOk() (*[]int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *HciAhvVmNic) SetVlanId(v []int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *HciAhvVmNic) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *HciAhvVmNic) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *HciAhvVmNic) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVlanMode

`func (o *HciAhvVmNic) GetVlanMode() string`

GetVlanMode returns the VlanMode field if non-nil, zero value otherwise.

### GetVlanModeOk

`func (o *HciAhvVmNic) GetVlanModeOk() (*string, bool)`

GetVlanModeOk returns a tuple with the VlanMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanMode

`func (o *HciAhvVmNic) SetVlanMode(v string)`

SetVlanMode sets VlanMode field to given value.

### HasVlanMode

`func (o *HciAhvVmNic) HasVlanMode() bool`

HasVlanMode returns a boolean if a field has been set.

### GetVmExtId

`func (o *HciAhvVmNic) GetVmExtId() string`

GetVmExtId returns the VmExtId field if non-nil, zero value otherwise.

### GetVmExtIdOk

`func (o *HciAhvVmNic) GetVmExtIdOk() (*string, bool)`

GetVmExtIdOk returns a tuple with the VmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmExtId

`func (o *HciAhvVmNic) SetVmExtId(v string)`

SetVmExtId sets VmExtId field to given value.

### HasVmExtId

`func (o *HciAhvVmNic) HasVmExtId() bool`

HasVmExtId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HciAhvVmNic) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciAhvVmNic) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciAhvVmNic) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciAhvVmNic) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciAhvVmNic) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciAhvVmNic) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetVm

`func (o *HciAhvVmNic) GetVm() HciAhvVmRelationship`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HciAhvVmNic) GetVmOk() (*HciAhvVmRelationship, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HciAhvVmNic) SetVm(v HciAhvVmRelationship)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HciAhvVmNic) HasVm() bool`

HasVm returns a boolean if a field has been set.

### SetVmNil

`func (o *HciAhvVmNic) SetVmNil(b bool)`

 SetVmNil sets the value for Vm to be an explicit nil

### UnsetVm
`func (o *HciAhvVmNic) UnsetVm()`

UnsetVm ensures that no value is present for Vm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


