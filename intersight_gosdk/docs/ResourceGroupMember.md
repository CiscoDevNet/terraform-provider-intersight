# ResourceGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.GroupMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.GroupMember"]
**Group** | Pointer to [**ResourceGroupRelationship**](ResourceGroupRelationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourceGroupMember

`func NewResourceGroupMember(classId string, objectType string, ) *ResourceGroupMember`

NewResourceGroupMember instantiates a new ResourceGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupMemberWithDefaults

`func NewResourceGroupMemberWithDefaults() *ResourceGroupMember`

NewResourceGroupMemberWithDefaults instantiates a new ResourceGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceGroupMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceGroupMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceGroupMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceGroupMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceGroupMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceGroupMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


