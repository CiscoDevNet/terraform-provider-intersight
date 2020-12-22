# AdapterPortChannelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.PortChannelSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.PortChannelSettings"]
**Enabled** | Pointer to **bool** | When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters. | [optional] [default to true]

## Methods

### NewAdapterPortChannelSettings

`func NewAdapterPortChannelSettings(classId string, objectType string, ) *AdapterPortChannelSettings`

NewAdapterPortChannelSettings instantiates a new AdapterPortChannelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterPortChannelSettingsWithDefaults

`func NewAdapterPortChannelSettingsWithDefaults() *AdapterPortChannelSettings`

NewAdapterPortChannelSettingsWithDefaults instantiates a new AdapterPortChannelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterPortChannelSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterPortChannelSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterPortChannelSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterPortChannelSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterPortChannelSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterPortChannelSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *AdapterPortChannelSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdapterPortChannelSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdapterPortChannelSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdapterPortChannelSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


