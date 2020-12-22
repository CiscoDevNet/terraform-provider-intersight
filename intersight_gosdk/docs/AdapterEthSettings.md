# AdapterEthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.EthSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.EthSettings"]
**LldpEnabled** | Pointer to **bool** | Status of LLDP protocol on the adapter interfaces. | [optional] [default to true]

## Methods

### NewAdapterEthSettings

`func NewAdapterEthSettings(classId string, objectType string, ) *AdapterEthSettings`

NewAdapterEthSettings instantiates a new AdapterEthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterEthSettingsWithDefaults

`func NewAdapterEthSettingsWithDefaults() *AdapterEthSettings`

NewAdapterEthSettingsWithDefaults instantiates a new AdapterEthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterEthSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterEthSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterEthSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterEthSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterEthSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterEthSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLldpEnabled

`func (o *AdapterEthSettings) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *AdapterEthSettings) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *AdapterEthSettings) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *AdapterEthSettings) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


