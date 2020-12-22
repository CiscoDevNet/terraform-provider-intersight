# AdapterFcSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.FcSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.FcSettings"]
**FipEnabled** | Pointer to **bool** | Status of FIP protocol on the adapter interfaces. | [optional] [default to true]

## Methods

### NewAdapterFcSettings

`func NewAdapterFcSettings(classId string, objectType string, ) *AdapterFcSettings`

NewAdapterFcSettings instantiates a new AdapterFcSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterFcSettingsWithDefaults

`func NewAdapterFcSettingsWithDefaults() *AdapterFcSettings`

NewAdapterFcSettingsWithDefaults instantiates a new AdapterFcSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterFcSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterFcSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterFcSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterFcSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterFcSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterFcSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFipEnabled

`func (o *AdapterFcSettings) GetFipEnabled() bool`

GetFipEnabled returns the FipEnabled field if non-nil, zero value otherwise.

### GetFipEnabledOk

`func (o *AdapterFcSettings) GetFipEnabledOk() (*bool, bool)`

GetFipEnabledOk returns a tuple with the FipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipEnabled

`func (o *AdapterFcSettings) SetFipEnabled(v bool)`

SetFipEnabled sets FipEnabled field to given value.

### HasFipEnabled

`func (o *AdapterFcSettings) HasFipEnabled() bool`

HasFipEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


