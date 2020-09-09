# MacpoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacBlocks** | Pointer to [**[]MacpoolBlock**](macpool.Block.md) |  | [optional] 
**BlockHeads** | Pointer to [**[]MacpoolIdBlockRelationship**](macpool.IdBlock.Relationship.md) | An array of relationships to macpoolIdBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewMacpoolPoolAllOf

`func NewMacpoolPoolAllOf() *MacpoolPoolAllOf`

NewMacpoolPoolAllOf instantiates a new MacpoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolPoolAllOfWithDefaults

`func NewMacpoolPoolAllOfWithDefaults() *MacpoolPoolAllOf`

NewMacpoolPoolAllOfWithDefaults instantiates a new MacpoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacBlocks

`func (o *MacpoolPoolAllOf) GetMacBlocks() []MacpoolBlock`

GetMacBlocks returns the MacBlocks field if non-nil, zero value otherwise.

### GetMacBlocksOk

`func (o *MacpoolPoolAllOf) GetMacBlocksOk() (*[]MacpoolBlock, bool)`

GetMacBlocksOk returns a tuple with the MacBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacBlocks

`func (o *MacpoolPoolAllOf) SetMacBlocks(v []MacpoolBlock)`

SetMacBlocks sets MacBlocks field to given value.

### HasMacBlocks

`func (o *MacpoolPoolAllOf) HasMacBlocks() bool`

HasMacBlocks returns a boolean if a field has been set.

### GetBlockHeads

`func (o *MacpoolPoolAllOf) GetBlockHeads() []MacpoolIdBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *MacpoolPoolAllOf) GetBlockHeadsOk() (*[]MacpoolIdBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *MacpoolPoolAllOf) SetBlockHeads(v []MacpoolIdBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *MacpoolPoolAllOf) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *MacpoolPoolAllOf) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *MacpoolPoolAllOf) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *MacpoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MacpoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MacpoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MacpoolPoolAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


