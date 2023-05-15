# NiatelemetryApicVisionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicVision"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicVision"]
**ApicVisionState** | Pointer to **string** | ApicVision App state. apicVisionState checks the current operational state of ApicVision app on APIC. | [optional] 
**ApicVisionStateLastUpdateTs** | Pointer to **string** | ApicVision App state last updated timestamp. It indicates the last updated timestamp for operational state of ApicVision app. | [optional] 
**ApicVisionVersion** | Pointer to **string** | ApicVision App version. apicVisionVersion is used to check compatibility with Nexus Cloud features. | [optional] 
**ConfigIssues** | Pointer to **string** | Configuration issues depicts the failures for ApicVision managed package upgrade on APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicVisionAllOf

`func NewNiatelemetryApicVisionAllOf(classId string, objectType string, ) *NiatelemetryApicVisionAllOf`

NewNiatelemetryApicVisionAllOf instantiates a new NiatelemetryApicVisionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicVisionAllOfWithDefaults

`func NewNiatelemetryApicVisionAllOfWithDefaults() *NiatelemetryApicVisionAllOf`

NewNiatelemetryApicVisionAllOfWithDefaults instantiates a new NiatelemetryApicVisionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicVisionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicVisionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicVisionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicVisionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicVisionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicVisionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApicVisionState

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionState() string`

GetApicVisionState returns the ApicVisionState field if non-nil, zero value otherwise.

### GetApicVisionStateOk

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionStateOk() (*string, bool)`

GetApicVisionStateOk returns a tuple with the ApicVisionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionState

`func (o *NiatelemetryApicVisionAllOf) SetApicVisionState(v string)`

SetApicVisionState sets ApicVisionState field to given value.

### HasApicVisionState

`func (o *NiatelemetryApicVisionAllOf) HasApicVisionState() bool`

HasApicVisionState returns a boolean if a field has been set.

### GetApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionStateLastUpdateTs() string`

GetApicVisionStateLastUpdateTs returns the ApicVisionStateLastUpdateTs field if non-nil, zero value otherwise.

### GetApicVisionStateLastUpdateTsOk

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionStateLastUpdateTsOk() (*string, bool)`

GetApicVisionStateLastUpdateTsOk returns a tuple with the ApicVisionStateLastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVisionAllOf) SetApicVisionStateLastUpdateTs(v string)`

SetApicVisionStateLastUpdateTs sets ApicVisionStateLastUpdateTs field to given value.

### HasApicVisionStateLastUpdateTs

`func (o *NiatelemetryApicVisionAllOf) HasApicVisionStateLastUpdateTs() bool`

HasApicVisionStateLastUpdateTs returns a boolean if a field has been set.

### GetApicVisionVersion

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionVersion() string`

GetApicVisionVersion returns the ApicVisionVersion field if non-nil, zero value otherwise.

### GetApicVisionVersionOk

`func (o *NiatelemetryApicVisionAllOf) GetApicVisionVersionOk() (*string, bool)`

GetApicVisionVersionOk returns a tuple with the ApicVisionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApicVisionVersion

`func (o *NiatelemetryApicVisionAllOf) SetApicVisionVersion(v string)`

SetApicVisionVersion sets ApicVisionVersion field to given value.

### HasApicVisionVersion

`func (o *NiatelemetryApicVisionAllOf) HasApicVisionVersion() bool`

HasApicVisionVersion returns a boolean if a field has been set.

### GetConfigIssues

`func (o *NiatelemetryApicVisionAllOf) GetConfigIssues() string`

GetConfigIssues returns the ConfigIssues field if non-nil, zero value otherwise.

### GetConfigIssuesOk

`func (o *NiatelemetryApicVisionAllOf) GetConfigIssuesOk() (*string, bool)`

GetConfigIssuesOk returns a tuple with the ConfigIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIssues

`func (o *NiatelemetryApicVisionAllOf) SetConfigIssues(v string)`

SetConfigIssues sets ConfigIssues field to given value.

### HasConfigIssues

`func (o *NiatelemetryApicVisionAllOf) HasConfigIssues() bool`

HasConfigIssues returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicVisionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicVisionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicVisionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicVisionAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


