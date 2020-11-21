# PkixEddsaKeySpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pkix.EddsaKeySpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pkix.EddsaKeySpec"]
**Algorithm** | Pointer to **string** | The EdDSA algorithm, as defined in RFC 8032. * &#x60;Ed25519&#x60; - The edwards25519 curve, as defined in RFC 8032 section 5.1. * &#x60;Ed25519ph&#x60; - The edwards25519 curve for the PureEdDSA variant. * &#x60;Ed25519ctx&#x60; - The edwards25519 curve for the HashEdDSA variant. | [optional] [default to "Ed25519"]

## Methods

### NewPkixEddsaKeySpecAllOf

`func NewPkixEddsaKeySpecAllOf(classId string, objectType string, ) *PkixEddsaKeySpecAllOf`

NewPkixEddsaKeySpecAllOf instantiates a new PkixEddsaKeySpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixEddsaKeySpecAllOfWithDefaults

`func NewPkixEddsaKeySpecAllOfWithDefaults() *PkixEddsaKeySpecAllOf`

NewPkixEddsaKeySpecAllOfWithDefaults instantiates a new PkixEddsaKeySpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PkixEddsaKeySpecAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PkixEddsaKeySpecAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PkixEddsaKeySpecAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PkixEddsaKeySpecAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PkixEddsaKeySpecAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PkixEddsaKeySpecAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlgorithm

`func (o *PkixEddsaKeySpecAllOf) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PkixEddsaKeySpecAllOf) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PkixEddsaKeySpecAllOf) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PkixEddsaKeySpecAllOf) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


