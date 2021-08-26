# NetworkVlanPortInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.VlanPortInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.VlanPortInfo"]
**AccessVlanPortCount** | Pointer to **int64** | The number of available VLAN access ports on a Fabric Interconnect. | [optional] [readonly] 
**BorderVlanPortCount** | Pointer to **int64** | The number of available VLAN border ports on a Fabric Interconnect. | [optional] [readonly] 
**CompressedOptimizationSetsValue** | Pointer to **int64** | The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library. | [optional] 
**CompressedVlanPortCount** | Pointer to **string** | The number of compressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**CompressedVlanPortCountValue** | Pointer to **int64** | The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. | [optional] 
**TotalVlanPortCount** | Pointer to **int64** | The total number of VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**UncompressedVlanPortCount** | Pointer to **string** | The number of uncompressed VLAN ports on a Fabric Interconnect. | [optional] [readonly] 
**UncompressedVlanPortCountValue** | Pointer to **int64** | The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library. | [optional] 
**VlanPortLimit** | Pointer to **int64** | The maximum number of VLAN ports allowed on a Fabric Interconnect. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVlanPortInfoAllOf

`func NewNetworkVlanPortInfoAllOf(classId string, objectType string, ) *NetworkVlanPortInfoAllOf`

NewNetworkVlanPortInfoAllOf instantiates a new NetworkVlanPortInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlanPortInfoAllOfWithDefaults

`func NewNetworkVlanPortInfoAllOfWithDefaults() *NetworkVlanPortInfoAllOf`

NewNetworkVlanPortInfoAllOfWithDefaults instantiates a new NetworkVlanPortInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVlanPortInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVlanPortInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVlanPortInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVlanPortInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVlanPortInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVlanPortInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfoAllOf) GetCompressedOptimizationSetsValue() int64`

GetCompressedOptimizationSetsValue returns the CompressedOptimizationSetsValue field if non-nil, zero value otherwise.

### GetCompressedOptimizationSetsValueOk

`func (o *NetworkVlanPortInfoAllOf) GetCompressedOptimizationSetsValueOk() (*int64, bool)`

GetCompressedOptimizationSetsValueOk returns a tuple with the CompressedOptimizationSetsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfoAllOf) SetCompressedOptimizationSetsValue(v int64)`

SetCompressedOptimizationSetsValue sets CompressedOptimizationSetsValue field to given value.

### HasCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfoAllOf) HasCompressedOptimizationSetsValue() bool`

HasCompressedOptimizationSetsValue returns a boolean if a field has been set.

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

### GetCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountValue() int64`

GetCompressedVlanPortCountValue returns the CompressedVlanPortCountValue field if non-nil, zero value otherwise.

### GetCompressedVlanPortCountValueOk

`func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountValueOk() (*int64, bool)`

GetCompressedVlanPortCountValueOk returns a tuple with the CompressedVlanPortCountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) SetCompressedVlanPortCountValue(v int64)`

SetCompressedVlanPortCountValue sets CompressedVlanPortCountValue field to given value.

### HasCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) HasCompressedVlanPortCountValue() bool`

HasCompressedVlanPortCountValue returns a boolean if a field has been set.

### GetTotalVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) GetTotalVlanPortCount() int64`

GetTotalVlanPortCount returns the TotalVlanPortCount field if non-nil, zero value otherwise.

### GetTotalVlanPortCountOk

`func (o *NetworkVlanPortInfoAllOf) GetTotalVlanPortCountOk() (*int64, bool)`

GetTotalVlanPortCountOk returns a tuple with the TotalVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) SetTotalVlanPortCount(v int64)`

SetTotalVlanPortCount sets TotalVlanPortCount field to given value.

### HasTotalVlanPortCount

`func (o *NetworkVlanPortInfoAllOf) HasTotalVlanPortCount() bool`

HasTotalVlanPortCount returns a boolean if a field has been set.

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

### GetUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountValue() int64`

GetUncompressedVlanPortCountValue returns the UncompressedVlanPortCountValue field if non-nil, zero value otherwise.

### GetUncompressedVlanPortCountValueOk

`func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountValueOk() (*int64, bool)`

GetUncompressedVlanPortCountValueOk returns a tuple with the UncompressedVlanPortCountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) SetUncompressedVlanPortCountValue(v int64)`

SetUncompressedVlanPortCountValue sets UncompressedVlanPortCountValue field to given value.

### HasUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfoAllOf) HasUncompressedVlanPortCountValue() bool`

HasUncompressedVlanPortCountValue returns a boolean if a field has been set.

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


