# OprsKvpair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "oprs.Kvpair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "oprs.Kvpair"]
**Key** | Pointer to **string** | Name of the property which is used as key. | [optional] 
**Value** | Pointer to **string** | Value of the property corresponding to the key. | [optional] 

## Methods

### NewOprsKvpair

`func NewOprsKvpair(classId string, objectType string, ) *OprsKvpair`

NewOprsKvpair instantiates a new OprsKvpair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOprsKvpairWithDefaults

`func NewOprsKvpairWithDefaults() *OprsKvpair`

NewOprsKvpairWithDefaults instantiates a new OprsKvpair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OprsKvpair) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OprsKvpair) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OprsKvpair) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OprsKvpair) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OprsKvpair) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OprsKvpair) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *OprsKvpair) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OprsKvpair) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OprsKvpair) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OprsKvpair) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *OprsKvpair) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OprsKvpair) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OprsKvpair) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OprsKvpair) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


