# BiosCbsCpuCcdCtrlC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCpuCcdCtrlC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCpuCcdCtrlC845BiosToken"]
**Value** | Pointer to **string** | Sets the number of active CCDs.  Once this option has been used to remove any CCDs, a POWER CYCLE is required in order for future selections to take effect. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsCpuCcdCtrl token. * &#x60;2 CCDs&#x60; - Value -- 2 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;4 CCDs&#x60; - Value -- 4 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;6 CCDs&#x60; - Value -- 6 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;8 CCDs&#x60; - Value -- 8 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;10 CCDs&#x60; - Value -- 10 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;12 CCDs&#x60; - Value -- 12 CCDs for configuring CbsCpuCcdCtrl token. * &#x60;14 CCDs&#x60; - Value -- 14 CCDs for configuring CbsCpuCcdCtrl token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsCpuCcdCtrlC845BiosToken

`func NewBiosCbsCpuCcdCtrlC845BiosToken(classId string, objectType string, ) *BiosCbsCpuCcdCtrlC845BiosToken`

NewBiosCbsCpuCcdCtrlC845BiosToken instantiates a new BiosCbsCpuCcdCtrlC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCpuCcdCtrlC845BiosTokenWithDefaults

`func NewBiosCbsCpuCcdCtrlC845BiosTokenWithDefaults() *BiosCbsCpuCcdCtrlC845BiosToken`

NewBiosCbsCpuCcdCtrlC845BiosTokenWithDefaults instantiates a new BiosCbsCpuCcdCtrlC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCpuCcdCtrlC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


