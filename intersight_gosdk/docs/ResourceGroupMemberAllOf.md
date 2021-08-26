# ResourceGroupMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.GroupMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.GroupMember"]
**Group** | Pointer to [**ResourceGroupRelationship**](ResourceGroupRelationship.md) |  | [optional] 
**Resource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourceGroupMemberAllOf

`func NewResourceGroupMemberAllOf(classId string, objectType string, ) *ResourceGroupMemberAllOf`

NewResourceGroupMemberAllOf instantiates a new ResourceGroupMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupMemberAllOfWithDefaults

`func NewResourceGroupMemberAllOfWithDefaults() *ResourceGroupMemberAllOf`

NewResourceGroupMemberAllOfWithDefaults instantiates a new ResourceGroupMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceGroupMemberAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceGroupMemberAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceGroupMemberAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceGroupMemberAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceGroupMemberAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceGroupMemberAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGroup

`func (o *ResourceGroupMemberAllOf) GetGroup() ResourceGroupRelationship`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ResourceGroupMemberAllOf) GetGroupOk() (*ResourceGroupRelationship, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ResourceGroupMemberAllOf) SetGroup(v ResourceGroupRelationship)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ResourceGroupMemberAllOf) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetResource

`func (o *ResourceGroupMemberAllOf) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceGroupMemberAllOf) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceGroupMemberAllOf) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceGroupMemberAllOf) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


