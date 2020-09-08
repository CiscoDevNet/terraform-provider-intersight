# MacpoolPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacBlocks** | Pointer to [**[]MacpoolBlock**](macpool.Block.md) |  | [optional] 
**BlockHeads** | Pointer to [**[]MacpoolIdBlockRelationship**](macpool.IdBlock.Relationship.md) | An array of relationships to macpoolIdBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewMacpoolPool

`func NewMacpoolPool() *MacpoolPool`

NewMacpoolPool instantiates a new MacpoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolPoolWithDefaults

`func NewMacpoolPoolWithDefaults() *MacpoolPool`

NewMacpoolPoolWithDefaults instantiates a new MacpoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacBlocks

`func (o *MacpoolPool) GetMacBlocks() []MacpoolBlock`

GetMacBlocks returns the MacBlocks field if non-nil, zero value otherwise.

### GetMacBlocksOk

`func (o *MacpoolPool) GetMacBlocksOk() (*[]MacpoolBlock, bool)`

GetMacBlocksOk returns a tuple with the MacBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocks

`func (o *MacpoolPool) SetMacBlocks(v []MacpoolBlock)`

SetMacBlocks sets MacBlocks field to given value.

### HasMacBlocks

`func (o *MacpoolPool) HasMacBlocks() bool`

HasMacBlocks returns a boolean if a field has been set.

### GetBlockHeads

`func (o *MacpoolPool) GetBlockHeads() []MacpoolIdBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *MacpoolPool) GetBlockHeadsOk() (*[]MacpoolIdBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *MacpoolPool) SetBlockHeads(v []MacpoolIdBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *MacpoolPool) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *MacpoolPool) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *MacpoolPool) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *MacpoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MacpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MacpoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MacpoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


