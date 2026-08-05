# BiosCbsCmnMemDramRefreshRateC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnMemDramRefreshRateC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnMemDramRefreshRateC845BiosToken"]
**Value** | Pointer to **string** | DRAM refresh rate: 1.95us or 3.9us (default). * &#x60;3.9 usec&#x60; - Value -- 3.9 usec for configuring CbsCmnMemDramRefreshRate token. * &#x60;1.95 usec&#x60; - Value -- 1.95 usec for configuring CbsCmnMemDramRefreshRate token. | [optional] [default to "3.9 usec"]

## Methods

### NewBiosCbsCmnMemDramRefreshRateC845BiosToken

`func NewBiosCbsCmnMemDramRefreshRateC845BiosToken(classId string, objectType string, ) *BiosCbsCmnMemDramRefreshRateC845BiosToken`

NewBiosCbsCmnMemDramRefreshRateC845BiosToken instantiates a new BiosCbsCmnMemDramRefreshRateC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnMemDramRefreshRateC845BiosTokenWithDefaults

`func NewBiosCbsCmnMemDramRefreshRateC845BiosTokenWithDefaults() *BiosCbsCmnMemDramRefreshRateC845BiosToken`

NewBiosCbsCmnMemDramRefreshRateC845BiosTokenWithDefaults instantiates a new BiosCbsCmnMemDramRefreshRateC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnMemDramRefreshRateC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


