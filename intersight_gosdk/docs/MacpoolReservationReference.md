# MacpoolReservationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.ReservationReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.ReservationReference"]
**ConsumerName** | Pointer to **string** | The consumer name for which the reserved MAC would be used. | [optional] 
**ConsumerType** | Pointer to **string** | The consumer type for which the reserved MAC would be used. * &#x60;Vnic&#x60; - MAC reservation would be used by VNIC. | [optional] [default to "Vnic"]

## Methods

### NewMacpoolReservationReference

`func NewMacpoolReservationReference(classId string, objectType string, ) *MacpoolReservationReference`

NewMacpoolReservationReference instantiates a new MacpoolReservationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolReservationReferenceWithDefaults

`func NewMacpoolReservationReferenceWithDefaults() *MacpoolReservationReference`

NewMacpoolReservationReferenceWithDefaults instantiates a new MacpoolReservationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolReservationReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolReservationReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolReservationReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolReservationReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolReservationReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolReservationReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsumerName

`func (o *MacpoolReservationReference) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *MacpoolReservationReference) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *MacpoolReservationReference) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *MacpoolReservationReference) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetConsumerType

`func (o *MacpoolReservationReference) GetConsumerType() string`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *MacpoolReservationReference) GetConsumerTypeOk() (*string, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *MacpoolReservationReference) SetConsumerType(v string)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *MacpoolReservationReference) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


