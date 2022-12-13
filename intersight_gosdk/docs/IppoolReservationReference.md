# IppoolReservationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.ReservationReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.ReservationReference"]
**ConsumerName** | Pointer to **string** | The consumer name for which the reserved IP would be used. | [optional] 
**ConsumerType** | Pointer to **string** | The consumer type for which the reserved IP would be used. * &#x60;OutofbandIpv4-Access&#x60; - IP reservation would be used for out of band management. * &#x60;InbandIpv4-Access&#x60; - IP reservation would be used for inband management. * &#x60;InbandIpv6-Access&#x60; - IP reservation would be used for inband management. * &#x60;ISCSI&#x60; - IP reservation would be used for ISCSI management. | [optional] [default to "OutofbandIpv4-Access"]

## Methods

### NewIppoolReservationReference

`func NewIppoolReservationReference(classId string, objectType string, ) *IppoolReservationReference`

NewIppoolReservationReference instantiates a new IppoolReservationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolReservationReferenceWithDefaults

`func NewIppoolReservationReferenceWithDefaults() *IppoolReservationReference`

NewIppoolReservationReferenceWithDefaults instantiates a new IppoolReservationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolReservationReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolReservationReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolReservationReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolReservationReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolReservationReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolReservationReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConsumerName

`func (o *IppoolReservationReference) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *IppoolReservationReference) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *IppoolReservationReference) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *IppoolReservationReference) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetConsumerType

`func (o *IppoolReservationReference) GetConsumerType() string`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *IppoolReservationReference) GetConsumerTypeOk() (*string, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *IppoolReservationReference) SetConsumerType(v string)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *IppoolReservationReference) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


