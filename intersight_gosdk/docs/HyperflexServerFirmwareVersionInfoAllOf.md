# HyperflexServerFirmwareVersionInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerFirmwareVersionInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerFirmwareVersionInfo"]
**ServerPlatform** | Pointer to **string** | The platform type for UCS server. * &#x60;M5&#x60; - M5 generation of UCS server. * &#x60;M3&#x60; - M3 generation of UCS server. * &#x60;M4&#x60; - M4 generation of UCS server. * &#x60;M6&#x60; - M6 generation of UCS server. | [optional] [default to "M5"]
**Version** | Pointer to **string** | The server firmware bundle version. | [optional] 

## Methods

### NewHyperflexServerFirmwareVersionInfoAllOf

`func NewHyperflexServerFirmwareVersionInfoAllOf(classId string, objectType string, ) *HyperflexServerFirmwareVersionInfoAllOf`

NewHyperflexServerFirmwareVersionInfoAllOf instantiates a new HyperflexServerFirmwareVersionInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionInfoAllOfWithDefaults

`func NewHyperflexServerFirmwareVersionInfoAllOfWithDefaults() *HyperflexServerFirmwareVersionInfoAllOf`

NewHyperflexServerFirmwareVersionInfoAllOfWithDefaults instantiates a new HyperflexServerFirmwareVersionInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerFirmwareVersionInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerFirmwareVersionInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServerPlatform

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetServerPlatform() string`

GetServerPlatform returns the ServerPlatform field if non-nil, zero value otherwise.

### GetServerPlatformOk

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetServerPlatformOk() (*string, bool)`

GetServerPlatformOk returns a tuple with the ServerPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPlatform

`func (o *HyperflexServerFirmwareVersionInfoAllOf) SetServerPlatform(v string)`

SetServerPlatform sets ServerPlatform field to given value.

### HasServerPlatform

`func (o *HyperflexServerFirmwareVersionInfoAllOf) HasServerPlatform() bool`

HasServerPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexServerFirmwareVersionInfoAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexServerFirmwareVersionInfoAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexServerFirmwareVersionInfoAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


