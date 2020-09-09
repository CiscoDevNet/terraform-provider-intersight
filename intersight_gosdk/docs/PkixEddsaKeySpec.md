# PkixEddsaKeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The EdDSA algorithm, as defined in RFC 8032. * &#x60;Ed25519&#x60; - The edwards25519 curve, as defined in RFC 8032 section 5.1. * &#x60;Ed25519ph&#x60; - The edwards25519 curve for the PureEdDSA variant. * &#x60;Ed25519ctx&#x60; - The edwards25519 curve for the HashEdDSA variant. | [optional] [default to "Ed25519"]

## Methods

### NewPkixEddsaKeySpec

`func NewPkixEddsaKeySpec() *PkixEddsaKeySpec`

NewPkixEddsaKeySpec instantiates a new PkixEddsaKeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixEddsaKeySpecWithDefaults

`func NewPkixEddsaKeySpecWithDefaults() *PkixEddsaKeySpec`

NewPkixEddsaKeySpecWithDefaults instantiates a new PkixEddsaKeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *PkixEddsaKeySpec) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PkixEddsaKeySpec) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PkixEddsaKeySpec) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PkixEddsaKeySpec) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


