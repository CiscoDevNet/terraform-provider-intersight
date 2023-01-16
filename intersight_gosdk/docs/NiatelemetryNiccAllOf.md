# NiatelemetryNiccAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Nicc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Nicc"]
**NiccVersion** | Pointer to **string** | NICC version. NiccVersion is used to check compatibility with Nexus Cloud features. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiccAllOf

`func NewNiatelemetryNiccAllOf(classId string, objectType string, ) *NiatelemetryNiccAllOf`

NewNiatelemetryNiccAllOf instantiates a new NiatelemetryNiccAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiccAllOfWithDefaults

`func NewNiatelemetryNiccAllOfWithDefaults() *NiatelemetryNiccAllOf`

NewNiatelemetryNiccAllOfWithDefaults instantiates a new NiatelemetryNiccAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiccAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiccAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiccAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiccAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiccAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiccAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNiccVersion

`func (o *NiatelemetryNiccAllOf) GetNiccVersion() string`

GetNiccVersion returns the NiccVersion field if non-nil, zero value otherwise.

### GetNiccVersionOk

`func (o *NiatelemetryNiccAllOf) GetNiccVersionOk() (*string, bool)`

GetNiccVersionOk returns a tuple with the NiccVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiccVersion

`func (o *NiatelemetryNiccAllOf) SetNiccVersion(v string)`

SetNiccVersion sets NiccVersion field to given value.

### HasNiccVersion

`func (o *NiatelemetryNiccAllOf) HasNiccVersion() bool`

HasNiccVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiccAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiccAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiccAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiccAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


