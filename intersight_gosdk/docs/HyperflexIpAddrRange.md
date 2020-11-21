# HyperflexIpAddrRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.IpAddrRange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.IpAddrRange"]
**EndAddr** | Pointer to **string** | The end IPv4 address of the range. | [optional] 
**Gateway** | Pointer to **string** | The default gateway for the start and end IPv4 addresses. | [optional] 
**Netmask** | Pointer to **string** | The netmask specified in dot decimal notation. The start address, end address, and gateway must all be within the network specified by this netmask. | [optional] 
**StartAddr** | Pointer to **string** | The start IPv4 address of the range. | [optional] 

## Methods

### NewHyperflexIpAddrRange

`func NewHyperflexIpAddrRange(classId string, objectType string, ) *HyperflexIpAddrRange`

NewHyperflexIpAddrRange instantiates a new HyperflexIpAddrRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexIpAddrRangeWithDefaults

`func NewHyperflexIpAddrRangeWithDefaults() *HyperflexIpAddrRange`

NewHyperflexIpAddrRangeWithDefaults instantiates a new HyperflexIpAddrRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexIpAddrRange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexIpAddrRange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexIpAddrRange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexIpAddrRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexIpAddrRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexIpAddrRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndAddr

`func (o *HyperflexIpAddrRange) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *HyperflexIpAddrRange) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *HyperflexIpAddrRange) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *HyperflexIpAddrRange) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetGateway

`func (o *HyperflexIpAddrRange) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *HyperflexIpAddrRange) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *HyperflexIpAddrRange) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *HyperflexIpAddrRange) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *HyperflexIpAddrRange) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *HyperflexIpAddrRange) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *HyperflexIpAddrRange) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *HyperflexIpAddrRange) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetStartAddr

`func (o *HyperflexIpAddrRange) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *HyperflexIpAddrRange) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *HyperflexIpAddrRange) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *HyperflexIpAddrRange) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


