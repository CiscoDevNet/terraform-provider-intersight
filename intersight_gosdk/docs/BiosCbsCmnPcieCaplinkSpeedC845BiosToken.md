# BiosCbsCmnPcieCaplinkSpeedC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnPcieCaplinkSpeedC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnPcieCaplinkSpeedC845BiosToken"]
**Value** | Pointer to **string** | Set all PCIe port speed capability. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;Maximum speed&#x60; - Value -- Maximum speed for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;Gen1&#x60; - Value -- Gen1 for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;Gen2&#x60; - Value -- Gen2 for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;GEN3&#x60; - Value -- GEN3 for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;GEN4&#x60; - Value -- GEN4 for configuring CbsCmnPcieCaplinkSpeed token. * &#x60;GEN5&#x60; - Value -- GEN5 for configuring CbsCmnPcieCaplinkSpeed token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsCmnPcieCaplinkSpeedC845BiosToken

`func NewBiosCbsCmnPcieCaplinkSpeedC845BiosToken(classId string, objectType string, ) *BiosCbsCmnPcieCaplinkSpeedC845BiosToken`

NewBiosCbsCmnPcieCaplinkSpeedC845BiosToken instantiates a new BiosCbsCmnPcieCaplinkSpeedC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnPcieCaplinkSpeedC845BiosTokenWithDefaults

`func NewBiosCbsCmnPcieCaplinkSpeedC845BiosTokenWithDefaults() *BiosCbsCmnPcieCaplinkSpeedC845BiosToken`

NewBiosCbsCmnPcieCaplinkSpeedC845BiosTokenWithDefaults instantiates a new BiosCbsCmnPcieCaplinkSpeedC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnPcieCaplinkSpeedC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


