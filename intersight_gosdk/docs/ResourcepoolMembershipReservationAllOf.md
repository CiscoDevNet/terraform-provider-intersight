# ResourcepoolMembershipReservationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.MembershipReservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.MembershipReservation"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Pools** | Pointer to [**[]ResourcepoolPoolRelationship**](ResourcepoolPoolRelationship.md) | An array of relationships to resourcepoolPool resources. | [optional] 

## Methods

### NewResourcepoolMembershipReservationAllOf

`func NewResourcepoolMembershipReservationAllOf(classId string, objectType string, ) *ResourcepoolMembershipReservationAllOf`

NewResourcepoolMembershipReservationAllOf instantiates a new ResourcepoolMembershipReservationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolMembershipReservationAllOfWithDefaults

`func NewResourcepoolMembershipReservationAllOfWithDefaults() *ResourcepoolMembershipReservationAllOf`

NewResourcepoolMembershipReservationAllOfWithDefaults instantiates a new ResourcepoolMembershipReservationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolMembershipReservationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolMembershipReservationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolMembershipReservationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolMembershipReservationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolMembershipReservationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolMembershipReservationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *ResourcepoolMembershipReservationAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourcepoolMembershipReservationAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourcepoolMembershipReservationAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourcepoolMembershipReservationAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPools

`func (o *ResourcepoolMembershipReservationAllOf) GetPools() []ResourcepoolPoolRelationship`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ResourcepoolMembershipReservationAllOf) GetPoolsOk() (*[]ResourcepoolPoolRelationship, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ResourcepoolMembershipReservationAllOf) SetPools(v []ResourcepoolPoolRelationship)`

SetPools sets Pools field to given value.

### HasPools

`func (o *ResourcepoolMembershipReservationAllOf) HasPools() bool`

HasPools returns a boolean if a field has been set.

### SetPoolsNil

`func (o *ResourcepoolMembershipReservationAllOf) SetPoolsNil(b bool)`

 SetPoolsNil sets the value for Pools to be an explicit nil

### UnsetPools
`func (o *ResourcepoolMembershipReservationAllOf) UnsetPools()`

UnsetPools ensures that no value is present for Pools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


