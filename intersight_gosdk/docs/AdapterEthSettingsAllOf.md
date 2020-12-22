# AdapterEthSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.EthSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.EthSettings"]
**LldpEnabled** | Pointer to **bool** | Status of LLDP protocol on the adapter interfaces. | [optional] [default to true]

## Methods

### NewAdapterEthSettingsAllOf

`func NewAdapterEthSettingsAllOf(classId string, objectType string, ) *AdapterEthSettingsAllOf`

NewAdapterEthSettingsAllOf instantiates a new AdapterEthSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterEthSettingsAllOfWithDefaults

`func NewAdapterEthSettingsAllOfWithDefaults() *AdapterEthSettingsAllOf`

NewAdapterEthSettingsAllOfWithDefaults instantiates a new AdapterEthSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterEthSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterEthSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterEthSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterEthSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterEthSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterEthSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLldpEnabled

`func (o *AdapterEthSettingsAllOf) GetLldpEnabled() bool`

GetLldpEnabled returns the LldpEnabled field if non-nil, zero value otherwise.

### GetLldpEnabledOk

`func (o *AdapterEthSettingsAllOf) GetLldpEnabledOk() (*bool, bool)`

GetLldpEnabledOk returns a tuple with the LldpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpEnabled

`func (o *AdapterEthSettingsAllOf) SetLldpEnabled(v bool)`

SetLldpEnabled sets LldpEnabled field to given value.

### HasLldpEnabled

`func (o *AdapterEthSettingsAllOf) HasLldpEnabled() bool`

HasLldpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


