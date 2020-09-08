# BootDeviceBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies if the boot device is enabled or disabled. | [optional] 
**Name** | Pointer to **string** | A name that helps identify a boot device. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters. | [optional] 

## Methods

### NewBootDeviceBaseAllOf

`func NewBootDeviceBaseAllOf() *BootDeviceBaseAllOf`

NewBootDeviceBaseAllOf instantiates a new BootDeviceBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootDeviceBaseAllOfWithDefaults

`func NewBootDeviceBaseAllOfWithDefaults() *BootDeviceBaseAllOf`

NewBootDeviceBaseAllOfWithDefaults instantiates a new BootDeviceBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BootDeviceBaseAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BootDeviceBaseAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BootDeviceBaseAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BootDeviceBaseAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *BootDeviceBaseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BootDeviceBaseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BootDeviceBaseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BootDeviceBaseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


