# IppoolPool

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

### NewIppoolPool

`func NewIppoolPool() *IppoolPool`

NewIppoolPool instantiates a new IppoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolWithDefaults

`func NewIppoolPoolWithDefaults() *IppoolPool`

NewIppoolPoolWithDefaults instantiates a new IppoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV4Blocks

`func (o *IppoolPool) GetIpV4Blocks() []IppoolIpBlock`

GetIpV4Blocks returns the IpV4Blocks field if non-nil, zero value otherwise.

### GetIpV4BlocksOk

`func (o *IppoolPool) GetIpV4BlocksOk() (*[]IppoolIpBlock, bool)`

GetIpV4BlocksOk returns a tuple with the IpV4Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Blocks

`func (o *IppoolPool) SetIpV4Blocks(v []IppoolIpBlock)`

SetIpV4Blocks sets IpV4Blocks field to given value.

### HasIpV4Blocks

`func (o *IppoolPool) HasIpV4Blocks() bool`

HasIpV4Blocks returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolPool) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolPool) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolPool) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolPool) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### GetV4Assigned

`func (o *IppoolPool) GetV4Assigned() int64`

GetV4Assigned returns the V4Assigned field if non-nil, zero value otherwise.

### GetV4AssignedOk

`func (o *IppoolPool) GetV4AssignedOk() (*int64, bool)`

GetV4AssignedOk returns a tuple with the V4Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Assigned

`func (o *IppoolPool) SetV4Assigned(v int64)`

SetV4Assigned sets V4Assigned field to given value.

### HasV4Assigned

`func (o *IppoolPool) HasV4Assigned() bool`

HasV4Assigned returns a boolean if a field has been set.

### GetV4Size

`func (o *IppoolPool) GetV4Size() int64`

GetV4Size returns the V4Size field if non-nil, zero value otherwise.

### GetV4SizeOk

`func (o *IppoolPool) GetV4SizeOk() (*int64, bool)`

GetV4SizeOk returns a tuple with the V4Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Size

`func (o *IppoolPool) SetV4Size(v int64)`

SetV4Size sets V4Size field to given value.

### HasV4Size

`func (o *IppoolPool) HasV4Size() bool`

HasV4Size returns a boolean if a field has been set.

### GetV6Assigned

`func (o *IppoolPool) GetV6Assigned() int64`

GetV6Assigned returns the V6Assigned field if non-nil, zero value otherwise.

### GetV6AssignedOk

`func (o *IppoolPool) GetV6AssignedOk() (*int64, bool)`

GetV6AssignedOk returns a tuple with the V6Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Assigned

`func (o *IppoolPool) SetV6Assigned(v int64)`

SetV6Assigned sets V6Assigned field to given value.

### HasV6Assigned

`func (o *IppoolPool) HasV6Assigned() bool`

HasV6Assigned returns a boolean if a field has been set.

### GetV6Size

`func (o *IppoolPool) GetV6Size() int64`

GetV6Size returns the V6Size field if non-nil, zero value otherwise.

### GetV6SizeOk

`func (o *IppoolPool) GetV6SizeOk() (*int64, bool)`

GetV6SizeOk returns a tuple with the V6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6Size

`func (o *IppoolPool) SetV6Size(v int64)`

SetV6Size sets V6Size field to given value.

### HasV6Size

`func (o *IppoolPool) HasV6Size() bool`

HasV6Size returns a boolean if a field has been set.

### GetOrganization

`func (o *IppoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IppoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IppoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IppoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetShadowPools

`func (o *IppoolPool) GetShadowPools() []IppoolShadowPoolRelationship`

GetShadowPools returns the ShadowPools field if non-nil, zero value otherwise.

### GetShadowPoolsOk

`func (o *IppoolPool) GetShadowPoolsOk() (*[]IppoolShadowPoolRelationship, bool)`

GetShadowPoolsOk returns a tuple with the ShadowPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowPools

`func (o *IppoolPool) SetShadowPools(v []IppoolShadowPoolRelationship)`

SetShadowPools sets ShadowPools field to given value.

### HasShadowPools

`func (o *IppoolPool) HasShadowPools() bool`

HasShadowPools returns a boolean if a field has been set.

### SetShadowPoolsNil

`func (o *IppoolPool) SetShadowPoolsNil(b bool)`

 SetShadowPoolsNil sets the value for ShadowPools to be an explicit nil

### UnsetShadowPools
`func (o *IppoolPool) UnsetShadowPools()`

UnsetShadowPools ensures that no value is present for ShadowPools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


