# CrdCustomResourceConfigProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "crd.CustomResourceConfigProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "crd.CustomResourceConfigProperty"]
**Key** | Pointer to **string** | Name of the key/value property. | [optional] 
**Value** | Pointer to **string** | Value of the key/value property. | [optional] 

## Methods

### NewCrdCustomResourceConfigProperty

`func NewCrdCustomResourceConfigProperty(classId string, objectType string, ) *CrdCustomResourceConfigProperty`

NewCrdCustomResourceConfigProperty instantiates a new CrdCustomResourceConfigProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrdCustomResourceConfigPropertyWithDefaults

`func NewCrdCustomResourceConfigPropertyWithDefaults() *CrdCustomResourceConfigProperty`

NewCrdCustomResourceConfigPropertyWithDefaults instantiates a new CrdCustomResourceConfigProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CrdCustomResourceConfigProperty) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CrdCustomResourceConfigProperty) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CrdCustomResourceConfigProperty) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CrdCustomResourceConfigProperty) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CrdCustomResourceConfigProperty) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CrdCustomResourceConfigProperty) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKey

`func (o *CrdCustomResourceConfigProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CrdCustomResourceConfigProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CrdCustomResourceConfigProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CrdCustomResourceConfigProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *CrdCustomResourceConfigProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CrdCustomResourceConfigProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CrdCustomResourceConfigProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CrdCustomResourceConfigProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


