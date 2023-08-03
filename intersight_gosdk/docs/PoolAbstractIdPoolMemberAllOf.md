# PoolAbstractIdPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AssignedByAnother** | Pointer to **bool** | Boolean to represent whether the ID is used either statically or by another pool. | [optional] [default to false]
**Reserved** | Pointer to **bool** | Boolean to represent whether the ID is reserved. | [optional] [default to false]

## Methods

### NewPoolAbstractIdPoolMemberAllOf

`func NewPoolAbstractIdPoolMemberAllOf(classId string, objectType string, ) *PoolAbstractIdPoolMemberAllOf`

NewPoolAbstractIdPoolMemberAllOf instantiates a new PoolAbstractIdPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractIdPoolMemberAllOfWithDefaults

`func NewPoolAbstractIdPoolMemberAllOfWithDefaults() *PoolAbstractIdPoolMemberAllOf`

NewPoolAbstractIdPoolMemberAllOfWithDefaults instantiates a new PoolAbstractIdPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractIdPoolMemberAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractIdPoolMemberAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractIdPoolMemberAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractIdPoolMemberAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractIdPoolMemberAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractIdPoolMemberAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssignedByAnother

`func (o *PoolAbstractIdPoolMemberAllOf) GetAssignedByAnother() bool`

GetAssignedByAnother returns the AssignedByAnother field if non-nil, zero value otherwise.

### GetAssignedByAnotherOk

`func (o *PoolAbstractIdPoolMemberAllOf) GetAssignedByAnotherOk() (*bool, bool)`

GetAssignedByAnotherOk returns a tuple with the AssignedByAnother field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedByAnother

`func (o *PoolAbstractIdPoolMemberAllOf) SetAssignedByAnother(v bool)`

SetAssignedByAnother sets AssignedByAnother field to given value.

### HasAssignedByAnother

`func (o *PoolAbstractIdPoolMemberAllOf) HasAssignedByAnother() bool`

HasAssignedByAnother returns a boolean if a field has been set.

### GetReserved

`func (o *PoolAbstractIdPoolMemberAllOf) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *PoolAbstractIdPoolMemberAllOf) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *PoolAbstractIdPoolMemberAllOf) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *PoolAbstractIdPoolMemberAllOf) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


