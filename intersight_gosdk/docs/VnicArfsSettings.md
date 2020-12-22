# VnicArfsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.ArfsSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.ArfsSettings"]
**Enabled** | Pointer to **bool** | Status of Accelerated Receive Flow Steering on the virtual ethernet interface. | [optional] [default to false]

## Methods

### NewVnicArfsSettings

`func NewVnicArfsSettings(classId string, objectType string, ) *VnicArfsSettings`

NewVnicArfsSettings instantiates a new VnicArfsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicArfsSettingsWithDefaults

`func NewVnicArfsSettingsWithDefaults() *VnicArfsSettings`

NewVnicArfsSettingsWithDefaults instantiates a new VnicArfsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicArfsSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicArfsSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicArfsSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicArfsSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicArfsSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicArfsSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *VnicArfsSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VnicArfsSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VnicArfsSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VnicArfsSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


