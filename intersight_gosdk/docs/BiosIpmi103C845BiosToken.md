# BiosIpmi103C845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.Ipmi103C845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.Ipmi103C845BiosToken"]
**Value** | Pointer to **string** | If enabled, starts a BIOS timer which can only be shut off by Management Software after the OS loads.  Helps determine that the OS successfully loaded or follows the OS Boot Watchdog Timer policy. * &#x60;Disabled&#x60; - Value -- Disabled for configuring Ipmi103 token. * &#x60;Enabled&#x60; - Value -- Enabled for configuring Ipmi103 token. | [optional] [default to "Disabled"]

## Methods

### NewBiosIpmi103C845BiosToken

`func NewBiosIpmi103C845BiosToken(classId string, objectType string, ) *BiosIpmi103C845BiosToken`

NewBiosIpmi103C845BiosToken instantiates a new BiosIpmi103C845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosIpmi103C845BiosTokenWithDefaults

`func NewBiosIpmi103C845BiosTokenWithDefaults() *BiosIpmi103C845BiosToken`

NewBiosIpmi103C845BiosTokenWithDefaults instantiates a new BiosIpmi103C845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosIpmi103C845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosIpmi103C845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosIpmi103C845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosIpmi103C845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosIpmi103C845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosIpmi103C845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosIpmi103C845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosIpmi103C845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosIpmi103C845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosIpmi103C845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


