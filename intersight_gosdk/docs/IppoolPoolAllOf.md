# IppoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.Pool"]
**IpV4Blocks** | Pointer to [**[]IppoolIpV4Block**](IppoolIpV4Block.md) |  | [optional] 
**IpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IpV6Blocks** | Pointer to [**[]IppoolIpV6Block**](IppoolIpV6Block.md) |  | [optional] 
**IpV6Config** | Pointer to [**NullableIppoolIpV6Config**](IppoolIpV6Config.md) |  | [optional] 
**V4Assigned** | Pointer to **int64** | Number of IPv4 addresses currently in use. | [optional] [readonly] 
**V4Size** | Pointer to **int64** | Number of IPv4 addresses in this pool. | [optional] [readonly] 
**V6Assigned** | Pointer to **int64** | Number of IPv6 addresses currently in use. | [optional] [readonly] 
**V6Size** | Pointer to **int64** | Number of IPv6 addresses in this pool. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**ShadowPools** | Pointer to [**[]IppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) | An array of relationships to ippoolShadowPool resources. | [optional] [readonly] 

## Methods

### NewIppoolPoolAllOf

`func NewIppoolPoolAllOf(classId string, objectType string, ) *IppoolPoolAllOf`

NewIppoolPoolAllOf instantiates a new IppoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolAllOfWithDefaults

`func NewIppoolPoolAllOfWithDefaults() *IppoolPoolAllOf`

NewIppoolPoolAllOfWithDefaults instantiates a new IppoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV4Blocks

`func (o *IppoolPoolAllOf) GetIpV4Blocks() []IppoolIpV4Block`

GetIpV4Blocks returns the IpV4Blocks field if non-nil, zero value otherwise.

### GetIpV4BlocksOk

`func (o *IppoolPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpV4Block, bool)`

GetIpV4BlocksOk returns a tuple with the IpV4Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Blocks

`func (o *IppoolPoolAllOf) SetIpV4Blocks(v []IppoolIpV4Block)`

SetIpV4Blocks sets IpV4Blocks field to given value.

### HasIpV4Blocks

`func (o *IppoolPoolAllOf) HasIpV4Blocks() bool`

HasIpV4Blocks returns a boolean if a field has been set.

### SetIpV4BlocksNil

`func (o *IppoolPoolAllOf) SetIpV4BlocksNil(b bool)`

 SetIpV4BlocksNil sets the value for IpV4Blocks to be an explicit nil

### UnsetIpV4Blocks
`func (o *IppoolPoolAllOf) UnsetIpV4Blocks()`

UnsetIpV4Blocks ensures that no value is present for IpV4Blocks, not even an explicit nil
### GetIpV4Config

`func (o *IppoolPoolAllOf) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolPoolAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolPoolAllOf) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolPoolAllOf) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### SetIpV4ConfigNil

`func (o *IppoolPoolAllOf) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *IppoolPoolAllOf) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
### GetIpV6Blocks

`func (o *IppoolPoolAllOf) GetIpV6Blocks() []IppoolIpV6Block`

GetIpV6Blocks returns the IpV6Blocks field if non-nil, zero value otherwise.

### GetIpV6BlocksOk

`func (o *IppoolPoolAllOf) GetIpV6BlocksOk() (*[]IppoolIpV6Block, bool)`

GetIpV6BlocksOk returns a tuple with the IpV6Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Blocks

`func (o *IppoolPoolAllOf) SetIpV6Blocks(v []IppoolIpV6Block)`

SetIpV6Blocks sets IpV6Blocks field to given value.

### HasIpV6Blocks

`func (o *IppoolPoolAllOf) HasIpV6Blocks() bool`

HasIpV6Blocks returns a boolean if a field has been set.

### SetIpV6BlocksNil

`func (o *IppoolPoolAllOf) SetIpV6BlocksNil(b bool)`

 SetIpV6BlocksNil sets the value for IpV6Blocks to be an explicit nil

### UnsetIpV6Blocks
`func (o *IppoolPoolAllOf) UnsetIpV6Blocks()`

UnsetIpV6Blocks ensures that no value is present for IpV6Blocks, not even an explicit nil
### GetIpV6Config

`func (o *IppoolPoolAllOf) GetIpV6Config() IppoolIpV6Config`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *IppoolPoolAllOf) GetIpV6ConfigOk() (*IppoolIpV6Config, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *IppoolPoolAllOf) SetIpV6Config(v IppoolIpV6Config)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *IppoolPoolAllOf) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *IppoolPoolAllOf) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *IppoolPoolAllOf) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
### GetV4Assigned

`func (o *IppoolPoolAllOf) GetV4Assigned() int64`

GetV4Assigned returns the V4Assigned field if non-nil, zero value otherwise.

### GetV4AssignedOk

`func (o *IppoolPoolAllOf) GetV4AssignedOk() (*int64, bool)`

GetV4AssignedOk returns a tuple with the V4Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Assigned

`func (o *IppoolPoolAllOf) SetV4Assigned(v int64)`

SetV4Assigned sets V4Assigned field to given value.

### HasV4Assigned

`func (o *IppoolPoolAllOf) HasV4Assigned() bool`

HasV4Assigned returns a boolean if a field has been set.

### GetV4Size

`func (o *IppoolPoolAllOf) GetV4Size() int64`

GetV4Size returns the V4Size field if non-nil, zero value otherwise.

### GetV4SizeOk

`func (o *IppoolPoolAllOf) GetV4SizeOk() (*int64, bool)`

GetV4SizeOk returns a tuple with the V4Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Size

`func (o *IppoolPoolAllOf) SetV4Size(v int64)`

SetV4Size sets V4Size field to given value.

### HasV4Size

`func (o *IppoolPoolAllOf) HasV4Size() bool`

HasV4Size returns a boolean if a field has been set.

### GetV6Assigned

`func (o *IppoolPoolAllOf) GetV6Assigned() int64`

GetV6Assigned returns the V6Assigned field if non-nil, zero value otherwise.

### GetV6AssignedOk

`func (o *IppoolPoolAllOf) GetV6AssignedOk() (*int64, bool)`

GetV6AssignedOk returns a tuple with the V6Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Assigned

`func (o *IppoolPoolAllOf) SetV6Assigned(v int64)`

SetV6Assigned sets V6Assigned field to given value.

### HasV6Assigned

`func (o *IppoolPoolAllOf) HasV6Assigned() bool`

HasV6Assigned returns a boolean if a field has been set.

### GetV6Size

`func (o *IppoolPoolAllOf) GetV6Size() int64`

GetV6Size returns the V6Size field if non-nil, zero value otherwise.

### GetV6SizeOk

`func (o *IppoolPoolAllOf) GetV6SizeOk() (*int64, bool)`

GetV6SizeOk returns a tuple with the V6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Size

`func (o *IppoolPoolAllOf) SetV6Size(v int64)`

SetV6Size sets V6Size field to given value.

### HasV6Size

`func (o *IppoolPoolAllOf) HasV6Size() bool`

HasV6Size returns a boolean if a field has been set.

### GetOrganization

`func (o *IppoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IppoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IppoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IppoolPoolAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetShadowPools

`func (o *IppoolPoolAllOf) GetShadowPools() []IppoolShadowPoolRelationship`

GetShadowPools returns the ShadowPools field if non-nil, zero value otherwise.

### GetShadowPoolsOk

`func (o *IppoolPoolAllOf) GetShadowPoolsOk() (*[]IppoolShadowPoolRelationship, bool)`

GetShadowPoolsOk returns a tuple with the ShadowPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowPools

`func (o *IppoolPoolAllOf) SetShadowPools(v []IppoolShadowPoolRelationship)`

SetShadowPools sets ShadowPools field to given value.

### HasShadowPools

`func (o *IppoolPoolAllOf) HasShadowPools() bool`

HasShadowPools returns a boolean if a field has been set.

### SetShadowPoolsNil

`func (o *IppoolPoolAllOf) SetShadowPoolsNil(b bool)`

 SetShadowPoolsNil sets the value for ShadowPools to be an explicit nil

### UnsetShadowPools
`func (o *IppoolPoolAllOf) UnsetShadowPools()`

UnsetShadowPools ensures that no value is present for ShadowPools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


