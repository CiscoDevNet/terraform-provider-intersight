# FcpoolPool

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

### NewFcpoolPool

`func NewFcpoolPool(classId string, objectType string, ) *FcpoolPool`

NewFcpoolPool instantiates a new FcpoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolPoolWithDefaults

`func NewFcpoolPoolWithDefaults() *FcpoolPool`

NewFcpoolPoolWithDefaults instantiates a new FcpoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdBlocks

`func (o *FcpoolPool) GetIdBlocks() []FcpoolBlock`

GetIdBlocks returns the IdBlocks field if non-nil, zero value otherwise.

### GetIdBlocksOk

`func (o *FcpoolPool) GetIdBlocksOk() (*[]FcpoolBlock, bool)`

GetIdBlocksOk returns a tuple with the IdBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdBlocks

`func (o *FcpoolPool) SetIdBlocks(v []FcpoolBlock)`

SetIdBlocks sets IdBlocks field to given value.

### HasIdBlocks

`func (o *FcpoolPool) HasIdBlocks() bool`

HasIdBlocks returns a boolean if a field has been set.

### SetIdBlocksNil

`func (o *FcpoolPool) SetIdBlocksNil(b bool)`

 SetIdBlocksNil sets the value for IdBlocks to be an explicit nil

### UnsetIdBlocks
`func (o *FcpoolPool) UnsetIdBlocks()`

UnsetIdBlocks ensures that no value is present for IdBlocks, not even an explicit nil
### GetPoolPurpose

`func (o *FcpoolPool) GetPoolPurpose() string`

GetPoolPurpose returns the PoolPurpose field if non-nil, zero value otherwise.

### GetPoolPurposeOk

`func (o *FcpoolPool) GetPoolPurposeOk() (*string, bool)`

GetPoolPurposeOk returns a tuple with the PoolPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPurpose

`func (o *FcpoolPool) SetPoolPurpose(v string)`

SetPoolPurpose sets PoolPurpose field to given value.

### HasPoolPurpose

`func (o *FcpoolPool) HasPoolPurpose() bool`

HasPoolPurpose returns a boolean if a field has been set.

### GetBlockHeads

`func (o *FcpoolPool) GetBlockHeads() []FcpoolFcBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *FcpoolPool) GetBlockHeadsOk() (*[]FcpoolFcBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *FcpoolPool) SetBlockHeads(v []FcpoolFcBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *FcpoolPool) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *FcpoolPool) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *FcpoolPool) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *FcpoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FcpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FcpoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FcpoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


