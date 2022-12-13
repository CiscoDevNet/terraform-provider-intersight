# PoolReservationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ReservationMoid** | Pointer to **string** | The moid of the reservation object. | [optional] 

## Methods

### NewPoolReservationReference

`func NewPoolReservationReference(classId string, objectType string, ) *PoolReservationReference`

NewPoolReservationReference instantiates a new PoolReservationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolReservationReferenceWithDefaults

`func NewPoolReservationReferenceWithDefaults() *PoolReservationReference`

NewPoolReservationReferenceWithDefaults instantiates a new PoolReservationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolReservationReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolReservationReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolReservationReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolReservationReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolReservationReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolReservationReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReservationMoid

`func (o *PoolReservationReference) GetReservationMoid() string`

GetReservationMoid returns the ReservationMoid field if non-nil, zero value otherwise.

### GetReservationMoidOk

`func (o *PoolReservationReference) GetReservationMoidOk() (*string, bool)`

GetReservationMoidOk returns a tuple with the ReservationMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationMoid

`func (o *PoolReservationReference) SetReservationMoid(v string)`

SetReservationMoid sets ReservationMoid field to given value.

### HasReservationMoid

`func (o *PoolReservationReference) HasReservationMoid() bool`

HasReservationMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


