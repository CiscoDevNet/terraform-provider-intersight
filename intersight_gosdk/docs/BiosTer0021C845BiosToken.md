# BiosTer0021C845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.Ter0021C845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.Ter0021C845BiosToken"]
**Value** | Pointer to **string** | Selects serial port transmission speed. The speed must be matched on the other side. Long or noisy lines may require lower speeds. * &#x60;115200&#x60; - Value -- 115200 for configuring Ter0021 token. * &#x60;9600&#x60; - Value -- 9600 for configuring Ter0021 token. * &#x60;19200&#x60; - Value -- 19200 for configuring Ter0021 token. * &#x60;38400&#x60; - Value -- 38400 for configuring Ter0021 token. * &#x60;57600&#x60; - Value -- 57600 for configuring Ter0021 token. * &#x60;230400&#x60; - Value -- 230400 for configuring Ter0021 token. * &#x60;460800&#x60; - Value -- 460800 for configuring Ter0021 token. * &#x60;921600&#x60; - Value -- 921600 for configuring Ter0021 token. | [optional] [default to "115200"]

## Methods

### NewBiosTer0021C845BiosToken

`func NewBiosTer0021C845BiosToken(classId string, objectType string, ) *BiosTer0021C845BiosToken`

NewBiosTer0021C845BiosToken instantiates a new BiosTer0021C845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosTer0021C845BiosTokenWithDefaults

`func NewBiosTer0021C845BiosTokenWithDefaults() *BiosTer0021C845BiosToken`

NewBiosTer0021C845BiosTokenWithDefaults instantiates a new BiosTer0021C845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosTer0021C845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosTer0021C845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosTer0021C845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosTer0021C845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosTer0021C845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosTer0021C845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosTer0021C845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosTer0021C845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosTer0021C845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosTer0021C845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


