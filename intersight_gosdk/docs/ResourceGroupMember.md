# ResourceGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**ResourceGroupRelationship**](resource.Group.Relationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 

## Methods

### NewResourceGroupMember

`func NewResourceGroupMember() *ResourceGroupMember`

NewResourceGroupMember instantiates a new ResourceGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupMemberWithDefaults

`func NewResourceGroupMemberWithDefaults() *ResourceGroupMember`

NewResourceGroupMemberWithDefaults instantiates a new ResourceGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ResourceGroupMember) GetGroup() ResourceGroupRelationship`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ResourceGroupMember) GetGroupOk() (*ResourceGroupRelationship, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ResourceGroupMember) SetGroup(v ResourceGroupRelationship)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ResourceGroupMember) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetResource

`func (o *ResourceGroupMember) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceGroupMember) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceGroupMember) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceGroupMember) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


