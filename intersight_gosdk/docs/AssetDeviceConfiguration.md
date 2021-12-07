# AssetDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.DeviceConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.DeviceConfiguration"]
**LocalConfigurationLocked** | Pointer to **bool** | Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled. | [optional] 
**LogLevel** | Pointer to **string** | The log level of the device connector service. | [optional] 
**TunneledKvm** | Pointer to **bool** | Specifies whether Tunneled vKVM is enabled on the device connector. | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAssetDeviceConfiguration

`func NewAssetDeviceConfiguration(classId string, objectType string, ) *AssetDeviceConfiguration`

NewAssetDeviceConfiguration instantiates a new AssetDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceConfigurationWithDefaults

`func NewAssetDeviceConfigurationWithDefaults() *AssetDeviceConfiguration`

NewAssetDeviceConfigurationWithDefaults instantiates a new AssetDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetDeviceConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetDeviceConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetDeviceConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetDeviceConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetDeviceConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetDeviceConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocalConfigurationLocked

`func (o *AssetDeviceConfiguration) GetLocalConfigurationLocked() bool`

GetLocalConfigurationLocked returns the LocalConfigurationLocked field if non-nil, zero value otherwise.

### GetLocalConfigurationLockedOk

`func (o *AssetDeviceConfiguration) GetLocalConfigurationLockedOk() (*bool, bool)`

GetLocalConfigurationLockedOk returns a tuple with the LocalConfigurationLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfigurationLocked

`func (o *AssetDeviceConfiguration) SetLocalConfigurationLocked(v bool)`

SetLocalConfigurationLocked sets LocalConfigurationLocked field to given value.

### HasLocalConfigurationLocked

`func (o *AssetDeviceConfiguration) HasLocalConfigurationLocked() bool`

HasLocalConfigurationLocked returns a boolean if a field has been set.

### GetLogLevel

`func (o *AssetDeviceConfiguration) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *AssetDeviceConfiguration) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *AssetDeviceConfiguration) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *AssetDeviceConfiguration) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetTunneledKvm

`func (o *AssetDeviceConfiguration) GetTunneledKvm() bool`

GetTunneledKvm returns the TunneledKvm field if non-nil, zero value otherwise.

### GetTunneledKvmOk

`func (o *AssetDeviceConfiguration) GetTunneledKvmOk() (*bool, bool)`

GetTunneledKvmOk returns a tuple with the TunneledKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunneledKvm

`func (o *AssetDeviceConfiguration) SetTunneledKvm(v bool)`

SetTunneledKvm sets TunneledKvm field to given value.

### HasTunneledKvm

`func (o *AssetDeviceConfiguration) HasTunneledKvm() bool`

HasTunneledKvm returns a boolean if a field has been set.

### GetDevice

`func (o *AssetDeviceConfiguration) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetDeviceConfiguration) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetDeviceConfiguration) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetDeviceConfiguration) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


