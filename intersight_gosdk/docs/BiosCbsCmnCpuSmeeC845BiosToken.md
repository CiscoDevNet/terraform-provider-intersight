# BiosCbsCmnCpuSmeeC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnCpuSmeeC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnCpuSmeeC845BiosToken"]
**Value** | Pointer to **string** | Control secure memory encryption enable  Enabling both SMEE and SME-MK is not supported. Results in #GP. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsCmnCpuSmee token. * &#x60;Disable&#x60; - Value -- Disable for configuring CbsCmnCpuSmee token. * &#x60;Enable&#x60; - Value -- Enable for configuring CbsCmnCpuSmee token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsCmnCpuSmeeC845BiosToken

`func NewBiosCbsCmnCpuSmeeC845BiosToken(classId string, objectType string, ) *BiosCbsCmnCpuSmeeC845BiosToken`

NewBiosCbsCmnCpuSmeeC845BiosToken instantiates a new BiosCbsCmnCpuSmeeC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnCpuSmeeC845BiosTokenWithDefaults

`func NewBiosCbsCmnCpuSmeeC845BiosTokenWithDefaults() *BiosCbsCmnCpuSmeeC845BiosToken`

NewBiosCbsCmnCpuSmeeC845BiosTokenWithDefaults instantiates a new BiosCbsCmnCpuSmeeC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnCpuSmeeC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


