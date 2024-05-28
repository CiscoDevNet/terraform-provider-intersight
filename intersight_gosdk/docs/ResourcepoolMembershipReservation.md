# ResourcepoolMembershipReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.MembershipReservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.MembershipReservation"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Pools** | Pointer to [**[]ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) | An array of relationships to resourcepoolPool resources. | [optional] 

## Methods

### NewResourcepoolMembershipReservation

`func NewResourcepoolMembershipReservation(classId string, objectType string, ) *ResourcepoolMembershipReservation`

NewResourcepoolMembershipReservation instantiates a new ResourcepoolMembershipReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolMembershipReservationWithDefaults

`func NewResourcepoolMembershipReservationWithDefaults() *ResourcepoolMembershipReservation`

NewResourcepoolMembershipReservationWithDefaults instantiates a new ResourcepoolMembershipReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolMembershipReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolMembershipReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolMembershipReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolMembershipReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolMembershipReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolMembershipReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *ResourcepoolMembershipReservation) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourcepoolMembershipReservation) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourcepoolMembershipReservation) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourcepoolMembershipReservation) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ResourcepoolMembershipReservation) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ResourcepoolMembershipReservation) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetPools

`func (o *ResourcepoolMembershipReservation) GetPools() []ResourcepoolPoolRelationship`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ResourcepoolMembershipReservation) GetPoolsOk() (*[]ResourcepoolPoolRelationship, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ResourcepoolMembershipReservation) SetPools(v []ResourcepoolPoolRelationship)`

SetPools sets Pools field to given value.

### HasPools

`func (o *ResourcepoolMembershipReservation) HasPools() bool`

HasPools returns a boolean if a field has been set.

### SetPoolsNil

`func (o *ResourcepoolMembershipReservation) SetPoolsNil(b bool)`

 SetPoolsNil sets the value for Pools to be an explicit nil

### UnsetPools
`func (o *ResourcepoolMembershipReservation) UnsetPools()`

UnsetPools ensures that no value is present for Pools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


