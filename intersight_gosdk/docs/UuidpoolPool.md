# UuidpoolPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx. | [optional] 
**UuidSuffixBlocks** | Pointer to [**[]UuidpoolUuidBlock**](uuidpool.UuidBlock.md) |  | [optional] 
**BlockHeads** | Pointer to [**[]UuidpoolBlockRelationship**](uuidpool.Block.Relationship.md) | An array of relationships to uuidpoolBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewUuidpoolPool

`func NewUuidpoolPool() *UuidpoolPool`

NewUuidpoolPool instantiates a new UuidpoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolPoolWithDefaults

`func NewUuidpoolPoolWithDefaults() *UuidpoolPool`

NewUuidpoolPoolWithDefaults instantiates a new UuidpoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *UuidpoolPool) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UuidpoolPool) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UuidpoolPool) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UuidpoolPool) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetUuidSuffixBlocks

`func (o *UuidpoolPool) GetUuidSuffixBlocks() []UuidpoolUuidBlock`

GetUuidSuffixBlocks returns the UuidSuffixBlocks field if non-nil, zero value otherwise.

### GetUuidSuffixBlocksOk

`func (o *UuidpoolPool) GetUuidSuffixBlocksOk() (*[]UuidpoolUuidBlock, bool)`

GetUuidSuffixBlocksOk returns a tuple with the UuidSuffixBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidSuffixBlocks

`func (o *UuidpoolPool) SetUuidSuffixBlocks(v []UuidpoolUuidBlock)`

SetUuidSuffixBlocks sets UuidSuffixBlocks field to given value.

### HasUuidSuffixBlocks

`func (o *UuidpoolPool) HasUuidSuffixBlocks() bool`

HasUuidSuffixBlocks returns a boolean if a field has been set.

### GetBlockHeads

`func (o *UuidpoolPool) GetBlockHeads() []UuidpoolBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *UuidpoolPool) GetBlockHeadsOk() (*[]UuidpoolBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *UuidpoolPool) SetBlockHeads(v []UuidpoolBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *UuidpoolPool) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *UuidpoolPool) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *UuidpoolPool) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *UuidpoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UuidpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UuidpoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UuidpoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


