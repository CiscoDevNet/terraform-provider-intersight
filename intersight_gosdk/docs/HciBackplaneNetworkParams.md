# HciBackplaneNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.BackplaneNetworkParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.BackplaneNetworkParams"]
**IsSegmentationEnabled** | Pointer to **bool** | Is segmentation enabled on the backplane. | [optional] [readonly] 
**Netmask** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**Subnet** | Pointer to [**NullableHciIpAddress**](HciIpAddress.md) |  | [optional] 
**VlanTag** | Pointer to **int32** | The vlan tag of this backplane. | [optional] [readonly] 

## Methods

### NewHciBackplaneNetworkParams

`func NewHciBackplaneNetworkParams(classId string, objectType string, ) *HciBackplaneNetworkParams`

NewHciBackplaneNetworkParams instantiates a new HciBackplaneNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciBackplaneNetworkParamsWithDefaults

`func NewHciBackplaneNetworkParamsWithDefaults() *HciBackplaneNetworkParams`

NewHciBackplaneNetworkParamsWithDefaults instantiates a new HciBackplaneNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciBackplaneNetworkParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciBackplaneNetworkParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciBackplaneNetworkParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciBackplaneNetworkParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciBackplaneNetworkParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciBackplaneNetworkParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsSegmentationEnabled

`func (o *HciBackplaneNetworkParams) GetIsSegmentationEnabled() bool`

GetIsSegmentationEnabled returns the IsSegmentationEnabled field if non-nil, zero value otherwise.

### GetIsSegmentationEnabledOk

`func (o *HciBackplaneNetworkParams) GetIsSegmentationEnabledOk() (*bool, bool)`

GetIsSegmentationEnabledOk returns a tuple with the IsSegmentationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSegmentationEnabled

`func (o *HciBackplaneNetworkParams) SetIsSegmentationEnabled(v bool)`

SetIsSegmentationEnabled sets IsSegmentationEnabled field to given value.

### HasIsSegmentationEnabled

`func (o *HciBackplaneNetworkParams) HasIsSegmentationEnabled() bool`

HasIsSegmentationEnabled returns a boolean if a field has been set.

### GetNetmask

`func (o *HciBackplaneNetworkParams) GetNetmask() HciIpAddress`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *HciBackplaneNetworkParams) GetNetmaskOk() (*HciIpAddress, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *HciBackplaneNetworkParams) SetNetmask(v HciIpAddress)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *HciBackplaneNetworkParams) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *HciBackplaneNetworkParams) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *HciBackplaneNetworkParams) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetSubnet

`func (o *HciBackplaneNetworkParams) GetSubnet() HciIpAddress`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *HciBackplaneNetworkParams) GetSubnetOk() (*HciIpAddress, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *HciBackplaneNetworkParams) SetSubnet(v HciIpAddress)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *HciBackplaneNetworkParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *HciBackplaneNetworkParams) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *HciBackplaneNetworkParams) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVlanTag

`func (o *HciBackplaneNetworkParams) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *HciBackplaneNetworkParams) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *HciBackplaneNetworkParams) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *HciBackplaneNetworkParams) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


