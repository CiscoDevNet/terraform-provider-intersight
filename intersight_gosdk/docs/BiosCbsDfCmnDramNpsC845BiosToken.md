# BiosCbsDfCmnDramNpsC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsDfCmnDramNpsC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsDfCmnDramNpsC845BiosToken"]
**Value** | Pointer to **string** | The number of desired NUMA nodes per socket. Zero will attempt to interleave the two sockets together. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsDfCmnDramNps token. * &#x60;NPS0&#x60; - Value -- NPS0 for configuring CbsDfCmnDramNps token. * &#x60;NPS1&#x60; - Value -- NPS1 for configuring CbsDfCmnDramNps token. * &#x60;NPS2&#x60; - Value -- NPS2 for configuring CbsDfCmnDramNps token. * &#x60;NPS4&#x60; - Value -- NPS4 for configuring CbsDfCmnDramNps token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsDfCmnDramNpsC845BiosToken

`func NewBiosCbsDfCmnDramNpsC845BiosToken(classId string, objectType string, ) *BiosCbsDfCmnDramNpsC845BiosToken`

NewBiosCbsDfCmnDramNpsC845BiosToken instantiates a new BiosCbsDfCmnDramNpsC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsDfCmnDramNpsC845BiosTokenWithDefaults

`func NewBiosCbsDfCmnDramNpsC845BiosTokenWithDefaults() *BiosCbsDfCmnDramNpsC845BiosToken`

NewBiosCbsDfCmnDramNpsC845BiosTokenWithDefaults instantiates a new BiosCbsDfCmnDramNpsC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsDfCmnDramNpsC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


