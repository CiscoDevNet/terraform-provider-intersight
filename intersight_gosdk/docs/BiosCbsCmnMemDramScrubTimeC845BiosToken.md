# BiosCbsCmnMemDramScrubTimeC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnMemDramScrubTimeC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnMemDramScrubTimeC845BiosToken"]
**Value** | Pointer to **string** | Provide a value that is the number of hours to scrub memory. * &#x60;24 hours&#x60; - Value -- 24 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;Disabled&#x60; - Value -- Disabled for configuring CbsCmnMemDramScrubTime token. * &#x60;1 hour&#x60; - Value -- 1 hour for configuring CbsCmnMemDramScrubTime token. * &#x60;4 hours&#x60; - Value -- 4 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;6 hours&#x60; - Value -- 6 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;8 hours&#x60; - Value -- 8 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;12 hours&#x60; - Value -- 12 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;16 hours&#x60; - Value -- 16 hours for configuring CbsCmnMemDramScrubTime token. * &#x60;48 hours&#x60; - Value -- 48 hours for configuring CbsCmnMemDramScrubTime token. | [optional] [default to "24 hours"]

## Methods

### NewBiosCbsCmnMemDramScrubTimeC845BiosToken

`func NewBiosCbsCmnMemDramScrubTimeC845BiosToken(classId string, objectType string, ) *BiosCbsCmnMemDramScrubTimeC845BiosToken`

NewBiosCbsCmnMemDramScrubTimeC845BiosToken instantiates a new BiosCbsCmnMemDramScrubTimeC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnMemDramScrubTimeC845BiosTokenWithDefaults

`func NewBiosCbsCmnMemDramScrubTimeC845BiosTokenWithDefaults() *BiosCbsCmnMemDramScrubTimeC845BiosToken`

NewBiosCbsCmnMemDramScrubTimeC845BiosTokenWithDefaults instantiates a new BiosCbsCmnMemDramScrubTimeC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnMemDramScrubTimeC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


