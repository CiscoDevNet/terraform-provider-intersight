# PoolAbstractPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Assigned** | Pointer to **bool** | Boolean to represent whether the ID is assigned or not. | [optional] [default to false]

## Methods

### NewPoolAbstractPoolMember

`func NewPoolAbstractPoolMember(classId string, objectType string, ) *PoolAbstractPoolMember`

NewPoolAbstractPoolMember instantiates a new PoolAbstractPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractPoolMemberWithDefaults

`func NewPoolAbstractPoolMemberWithDefaults() *PoolAbstractPoolMember`

NewPoolAbstractPoolMemberWithDefaults instantiates a new PoolAbstractPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssigned

`func (o *PoolAbstractPoolMember) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *PoolAbstractPoolMember) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *PoolAbstractPoolMember) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *PoolAbstractPoolMember) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


