# BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.CbsCmnCpuGlobalCstateCtrlC845BiosToken"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.CbsCmnCpuGlobalCstateCtrlC845BiosToken"]
**Value** | Pointer to **string** | Controls IO based C-state generation and DF C-states. * &#x60;Auto&#x60; - Value -- Auto for configuring CbsCmnCpuGlobalCstateCtrl token. * &#x60;Disabled&#x60; - Value -- Disabled for configuring CbsCmnCpuGlobalCstateCtrl token. * &#x60;Enabled&#x60; - Value -- Enabled for configuring CbsCmnCpuGlobalCstateCtrl token. | [optional] [default to "Auto"]

## Methods

### NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosToken

`func NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosToken(classId string, objectType string, ) *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken`

NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosToken instantiates a new BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosTokenWithDefaults

`func NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosTokenWithDefaults() *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken`

NewBiosCbsCmnCpuGlobalCstateCtrlC845BiosTokenWithDefaults instantiates a new BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetValue

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BiosCbsCmnCpuGlobalCstateCtrlC845BiosToken) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


