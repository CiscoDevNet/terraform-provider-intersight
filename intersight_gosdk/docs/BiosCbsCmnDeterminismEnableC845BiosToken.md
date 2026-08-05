# BiosCbsCmnDeterminismEnableC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnDeterminismEnableC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnDeterminismEnableC845BiosToken"]
**Value** | Pointer to **string** | [0 &#x3D; Power; 1 &#x3D; Performance]. * &#x60;Power&#x60; - Value -- Power for configuring CbsCmnDeterminismEnable token. * &#x60;Performance&#x60; - Value -- Performance for configuring CbsCmnDeterminismEnable token. | [optional] [default to "Power"]

## Methods

### NewBiosCbsCmnDeterminismEnableC845BiosToken

`func NewBiosCbsCmnDeterminismEnableC845BiosToken(classId string, objectType string, ) *BiosCbsCmnDeterminismEnableC845BiosToken`

NewBiosCbsCmnDeterminismEnableC845BiosToken instantiates a new BiosCbsCmnDeterminismEnableC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnDeterminismEnableC845BiosTokenWithDefaults

`func NewBiosCbsCmnDeterminismEnableC845BiosTokenWithDefaults() *BiosCbsCmnDeterminismEnableC845BiosToken`

NewBiosCbsCmnDeterminismEnableC845BiosTokenWithDefaults instantiates a new BiosCbsCmnDeterminismEnableC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnDeterminismEnableC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


