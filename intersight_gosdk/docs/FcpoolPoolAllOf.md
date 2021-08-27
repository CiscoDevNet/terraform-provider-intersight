# FcpoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Pool"]
**IdBlocks** | Pointer to [**[]FcpoolBlock**](FcpoolBlock.md) |  | [optional] 
**PoolPurpose** | Pointer to **string** | Purpose of this WWN pool. | [optional] 
**BlockHeads** | Pointer to [**[]FcpoolFcBlockRelationship**](FcpoolFcBlockRelationship.md) | An array of relationships to fcpoolFcBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFcpoolPoolAllOf

`func NewFcpoolPoolAllOf(classId string, objectType string, ) *FcpoolPoolAllOf`

NewFcpoolPoolAllOf instantiates a new FcpoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolPoolAllOfWithDefaults

`func NewFcpoolPoolAllOfWithDefaults() *FcpoolPoolAllOf`

NewFcpoolPoolAllOfWithDefaults instantiates a new FcpoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdBlocks

`func (o *FcpoolPoolAllOf) GetIdBlocks() []FcpoolBlock`

GetIdBlocks returns the IdBlocks field if non-nil, zero value otherwise.

### GetIdBlocksOk

`func (o *FcpoolPoolAllOf) GetIdBlocksOk() (*[]FcpoolBlock, bool)`

GetIdBlocksOk returns a tuple with the IdBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdBlocks

`func (o *FcpoolPoolAllOf) SetIdBlocks(v []FcpoolBlock)`

SetIdBlocks sets IdBlocks field to given value.

### HasIdBlocks

`func (o *FcpoolPoolAllOf) HasIdBlocks() bool`

HasIdBlocks returns a boolean if a field has been set.

### SetIdBlocksNil

`func (o *FcpoolPoolAllOf) SetIdBlocksNil(b bool)`

 SetIdBlocksNil sets the value for IdBlocks to be an explicit nil

### UnsetIdBlocks
`func (o *FcpoolPoolAllOf) UnsetIdBlocks()`

UnsetIdBlocks ensures that no value is present for IdBlocks, not even an explicit nil
### GetPoolPurpose

`func (o *FcpoolPoolAllOf) GetPoolPurpose() string`

GetPoolPurpose returns the PoolPurpose field if non-nil, zero value otherwise.

### GetPoolPurposeOk

`func (o *FcpoolPoolAllOf) GetPoolPurposeOk() (*string, bool)`

GetPoolPurposeOk returns a tuple with the PoolPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPurpose

`func (o *FcpoolPoolAllOf) SetPoolPurpose(v string)`

SetPoolPurpose sets PoolPurpose field to given value.

### HasPoolPurpose

`func (o *FcpoolPoolAllOf) HasPoolPurpose() bool`

HasPoolPurpose returns a boolean if a field has been set.

### GetBlockHeads

`func (o *FcpoolPoolAllOf) GetBlockHeads() []FcpoolFcBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *FcpoolPoolAllOf) GetBlockHeadsOk() (*[]FcpoolFcBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *FcpoolPoolAllOf) SetBlockHeads(v []FcpoolFcBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *FcpoolPoolAllOf) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *FcpoolPoolAllOf) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *FcpoolPoolAllOf) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *FcpoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FcpoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FcpoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FcpoolPoolAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


