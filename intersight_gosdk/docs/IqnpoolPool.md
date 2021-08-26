# IqnpoolPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.Pool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.Pool"]
**IqnSuffixBlocks** | Pointer to [**[]IqnpoolIqnSuffixBlock**](IqnpoolIqnSuffixBlock.md) |  | [optional] 
**Prefix** | Pointer to **string** | The prefix for any IQN blocks created for this pool. IQN Prefix must have the following format \&quot;iqn.yyyy-mm.naming-authority\&quot;, where naming-authority is usually the reverse syntax of the Internet domain name of the naming authority. | [optional] 
**BlockHeads** | Pointer to [**[]IqnpoolBlockRelationship**](IqnpoolBlockRelationship.md) | An array of relationships to iqnpoolBlock resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolPool

`func NewIqnpoolPool(classId string, objectType string, ) *IqnpoolPool`

NewIqnpoolPool instantiates a new IqnpoolPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolPoolWithDefaults

`func NewIqnpoolPoolWithDefaults() *IqnpoolPool`

NewIqnpoolPoolWithDefaults instantiates a new IqnpoolPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnSuffixBlocks

`func (o *IqnpoolPool) GetIqnSuffixBlocks() []IqnpoolIqnSuffixBlock`

GetIqnSuffixBlocks returns the IqnSuffixBlocks field if non-nil, zero value otherwise.

### GetIqnSuffixBlocksOk

`func (o *IqnpoolPool) GetIqnSuffixBlocksOk() (*[]IqnpoolIqnSuffixBlock, bool)`

GetIqnSuffixBlocksOk returns a tuple with the IqnSuffixBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffixBlocks

`func (o *IqnpoolPool) SetIqnSuffixBlocks(v []IqnpoolIqnSuffixBlock)`

SetIqnSuffixBlocks sets IqnSuffixBlocks field to given value.

### HasIqnSuffixBlocks

`func (o *IqnpoolPool) HasIqnSuffixBlocks() bool`

HasIqnSuffixBlocks returns a boolean if a field has been set.

### SetIqnSuffixBlocksNil

`func (o *IqnpoolPool) SetIqnSuffixBlocksNil(b bool)`

 SetIqnSuffixBlocksNil sets the value for IqnSuffixBlocks to be an explicit nil

### UnsetIqnSuffixBlocks
`func (o *IqnpoolPool) UnsetIqnSuffixBlocks()`

UnsetIqnSuffixBlocks ensures that no value is present for IqnSuffixBlocks, not even an explicit nil
### GetPrefix

`func (o *IqnpoolPool) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IqnpoolPool) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IqnpoolPool) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IqnpoolPool) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetBlockHeads

`func (o *IqnpoolPool) GetBlockHeads() []IqnpoolBlockRelationship`

GetBlockHeads returns the BlockHeads field if non-nil, zero value otherwise.

### GetBlockHeadsOk

`func (o *IqnpoolPool) GetBlockHeadsOk() (*[]IqnpoolBlockRelationship, bool)`

GetBlockHeadsOk returns a tuple with the BlockHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeads

`func (o *IqnpoolPool) SetBlockHeads(v []IqnpoolBlockRelationship)`

SetBlockHeads sets BlockHeads field to given value.

### HasBlockHeads

`func (o *IqnpoolPool) HasBlockHeads() bool`

HasBlockHeads returns a boolean if a field has been set.

### SetBlockHeadsNil

`func (o *IqnpoolPool) SetBlockHeadsNil(b bool)`

 SetBlockHeadsNil sets the value for BlockHeads to be an explicit nil

### UnsetBlockHeads
`func (o *IqnpoolPool) UnsetBlockHeads()`

UnsetBlockHeads ensures that no value is present for BlockHeads, not even an explicit nil
### GetOrganization

`func (o *IqnpoolPool) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IqnpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IqnpoolPool) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IqnpoolPool) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


