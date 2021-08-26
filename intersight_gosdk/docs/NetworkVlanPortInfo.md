# NetworkVlanPortInfo

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

### NewNetworkVlanPortInfo

`func NewNetworkVlanPortInfo(classId string, objectType string, ) *NetworkVlanPortInfo`

NewNetworkVlanPortInfo instantiates a new NetworkVlanPortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlanPortInfoWithDefaults

`func NewNetworkVlanPortInfoWithDefaults() *NetworkVlanPortInfo`

NewNetworkVlanPortInfoWithDefaults instantiates a new NetworkVlanPortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVlanPortInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVlanPortInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVlanPortInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVlanPortInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVlanPortInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVlanPortInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfo) GetCompressedOptimizationSetsValue() int64`

GetCompressedOptimizationSetsValue returns the CompressedOptimizationSetsValue field if non-nil, zero value otherwise.

### GetCompressedOptimizationSetsValueOk

`func (o *NetworkVlanPortInfo) GetCompressedOptimizationSetsValueOk() (*int64, bool)`

GetCompressedOptimizationSetsValueOk returns a tuple with the CompressedOptimizationSetsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfo) SetCompressedOptimizationSetsValue(v int64)`

SetCompressedOptimizationSetsValue sets CompressedOptimizationSetsValue field to given value.

### HasCompressedOptimizationSetsValue

`func (o *NetworkVlanPortInfo) HasCompressedOptimizationSetsValue() bool`

HasCompressedOptimizationSetsValue returns a boolean if a field has been set.

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

### GetCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountValue() int64`

GetCompressedVlanPortCountValue returns the CompressedVlanPortCountValue field if non-nil, zero value otherwise.

### GetCompressedVlanPortCountValueOk

`func (o *NetworkVlanPortInfo) GetCompressedVlanPortCountValueOk() (*int64, bool)`

GetCompressedVlanPortCountValueOk returns a tuple with the CompressedVlanPortCountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) SetCompressedVlanPortCountValue(v int64)`

SetCompressedVlanPortCountValue sets CompressedVlanPortCountValue field to given value.

### HasCompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) HasCompressedVlanPortCountValue() bool`

HasCompressedVlanPortCountValue returns a boolean if a field has been set.

### GetTotalVlanPortCount

`func (o *NetworkVlanPortInfo) GetTotalVlanPortCount() int64`

GetTotalVlanPortCount returns the TotalVlanPortCount field if non-nil, zero value otherwise.

### GetTotalVlanPortCountOk

`func (o *NetworkVlanPortInfo) GetTotalVlanPortCountOk() (*int64, bool)`

GetTotalVlanPortCountOk returns a tuple with the TotalVlanPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVlanPortCount

`func (o *NetworkVlanPortInfo) SetTotalVlanPortCount(v int64)`

SetTotalVlanPortCount sets TotalVlanPortCount field to given value.

### HasTotalVlanPortCount

`func (o *NetworkVlanPortInfo) HasTotalVlanPortCount() bool`

HasTotalVlanPortCount returns a boolean if a field has been set.

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

### GetUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountValue() int64`

GetUncompressedVlanPortCountValue returns the UncompressedVlanPortCountValue field if non-nil, zero value otherwise.

### GetUncompressedVlanPortCountValueOk

`func (o *NetworkVlanPortInfo) GetUncompressedVlanPortCountValueOk() (*int64, bool)`

GetUncompressedVlanPortCountValueOk returns a tuple with the UncompressedVlanPortCountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) SetUncompressedVlanPortCountValue(v int64)`

SetUncompressedVlanPortCountValue sets UncompressedVlanPortCountValue field to given value.

### HasUncompressedVlanPortCountValue

`func (o *NetworkVlanPortInfo) HasUncompressedVlanPortCountValue() bool`

HasUncompressedVlanPortCountValue returns a boolean if a field has been set.

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


