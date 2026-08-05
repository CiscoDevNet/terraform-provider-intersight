# BiosTer05dC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.Ter05dC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.Ter05dC845BiosToken"]
**Value** | Pointer to **string** | Flow control can prevent data loss from buffer overflow. When sending data, if the receiving buffers are full, a &#39;stop&#39; signal can be sent to stop the data flow. Once the buffers are empty, a &#39;start&#39; signal can be sent to re-start the flow. Hardware flow control uses two wires to send start/stop signals. * &#x60;None&#x60; - Value -- None for configuring Ter05d token. * &#x60;Hardware RTS/CTS&#x60; - Value -- Hardware RTS/CTS for configuring Ter05d token. | [optional] [default to "None"]

## Methods

### NewBiosTer05dC845BiosToken

`func NewBiosTer05dC845BiosToken(classId string, objectType string, ) *BiosTer05dC845BiosToken`

NewBiosTer05dC845BiosToken instantiates a new BiosTer05dC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosTer05dC845BiosTokenWithDefaults

`func NewBiosTer05dC845BiosTokenWithDefaults() *BiosTer05dC845BiosToken`

NewBiosTer05dC845BiosTokenWithDefaults instantiates a new BiosTer05dC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosTer05dC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosTer05dC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosTer05dC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosTer05dC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosTer05dC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosTer05dC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosTer05dC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosTer05dC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosTer05dC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosTer05dC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


