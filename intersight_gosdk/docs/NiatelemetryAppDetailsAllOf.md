# NiatelemetryAppDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.AppDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.AppDetails"]
**AppName** | Pointer to **string** | Names of apps running on ND. | [optional] 
**AppStatus** | Pointer to **string** | Status of apps running on ND. | [optional] 
**AppVersion** | Pointer to **string** | Versions of apps running on ND. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryAppDetailsAllOf

`func NewNiatelemetryAppDetailsAllOf(classId string, objectType string, ) *NiatelemetryAppDetailsAllOf`

NewNiatelemetryAppDetailsAllOf instantiates a new NiatelemetryAppDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryAppDetailsAllOfWithDefaults

`func NewNiatelemetryAppDetailsAllOfWithDefaults() *NiatelemetryAppDetailsAllOf`

NewNiatelemetryAppDetailsAllOfWithDefaults instantiates a new NiatelemetryAppDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryAppDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryAppDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryAppDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryAppDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryAppDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryAppDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppName

`func (o *NiatelemetryAppDetailsAllOf) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *NiatelemetryAppDetailsAllOf) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *NiatelemetryAppDetailsAllOf) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *NiatelemetryAppDetailsAllOf) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppStatus

`func (o *NiatelemetryAppDetailsAllOf) GetAppStatus() string`

GetAppStatus returns the AppStatus field if non-nil, zero value otherwise.

### GetAppStatusOk

`func (o *NiatelemetryAppDetailsAllOf) GetAppStatusOk() (*string, bool)`

GetAppStatusOk returns a tuple with the AppStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStatus

`func (o *NiatelemetryAppDetailsAllOf) SetAppStatus(v string)`

SetAppStatus sets AppStatus field to given value.

### HasAppStatus

`func (o *NiatelemetryAppDetailsAllOf) HasAppStatus() bool`

HasAppStatus returns a boolean if a field has been set.

### GetAppVersion

`func (o *NiatelemetryAppDetailsAllOf) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *NiatelemetryAppDetailsAllOf) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *NiatelemetryAppDetailsAllOf) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *NiatelemetryAppDetailsAllOf) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryAppDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryAppDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryAppDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryAppDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


