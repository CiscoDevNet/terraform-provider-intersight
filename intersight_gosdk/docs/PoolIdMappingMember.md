# PoolIdMappingMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pool.IdMappingMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pool.IdMappingMember"]
**IdMappingPolicy** | Pointer to [**NullablePoolIdMappingPolicyRelationship**](PoolIdMappingPolicyRelationship.md) |  | [optional] 
**Resource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewPoolIdMappingMember

`func NewPoolIdMappingMember(classId string, objectType string, ) *PoolIdMappingMember`

NewPoolIdMappingMember instantiates a new PoolIdMappingMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolIdMappingMemberWithDefaults

`func NewPoolIdMappingMemberWithDefaults() *PoolIdMappingMember`

NewPoolIdMappingMemberWithDefaults instantiates a new PoolIdMappingMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolIdMappingMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolIdMappingMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolIdMappingMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolIdMappingMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolIdMappingMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolIdMappingMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdMappingPolicy

`func (o *PoolIdMappingMember) GetIdMappingPolicy() PoolIdMappingPolicyRelationship`

GetIdMappingPolicy returns the IdMappingPolicy field if non-nil, zero value otherwise.

### GetIdMappingPolicyOk

`func (o *PoolIdMappingMember) GetIdMappingPolicyOk() (*PoolIdMappingPolicyRelationship, bool)`

GetIdMappingPolicyOk returns a tuple with the IdMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingPolicy

`func (o *PoolIdMappingMember) SetIdMappingPolicy(v PoolIdMappingPolicyRelationship)`

SetIdMappingPolicy sets IdMappingPolicy field to given value.

### HasIdMappingPolicy

`func (o *PoolIdMappingMember) HasIdMappingPolicy() bool`

HasIdMappingPolicy returns a boolean if a field has been set.

### SetIdMappingPolicyNil

`func (o *PoolIdMappingMember) SetIdMappingPolicyNil(b bool)`

 SetIdMappingPolicyNil sets the value for IdMappingPolicy to be an explicit nil

### UnsetIdMappingPolicy
`func (o *PoolIdMappingMember) UnsetIdMappingPolicy()`

UnsetIdMappingPolicy ensures that no value is present for IdMappingPolicy, not even an explicit nil
### GetResource

`func (o *PoolIdMappingMember) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PoolIdMappingMember) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PoolIdMappingMember) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PoolIdMappingMember) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *PoolIdMappingMember) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *PoolIdMappingMember) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


