# MacpoolMemberOfAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.MemberOf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.MemberOf"]
**IdBlockMoid** | Pointer to **string** | The moid of the block of a pool to which this ID belongs. | [optional] [readonly] 
**PoolMoid** | Pointer to **string** | The moid of the pool to which this ID belongs. | [optional] [readonly] 

## Methods

### NewMacpoolMemberOfAllOf

`func NewMacpoolMemberOfAllOf(classId string, objectType string, ) *MacpoolMemberOfAllOf`

NewMacpoolMemberOfAllOf instantiates a new MacpoolMemberOfAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolMemberOfAllOfWithDefaults

`func NewMacpoolMemberOfAllOfWithDefaults() *MacpoolMemberOfAllOf`

NewMacpoolMemberOfAllOfWithDefaults instantiates a new MacpoolMemberOfAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolMemberOfAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolMemberOfAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolMemberOfAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolMemberOfAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolMemberOfAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolMemberOfAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdBlockMoid

`func (o *MacpoolMemberOfAllOf) GetIdBlockMoid() string`

GetIdBlockMoid returns the IdBlockMoid field if non-nil, zero value otherwise.

### GetIdBlockMoidOk

`func (o *MacpoolMemberOfAllOf) GetIdBlockMoidOk() (*string, bool)`

GetIdBlockMoidOk returns a tuple with the IdBlockMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdBlockMoid

`func (o *MacpoolMemberOfAllOf) SetIdBlockMoid(v string)`

SetIdBlockMoid sets IdBlockMoid field to given value.

### HasIdBlockMoid

`func (o *MacpoolMemberOfAllOf) HasIdBlockMoid() bool`

HasIdBlockMoid returns a boolean if a field has been set.

### GetPoolMoid

`func (o *MacpoolMemberOfAllOf) GetPoolMoid() string`

GetPoolMoid returns the PoolMoid field if non-nil, zero value otherwise.

### GetPoolMoidOk

`func (o *MacpoolMemberOfAllOf) GetPoolMoidOk() (*string, bool)`

GetPoolMoidOk returns a tuple with the PoolMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMoid

`func (o *MacpoolMemberOfAllOf) SetPoolMoid(v string)`

SetPoolMoid sets PoolMoid field to given value.

### HasPoolMoid

`func (o *MacpoolMemberOfAllOf) HasPoolMoid() bool`

HasPoolMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


