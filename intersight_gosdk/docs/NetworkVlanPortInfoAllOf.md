# NetworkVlanPortInfoAllOf

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

### NewNetworkVlanPortInfoAllOf

`func NewNetworkVlanPortInfoAllOf() *NetworkVlanPortInfoAllOf`

NewNetworkVlanPortInfoAllOf instantiates a new NetworkVlanPortInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlanPortInfoAllOfWithDefaults

`func NewNetworkVlanPortInfoAllOfWithDefaults() *NetworkVlanPortInfoAllOf`

NewNetworkVlanPortInfoAllOfWithDefaults instantiates a new NetworkVlanPortInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) GetAccessVlanPortCount() int64`

GetAccessVlanPortCount returns the AccessVlanPortCount field if non-nil, zero value otherwise.

### GetAccessVlanPortCountOk

`func (o *NetworkVlanPortInfoAllOf) GetAccessVlanPortCountOk() (*int64, bool)`

GetAccessVlanPortCountOk returns a tuple with the AccessVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) SetAccessVlanPortCount(v int64)`

SetAccessVlanPortCount sets AccessVlanPortCount field to given value.

### HasAccessVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) HasAccessVlanPortCount() bool`

HasAccessVlanPortCount returns a boolean if a field has been set.

### GetBorderVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) GetBorderVlanPortCount() int64`

GetBorderVlanPortCount returns the BorderVlanPortCount field if non-nil, zero value otherwise.

### GetBorderVlanPortCountOk

`func (o *NetworkVlanPortInfoAllOf) GetBorderVlanPortCountOk() (*int64, bool)`

GetBorderVlanPortCountOk returns a tuple with the BorderVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) SetBorderVlanPortCount(v int64)`

SetBorderVlanPortCount sets BorderVlanPortCount field to given value.

### HasBorderVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) HasBorderVlanPortCount() bool`

HasBorderVlanPortCount returns a boolean if a field has been set.

### GetCompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCount() string`

GetCompressedVlanPortCount returns the CompressedVlanPortCount field if non-nil, zero value otherwise.

### GetCompressedVlanPortCountOk

`func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountOk() (*string, bool)`

GetCompressedVlanPortCountOk returns a tuple with the CompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) SetCompressedVlanPortCount(v string)`

SetCompressedVlanPortCount sets CompressedVlanPortCount field to given value.

### HasCompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) HasCompressedVlanPortCount() bool`

HasCompressedVlanPortCount returns a boolean if a field has been set.

### GetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCount() string`

GetUncompressedVlanPortCount returns the UncompressedVlanPortCount field if non-nil, zero value otherwise.

### GetUncompressedVlanPortCountOk

`func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountOk() (*string, bool)`

GetUncompressedVlanPortCountOk returns a tuple with the UncompressedVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) SetUncompressedVlanPortCount(v string)`

SetUncompressedVlanPortCount sets UncompressedVlanPortCount field to given value.

### HasUncompressedVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) HasUncompressedVlanPortCount() bool`

HasUncompressedVlanPortCount returns a boolean if a field has been set.

### GetVlanPortLimit

`func (o *NetworkVlanPortInfoAllOf) GetVlanPortLimit() int64`

GetVlanPortLimit returns the VlanPortLimit field if non-nil, zero value otherwise.

### GetVlanPortLimitOk

`func (o *NetworkVlanPortInfoAllOf) GetVlanPortLimitOk() (*int64, bool)`

GetVlanPortLimitOk returns a tuple with the VlanPortLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortLimit

`func (o *NetworkVlanPortInfoAllOf) SetVlanPortLimit(v int64)`

SetVlanPortLimit sets VlanPortLimit field to given value.

### HasVlanPortLimit

`func (o *NetworkVlanPortInfoAllOf) HasVlanPortLimit() bool`

HasVlanPortLimit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *NetworkVlanPortInfoAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *NetworkVlanPortInfoAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *NetworkVlanPortInfoAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *NetworkVlanPortInfoAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkVlanPortInfoAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVlanPortInfoAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVlanPortInfoAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVlanPortInfoAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkVlanPortInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVlanPortInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVlanPortInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVlanPortInfoAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


