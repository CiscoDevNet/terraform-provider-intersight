# IppoolShadowPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpV4Blocks** | Pointer to [**[]IppoolIpBlock**](ippool.IpBlock.md) |  | [optional] 
**IpV4Config** | Pointer to [**IppoolIpV4Config**](ippool.IpV4Config.md) |  | [optional] 
**V4Assigned** | Pointer to **int64** | Number of IPv4 addresses currently in use. | [optional] [readonly] 
**V4Size** | Pointer to **int64** | Number of IPv4 addresses in this pool. | [optional] [readonly] 
**V6Assigned** | Pointer to **int64** | Number of IPv6 addresses currently in use. | [optional] [readonly] 
**V6Size** | Pointer to **int64** | Number of IPv6 addresses in this pool. | [optional] [readonly] 
**Pool** | Pointer to [**IppoolPoolRelationship**](ippool.Pool.Relationship.md) |  | [optional] 
**V4BlockHeads** | Pointer to [**[]IppoolShadowBlockRelationship**](ippool.ShadowBlock.Relationship.md) | An array of relationships to ippoolShadowBlock resources. | [optional] [readonly] 
**Vrf** | Pointer to [**VrfVrfRelationship**](vrf.Vrf.Relationship.md) |  | [optional] 

## Methods

### NewIppoolShadowPoolAllOf

`func NewIppoolShadowPoolAllOf() *IppoolShadowPoolAllOf`

NewIppoolShadowPoolAllOf instantiates a new IppoolShadowPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolShadowPoolAllOfWithDefaults

`func NewIppoolShadowPoolAllOfWithDefaults() *IppoolShadowPoolAllOf`

NewIppoolShadowPoolAllOfWithDefaults instantiates a new IppoolShadowPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV4Blocks

`func (o *IppoolShadowPoolAllOf) GetIpV4Blocks() []IppoolIpBlock`

GetIpV4Blocks returns the IpV4Blocks field if non-nil, zero value otherwise.

### GetIpV4BlocksOk

`func (o *IppoolShadowPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpBlock, bool)`

GetIpV4BlocksOk returns a tuple with the IpV4Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Blocks

`func (o *IppoolShadowPoolAllOf) SetIpV4Blocks(v []IppoolIpBlock)`

SetIpV4Blocks sets IpV4Blocks field to given value.

### HasIpV4Blocks

`func (o *IppoolShadowPoolAllOf) HasIpV4Blocks() bool`

HasIpV4Blocks returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolShadowPoolAllOf) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolShadowPoolAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolShadowPoolAllOf) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolShadowPoolAllOf) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### GetV4Assigned

`func (o *IppoolShadowPoolAllOf) GetV4Assigned() int64`

GetV4Assigned returns the V4Assigned field if non-nil, zero value otherwise.

### GetV4AssignedOk

`func (o *IppoolShadowPoolAllOf) GetV4AssignedOk() (*int64, bool)`

GetV4AssignedOk returns a tuple with the V4Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Assigned

`func (o *IppoolShadowPoolAllOf) SetV4Assigned(v int64)`

SetV4Assigned sets V4Assigned field to given value.

### HasV4Assigned

`func (o *IppoolShadowPoolAllOf) HasV4Assigned() bool`

HasV4Assigned returns a boolean if a field has been set.

### GetV4Size

`func (o *IppoolShadowPoolAllOf) GetV4Size() int64`

GetV4Size returns the V4Size field if non-nil, zero value otherwise.

### GetV4SizeOk

`func (o *IppoolShadowPoolAllOf) GetV4SizeOk() (*int64, bool)`

GetV4SizeOk returns a tuple with the V4Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Size

`func (o *IppoolShadowPoolAllOf) SetV4Size(v int64)`

SetV4Size sets V4Size field to given value.

### HasV4Size

`func (o *IppoolShadowPoolAllOf) HasV4Size() bool`

HasV4Size returns a boolean if a field has been set.

### GetV6Assigned

`func (o *IppoolShadowPoolAllOf) GetV6Assigned() int64`

GetV6Assigned returns the V6Assigned field if non-nil, zero value otherwise.

### GetV6AssignedOk

`func (o *IppoolShadowPoolAllOf) GetV6AssignedOk() (*int64, bool)`

GetV6AssignedOk returns a tuple with the V6Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Assigned

`func (o *IppoolShadowPoolAllOf) SetV6Assigned(v int64)`

SetV6Assigned sets V6Assigned field to given value.

### HasV6Assigned

`func (o *IppoolShadowPoolAllOf) HasV6Assigned() bool`

HasV6Assigned returns a boolean if a field has been set.

### GetV6Size

`func (o *IppoolShadowPoolAllOf) GetV6Size() int64`

GetV6Size returns the V6Size field if non-nil, zero value otherwise.

### GetV6SizeOk

`func (o *IppoolShadowPoolAllOf) GetV6SizeOk() (*int64, bool)`

GetV6SizeOk returns a tuple with the V6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Size

`func (o *IppoolShadowPoolAllOf) SetV6Size(v int64)`

SetV6Size sets V6Size field to given value.

### HasV6Size

`func (o *IppoolShadowPoolAllOf) HasV6Size() bool`

HasV6Size returns a boolean if a field has been set.

### GetPool

`func (o *IppoolShadowPoolAllOf) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolShadowPoolAllOf) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolShadowPoolAllOf) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolShadowPoolAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetV4BlockHeads

`func (o *IppoolShadowPoolAllOf) GetV4BlockHeads() []IppoolShadowBlockRelationship`

GetV4BlockHeads returns the V4BlockHeads field if non-nil, zero value otherwise.

### GetV4BlockHeadsOk

`func (o *IppoolShadowPoolAllOf) GetV4BlockHeadsOk() (*[]IppoolShadowBlockRelationship, bool)`

GetV4BlockHeadsOk returns a tuple with the V4BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4BlockHeads

`func (o *IppoolShadowPoolAllOf) SetV4BlockHeads(v []IppoolShadowBlockRelationship)`

SetV4BlockHeads sets V4BlockHeads field to given value.

### HasV4BlockHeads

`func (o *IppoolShadowPoolAllOf) HasV4BlockHeads() bool`

HasV4BlockHeads returns a boolean if a field has been set.

### SetV4BlockHeadsNil

`func (o *IppoolShadowPoolAllOf) SetV4BlockHeadsNil(b bool)`

 SetV4BlockHeadsNil sets the value for V4BlockHeads to be an explicit nil

### UnsetV4BlockHeads
`func (o *IppoolShadowPoolAllOf) UnsetV4BlockHeads()`

UnsetV4BlockHeads ensures that no value is present for V4BlockHeads, not even an explicit nil
### GetVrf

`func (o *IppoolShadowPoolAllOf) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolShadowPoolAllOf) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolShadowPoolAllOf) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolShadowPoolAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


