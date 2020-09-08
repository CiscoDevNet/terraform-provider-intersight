# NetworkVlanPortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessVlanPortCount** | Pointer to **int64** | The number of available VLAN access ports on a Fabric Interconnect. | [optional] [readonly] 
**BorderVlanPortCount** | Pointer to **int64** | The number of available VLAN border ports on a Fabric Interconnect. | [optional] [readonly] 
**CompressedVlanPortCount** | Pointer to **string** | The number of compressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**UncompressedVlanPortCount** | Pointer to **string** | The number of uncompressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**VlanPortLimit** | Pointer to **int64** | The maximum number of VLAN ports allowed on a Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNetworkVlanPortInfo

`func NewNetworkVlanPortInfo() *NetworkVlanPortInfo`

NewNetworkVlanPortInfo instantiates a new NetworkVlanPortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlanPortInfoWithDefaults

`func NewNetworkVlanPortInfoWithDefaults() *NetworkVlanPortInfo`

NewNetworkVlanPortInfoWithDefaults instantiates a new NetworkVlanPortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessVlanPortCount

`func (o *NetworkVlanPortInfo) GetAccessVlanPortCount() int64`

GetAccessVlanPortCount returns the AccessVlanPortCount field if non-nil, zero value otherwise.

### GetAccessVlanPortCountOk

`func (o *NetworkVlanPortInfo) GetAccessVlanPortCountOk() (*int64, bool)`

GetAccessVlanPortCountOk returns a tuple with the AccessVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlanPortCount

`func (o *NetworkVlanPortInfo) SetAccessVlanPortCount(v int64)`

SetAccessVlanPortCount sets AccessVlanPortCount field to given value.

### HasAccessVlanPortCount

`func (o *NetworkVlanPortInfo) HasAccessVlanPortCount() bool`

HasAccessVlanPortCount returns a boolean if a field has been set.

### GetBorderVlanPortCount

`func (o *NetworkVlanPortInfo) GetBorderVlanPortCount() int64`

GetBorderVlanPortCount returns the BorderVlanPortCount field if non-nil, zero value otherwise.

### GetBorderVlanPortCountOk

`func (o *NetworkVlanPortInfo) GetBorderVlanPortCountOk() (*int64, bool)`

GetBorderVlanPortCountOk returns a tuple with the BorderVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderVlanPortCount

`func (o *NetworkVlanPortInfo) SetBorderVlanPortCount(v int64)`

SetBorderVlanPortCount sets BorderVlanPortCount field to given value.

### HasBorderVlanPortCount

`func (o *NetworkVlanPortInfo) HasBorderVlanPortCount() bool`

HasBorderVlanPortCount returns a boolean if a field has been set.

### GetCompressedVlanPortCount

`func (o *NetworkVlanPortInfo) GetCompressedVlanPortCount() string`

GetCompressedVlanPortCount returns the CompressedVlanPortCount field if non-nil, zero value otherwise.

### GetCompressedVlanPortCountOk

`func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountOk() (*string, bool)`

GetCompressedVlanPortCountOk returns a tuple with the CompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedVlanPortCount

`func (o *NetworkVlanPortInfo) SetCompressedVlanPortCount(v string)`

SetCompressedVlanPortCount sets CompressedVlanPortCount field to given value.

### HasCompressedVlanPortCount

`func (o *NetworkVlanPortInfo) HasCompressedVlanPortCount() bool`

HasCompressedVlanPortCount returns a boolean if a field has been set.

### GetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCount() string`

GetUncompressedVlanPortCount returns the UncompressedVlanPortCount field if non-nil, zero value otherwise.

### GetUncompressedVlanPortCountOk

`func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountOk() (*string, bool)`

GetUncompressedVlanPortCountOk returns a tuple with the UncompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfo) SetUncompressedVlanPortCount(v string)`

SetUncompressedVlanPortCount sets UncompressedVlanPortCount field to given value.

### HasUncompressedVlanPortCount

`func (o *NetworkVlanPortInfo) HasUncompressedVlanPortCount() bool`

HasUncompressedVlanPortCount returns a boolean if a field has been set.

### GetVlanPortLimit

`func (o *NetworkVlanPortInfo) GetVlanPortLimit() int64`

GetVlanPortLimit returns the VlanPortLimit field if non-nil, zero value otherwise.

### GetVlanPortLimitOk

`func (o *NetworkVlanPortInfo) GetVlanPortLimitOk() (*int64, bool)`

GetVlanPortLimitOk returns a tuple with the VlanPortLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortLimit

`func (o *NetworkVlanPortInfo) SetVlanPortLimit(v int64)`

SetVlanPortLimit sets VlanPortLimit field to given value.

### HasVlanPortLimit

`func (o *NetworkVlanPortInfo) HasVlanPortLimit() bool`

HasVlanPortLimit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *NetworkVlanPortInfo) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkVlanPortInfo) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkVlanPortInfo) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkVlanPortInfo) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVlanPortInfo) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVlanPortInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVlanPortInfo) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVlanPortInfo) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVlanPortInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVlanPortInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVlanPortInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVlanPortInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


