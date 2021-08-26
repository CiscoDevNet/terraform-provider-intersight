# MacpoolPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.Pool"]
**MacBlocks** | Pointer to [**[]MacpoolBlock**](MacpoolBlock.md) |  | [optional] 
**BlockHeads** | Pointer to [**[]MacpoolIdBlockRelationship**](MacpoolIdBlockRelationship.md) | An array of relationships to macpoolIdBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMacpoolPoolAllOf

`func NewMacpoolPoolAllOf(classId string, objectType string, ) *MacpoolPoolAllOf`

NewMacpoolPoolAllOf instantiates a new MacpoolPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolPoolAllOfWithDefaults

`func NewMacpoolPoolAllOfWithDefaults() *MacpoolPoolAllOf`

NewMacpoolPoolAllOfWithDefaults instantiates a new MacpoolPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetMacBlocksNil

`func (o *MacpoolPoolAllOf) SetMacBlocksNil(b bool)`

 SetMacBlocksNil sets the value for MacBlocks to be an explicit nil

### UnsetMacBlocks
`func (o *MacpoolPoolAllOf) UnsetMacBlocks()`

UnsetMacBlocks ensures that no value is present for MacBlocks, not even an explicit nil
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


