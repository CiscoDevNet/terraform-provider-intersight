# AssetDeviceConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalConfigurationLocked** | Pointer to **bool** | Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled. | [optional] 
**LogLevel** | Pointer to **string** | The log level of the device connector service. | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeviceConfigurationAllOf

`func NewAssetDeviceConfigurationAllOf() *AssetDeviceConfigurationAllOf`

NewAssetDeviceConfigurationAllOf instantiates a new AssetDeviceConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceConfigurationAllOfWithDefaults

`func NewAssetDeviceConfigurationAllOfWithDefaults() *AssetDeviceConfigurationAllOf`

NewAssetDeviceConfigurationAllOfWithDefaults instantiates a new AssetDeviceConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalConfigurationLocked

`func (o *AssetDeviceConfigurationAllOf) GetLocalConfigurationLocked() bool`

GetLocalConfigurationLocked returns the LocalConfigurationLocked field if non-nil, zero value otherwise.

### GetLocalConfigurationLockedOk

`func (o *AssetDeviceConfigurationAllOf) GetLocalConfigurationLockedOk() (*bool, bool)`

GetLocalConfigurationLockedOk returns a tuple with the LocalConfigurationLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigurationLocked

`func (o *AssetDeviceConfigurationAllOf) SetLocalConfigurationLocked(v bool)`

SetLocalConfigurationLocked sets LocalConfigurationLocked field to given value.

### HasLocalConfigurationLocked

`func (o *AssetDeviceConfigurationAllOf) HasLocalConfigurationLocked() bool`

HasLocalConfigurationLocked returns a boolean if a field has been set.

### GetLogLevel

`func (o *AssetDeviceConfigurationAllOf) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *AssetDeviceConfigurationAllOf) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *AssetDeviceConfigurationAllOf) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *AssetDeviceConfigurationAllOf) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetDevice

`func (o *AssetDeviceConfigurationAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetDeviceConfigurationAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetDeviceConfigurationAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetDeviceConfigurationAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


