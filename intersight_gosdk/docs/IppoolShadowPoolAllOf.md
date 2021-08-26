# IppoolShadowPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.ShadowPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.ShadowPool"]
**IpV4Blocks** | Pointer to [**[]IppoolIpV4Block**](IppoolIpV4Block.md) |  | [optional] 
**IpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IpV6Blocks** | Pointer to [**[]IppoolIpV6Block**](IppoolIpV6Block.md) |  | [optional] 
**IpV6Config** | Pointer to [**NullableIppoolIpV6Config**](IppoolIpV6Config.md) |  | [optional] 
**V4Assigned** | Pointer to **int64** | Number of IPv4 addresses currently in use. | [optional] [readonly] 
**V4Size** | Pointer to **int64** | Number of IPv4 addresses in this pool. | [optional] [readonly] 
**V6Assigned** | Pointer to **int64** | Number of IPv6 addresses currently in use. | [optional] [readonly] 
**V6Size** | Pointer to **int64** | Number of IPv6 addresses in this pool. | [optional] [readonly] 
**IpBlockHeads** | Pointer to [**[]IppoolShadowBlockRelationship**](IppoolShadowBlockRelationship.md) | An array of relationships to ippoolShadowBlock resources. | [optional] [readonly] 
**Pool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolShadowPoolAllOf

`func NewIppoolShadowPoolAllOf(classId string, objectType string, ) *IppoolShadowPoolAllOf`

NewIppoolShadowPoolAllOf instantiates a new IppoolShadowPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolShadowPoolAllOfWithDefaults

`func NewIppoolShadowPoolAllOfWithDefaults() *IppoolShadowPoolAllOf`

NewIppoolShadowPoolAllOfWithDefaults instantiates a new IppoolShadowPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolShadowPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolShadowPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolShadowPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolShadowPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolShadowPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolShadowPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV4Blocks

`func (o *IppoolShadowPoolAllOf) GetIpV4Blocks() []IppoolIpV4Block`

GetIpV4Blocks returns the IpV4Blocks field if non-nil, zero value otherwise.

### GetIpV4BlocksOk

`func (o *IppoolShadowPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpV4Block, bool)`

GetIpV4BlocksOk returns a tuple with the IpV4Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Blocks

`func (o *IppoolShadowPoolAllOf) SetIpV4Blocks(v []IppoolIpV4Block)`

SetIpV4Blocks sets IpV4Blocks field to given value.

### HasIpV4Blocks

`func (o *IppoolShadowPoolAllOf) HasIpV4Blocks() bool`

HasIpV4Blocks returns a boolean if a field has been set.

### SetIpV4BlocksNil

`func (o *IppoolShadowPoolAllOf) SetIpV4BlocksNil(b bool)`

 SetIpV4BlocksNil sets the value for IpV4Blocks to be an explicit nil

### UnsetIpV4Blocks
`func (o *IppoolShadowPoolAllOf) UnsetIpV4Blocks()`

UnsetIpV4Blocks ensures that no value is present for IpV4Blocks, not even an explicit nil
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

### SetIpV4ConfigNil

`func (o *IppoolShadowPoolAllOf) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *IppoolShadowPoolAllOf) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
### GetIpV6Blocks

`func (o *IppoolShadowPoolAllOf) GetIpV6Blocks() []IppoolIpV6Block`

GetIpV6Blocks returns the IpV6Blocks field if non-nil, zero value otherwise.

### GetIpV6BlocksOk

`func (o *IppoolShadowPoolAllOf) GetIpV6BlocksOk() (*[]IppoolIpV6Block, bool)`

GetIpV6BlocksOk returns a tuple with the IpV6Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Blocks

`func (o *IppoolShadowPoolAllOf) SetIpV6Blocks(v []IppoolIpV6Block)`

SetIpV6Blocks sets IpV6Blocks field to given value.

### HasIpV6Blocks

`func (o *IppoolShadowPoolAllOf) HasIpV6Blocks() bool`

HasIpV6Blocks returns a boolean if a field has been set.

### SetIpV6BlocksNil

`func (o *IppoolShadowPoolAllOf) SetIpV6BlocksNil(b bool)`

 SetIpV6BlocksNil sets the value for IpV6Blocks to be an explicit nil

### UnsetIpV6Blocks
`func (o *IppoolShadowPoolAllOf) UnsetIpV6Blocks()`

UnsetIpV6Blocks ensures that no value is present for IpV6Blocks, not even an explicit nil
### GetIpV6Config

`func (o *IppoolShadowPoolAllOf) GetIpV6Config() IppoolIpV6Config`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *IppoolShadowPoolAllOf) GetIpV6ConfigOk() (*IppoolIpV6Config, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *IppoolShadowPoolAllOf) SetIpV6Config(v IppoolIpV6Config)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *IppoolShadowPoolAllOf) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *IppoolShadowPoolAllOf) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *IppoolShadowPoolAllOf) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
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

### GetIpBlockHeads

`func (o *IppoolShadowPoolAllOf) GetIpBlockHeads() []IppoolShadowBlockRelationship`

GetIpBlockHeads returns the IpBlockHeads field if non-nil, zero value otherwise.

### GetIpBlockHeadsOk

`func (o *IppoolShadowPoolAllOf) GetIpBlockHeadsOk() (*[]IppoolShadowBlockRelationship, bool)`

GetIpBlockHeadsOk returns a tuple with the IpBlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlockHeads

`func (o *IppoolShadowPoolAllOf) SetIpBlockHeads(v []IppoolShadowBlockRelationship)`

SetIpBlockHeads sets IpBlockHeads field to given value.

### HasIpBlockHeads

`func (o *IppoolShadowPoolAllOf) HasIpBlockHeads() bool`

HasIpBlockHeads returns a boolean if a field has been set.

### SetIpBlockHeadsNil

`func (o *IppoolShadowPoolAllOf) SetIpBlockHeadsNil(b bool)`

 SetIpBlockHeadsNil sets the value for IpBlockHeads to be an explicit nil

### UnsetIpBlockHeads
`func (o *IppoolShadowPoolAllOf) UnsetIpBlockHeads()`

UnsetIpBlockHeads ensures that no value is present for IpBlockHeads, not even an explicit nil
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


