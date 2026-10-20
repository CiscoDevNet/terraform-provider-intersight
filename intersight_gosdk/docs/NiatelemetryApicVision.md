# NiatelemetryApicVision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicVision"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicVision"]
**ApicVisionState** | Pointer to **string** | ApicVision App state. apicVisionState checks the current operational state of ApicVision app on APIC. | [optional] 
**ApicVisionStateLastUpdateTs** | Pointer to **string** | ApicVision App state last updated timestamp. It indicates the last updated timestamp for operational state of ApicVision app. | [optional] 
**ApicVisionVersion** | Pointer to **string** | ApicVision App version. apicVisionVersion is used to check compatibility with Nexus Cloud features. | [optional] 
**ConfigIssues** | Pointer to **string** | Configuration issues depicts the failures for ApicVision managed package upgrade on APIC. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicVision

`func NewNiatelemetryApicVision(classId string, objectType string, ) *NiatelemetryApicVision`

NewNiatelemetryApicVision instantiates a new NiatelemetryApicVision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicVisionWithDefaults

`func NewNiatelemetryApicVisionWithDefaults() *NiatelemetryApicVision`

NewNiatelemetryApicVisionWithDefaults instantiates a new NiatelemetryApicVision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicVision) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicVision) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicVision) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicVision) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicVision) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicVision) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApicVisionState

`func (o *NiatelemetryApicVision) GetApicVisionState() string`

GetApicVisionState returns the ApicVisionState field if non-nil, zero value otherwise.

### GetApicVisionStateOk

`func (o *NiatelemetryApicVision) GetApicVisionStateOk() (*string, bool)`

GetApicVisionStateOk returns a tuple with the ApicVisionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionState

`func (o *NiatelemetryApicVision) SetApicVisionState(v string)`

SetApicVisionState sets ApicVisionState field to given value.

### HasApicVisionState

`func (o *NiatelemetryApicVision) HasApicVisionState() bool`

HasApicVisionState returns a boolean if a field has been set.

### GetApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVision) GetApicVisionStateLastUpdateTs() string`

GetApicVisionStateLastUpdateTs returns the ApicVisionStateLastUpdateTs field if non-nil, zero value otherwise.

### GetApicVisionStateLastUpdateTsOk

`func (o *NiatelemetryApicVision) GetApicVisionStateLastUpdateTsOk() (*string, bool)`

GetApicVisionStateLastUpdateTsOk returns a tuple with the ApicVisionStateLastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVision) SetApicVisionStateLastUpdateTs(v string)`

SetApicVisionStateLastUpdateTs sets ApicVisionStateLastUpdateTs field to given value.

### HasApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVision) HasApicVisionStateLastUpdateTs() bool`

HasApicVisionStateLastUpdateTs returns a boolean if a field has been set.

### GetApicVisionVersion

`func (o *NiatelemetryApicVision) GetApicVisionVersion() string`

GetApicVisionVersion returns the ApicVisionVersion field if non-nil, zero value otherwise.

### GetApicVisionVersionOk

`func (o *NiatelemetryApicVision) GetApicVisionVersionOk() (*string, bool)`

GetApicVisionVersionOk returns a tuple with the ApicVisionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionVersion

`func (o *NiatelemetryApicVision) SetApicVisionVersion(v string)`

SetApicVisionVersion sets ApicVisionVersion field to given value.

### HasApicVisionVersion

`func (o *NiatelemetryApicVision) HasApicVisionVersion() bool`

HasApicVisionVersion returns a boolean if a field has been set.

### GetConfigIssues

`func (o *NiatelemetryApicVision) GetConfigIssues() string`

GetConfigIssues returns the ConfigIssues field if non-nil, zero value otherwise.

### GetConfigIssuesOk

`func (o *NiatelemetryApicVision) GetConfigIssuesOk() (*string, bool)`

GetConfigIssuesOk returns a tuple with the ConfigIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIssues

`func (o *NiatelemetryApicVision) SetConfigIssues(v string)`

SetConfigIssues sets ConfigIssues field to given value.

### HasConfigIssues

`func (o *NiatelemetryApicVision) HasConfigIssues() bool`

HasConfigIssues returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicVision) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicVision) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicVision) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicVision) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryApicVision) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryApicVision) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


