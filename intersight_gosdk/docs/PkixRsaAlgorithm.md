# PkixRsaAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modulus** | Pointer to **int32** | The length of the RSA key, expressed in bits, for both public and private keys. * &#x60;2048&#x60; - A key length of 2048 bits. * &#x60;2560&#x60; - A key length of 2560 bits. * &#x60;3072&#x60; - A key length of 3072 bits. * &#x60;3584&#x60; - A key length of 3584 bits. * &#x60;4096&#x60; - A key length of 4096 bits. | [optional] [default to 2048]

## Methods

### NewPkixRsaAlgorithm

`func NewPkixRsaAlgorithm() *PkixRsaAlgorithm`

NewPkixRsaAlgorithm instantiates a new PkixRsaAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixRsaAlgorithmWithDefaults

`func NewPkixRsaAlgorithmWithDefaults() *PkixRsaAlgorithm`

NewPkixRsaAlgorithmWithDefaults instantiates a new PkixRsaAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModulus

`func (o *PkixRsaAlgorithm) GetModulus() int32`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *PkixRsaAlgorithm) GetModulusOk() (*int32, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulus

`func (o *PkixRsaAlgorithm) SetModulus(v int32)`

SetModulus sets Modulus field to given value.

### HasModulus

`func (o *PkixRsaAlgorithm) HasModulus() bool`

HasModulus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


