# PkixEcdsaKeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pkix.EcdsaKeySpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pkix.EcdsaKeySpec"]
**Curve** | Pointer to **string** | A specific set of Elliptic Curve parameters, as recommended by NIST in FIPS 186-4. * &#x60;P256&#x60; - P256 returns a Curve which implements P-256, as defined in FIPS 186-4, section D.2.3. * &#x60;P224&#x60; - P224 returns a Curve which implements P-224, as defined in FIPS 186-4, section D.2.2. * &#x60;P384&#x60; - P384 returns a Curve which implements P-384, as defined in FIPS 186-4, section D.2.4. * &#x60;P521&#x60; - P521 returns a Curve which implements P-521, as defined in FIPS 186-4, section D.2.5. | [optional] [default to "P256"]

## Methods

### NewPkixEcdsaKeySpec

`func NewPkixEcdsaKeySpec(classId string, objectType string, ) *PkixEcdsaKeySpec`

NewPkixEcdsaKeySpec instantiates a new PkixEcdsaKeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixEcdsaKeySpecWithDefaults

`func NewPkixEcdsaKeySpecWithDefaults() *PkixEcdsaKeySpec`

NewPkixEcdsaKeySpecWithDefaults instantiates a new PkixEcdsaKeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PkixEcdsaKeySpec) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PkixEcdsaKeySpec) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PkixEcdsaKeySpec) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PkixEcdsaKeySpec) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PkixEcdsaKeySpec) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PkixEcdsaKeySpec) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurve

`func (o *PkixEcdsaKeySpec) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *PkixEcdsaKeySpec) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *PkixEcdsaKeySpec) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *PkixEcdsaKeySpec) HasCurve() bool`

HasCurve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


