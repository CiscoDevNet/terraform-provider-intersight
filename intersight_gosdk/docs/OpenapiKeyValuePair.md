# OpenapiKeyValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "openapi.KeyValuePair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "openapi.KeyValuePair"]
**Key** | Pointer to **string** | Name of the key variable. | [optional] 
**Value** | Pointer to **string** | Value of the key variable. | [optional] 

## Methods

### NewOpenapiKeyValuePair

`func NewOpenapiKeyValuePair(classId string, objectType string, ) *OpenapiKeyValuePair`

NewOpenapiKeyValuePair instantiates a new OpenapiKeyValuePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiKeyValuePairWithDefaults

`func NewOpenapiKeyValuePairWithDefaults() *OpenapiKeyValuePair`

NewOpenapiKeyValuePairWithDefaults instantiates a new OpenapiKeyValuePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OpenapiKeyValuePair) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OpenapiKeyValuePair) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OpenapiKeyValuePair) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OpenapiKeyValuePair) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenapiKeyValuePair) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenapiKeyValuePair) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *OpenapiKeyValuePair) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OpenapiKeyValuePair) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OpenapiKeyValuePair) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OpenapiKeyValuePair) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *OpenapiKeyValuePair) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OpenapiKeyValuePair) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OpenapiKeyValuePair) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OpenapiKeyValuePair) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


