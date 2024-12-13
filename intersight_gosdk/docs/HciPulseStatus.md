# HciPulseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.PulseStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.PulseStatus"]
**Enabled** | Pointer to **bool** | The status of the pulse service. | [optional] [readonly] 
**PiiScrubbingLevel** | Pointer to **string** | PII Scrubbing Level for pulse. It describe at what level the pulse data is collected. | [optional] [readonly] 

## Methods

### NewHciPulseStatus

`func NewHciPulseStatus(classId string, objectType string, ) *HciPulseStatus`

NewHciPulseStatus instantiates a new HciPulseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciPulseStatusWithDefaults

`func NewHciPulseStatusWithDefaults() *HciPulseStatus`

NewHciPulseStatusWithDefaults instantiates a new HciPulseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciPulseStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciPulseStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciPulseStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciPulseStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciPulseStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciPulseStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *HciPulseStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HciPulseStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HciPulseStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HciPulseStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPiiScrubbingLevel

`func (o *HciPulseStatus) GetPiiScrubbingLevel() string`

GetPiiScrubbingLevel returns the PiiScrubbingLevel field if non-nil, zero value otherwise.

### GetPiiScrubbingLevelOk

`func (o *HciPulseStatus) GetPiiScrubbingLevelOk() (*string, bool)`

GetPiiScrubbingLevelOk returns a tuple with the PiiScrubbingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiScrubbingLevel

`func (o *HciPulseStatus) SetPiiScrubbingLevel(v string)`

SetPiiScrubbingLevel sets PiiScrubbingLevel field to given value.

### HasPiiScrubbingLevel

`func (o *HciPulseStatus) HasPiiScrubbingLevel() bool`

HasPiiScrubbingLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


