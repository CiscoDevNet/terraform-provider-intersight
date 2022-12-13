# NetworkTelemetryCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.TelemetryCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.TelemetryCheck"]
**Status** | Pointer to **string** | Failure status for telemetry configured. | [optional] [readonly] 
**TelemetryConfig** | Pointer to **string** | The telemetry configuration details from endpoint. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkTelemetryCheck

`func NewNetworkTelemetryCheck(classId string, objectType string, ) *NetworkTelemetryCheck`

NewNetworkTelemetryCheck instantiates a new NetworkTelemetryCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTelemetryCheckWithDefaults

`func NewNetworkTelemetryCheckWithDefaults() *NetworkTelemetryCheck`

NewNetworkTelemetryCheckWithDefaults instantiates a new NetworkTelemetryCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkTelemetryCheck) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkTelemetryCheck) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkTelemetryCheck) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkTelemetryCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkTelemetryCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkTelemetryCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *NetworkTelemetryCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkTelemetryCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkTelemetryCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkTelemetryCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTelemetryConfig

`func (o *NetworkTelemetryCheck) GetTelemetryConfig() string`

GetTelemetryConfig returns the TelemetryConfig field if non-nil, zero value otherwise.

### GetTelemetryConfigOk

`func (o *NetworkTelemetryCheck) GetTelemetryConfigOk() (*string, bool)`

GetTelemetryConfigOk returns a tuple with the TelemetryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryConfig

`func (o *NetworkTelemetryCheck) SetTelemetryConfig(v string)`

SetTelemetryConfig sets TelemetryConfig field to given value.

### HasTelemetryConfig

`func (o *NetworkTelemetryCheck) HasTelemetryConfig() bool`

HasTelemetryConfig returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkTelemetryCheck) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkTelemetryCheck) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkTelemetryCheck) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkTelemetryCheck) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


