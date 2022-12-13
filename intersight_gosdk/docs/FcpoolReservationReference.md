# FcpoolReservationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.ReservationReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.ReservationReference"]
**ConsumerName** | Pointer to **string** | The consumer name for which the reserved fc pool would be used. | [optional] 
**ConsumerType** | Pointer to **string** | The consumer type for which the reserved fc pool would be used. * &#x60;Vhba&#x60; - FC reservation would be used by Vhba. * &#x60;WWNN&#x60; - FC reservation would be used by WWNN. | [optional] [default to "Vhba"]

## Methods

### NewFcpoolReservationReference

`func NewFcpoolReservationReference(classId string, objectType string, ) *FcpoolReservationReference`

NewFcpoolReservationReference instantiates a new FcpoolReservationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolReservationReferenceWithDefaults

`func NewFcpoolReservationReferenceWithDefaults() *FcpoolReservationReference`

NewFcpoolReservationReferenceWithDefaults instantiates a new FcpoolReservationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolReservationReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolReservationReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolReservationReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolReservationReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolReservationReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolReservationReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsumerName

`func (o *FcpoolReservationReference) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *FcpoolReservationReference) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *FcpoolReservationReference) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *FcpoolReservationReference) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetConsumerType

`func (o *FcpoolReservationReference) GetConsumerType() string`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *FcpoolReservationReference) GetConsumerTypeOk() (*string, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *FcpoolReservationReference) SetConsumerType(v string)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *FcpoolReservationReference) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


