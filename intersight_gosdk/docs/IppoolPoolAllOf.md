# IppoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpV4Blocks** | Pointer to [**[]IppoolIpBlock**](ippool.IpBlock.md) |  | [optional] 
**IpV4Config** | Pointer to [**IppoolIpV4Config**](ippool.IpV4Config.md) |  | [optional] 
**V4Assigned** | Pointer to **int64** | Number of IPv4 addresses currently in use. | [optional] [readonly] 
**V4Size** | Pointer to **int64** | Number of IPv4 addresses in this pool. | [optional] [readonly] 
**V6Assigned** | Pointer to **int64** | Number of IPv6 addresses currently in use. | [optional] [readonly] 
**V6Size** | Pointer to **int64** | Number of IPv6 addresses in this pool. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ShadowPools** | Pointer to [**[]IppoolShadowPoolRelationship**](ippool.ShadowPool.Relationship.md) | An array of relationships to ippoolShadowPool resources. | [optional] [readonly] 

## Methods

### NewIppoolPoolAllOf

`func NewIppoolPoolAllOf() *IppoolPoolAllOf`

NewIppoolPoolAllOf instantiates a new IppoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolAllOfWithDefaults

`func NewIppoolPoolAllOfWithDefaults() *IppoolPoolAllOf`

NewIppoolPoolAllOfWithDefaults instantiates a new IppoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV4Blocks

`func (o *IppoolPoolAllOf) GetIpV4Blocks() []IppoolIpBlock`

GetIpV4Blocks returns the IpV4Blocks field if non-nil, zero value otherwise.

### GetIpV4BlocksOk

`func (o *IppoolPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpBlock, bool)`

GetIpV4BlocksOk returns a tuple with the IpV4Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Blocks

`func (o *IppoolPoolAllOf) SetIpV4Blocks(v []IppoolIpBlock)`

SetIpV4Blocks sets IpV4Blocks field to given value.

### HasIpV4Blocks

`func (o *IppoolPoolAllOf) HasIpV4Blocks() bool`

HasIpV4Blocks returns a boolean if a field has been set.

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


