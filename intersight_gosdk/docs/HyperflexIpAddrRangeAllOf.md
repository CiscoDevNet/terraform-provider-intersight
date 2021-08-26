# HyperflexIpAddrRangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.IpAddrRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.IpAddrRange"]
**EndAddr** | Pointer to **string** | The end IPv4 address of the range. | [optional] 
**Gateway** | Pointer to **string** | The default gateway for the start and end IPv4 addresses. | [optional] 
**IpAddrBlocks** | Pointer to [**[]CommIpV4AddressBlock**](CommIpV4AddressBlock.md) |  | [optional] 
**Netmask** | Pointer to **string** | The netmask specified in dot decimal notation. The start address, end address, and gateway must all be within the network specified by this netmask. | [optional] 
**StartAddr** | Pointer to **string** | The start IPv4 address of the range. | [optional] 

## Methods

### NewHyperflexIpAddrRangeAllOf

`func NewHyperflexIpAddrRangeAllOf(classId string, objectType string, ) *HyperflexIpAddrRangeAllOf`

NewHyperflexIpAddrRangeAllOf instantiates a new HyperflexIpAddrRangeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexIpAddrRangeAllOfWithDefaults

`func NewHyperflexIpAddrRangeAllOfWithDefaults() *HyperflexIpAddrRangeAllOf`

NewHyperflexIpAddrRangeAllOfWithDefaults instantiates a new HyperflexIpAddrRangeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexIpAddrRangeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexIpAddrRangeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexIpAddrRangeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexIpAddrRangeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexIpAddrRangeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexIpAddrRangeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddr

`func (o *HyperflexIpAddrRangeAllOf) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *HyperflexIpAddrRangeAllOf) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *HyperflexIpAddrRangeAllOf) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *HyperflexIpAddrRangeAllOf) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetGateway

`func (o *HyperflexIpAddrRangeAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *HyperflexIpAddrRangeAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *HyperflexIpAddrRangeAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *HyperflexIpAddrRangeAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddrBlocks

`func (o *HyperflexIpAddrRangeAllOf) GetIpAddrBlocks() []CommIpV4AddressBlock`

GetIpAddrBlocks returns the IpAddrBlocks field if non-nil, zero value otherwise.

### GetIpAddrBlocksOk

`func (o *HyperflexIpAddrRangeAllOf) GetIpAddrBlocksOk() (*[]CommIpV4AddressBlock, bool)`

GetIpAddrBlocksOk returns a tuple with the IpAddrBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddrBlocks

`func (o *HyperflexIpAddrRangeAllOf) SetIpAddrBlocks(v []CommIpV4AddressBlock)`

SetIpAddrBlocks sets IpAddrBlocks field to given value.

### HasIpAddrBlocks

`func (o *HyperflexIpAddrRangeAllOf) HasIpAddrBlocks() bool`

HasIpAddrBlocks returns a boolean if a field has been set.

### SetIpAddrBlocksNil

`func (o *HyperflexIpAddrRangeAllOf) SetIpAddrBlocksNil(b bool)`

 SetIpAddrBlocksNil sets the value for IpAddrBlocks to be an explicit nil

### UnsetIpAddrBlocks
`func (o *HyperflexIpAddrRangeAllOf) UnsetIpAddrBlocks()`

UnsetIpAddrBlocks ensures that no value is present for IpAddrBlocks, not even an explicit nil
### GetNetmask

`func (o *HyperflexIpAddrRangeAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *HyperflexIpAddrRangeAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *HyperflexIpAddrRangeAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *HyperflexIpAddrRangeAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetStartAddr

`func (o *HyperflexIpAddrRangeAllOf) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *HyperflexIpAddrRangeAllOf) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *HyperflexIpAddrRangeAllOf) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *HyperflexIpAddrRangeAllOf) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


