# FabricMacAgingSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.MacAgingSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.MacAgingSettings"]
**MacAgingOption** | Pointer to **string** | This specifies one of the option to configure the MAC address aging time. * &#x60;Default&#x60; - This option sets the default MAC address aging time to 14500 seconds for End Host mode. * &#x60;Custom&#x60; - This option allows the the user to configure the MAC address aging time on the switch. For Switch Model UCS-FI-6454 or higher, the valid range is 120 to 918000 seconds and the switch will set the lower multiple of 5 of the given time. * &#x60;Never&#x60; - This option disables the MAC address aging process and never allows the MAC address entries to get removed from the table. | [optional] [default to "Default"]
**MacAgingTime** | Pointer to **int32** | Define the MAC address aging time in seconds. This field is valid when the \&quot;Custom\&quot; MAC address aging option is selected. | [optional] [default to 14500]

## Methods

### NewFabricMacAgingSettingsAllOf

`func NewFabricMacAgingSettingsAllOf(classId string, objectType string, ) *FabricMacAgingSettingsAllOf`

NewFabricMacAgingSettingsAllOf instantiates a new FabricMacAgingSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMacAgingSettingsAllOfWithDefaults

`func NewFabricMacAgingSettingsAllOfWithDefaults() *FabricMacAgingSettingsAllOf`

NewFabricMacAgingSettingsAllOfWithDefaults instantiates a new FabricMacAgingSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMacAgingSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMacAgingSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMacAgingSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMacAgingSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMacAgingSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMacAgingSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAgingOption

`func (o *FabricMacAgingSettingsAllOf) GetMacAgingOption() string`

GetMacAgingOption returns the MacAgingOption field if non-nil, zero value otherwise.

### GetMacAgingOptionOk

`func (o *FabricMacAgingSettingsAllOf) GetMacAgingOptionOk() (*string, bool)`

GetMacAgingOptionOk returns a tuple with the MacAgingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAgingOption

`func (o *FabricMacAgingSettingsAllOf) SetMacAgingOption(v string)`

SetMacAgingOption sets MacAgingOption field to given value.

### HasMacAgingOption

`func (o *FabricMacAgingSettingsAllOf) HasMacAgingOption() bool`

HasMacAgingOption returns a boolean if a field has been set.

### GetMacAgingTime

`func (o *FabricMacAgingSettingsAllOf) GetMacAgingTime() int32`

GetMacAgingTime returns the MacAgingTime field if non-nil, zero value otherwise.

### GetMacAgingTimeOk

`func (o *FabricMacAgingSettingsAllOf) GetMacAgingTimeOk() (*int32, bool)`

GetMacAgingTimeOk returns a tuple with the MacAgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAgingTime

`func (o *FabricMacAgingSettingsAllOf) SetMacAgingTime(v int32)`

SetMacAgingTime sets MacAgingTime field to given value.

### HasMacAgingTime

`func (o *FabricMacAgingSettingsAllOf) HasMacAgingTime() bool`

HasMacAgingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


