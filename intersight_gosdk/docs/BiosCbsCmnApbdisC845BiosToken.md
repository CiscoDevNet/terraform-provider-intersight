# BiosCbsCmnApbdisC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnApbdisC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnApbdisC845BiosToken"]
**Value** | Pointer to **string** | 0 &#x3D; not APBDIS (mission mode)  1 &#x3D; APBDIS. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsCmnApbdis token. * &#x60;0&#x60; - Value -- 0 for configuring CbsCmnApbdis token. * &#x60;1&#x60; - Value -- 1 for configuring CbsCmnApbdis token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsCmnApbdisC845BiosToken

`func NewBiosCbsCmnApbdisC845BiosToken(classId string, objectType string, ) *BiosCbsCmnApbdisC845BiosToken`

NewBiosCbsCmnApbdisC845BiosToken instantiates a new BiosCbsCmnApbdisC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnApbdisC845BiosTokenWithDefaults

`func NewBiosCbsCmnApbdisC845BiosTokenWithDefaults() *BiosCbsCmnApbdisC845BiosToken`

NewBiosCbsCmnApbdisC845BiosTokenWithDefaults instantiates a new BiosCbsCmnApbdisC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnApbdisC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnApbdisC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnApbdisC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnApbdisC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnApbdisC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnApbdisC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnApbdisC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnApbdisC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnApbdisC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnApbdisC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


