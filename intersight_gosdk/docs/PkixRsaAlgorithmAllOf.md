# PkixRsaAlgorithmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pkix.RsaAlgorithm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pkix.RsaAlgorithm"]
**Modulus** | Pointer to **int32** | The length of the RSA key, expressed in bits, for both public and private keys. * &#x60;2048&#x60; - A key length of 2048 bits. * &#x60;2560&#x60; - A key length of 2560 bits. * &#x60;3072&#x60; - A key length of 3072 bits. * &#x60;3584&#x60; - A key length of 3584 bits. * &#x60;4096&#x60; - A key length of 4096 bits. | [optional] [default to 2048]

## Methods

### NewPkixRsaAlgorithmAllOf

`func NewPkixRsaAlgorithmAllOf(classId string, objectType string, ) *PkixRsaAlgorithmAllOf`

NewPkixRsaAlgorithmAllOf instantiates a new PkixRsaAlgorithmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixRsaAlgorithmAllOfWithDefaults

`func NewPkixRsaAlgorithmAllOfWithDefaults() *PkixRsaAlgorithmAllOf`

NewPkixRsaAlgorithmAllOfWithDefaults instantiates a new PkixRsaAlgorithmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PkixRsaAlgorithmAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PkixRsaAlgorithmAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PkixRsaAlgorithmAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PkixRsaAlgorithmAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PkixRsaAlgorithmAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PkixRsaAlgorithmAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModulus

`func (o *PkixRsaAlgorithmAllOf) GetModulus() int32`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *PkixRsaAlgorithmAllOf) GetModulusOk() (*int32, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulus

`func (o *PkixRsaAlgorithmAllOf) SetModulus(v int32)`

SetModulus sets Modulus field to given value.

### HasModulus

`func (o *PkixRsaAlgorithmAllOf) HasModulus() bool`

HasModulus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


