# UuidpoolPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.Pool"]
**Prefix** | Pointer to **string** | The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx. | [optional] 
**UuidSuffixBlocks** | Pointer to [**[]UuidpoolUuidBlock**](UuidpoolUuidBlock.md) |  | [optional] 
**BlockHeads** | Pointer to [**[]UuidpoolBlockRelationship**](UuidpoolBlockRelationship.md) | An array of relationships to uuidpoolBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewUuidpoolPool

`func NewUuidpoolPool(classId string, objectType string, ) *UuidpoolPool`

NewUuidpoolPool instantiates a new UuidpoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolPoolWithDefaults

`func NewUuidpoolPoolWithDefaults() *UuidpoolPool`

NewUuidpoolPoolWithDefaults instantiates a new UuidpoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetUuidSuffixBlocksNil

`func (o *UuidpoolPool) SetUuidSuffixBlocksNil(b bool)`

 SetUuidSuffixBlocksNil sets the value for UuidSuffixBlocks to be an explicit nil

### UnsetUuidSuffixBlocks
`func (o *UuidpoolPool) UnsetUuidSuffixBlocks()`

UnsetUuidSuffixBlocks ensures that no value is present for UuidSuffixBlocks, not even an explicit nil
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


